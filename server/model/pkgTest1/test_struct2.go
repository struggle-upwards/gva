// 自动生成模板Test_struct2
package pkgTest1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 测试结构体2 结构体  Test_struct2
type Test_struct2 struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名;"`  //姓名 
      Score  *int `json:"score" form:"score" gorm:"column:score;comment:成绩;"`  //成绩 
}


// TableName 测试结构体2 Test_struct2自定义表名 test_struct2
func (Test_struct2) TableName() string {
  return "test_struct2"
}

