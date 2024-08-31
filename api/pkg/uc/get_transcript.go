package uc

import (
	infraDynamo "api/pkg/infra/dynamo"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type GetTranscript struct {
	d *infraDynamo.Dynamo
}

func NewGetTranscript(d *infraDynamo.Dynamo) *GetTranscript {
	return &GetTranscript{
		d: d,
	}
}

func (uc *GetTranscript) Execute(c context.Context,
	ID string, StartTime float32, EndTime float32,
) (*dynamodb.QueryOutput, error) {
	resp, err := uc.d.Client.Query(c, &dynamodb.QueryInput{
		IndexName:              aws.String("media-id-id-index"),
		TableName:              aws.String("transcribe"),
		KeyConditionExpression: aws.String("#media_id = :media_id"),
		FilterExpression:       aws.String("start_time > :start_time AND end_time < :end_time"),
		ExpressionAttributeNames: map[string]string{
			"#media_id": "media-id",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":media_id":   &types.AttributeValueMemberS{Value: ID},
			":start_time": &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", StartTime)},
			":end_time":   &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", EndTime)},
		},
	})
	return resp, err
}
