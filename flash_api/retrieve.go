package flash_api

import (
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"fmt"
	"os"
)

type awsCreds struct {
	accessKey		string
	secretKey		string
	region			string
	token 			string
	bucket			string	
}

func GetS3ObjectSize(bucket, item string) int64 {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(constants.AWS_REGION), Credentials: credentials.AnonymousCredentials},
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
}

func GetVideoObject(w http.ResponseWriter, r *http.Request, videoId string, channelId string) {
	err := fetchFile(videoId, channelId)
	if err != nil {
		log.Errorf("Got an error while fetching the file. Error is %v", err)
		return
	}
	vid := videoId + ".ts"
	http.ServeFile(w, r, vid)
	os.Remove(vid)
}

func fetchFile(vid string, chanId string) error {
	aws := awsCreds{}
	svc := s3.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	//context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	params := &s3.ListObjectsInput {
		Bucket: aws.String(aws.Bucket),
		Prefix: aws.String(chanId),
	}
	var num string
	resp, err1 := svc.ListObjects(params)
	if err1 != nil {
		log.Infof("Failed to list s3 objects")
		return err1
	}
	for _, key := range resp.Contents {
		if strings.Contains(*key.Key, vid)
	}
	log.Infof("File found in S3. Key: %v", *key.Key)

	if num == "" {
		return fmt.Errorf("File %v does not exist in S3", vid)
	}
	path := chanId + "/" + "video"
	item := vid + ".ts"
	err := DownloadFromS3Bucket(aws.Bucket string, path string, item string)
	if err != nil {
		log.Errorf("Could not download video from S3. Error is %v", err)
		return err
	}
	return nil
}