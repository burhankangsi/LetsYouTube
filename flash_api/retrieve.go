package flash_api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/burhankangsi/LetsYouTube/content"
	"github.com/cheggaaa/pb"
	log "github.com/sirupsen/logrus"
)

type awsCreds struct {
	accessKey  string
	secretKey  string
	region     string
	token      string
	bucket     string
	AWS_REGION string
}

type progressWriter struct {
	writer  io.WriterAt
	size    int64
	bar     *pb.ProgressBar
	display bool
}

func GetS3ObjectSize(bucket, item string) int64 {
	var curr_credentials awsCreds
	curr_credentials.bucket = "youtube-clone-bk"
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(curr_credentials.AWS_REGION), Credentials: credentials.AnonymousCredentials},
	)

	svc := s3.New(sess)
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	}

	result, err := svc.HeadObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Errorf("Error getting size of file", aerr)
		} else {
			fmt.Errorf("Error getting size of file", err)
		}
	}
	return *result.ContentLength
}

func (pw *progressWriter) init(s3ObjectSize int64) {
	if pw.display {
		pw.bar = pb.StartNew(int(s3ObjectSize))
		pw.bar.ShowSpeed = true
		pw.bar.Format("[=>_]")
		pw.bar.SetUnits(pb.U_BYTES_DEC)
	}
}

func DownloadFromS3Bucket(bucket, path, item string) error {

	// os.Setenv("AWS_ACCESS_KEY", "my-key")
	// os.Setenv("AWS_SECRET_KEY", "my-secret")

	file, err := os.Create(filepath.Join(path, item))
	if err != nil {
		fmt.Printf("Error in creating file path: %v \n", err)
		os.Exit(1)
	}
	defer file.Close()

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	// Get the object size
	s3ObjectSize := GetS3ObjectSize(bucket, item)
	downloader := s3manager.NewDownloader(sess, func(d *s3manager.Downloader) {
		d.PartSize = 64 * 1024 * 1024 // 64MB per part
		d.Concurrency = 6
	})
	var displayProgressBar bool
	if s3ObjectSize > 64 {
		displayProgressBar = true
	}
	// Initialize progress writer
	writer := &progressWriter{writer: file, size: s3ObjectSize}
	writer.display = displayProgressBar
	writer.init(s3ObjectSize)

	// Start the download
	numBytes, err1 := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(path + item),
		})
	if err1 != nil {
		log.Errorf("Error while downloading the file from bucket %v, Error is %v", item, err1)
		return err1
	}
	writer.finish()
	fmt.Println("Download complete", file.Name(), numBytes, "bytes")
	return nil
}

func GetVideoObject(w http.ResponseWriter, r *http.Request, videoId string, channelId string) (content.File, error) {
	var file content.File
	var err error
	file, err = fetchFile(videoId, channelId)
	if err != nil {
		log.Errorf("Got an error while fetching the file. Error is %v", err)
		return file, err
	}
	// vid := videoId + ".ts"
	vid := videoId + ".mp4"
	http.ServeFile(w, r, vid)
	os.Remove(vid)
	return file, nil
}

func fetchFile(vid string, chanId string) (content.File, error) {
	awsCred := awsCreds{}
	svc := s3.New(session.New(), &aws.Config{
		Region: aws.String("us-east-1"),
	})

	awsCred.bucket = "youtube-clone-bk"
	//context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	params := &s3.ListObjectsInput{
		Bucket: aws.String(awsCred.bucket),
		Prefix: aws.String(chanId + "/" + "video" + "/"),
	}
	var num string
	var count int

	var file content.File
	resp, err1 := svc.ListObjects(params)
	if err1 != nil {
		log.Infof("Failed to list s3 objects")
		return file, err1
	}
	for _, obj := range resp.Contents {
		if strings.Contains(*obj.Key, vid) {
			log.Infof("File found in S3. Key: %v", *obj.Key)
			num = *obj.Key
			count++
		}
	}

	if num == "" {
		log.Info("Video does not exist")
		return file, nil
	}
	path := "s3a://" + chanId + "/" + "video" + "/"
	//item := vid + ".ts"
	item := vid + ".mp4"
	if count == 1 {
		log.Infof("Video exists but json file doesn't")
	} else {
		var err3 error
		file, err3 = DownloadJsonFromS3(awsCred.bucket, path, item)
		if err3 != nil {
			log.Infof("Error occured while downloading json file, %v", err3)
			return file, err3
		}
	}

	err2 := DownloadFromS3Bucket(awsCred.bucket, path, item)
	if err2 != nil {
		log.Errorf("Could not initiate download video from S3. Error is %v", err2)
		return file, err2
	}
	return file, nil
}

func (pw *progressWriter) WriteAt(p []byte, off int64) (int, error) {
	if pw.display {
		pw.bar.Add64(int64(len(p)))
	}
	return pw.writer.WriteAt(p, off)
}

func (pw *progressWriter) finish() {
	if pw.display {
		pw.bar.Finish()
	}
}
func DownloadJsonFromS3(bucket string, path string, item string) (content.File, error) {

	sess, err1 := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err1 != nil {
		log.Infof("Could not create an aws session. %v", err1)
	}
	// 3) Create a new AWS S3 downloader
	downloader := s3manager.NewDownloader(sess)

	// 4) Download the item from the bucket. If an error occurs, log it and exit. Otherwise, notify the user that the download succeeded.
	file, err := os.Create(item)
	defer file.Close()

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(path + item),
		})
	var outputFile content.File
	if err != nil {
		log.Fatalf("Unable to download item %q, %v", item, err)
		return outputFile, err
	}

	fmt.Println("Downloaded json", file.Name(), numBytes, "bytes")
	//take data, and put in struct
	byteValue, err3 := ioutil.ReadAll(file)
	if err3 != nil {
		log.Errorf("Could not unmarshal json file. error is %v", err3)
	}

	json.Unmarshal(byteValue, &outputFile)
	return outputFile, nil
}
