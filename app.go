package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	// writes the expected string to the responsewriter
	w.Write([]byte("Hello World! \n(from ApoorvaJ)\n"))
}

func main() {
	// create an HTTP request multiplexer
	mux := http.NewServeMux()

	// The request to localhost:8080/ will be handled by 'sayHelloHandler()'
	mux.HandleFunc("/", sayHelloHandler)

	// Making the server listen at port '8080'
	fmt.Println("Starting the server now...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
