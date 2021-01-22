package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
// Request represents an HTTP request received by a server or to be sent by a client.
func first(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello world!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "OOPS", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Learn Microservices with Go %s\n", d)
}

func second(rw http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to learn Microservices!")

	fmt.Fprintf(rw, "Learn Microservices with Go")
}

func main() {
	//HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	http.HandleFunc("/first", first)
	http.HandleFunc("/second", second)

	/*ListenAndServe listens on the TCP newtwork address addr and
	  then calls Serve with handler to handle requests on incoming connections*/
	http.ListenAndServe(":8090", nil)

}

/* ServeMux is an HTTP request multiplexer
   It matches the URL of each incoming request against a list of registered
   patterns and calls the handler for the pattern that
   most closely matches the URL.*/
