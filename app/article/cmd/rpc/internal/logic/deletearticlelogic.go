package logic

import (
	"context"

	"backend-learning/blog/app/article/cmd/rpc/article"
	"backend-learning/blog/app/article/cmd/rpc/internal/svc"

	"manlu.org/tao/core/logx"
)

type DeleteArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteArticleLogic) DeleteArticle(in *article.DeleteArticleReq) (*article.DeleteArticleResp, error) {
	// todo: add your logic here and delete this line

	return &article.DeleteArticleResp{}, nil
}
