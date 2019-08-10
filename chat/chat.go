package chat

import (
	. "../generated/protos"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

var streams = make(map[*User]ChatService_ReceiveMessagesServer)
var finished = false
var mutex = &sync.Mutex{}

type chatServiceServer struct{}

func (s *chatServiceServer) SendMessage(ctx context.Context, message *ChatMessage) (*Status, error) {
	println("Got request")
	for user, stream := range streams {
		println(user.Uuid)
		//if user.Uuid != message.Sender.Uuid {
		if stream != nil {
			if err := stream.Send(message); err != nil {
				return &Status{Code: Status_DENIED, Message: "Error", Details: "Error"}, err
			}
		}
		println("SENT", user.Name, stream)
		//} else {
		//	println("Same uuid")
		//}
	}

	return &Status{Code: Status_OK, Message: "Sent", Details: "XD"}, nil
}

func (s *chatServiceServer) ReceiveMessages(user *User, stream ChatService_ReceiveMessagesServer) error {
	mutex.Lock()
	streams[user] = stream
	mutex.Unlock()

	for {
		if finished {
			break
		}
	}

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
