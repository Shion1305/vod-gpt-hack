package handler

import (
	infraDynamo "api/pkg/infra/dynamo"
	"api/pkg/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
)

type ListHandler struct {
	d *infraDynamo.Dynamo
}

func NewListHandler(
	d *infraDynamo.Dynamo,
) *ListHandler {
	return &ListHandler{
		d: d,
	}
}

func (h *ListHandler) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req schema.ListRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.UserID == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "missing user id"})
			return
		}
		input := &dynamodb.QueryInput{
			IndexName:              aws.String("userId-s3-index"),
			TableName:              aws.String("media"),
			KeyConditionExpression: aws.String("userId = :userId"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":userId": &types.AttributeValueMemberS{Value: req.UserID},
			},
		}
		resp, err := h.d.Client.Query(c, input)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		vList := make([]schema.ListEntry, 0, len(resp.Items))

		for _, item := range resp.Items {
			id := item["id"].(*types.AttributeValueMemberS).Value
			s3 := item["s3"].(*types.AttributeValueMemberS).Value
			status := item["status"].(*types.AttributeValueMemberS).Value
			title := item["title"].(*types.AttributeValueMemberS).Value
			vList = append(vList, schema.ListEntry{
				VID:    id,
				S3:     s3,
				Status: status,
				Title:  title,
			})
		}
		c.JSON(200, schema.ListResponse{
			Videos: vList,
		})
	}
}
