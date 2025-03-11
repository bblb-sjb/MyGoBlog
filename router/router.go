package router

import (
	"myWeb/api"
	"myWeb/views"
	"net/http"
)

func Router() {
	//http.Handle("/", context.Context)
	//
	//context.Context.Handler("/", views.HTML.IndexNew)
	//context.Context.Handler("/c/{cid}", views.HTML.CategoryNew)
	//context.Context.Handler("/login", views.HTML.LoginNew)
	//context.Context.Handler("/p/{pid}.html", views.HTML.DetailNew)
	//context.Context.Handler("/writing", views.HTML.WritingNew)
	//context.Context.Handler("/pigeonhole", views.HTML.PigeonholeNew)
	//context.Context.Handler("/golang", views.HTML.SlugNew)
	//
	//context.Context.Handler("/api/v1/post", api.API.SaveAndUpdatePostNew)
	////context.Context.Handler("/api/v1/post/search/val={value}", api.API.SearchNew)
	//context.Context.Handler("/api/v1/post/{pid}", api.API.GetPostNew)
	//context.Context.Handler("/api/v1/login", api.API.LoginNew)
	//context.Context.Handler("/api/v1/upload/oss", api.API.UploadImageNew)

	//1. 页面  views 2. api 数据（json） 3. 静态资源
	//views
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login/", views.HTML.Login)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)

	//apis
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/upload/oss", api.API.UploadImage)
	http.HandleFunc("/api/v1/post/search", api.API.Search)

	//静态资源
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
