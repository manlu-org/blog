package config

import (
	"manlu.org/tao/rest"
	"manlu.org/tao/zrpc"
)

type Config struct {
	rest.RestConf
	Tag zrpc.RpcClientConf
}
