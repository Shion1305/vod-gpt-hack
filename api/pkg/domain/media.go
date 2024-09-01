package domain

const (
	Completed  string = "COMPLETED"
	Failed     string = "FAILED"
	InProgress string = "IN_PROGRESS"
)

type Media struct {
	// NOTE: 不足してるかも。
	ID     string `dynamodbav:"id"`
	UserID string `dynamodbav:"userId"`
	Status string `dynamodbav:"status"`
}
