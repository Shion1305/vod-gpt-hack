package main

import (
	"log"
	"net/http"
	"server/pkg/infra/dynamo"
	"server/pkg/infra/s3"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := s3.NewS3()
	if err != nil {
		log.Fatalf("failed to connect to s3, err: %v", err)
	}
	_, err = dynamo.NewDynamo()
	if err != nil {
		log.Fatalf("failed to connect to dynamo, err: %v", err)
	}

	engine := gin.Default()

	// Health Check
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "server is running",
		})
	})

	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("failed to start server... %v", err)
		return
	}
}
