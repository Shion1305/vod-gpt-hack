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
		log.Fatalln("failed to load .env file, ENV_LOCATION is not set")
	}
	reader, err := os.Open(envLocation)
	if err != nil {
		dir, _ := os.Getwd()
		log.Fatalf("failed to open setting file: %v, %v\n", dir, err)
	}
	decoder := yaml.NewDecoder(reader)
	config = &Config{}
	if err = decoder.Decode(config); err != nil {
		log.Fatalf("failed to decode setting file: %v\n", err)
	}
	if sqsUrl := os.Getenv("SQS_URL"); sqsUrl != "" {
		config.Infrastructure.SQS.URL = sqsUrl
	}
}

func Get() Config {
	if config == nil {
		panic("setting is nil")
	}
	return *config
}
