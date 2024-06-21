package pkgTest1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest1"
    pkgTest1Req "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest1/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TestStruct1Api struct {
}

var testStruct1Service = service.ServiceGroupApp.PkgTest1ServiceGroup.TestStruct1Service


// CreateTestStruct1 创建测试结构体1
// @Tags TestStruct1
// @Summary 创建测试结构体1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest1.TestStruct1 true "创建测试结构体1"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /testStruct1/createTestStruct1 [post]
func (testStruct1Api *TestStruct1Api) CreateTestStruct1(c *gin.Context) {
	var testStruct1 pkgTest1.TestStruct1
	err := c.ShouldBindJSON(&testStruct1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := testStruct1Service.CreateTestStruct1(&testStruct1); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestStruct1 删除测试结构体1
// @Tags TestStruct1
// @Summary 删除测试结构体1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest1.TestStruct1 true "删除测试结构体1"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /testStruct1/deleteTestStruct1 [delete]
func (testStruct1Api *TestStruct1Api) DeleteTestStruct1(c *gin.Context) {
	ID := c.Query("ID")
	if err := testStruct1Service.DeleteTestStruct1(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestStruct1ByIds 批量删除测试结构体1
// @Tags TestStruct1
// @Summary 批量删除测试结构体1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /testStruct1/deleteTestStruct1ByIds [delete]
func (testStruct1Api *TestStruct1Api) DeleteTestStruct1ByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := testStruct1Service.DeleteTestStruct1ByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestStruct1 更新测试结构体1
// @Tags TestStruct1
// @Summary 更新测试结构体1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest1.TestStruct1 true "更新测试结构体1"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /testStruct1/updateTestStruct1 [put]
func (testStruct1Api *TestStruct1Api) UpdateTestStruct1(c *gin.Context) {
	var testStruct1 pkgTest1.TestStruct1
	err := c.ShouldBindJSON(&testStruct1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := testStruct1Service.UpdateTestStruct1(testStruct1); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestStruct1 用id查询测试结构体1
// @Tags TestStruct1
// @Summary 用id查询测试结构体1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest1.TestStruct1 true "用id查询测试结构体1"
// @Success 200 {object} response.Response{data=object{retestStruct1=pkgTest1.TestStruct1},msg=string} "查询成功"
// @Router /testStruct1/findTestStruct1 [get]
func (testStruct1Api *TestStruct1Api) FindTestStruct1(c *gin.Context) {
	ID := c.Query("ID")
	if retestStruct1, err := testStruct1Service.GetTestStruct1(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestStruct1": retestStruct1}, c)
	}
}

// GetTestStruct1List 分页获取测试结构体1列表
// @Tags TestStruct1
// @Summary 分页获取测试结构体1列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest1Req.TestStruct1Search true "分页获取测试结构体1列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /testStruct1/getTestStruct1List [get]
func (testStruct1Api *TestStruct1Api) GetTestStruct1List(c *gin.Context) {
	var pageInfo pkgTest1Req.TestStruct1Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testStruct1Service.GetTestStruct1InfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// GetTestStruct1Public 不需要鉴权的测试结构体1接口
// @Tags TestStruct1
// @Summary 不需要鉴权的测试结构体1接口
// @accept application/json
// @Produce application/json
// @Param data query pkgTest1Req.TestStruct1Search true "分页获取测试结构体1列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /testStruct1/getTestStruct1Public [get]
func (testStruct1Api *TestStruct1Api) GetTestStruct1Public(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的测试结构体1接口信息",
    }, "获取成功", c)
}
