package handlers

import (
	"work/pkg/errmsg"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SendResponse(c *app.RequestContext, err error, data interface{}) {
	errCopy := errmsg.Convert(err)
	c.JSON(consts.StatusOK, Response{
		Code: errCopy.ErrorCode,
		Msg:  errCopy.ErrorMsg,
		Data: data,
	})
}
