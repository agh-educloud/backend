package quiz

import (
	"../class/class_commons"
	. "../generated/protos/grpc"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var StudentQuizStream = make(map[string]QuizService_WaitForQuestionsServer)
var StudentCorrectAnswer = make(map[string]string)
var StudentPhotoAnswer = make(map[string]string)

var mutex = &sync.Mutex{}

type quizServiceServer struct{}

func (s *quizServiceServer) WaitForQuestions(request *QuizRequest, stream QuizService_WaitForQuestionsServer) error {
	mutex.Lock()
	StudentQuizStream[request.UserId] = stream
	mutex.Unlock()

	time.Sleep(24 * time.Hour)

	return nil
}

func (s *quizServiceServer) AnswerQuestion(ctx context.Context, answer *QuizAnswer) (*Status, error) {
	if _, ok := StudentCorrectAnswer[answer.UserId]; ok {
		var expectedAnswer = StudentCorrectAnswer[answer.UserId]
		if expectedAnswer == answer.Answer {
			return &Status{Code: Status_OK}, nil
		}
		if strings.HasPrefix(answer.Answer, "http") {
			StudentPhotoAnswer[answer.UserId] = answer.Answer
			return &Status{Code: Status_OK}, nil
		}
	}
	return &Status{Code: Status_DENIED}, nil
}

func SendClosedQuestion(classId string, possibleAnswers []string, correctAnswer string, assignToSubGroup bool, numberOfGroups int) {
	var students = class_commons.AllStudentsInClass[classId]

	for index, student := range students {
		StudentCorrectAnswer[student] = correctAnswer
		var stream = StudentQuizStream[student]
		if assignToSubGroup {
			_ = stream.Send(&Question{
				ClassId:        class_commons.ClassUuidToCode[classId], //mobile uses SecretCode instead of ClassUUID
				ClosedQuestion: true,
				Answers:        possibleAnswers,
				GroupId:        class_commons.ClassUuidToCode[classId] + string(index%numberOfGroups),
			})
		} else {
			_ = stream.Send(&Question{
				ClassId:        class_commons.ClassUuidToCode[classId], //mobile uses SecretCode instead of ClassUUID
				ClosedQuestion: true,
				Answers:        possibleAnswers,
			})
		}
	}
}

func SendPhotoQuestion(classId string, assignToSubGroup bool, numberOfGroups int) {
	var students = class_commons.AllStudentsInClass[classId]

	for index, student := range students {
		var stream = StudentQuizStream[student]
		if assignToSubGroup {
			_ = stream.Send(&Question{
				ClassId:       class_commons.ClassUuidToCode[classId], //mobile uses SecretCode instead of ClassUUID
				PhotoQuestion: true,
				GroupId:       class_commons.ClassUuidToCode[classId] + string(index%numberOfGroups),
			})
		} else {
			_ = stream.Send(&Question{
				ClassId:       class_commons.ClassUuidToCode[classId], //mobile uses SecretCode instead of ClassUUID
				PhotoQuestion: true,
			})
		}
	}
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
