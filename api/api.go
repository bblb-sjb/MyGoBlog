package api

import (
	"myWeb/context"
	"net/http"
)

var API APIResponder = &Api{}

type APIResponder interface {
	SaveAndUpdatePost(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetPost(w http.ResponseWriter, r *http.Request)
	UploadImage(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)

	LoginNew(ctx *context.MsContext)
	GetPostNew(ctx *context.MsContext)
	UploadImageNew(ctx *context.MsContext)
	SearchNew(ctx *context.MsContext)
	SaveAndUpdatePostNew(ctx *context.MsContext)
}
type Api struct {
}
