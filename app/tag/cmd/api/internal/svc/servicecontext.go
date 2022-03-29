package svc

import (
	"backend-learning/blog/app/tag/cmd/api/internal/config"
	"backend-learning/blog/app/tag/cmd/rpc/tagclient"
	"manlu.org/tao/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Tag    tagclient.Tag
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Tag:    tagclient.NewTag(zrpc.MustNewClient(c.Tag)),
	}
}
