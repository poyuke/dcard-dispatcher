package main

import (
	"fmt"
	"time"
	"os"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"

	"dispatcher/pkg/config"
	"dispatcher/pkg/controller"
	"dispatcher/pkg/dao"
	"dispatcher/pkg/log"
)

func main() {
	config.SetConfig()

	logger, err := log.InitLog()
	if err != nil {
		panic("Init log system fail")
	}
	defer logger.Sync()

	debug := viper.GetBool("debug")
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	redisConn := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: "",
		DB:       0,
	})

	err = redisConn.Ping().Err()
    if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	daoEnv := &dao.Env{
		Logger: logger,
		Redis:  redisConn,
	}

	controllerEnv := controller.Env{
		Logger: logger,
		DAO:    daoEnv,
	}

	dispatcherDao := dao.NewDAO(logger)
	dispatcherDao.Redis = redisConn

	r := gin.New()
	if !viper.GetBool("log.disable_json") {
		r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
		r.Use(ginzap.RecoveryWithZap(logger, true))
	} else {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}

	r.Use(cors.Default())

	controller.InitRouter(controllerEnv, r)

	port := viper.GetInt("port")
	r.Run(fmt.Sprintf(":%d", port)) // listen and serve on 0.0.0.0
}
