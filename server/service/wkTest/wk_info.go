package wkTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wkTest"
	wkTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/wkTest/request"
	"gorm.io/gorm"
)

type WkInfoService struct {
}

// CreateWkInfo 创建WkInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkinfoService *WkInfoService) CreateWkInfo(wkinfo *wkTest.WkInfo) (err error) {
	err = global.GVA_DB.Create(wkinfo).Error
	return err
}

// DeleteWkInfo 删除WkInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkinfoService *WkInfoService) DeleteWkInfo(wkinfo wkTest.WkInfo) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wkTest.WkInfo{}).Where("id = ?", wkinfo.ID).Update("deleted_by", wkinfo.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&wkinfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteWkInfoByIds 批量删除WkInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkinfoService *WkInfoService) DeleteWkInfoByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wkTest.WkInfo{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&wkTest.WkInfo{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateWkInfo 更新WkInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkinfoService *WkInfoService) UpdateWkInfo(wkinfo wkTest.WkInfo) (err error) {
	err = global.GVA_DB.Save(&wkinfo).Error
	return err
}

// GetWkInfo 根据id获取WkInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkinfoService *WkInfoService) GetWkInfo(id uint) (wkinfo wkTest.WkInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&wkinfo).Error
	return
}

// GetWkInfoInfoList 分页获取WkInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkinfoService *WkInfoService) GetWkInfoInfoList(info wkTestReq.WkInfoSearch) (list []wkTest.WkInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wkTest.WkInfo{})
	var wkinfos []wkTest.WkInfo
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

	err = db.Preload("Student").Find(&wkinfos).Error
	return wkinfos, total, err
}
