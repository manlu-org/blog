package config

import (
	"manlu.org/tao/rest"
	"manlu.org/tao/zrpc"
)

type Config struct {
	rest.RestConf
	Reply zrpc.RpcClientConf
	Tag   zrpc.RpcClientConf
}
