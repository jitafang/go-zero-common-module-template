package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	JWT Jwt
}

type Jwt struct {
	Issuer              string
	AccessSecret        string
	AccessExpire        int
	AccessRefreshExpire int
}
