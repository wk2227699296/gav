// 自动生成模板WkClass
package wkTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// WkClass 结构体  WkClass
type WkClass struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"` //name
}

// TableName WkClass WkClass自定义表名 wk_class
func (WkClass) TableName() string {
	return "wk_class"
}
