package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Object(bucket string, objectName string) (*s3.GetObjectOutput, error) {
	newSession, err := session.NewSession(&aws.Config{})

	if err != nil {
		return nil, fmt.Errorf("could not initialize new aws session: %v", err)
	}

	s3Client := s3.New(newSession)
	object, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectName),
	})

	if err != nil {
		return nil, fmt.Errorf("could not retreive s3 object: %v", err)
	}

	return object, nil
}
