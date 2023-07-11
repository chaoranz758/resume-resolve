package code

const (
	CodeSuccess = 1000 + iota
	CodeInternal
	CodeUsernameInputFailed
	CodePasswordInputFailed
	CodeUserExist
)

var codeMsgMap = map[int]string{
	CodeSuccess:             "success",
	CodeInternal:            "internal error",
	CodeUsernameInputFailed: "用户名不存在",
	CodePasswordInputFailed: "密码输入错误",
	CodeUserExist:           "用户已存在",
}

func GetMsg(code int) string {
	return codeMsgMap[code]
}
