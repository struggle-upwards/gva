package pkgTest1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest1"
    pkgTest1Req "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest1/request"
)

type Test_struct2Service struct {
}

// CreateTest_struct2 创建测试结构体2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ts2Service *Test_struct2Service) CreateTest_struct2(ts2 *pkgTest1.Test_struct2) (err error) {
	err = global.GVA_DB.Create(ts2).Error
	return err
}

// DeleteTest_struct2 删除测试结构体2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ts2Service *Test_struct2Service)DeleteTest_struct2(ID string) (err error) {
	err = global.GVA_DB.Delete(&pkgTest1.Test_struct2{},"id = ?",ID).Error
	return err
}

// DeleteTest_struct2ByIds 批量删除测试结构体2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ts2Service *Test_struct2Service)DeleteTest_struct2ByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]pkgTest1.Test_struct2{},"id in ?",IDs).Error
	return err
}

// UpdateTest_struct2 更新测试结构体2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ts2Service *Test_struct2Service)UpdateTest_struct2(ts2 pkgTest1.Test_struct2) (err error) {
	err = global.GVA_DB.Model(&pkgTest1.Test_struct2{}).Where("id = ?",ts2.ID).Updates(&ts2).Error
	return err
}

// GetTest_struct2 根据ID获取测试结构体2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ts2Service *Test_struct2Service)GetTest_struct2(ID string) (ts2 pkgTest1.Test_struct2, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ts2).Error
	return
}

// GetTest_struct2InfoList 分页获取测试结构体2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ts2Service *Test_struct2Service)GetTest_struct2InfoList(info pkgTest1Req.Test_struct2Search) (list []pkgTest1.Test_struct2, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgTest1.Test_struct2{})
    var ts2s []pkgTest1.Test_struct2
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
	
	err = db.Find(&ts2s).Error
	return  ts2s, total, err
}