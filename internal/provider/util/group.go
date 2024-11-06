package util

import (
	"github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func IsRoot(group config_api_client.GroupsGetItem) bool {
	_, set := group.Parent.Get().GetIdOk()
	return !set
}