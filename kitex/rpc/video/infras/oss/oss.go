package oss

import (
	"context"
	"os"

	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadVideo(path, vid string) (string, error) {
	recoder, err := storage.NewFileRecorder(os.TempDir())
	if err != nil {
		return ``, err
	}
	ret := storage.PutRet{}
	putExtra := storage.RputV2Extra{
		Recorder: recoder,
	}
	err = resumeUploader.PutFile(
		context.Background(),
		&ret,
		upToken,
		`video/`+vid+`/video.mp4`,
		path,
		&putExtra,
	)
	if err != nil {
		return ``, err
	}
	return Url + `/video/` + vid + `/video.mp4`, nil
}

func UploadVideoCover(path, vid string) (string, error) {
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.PutFile(
		context.Background(),
		&ret,
		upToken,
		`video/`+vid+`/cover.jpg`,
		path,
		&putExtra,
	)
	if err != nil {
		return ``, err
	}
	return Url + `/video/` + vid + `/cover.jpg`, nil
}
