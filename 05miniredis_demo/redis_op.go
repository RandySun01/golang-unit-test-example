package miniredis_demo

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入版本
)

/*
@author RandySun
@create 2022-05-01-11:35
*/

const KeyValidWebSite = "app:valid:website:list"

func DoSomethingWithRedis(rdb *redis.Client, key string) bool {
	// 这里可以对redis操作的一些逻辑
	ctx := context.TODO()
	if !rdb.SIsMember(ctx, KeyValidWebSite, key).Val() {
		return false
	}
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	if !strings.HasPrefix(val, "https://") {
		val = "https://" + val
	}
	// 设置blog key 五秒过期
	if err := rdb.Set(ctx, "blog", val, 5*time.Second).Err(); err != nil {
		return false
	}
	return true
}
