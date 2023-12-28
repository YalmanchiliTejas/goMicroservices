package handlers

import (
	"log"
	"net/http"
)

type p_handles struct {
	l *log.Logger
}

func NewProdcut(l *log.Logger) *p_handles {

	return &p_handles{l}
}

func (*p_handles) ServeHttp(w http.ResponseWriter, r *http.Request) {

}
