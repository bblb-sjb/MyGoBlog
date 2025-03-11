package views

import (
	"myWeb/common"
	"myWeb/config"
	"myWeb/context"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}

func (*HTMLApi) LoginNew(ctx *context.MsContext) {
	login := common.Template.Login
	login.WriteData(ctx.W, config.Cfg.Viewer)
}
