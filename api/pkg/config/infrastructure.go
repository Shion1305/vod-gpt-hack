package config

type Infrastructure struct {
	S3       S3       `yaml:"s3"`
	DynamoDB DynamoDB `yaml:"dynamodb"`
	SQS      SQS      `yaml:"sqs"`
}

type S3 struct {
	Region string `yaml:"region"`
	Bucket string `yaml:"bucket"`
}

type DynamoDB struct {
	Region string `yaml:"region"`
}

type SQS struct {
	Region string `yaml:"sqs"`
	URL    string `yaml:"url"`
}
