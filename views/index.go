package views

import (
	"errors"
	"log"
	"myWeb/common"
	"myWeb/context"
	"myWeb/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	err := r.ParseForm()
	if err != nil {
		log.Printf("解析请求参数异常：%v", err)
		index.WriteError(w, errors.New("解析请求参数异常,请联系管理员"))
		return
	}
	//获取分页信息
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//获取每页显示的条数
	limitStr := r.Form.Get("limit")
	limit := 10
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, limit)
	if err != nil {
		log.Printf("查询Index信息异常：%v", err)
		index.WriteError(w, errors.New("查询Index信息异常,请联系管理员"))
	}

	index.WriteData(w, hr)
}

func (*HTMLApi) IndexNew(ctx *context.MsContext) {
	index := common.Template.Index

	pageStr, _ := ctx.GetForm("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//获取每页显示的条数
	limitStr, _ := ctx.GetForm("limit")
	limit := 10
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}
	slug := ctx.GetPathVariable("slug")
	hr, err := service.GetAllIndexInfo(slug, page, limit)
	if err != nil {
		log.Printf("查询Index信息异常：%v", err)
		index.WriteError(ctx.W, errors.New("查询Index信息异常,请联系管理员"))
	}

	index.WriteData(ctx.W, hr)
}

func (*HTMLApi) SlugNew(ctx *context.MsContext) {
	index := common.Template.Index

	pageStr, _ := ctx.GetForm("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//获取每页显示的条数
	limitStr, _ := ctx.GetForm("limit")
	limit := 10
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}
	slug := ctx.GetPathVariable("slug")
	hr, err := service.GetAllIndexInfo(slug, page, limit)
	if err != nil {
		log.Printf("查询Index信息异常：%v", err)
		index.WriteError(ctx.W, errors.New("查询Index信息异常,请联系管理员"))
	}

	index.WriteData(ctx.W, hr)

}
