package setup

import (
	"context"
	"github.com/lishimeng/openapi/internal/server"
)

type componentHandler func(ctx context.Context) (err error)

func BeforeApp(ctx context.Context) (err error) {

	var comps []componentHandler
	comps = append(comps, initServer)

	for _, com := range comps {
		err = com(ctx)
		if err != nil {
			break
		}
	}
	return
}

func initServer(_ context.Context) (err error) {
	server.Init()
	return
}
