package main

import (
	"context"
	"work/kitex_gen/base"
	message "work/kitex_gen/message"
	"work/pkg/errno"
	"work/pkg/utils"
	"work/rpc/message/service"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// InsertMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) InsertMessage(ctx context.Context, request *message.InsertMessageRequest) (resp *message.InsertMessageResponse, err error) {
	// TODO: Your code here...
	resp = new(message.InsertMessageResponse)

	err = service.NewMessageService(ctx).PushMessage(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	return resp, nil
}

// PopMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) PopMessage(ctx context.Context, request *message.PopMessageRequest) (resp *message.PopMessageResponse, err error) {
	// TODO: Your code here...
	resp = new(message.PopMessageResponse)
	resp.Data = &message.PopMessageResponseData{}

	data, err := service.NewMessageService(ctx).PopMessage(request)
	if err != nil {
		respErr := utils.CreateBaseHttpResponse(err)
		resp.Base = &base.Status{
			Code: respErr.StatusCode,
			Msg:  respErr.StatusMsg,
		}
		return resp, nil
	}

	resp.Base = &base.Status{
		Code: errno.NoError.Code,
		Msg:  errno.NoError.Message,
	}
	resp.Data = data
	return resp, nil
}
