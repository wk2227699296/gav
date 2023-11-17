// 自动生成模板WkStudent
package wkTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// wkStudent 结构体  WkStudent
type WkStudent struct {
	global.GVA_MODEL
	Name   string `json:"name" form:"name" gorm:"column:name;comment:;"`       //name
	Age    string `json:"age" form:"age" gorm:"column:age;comment:;"`          //age
	Gender *int   `json:"gender" form:"gender" gorm:"column:gender;comment:;"` //gender
}

// TableName wkStudent WkStudent自定义表名 wk_student
func (WkStudent) TableName() string {
	return "wk_student"
}

type WkInfoWkStudent struct {
}

func (WkInfoWkStudent) TableName() string {
	return "wk_info_wk_student"
}
