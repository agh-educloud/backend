package quiz

import (
	. "../generated/protos"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type quizServiceServer struct{}

func (q quizServiceServer) CreateClass(context context.Context, request *ClassCreationRequest) (*ClassCreationResponse, error) {

	print("Class name: ", request.Class.Name)
	print("Class topic: ", request.Class.Name)
	if len(request.Class.QuizQuestion) > 0 {
		print("Class first question: ", request.Class.QuizQuestion[0].Question)
	}

	return nil, nil
}

func (q quizServiceServer) UpdateClass(context.Context, *ClassUpdateRequest) (*Status, error) {
	panic("implement me")
}

func (q quizServiceServer) DeleteClass(context.Context, *ClassDeleteRequest) (*Status, error) {
	panic("implement me")
}

func (q quizServiceServer) StartClass(*ClassStartRequest, PresenterClassService_StartClassServer) error {
	panic("implement me")
}

func (q quizServiceServer) EndClass(context.Context, *ClassEndRequest) (*Status, error) {
	panic("implement me")
}

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	println("wii")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	println("wii")

	classServer := grpc.NewServer()
	RegisterPresenterClassServiceServer(classServer, &quizServiceServer{})
	println("wii")

	if err := classServer.Serve(lis); err != nil {
		println("Chat server failed")
	} else {
		println("Running")
	}
}
