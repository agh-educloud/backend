package class

import (
	"context"
	"generated/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

var classes = make([]string, 0, 10)

type classUserService struct{}

func (s *classUserService) JoinClass(ctx context.Context, request *educloud.JoinClassRequest) (*educloud.Status, error) {
	classes = append(classes, "11111")

	return &educloud.Status{Code: stringInSlice(request.SecretCode, classes)}, nil
}

func stringInSlice(a string, list []string) educloud.Status_Code {
	for _, b := range list {
		if b == a {
			return educloud.Status_OK
		}
	}
	return educloud.Status_DENIED
}

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50099")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	chatServer := grpc.NewServer()
	educloud.RegisterUserClassServiceServer(chatServer, &classUserService{})

	if err := chatServer.Serve(lis); err != nil {
		println("Chat server failed")
	}
}
