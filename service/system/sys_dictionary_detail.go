package system

import (
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/lxhcaicai/gin-vue-admin/server/model/system"
	"github.com/lxhcaicai/gin-vue-admin/server/model/system/request"
)

type DictionaryDetailService struct {
}

// CreateSysDictionaryDetail
//
//	@Description: 创建字典详情数据
func (dictionaryDetailService *DictionaryDetailService) CreateSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Create(&sysDictionaryDetail).Error
	return err
}

// DeleteSysDictionaryDetail
//
//	@Description: 删除字典详情数据
func (dictionaryDetailService *DictionaryDetailService) DeleteSysDictionaryDetail(syssysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Delete(&syssysDictionaryDetail).Error
	return err
}

// UpdateSysDictionaryDetail
//
//	@Description: 更新字典详情数据
func (dictionaryDetailService *DictionaryDetailService) UpdateSysDictionaryDetail(sysDictionaryDetail *system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Save(sysDictionaryDetail).Error
	return err
}

// GetSysDictionaryDetail
//
//	@Description: 根据id获取字典详情单条数据
func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetail(id uint) (sysDictionaryDetail system.SysDictionaryDetail, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

// GetSysDictionaryDetailInfoList
//
//	@Description: 分页获取字典详情列表
func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetailInfoList(info request.SysDictionaryDetailSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{})
	var sysDictionaryDetails []system.SysDictionaryDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != "" {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort").Find(&sysDictionaryDetails).Error
	return sysDictionaryDetails, total, err
}
