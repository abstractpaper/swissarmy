package s3

import (
	"bytes"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Upload uploads a slice of bytes to an S3 bucket.
func Upload(session *session.Session, body []byte, bucket string, key string) (location string, err error) {
	uploader := s3manager.NewUploader(session)
	// upload the file to S3
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
	})
	if err != nil {
		log.Errorln("Failed to upload file: ", key)
	}

	log.Info("Uploaded ", result.Location)
	location = result.Location

	return
}
