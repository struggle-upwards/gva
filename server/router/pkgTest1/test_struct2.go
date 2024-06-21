package pkgTest1

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Test_struct2Router struct {
}

// InitTest_struct2Router 初始化 测试结构体2 路由信息
func (s *Test_struct2Router) InitTest_struct2Router(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	ts2Router := Router.Group("ts2").Use(middleware.OperationRecord())
	ts2RouterWithoutRecord := Router.Group("ts2")
	ts2RouterWithoutAuth := PublicRouter.Group("ts2")

	var ts2Api = v1.ApiGroupApp.PkgTest1ApiGroup.Test_struct2Api
	{
		ts2Router.POST("createTest_struct2", ts2Api.CreateTest_struct2)             // 新建测试结构体2
		ts2Router.DELETE("deleteTest_struct2", ts2Api.DeleteTest_struct2)           // 删除测试结构体2
		ts2Router.DELETE("deleteTest_struct2ByIds", ts2Api.DeleteTest_struct2ByIds) // 批量删除测试结构体2
		ts2Router.PUT("updateTest_struct2", ts2Api.UpdateTest_struct2)              // 更新测试结构体2
	}
	{
		ts2RouterWithoutRecord.GET("findTest_struct2", ts2Api.FindTest_struct2)       // 根据ID获取测试结构体2
		ts2RouterWithoutRecord.GET("getTest_struct2List", ts2Api.GetTest_struct2List) // 获取测试结构体2列表
	}
	{
		ts2RouterWithoutAuth.GET("getTest_struct2Public", ts2Api.GetTest_struct2Public) // 获取测试结构体2列表
	}
}
