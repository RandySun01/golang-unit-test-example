package monkey_demo

import (
	"golang-unit-test-example/varys"
	"strings"
	"testing"

	"bou.ke/monkey"
)

/*
@author RandySun
@create 2022-05-01-17:37
*/
//
//  TestMyFunc
//  @Description: 为函数打桩
//  @param t
//
func TestMyFunc(t *testing.T) {
	// 对 varys.GetInfoByUID 进行打桩
	// 无论传入uid是多少,都返回 &varys.UserInfo{Name: "RandySun"}, nil

	monkey.Patch(varys.GetInfoByUID, func(int64) (*varys.UserInfo, error) {
		return &varys.UserInfo{Name: "RandySun"}, nil
	})

	res := MyFunc(123)

	if !strings.Contains(res, "RandySun") {
		t.Fatal()
	}
	t.Logf("name: %s", res)
}
