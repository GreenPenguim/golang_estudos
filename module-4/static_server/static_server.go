package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		You need to be on the root of this directory to work the server opening with  
	*/
	file_server := http.FileServer(http.Dir("./site"))
	http.Handle("/", file_server)
	log.Print("Listening on: 3000")
	log.Print("Open on http://localhost:3000/home.html")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal()
	}
}

