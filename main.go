package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/saimohanreddyt/Microservices_with_Go.git/handlers"
)

func main() {

	l := log.New(os.Stdout, "Miceroservices", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/bye", gh)

	s := &http.Server{
		Addr:         ":8090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}

//A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
// Request represents an HTTP request received by a server or to be sent by a client.

//HandleFunc registers the handler function for the given pattern in the DefaultServeMux.

/*ListenAndServe listens on the TCP newtwork address addr and
  then calls Serve with handler to handle requests on incoming connections*/

/* ServeMux is an HTTP request multiplexer
   It matches the URL of each incoming request against a list of registered
   patterns and calls the handler for the pattern that
   most closely matches the URL.*/
