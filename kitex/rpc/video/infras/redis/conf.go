package redis

type _Redis struct {
	Addr string
	Pwd  string
	DB   int
}

var (
	VideoInfo, VideoUpload _Redis
)
