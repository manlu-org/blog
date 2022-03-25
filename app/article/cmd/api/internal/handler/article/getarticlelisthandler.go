package article

import (
	"net/http"

	"backend-learning/blog/app/article/cmd/api/internal/logic/article"
	"backend-learning/blog/app/article/cmd/api/internal/svc"
	"backend-learning/blog/app/article/cmd/api/internal/types"
	"manlu.org/tao/rest/httpx"
)

func GetArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := article.NewGetArticleListLogic(r.Context(), svcCtx)
		resp, err := l.GetArticleList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
