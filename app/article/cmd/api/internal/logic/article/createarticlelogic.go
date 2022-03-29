package article

import (
	"context"

	"backend-learning/blog/app/article/cmd/api/internal/svc"
	"backend-learning/blog/app/article/cmd/api/internal/types"

	"manlu.org/tao/core/logx"
)

type CreateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateArticleLogic {
	return CreateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateArticleLogic) CreateArticle(req types.CreateArticleReq) (resp *types.CreateArticleResp, err error) {
	resp = &types.CreateArticleResp{Success: true}
	return
}
