package http_test_demo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_helloHandler(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	// 创建测试用例
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"base case1", `{"name": "RandySun"}`, "hello RandySun"},
		{"bad case2", "", "we need a name"},
		{"bad case3", "", "name"},
	}

	// 注册路由
	r := SetupRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock一个HTTP请求
			req := httptest.NewRequest(
				"POST",                      // 请求方法
				"/hello",                    // 请求URL
				strings.NewReader(tt.param), // 请求参数
			)

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)

			// 校验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)

			// 解析并检验响应内容是否复合预期
			var resp map[string]string
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			t.Logf("response msg: %#v", resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp["msg"])

		})
	}
}

//func TestSetupRoute(t *testing.T) {
//	tests := []struct {
//		name string
//		want *gin.Engine
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := SetupRouter(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("SetupRoute() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
