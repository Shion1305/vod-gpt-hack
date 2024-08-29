package handler

import (
	"api/pkg/config"
	"api/pkg/domain"
	infraDynamo "api/pkg/infra/dynamo"
	infraS3 "api/pkg/infra/s3"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MediaHandler struct {
	s *infraS3.S3
	d *infraDynamo.Dynamo
}

func NewMediaHandler(s *infraS3.S3, d *infraDynamo.Dynamo) MediaHandler {
	return MediaHandler{
		s,
		d,
	}
}

func (m MediaHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New().String()
		now := time.Now()
		media := domain.Media{
			ID:        id,
			Status:    domain.InProgress,
			CreatedAt: now,
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

		input := &s3.PutObjectInput{
			Bucket:      aws.String(conf.Infrastructure.S3.Bucket),
			Key:         aws.String(fmt.Sprintf("%s/%s.mp4", id, uuid.New().String())),
			Body:        file,
			ContentType: aws.String(contentType),
		}

		_, err = m.s.Client.PutObject(c, input)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "upload mp4 file successfully.",
		})
	}
}
