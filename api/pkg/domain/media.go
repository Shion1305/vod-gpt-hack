package domain

import "time"

const (
	Completed  string = "COMPLETED"
	Failed     string = "FAILED"
	InProgress string = "IN_PROGRESS"
)

type Media struct {
	// NOTE: 不足してるかも。
	ID        string
	Status    string
	CreatedAt time.Time
}
