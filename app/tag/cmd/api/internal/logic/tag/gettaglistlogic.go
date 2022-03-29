package tag

import (
	"backend-learning/blog/app/tag/cmd/rpc/tagclient"
	"context"
	"manlu.org/tao/core/xcopy"

	"backend-learning/blog/app/tag/cmd/api/internal/svc"
	"backend-learning/blog/app/tag/cmd/api/internal/types"

	"manlu.org/tao/core/logx"
)

type GetTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetTagListLogic {
	return GetTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagListLogic) GetTagList() (resp *types.GetTagListResp, err error) {
	req := &tagclient.GetTagListReq{}

	tagLogic, err := l.svcCtx.Tag.GetTagList(l.ctx, req)
	if err != nil {
		return nil, err
	}

	resp = new(types.GetTagListResp)
	xcopy.Copy(&resp, tagLogic)

	return
}
