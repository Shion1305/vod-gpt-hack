package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Dynamo struct {
	Client *dynamodb.Client
}

func NewDynamo() (*Dynamo, error) {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"))
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	return &Dynamo{
		Client: client,
	}, nil
}
