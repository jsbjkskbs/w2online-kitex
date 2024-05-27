package errno

import (
	"errors"
	"fmt"
)

type Errno struct {
	Code    int64
	Message string
}

func (err Errno) Error() string {
	return fmt.Sprintf("%v, Code:%v", err.Message, err.Code)
}

func (err Errno) WithMessage(msg string) Errno {
	return Errno{
		Code:    err.Code,
		Message: msg,
	}
}

func NewErrorMessage(code int64, msg string) Errno {
	return Errno{
		Code:    code,
		Message: msg,
	}
}

// 自定义错误
var (
	BaseError = Errno{Code: serviceErrCode}

	ServiceError                = Errno{Code: serviceErrCode, Message: serviceErrMsg}
	TokenInvailed               = Errno{Code: tokenInvailedErrCode, Message: tokenInvailedErrMsg}
	InfomationNotExist          = Errno{Code: infomationNotExistErrCode, Message: infomationNotExistErrMsg}
	InfomationAlreadyExistError = Errno{Code: infomationAlreadyExistErrCode, Message: infomationAlreadyExistErrMsg}
	TOTPAuthenticatedFailed     = Errno{Code: authErrCode, Message: totpAuthErrMsg}
	DataProcessFailed           = Errno{Code: dataProcessErrCode, Message: dataProcessErrMsg}
	RequestError                = Errno{Code: requestErrCode, Message: requestErrMsg}
	RPCError                    = Errno{Code: rpcErrCode, Message: rpcErrMsg}
	NoError                     = Errno{Code: noErrorCode, Message: noErrorMsg}

	MySQLError    = Errno{Code: mySQLErrCode, Message: mySQLErrMsg}
	RedisError    = Errno{Code: redisErrCode, Message: redisErrMsg}
	ElasticError  = Errno{Code: elasticErrCode, Message: elasticErrMsg}
	RabbitMQError = Errno{Code: rabbitMQErrCode, Message: rabbitMQErrMsg}
	OSSError      = Errno{Code: ossErrCode, Message: ossErrMsg}
	MilvusError   = Errno{Code: milvusErrCode, Message: milvusErrMsg}
)

// 提供转换方法
func Convert(err error) Errno {
	var e Errno
	if errors.As(err, &e) {
		return e
	}

	s := BaseError
	s.Message = err.Error()
	return s
}
