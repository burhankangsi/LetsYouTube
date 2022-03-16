package flash_api

import (
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"fmt"
	"os"
)

func DownloadFromS3Bucket(bucket, item, path string) {

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
		fmt.Printf("Error in downloading from file: %v \n", err)
		os.Exit(1)
	}
	defer file.Close()

	sess, _ := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	//For anonymous credentials. Use if you don't have access key or secret key
	// sess, _ := session.NewSession(&aws.Config{
	//     Region: aws.String(constants.AWS_REGION), Credentials: credentials.AnonymousCredentials},
	// )

	downloader := s3manager.NewDownloader(sess, func(d *s3manager.Downloader) {
		d.PartSize = 64 * 1024 * 1024 // 64MB per part
		d.Concurrency = 6
	})
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func GetVideoObject(W http.ResponseWriter, R *http.Request, videoId string, channelId string) {
	ts := videoId + ".ts"
}
