package tag

import (
	"net/http"

	"backend-learning/blog/app/tag/cmd/api/internal/logic/tag"
	"backend-learning/blog/app/tag/cmd/api/internal/svc"
	"manlu.org/tao/rest/httpx"
)

func GetTagListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tag.NewGetTagListLogic(r.Context(), svcCtx)
		resp, err := l.GetTagList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
