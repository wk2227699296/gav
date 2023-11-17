// 自动生成模板WkInfo
package wkTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// WkInfo 结构体  WkInfo
type WkInfo struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:;"`    //name
	Age       string `json:"age" form:"age" gorm:"column:age;comment:;"`       //age
	Email     string `json:"email" form:"email" gorm:"column:email;comment:;"` //email
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName WkInfo WkInfo自定义表名 wk_info
func (WkInfo) TableName() string {
	return "wk_info"
}
