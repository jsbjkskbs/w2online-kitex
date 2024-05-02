package handler_video

import (
	"context"
	"sync"
	"work/kitex_gen/video"
	"work/pkg/constants"
	"work/pkg/errno"
	"work/rpc/facade/handlers"
	"work/rpc/facade/handlers/handler_video/convert"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/model/base"
	facade_video "work/rpc/facade/model/base/video"

	"github.com/cloudwego/hertz/pkg/app"
)

func VideoSearch(ctx context.Context, c *app.RequestContext) {
	var facadeReq facade_video.VideoSearchRequest
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		if _, exist := c.Get(`keywords`); !exist {
			facadeReq.Keywords = constants.ESNoKeywordsFlag
		}
		wg.Done()
	}()
	go func() {
		if _, exist := c.Get(`page_num`); !exist {
			facadeReq.PageNum = constants.ESNoPageParamFlag
		}
		wg.Done()
	}()
	go func() {
		if _, exist := c.Get(`page_size`); !exist {
			facadeReq.PageSize = constants.ESNoPageParamFlag
		}
		wg.Done()
	}()
	go func() {
		if _, exist := c.Get(`from_date`); !exist {
			facadeReq.FromDate = constants.ESNoTimeFilterFlag
		}
		wg.Done()
	}()
	go func() {
		if _, exist := c.Get(`to_date`); !exist {
			facadeReq.ToDate = constants.ESNoTimeFilterFlag
		}
		wg.Done()
	}()
	go func() {
		if _, exist := c.Get(`username`); !exist {
			facadeReq.Username = constants.ESNoUsernameFilterFlag
		}
		wg.Done()
	}()
	wg.Wait()

	data, err := client.VideoSearch(ctx, &video.VideoSearchRequest{
		Keywords: facadeReq.Keywords,
		PageNum:  facadeReq.PageNum,
		PageSize: facadeReq.PageSize,
		FromDate: facadeReq.FromDate,
		ToDate:   facadeReq.ToDate,
		Username: facadeReq.Username,
	})
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_video.VideoSearchResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
		Data: &facade_video.VideoSearchResponse_VideoSearchResponseData{
			Items: *convert.KitexGenToRespVideo(&data.Items),
			Total: data.Total,
		},
	})
}
