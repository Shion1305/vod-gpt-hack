package main

import (
	"api/pkg/domain"
	"api/pkg/service"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bedrock, err := service.NewBedrockService()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	for {
		fmt.Println("Enter your message:")
		userMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from stdin: %v", err)
		}

		userMessage = strings.TrimSpace(userMessage)
		systemPrompt := "あなたは魚大好き星人です。なんでも魚に例えて話してみてください"
		response, err := bedrock.ClaudeMessageCompletion(userMessage, systemPrompt)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Println("Response:", response)
		streamResp, err := bedrock.ClaudeMessageStreamCompletion(response, systemPrompt)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		for {
			if mes, ok := <-streamResp; ok {
				fmt.Println("Response:", mes)
			} else {
				break
			}
		}
	}
}

func claudeMessageCompletion(userMessage string, systemPrompt string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"))
	if err != nil {
		return "", fmt.Errorf("error Config Load: %v", err)
	}

	// Bedrockクライアントを初期化
	brc := bedrockruntime.NewFromConfig(cfg)

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

	output, err := brc.InvokeModel(context.Background(), &bedrockruntime.InvokeModelInput{
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
