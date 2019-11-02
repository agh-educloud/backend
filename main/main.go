package main

import (
	. "../generated/protos"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"../chat"
	"../class"
	"../feedback"
	"../homework"
	"sync"
)

var classes []*ClassWithUuid

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go homework.StartServer()

	wg.Add(2)
	go chat.StartServer()

	wg.Add(3)
	go class.StartServer()

	wg.Add(4)
	go feedback.StartServer()

	wg.Wait()

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

	for i, v := range classes {
		if v.ClassUuid == int32(classUuidInt) {
			classes = append(classes[:i], classes[i+1:]...)
			log.Println("Deleted class")
			writer.WriteHeader(http.StatusOK)
			return
		}
	}

	writer.WriteHeader(http.StatusNotFound)
}

func getAllClasses(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	var response = GetClassesResponse{Classes: classes}
	_ = json.NewEncoder(writer).Encode(response)

}
func getClass(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	var classUuid = mux.Vars(request)["id"]
	var classUuidInt, _ = strconv.ParseInt(classUuid, 10, 32)

	for _, v := range classes {
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
	var wg sync.WaitGroup

	wg.Add(1)
	go homework.StartServer()

	wg.Add(2)
	go chat.StartServer()

	wg.Add(3)
	go class.StartServer()

	wg.Add(4)
	go feedback.StartServer()

	wg.Wait()

	var nextUuid = int32(len(classes) + 1)

	classes = append(classes, &ClassWithUuid{
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
	for _, v := range classes {
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
