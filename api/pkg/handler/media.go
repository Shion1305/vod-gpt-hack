package handler

import (
	"api/pkg/config"
	"api/pkg/domain"
	infraDynamo "api/pkg/infra/dynamo"
	infraS3 "api/pkg/infra/s3"
	infraSQS "api/pkg/infra/sqs"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MediaHandler struct {
	s  *infraS3.S3
	d  *infraDynamo.Dynamo
	sq *infraSQS.SQS
}

type SQSBody struct {
	ID       string `json:"id"`
	FileName string `json:"fileName"`
}

func NewMediaHandler(s *infraS3.S3, d *infraDynamo.Dynamo, sq *infraSQS.SQS) MediaHandler {
	return MediaHandler{
		s,
		d,
		sq,
	}
}

func (m MediaHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userId")
		id := uuid.New().String()
		media := domain.Media{
			ID:     id,
			UserID: userID,
			Status: domain.InProgress,
		}
		item, err := attributevalue.MarshalMap(media)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		_, err = m.d.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
			TableName: aws.String("media"), Item: item,
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"id": id,
		})
	}
}

func (m MediaHandler) UploadMP4() gin.HandlerFunc {
	conf := config.Get()
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Missing 'id' url parameter.",
			})
			return
		}

		file, fileHeader, err := c.Request.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		contentType := fileHeader.Header.Get("Content-Type")
		if contentType != "video/mp4" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Missing Content-Type Header.",
			})
			return
		}

		fileName := uuid.New().String()
		input := &s3.PutObjectInput{
			Bucket:      aws.String(conf.Infrastructure.S3.Bucket),
			Key:         aws.String(fmt.Sprintf("%s/%s.mp4", id, fileName)),
			Body:        file,
			ContentType: aws.String(contentType),
		}

		_, err = m.s.Client.PutObject(c, input)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		messageBody := SQSBody{
			ID:       id,
			FileName: fileName,
		}

		jsonData, err := json.Marshal(messageBody)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		sqsParams := &sqs.SendMessageInput{
			MessageBody: aws.String(string(jsonData)),
			QueueUrl:    aws.String(conf.Infrastructure.SQS.URL),
		}

		_, err = m.sq.Client.SendMessage(c, sqsParams)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "upload mp4 file successfully.",
		})
	}
}
