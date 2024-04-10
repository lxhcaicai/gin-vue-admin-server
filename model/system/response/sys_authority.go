package response

import "github.com/lxhcaicai/gin-vue-admin/server/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}
