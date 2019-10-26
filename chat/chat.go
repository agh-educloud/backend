package chat

import (
	. "../generated/protos"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

var streams = make(map[string]ChatService_ReceiveMessagesServer)
var finished = false
var mutex = &sync.Mutex{}

type chatServiceServer struct{}

func (s *chatServiceServer) SendMessage(ctx context.Context, message *ChatMessage) (*Status, error) {
	println("Got request", len(streams))
	for _, stream := range streams {
		if stream != nil {
			if err := stream.Send(message); err != nil {
				return &Status{Code: Status_DENIED, Message: "Error", Details: "Error"}, err
			}
		}
		println("SENT", message.Message.Content)
	}
	println("EXIT")
	return &Status{}, nil
}

func (s *chatServiceServer) ReceiveMessages(user *User, stream ChatService_ReceiveMessagesServer) error {
	mutex.Lock()
	streams[user.Uuid] = stream
	mutex.Unlock()

	time.Sleep(24 * time.Hour)

	return nil
}

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	chatServer := grpc.NewServer()
	RegisterChatServiceServer(chatServer, &chatServiceServer{})

	if err := chatServer.Serve(lis); err != nil {
		println("Chat server failed")
	}
}

func foo() {
	for {
		continue
	}
}
