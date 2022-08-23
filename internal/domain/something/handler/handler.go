package handler

import "net/http"

type SomethingHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
