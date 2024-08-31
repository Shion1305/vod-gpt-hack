package schema

type ChatRequest struct {
	Question string  `json:"question"`
	From     float32 `json:"from"`
	To       float32 `json:"to"`
	VID      string  `json:"vid"`
}
