package schema

type ChatRequest struct {
	Question string  `json:"question"`
	From     float32 `json:"from"`
	To       float32 `json:"to"`
	VID      string  `json:"vid"`
}

type SummaryRequest struct {
	VID  string  `json:"vid"`
	From float32 `json:"from"`
	To   float32 `json:"to"`
}

type ListRequest struct {
	UserID string `json:"userId"`
}

type ListResponse struct {
	Videos []ListEntry `json:"videos"`
}

type ListEntry struct {
	VID    string `json:"vid"`
	S3     string `json:"s3"`
	Status string `json:"status"`
	Title  string `json:"title"`
}
