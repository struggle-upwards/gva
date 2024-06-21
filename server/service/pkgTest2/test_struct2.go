package pkgTest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest2"
    pkgTest2Req "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest2/request"
)

type TestStruct2Service struct {
}

// CreateTestStruct2 创建测试包中文2记录
// Author [piexlmax](https://github.com/piexlmax)
func (TS2Service *TestStruct2Service) CreateTestStruct2(TS2 *pkgTest2.TestStruct2) (err error) {
	err = global.GVA_DB.Create(TS2).Error
	return err
}

// DeleteTestStruct2 删除测试包中文2记录
// Author [piexlmax](https://github.com/piexlmax)
func (TS2Service *TestStruct2Service)DeleteTestStruct2(ID string) (err error) {
	err = global.GVA_DB.Delete(&pkgTest2.TestStruct2{},"id = ?",ID).Error
	return err
}

// DeleteTestStruct2ByIds 批量删除测试包中文2记录
// Author [piexlmax](https://github.com/piexlmax)
func (TS2Service *TestStruct2Service)DeleteTestStruct2ByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]pkgTest2.TestStruct2{},"id in ?",IDs).Error
	return err
}

// UpdateTestStruct2 更新测试包中文2记录
// Author [piexlmax](https://github.com/piexlmax)
func (TS2Service *TestStruct2Service)UpdateTestStruct2(TS2 pkgTest2.TestStruct2) (err error) {
	err = global.GVA_DB.Model(&pkgTest2.TestStruct2{}).Where("id = ?",TS2.ID).Updates(&TS2).Error
	return err
}

// GetTestStruct2 根据ID获取测试包中文2记录
// Author [piexlmax](https://github.com/piexlmax)
func (TS2Service *TestStruct2Service)GetTestStruct2(ID string) (TS2 pkgTest2.TestStruct2, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&TS2).Error
	return
}

// GetTestStruct2InfoList 分页获取测试包中文2记录
// Author [piexlmax](https://github.com/piexlmax)
func (TS2Service *TestStruct2Service)GetTestStruct2InfoList(info pkgTest2Req.TestStruct2Search) (list []pkgTest2.TestStruct2, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgTest2.TestStruct2{})
    var TS2s []pkgTest2.TestStruct2
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&TS2s).Error
	return  TS2s, total, err
}