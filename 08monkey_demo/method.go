package monkey_demo

import (
	"fmt"
	"time"
)

/*
@author RandySun
@create 2022-05-01-17:47
*/

type User struct {
	Name     string
	Birthday string
}

// CalCage 计算用户年龄
func (u *User) CalcAge() int {
	t, err := time.Parse("2006-01-02", u.Birthday)
	if err != nil {
		return -1
	}
	return int(time.Now().Sub(t).Hours()/24.0) / 365

}

func (u *User) GetInfo() string {
	age := u.CalcAge()
	if age <= 0 {
		return fmt.Sprintf("%s很神秘, 我们还不了解ta。", u.Name)
	}
	return fmt.Sprintf("%s今年%d岁,ta是我们的朋友。", u.Name, age)

}
