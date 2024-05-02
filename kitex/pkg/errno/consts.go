package errno

const (
	noErrorCode = 0

	_ = iota + 9999

	serviceErrCode
	authErrCode
	requestErrCode

	mySQLErrCode
	redisErrCode
	elasticErrCode
	ossErrCode
	rabbitMQErrCode

	rpcErrCode

	tokenInvailedErrCode
	infomationNotExistErrCode
	infomationAlreadyExistErrCode
	totpAuthErrCode
	dataProcessErrCode
)

const (
	noErrorMsg = `success`

	serviceErrMsg = `service failed`
	authErrMsg    = `authenticate failed`
	requestErrMsg = `request wrong`

	mySQLErrMsg    = `something wrong with MySQL`
	redisErrMsg    = `something wrong with Redis`
	elasticErrMsg  = `something wrong with ElasticSearch`
	ossErrMsg      = `something wrong with OSS`
	rabbitMQErrMsg = `something wrong with RabbitMQ`

	rpcErrMsg = `something wrong with rpc`

	tokenInvailedErrMsg          = `token is invailed`
	infomationNotExistErrMsg     = `infomation not exist`
	infomationAlreadyExistErrMsg = `infomation already exist`
	totpAuthErrMsg               = `code wrong or expired`
	dataProcessErrMsg            = `failed to solve`
)
