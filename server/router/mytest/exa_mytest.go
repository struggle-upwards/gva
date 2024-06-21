package mytest

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

var exa_MyTestApi = v1.ApiGroupApp.MyTestApiGroup.MyTestApi

func (t *TestRouter) InitTestGroup(Router *gin.RouterGroup) {

	mytestRouter := Router.Group("mytest").Use(middleware.OperationRecord())
	{
		mytestRouter.GET("testapi", exa_MyTestApi.TestAPI)
	}
}
