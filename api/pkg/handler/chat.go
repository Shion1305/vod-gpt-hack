package handler

import (
	infraDynamo "api/pkg/infra/dynamo"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
)

type SendRequest struct {
	ID        string  `json:"id"`
	StartTime float32 `json:"startTime"`
	EndTime   float32 `json:"endTime"`
	Question  string  `json:"question"`
}

type ChatHandler struct {
	d *infraDynamo.Dynamo
}

func NewChatHandler(d *infraDynamo.Dynamo) *ChatHandler {
	return &ChatHandler{
		d: d,
	}
}

func (ch ChatHandler) Send() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SendRequest
		if err := c.Bind(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// クエリの実行
		resp, err := ch.d.Client.Query(c, &dynamodb.QueryInput{
			IndexName:              aws.String("media-id-id-index"),
			TableName:              aws.String("transcribe"),
			KeyConditionExpression: aws.String("#media_id = :media_id"),
			FilterExpression:       aws.String("start_time > :start_time AND end_time < :end_time"),
			ExpressionAttributeNames: map[string]string{
				"#media_id": "media-id",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":media_id":   &types.AttributeValueMemberS{Value: req.ID},
				":start_time": &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", req.StartTime)},
				":end_time":   &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", req.EndTime)},
			},
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
