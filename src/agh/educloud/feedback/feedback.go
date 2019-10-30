package feedback

import (
	. "agh/educloud/generated/protos"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

var feedbacks = make([]*Feedback, 0, 10)

type feedbackServiceServer struct{}

func (s *feedbackServiceServer) SendFeedback(ctx context.Context, fb *Feedback) (*Status, error) {
	feedbacks = append(feedbacks, fb)
	print("GOT FEEDBACK", fb.Name, fb.Note)
	return &Status{}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50060")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	chatServer := grpc.NewServer()
	RegisterFeedbackServiceServer(chatServer, &feedbackServiceServer{})
	if err := chatServer.Serve(lis); err != nil {
		println("Chat server failed")
	}
}
