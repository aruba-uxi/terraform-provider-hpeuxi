package data_source_test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-configuration/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-configuration/test/mocked/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/nbio/st"
)

func TestSensorDataSource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test Read
			{
				PreConfig: func() {
					util.MockGetSensor(
						"uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateSensorResponseModel("uid", "")},
						),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_sensor" "my_sensor" {
						filter = {
							sensor_id = "uid"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_sensor.my_sensor", "id", "uid"),
					resource.TestCheckResourceAttr("data.uxi_sensor.my_sensor", "name", "name"),
					resource.TestCheckResourceAttr("data.uxi_sensor.my_sensor", "serial", "serial"),
					resource.TestCheckResourceAttr(
						"data.uxi_sensor.my_sensor",
						"model_number",
						"model_number",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_sensor.my_sensor",
						"wifi_mac_address",
						"wifi_mac_address",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_sensor.my_sensor",
						"ethernet_mac_address",
						"ethernet_mac_address",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_sensor.my_sensor",
						"address_note",
						"address_note",
					),
					resource.TestCheckResourceAttr("data.uxi_sensor.my_sensor", "latitude", "0"),
					resource.TestCheckResourceAttr("data.uxi_sensor.my_sensor", "longitude", "0"),
					resource.TestCheckResourceAttr("data.uxi_sensor.my_sensor", "notes", "notes"),
					resource.TestCheckResourceAttr(
						"data.uxi_sensor.my_sensor",
						"pcap_mode",
						"light",
					),
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestSensorDataSource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mock429 *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			// Test Read
			{
				PreConfig: func() {
					mock429 = gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/sensors").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetSensor(
						"uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateSensorResponseModel("uid", "")},
						),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_sensor" "my_sensor" {
						filter = {
							sensor_id = "uid"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_sensor.my_sensor", "id", "uid"),
					func(s *terraform.State) error {
						st.Assert(t, mock429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestSensorDataSourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// 5xx error
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/sensors").
						Reply(500).
						JSON(map[string]interface{}{
							"httpStatusCode": 500,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					data "uxi_sensor" "my_sensor" {
						filter = {
							sensor_id = "uid"
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
					util.MockGetSensor(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_sensor" "my_sensor" {
						filter = {
							sensor_id = "uid"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Could not find specified data source`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
