package feedback

import (
	"context"
	"generated/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

var feedbacks = make([]*educloud.Feedback, 0, 10)

type feedbackServiceServer struct{}

func (s *feedbackServiceServer) SendFeedback(ctx context.Context, fb *educloud.Feedback) (*educloud.Status, error) {
	feedbacks = append(feedbacks, fb)
	print("GOT FEEDBACK", fb.Name, fb.Note)
	return &educloud.Status{}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50060")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	chatServer := grpc.NewServer()
	educloud.RegisterFeedbackServiceServer(chatServer, &feedbackServiceServer{})
	if err := chatServer.Serve(lis); err != nil {
		println("Chat server failed")
	}
}
