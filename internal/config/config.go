package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Rest rest.RestConf // RestConf 里已经包含了 LogConf
	Auth struct {
		AccessSecret string `json:"AccessSecret"`
		AccessExpire int64  `json:"AccessExpire"`
	}
	UserRpc     zrpc.RpcClientConf `json:"UserRpc"`
	CustomerRpc zrpc.RpcClientConf `json:"CustomerRpc"`
	Mysql       struct {
		DataSource string `json:"DataSource"`
	}
	// CacheRedis cache.CacheConf `json:"CacheRedis"`
}
