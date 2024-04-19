package request

import (
	"github.com/lxhcaicai/gin-vue-admin/server/model/common/request"
	"github.com/lxhcaicai/gin-vue-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
