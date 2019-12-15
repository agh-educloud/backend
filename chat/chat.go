package chat

import (
	class "../class/class_commons"
	. "../generated/protos/grpc"
	web_gen "../generated/protos/rest"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var StudentChatStreams = make(map[string]ChatService_ReceiveMessagesServer)
var mutex = &sync.Mutex{}

type chatServiceServer struct{}

func (s *chatServiceServer) SendMessage(ctx context.Context, message *ChatMessage) (*Status, error) {
	for _, stream := range StudentChatStreams {
		_ = stream.Send(message)
	}
	println("Message -> ", message.Message.Content)
	//Question to presenter
	if strings.HasPrefix(message.Message.Content, "/zapytaj") {
		println("Sending message to presenter")
		class.MessagesToPresenter = append(class.MessagesToPresenter, &web_gen.RestChatMessage{
			Message: &web_gen.RestMessage{
				Content: message.Message.Content,
			},
		})
	}
	return &Status{}, nil
}

func (s *chatServiceServer) ReceiveMessages(user *User, stream ChatService_ReceiveMessagesServer) error {
	mutex.Lock()
	StudentChatStreams[user.Uuid] = stream
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
