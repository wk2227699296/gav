package wkTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wkTest"
	wkTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/wkTest/request"
)

type WkClassService struct {
}

// CreateWkClass 创建WkClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (wClassService *WkClassService) CreateWkClass(wClass *wkTest.WkClass) (err error) {
	err = global.GVA_DB.Create(wClass).Error
	return err
}

// DeleteWkClass 删除WkClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (wClassService *WkClassService) DeleteWkClass(wClass wkTest.WkClass) (err error) {
	err = global.GVA_DB.Delete(&wClass).Error
	return err
}

// DeleteWkClassByIds 批量删除WkClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (wClassService *WkClassService) DeleteWkClassByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wkTest.WkClass{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWkClass 更新WkClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (wClassService *WkClassService) UpdateWkClass(wClass wkTest.WkClass) (err error) {
	err = global.GVA_DB.Save(&wClass).Error
	return err
}

// GetWkClass 根据id获取WkClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (wClassService *WkClassService) GetWkClass(id uint) (wClass wkTest.WkClass, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&wClass).Error
	return
}

// GetWkClassInfoList 分页获取WkClass记录
// Author [piexlmax](https://github.com/piexlmax)
func (wClassService *WkClassService) GetWkClassInfoList(info wkTestReq.WkClassSearch) (list []wkTest.WkClass, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wkTest.WkClass{})
	var wClasss []wkTest.WkClass
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wClasss).Error
	return wClasss, total, err
}
