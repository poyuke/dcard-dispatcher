package controller

import (
	"dispatcher/pkg/dao"
	"go.uber.org/zap"
)

type Env struct {
	Logger *zap.Logger
	DAO    *dao.Env
}