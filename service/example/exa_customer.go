package example

import (
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/lxhcaicai/gin-vue-admin/server/model/example"
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
