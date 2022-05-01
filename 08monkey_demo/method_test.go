package monkey_demo

import (
	"reflect"
	"strings"
	"testing"

	"bou.ke/monkey"
)

//
//  TestUser_GetInfo
//  @Description: 为方法打桩
//  @param t
//
func TestUser_GetInfo(t *testing.T) {
	var u = &User{
		Name:     "RandySun",
		Birthday: "2004-06-30",
	}
	// 为对象方法打桩
	monkey.PatchInstanceMethod(reflect.TypeOf(u), "CalcAge", func(*User) int {
		return 18
	})

	res := u.GetInfo() // 内部调用u.CalcAge方法会返回18
	if !strings.Contains(res, "朋友") {
		t.Fatal()
	}

	t.Logf("%s", res)
}
