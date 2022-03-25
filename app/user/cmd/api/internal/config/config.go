package config

import "manlu.org/tao/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessExpire int64
		AccessSecret string
	}
}
