/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package data_source_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func TestWirelessNetworkDataSource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	wirelessNetwork := util.GenerateWirelessNetworksGetResponse("id", "").Items[0]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"id",
						util.GenerateWirelessNetworksGetResponse("id", ""),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_wireless_network" "my_wireless_network" {
						filter = {
							id = "id"
						}
					}
				`,
				Check: shared.CheckStateAgainstWirelessNetwork(
					t,
					"data.hpeuxi_wireless_network.my_wireless_network",
					wirelessNetwork,
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestWirelessNetworkDataSourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	wirelessNetwork := util.GenerateWirelessNetworksGetResponse("id", "").Items[0]
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Get(shared.WirelessNetworkPath).
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetWirelessNetwork(
						"id",
						util.GenerateWirelessNetworksGetResponse("id", ""),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_wireless_network" "my_wireless_network" {
						filter = {
							id = "id"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					shared.CheckStateAgainstWirelessNetwork(
						t,
						"data.hpeuxi_wireless_network.my_wireless_network",
						wirelessNetwork,
					),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestWirelessNetworkAssignmentDataSourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					gock.New(util.MockUXIURL).
						Get(shared.WirelessNetworkPath).
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_wireless_network" "my_wireless_network" {
						filter = {
							id = "id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork("id", util.EmptyGetListResponse, 1)
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_wireless_network" "my_wireless_network" {
						filter = {
							id = "id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Could not find specified data source`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
