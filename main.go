package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

var OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")

func main() {
	client := openai.NewClient(OPENAI_API_KEY)

	// 創建 ChatGPT 請求
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo, // 模型名稱
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: "What is blockchain?"},
		},
		Temperature: 0.5,
		MaxTokens:   1000,
	}

	// 發送請求並獲取回應
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印回應
	fmt.Println("Response:", resp.Choices[0].Message.Content)
}
