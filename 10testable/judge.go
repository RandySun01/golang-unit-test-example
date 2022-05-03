package main

import "time"

/*
@author RandySun
@create 2022-05-01-20:32
*/

//
//  judgeRate
//  @Description: 报警速率决策函数
//  @return int
//
func judgeRate() int {
	now := time.Now()
	switch hour := now.Hour(); {
	case hour >= 8 && hour < 20:
		return 10
	case hour >= 20 && hour <= 20:
		return 1
	}
	return -1
}

//
//  judgeRateByTime
//  @Description: 报警速率决策函数,时间作为参数
//  @param now
//  @return int
//
func judgeRateByTime(now time.Time) int {
	switch hour := now.Hour(); {
	case hour >= 8 && hour < 20:
		return 10
	case hour >= 20 && hour <= 23:
		return 1
	}
	return -1

}
