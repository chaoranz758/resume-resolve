package code

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

const (
	CodeSuccess = 1000 + iota
	CodeRPCInternal
	CodeLoginout
	CodeInvalidParam
	CodeServerBusy
	CodeInvalidToken
	CodePasswordInputRepeat
	CodePasswordInputNotRepeat
)

var codeMsgMap = map[int]string{
	CodeSuccess:                "success",
	CodeLoginout:               "login out",
	CodeRPCInternal:            "rpc internal error",
	CodeInvalidParam:           "请求参数错误",
	CodeServerBusy:             "服务繁忙",
	CodeInvalidToken:           "token校验失败",
	CodePasswordInputRepeat:    "旧密码与原密码相同，请重新输入",
	CodePasswordInputNotRepeat: "两次输入密码不一致，请重新输入",
}

func GetMsg(code int) string {
	return codeMsgMap[code]
}

func RespWithJson(code int32, message string, c *app.RequestContext, err error) {
	if err != nil {
		hlog.Error("rpc service error", err)
		c.JSON(consts.StatusOK, map[string]interface{}{
			"code":    CodeServerBusy,
			"message": GetMsg(CodeServerBusy),
		})
		c.Abort()
		return
	}

	if code == CodeRPCInternal {
		hlog.Error("rpc service biz error")
		c.JSON(consts.StatusOK, map[string]interface{}{
			"code":    CodeServerBusy,
			"message": GetMsg(CodeServerBusy),
		})
		c.Abort()
		return
	}

	c.JSON(consts.StatusOK, map[string]interface{}{
		"code":    code,
		"message": message,
	})
	c.Abort()
	return
}

func RespWithJsonWithData(code int32, message string, c *app.RequestContext, data interface{}, err error) {
	if err != nil {
		hlog.Error("rpc service error", err)
		c.JSON(consts.StatusOK, map[string]interface{}{
			"code":    CodeServerBusy,
			"message": GetMsg(CodeServerBusy),
			"data":    nil,
		})
		c.Abort()
		return
	}

	if code == CodeRPCInternal {
		hlog.Error("rpc service biz error")
		c.JSON(consts.StatusOK, map[string]interface{}{
			"code":    CodeServerBusy,
			"message": GetMsg(CodeServerBusy),
			"data":    nil,
		})
		c.Abort()
		return
	}

	c.JSON(consts.StatusOK, map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	})
	c.Abort()
	return
}
