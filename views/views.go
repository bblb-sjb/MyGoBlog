package views

import (
	"myWeb/context"
	"net/http"
)

var HTML HTMLRenderer = &HTMLApi{}

// 定义接口
type HTMLRenderer interface {
	Index(w http.ResponseWriter, r *http.Request)
	Category(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
	Writing(w http.ResponseWriter, r *http.Request)
	Pigeonhole(w http.ResponseWriter, r *http.Request)

	IndexNew(ctx *context.MsContext)
	LoginNew(ctx *context.MsContext)
	CategoryNew(ctx *context.MsContext)
	DetailNew(ctx *context.MsContext)
	WritingNew(ctx *context.MsContext)
	PigeonholeNew(ctx *context.MsContext)
	SlugNew(ctx *context.MsContext)
}

type HTMLApi struct {
}
