package base_demo

import "strings"

/*
@author RandySun
@create 2022-04-30-14:48
*/

//
// Split
//  @Description: 把字符串s按照给定的分隔符sep进行分割返回字符串切片
//  @param s
//  @param sep
//  @return result
//
//func Split(s, sep string) (result []string) {
//	//返回子串str在字符串s中第一次出现的位置。
//	//如果找不到则返回-1；如果str为空，则返回0
//	i := strings.Index(s, sep)
//	for i > -1 {
//		result = append(result, s[:i])
//		s = s[i+1:]
//		i = strings.Index(s, sep)
//	}
//	result = append(result, s)
//	return
//}

// Split
//  @Description: 把字符串s按照给定的分隔符sep进行分割返回字符串切片
//  @param s
//  @param sep
//  @return result
//
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	result = make([]string, 0, strings.Count(s, sep)+1)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)

	}
	// 去除首尾作为分隔符
	//if s == "" {
	//	return
	//}
	result = append(result, s)
	return

}
