package main

import (
	handler "api/pkg/handler"
	"api/pkg/infra/dynamo"
	"api/pkg/infra/s3"
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

	engine := gin.Default()

	// Health Check
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "api is running",
		})
	})

	apiV1 := engine.Group("/api/v1")

	if err := implement(apiV1, s, d); err != nil {
		log.Fatalf("failed to start server... %v", err)
		return
	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("failed to start api... %v", err)
		return
	}
}

func implement(g *gin.RouterGroup, s *s3.S3, d *dynamo.Dynamo) error {
	mediaHandler := handler.NewMediaHandler(s, d)
	g.Handle("POST", "/media/upload/:id", mediaHandler.UploadMP4())
	g.Handle("GET", "/media/id", mediaHandler.Create())

	return nil
}
