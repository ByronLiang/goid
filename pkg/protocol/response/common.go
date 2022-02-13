package response

type RespInfo struct {
	Code ErrorCode   `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(data interface{}) RespInfo {
	return RespInfo{Code: CodeOK, Msg: "success", Data: data}
}

func Error(errCode ErrorCode, errMsg string) RespInfo {
	if errMsg == "" {
		errMsg = errCode.String()
	} else {
		errMsg = errCode.String() + ", " + errMsg
	}
	return RespInfo{Code: errCode, Msg: errMsg}
}

const (
	CodeOK = 0
	// 服务器的错误
	CodeInternalError = 10
	CodeDBError       = 11

	// 用户的错误
	CodeInvalidToken             = 12
	CodeRequestHeaderParamsError = 13
	CodeSignatureError           = 14
	CodeInvalidRequestParams     = 15
	CodeRepeatSubmitError        = 16

	// 业务错
)

type ErrorCode int

func (err ErrorCode) String() string {
	switch err {
	case CodeOK:
		return "success"
	case CodeInvalidRequestParams:
		return "invalid request params"
	case CodeInternalError:
		return "unknown internal error"
	case CodeInvalidToken:
		return "invalid token"
	case CodeSignatureError:
		return "signature error"
	case CodeDBError:
		return "db error"
	case CodeRepeatSubmitError:
		return "form repeat submit"
	}
	return "error"
}
