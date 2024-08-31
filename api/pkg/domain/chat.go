package domain

import "github.com/google/uuid"

type ChatRequest struct {
	Question string
	From     float32
	To       float32
	VID      uuid.UUID
}
