package views

import (
	"myWeb/common"
	"myWeb/context"
	"myWeb/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole

	pigeonholeRes := service.FindPostByPigeonhole()

	pigeonhole.WriteData(w, pigeonholeRes)
}

func (*HTMLApi) PigeonholeNew(ctx *context.MsContext) {
	pigeonhole := common.Template.Pigeonhole

	pigeonholeRes := service.FindPostByPigeonhole()

	pigeonhole.WriteData(ctx.W, pigeonholeRes)
}
