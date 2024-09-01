package uc

import (
	"api/pkg/domain"
	infraDynamo "api/pkg/infra/dynamo"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"sort"
	"strconv"
)

type GetTranscript struct {
	d *infraDynamo.Dynamo
}

func NewGetTranscript(d *infraDynamo.Dynamo) *GetTranscript {
	return &GetTranscript{
		d: d,
	}
}

func (uc *GetTranscript) Execute(
	c context.Context,
	ID string,
	StartTime float32,
	EndTime float32,
) ([]domain.TranscriptionPartial, error) {
	resp, err := uc.exec(c, ID, StartTime, EndTime)
	if err != nil {
		return nil, err
	}
	partials := make([]domain.TranscriptionPartial, len(resp.Items))
	for _, item := range resp.Items {
		content := item["content"].(*types.AttributeValueMemberS).Value
		idS := item["id"].(*types.AttributeValueMemberN).Value
		id, err := strconv.Atoi(idS)
		if err != nil {
			fmt.Println("UNEXPECTED ERROR: ", err)
			continue
		}
		startTimeS := item["start_time"].(*types.AttributeValueMemberN).Value
		startTime, err := strconv.ParseFloat(startTimeS, 64)
		if err != nil {
			fmt.Println("UNEXPECTED ERROR: ", err)
			continue
		}
		endTimeS := item["end_time"].(*types.AttributeValueMemberN).Value
		endTime, err := strconv.ParseFloat(endTimeS, 64)
		if err != nil {
			fmt.Println("UNEXPECTED ERROR: ", err)
			continue
		}
		fmt.Printf("item: %d %s %f %f\n", id, content, startTime, endTime)
		partials = append(partials, domain.TranscriptionPartial{
			ID:      id,
			Content: content,
			Start:   float32(startTime),
			End:     float32(endTime),
		})
	}
	sort.Slice(partials, func(i, j int) bool {
		return partials[i].Start < partials[j].Start
	})
	return partials, nil
}

func (uc *GetTranscript) ExecuteGrouped(
	c context.Context,
	ID string,
	StartTime float32,
	EndTime float32,
) (*domain.TranscriptionGrouped, error) {
	resp, err := uc.Execute(c, ID, StartTime, EndTime)
	if err != nil {
		return nil, err
	}
	content := ""
	for _, item := range resp {
		content += item.Content
	}
	return &domain.TranscriptionGrouped{
		Content: content,
		Start:   resp[0].Start,
		End:     resp[len(resp)-1].End,
	}, nil
}

func (uc *GetTranscript) exec(c context.Context,
	ID string, StartTime float32, EndTime float32,
) (*dynamodb.QueryOutput, error) {
	resp, err := uc.d.Client.Query(c, &dynamodb.QueryInput{
		IndexName:              aws.String("media_id-id-index"),
		TableName:              aws.String("transcribe"),
		KeyConditionExpression: aws.String("#media_id = :media_id"),
		FilterExpression:       aws.String("start_time > :start_time AND end_time < :end_time"),
		ExpressionAttributeNames: map[string]string{
			"#media_id": "media_id",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":media_id":   &types.AttributeValueMemberS{Value: ID},
			":start_time": &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", StartTime)},
			":end_time":   &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", EndTime)},
		},
	})
	return resp, err
}
