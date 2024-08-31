package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

var config *Config

type Config struct {
	Infrastructure Infrastructure `yaml:"infrastructure"`
}

func init() {
	var envLocation string
	_ = godotenv.Load(".env")
	envLocation = os.Getenv("ENV_LOCATION")
	if envLocation == "" {
		log.Printf("failed to load .env file, ENV_LOCATION is not set")
	} else {
		reader, err := os.Open(envLocation)
		if err != nil {
			dir, _ := os.Getwd()
			log.Printf("failed to open setting file: %v, %v\n", dir, err)
		}
		decoder := yaml.NewDecoder(reader)
		config = &Config{}
		if err = decoder.Decode(config); err != nil {
			log.Printf("failed to decode setting file: %v\n", err)
		}
	}
	if sqsUrl := os.Getenv("SQS_URL"); sqsUrl != "" {
		config.Infrastructure.SQS.URL = sqsUrl
	}
	if s3Bucket := os.Getenv("S3_BUCKET"); s3Bucket != "" {
		config.Infrastructure.S3.Bucket = s3Bucket
	}
	if region := os.Getenv("REGION"); region != "" {
		config.Infrastructure.SQS.Region = region
		config.Infrastructure.S3.Region = region
		config.Infrastructure.DynamoDB.Region = region
	}
}

func Get() Config {
	if config == nil {
		panic("setting is nil")
	}
	return *config
}
