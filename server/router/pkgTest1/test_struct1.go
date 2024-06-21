package pkgTest1

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestStruct1Router struct {
}

// InitTestStruct1Router 初始化 测试结构体1 路由信息
func (s *TestStruct1Router) InitTestStruct1Router(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	testStruct1Router := Router.Group("testStruct1").Use(middleware.OperationRecord())
	testStruct1RouterWithoutRecord := Router.Group("testStruct1")
	testStruct1RouterWithoutAuth := PublicRouter.Group("testStruct1")

	var testStruct1Api = v1.ApiGroupApp.PkgTest1ApiGroup.TestStruct1Api
	{
		testStruct1Router.POST("createTestStruct1", testStruct1Api.CreateTestStruct1)             // 新建测试结构体1
		testStruct1Router.DELETE("deleteTestStruct1", testStruct1Api.DeleteTestStruct1)           // 删除测试结构体1
		testStruct1Router.DELETE("deleteTestStruct1ByIds", testStruct1Api.DeleteTestStruct1ByIds) // 批量删除测试结构体1
		testStruct1Router.PUT("updateTestStruct1", testStruct1Api.UpdateTestStruct1)              // 更新测试结构体1
	}
	{
		testStruct1RouterWithoutRecord.GET("findTestStruct1", testStruct1Api.FindTestStruct1)       // 根据ID获取测试结构体1
		testStruct1RouterWithoutRecord.GET("getTestStruct1List", testStruct1Api.GetTestStruct1List) // 获取测试结构体1列表
	}
	{
		testStruct1RouterWithoutAuth.GET("getTestStruct1Public", testStruct1Api.GetTestStruct1Public) // 获取测试结构体1列表
	}
}
