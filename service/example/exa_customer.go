package example

import (
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/lxhcaicai/gin-vue-admin/server/model/common/request"
	"github.com/lxhcaicai/gin-vue-admin/server/model/example"
	"github.com/lxhcaicai/gin-vue-admin/server/model/system"
	systemService "github.com/lxhcaicai/gin-vue-admin/server/service/system"
)

type CustomerService struct {
}

// CreateExaCustomer
//
//	@Description: 创建客户
func (exa *CustomerService) CreateExaCustomer(e example.ExaCustomer) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

// UpdateExaCustomer
//
//	@Description: 更新客户
func (exa *CustomerService) UpdateExaCustomer(e *example.ExaCustomer) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

// DeleteExaCustomer
//
//	@Description: 删除客户
func (exa *CustomerService) DeleteExaCustomer(e example.ExaCustomer) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

// GetExaCustomer
//
//	@Description: 获取客户信息
func (exa *CustomerService) GetExaCustomer(id uint) (customer example.ExaCustomer, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

// GetCustomerInfoList
//
//	@Description: 分页获取客户列表
func (exa *CustomerService) GetCustomerInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.Page
	offset := info.PageSize
	db := global.GVA_DB.Model(&example.ExaCustomer{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []uint
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []example.ExaCustomer
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return CustomerList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", dataId).Find(&CustomerList).Error
	}
	return CustomerList, total, err
}
