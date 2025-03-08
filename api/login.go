package api

import (
	"log"
	"myWeb/common"
	"myWeb/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	params, err := common.GetRequestJsonParam(r)
	if err != nil {
		log.Printf("解析请求参数异常：%v", err)
		return
	}
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	data, err := service.Login(userName, passwd)
	if err != nil {
		log.Printf("登录异常：%v", err)
		common.ErrorResult(w, err)
		return
	}
	common.SuccessResult(w, data)
}
