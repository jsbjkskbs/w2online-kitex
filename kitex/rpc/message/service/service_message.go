package service

import (
	"context"
	"work/kitex_gen/message"
	"work/pkg/errno"
	"work/rpc/message/dal/db"
	"work/rpc/message/service/convert"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{
		ctx: ctx,
	}
}

func (s *MessageService) PushMessage(request *message.InsertMessageRequest) error {
	if err := db.InsertMessage(request.Message.FromUid, request.Message.ToUid, request.Message.Content); err != nil {
		return errno.MySQLError
	}
	return nil
}

func (s *MessageService) PopMessage(request *message.PopMessageRequest) (*message.PopMessageResponseData, error) {
	if data, err := db.PopMessage(request.Uid); err != nil {
		return nil, errno.MySQLError
	} else {
		return &message.PopMessageResponseData{
			Items: *convert.DBRespToKitexGen(data),
		}, nil
	}
}
