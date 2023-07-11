package code

const (
	CodeSuccess = 1000 + iota
	CodeInternal
)

var codeMsgMap = map[int]string{
	CodeSuccess:  "success",
	CodeInternal: "internal error",
}

func GetMsg(code int) string {
	return codeMsgMap[code]
}
