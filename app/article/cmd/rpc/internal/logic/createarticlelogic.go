package logic

import (
	"context"

	"backend-learning/blog/app/article/cmd/rpc/article"
	"backend-learning/blog/app/article/cmd/rpc/internal/svc"

	"manlu.org/tao/core/logx"
)

type CreateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  create article
func (l *CreateArticleLogic) CreateArticle(in *article.CreateArticleReq) (*article.CreateArticleResp, error) {
	// todo: add your logic here and delete this line

	return &article.CreateArticleResp{}, nil
}
