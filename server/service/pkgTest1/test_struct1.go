package pkgTest1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest1"
    pkgTest1Req "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest1/request"
)

type TestStruct1Service struct {
}

// CreateTestStruct1 创建测试结构体1记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct1Service *TestStruct1Service) CreateTestStruct1(testStruct1 *pkgTest1.TestStruct1) (err error) {
	err = global.GVA_DB.Create(testStruct1).Error
	return err
}

// DeleteTestStruct1 删除测试结构体1记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct1Service *TestStruct1Service)DeleteTestStruct1(ID string) (err error) {
	err = global.GVA_DB.Delete(&pkgTest1.TestStruct1{},"id = ?",ID).Error
	return err
}

// DeleteTestStruct1ByIds 批量删除测试结构体1记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct1Service *TestStruct1Service)DeleteTestStruct1ByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]pkgTest1.TestStruct1{},"id in ?",IDs).Error
	return err
}

// UpdateTestStruct1 更新测试结构体1记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct1Service *TestStruct1Service)UpdateTestStruct1(testStruct1 pkgTest1.TestStruct1) (err error) {
	err = global.GVA_DB.Model(&pkgTest1.TestStruct1{}).Where("id = ?",testStruct1.ID).Updates(&testStruct1).Error
	return err
}

// GetTestStruct1 根据ID获取测试结构体1记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct1Service *TestStruct1Service)GetTestStruct1(ID string) (testStruct1 pkgTest1.TestStruct1, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&testStruct1).Error
	return
}

// GetTestStruct1InfoList 分页获取测试结构体1记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct1Service *TestStruct1Service)GetTestStruct1InfoList(info pkgTest1Req.TestStruct1Search) (list []pkgTest1.TestStruct1, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgTest1.TestStruct1{})
    var testStruct1s []pkgTest1.TestStruct1
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Age != nil {
        db = db.Where("age = ?",info.Age)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&testStruct1s).Error
	return  testStruct1s, total, err
}