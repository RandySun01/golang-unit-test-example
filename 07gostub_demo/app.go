package gostub_demo

import "io/ioutil"

/*
@author RandySun
@create 2022-05-01-16:39
*/

var (
	configFile = "config.json"
	maxNum     = 10
)

func GetConfig() ([]byte, error) {
	return ioutil.ReadFile(configFile)
}

func ShowNumBer() int {
	return maxNum
}
