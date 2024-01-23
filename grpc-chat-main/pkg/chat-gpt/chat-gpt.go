package chat_gpt

import (
	"time"
)

type ChatGptRequest struct {
	Model    string           `json:"model"`
	Messages []RequestMessage `json:"messages"`
	Stream   bool             `json:"stream"`
}

type RequestMessage struct {
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"-"`
}

type ChatGptResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choice struct {
	Message      RequestMessage `json:"message"`
	FinishReason string         `json:"finish_reason"`
	Index        int            `json:"index"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
