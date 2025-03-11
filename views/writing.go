package views

import (
	"myWeb/common"
	"myWeb/context"
	"myWeb/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}

func (*HTMLApi) WritingNew(ctx *context.MsContext) {
	writing := common.Template.Writing

	wr := service.Writing()
	writing.WriteData(ctx.W, wr)
}
