package article

import (
	"context"

	"backend-learning/blog/app/article/cmd/api/internal/svc"
	"backend-learning/blog/app/article/cmd/api/internal/types"

	"manlu.org/tao/core/logx"
)

type ShowArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShowArticleLogic {
	return ShowArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowArticleLogic) ShowArticle(req types.ShowArticleReq) (resp *types.ShowArticleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
