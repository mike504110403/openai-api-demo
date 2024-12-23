package chatcompletions

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

var OPENAI_API_KEY string

const Poem_System_Role_Content = "你是一個專業的籤詩解說專家，請根據籤詩的內容回答用戶的問題，並提供詳細的解釋。"

// Init 初始化
func Init(key string) {
	OPENAI_API_KEY = key
}

// SendQuestion 用戶問題
func SendQuestion(question string, poem string, key string) string {
	client := openai.NewClient(OPENAI_API_KEY)
	prompt := fmt.Sprintf("以下是籤詩內容：\n%s\n問題：%s", poem, question)
	// 創建 ChatGPT 請求
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo, // 模型名稱
		Messages: []openai.ChatCompletionMessage{
			// 系統角色設定
			{Role: "system", Content: Poem_System_Role_Content},
			// 用戶輸入的籤詩和問題
			{Role: "user", Content: prompt},
		},
		Temperature: 0.7,
		MaxTokens:   500,
	}

	// 發送請求並獲取回應
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	// 打印回應
	content := resp.Choices[0].Message.Content
	if err := saveResponse(resp.ID, prompt, content); err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return content
}

// SaveResponse 保存回應
func saveResponse(id string, prompt string, response string) error {
	data := map[string]interface{}{
		"id":       id,
		"prompt":   prompt,
		"response": response,
	}

	file, _ := json.MarshalIndent(data, "", "  ")
	if err := os.WriteFile("response.json", file, 0644); err != nil {
		return err
	}
	return nil
}
