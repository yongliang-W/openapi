package pub

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/openapi/internal/server"
)

// AuthorizeHandler 认证(支持client_credential
func AuthorizeHandler(ctx iris.Context) {
	var err error

	err = server.GetServer().HandleAuthorizeRequest(ctx.ResponseWriter(), ctx.Request())
	if err != nil {
		log.Info("handle authorize request failed")
		log.Info(err)
		// TODO resp <- 401
	}
}


