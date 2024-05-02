package oss

import (
	"bytes"
	"context"

	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadAvatar(data *[]byte, dataSize int64, uid string, tag string) (string, error) {
	deleteAvatar(uid)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	var suffix string
	switch tag {
	case `image/jpeg`, `image/jpg`:
		suffix = `.jpg`
	case `image/png`:
		suffix = `.png`
	}
	err := formUploader.Put(
		context.Background(),
		&ret,
		upToken,
		`avatar/`+uid+suffix,
		bytes.NewReader(*data),
		dataSize,
		&putExtra,
	)
	if err != nil {
		return ``, err
	}
	return Url + `/avatar/` + uid + suffix, nil
}

func deleteAvatar(uid string) {
	keys := []string{
		`avatar/` + uid + `.jpg`,
		`avatar/` + uid + `.jpeg`,
		`avatar/` + uid + `.png`,
	}
	for _, key := range keys {
		bucketManager.Delete(Bucket, key)
	}
}
