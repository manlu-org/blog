package svc

import (
	"backend-learning/blog/app/article/cmd/api/internal/config"
	"backend-learning/blog/app/reply/cmd/rpc/replyclient"
	"backend-learning/blog/app/tag/cmd/rpc/tagclient"
	"manlu.org/tao/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Tag    tagclient.Tag
	Reply  replyclient.Reply
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Tag:    tagclient.NewTag(zrpc.MustNewClient(c.Tag)),
		Reply:  replyclient.NewReply(zrpc.MustNewClient(c.Reply)),
	}
}
