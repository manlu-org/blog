package logic

import (
	"context"
	"fmt"

	"backend-learning/blog/app/tag/cmd/rpc/internal/svc"
	"backend-learning/blog/app/tag/cmd/rpc/tag"

	"manlu.org/tao/core/logx"
)

type GetTagListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagListLogic {
	return &GetTagListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTagListLogic) GetTagList(in *tag.GetTagListReq) (*tag.GetTagListResp, error) {
	tagList := make([]*tag.TagModel, 0)

	for i := 0; i < 3; i++ {
		tm := &tag.TagModel{
			Id:   int64(i + 1),
			Name: fmt.Sprintf("tag%d", i+1),
		}
		tagList = append(tagList, tm)
	}

	return &tag.GetTagListResp{Tags: tagList}, nil
}
