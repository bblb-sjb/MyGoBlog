package views

import (
	"net/http"
)

var HTML HTMLRenderer = &HTMLApi{}

// 定义接口
type HTMLRenderer interface {
	Index(w http.ResponseWriter, r *http.Request)
	Category(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
}

type HTMLApi struct {
}
