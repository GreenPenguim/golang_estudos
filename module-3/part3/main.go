package main

import (
	"fmt"
	"time"
)

/*
	Concurrency is a powerful feature in Go that allows you to run multiple tasks simultaneously.
	It is achieved using goroutines, which are lightweight threads managed by the Go runtime.

	- Goroutines;
	- Channels;
	- Select statements;

*/

/*
	Goroutines:
	Capable of running functions concurrently, they are created using the `go` keyword.

*/

	func goRoutine(n int){
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

/*
	Channels:
	Used to communicate between goroutines, allowing them to send and receive messages.
	
	`chan` is the keyword used to create a channel.
	`<-` is used to send and receive messages from a channel.

*/

func ping(c chan string){
	for i:=0; ;i++{
		c <- "ping"
	}
}

func print(c chan string){
	for {
		msg:= <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

/*
	Select statements:
	Used to wait on multiple channel operations, allowing you to handle multiple channels in a single goroutine.
	Select statements block until one of the cases can proceed, making it a powerful tool for managing concurrency.
*/

func selectStatement() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select{
			case msg1 := <-c1:
				fmt.Println("Received", msg1)
			
			case msg2 := <-c2:
				fmt.Println("Received", msg2)
		}
	}

}

func main(){
	if false {
		// goRoutine
		go goRoutine(0)
		var write string
		fmt.Scanln(&write)
		
		// Channels
		var c1 chan string = make(chan string)
		
		go ping(c1)
		go print(c1)
		
		var input string
		fmt.Scanln(&input)
	}

	// Select statements
	selectStatement()
}