import service from '@/utils/request'

// @Tags TestStruct1
// @Summary 创建测试结构体1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStruct1 true "创建测试结构体1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /testStruct1/createTestStruct1 [post]
export const createTestStruct1 = (data) => {
  return service({
    url: '/testStruct1/createTestStruct1',
    method: 'post',
    data
  })
}

// @Tags TestStruct1
// @Summary 删除测试结构体1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStruct1 true "删除测试结构体1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testStruct1/deleteTestStruct1 [delete]
export const deleteTestStruct1 = (params) => {
  return service({
    url: '/testStruct1/deleteTestStruct1',
    method: 'delete',
    params
  })
}

// @Tags TestStruct1
// @Summary 批量删除测试结构体1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试结构体1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testStruct1/deleteTestStruct1 [delete]
export const deleteTestStruct1ByIds = (params) => {
  return service({
    url: '/testStruct1/deleteTestStruct1ByIds',
    method: 'delete',
    params
  })
}

// @Tags TestStruct1
// @Summary 更新测试结构体1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStruct1 true "更新测试结构体1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testStruct1/updateTestStruct1 [put]
export const updateTestStruct1 = (data) => {
  return service({
    url: '/testStruct1/updateTestStruct1',
    method: 'put',
    data
  })
}

// @Tags TestStruct1
// @Summary 用id查询测试结构体1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestStruct1 true "用id查询测试结构体1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testStruct1/findTestStruct1 [get]
export const findTestStruct1 = (params) => {
  return service({
    url: '/testStruct1/findTestStruct1',
    method: 'get',
    params
  })
}

// @Tags TestStruct1
// @Summary 分页获取测试结构体1列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试结构体1列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testStruct1/getTestStruct1List [get]
export const getTestStruct1List = (params) => {
  return service({
    url: '/testStruct1/getTestStruct1List',
    method: 'get',
    params
  })
}
