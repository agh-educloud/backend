package feedback

import (
	. "../generated/protos/grpc"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

var FeedbackPerSurname = make(map[string][]Feedback)

type feedbackServiceServer struct{}

func (s *feedbackServiceServer) SendFeedback(ctx context.Context, fb *Feedback) (*Status, error) {
	var surname = strings.ToTitle(fb.Name)
	FeedbackPerSurname[surname] = append(FeedbackPerSurname[surname], *fb)

	println("Saved feedback for", surname)
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
