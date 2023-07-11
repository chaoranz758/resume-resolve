package code

const (
	CodeSuccess = 1000 + iota
	CodeInternal
	CodeDeliveryOverflow
	CodeDeliveryPostProcessing
	CodePostProcessing
)

var codeMsgMap = map[int]string{
	CodeSuccess:                "success",
	CodeInternal:               "internal error",
	CodeDeliveryOverflow:       "您已投递岗位，请耐心等待",
	CodeDeliveryPostProcessing: "您投递的简历正在处理或已经处理完成，请不要取消投递",
	CodePostProcessing:         "您的简历正在处理，不能继续投递",
}

func GetMsg(code int) string {
	return codeMsgMap[code]
}
