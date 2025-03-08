package common

import (
	"encoding/json"
	"log"
	"myWeb/config"
	"myWeb/models"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//耗时
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

func SuccessResult(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Printf("返回数据失败：%v", err)
		return
	}
}

func ErrorResult(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = 500
	result.Error = err.Error()
	result.Data = nil
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Printf("返回数据失败：%v", err)
		return
	}
}

func GetRequestJsonParam(r *http.Request) (map[string]interface{}, error) {
	var params map[string]interface{}
	// 使用 json.NewDecoder 来逐步解码请求体
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("解析请求参数失败：%v", err)
		return nil, err
	}
	return params, nil
}
