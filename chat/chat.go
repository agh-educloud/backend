package chat

import (
	. "../generated/protos"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
)

var streams = make([]ChatService_ExchangeMessagesServer, 0)
var mutex = &sync.Mutex{}

type chatServiceServer struct{}

func (s *chatServiceServer) ExchangeMessages(stream ChatService_ExchangeMessagesServer) error {

	mutex.Lock()
	streams = append(streams, stream)
	mutex.Unlock()

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		for _, streamFromSlice := range streams {
			if streamFromSlice != nil {
				if err := streamFromSlice.Send(in); err != nil {
					return err
				}
			}

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
