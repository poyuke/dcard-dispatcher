package dao

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var d *DAO

type Env struct {
	Logger      *zap.Logger
	Redis       *redis.Client
}

// DAO 通用連線物件
type DAO struct {
	Logger      *zap.Logger
	Redis       *redis.Client
}
   