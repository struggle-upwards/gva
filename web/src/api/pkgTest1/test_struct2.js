import service from '@/utils/request'

// @Tags Test_struct2
// @Summary 创建测试结构体2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test_struct2 true "创建测试结构体2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ts2/createTest_struct2 [post]
export const createTest_struct2 = (data) => {
  return service({
    url: '/ts2/createTest_struct2',
    method: 'post',
    data
  })
}

// @Tags Test_struct2
// @Summary 删除测试结构体2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test_struct2 true "删除测试结构体2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ts2/deleteTest_struct2 [delete]
export const deleteTest_struct2 = (params) => {
  return service({
    url: '/ts2/deleteTest_struct2',
    method: 'delete',
    params
  })
}

// @Tags Test_struct2
// @Summary 批量删除测试结构体2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试结构体2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ts2/deleteTest_struct2 [delete]
export const deleteTest_struct2ByIds = (params) => {
  return service({
    url: '/ts2/deleteTest_struct2ByIds',
    method: 'delete',
    params
  })
}

// @Tags Test_struct2
// @Summary 更新测试结构体2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test_struct2 true "更新测试结构体2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ts2/updateTest_struct2 [put]
export const updateTest_struct2 = (data) => {
  return service({
    url: '/ts2/updateTest_struct2',
    method: 'put',
    data
  })
}

// @Tags Test_struct2
// @Summary 用id查询测试结构体2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Test_struct2 true "用id查询测试结构体2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ts2/findTest_struct2 [get]
export const findTest_struct2 = (params) => {
  return service({
    url: '/ts2/findTest_struct2',
    method: 'get',
    params
  })
}

// @Tags Test_struct2
// @Summary 分页获取测试结构体2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试结构体2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ts2/getTest_struct2List [get]
export const getTest_struct2List = (params) => {
  return service({
    url: '/ts2/getTest_struct2List',
    method: 'get',
    params
  })
}
