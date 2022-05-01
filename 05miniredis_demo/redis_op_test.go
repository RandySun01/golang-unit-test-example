package miniredis_demo

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/alicebob/miniredis/v2"
)

/*
@author RandySun
@create 2022-05-01-11:42
*/

func TestDoSomethingWithRedis(t *testing.T) {
	// mock 一个redis server
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	defer s.Close()
	// 准备数据
	s.Set("RandySun", "www.cnblogs.com/randysun/")
	s.SAdd(KeyValidWebSite, "RandySun")
	// 连接mock的redis server
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(), // mock redis server的地址
	})

	// 调用函数
	ok := DoSomethingWithRedis(rdb, "RandySun")
	if !ok {
		t.Fatal()
	}

	// 检查redis中的只是否复合预期
	if got, err := s.Get("blog"); err != nil || got != "https://www.cnblogs.com/randysun/" {
		t.Fatalf("'blog' has the wrong value")
	}

	// 使用帮助工具检查
	s.CheckGet(t, "blog", "https://www.cnblogs.com/randysun/")

	// 过期检查
	s.FastForward(5 * time.Second) // 快进五秒
	if s.Exists("blog") {
		t.Fatal("'blog' should not have existed any more")
	}
}
