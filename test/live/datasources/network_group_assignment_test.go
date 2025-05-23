/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package data_source_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
)

func TestNetworkGroupAssignmentDataSource(t *testing.T) {
	const groupName = "tf_provider_acceptance_test_network_group_assignment_datasource"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					// create the resource to be used as a datasource
					resource "hpeuxi_group" "my_group" {
						name = "` + groupName + `"
					}

					data "hpeuxi_wired_network" "my_network" {
						filter = {
							id = "` + config.WiredNetworkID + `"
						}
					}

					resource "hpeuxi_network_group_assignment" "my_network_group_assignment" {
						network_id = data.hpeuxi_wired_network.my_network.id
						group_id   = hpeuxi_group.my_group.id
					}

					// the actual datasource
					data "hpeuxi_network_group_assignment" "my_network_group_assignment" {
						filter = {
							id = hpeuxi_network_group_assignment.my_network_group_assignment.id
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						resourceName := "hpeuxi_network_group_assignment.my_network_group_assignment"
						rs := s.RootModule().Resources[resourceName]

						return util.CheckStateAgainstNetworkGroupAssignment(
							t,
							"data.hpeuxi_network_group_assignment.my_network_group_assignment",
							util.GetNetworkGroupAssignment(rs.Primary.ID),
						)(s)
					},
				),
			},
		},
	})
}
