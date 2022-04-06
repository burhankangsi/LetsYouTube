package flash_api

import (
	"fmt"
	"io"
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

	os.Setenv("AWS_ACCESS_KEY", "my-key")
	os.Setenv("AWS_SECRET_KEY", "my-secret")

	// bucket := "cellery-runtime-installation"
	// item := "hello-world.txt"

	// file, err := os.Create(item)
	// if err != nil {
	//     fmt.Println(err)
	//}

	file, err := os.Create(filepath.Join(path, item))
	if err != nil {
		fmt.Printf("Error in creating file path: %v \n", err)
		os.Exit(1)
	}
	defer file.Close()

	sess, _ := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	//For anonymous credentials. Use if you don't have access key or secret key
	// sess, _ := session.NewSession(&aws.Config{
	//     Region: aws.String(constants.AWS_REGION), Credentials: credentials.AnonymousCredentials},
	// )

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
			Key:    aws.String(item),
		})
	if err1 != nil {
		log.Errorf("Error while downloading the file %v, Error is %v", item, err1)
		return err1
	}
	writer.finish()
	fmt.Println("Download complete", file.Name(), numBytes, "bytes")
	return nil
}

func GetVideoObject(w http.ResponseWriter, r *http.Request, videoId string, channelId string) (json, err) {
	err := fetchFile(videoId, channelId)
	if err != nil {
		log.Errorf("Got an error while fetching the file. Error is %v", err)
		return err
	}
	vid := videoId + ".ts"
	http.ServeFile(w, r, vid)
	os.Remove(vid)
}

func fetchFile(vid string, chanId string) error {
	awsCred := awsCreds{}
	svc := s3.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	//context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	params := &s3.ListObjectsInput{
		Bucket: aws.String(awsCred.bucket),
		Prefix: aws.String(chanId),
	}
	var num string
	resp, err1 := svc.ListObjects(params)
	if err1 != nil {
		log.Infof("Failed to list s3 objects")
		return err1
	}
	for _, key := range resp.Contents {
		if strings.Contains(*key.Key, vid) {
			log.Infof("File found in S3. Key: %v", *key.Key)
		}
	}

	if num == "" {
		return fmt.Errorf("File %v does not exist in S3", vid)
	}
	path := chanId + "/" + "video"
	item := vid + ".ts"
	err := DownloadFromS3Bucket(awsCred.bucket, path, item)
	if err != nil {
		log.Errorf("Could not download video from S3. Error is %v", err)
		return err
	}
	return nil
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
