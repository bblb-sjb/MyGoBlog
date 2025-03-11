package views

import (
	"errors"
	"myWeb/common"
	"myWeb/context"
	"myWeb/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	//http://localhost:8080/p/7.html  7参数 文章的id
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	//7.html
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("Detail不识别此请求路径"))
		return
	}
	postRes, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)
}

func (*HTMLApi) DetailNew(ctx *context.MsContext) {
	detail := common.Template.Detail
	//http://localhost:8080/p/7.html  7参数 文章的id
	pIdStr := ctx.GetPathVariable("pid")
	//7.html
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pid, _ := strconv.Atoi(pIdStr)
	postRes, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteError(ctx.W, errors.New("DetailNew查询出错"))
		return
	}
	detail.WriteData(ctx.W, postRes)
}
