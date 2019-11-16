package quiz

import (
	. "backend/generated/protos"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

var streams = make(map[string]QuizService_WaitForQuestionsServer)
var mutex = &sync.Mutex{}

type quizServiceServer struct{}

func (s *quizServiceServer) WaitForQuestions(request *QuizRequest, stream QuizService_WaitForQuestionsServer) error {
	mutex.Lock()
	streams[request.UserId] = stream
	mutex.Unlock()

	time.Sleep(24 * time.Hour)

	return nil
}

func (s *quizServiceServer) AnswerQuestion(ctx context.Context, answer *QuizAnswer) (*Status, error) {
	if answer.Answer == "A" {
		return &Status{Code: Status_OK}, nil
	}
	return &Status{Code: Status_DENIED}, nil
}

func Simulate() {
	//for {
	time.Sleep(30 * time.Second)
	println("SENDING QUIZ")
	for _, stream := range streams {
		_ = stream.Send(&Question{ClassId: "11111", PhotoQuestion: true, ClosedQuestion: false, Answers: []string{"A", "BE", "CE", "DE"}})
	}
	//}
}

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50062")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	chatServer := grpc.NewServer()
	RegisterQuizServiceServer(chatServer, &quizServiceServer{})

	if err := chatServer.Serve(lis); err != nil {
		println("Chat server failed")
	}
}
