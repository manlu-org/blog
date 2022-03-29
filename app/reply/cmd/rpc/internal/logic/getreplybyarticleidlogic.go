package logic

import (
	"context"
	"fmt"
	"time"

	"backend-learning/blog/app/reply/cmd/rpc/internal/svc"
	"backend-learning/blog/app/reply/cmd/rpc/reply"

	"manlu.org/tao/core/logx"
)

type GetReplyByArticleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetReplyByArticleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReplyByArticleIdLogic {
	return &GetReplyByArticleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetReplyByArticleIdLogic) GetReplyByArticleId(in *reply.GetReplyByArticleIdReq) (*reply.GetReplyByArticleIdResp, error) {

	replyList := make([]*reply.ReplyModel, 0)

	for i := 0; i < 3; i++ {
		rm := &reply.ReplyModel{
			Id:        0,
			Content:   fmt.Sprintf("content: %d", i),
			CreatedAt: fmt.Sprintf("%s", time.Now()),
		}

		replyList = append(replyList, rm)
	}

	resp := &reply.GetReplyByArticleIdResp{Reply: replyList}

	return resp, nil
}
