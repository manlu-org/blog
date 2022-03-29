package article

import (
	"context"

	"backend-learning/blog/app/article/cmd/api/internal/svc"
	"backend-learning/blog/app/article/cmd/api/internal/types"

	"manlu.org/tao/core/logx"
)

type GetArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetArticleListLogic {
	return GetArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleListLogic) GetArticleList(req types.PageReq) (resp *types.GetArticleListResp, err error) {

	return
}
