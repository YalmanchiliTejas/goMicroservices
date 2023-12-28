package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type hello struct {
	l *log.Logger
}

func MakeHello(l *log.Logger) *hello {

	return &hello{l}
}

func (h *hello) ServeHttp(w http.ResponseWriter, r *http.Request) {

	h.l.Println("Hello World")
	d, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}
	h.l.Printf("Data %s", d)

	fmt.Fprintf(w, "Hello World %s", d)
}
