package oss

import (
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	formUploader  *storage.FormUploader
	bucketManager *storage.BucketManager
	resumeUploader *storage.ResumeUploaderV2
	mac           *auth.Credentials
	upToken       string
)

func Load() {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac = qbox.NewMac(AccessKey, SecretKey)
	upToken = putPolicy.UploadToken(mac)
	cfg := storage.Config{
		UseHTTPS:      false,
		UseCdnDomains: false,

		Zone: &storage.ZoneHuanan, //your bucket zone
	}
	formUploader = storage.NewFormUploader(&cfg)
	bucketManager = storage.NewBucketManager(mac, &cfg)
	resumeUploader = storage.NewResumeUploaderV2(&cfg)
}
