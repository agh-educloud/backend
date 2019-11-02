package class

import (
	. "backend/generated/protos"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
)

var webCreatedClasses []*ClassWithUuid

var classes = make([]string, 0, 10)

type classUserService struct{}

func (s *classUserService) JoinClass(ctx context.Context, request *JoinClassRequest) (*Status, error) {
	classes = append(classes, "11111")

	return &Status{Code: stringInSlice(request.SecretCode, classes)}, nil
}

func stringInSlice(a string, list []string) Status_Code {
	for _, b := range list {
		if b == a {
			return Status_OK
		}
	}
	return Status_DENIED
}

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50099")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	chatServer := grpc.NewServer()
	RegisterUserClassServiceServer(chatServer, &classUserService{})

	if err := chatServer.Serve(lis); err != nil {
		println("Chat server failed")
	}
}

func Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/class", createClass).Methods("POST")
	router.HandleFunc("/class/{id}", getClass).Methods("GET")
	router.HandleFunc("/class", getAllClasses).Methods("GET")
	router.HandleFunc("/class/{id}", updateClass).Methods("PATCH")
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

	var response = GetClassesResponse{Classes: webCreatedClasses}
	_ = json.NewEncoder(writer).Encode(response)

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

	var class Class

	err = class.XXX_Unmarshal(reqBody)

	err = json.Unmarshal(reqBody, &class)
	if err != nil {
		fmt.Println("Error while unmarshal")
	}
	log.Println(class.Name)

	var nextUuid = int32(len(classes) + 1)

	webCreatedClasses = append(webCreatedClasses, &ClassWithUuid{
		ClassUuid: nextUuid,
		Class:     &class,
	})
	log.Println("Created class")

	var response = ClassCreationResponse{
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

	var class ClassUpdateRequest

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
			} else if len(class.Class.Topic) > 0 {
				v.Class.Topic = class.Class.Topic
				log.Println("Updated class topic")
			} else if len(class.Class.QuizQuestion) > 0 {
				v.Class.QuizQuestion = class.Class.QuizQuestion
				log.Println("Updated class quiz questions")
			}
		}
	}

	log.Println("Updated class")

	var response = Status{Code: Status_OK}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(response)
}
