package bucket_api

import (
	"bytes"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func saveToBucket() {

	ret := retrieve{}
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(ret.pipeReader)
	res := buffer.Bytes()

	// Upload Files
	err = uploadFile(session, res)
	if err != nil {
		log.Fatal(err)
	}
}

func createAWSSession() {
	session, err := session.NewSession(&aws.Config{Region: aws.String("ap-south-1")})
	if err != nil {
		log.Fatal(err)
	}
}

func uploadFile(session *session.Session, video []byte) {

	upFile, err := os.Open(video)
	if err != nil {
		return err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	// _, err = s3.New(session).PutObject(&s3.PutObjectInput{
	// 	Bucket:               aws.String(AWS_S3_BUCKET),
	// 	Key:                  aws.String(video),
	// 	ACL:                  aws.String("private"),
	// 	Body:                 bytes.NewReader(fileBuffer),
	// 	ContentLength:        aws.Int64(fileSize),
	// 	ContentType:          aws.String(http.DetectContentType(fileBuffer)),
	// 	ContentDisposition:   aws.String("attachment"),
	// 	ServerSideEncryption: aws.String("AES256"),
	// })
	// return err

	uploader := s3manager.NewUploader(session)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWS_S3_BUCKET), // Bucket to be used
		Key:    aws.String(video),         // Name of the file to be saved
		Body:   upFile,                    // File
	})
	if err != nil {
		// Do your error handling here
		return err
	}
}
