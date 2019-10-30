package main

import (
	"github.com/agh-eduweb/chat"
	"github.com/agh-eduweb/class"
	"github.com/agh-eduweb/feedback"
	"github.com/agh-eduweb/homework"
	"sync"
)

type User struct {
	Id           int
	OAuthService string
}

//func init() {
//	err := gotenv.Load("./../.env")
//	if err != nil {
//		panic("Error loading .env!")
//	}
//}

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

}
