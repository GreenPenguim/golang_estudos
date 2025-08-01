package main

import(
	"fmt"
	"net/http"
)

/*
	Static Server:
	This code sets up a simple HTTP server that listens on port 8090 and serves two endpoints
	- `/hello`: Responds with "Hello, World!"
	- `/headers`: Displays the headers of the incoming request.
*/

/*
	http.ResponseWriter: Used to construct the HTTP response.
	http.Request: Represents the incoming HTTP request.
*/
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

/*
	Headers function:
	This function iterates over the headers of the incoming request and writes them to the response.
	It uses the `http.ResponseWriter` to send the headers back to the client.
	`req.Header` is a map of header names to values, allowing you to access and display the headers in the response.
*/
func headers(w http.ResponseWriter, req *http.Request){
	for name, headers := range req.Header{
		for _, h:=range headers{
			fmt.Fprintf(w, "%s: %s\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	/*
		http.ListenAndServe: Starts an HTTP server on the specified address (in this case, ":8090").
		It listens for incoming HTTP requests and routes them to the appropriate handler functions defined above.
		It blocks until the server is stopped, allowing it to handle requests concurrently.

		Checkout on [http://localhost:8090/hello]
		Checkout on [http://localhost:8090/headers]

	*/
	http.ListenAndServe(":8090", nil)
}