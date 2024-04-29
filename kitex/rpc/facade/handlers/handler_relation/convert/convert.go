package convert

import (
	"work/kitex_gen/base"
	resp "work/rpc/facade/model/base"
)

func KitexGenToRespUserLite(items *[]*base.UserLite) *[]*resp.UserLite {
	result := make([]*resp.UserLite, 0)
	for _, item := range *items {
		result = append(result, &resp.UserLite{
			Uid:       item.Uid,
			Username:  item.Username,
			AvatarUrl: item.AvatarUrl,
		})
	}
	return &result
}
