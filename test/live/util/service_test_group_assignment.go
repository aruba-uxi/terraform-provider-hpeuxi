/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func GetServiceTestGroupAssignment(id string) config_api_client.ServiceTestGroupAssignmentsGetItem {
	result, response, err := Client.ConfigurationAPI.
		ServiceTestGroupAssignmentsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if len(result.Items) != 1 {
		panic("service_test_group_assignment with id `" + id + "` could not be found")
	}

	return result.Items[0]
}

func CheckStateAgainstServiceTestGroupAssignment(
	t *testing.T,
	entity string,
	serviceTestGroupAssignment config_api_client.ServiceTestGroupAssignmentsGetItem,
) resource.TestCheckFunc {
	t.Helper()

	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", serviceTestGroupAssignment.Id),
		resource.TestCheckResourceAttr(entity, "group_id", serviceTestGroupAssignment.Group.Id),
		resource.TestCheckResourceAttr(
			entity,
			"service_test_id",
			serviceTestGroupAssignment.ServiceTest.Id,
		),
	)
}
