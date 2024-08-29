package dynamo

import (
	"api/pkg/config"
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Dynamo struct {
	Client *dynamodb.Client
}

func NewDynamo() (*Dynamo, error) {
	conf := config.Get()
	ctx := context.Background()

	cfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithRegion(conf.Infrastructure.DynamoDB.Region))
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	return &Dynamo{
		Client: client,
	}, nil
}
