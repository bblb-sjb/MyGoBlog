package api

import (
	"errors"
	"log"
	"myWeb/common"
	"myWeb/context"
	"myWeb/models"
	"myWeb/service"
	"myWeb/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.ErrorResult(w, errors.New("GetPost不识别此请求路径"))
		return
	}
	post, err := service.GetPostById(pid)
	if err != nil {
		common.ErrorResult(w, errors.New("GetPost查询出错"))
		return
	}
	common.SuccessResult(w, post)
}

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//判断用户是否登录
	token := r.Header.Get("Authorization")
	if token == "" {
		log.Printf("SaveAndUpdatePost的token为空")
		return
	}
	//解析token
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.ErrorResult(w, errors.New("token解析失败"))
		return
	}
	uid := claim.Uid

	method := r.Method
	switch method {
	case http.MethodPost:
		params, err := common.GetRequestJsonParam(r)
		if err != nil {
			log.Printf("SaveAndUpdatePost的POST解析请求参数异常：%v", err)
			return
		}
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := 0
		if params["type"] != nil {
			postType, _ = params["type"].(int)
		}
		post := &models.Post{
			CategoryId: categoryId,
			Content:    content,
			Markdown:   markdown,
			Slug:       slug,
			Title:      title,
			Type:       postType,
			UserId:     uid,
			Pid:        -1,
			ViewCount:  0,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.SuccessResult(w, post)
	case http.MethodPut:
		params, err := common.GetRequestJsonParam(r)
		if err != nil {
			log.Printf("SaveAndUpdatePost的PUT解析请求参数异常：%v", err)
			return
		}
		pid := params["pid"].(float64)
		pId := int(pid)
		post, _ := service.GetPostById(pId)
		if post == nil {
			common.ErrorResult(w, errors.New("SaveAndUpdatePost查询不到文章"))
			return
		}
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := 0
		if params["type"] != nil {
			postType, _ = params["type"].(int)
		}
		post.CategoryId = categoryId
		post.Content = content
		post.Markdown = markdown
		post.Slug = slug
		post.Title = title
		post.Type = postType
		post.UpdateAt = time.Now()
		service.UpdatePost(post)
		common.SuccessResult(w, post)
	}
}

func (*Api) GetPostNew(ctx *context.MsContext) {
	pIdStr := ctx.GetPathVariable("pid")
	pid, _ := strconv.Atoi(pIdStr)
	post, err := service.GetPostById(pid)
	if err != nil {
		common.ErrorResult(ctx.W, errors.New("GetPost查询出错"))
		return
	}
	common.SuccessResult(ctx.W, post)
}

func (*Api) SaveAndUpdatePostNew(ctx *context.MsContext) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		log.Printf("SaveAndUpdatePost的token为空")
		return
	}
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.ErrorResult(ctx.W, errors.New("token解析失败"))
		return
	}
	uid := claim.Uid

	r := ctx.Request
	w := ctx.W
	method := r.Method
	switch method {
	case http.MethodPost:
		params, err := common.GetRequestJsonParam(r)
		if err != nil {
			log.Printf("SaveAndUpdatePost的POST解析请求参数异常：%v", err)
			return
		}
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := 0
		if params["type"] != nil {
			postType, _ = params["type"].(int)
		}
		post := &models.Post{
			CategoryId: categoryId,
			Content:    content,
			Markdown:   markdown,
			Slug:       slug,
			Title:      title,
			Type:       postType,
			UserId:     uid,
			Pid:        -1,
			ViewCount:  0,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.SuccessResult(w, post)
	case http.MethodPut:
		params, err := common.GetRequestJsonParam(r)
		if err != nil {
			log.Printf("SaveAndUpdatePost的PUT解析请求参数异常：%v", err)
			return
		}
		pid := params["pid"].(float64)
		pId := int(pid)
		post, _ := service.GetPostById(pId)
		if post == nil {
			common.ErrorResult(w, errors.New("SaveAndUpdatePost查询不到文章"))
			return
		}
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := 0
		if params["type"] != nil {
			postType, _ = params["type"].(int)
		}
		post.CategoryId = categoryId
		post.Content = content
		post.Markdown = markdown
		post.Slug = slug
		post.Title = title
		post.Type = postType
		post.UpdateAt = time.Now()
		service.UpdatePost(post)
		common.SuccessResult(w, post)
	}

}
