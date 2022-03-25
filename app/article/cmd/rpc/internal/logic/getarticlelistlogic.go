package logic

import (
	"context"

	"backend-learning/blog/app/article/cmd/rpc/article"
	"backend-learning/blog/app/article/cmd/rpc/internal/svc"

	"manlu.org/tao/core/logx"
)

type GetArticleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleListLogic) GetArticleList(in *article.GetArticleListReq) (*article.GetArticleListResp, error) {
	// todo: add your logic here and delete this line

	return &article.GetArticleListResp{}, nil
}
