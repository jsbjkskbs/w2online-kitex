package conf

var (
	EtcdAddress = "localhost:2379"

	FacadeServiceName   = "work.facade"
	InteractServiceName = "work.interact"
	RelationServiceName = "work.relation"
	UserServiceName     = "work.user"
	VideoServiceName    = "work.video"
	MessageServiceName  = "work.message"

	InteractServiceAddress = "localhost:8888"
	RelationServiceAddress = "localhost:8889"
	UserServiceAddress     = "localhost:8890"
	VideoServiceAddress    = "localhost:8891"
	MessageServiceAddress  = "localhost:8892"

	ExportEndpoint = `localhost:4317`
)
