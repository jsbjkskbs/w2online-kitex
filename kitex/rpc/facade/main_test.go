package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"os"
	"testing"
	"work/pkg/utils"
	"work/rpc/facade/infras/client"
	"work/rpc/facade/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func TestFunc(t *testing.T) {
	// Any func here
	b := utils.NewRsaService()
	hlog.Info(b.Build([]byte(`MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6KlV8XKwncTCLoCiFDzpTObECUdSOkBFBrek/Pz0zDZ0G6Llf9NUXiA+3IvOQ2uvaIOVJBbwXYs3GAnhM3H+fWQ9AcbEcR+jXBm+yrqrVQzb7tANvy1V2w0+UZoVB5AKpOYURvCin1qU65X26Q2sg96mh2i+utwkDNMwmHdpVO4wTFsO4iwFPBJNhekM/+WleiyROQaEqUwY1Xxbwfv0GnsoqhpRM/yxDgtexrDKSmnXRWNsce/7ReyqCDLC9osJtDigCIOkUtYQ/6qs9tWg+jaAMQ4/KgnhFreJn6J0vykUXqIM2HTexgtC4nEKbsJaMVeg5u5Uqg5NJlyBUgZpgwIDAQAB`)))
	msg, _ := b.Encode([]byte(`Hello World`))
	hlog.Info(base64.StdEncoding.EncodeToString(msg))
	hlog.Info(b.GetPublicKeyPemFormat())
	hlog.Info(b.GetPrivateKeyPemFormat())
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	hlog.Info(input)
	hlog.Info(b.Decode([]byte(input)))
}

func hInit() *server.Hertz {
	client.Init()
	jwt.AccessTokenJwtInit()
	jwt.RefreshTokenJwtInit()
	h := server.Default(
		server.WithHostPorts(`:10001`),
	)
	return h
}

func BenchmarkPing(b *testing.B) {
	b.StopTimer()
	h := server.Default()
	b.StartTimer()
	ut.PerformRequest(h.Engine, "GET", "/ping", &ut.Body{Body: bytes.NewBufferString("1"), Len: 1},
		ut.Header{Key: "Connection", Value: "close"})
}

func TestUserRegister(t *testing.T) {
	h := hInit()
	req := `{
		"username":"test",
		"password":"123456789"
	}
	`
	ut.PerformRequest(h.Engine, "POST", "/user/register/", &ut.Body{Body: bytes.NewBufferString(req), Len: len(req)},
		ut.Header{Key: `Content-Type`, Value: `application/json`},
	)
}

