package homework

import (
	. "../generated/protos"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

var homeworks = make([]*Homework, 0, 10)
var count = 0

type userHomeworkServiceServer struct{}

func (s *userHomeworkServiceServer) GetHomeworks(context.Context, *Status) (*Homeworks, error) {
	homeworks = nil
	homeworks = append(homeworks, &Homework{HomeworkUuid: 123, Url: "https://i.imgur.com/KUJmgtE.jpg", MailTo: "teach@agh.edu.pl"})
	homeworks = append(homeworks, &Homework{HomeworkUuid: 321, Url: "https://i.imgur.com/yyQn13a.jpg", MailTo: "teach@agh.edu.pl"})
	return &Homeworks{Homework: homeworks}, nil
}

func (s *userHomeworkServiceServer) SendHomeworkSolution(context.Context, *Homework) (*Status, error) {
	return nil, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	chatServer := grpc.NewServer()
	RegisterUserHomeworkServiceServer(chatServer, &userHomeworkServiceServer{})
	print("XD")
	if err := chatServer.Serve(lis); err != nil {
		println("Chat server failed")
	}
}
