package main

import (
	"fmt"
	"log"
	"net/http"
)

func first(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello world!")

	fmt.Fprintf(rw, "Learn Microservices with Go")
}

func second(rw http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to learn Microservices!")

	fmt.Fprintf(rw, "Learn Microservices with Go")
}

func main() {
	http.HandleFunc("/first", first)
	http.HandleFunc("/second", second)

	/*ListenAndServe listens on the TCP newtwork address addr and
	  then calls Serve with handler to handle requests on incoming connections*/
	http.ListenAndServe(":8090", nil)

}
