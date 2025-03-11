package server

import (
	"log"
	"myWeb/router"
	"net/http"
)

var App = &Server{}

type Server struct {
}

func (*Server) Start(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
