// 自动生成模板TestStruct1
package pkgTest1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 测试结构体1 结构体  TestStruct1
type TestStruct1 struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"default:"";column:name;comment:姓名;size:191;"`  //姓名 
      Age  *int `json:"age" form:"age" gorm:"default:0;column:age;comment:年龄;size:3;"`  //年龄 
}


// TableName 测试结构体1 TestStruct1自定义表名 testStruct1
func (TestStruct1) TableName() string {
  return "testStruct1"
}

