/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func TestSensorResource(t *testing.T) {
	originalSensor := util.GetSensor(config.SensorID)
	updatedSensor := originalSensor
	updatedSensor.Notes = *config_api_client.NewNullableString(
		config_api_client.PtrString("tf_provider_acceptance_test_update_notes"),
	)
	updatedSensor.AddressNote = *config_api_client.NewNullableString(
		config_api_client.PtrString("tf_provider_acceptance_test_update_address_note"),
	)
	updatedSensor.PcapMode = *config_api_client.NewNullableSensorPcapMode(
		config_api_client.SENSORPCAPMODE_OFF.Ptr(),
	)

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
					resource "hpeuxi_sensor" "my_sensor" {
						name = "` + originalSensor.Name + `"
					}`,

				ExpectError: regexp.MustCompile(
					`creating a sensor is not supported; sensors can only be imported`,
				),
			},
			// Importing a sensor
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_sensor" "my_sensor" {
						name = "` + originalSensor.Name + `"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "` + config.SensorID + `"
					}`,

				Check: shared.CheckStateAgainstSensor(t, "hpeuxi_sensor.my_sensor", originalSensor),
			},
			// ImportState testing
			{
				ResourceName:      "hpeuxi_sensor.my_sensor",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update testing
			{
				Config: provider.ProviderConfig + `
				resource "hpeuxi_sensor" "my_sensor" {
					name 		 = "` + updatedSensor.Name + `"
					address_note = "` + updatedSensor.GetAddressNote() + `"
					notes 		 = "` + updatedSensor.GetNotes() + `"
					pcap_mode 	 = "` + string(updatedSensor.GetPcapMode()) + `"
				}`,
				Check: shared.CheckStateAgainstSensor(t, "hpeuxi_sensor.my_sensor", updatedSensor),
			},
			// Update sensor back to original
			{
				Config: provider.ProviderConfig + `
				resource "hpeuxi_sensor" "my_sensor" {
					name 		 = "` + originalSensor.Name + `"
					address_note = "` + originalSensor.GetAddressNote() + `"
					notes 		 = "` + originalSensor.GetNotes() + `"
					pcap_mode 	 = "` + string(originalSensor.GetPcapMode()) + `"
				}`,
				Check: shared.CheckStateAgainstSensor(t, "hpeuxi_sensor.my_sensor", originalSensor),
			},
			// Deleting a sensor is not allowed
			{
				Config: provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(
					`deleting a sensor is not supported; sensors can only removed from state`,
				),
			},
			// Remove sensor from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}
