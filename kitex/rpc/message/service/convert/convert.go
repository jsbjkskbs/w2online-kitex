package convert

import (
	"work/kitex_gen/message"
	"work/rpc/message/dal/db"
)

func DBRespToKitexGen(items *[]db.Message) *[]*message.MessageInfo {
	result := make([]*message.MessageInfo, 0)
	for _, item := range *items {
		result = append(result, &message.MessageInfo{
			FromUid: item.FromUserId,
			ToUid:   item.ToUserId,
			Content: item.Content,
		})
	}
	return &result
}
