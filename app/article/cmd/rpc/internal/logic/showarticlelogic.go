package logic

import (
	"context"

	"backend-learning/blog/app/article/cmd/rpc/article"
	"backend-learning/blog/app/article/cmd/rpc/internal/svc"

	"manlu.org/tao/core/logx"
)

type ShowArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShowArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowArticleLogic {
	return &ShowArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShowArticleLogic) ShowArticle(in *article.ShowArticleReq) (*article.ShowArticleResp, error) {
	// todo: add your logic here and delete this line

	return &article.ShowArticleResp{}, nil
}
