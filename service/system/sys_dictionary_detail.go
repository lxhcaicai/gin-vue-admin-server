package system

import (
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/lxhcaicai/gin-vue-admin/server/model/system"
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
