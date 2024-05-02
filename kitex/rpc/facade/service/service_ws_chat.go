package service

import (
	"context"
	"time"
	"work/kitex_gen/message"
	"work/kitex_gen/user"
	"work/pkg/errno"
	"work/pkg/utils"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
)

type ChatService struct {
	ctx  context.Context
	c    *app.RequestContext
	conn *websocket.Conn
}

type _user struct {
	username string
	conn     *websocket.Conn
	rsa      *utils.RsaService
}

var userMap = make(map[string]*_user)

func NewChatService(ctx context.Context, c *app.RequestContext, conn *websocket.Conn) *ChatService {
	return &ChatService{
		ctx:  ctx,
		c:    c,
		conn: conn,
	}
}

func (service ChatService) Login() error {
	uid, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return err
	}
	user, err := client.UserInfo(service.ctx, &user.UserInfoRequest{UserId: uid})
	if err != nil {
		return err
	}
	rsaClientKey := service.c.GetHeader(`rsa_public_key`)
	r := utils.NewRsaService()
	if err := r.Build(rsaClientKey); err != nil {
		hlog.Info(err)
		return errno.ServiceError
	}
	userMap[uid] = &_user{conn: service.conn, username: user.Username, rsa: r}
	publicKey, err := r.GetPublicKeyPemFormat()
	if err != nil {
		return errno.ServiceError
	}
	if err := service.conn.WriteMessage(websocket.TextMessage, []byte(publicKey)); err != nil {
		return errno.ServiceError
	}

	return nil
}

func (service ChatService) Logout() error {
	uid, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return err
	}
	userMap[uid] = nil
	return nil
}

func (service ChatService) SendMessage(content []byte) error {
	from, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return errno.TokenInvailed
	}
	to := service.c.Query(`to_user_id`)
	user, err := client.UserInfo(service.ctx, &user.UserInfoRequest{UserId: to})
	if err != nil {
		return errno.ServiceError
	}
	if user == nil {
		return errno.InfomationNotExist
	}
	toConn := userMap[to]
	fromConn := userMap[from]
	switch toConn {
	case nil: // 离线
		{
			plainText, err := fromConn.rsa.Decode(content)
			if err != nil {
				return errno.ServiceError
			}
			if err := client.InsertMessage(context.Background(), &message.InsertMessageRequest{
				Message: &message.MessageInfo{
					FromUid: from,
					ToUid:   to,
					Content: string(plainText),
				},
			}); err != nil {
				return errno.RedisError
			}
		}
	default: // 在线
		{
			plainText, err := fromConn.rsa.Decode(content)
			if err != nil {
				return errno.ServiceError
			}
			ciphertext, err := toConn.rsa.Encode(userinfoAppend(plainText, from))
			if err != nil {
				return errno.ServiceError
			}
			err = toConn.conn.WriteMessage(websocket.BinaryMessage, ciphertext)
			if err != nil {
				return errno.ServiceError
			}
		}
	}
	return nil
}

func (service ChatService) ReadOfflineMessage() error {
	uid, err := jwt.CovertJWTPayloadToString(service.ctx, service.c)
	if err != nil {
		return errno.TokenInvailed
	}
	data, err := client.PopMessage(context.Background(), &message.PopMessageRequest{
		Uid: uid,
	})
	if err != nil {
		return errno.ServiceError
	}
	toConn := userMap[uid]
	for _, item := range (*data).Items {
		ciphertext, err := toConn.rsa.Encode(userinfoAppend([]byte(item.Content), item.FromUid))
		if err != nil {
			return errno.ServiceError
		}
		err = service.conn.WriteMessage(websocket.BinaryMessage, ciphertext)
		if err != nil {
			return errno.ServiceError
		}
	}
	return nil
}
func userinfoAppend(rawText []byte, from string) []byte {
	return []byte(utils.ConvertTimestampToStringDefault(time.Now().Unix()) + ` [` + from + `]: ` + string(rawText))
}
