package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gitlab.com/Alexandrhub/grpc-chat/config"
	"gitlab.com/Alexandrhub/grpc-chat/gen/pb"
)

func main() {
	cfg := config.MustConfig()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.GRPC.Host, cfg.GRPC.Port), grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer func() { _ = conn.Close() }()

	client := pb.NewChatServiceClient(conn)
	ctx := context.Background()

	msg := &pb.MessageRequest{
		Content: "Hello, write please hello world in golang",
	}

	// unary, err := client.Chat(ctx, msg)
	// if err != nil {
	// 	log.Fatalf("failed to send message: %v", err)
	// }
	//
	// log.Printf("message received: %s", unary.GetContent())

	stream, err := client.ChatStream(ctx, msg)
	if err != nil {
		log.Fatalf("failed to send message: %v", err)
	}

	var currentLine string
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("\nend of stream")
				break
			}

			log.Fatalf("failed to receive message: %v", err.Error())
		}

		currentLine += msg.GetContent()

		if strings.HasSuffix(currentLine, "\n") {
			log.Println(currentLine)
			currentLine = ""
		}
	}

}
