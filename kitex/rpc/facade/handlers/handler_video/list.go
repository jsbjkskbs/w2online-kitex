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

func VideoList(ctx context.Context, c *app.RequestContext) {
	var (
		err       error
		facadeReq facade_video.VideoListRequest
		wg        sync.WaitGroup
		errChan   = make(chan error, 3)
	)
	if err = c.BindAndValidate(&facadeReq); err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}
	wg.Add(3)
	go func() {
		if _, exist := c.GetQuery(`page_num`); !exist {
			facadeReq.PageNum = constants.ESNoPageParamFlag
		}
		wg.Done()
	}()
	go func() {
		if _, exist := c.GetQuery(`page_size`); !exist {
			facadeReq.PageSize = constants.ESNoPageParamFlag
		}
		wg.Done()
	}()
	go func() {
		if _, exist := c.GetQuery(`user_id`); !exist {
			errChan <- errno.ServiceError
		}
		wg.Done()
	}()
	wg.Wait()
	select {
	case err := <-errChan:
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	default:
	}

	data, err := client.VideoList(ctx, &video.VideoListRequest{
		UserId:   facadeReq.UserId,
		PageNum:  facadeReq.PageNum,
		PageSize: facadeReq.PageSize,
	})
	if err != nil {
		handlers.SendResponse(c, errno.Convert(err), nil)
		return
	}

	handlers.SendFormedResponse(c, &facade_video.VideoListResponse{
		Base: &base.Status{
			Code: errno.NoError.Code,
			Msg:  errno.NoError.Message,
		},
		Data: &facade_video.VideoListResponse_VideoListResponseData{
			Items: *convert.KitexGenToRespVideo(&data.Data),
			Total: data.Total,
		},
	})
}
