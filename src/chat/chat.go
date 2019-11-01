package chat

import (
	"context"
	"generated/protos"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

var streams = make(map[string]educloud.ChatService_ReceiveMessagesServer)
var finished = false
var mutex = &sync.Mutex{}

type chatServiceServer struct{}

func (s *chatServiceServer) SendMessage(ctx context.Context, message *educloud.ChatMessage) (*educloud.Status, error) {
	println("Got request", len(streams))
	for _, stream := range streams {
		if stream != nil {
			if err := stream.Send(message); err != nil {
				return &educloud.Status{Code: educloud.Status_DENIED, Message: "Error", Details: "Error"}, err
			}
		}
		println("SENT", message.Message.Content)
	}
	println("EXIT")
	return &educloud.Status{}, nil
}

func (s *chatServiceServer) ReceiveMessages(user *educloud.User, stream educloud.ChatService_ReceiveMessagesServer) error {
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
	educloud.RegisterChatServiceServer(chatServer, &chatServiceServer{})

	if err := chatServer.Serve(lis); err != nil {
		println("Chat server failed")
	}
}

func foo() {
	for {
		continue
	}
}
