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

type Test_struct2Api struct {
}

var ts2Service = service.ServiceGroupApp.PkgTest1ServiceGroup.Test_struct2Service


// CreateTest_struct2 创建测试结构体2
// @Tags Test_struct2
// @Summary 创建测试结构体2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest1.Test_struct2 true "创建测试结构体2"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ts2/createTest_struct2 [post]
func (ts2Api *Test_struct2Api) CreateTest_struct2(c *gin.Context) {
	var ts2 pkgTest1.Test_struct2
	err := c.ShouldBindJSON(&ts2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := ts2Service.CreateTest_struct2(&ts2); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTest_struct2 删除测试结构体2
// @Tags Test_struct2
// @Summary 删除测试结构体2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest1.Test_struct2 true "删除测试结构体2"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ts2/deleteTest_struct2 [delete]
func (ts2Api *Test_struct2Api) DeleteTest_struct2(c *gin.Context) {
	ID := c.Query("ID")
	if err := ts2Service.DeleteTest_struct2(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTest_struct2ByIds 批量删除测试结构体2
// @Tags Test_struct2
// @Summary 批量删除测试结构体2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ts2/deleteTest_struct2ByIds [delete]
func (ts2Api *Test_struct2Api) DeleteTest_struct2ByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := ts2Service.DeleteTest_struct2ByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTest_struct2 更新测试结构体2
// @Tags Test_struct2
// @Summary 更新测试结构体2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest1.Test_struct2 true "更新测试结构体2"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ts2/updateTest_struct2 [put]
func (ts2Api *Test_struct2Api) UpdateTest_struct2(c *gin.Context) {
	var ts2 pkgTest1.Test_struct2
	err := c.ShouldBindJSON(&ts2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := ts2Service.UpdateTest_struct2(ts2); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTest_struct2 用id查询测试结构体2
// @Tags Test_struct2
// @Summary 用id查询测试结构体2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest1.Test_struct2 true "用id查询测试结构体2"
// @Success 200 {object} response.Response{data=object{rets2=pkgTest1.Test_struct2},msg=string} "查询成功"
// @Router /ts2/findTest_struct2 [get]
func (ts2Api *Test_struct2Api) FindTest_struct2(c *gin.Context) {
	ID := c.Query("ID")
	if rets2, err := ts2Service.GetTest_struct2(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rets2": rets2}, c)
	}
}

// GetTest_struct2List 分页获取测试结构体2列表
// @Tags Test_struct2
// @Summary 分页获取测试结构体2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest1Req.Test_struct2Search true "分页获取测试结构体2列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ts2/getTest_struct2List [get]
func (ts2Api *Test_struct2Api) GetTest_struct2List(c *gin.Context) {
	var pageInfo pkgTest1Req.Test_struct2Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ts2Service.GetTest_struct2InfoList(pageInfo); err != nil {
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

// GetTest_struct2Public 不需要鉴权的测试结构体2接口
// @Tags Test_struct2
// @Summary 不需要鉴权的测试结构体2接口
// @accept application/json
// @Produce application/json
// @Param data query pkgTest1Req.Test_struct2Search true "分页获取测试结构体2列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ts2/getTest_struct2Public [get]
func (ts2Api *Test_struct2Api) GetTest_struct2Public(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的测试结构体2接口信息",
    }, "获取成功", c)
}
