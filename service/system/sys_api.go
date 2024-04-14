package system

import (
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/lxhcaicai/gin-vue-admin/server/model/common/request"
	"github.com/lxhcaicai/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ApiService struct {
}

func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GVA_DB.Create(&api).Error
}

// DeleteApi
//
//	@Description: 删除基础api
func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	var entity system.SysApi
	err = global.GVA_DB.Where("id = ?", api.ID).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method)
	if err != nil {
		return err
	}
	return nil
}

// GetApiById
//
//	@Description: 根据id获取api
func (apiService *ApiService) GetApiById(id int) (api system.SysApi, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: UpdateApi
// @description: 根据id更新api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.GVA_DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		var duplicateApi system.SysApi
		if err := global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&duplicateApi).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
		}
	}
	if err != nil {
		err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.GVA_DB.Save(&api).Error
		}
	}
	return err
}

// DeleteApisByIds
//
//	@Description: 删除选中API
func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	var apis []system.SysApi
	err = global.GVA_DB.Find(apis, "id in ?", ids.Ids).Delete(&apis).Error
	if err != nil {
		return err
	} else {
		for _, sysApi := range apis {
			CasbinServiceApp.ClearCasbin(1, sysApi.Path, sysApi.Method)
		}
		if err != nil {
			return err
		}
	}
	return err
}
