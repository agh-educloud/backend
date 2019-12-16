package quiz

import (
	"../class/class_commons"
	. "../generated/protos/grpc"
	web_gen "../generated/protos/rest"
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

var currentQuizId = int32(0)
var mutex = &sync.Mutex{}

type quizServiceServer struct{}

func (s *quizServiceServer) WaitForQuestions(request *QuizRequest, stream QuizService_WaitForQuestionsServer) error {
	println("JOINED QUIZ")
	mutex.Lock()
	StudentQuizStream[request.UserId] = stream
	mutex.Unlock()
	time.Sleep(24 * time.Hour)

	return nil
}

func (s *quizServiceServer) AnswerQuestion(ctx context.Context, answer *QuizAnswer) (*Status, error) {

	if _, ok := StudentCorrectAnswer[answer.UserId]; ok {
		var expectedAnswer = StudentCorrectAnswer[answer.UserId]
		println("EXPECTED ANSWER", expectedAnswer)

		if expectedAnswer == answer.Answer {
			go updateStatistics()
			return &Status{Code: Status_OK}, nil
		}
		if strings.HasPrefix(answer.Answer, "http") {
			StudentPhotoAnswer[answer.UserId] = answer.Answer
			return &Status{Code: Status_OK}, nil
		}
	}
	return &Status{Code: Status_DENIED}, nil
}

func updateStatistics() {
	println("Updating stats!")
	quizStats := class_commons.Statistics[currentQuizId]
	correctAnswersNumber := float32(quizStats.Participants)*quizStats.PercentageOfCorrectAnswers + 1
	quizStats.Participants += 1
	quizStats.PercentageOfCorrectAnswers = correctAnswersNumber / float32(quizStats.Participants)
	class_commons.Statistics[currentQuizId] = quizStats
}

func SendClosedQuestion(classId string, quizId int32, possibleAnswers []string, correctAnswer string, assignToSubGroup bool, numberOfGroups int) {
	var students = class_commons.AllStudentsInClass[classId]
	currentQuizId = quizId
	class_commons.Statistics = append(class_commons.Statistics, web_gen.QuizQuestionStatistics{
		QuestionUuid:               quizId,
		PercentageOfCorrectAnswers: 0.0,
		Participants:               0,
	})

	println("Sending closed question")

	for index, student := range students {
		StudentCorrectAnswer[student] = correctAnswer

		println("CORRECT ANSWER", correctAnswer)
		var stream = StudentQuizStream[student]

		if assignToSubGroup {
			var result = stream.Send(&Question{ClassId: class_commons.ClassUuidToCode[classId], //mobile uses SecretCode instead of ClassUUID
				ClosedQuestion: true,
				Answers:        possibleAnswers,
				GroupId:        class_commons.ClassUuidToCode[classId] + string(index%numberOfGroups),
			})
			println(result)
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
