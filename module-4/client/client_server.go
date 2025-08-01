package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main(){
	/*
		HTTP Client:
		This code demonstrates how to create a simple HTTP client in Go that makes a GET request to a specified URL.
		It uses the `http` package to perform the request and handle the response.
	*/


	/*
		http.Get: Makes a GET request to the specified URL and returns the response.
		It returns an `http.Response` object and an error if the request fails.
		http.Response: Represents the response from the server, including status code, headers, and body.
	*/
	resp, err := http.Get("https://gobyexample.com")
	if(err != nil){
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)


	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i< 5;i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}