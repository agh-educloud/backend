package main

import (
	"../class"
)

func main() {

	//var wg sync.WaitGroup
	//
	//wg.Add(1)
	//go homework.StartServer()
	//
	//wg.Add(2)
	//go chat.StartServer()
	//
	//wg.Add(3)
	//go class.StartServer()
	//
	//wg.Add(4)
	//go feedback.StartServer()
	//
	//wg.Add(5)
	class.Start()

	//wg.Add(6)
	//go quiz.StartServer()
	//
	//wg.Wait()
}
