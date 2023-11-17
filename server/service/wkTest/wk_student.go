package wkTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wkTest"
	wkTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/wkTest/request"
)

type WkStudentService struct {
}

// CreateWkStudent 创建wkStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkStudentService *WkStudentService) CreateWkStudent(wkStudent *wkTest.WkStudent) (err error) {
	err = global.GVA_DB.Create(wkStudent).Error
	return err
}

// DeleteWkStudent 删除wkStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkStudentService *WkStudentService) DeleteWkStudent(wkStudent wkTest.WkStudent) (err error) {
	err = global.GVA_DB.Delete(&wkStudent).Error
	return err
}

// DeleteWkStudentByIds 批量删除wkStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkStudentService *WkStudentService) DeleteWkStudentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]wkTest.WkStudent{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWkStudent 更新wkStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkStudentService *WkStudentService) UpdateWkStudent(wkStudent wkTest.WkStudent) (err error) {
	err = global.GVA_DB.Save(&wkStudent).Error
	return err
}

// GetWkStudent 根据id获取wkStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkStudentService *WkStudentService) GetWkStudent(id uint) (wkStudent wkTest.WkStudent, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&wkStudent).Error
	return
}

// GetWkStudentInfoList 分页获取wkStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func (wkStudentService *WkStudentService) GetWkStudentInfoList(info wkTestReq.WkStudentSearch) (list []wkTest.WkStudent, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wkTest.WkStudent{})
	var wkStudents []wkTest.WkStudent
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

	err = db.Find(&wkStudents).Error
	return wkStudents, total, err
}
