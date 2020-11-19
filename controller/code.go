package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidToken
	CodeNeedLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "Success",
	CodeInvalidParam:    "Wrong Parameter",
	CodeUserExist:       "User Exist",
	CodeUserNotExist:    "User Not Exist",
	CodeInvalidPassword: "Wrong Username or Password",
	CodeServerBusy:      "Server Busy",

	CodeInvalidToken: "Invalid Token",
	CodeNeedLogin:    "Need Login",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
