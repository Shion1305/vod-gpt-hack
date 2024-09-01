package service

import (
	"api/pkg/domain"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

type BedrockService struct {
	brc *bedrockruntime.Client
}

func NewBedrockService() (*BedrockService, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"))
	if err != nil {
		return nil, fmt.Errorf("error Config Load: %v", err)
	}
	brc := bedrockruntime.NewFromConfig(cfg)
	return &BedrockService{brc}, nil
}

func (b BedrockService) ClaudeMessageCompletion(userMessage string, systemPrompt string) (string, error) {
	content := domain.BRContent{Type: "text", Text: userMessage}
	msg := domain.BRMessage{Role: "user", Content: []domain.BRContent{content}}
	payload := domain.BRRequest{
		AnthropicVersion: "bedrock-2023-05-31",
		MaxTokens:        512,
		System:           systemPrompt,
		Messages:         []domain.BRMessage{msg},
		Temperature:      0.5,
		TopP:             0.9,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshalling payload: %v", err)
	}

	output, err := b.brc.InvokeModel(context.Background(), &bedrockruntime.InvokeModelInput{
		Body:        payloadBytes,
		ModelId:     aws.String("anthropic.claude-3-sonnet-20240229-v1:0"),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return "", fmt.Errorf("error Claude response: %v", err)
	}

	var resp domain.BRResponse

	err = json.Unmarshal(output.Body, &resp)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	return resp.ContentItem[len(resp.ContentItem)-1].Text, nil

}

func (b BedrockService) ClaudeMessageStreamCompletion(userMessage string, systemPrompt string) (chan string, error) {
	ch := make(chan string)
	content := domain.BRContent{Type: "text", Text: userMessage}
	msg := domain.BRMessage{Role: "user", Content: []domain.BRContent{content}}
	payload := domain.BRRequest{
		AnthropicVersion: "bedrock-2023-05-31",
		MaxTokens:        512,
		System:           systemPrompt,
		Messages:         []domain.BRMessage{msg},
		Temperature:      0.5,
		TopP:             0.9,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	go func() {
		defer close(ch)
		out, err := b.brc.InvokeModelWithResponseStream(context.Background(), &bedrockruntime.InvokeModelWithResponseStreamInput{
			Body:        payloadBytes,
			ModelId:     aws.String("anthropic.claude-3-sonnet-20240229-v1:0"),
			ContentType: aws.String("application/json"),
		})
		if err != nil {
			return
		}
		for event := range out.GetStream().Events() {
			switch v := event.(type) {
			case *types.ResponseStreamMemberChunk:
				fmt.Printf("data: %s\n", string(v.Value.Bytes))
				var resp domain.BRStreamResponse
				err := json.NewDecoder(bytes.NewReader(v.Value.Bytes)).Decode(&resp)
				if err != nil {
					return
				}
				ch <- resp.Delta.Text
			case *types.UnknownUnionMember:
				fmt.Println("unknown tag:", v.Tag)
			default:
				fmt.Println("union is nil or unknown type")
			}
		}
	}()
	return ch, nil
}
