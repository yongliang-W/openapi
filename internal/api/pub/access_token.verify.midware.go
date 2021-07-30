package pub

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/openapi/internal/server"
)

func AccessTokenVerifyHandler(ctx iris.Context) {
	ti, err := server.GetServer().ValidationBearerToken(ctx.Request())
	if err != nil {
		log.Info(err)
		// TODO resp <- err
		data, _, headers := server.GetServer().GetErrorData(err)
		data["code"] = 401
		for key := range headers {
			ctx.Header(key, headers.Get(key))
		}
		_, _ = ctx.JSON(data)
		return
	}
	ctx.Values().Set("tokenKey", ti)
	ctx.Next()
}