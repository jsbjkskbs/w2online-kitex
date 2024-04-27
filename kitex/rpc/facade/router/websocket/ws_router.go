package webs

import (
	handler_wschat "work/rpc/facade/handlers/handler_ws_chat"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func register(h *server.Hertz) {
	h.GET(`/`, append(_homeMW(), handler_wschat.Handler)...)
}
