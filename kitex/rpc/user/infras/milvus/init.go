package milvus

import (
	"context"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

var (
	mClient *client.Client
)

func Load() {
	if mClient != nil {
		(*mClient).Close()
	}

	c, err := client.NewClient(context.Background(), client.Config{
		Address: Address,
	})
	if err != nil {
		panic(err)
	}

	mClient = &c
}
