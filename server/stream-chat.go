package server

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"

	"github.com/Bubotka/grpc-chat/gen/pb"
)

func (s *ChatGptServer) ChatStream(messageRequest *pb.MessageRequest, stream pb.ChatService_ChatStreamServer) error {
	ctx := context.Background()

	aiReq := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 4000,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: messageRequest.Content,
			},
		},
		Stream: true,
	}

	aiStream, err := s.aiClient.CreateChatCompletionStream(ctx, aiReq)
	if err != nil {
		log.Println("failed to create chat completion stream: ", err)

		return err
	}

	for {
		aiRes, err := aiStream.Recv()
		if errors.Is(err, io.EOF) {
			log.Println("\nEnd of stream")

			aiStream.Close()

			break
		}
		if err != nil {
			log.Println("failed to receive chat completion stream: ", err)

			aiStream.Close()

			return err
		}

		res := &pb.MessageResponse{
			Content: aiRes.Choices[0].Delta.Content,
		}

		if err := stream.Send(res); err != nil {
			log.Println("failed to send message: ", err)

			aiStream.Close()

			return err
		}
	}

	return nil
}
