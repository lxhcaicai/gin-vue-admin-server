package system

import (
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/lxhcaicai/gin-vue-admin/server/model/common/request"
	"github.com/lxhcaicai/gin-vue-admin/server/model/system"
	systemReq "github.com/lxhcaicai/gin-vue-admin/server/model/system/request"
	"github.com/lxhcaicai/gin-vue-admin/server/model/system/response"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

var ErrRoleExistence = errors.New("存在相同角色id")

type AuthorityService struct {
}

func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysAuthority{})
	if err = db.Where("parent_id = ?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var authority []system.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = ?", "0").Find(&authority).Error
	for k := range authority {
		err = authorityService.findChildrenAuthority(&authority[k])
	}
	return authority, total, err
}

// findChildrenAuthority
//
//	@Description: 查询子角色
func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {

	if err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}

	e := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&auth).Error; err != nil {
			return err
		}

		auth.SysBaseMenus = systemReq.DefaultMenu()
		if err = tx.Model(&auth).Association("SysBaseMenus").Replace(&auth.SysBaseMenus); err != nil {
			return err
		}
		casbinInfos := systemReq.DefaultCasbin()
		authorityId := strconv.Itoa(int(auth.AuthorityId))
		rules := [][]string{}
		for _, v := range casbinInfos {
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
		return CasbinServiceApp.AddPolicies(tx, rules)
	})

	return auth, e
}

// DeleteAuthority
//
//	@Description: 删除角色
func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority) error {
	if errors.Is(global.GVA_DB.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("parent_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var err error
		if err = tx.Preload("SysBaseMenus").Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(auth).Unscoped().Delete(auth).Error; err != nil {
			return err
		}

		if len(auth.SysBaseMenus) > 0 {
			if err = tx.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus); err != nil {
				return err
			}
		}

		if len(auth.DataAuthorityId) > 0 {
			if err = tx.Model(auth).Association("DataAuthorityId").Delete(auth.DataAuthorityId); err != nil {
				return err
			}
		}

		if err = tx.Delete(&system.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error; err != nil {
			return err
		}
		if err = tx.Where("authority_id = ?", auth.AuthorityId).Delete(&[]system.SysAuthorityBtn{}).Error; err != nil {
			return err
		}

		authorityId := strconv.Itoa(int(auth.AuthorityId))

		if err = CasbinServiceApp.RemoveFilteredPolicy(tx, authorityId); err != nil {
			return err
		}

		return nil
	})
}

// UpdateAuthority
//
//	@Description: 更改一个角色
func (authorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Updates(&auth).Error
	return auth, err
}

// CopyAuthority
//
//	@Description: 复制一个角色
func (authorityService *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (authority system.SysAuthority, err error) {
	var authorityBox system.SysAuthority
	if errors.Is(global.GVA_DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return authority, ErrRoleExistence
	}
	copyInfo.Authority.Children = []system.SysAuthority{}
	menus, err := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	if err != nil {
		return
	}

	var baseMenu []system.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.GVA_DB.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}

	var btns []system.SysAuthorityBtn

	err = global.GVA_DB.Find(&btns, "authority_id = ?", copyInfo.OldAuthorityId).Error
	if err != nil {
		return
	}

	if len(btns) > 0 {
		for i := range btns {
			btns[i].AuthorityId = copyInfo.Authority.AuthorityId
		}
		err = global.GVA_DB.Create(&btns).Error

		if err != nil {
			return

		}
	}

	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = CasbinServiceApp.UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return copyInfo.Authority, err
}