func BenchmarkUserRegister(b *testing.B) {
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	req := `{
		"username":"test",
		"password":"123456789"
	}
	`
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "POST", "/user/register/", &ut.Body{Body: bytes.NewBufferString(req), Len: len(req)},
			ut.Header{Key: `Content-Type`, Value: `application/json`},
		)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkUserLoginNoMFA(b *testing.B) {
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	req := `{
		"username":"test",
		"password":"123456789"
	}`
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "POST", "/user/login/", &ut.Body{Body: bytes.NewBufferString(req), Len: len(req)},
			ut.Header{Key: `Content-Type`, Value: `application/json`},
		)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkUserLoginMFA(b *testing.B) {
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	req := `{
		"username":"cyk",
		"password":"123456789",
		"code":"123456"
	}`
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "POST", "/user/login/", &ut.Body{Body: bytes.NewBufferString(req), Len: len(req)},
			ut.Header{Key: `Content-Type`, Value: `application/json`},
		)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkUserInfo(b *testing.B) {
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5fZmllbGQiOnsiVWlkIjoiMTAwMDQifSwiZXhwIjoxNzExMDMxODg5LCJvcmlnX2lhdCI6MTcxMTAyODI4OX0.gCpVOMoKmDS4iEzzl2ZKlxKwM9ivBX8XBWqsoS5h1GQ`
	uid := `10005`
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "GET", `/user/info/?token=`+token+`&user_id=`+uid, nil)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkVideoFeed(b *testing.B) {
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "GET", `/video/feed/`, nil)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkVideoList(b *testing.B) {
	const (
		uid      = `10003`
		pagenum  = `0`
		pagesize = `10`
		token    = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5fZmllbGQiOnsiVWlkIjoiMTAwMDQifSwiZXhwIjoxNzExMDMxODg5LCJvcmlnX2lhdCI6MTcxMTAyODI4OX0.gCpVOMoKmDS4iEzzl2ZKlxKwM9ivBX8XBWqsoS5h1GQ`
	)
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "GET", `/video/feed/?user_id=`+uid+`&page_num=`+pagenum+`&page_size=`+pagesize, nil,
			ut.Header{Key: "Access-Token", Value: token},
		)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkVideoPopular(b *testing.B) {
	const (
		pagenum  = `0`
		pagesize = `10`
	)
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "GET", `/video/popular/?page_num=`+pagenum+`&page_size=`+pagesize, nil)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkVideoSearch(b *testing.B) {
	const (
		token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5fZmllbGQiOnsiVWlkIjoiMTAwMDQifSwiZXhwIjoxNzExMDMxODg5LCJvcmlnX2lhdCI6MTcxMTAyODI4OX0.gCpVOMoKmDS4iEzzl2ZKlxKwM9ivBX8XBWqsoS5h1GQ`
	)
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	req := `{
		"page_num":0,
		"page_size":10
	}`
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "POST", `/video/search/`, &ut.Body{Body: bytes.NewBufferString(req), Len: len(req)},
			ut.Header{Key: "Access-Token", Value: token},
			ut.Header{Key: `Content-Type`, Value: `application/json`},
		)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkLikeAction(b *testing.B) {
	const (
		token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5fZmllbGQiOnsiVWlkIjoiMTAwMDQifSwiZXhwIjoxNzExMTMwNzY3LCJvcmlnX2lhdCI6MTcxMTEyNzE2N30.dhkp9QS4HgCobTmIKy7mREEcEz0g8xkkmI_VqQUqy1A`
	)
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	req := `{
		"video_id":"10003",
		"action_type":"1"
	}`
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "POST", `/like/action/`, &ut.Body{Body: bytes.NewBufferString(req), Len: len(req)},
			ut.Header{Key: "Access-Token", Value: token},
			ut.Header{Key: `Content-Type`, Value: `application/json`},
		)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkLikeList(b *testing.B) {
	const (
		pagenum  = `0`
		pagesize = `10`
		uid      = `10004`
		token    = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5fZmllbGQiOnsiVWlkIjoiMTAwMDQifSwiZXhwIjoxNzExMDMxODg5LCJvcmlnX2lhdCI6MTcxMTAyODI4OX0.gCpVOMoKmDS4iEzzl2ZKlxKwM9ivBX8XBWqsoS5h1GQ`
	)
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "GET", `/like/list/?page_num=`+pagenum+`&page_size=`+pagesize+`&user_id=`+uid, nil,
			ut.Header{Key: "Access-Token", Value: token},
		)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkCommentList(b *testing.B) {
	const (
		pagenum  = `0`
		pagesize = `10`
		vid      = `10003`
		token    = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5fZmllbGQiOnsiVWlkIjoiMTAwMDQifSwiZXhwIjoxNzExMDMxODg5LCJvcmlnX2lhdCI6MTcxMTAyODI4OX0.gCpVOMoKmDS4iEzzl2ZKlxKwM9ivBX8XBWqsoS5h1GQ`
	)
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "GET", `/comment/list/?page_num=`+pagenum+`&page_size=`+pagesize+`&video_id=`+vid, nil,
			ut.Header{Key: "Access-Token", Value: token},
		)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkRelationAction(b *testing.B) {
	const (
		token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5fZmllbGQiOnsiVWlkIjoiMTAwMDQifSwiZXhwIjoxNzExMDM2NTY4LCJvcmlnX2lhdCI6MTcxMTAzMjk2OH0.Tcg4oZyUm-p3e-24OPpcsjqgM4frEmmhQ8HyItvU4tc`
	)
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	req := `{
		"to_user_id":"10003",
		"action_type":0
	}`
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "POST", `/relation/action/`, &ut.Body{Body: bytes.NewBufferString(req), Len: len(req)},
			ut.Header{Key: "Access-Token", Value: token},
			ut.Header{Key: `Content-Type`, Value: `application/json`},
		)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkFollowingList(b *testing.B) {
	const (
		uid   = `10003`
		token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5fZmllbGQiOnsiVWlkIjoiMTAwMDQifSwiZXhwIjoxNzExMDM2NTY4LCJvcmlnX2lhdCI6MTcxMTAzMjk2OH0.Tcg4oZyUm-p3e-24OPpcsjqgM4frEmmhQ8HyItvU4tc`
	)
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "GET", `/following/list/?user_id=`+uid, nil,
			ut.Header{Key: "Access-Token", Value: token},
		)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}

func BenchmarkFriendList(b *testing.B) {
	const (
		uid   = `10003`
		token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW5fZmllbGQiOnsiVWlkIjoiMTAwMDQifSwiZXhwIjoxNzExMDM2NTY4LCJvcmlnX2lhdCI6MTcxMTAzMjk2OH0.Tcg4oZyUm-p3e-24OPpcsjqgM4frEmmhQ8HyItvU4tc`
	)
	b.StopTimer()
	h := hInit()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "GET", `/friend/list/?user_id=`+uid, nil,
			ut.Header{Key: "Access-Token", Value: token},
		)
		assert.DeepEqual(b, consts.StatusOK, w.Result().StatusCode())
	}
}
