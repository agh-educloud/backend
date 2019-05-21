package chat

import (
	. "../generated/protos"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type chatServiceServer struct{}

func (s *chatServiceServer) ExchangeMessages(stream ChatService_ExchangeMessagesServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		println("Got", in.Message.Content, "sending back")

		if err := stream.Send(in); err != nil {
			return err
		}

	}
	return nil
}

func StartServer() {
	lis, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	println("Server starting")
	grpcServer := grpc.NewServer()
	RegisterChatServiceServer(grpcServer, &chatServiceServer{})
	_ = grpcServer.Serve(lis)

}
