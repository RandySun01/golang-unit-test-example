package http_test_demo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-04-30-21:06
*/

type Param struct {
	Name string `json:"name"`
}

//
//  helloHandler
//  @Description: hello请求处理函数
//  @param c
//
func helloHandler(c *gin.Context) {
	var p Param
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "we need a name",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("hello %s", p.Name),
	})
}

//
// SetupRoute
//  @Description: 路由
//  @return *gin.Engine
//
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/hello", helloHandler)
	return router
}
