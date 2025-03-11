package api

import (
	"errors"
	"myWeb/common"
	"myWeb/context"
	"myWeb/service"
	"net/http"
)

func (*Api) Search(w http.ResponseWriter, r *http.Request) {
	//api/v1/post/search?val=xxxx
	val := r.URL.Query().Get("val")
	if val == "" {
		common.ErrorResult(w, errors.New("搜索参数不能为空"))
		return
	}
	posts, err := service.SearchPost(val)
	if err != nil {
		common.ErrorResult(w, errors.New("搜索出错"))
		return
	}
	common.SuccessResult(w, posts)
}

func (*Api) SearchNew(ctx *context.MsContext) {
	//api/v1/post/search?val=xxxx
	val := ctx.Request.URL.Query().Get("value")
	if val == "" {
		common.ErrorResult(ctx.W, errors.New("搜索参数不能为空"))
		return
	}
	posts, err := service.SearchPost(val)
	if err != nil {
		common.ErrorResult(ctx.W, errors.New("搜索出错"))
		return
	}
	common.SuccessResult(ctx.W, posts)
}
