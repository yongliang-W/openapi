package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/openapi/internal/api/pub"
)

func Route(app *iris.Application) {
	oauth2(app.Party("/oauth2"))
	test(app.Party("/test"))
}

func oauth2(p iris.Party) {
	p.Get("/authorize", pub.AuthorizeHandler)
	p.Get("/token", pub.TokenHandler)
}

func test(p iris.Party) {
	p.Get("/logic", pub.AccessTokenVerifyHandler, pub.TestLogicApi)
}