package e

//MsgFlags message definitions
var MsgFlags = map[int]string{
	Success:                    "ok",
	Error:                      "fail",
	InvalidParams:              "Invalid parameters",
	ErrorAuthCheckTokenFail:    "Authentication Token Failed",
	ErrorAuthCheckTokenTimeout: "Authentication Token Expired",
}

//GetMsg get error information based on code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
