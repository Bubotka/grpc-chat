package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sashabaranov/go-openai"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Bubotka/grpc-chat/gen/pb"
	gpt "github.com/Bubotka/grpc-chat/pkg/chat-gpt"
)

func (s *ChatGptServer) Chat(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("Recieved message: %s", req.Content)
	prompt := req.Content

	var messages []gpt.RequestMessage
	messages = append(messages, s.systemPrompt)
	messages = append(messages, s.prevPrompts...)

	question := gpt.RequestMessage{
		Role:      "user",
		Content:   prompt,
		CreatedAt: time.Now(),
	}

	messages = append(messages, question)

	resp, err := s.aiClient.CreateChatCompletion(
		ctx, openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		log.Print("Error in server.Chat.CreateChatCompletion(): ", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error processing prompt: %v", err))
	}
	respStr := resp.Choices[0].Message.Content
	log.Printf("Response finished: %s\n", respStr)
	answer := gpt.RequestMessage{
		Role:      "assistant",
		Content:   respStr,
		CreatedAt: time.Now(),
	}
	s.prevPrompts = append(s.prevPrompts, []gpt.RequestMessage{question, answer}...)

	twoHoursAgo := time.Now().Add(-2 * time.Hour)
	var filteredPrompts []gpt.RequestMessage
	for _, message := range s.prevPrompts {
		// check if the message is within the last two hours
		if message.CreatedAt.After(twoHoursAgo) {
			// keep the message in the filteredPrompts slice
			filteredPrompts = append(filteredPrompts, message)
		}
	}
	s.prevPrompts = filteredPrompts

	return &pb.MessageResponse{
		Content: answer.Content,
	}, nil
}
