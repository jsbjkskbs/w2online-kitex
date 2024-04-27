package main

import (
	"work/rpc/facade/infras/client"
	"work/rpc/facade/router"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/pprof"
)

func main() {
	client.Init()

	h := server.Default(server.WithHostPorts(`:10001`))
	pprof.Register(h)
	router.Register(h)

	h.Spin()
}
