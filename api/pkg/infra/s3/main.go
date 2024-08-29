package s3

import (
	"api/pkg/config"
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	Client *s3.Client
}

func NewS3() (*S3, error) {
	conf := config.Get()
	ctx := context.Background()

	cfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithRegion(conf.Infrastructure.S3.Region))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &S3{
		Client: client,
	}, nil
}
