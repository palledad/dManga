package services

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	bucketName = string("dev-dmanga-resource")
)

type ImageLinksService struct {
	s3PresignClient *s3.PresignClient
}

func NewImageLinksService(s3Client *s3.Client) *ImageLinksService {
	return &ImageLinksService{
		s3PresignClient: s3.NewPresignClient(s3Client),
	}
}

func (s *ImageLinksService) GetLink(c context.Context, objectKey string) (string, error) {
	linkExpire := time.Now().Add(time.Minute * 5)
	presignedPutRequest, err := s.s3PresignClient.PresignPutObject(c, &s3.PutObjectInput{
		Bucket:  &bucketName,
		Key:     &objectKey,
		Expires: &linkExpire,
	})
	if err != nil {
		return "", err
	}
	return presignedPutRequest.URL, nil
}

func String(s string) {
	panic("unimplemented")
}
