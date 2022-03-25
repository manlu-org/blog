package user

import (
	"backend-learning/blog/app/user/cmd/api/internal/svc"
	"backend-learning/blog/app/user/cmd/api/internal/types"
	"context"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"

	"manlu.org/tao/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req types.UserLoginReq) (resp *types.UserLoginResp, err error) {

	if !(req.Username == "admin" && req.Password == "123456") {
		return &types.UserLoginResp{
			Success: false,
			Message: "username or password error",
		}, nil
	}

	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	now := time.Now().Unix()
	accessToken, err := l.GenToken(now, l.svcCtx.Config.Auth.AccessSecret, nil, accessExpire)
	if err != nil {
		return nil, err
	}

	return &types.UserLoginResp{
		Success:   true,
		Message:   "ok",
		Token:     accessToken,
		ExpiredAt: now + accessExpire,
	}, nil
}

func (l *UserLoginLogic) GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
