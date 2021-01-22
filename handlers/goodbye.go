package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	g.l.Println("Thanking you!")

	e, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "OOPS", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Thanks for Learning Microservices with Go %s\n", e)
}
