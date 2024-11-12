package data_source_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestWirelessNetworkDataSource(t *testing.T) {
	wireless_network := util.GetWirelessNetwork(config.WirelessNetworkUid)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					data "uxi_wireless_network" "my_wireless_network" {
						filter = {
							wireless_network_id = "` + config.WirelessNetworkUid + `"
						}
					}
				`,
				Check: util.CheckStateAgainstWirelessNetwork(t, wireless_network),
			},
		},
	})
}