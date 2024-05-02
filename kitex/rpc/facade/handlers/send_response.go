package handlers

import (
	"work/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type _Base struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
type Response struct {
	Base _Base `json:"base"`
}
type ResponseWithExtraMsg struct {
	Base _Base       `json:"base"`
	Data interface{} `json:"data"`
}

func SendResponse(c *app.RequestContext, err error, extra *map[string]interface{}) {
	errCopy := errno.Convert(err)
	if extra == nil {
		c.JSON(consts.StatusOK, Response{
			Base: _Base{
				Code: errCopy.Code,
				Msg:  errCopy.Message,
			},
		})
		return
	}

	c.JSON(consts.StatusOK, ResponseWithExtraMsg{
		Base: _Base{
			Code: errCopy.Code,
			Msg:  errCopy.Message,
		},
		Data: extra,
	})
}

func SendFormedResponse(c *app.RequestContext, resp interface{}) {
	c.JSON(consts.StatusOK, resp)
}
