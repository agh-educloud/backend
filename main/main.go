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
)

var classes []Class

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/class", createClass).Methods("POST")
	router.HandleFunc("/class/{id}", getClass).Methods("GET")
	router.HandleFunc("/class", getAllClasses).Methods("GET")
	router.HandleFunc("/class/{id}", updateClass).Methods("PATCH")
	router.HandleFunc("/class/{id}", deleteClass).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func deleteClass(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	var classUuid = mux.Vars(request)["id"]

	for i, v := range classes {
		if v.Uuid == classUuid {
			classes = append(classes[:i], classes[i+1:]...)
			writer.WriteHeader(http.StatusOK)
			return
		}
	}

	writer.WriteHeader(http.StatusNotFound)
}

func getAllClasses(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(interface{}(json.Marshal(classes)))

}
func getClass(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	var classUuid = mux.Vars(request)["id"]

	for _, v := range classes {
		if v.Uuid == classUuid {
			writer.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(writer).Encode(interface{}(json.Marshal(v)))
			return
		}
	}

	writer.WriteHeader(http.StatusNotFound)
}

func createClass(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	var class = ClassCreationRequest{}
	reqBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		_, _ = fmt.Fprintf(writer, "Kindly enter data with the event title and description only in order to update")
	}
	_ = json.Unmarshal(reqBody, &class)

	var nextUuid = int32(len(classes) + 1)
	var classUuid = strconv.FormatInt(int64(nextUuid), 10)
	classes = append(classes, Class{Uuid: classUuid})

	log.Println("Created class")
	writer.WriteHeader(http.StatusCreated)
	var response = ClassCreationResponse{ClassUuid: nextUuid}
	_ = json.NewEncoder(writer).Encode(response)
}

func updateClass(writer http.ResponseWriter, request *http.Request) {
	var class = ClassUpdateRequest{}
	var classUuid = mux.Vars(request)["id"]

	reqBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		_, _ = fmt.Fprintf(writer, "Kindly enter data with the event title and description only in order to update")
	}
	_ = json.Unmarshal(reqBody, &class)

	// TODO now updating only class name
	for _, v := range classes {
		if v.Uuid == classUuid {
			if len(class.Class.Name) > 0 {
				v.Name = class.Class.Name
				log.Println("Updated class name")
			} else if len(class.Class.Topic) > 0 {
				v.Topic = class.Class.Topic
				log.Println("Updated class topic")
			} else if len(class.Class.QuizQuestion) > 0 {
				v.QuizQuestion = class.Class.QuizQuestion
				log.Println("Updated class quiz questions")
			}
		}
	}

	log.Println("Updated class")
	writer.WriteHeader(http.StatusAccepted)
	var response = Status{Code: Status_OK}

	_ = json.NewEncoder(writer).Encode(response)
}
