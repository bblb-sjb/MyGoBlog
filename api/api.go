package api

import "net/http"

var API APIResponder = &Api{}

type APIResponder interface {
	SaveAndUpdatePost(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}
type Api struct {
}
