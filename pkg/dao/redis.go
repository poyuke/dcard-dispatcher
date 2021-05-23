package dao

import (
	"fmt"
	"time"
	"go.uber.org/zap"
	"github.com/speps/go-hashids/v2"
)

func NewDAO(logger *zap.Logger) *DAO {
	d = &DAO{Logger: logger}
	return d
}

func (e Env) SetRedis(fileContent string) (string, error) {
	hashId, err := e.HashId()
	if err != nil {
		e.Logger.Error("Hash Id Error", zap.String("SetRedis", "Hash Id"), zap.String("error", err.Error()))
		return "", err
	}
	data := make(map[string]interface{})
	data["fileContent"] = fileContent
	data["status"] = "waiting"
	data["sha"] = ""
	err = e.Redis.HMSet(hashId, data).Err()
	if err != nil {
		e.Logger.Error("Redis Error", zap.String("SetRedis", "set file"), zap.String("error", err.Error()))
		return "", err
	}
	return hashId, nil
}

func (e Env) GetRedisByKey(key string, field string) (string, error) {
	data, err := e.Redis.HMGet(key, field).Result()
	if err != nil {
		e.Logger.Error("Redis Error", zap.String("GetRedisByKey", "hmget key"), zap.String("error", err.Error()))
		return "", err
	}
	value := ""
	if data != nil && len(data) > 0 {
		value = fmt.Sprintf("%v", data[0])
	}
	return value, nil
}

func (e Env) GetRedisAllStatus() (map[string]interface{}, error) {
	keys, err := e.Redis.Keys("*").Result()
	if err != nil {
		e.Logger.Error("Redis Error", zap.String("GetRedisAllStatus", "get keys"), zap.String("error", err.Error()))
		return nil, err
	}
	jobs := make(map[string]interface{})
	for _, key := range keys {
		data, err := e.Redis.HMGet(key, "status").Result()
		if err != nil {
			e.Logger.Error("Redis Error", zap.String("GetRedisAllStatus", "hmget key"), zap.String("error", err.Error()))
			return nil, err
		}
		jobs[key] = data[0]
	}
	return jobs, nil
}

func (e Env) HashId() (string, error) {
	hdata := hashids.NewData()
	hdata.MinLength = 11
	hdata.Salt = "dcard dispatcher"

	hid, _ := hashids.NewWithData(hdata)

	epoch := time.Now().Unix()
	numbers := []int64{epoch}
	hash, err := hid.EncodeInt64(numbers)
	if err != nil {
		return "", err
	}
	return hash, nil
}
