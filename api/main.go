package main

import (
	handler "api/pkg/handler"
	"api/pkg/infra/dynamo"
	"api/pkg/infra/s3"
	"api/pkg/infra/sqs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	s, err := s3.NewS3()
	if err != nil {
		log.Fatalf("failed to connect to s3, err: %v", err)
	}
	d, err := dynamo.NewDynamo()
	if err != nil {
		log.Fatalf("failed to connect to dynamo, err: %v", err)
	}
	sq, err := sqs.NewSQS()
	if err != nil {
		log.Fatalf("failed to connect to sqs, err: %v", err)
	}

	engine := gin.Default()

	// Health Check
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "api is running",
		})
	})

	apiV1 := engine.Group("/api/v1")

	if err := implement(apiV1, s, d, sq); err != nil {
		log.Fatalf("failed to start server... %v", err)
		return
	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("failed to start api... %v", err)
		return
	}
}

func implement(g *gin.RouterGroup, s *s3.S3, d *dynamo.Dynamo, sq *sqs.SQS) error {
	mediaHandler := handler.NewMediaHandler(s, d, sq)
	chatHandler := handler.NewChatHandler(d)
	g.Handle("POST", "/media/upload/:id", mediaHandler.UploadMP4())
	g.Handle("GET", "/media/id", mediaHandler.Create())
	g.Handle("POST", "/chat", chatHandler.Start())
	g.Handle("GET", "/chat/:id", chatHandler.SendDummy())

	return nil
}
