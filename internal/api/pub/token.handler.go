package pub

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/openapi/internal/server"
)

func TokenHandler(ctx iris.Context) {
	var err error

	err = server.GetServer().HandleTokenRequest(ctx.ResponseWriter(), ctx.Request())
	if err != nil {
		// TODO resp <- 401
	}
}

