package logic

import (
	"context"
	"fmt"

	"backend-learning/blog/app/tag/cmd/rpc/internal/svc"
	"backend-learning/blog/app/tag/cmd/rpc/tag"

	"manlu.org/tao/core/logx"
)

type GetTagInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagInfoLogic {
	return &GetTagInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTagInfoLogic) GetTagInfo(in *tag.GetTagInfoReq) (*tag.GetTagInfoResp, error) {
	// todo: add your logic here and delete this line
	tagList := make([]*tag.TagModel, 0)

	for i := 0; i < 2; i++ {
		tm := &tag.TagModel{
			Id:   int64(i + 1),
			Name: fmt.Sprintf("tag%d", i+1),
		}
		tagList = append(tagList, tm)
	}
	return &tag.GetTagInfoResp{Tag: tagList}, nil
}
