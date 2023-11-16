package utils

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
)

var Client *redis.Client
var Ctx context.Context

// ExpireTime 过期时间
const ExpireTime = 60 * 60 * time.Second

// Set 设置缓存
func Set(key string, value interface{}) {
	valueJson, _ := json.Marshal(value)
	err := Client.Set(context.Background(), key, valueJson, ExpireTime).Err()
	if err != nil {
		logrus.Error("redis set 失败。-- err: ", err)
	}
}

/**
 * @Description: 获取缓存
 * @param key string 键
 */
func Get(key string) (bool, interface{}) {
	var data interface{}
	val, err := Client.Get(context.Background(), key).Result()
	if err != nil {
		logrus.Error("redis get 失败。-- err: ", err)
		return false, nil
	}
	// 反序列化
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		logrus.Error("redis get 反序列化失败。-- err: ", err)
		return false, nil
	}
	return true, data
}

// 删除缓存
func Del(key string) {
	// 删除已 key 为前缀的所有缓存
	key = key + "*"
	result, err := Client.Keys(context.Background(), key).Result()
	logrus.Info("result: ", result)
	if err != nil {
		logrus.Error("redis del 失败。-- err: ", err)
	}
	for _, v := range result {
		Client.Del(context.Background(), v)
	}
}

// getKey
func GetKey(tableName string, q interface{}) string {
	key := tableName
	marshal, err := json.Marshal(q)
	if err != nil {
		return key
	}
	key += "-" + string(marshal)
	return key
}
