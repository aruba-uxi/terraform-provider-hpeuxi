package resource_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/stretchr/testify/assert"
)

func TestSensorResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a sensor is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "uxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "note"
						pcap_mode = "light"
					}`,

				ExpectError: regexp.MustCompile(
					`creating a sensor is not supported; sensors can only be imported`,
				),
			},
			// Importing a sensor
			{
				PreConfig: func() {
					util.MockGetSensor("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("id", "")}),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "name", "name"),
					resource.TestCheckResourceAttr(
						"uxi_sensor.my_sensor",
						"address_note",
						"address_note",
					),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "notes", "notes"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "pcap_mode", "light"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "id", "id"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetSensor("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("id", "")}),
						1,
					)
				},
				ResourceName:      "uxi_sensor.my_sensor",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					// existing sensor
					util.MockGetSensor("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("id", "")}),
						1,
					)
					util.MockUpdateSensor(
						"id",
						util.GenerateSensorRequestUpdateModel("_2"),
						util.GenerateSensorResponseModel("id", "_2"),
						1,
					)
					// updated sensor
					util.MockGetSensor("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("id", "_2")}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_sensor" "my_sensor" {
					name = "name_2"
					address_note = "address_note_2"
					notes = "notes_2"
					pcap_mode = "light"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "name", "name_2"),
					resource.TestCheckResourceAttr(
						"uxi_sensor.my_sensor",
						"address_note",
						"address_note_2",
					),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "notes", "notes_2"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "pcap_mode", "light"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "id", "id"),
				),
			},
			// Deleting a sensor is not allowed
			{
				PreConfig: func() {
					util.MockGetSensor("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("id", "_2")}),
						1,
					)
				},
				Config: provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(
					`deleting a sensor is not supported; sensors can only removed from state`,
				),
			},
			// Remove sensor from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestSensorResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Importing a sensor
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/sensors").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetSensor("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("id", "")}),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "id", "id"),
					func(s *terraform.State) error {
						assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Update and Read testing
			{
				PreConfig: func() {
					// existing sensor
					util.MockGetSensor("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("id", "")}),
						1,
					)
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Patch("/networking-uxi/v1alpha1/sensors/id").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockUpdateSensor(
						"id",
						util.GenerateSensorRequestUpdateModel("_2"),
						util.GenerateSensorResponseModel("id", "_2"),
						1,
					)
					// updated sensor
					util.MockGetSensor("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("id", "_2")}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_sensor" "my_sensor" {
					name = "name_2"
					address_note = "address_note_2"
					notes = "notes_2"
					pcap_mode = "light"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "name", "name_2"),
					func(s *terraform.State) error {
						assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Remove sensor from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()

}
func TestSensorResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Read 5xx error
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/sensors").
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "id"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Read not found
			{
				PreConfig: func() {
					util.MockGetSensor(
						"id",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "id"
					}`,

				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Actually import a sensor for subsequent testing
			{
				PreConfig: func() {
					util.MockGetSensor("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("id", "")}),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "name", "name"),
					resource.TestCheckResourceAttr(
						"uxi_sensor.my_sensor",
						"address_note",
						"address_note",
					),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "notes", "notes"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "pcap_mode", "light"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "id", "id"),
				),
			},
			// Update 4xx
			{
				PreConfig: func() {
					// existing sensor
					util.MockGetSensor("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("id", "")}),
						1,
					)
					// patch sensor - with error
					gock.New("https://test.api.capenetworks.com").
						Patch("/networking-uxi/v1alpha1/sensors/id").
						Reply(http.StatusUnprocessableEntity).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusUnprocessableEntity,
							"errorCode":      "HPE_GL_UXI_INVALID_PCAP_MODE_ERROR",
							"message":        "Unable to update sensor - pcap_mode must be one the following ['light', 'full', 'off'].",
							"debugId":        "12312-123123-123123-1231212",
							"type":           "hpe.greenlake.uxi.invalid_pcap_mode",
						})
				},
				Config: provider.ProviderConfig + `
				resource "uxi_sensor" "my_sensor" {
					name = "name_2"
					address_note = "address_note_2"
					notes = "notes_2"
					pcap_mode = "light"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)Unable to update sensor - pcap_mode must be one the following \['light',\s*'full', 'off'\].\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Remove sensor from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
