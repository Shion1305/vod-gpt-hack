package domain

type BRRequest struct {
	AnthropicVersion string      `json:"anthropic_version"`
	MaxTokens        int         `json:"max_tokens"`
	System           string      `json:"system"`
	Messages         []BRMessage `json:"messages"`
	Temperature      float64     `json:"temperature"`
	TopP             float64     `json:"top_p"`
}

type BRMessage struct {
	Role    string      `json:"role"`
	Content []BRContent `json:"content"`
}

type BRContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type BRResponse struct {
	ID           string          `json:"id"`
	Model        string          `json:"model"`
	Type         string          `json:"type"`
	Role         string          `json:"role"`
	ContentItem  []BRContentItem `json:"content"`
	StopReason   string          `json:"stop_reason,omitempty"`
	StopSequence string          `json:"stop_sequence,omitempty"`
	Usage        BRUsageDetails  `json:"usage"`
}

type BRStreamResponse struct {
	Type  string `json:"type"`
	Index int    `json:"index"`
	Delta struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"delta"`
}

type BRContentItem struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type BRUsageDetails struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}
