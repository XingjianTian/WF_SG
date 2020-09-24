package middlewares

import (
	"WF_SG/Web/common"
	"github.com/kataras/iris/context"
)

func SessionLoginAuth(Ctx context.Context) {
	if auth := common.SessManager.Start(Ctx).Get("user_session"); auth == nil {
		Ctx.Redirect("/login")
		return
	}
	Ctx.Next()
}
