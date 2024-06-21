// 自动生成模板TestStruct2
package pkgTest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 测试包中文2 结构体  TestStruct2
type TestStruct2 struct {
 global.GVA_MODEL 
      Score  *int `json:"score" form:"score" gorm:"column:score;comment:;"`  //成绩 
}


// TableName 测试包中文2 TestStruct2自定义表名 test_struct2
func (TestStruct2) TableName() string {
  return "test_struct2"
}

