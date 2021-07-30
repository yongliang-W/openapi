package pub

import "github.com/kataras/iris/v12"

func TestLogicApi(ctx iris.Context) {
	var resp = make(map[string]string)
	resp["code"] = "ok"
	resp["message"] = "this is a test"
	_, _ = ctx.JSON(resp)
}
