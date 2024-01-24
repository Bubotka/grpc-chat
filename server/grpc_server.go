package server

import (
	"log"
	"net"

	"github.com/sashabaranov/go-openai"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Bubotka/grpc-chat/config"
	"github.com/Bubotka/grpc-chat/gen/pb"
	gpt "github.com/Bubotka/grpc-chat/pkg/chat-gpt"
)

type ChatGptServer struct {
	pb.UnimplementedChatServiceServer
	systemPrompt gpt.RequestMessage
	aiClient     *openai.Client
	prevPrompts  []gpt.RequestMessage
	cfg          config.Config
	//client       *gpt.ChatGptClient
}

func NewChatGptServer(cfg config.Config, systemPrompt string) *ChatGptServer {
	aiClient := openai.NewClient(cfg.OpenAIKey)
	sp := gpt.RequestMessage{
		Role:    "system",
		Content: systemPrompt,
	}

	return &ChatGptServer{
		systemPrompt: sp,
		aiClient:     aiClient,
		cfg:          cfg,
	}
}

func (s *ChatGptServer) Start(listener net.Listener) error {
	srv := grpc.NewServer()
	pb.RegisterChatServiceServer(srv, s)

	reflection.Register(srv)

	log.Println("Starting server... on port", s.cfg.GRPC.Port)

	return srv.Serve(listener)
}
