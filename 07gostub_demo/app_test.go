package gostub_demo

import (
	"testing"

	"github.com/prashantv/gostub"
)

func TestGetConfig(t *testing.T) {
	// 为全局变量configFile打桩, 给发赋值一个指定文件
	stubs := gostub.Stub(&configFile, "./test.toml")
	defer stubs.Reset() // 测试结束重置

	// 测试代码
	data, err := GetConfig()
	if err != nil {
		t.Fatal()
	}
	// 返回的data的内容就是上面test.config文件的内容
	t.Logf("data: %s\n", data)
}

func TestShowNumBer(t *testing.T) {
	stubs := gostub.Stub(&maxNum, 20)
	defer stubs.Reset()
	// 测试代码
	res := ShowNumBer()
	if res != 20 {
		t.Fatal()
	}
	t.Logf("res: %d", res)
}
