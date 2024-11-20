/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestSensorGroupAssignmentResource(t *testing.T) {
	const (
		groupName  = "tf_provider_acceptance_test_sensor_assignment_resource"
		group2Name = "tf_provider_acceptance_test_sensor_assignment_resource_two"
	)

	var (
		existingSensorProperties = util.GetSensor(config.SensorId)
		resourceIdBeforeRecreate string
		resourceIdAfterRecreate  string
	)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_sensor" "my_sensor" {
						name 			= "` + existingSensorProperties.Name + `"
						` + util.ConditionalProperty("address_note", existingSensorProperties.AddressNote.Get()) + `
						` + util.ConditionalProperty("notes", existingSensorProperties.Notes.Get()) + `
						` + util.ConditionalProperty("pcap_mode", existingSensorProperties.PcapMode.Get()) + `
					}

					import {
						to = uxi_sensor.my_sensor
						id = "` + config.SensorId + `"
					}

					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id = uxi_sensor.my_sensor.id
						group_id  = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"uxi_sensor_group_assignment.my_sensor_group_assignment",
						"sensor_id",
						config.SensorId,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_sensor_group_assignment.my_sensor_group_assignment",
						"group_id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupName).Id)
							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "uxi_sensor_group_assignment.my_sensor_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIdBeforeRecreate = rs.Primary.ID
						return util.CheckStateAgainstSensorGroupAssignment(
							t,
							"uxi_sensor_group_assignment.my_sensor_group_assignment",
							util.GetSensorGroupAssignment(resourceIdBeforeRecreate),
						)(s)
					},
				),
			},
			// ImportState
			{
				ResourceName:      "uxi_sensor_group_assignment.my_sensor_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_sensor" "my_sensor" {
						name 			= "` + existingSensorProperties.Name + `"
						` + util.ConditionalProperty("address_note", existingSensorProperties.AddressNote.Get()) + `
						` + util.ConditionalProperty("notes", existingSensorProperties.Notes.Get()) + `
						` + util.ConditionalProperty("pcap_mode", existingSensorProperties.PcapMode.Get()) + `
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name = "` + group2Name + `"
					}

					// the assignment update, updated from sensor/group to sensor/group_2
					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id = uxi_sensor.my_sensor.id
						group_id  = uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"uxi_sensor_group_assignment.my_sensor_group_assignment",
						"sensor_id",
						config.SensorId,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_sensor_group_assignment.my_sensor_group_assignment",
						"group_id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(group2Name).Id)
							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "uxi_sensor_group_assignment.my_sensor_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIdAfterRecreate = rs.Primary.ID
						return util.CheckStateAgainstSensorGroupAssignment(
							t,
							"uxi_sensor_group_assignment.my_sensor_group_assignment",
							util.GetSensorGroupAssignment(resourceIdAfterRecreate),
						)(s)
					},
				),
			},
			// Delete sensor-group assignments and remove sensors from state
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
		CheckDestroy: func(s *terraform.State) error {
			assert.Equal(t, util.GetGroupByName(groupName), nil)
			assert.Equal(t, util.GetGroupByName(group2Name), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIdBeforeRecreate), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIdAfterRecreate), nil)
			return nil
		},
	})
}
