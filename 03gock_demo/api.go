package gock_demo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
@author RandySun
@create 2022-04-30-22:06
*/

//
// ReqParam
//  @Description: Api请求参数
//
type ReqParam struct {
	X int `json:"x"`
}

//
// Result
//  @Description: API返回结果
//
type Result struct {
	Value int `json:"value"`
}

//
// GetResultByApi
//  @Description: 调用外部api
//  @param x
//  @param y
//  @return int
//
func GetResultByAPI(x, y int) int {
	p := &ReqParam{
		X: x,
	}
	b, _ := json.Marshal(p)
	// 调用其它服务的API
	response, err := http.Post(
		"http://your-api.com/post",
		"application/json",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return -1
	}
	body, _ := ioutil.ReadAll(response.Body)
	var rest Result
	if err != json.Unmarshal(body, &rest) {
		return -1
	}
	// 这里是对API返回的数据做一些逻辑处理
	return rest.Value + y
}
