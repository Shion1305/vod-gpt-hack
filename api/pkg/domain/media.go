package domain

const (
	Completed  string = "COMPLETED"
	Failed     string = "FAILED"
	InProgress string = "IN_PROGRESS"
)

type Media struct {
	// NOTE: 不足してるかも。
	ID     string `dynamo:"id"`
	UserID string `dynamo:"userId"`
	Status string `dynamo:"status"`
}
