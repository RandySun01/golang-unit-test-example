package goconvey_demo

import "strings"

/*
@author RandySun
@create 2022-05-01-19:19
*/

//
// Split
//  @Description: 把字符串s按照给定的分隔符sep进行分割返回字符串切片
//  @param s
//  @param sep
//  @return result
//
func Split(s, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}

	result = append(result, s)
	return
}
