package class

import (
	grpc_gen "../generated/protos/grpc"
	web_gen "../generated/protos/rest"
	"../quiz"
	"./class_commons"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
)

var webCreatedClasses []*web_gen.ClassWithUuid

type classUserService struct{}

func (s *classUserService) SendMessageToPresenter(context.Context, *grpc_gen.ChatMessage) (*grpc_gen.Status, error) {
	panic("implement me")
}

func (s *classUserService) JoinClass(ctx context.Context, request *grpc_gen.JoinClassRequest) (*grpc_gen.Status, error) {
	if _, ok := class_commons.CodesToClassUuid[request.SecretCode]; ok {
		var classId = class_commons.CodesToClassUuid[request.SecretCode]
		class_commons.AllStudentsInClass[classId] = append(class_commons.AllStudentsInClass[classId], request.User.Uuid)

		return &grpc_gen.Status{Code: grpc_gen.Status_OK}, nil
	} else {
		return &grpc_gen.Status{Code: grpc_gen.Status_DENIED}, nil
	}
}

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50099")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	chatServer := grpc.NewServer()
	grpc_gen.RegisterUserClassServiceServer(chatServer, &classUserService{})

	if err := chatServer.Serve(lis); err != nil {
		println("Chat server failed")
	}
}

func Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/class", createClass).Methods("POST")
	router.HandleFunc("/class/{id}", getClass).Methods("GET")
	router.HandleFunc("/quizStatistics/{id}", getQuizStatistics).Methods("GET")
	router.HandleFunc("/studentsQuestion/{id}", getStudentsQuestion).Methods("GET")
	router.HandleFunc("/quizHistoryStatistics/{id}", getQuizHistoryStatistics).Methods("GET")
	router.HandleFunc("/openQuizQuestionAnswers/{id}", getOpenQuizQuestionAnswers).Methods("GET")
	router.HandleFunc("/class", getAllClasses).Methods("GET")
	router.HandleFunc("/class/{id}", updateClass).Methods("PATCH")
	router.HandleFunc("/startClass/{id}", startClass).Methods("POST")
	router.HandleFunc("/endClass/{id}", endClass).Methods("POST")
	router.HandleFunc("/quizToDelegate/{id}", delegateQuizQuestion).Methods("POST")
	router.HandleFunc("/class/{id}", deleteClass).Methods("DELETE")
	log.Fatal(http.ListenAndServe("localhost:8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func deleteClass(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	log.Println("Deleted class method")

	var classUuid = mux.Vars(request)["id"]
	var classUuidInt, _ = strconv.ParseInt(classUuid, 10, 32)

	for i, v := range webCreatedClasses {
		if v.ClassUuid == int32(classUuidInt) {
			webCreatedClasses = append(webCreatedClasses[:i], webCreatedClasses[i+1:]...)
			log.Println("Deleted class")
			writer.WriteHeader(http.StatusOK)
			return
		}
	}

	writer.WriteHeader(http.StatusNotFound)
}

func getAllClasses(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	var response = web_gen.GetClassesResponse{Classes: webCreatedClasses}
	_ = json.NewEncoder(writer).Encode(response)

}
func getQuizStatistics(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	var questionUuid = mux.Vars(request)["id"]

	var questionUuidInt, _ = strconv.ParseInt(questionUuid, 10, 32)

	if len(class_commons.Statistics) > int(questionUuidInt) {
		writer.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(writer).Encode(class_commons.Statistics[questionUuidInt])
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func getStudentsQuestion(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(
		web_gen.StudentQuestions{
			Message: class_commons.MessagesToPresenter,
		})
}

func getQuizHistoryStatistics(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	var classUuid = mux.Vars(request)["id"]

	var classUuidInt, _ = strconv.ParseInt(classUuid, 10, 32)

	println(len(class_commons.PresentationHistoryData))
	if len(class_commons.PresentationHistoryData) > int(classUuidInt) {
		writer.WriteHeader(http.StatusOK)

		var statsList []*web_gen.QuizQuestionStatistics
		for _, v := range class_commons.PresentationHistoryData[int32(classUuidInt)] {
			statsList = append(statsList, &v)
		}

		stats := web_gen.QuizzesHistoryStatistics{
			QuizQuestionStatistics: statsList,
		}
		_ = json.NewEncoder(writer).Encode(stats)
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func getOpenQuizQuestionAnswers(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	urls := []string{
		"https://images.unsplash.com/photo-1522205408450-add114ad53fe?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=368f45b0888aeb0b7b08e3a1084d3ede&auto=format&fit=crop&w=1950&q=80",
	}
	for _, val := range quiz.StudentPhotoAnswer {
		urls = append(urls, val)
	}

	res := web_gen.OpenQuizQuestionAnswers{
		Url: urls,
	}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(res)
}

func getClass(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	var classUuid = mux.Vars(request)["id"]
	var classUuidInt, _ = strconv.ParseInt(classUuid, 10, 32)

	for _, v := range webCreatedClasses {
		if v.ClassUuid == int32(classUuidInt) {
			writer.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(writer).Encode(v)
			return
		}
	}

	writer.WriteHeader(http.StatusNotFound)
}

func createClass(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		_, _ = fmt.Fprintf(writer, "Kindly enter data with the event title and description only in order to update")
	}

	var class web_gen.RestClass

	err = json.Unmarshal(reqBody, &class)
	if err != nil {
		fmt.Println("Error while unmarshal")
	}
	log.Println(class.Name)
	log.Println(class.QuizQuestion)

	var nextUuid = int32(len(webCreatedClasses))

	webCreatedClasses = append(webCreatedClasses, &web_gen.ClassWithUuid{
		ClassUuid: nextUuid,
		Class:     &class,
	})
	log.Println("Created class")

	var response = web_gen.ClassCreationResponse{
		ClassUuid:  nextUuid,
		SecretCode: 12345,
		Error:      nil,
	}
	//marshalled, _ := response.XXX_Marshal(reqBody,false)
	writer.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(writer).Encode(response)
}

func updateClass(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	var classUuid = mux.Vars(request)["id"]
	var classUuidInt, _ = strconv.ParseInt(classUuid, 10, 32)

	reqBody, err := ioutil.ReadAll(request.Body)

	var class web_gen.ClassUpdateRequest

	err = class.XXX_Unmarshal(reqBody)
	err = json.Unmarshal(reqBody, &class)
	if err != nil {
		fmt.Println("Error while unmarshal")
	}

	//TODO more fields?
	for _, v := range webCreatedClasses {
		if v.ClassUuid == int32(classUuidInt) {
			if len(class.Class.Name) > 0 {
				v.Class.Name = class.Class.Name
				log.Println("Updated class name")
			}
			if len(class.Class.Topic) > 0 {
				v.Class.Topic = class.Class.Topic
				log.Println("Updated class topic")
			}
			if len(class.Class.QuizQuestion) > 0 {
				v.Class.QuizQuestion = class.Class.QuizQuestion
				log.Println("Updated class quiz questions")
			}
		}
	}

	log.Println("Updated class")

	var response = web_gen.RestStatus{Code: web_gen.RestStatus_OK}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(response)
}

func startClass(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	log.Println("Starting class")

	var classUuid = mux.Vars(request)["id"]
	var classUuidInt, _ = strconv.ParseInt(classUuid, 10, 32)

	for _, v := range webCreatedClasses {
		if v.ClassUuid == int32(classUuidInt) {
			writer.WriteHeader(http.StatusOK)
			var code = generateCode()
			class_commons.CodesToClassUuid[code] = classUuid
			class_commons.ClassUuidToCode[classUuid] = code
			log.Println("Created secret code for class " + classUuid + " : " + code)
			var response = web_gen.ClassCode{
				Code: code,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}
	}

	writer.WriteHeader(http.StatusBadRequest)
}

func endClass(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	log.Println("End class")

	var classUuid = mux.Vars(request)["id"]
	var classUuidInt, _ = strconv.ParseInt(classUuid, 10, 32)

	for _, v := range webCreatedClasses {
		if v.ClassUuid == int32(classUuidInt) {
			writer.WriteHeader(http.StatusOK)
			class_commons.PresentationHistoryData[int32(classUuidInt)] = class_commons.Statistics
			class_commons.Statistics = nil
			return
		}
	}

	writer.WriteHeader(http.StatusBadRequest)
}

func delegateQuizQuestion(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	var classUuid = mux.Vars(request)["id"]
	log.Println("Delegating quizToSend for class " + classUuid)
	var classUuidInt, _ = strconv.ParseInt(classUuid, 10, 32)

	reqBody, err := ioutil.ReadAll(request.Body)

	var quizId web_gen.RestQuizQuestionUuid

	err = json.Unmarshal(reqBody, &quizId)
	if err != nil {
		fmt.Println("Error while unmarshal")
	}

	for _, v := range webCreatedClasses {
		if v.ClassUuid == int32(classUuidInt) {
			for _, q := range v.Class.QuizQuestion {
				if q.Uuid == quizId.Uuid {
					writer.WriteHeader(http.StatusOK)
					possibleAnswers := getPossibleAnswers(q.Question.Option)

					println("classUUId", classUuid)
					println("quizId.Uuid", quizId.Uuid)
					println("possibleAnswers", possibleAnswers)
					//println(q.Question.Answer.Value)
					if len(possibleAnswers) == 0 {
						quiz.SendPhotoQuestion(classUuid, true, 1)
					} else {
						quiz.SendClosedQuestion(classUuid, quizId.Uuid, possibleAnswers, q.Question.Answer.Value, true, 1)
					}
					return
				}
			}
		}
	}

	writer.WriteHeader(http.StatusBadRequest)
}

func getPossibleAnswers(options []*web_gen.RestOption) []string {
	var opt = make([]string, len(options))
	for i, v := range options {
		opt[i] = v.Value
	}

	return opt
}

var letterRunes = []rune("0123456789")

func generateCode() string {
	str := make([]rune, 5)
	for i := range str {
		str[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(str)
}
