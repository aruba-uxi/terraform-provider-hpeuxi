/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func TestServiceTestResource(t *testing.T) {
	serviceTest := util.GetServiceTest(config.ServiceTestID)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a service_test is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_service_test" "my_service_test" {
						name = "` + serviceTest.Name + `"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)creating a service_test is not supported; service_tests can only be\s*imported`,
				),
			},
			// Importing a service_test
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_service_test" "my_service_test" {
						name = "` + serviceTest.Name + `"
					}

					import {
						to = hpeuxi_service_test.my_service_test
						id = "` + config.ServiceTestID + `"
					}`,

				Check: shared.CheckStateAgainstServiceTest(
					t,
					"hpeuxi_service_test.my_service_test",
					serviceTest,
				),
			},
			// ImportState testing
			{
				ResourceName:      "hpeuxi_service_test.my_service_test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Updating a service_test is not allowed
			{
				Config: provider.ProviderConfig + `
				resource "hpeuxi_service_test" "my_service_test" {
					name = "` + serviceTest.Name + `-updated-name"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)updating a service_test is not supported; service_tests can only be updated\s*through the dashboard`,
				),
			},
			// Deleting a service_test is not allowed
			{
				Config: provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(
					`(?s)deleting a service_test is not supported; service_tests can only removed from\s*state`,
				),
			},
			// Remove service_test from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_service_test.my_service_test

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}
