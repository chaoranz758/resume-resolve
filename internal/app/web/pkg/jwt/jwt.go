package jwt

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"resume-resolving/api/idl/service/user/kitex_gen/user"
	"resume-resolving/internal/app/web/config"
	rpc2 "resume-resolving/internal/app/web/pkg/client/rpc"
	"resume-resolving/internal/app/web/pkg/code"
	"time"
)

var IdentityKey = "role"

type login struct {
	Username string `form:"username,required" json:"username,required"`
	Password string `form:"password,required" json:"password,required"`
	Role     int8   `form:"role,required" json:"role,required"`
}

type User struct {
	UserName string
	Role     int8
}

type JWT struct {
	config         *config.Config
	AuthMiddleware *jwt.HertzJWTMiddleware
}

func (j *JWT) Init(UserClient *rpc2.UserClient) error {
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte(j.config.ConfigInNacos.Jwt.Key),
		Timeout:     time.Duration(j.config.ConfigInNacos.Jwt.TimeOut) * time.Hour,
		MaxRefresh:  time.Duration(j.config.ConfigInNacos.Jwt.MaxRefresh) * time.Hour,
		IdentityKey: IdentityKey,
		//向token中添加自定义负载信息的函数
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					IdentityKey: v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		//获取身份信息的函数
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &User{
				Role: claims[IdentityKey].(int8),
			}
		},
		//登录时认证用户信息的函数
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals login
			if err := c.BindAndValidate(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			resp, err := UserClient.Client.UserLogin(ctx, &user.UserLoginRPCRequest{
				Username: loginVals.Username,
				Password: loginVals.Password,
				Role:     loginVals.Role,
			})

			if err != nil {
				//if resp.code ==0 => web client doesn't connect to rpc server
				if resp.Code == code.CodeRPCInternal || resp.Code == 0 {
					hlog.Error(code.GetMsg(code.CodeServerBusy), err)
					return code.GetMsg(code.CodeServerBusy), err
				} else {
					hlog.Error(code.GetMsg(int(resp.Code)), err)
					return resp.Message, err
				}
			}

			return nil, nil
		},
		//jwt校验通过后，在这个函数中通过role字段验证身份字段验证
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			//直接return true，后续再通过一个中间件从上下文中读出role值验证不同用户的身份
			return true
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code1 int, message string) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    code.CodeInvalidToken,
				"message": code.GetMsg(code.CodeInvalidToken),
			})
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code1 int, token string, expire time.Time) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    code1,
				"message": code.GetMsg(code.CodeSuccess),
				"token":   token,
				"expire":  expire.UnixMilli(),
			})
		},
		RefreshResponse: func(ctx context.Context, c *app.RequestContext, code1 int, token string, expire time.Time) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    code.CodeSuccess,
				"message": code.GetMsg(code.CodeSuccess),
				"token":   token,
				"expire":  expire.UnixMilli(),
			})
		},
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code1 int) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    code.CodeLoginout,
				"message": code.GetMsg(code.CodeLoginout),
			})
		},
	})
	if err != nil {
		hlog.Error(err)
		return err
	}
	err = authMiddleware.MiddlewareInit()

	if err != nil {
		hlog.Error("authMiddleware.MiddlewareInit() Error:" + err.Error())
		return err
	}
	j.AuthMiddleware = authMiddleware
	return nil
}

func NewJWT(config *config.Config) *JWT {
	return &JWT{
		config: config,
	}
}
