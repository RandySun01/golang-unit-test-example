package monkey_demo

import (
	"fmt"
	"golang-unit-test-example/varys"
)

/*
@author RandySun
@create 2022-05-01-17:31
*/

// MyFunc
//  @Description: 获取用户名
//  @param uid
//  @return string
//
func MyFunc(uid int64) string {
	u, err := varys.GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	// 这里处理逻辑代码

	return fmt.Sprintf("hello %s\n", u.Name)

}
