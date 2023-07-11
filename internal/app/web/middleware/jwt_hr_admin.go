package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"resume-resolving/internal/app/web/pkg/code"
	"resume-resolving/internal/app/web/pkg/jwt"
)

func JWTHrAdmin() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		data, ok := ctx.Get(jwt.IdentityKey)
		if !ok {
			hlog.Error(errGetUserFromContext)
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"code":    code.CodeInvalidToken,
				"message": code.GetMsg(code.CodeInvalidToken),
			})
			ctx.Abort()
			return
		}
		user, ok := data.(*jwt.User)
		if !ok {
			hlog.Error(errCtxValue2User)
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"code":    code.CodeInvalidToken,
				"message": code.GetMsg(code.CodeInvalidToken),
			})
			ctx.Abort()
			return
		}
		if user.Role != adminRole && user.Role != hrRole {
			hlog.Error(errWrongRole)
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"code":    code.CodeInvalidToken,
				"message": code.GetMsg(code.CodeInvalidToken),
			})
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
