package mytest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyTestApi struct {
}

// Login
// @Tags     Test
// @Summary  测试api
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /mytest/testapi [get]
func (m *MyTestApi) TestAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
	})
}
