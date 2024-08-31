package sqs

import (
	"api/pkg/config"
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQS struct {
	Client *sqs.Client
}

func NewSQS() (*SQS, error) {
	conf := config.Get()
	ctx := context.Background()

	cfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithRegion(conf.Infrastructure.SQS.Region))
	if err != nil {
		return nil, err
	}

	client := sqs.NewFromConfig(cfg)

	return &SQS{
		Client: client,
	}, nil
}
