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
)

func TestGroupDataSource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test Read
			{
				PreConfig: func() {
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 3)
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_group" "my_group" {
						filter = {
							id = "id"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.hpeuxi_group.my_group", "id", "id"),
				),
			},
			{
				PreConfig: func() {
					util.MockGetGroup(util.MockRootGroupID, util.GenerateRootGroupGetResponse(), 1)
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_group" "my_group" {
						filter = {
							id = "my_root_group_id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`The root group cannot be used as a data source`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestGroupDataSourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test Read
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Get("/networking-uxi/v1alpha1/groups").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 3)
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_group" "my_group" {
						filter = {
							id = "id"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.hpeuxi_group.my_group", "id", "id"),
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

func TestGroupDataSourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// HTTP error
			{
				PreConfig: func() {
					gock.New(util.MockUXIURL).
						Get("/networking-uxi/v1alpha1/groups").
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_group" "my_group" {
						filter = {
							id = "id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Not found error
			{
				PreConfig: func() {
					util.MockGetGroup("id", util.EmptyGetListResponse, 1)
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_group" "my_group" {
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
