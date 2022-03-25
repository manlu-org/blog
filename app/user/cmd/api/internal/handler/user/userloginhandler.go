package user

import (
	"net/http"

	"backend-learning/blog/app/user/cmd/api/internal/logic/user"
	"backend-learning/blog/app/user/cmd/api/internal/svc"
	"backend-learning/blog/app/user/cmd/api/internal/types"
	"manlu.org/tao/rest/httpx"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
