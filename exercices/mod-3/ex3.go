package main

import (
	"fmt"
	"time"
)

/*
	Ping-Pong:
	This example demonstrates how to use goroutines and channels to create a simple ping-pong game.
*/

func main(){
	var c = make(chan string)
	go ping(c)
	go print(c)
	go pong(c)

	var exit string
	fmt.Println("Press Enter to exit...")
	fmt.Scanln(&exit)
}

func ping(c chan string){
	for i:= 0; ; i++{
		c <- "ping"
	}
}

func pong(c chan string){
	for i:= 0; ; i++{
		c <- "pong"
	}
}

func print(c chan string){
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}