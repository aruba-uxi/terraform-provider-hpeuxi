/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"context"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func GetWiredNetwork(id string) config_api_client.WiredNetworksGetItem {
	result, response, err := Client.ConfigurationAPI.
		WiredNetworksGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if len(result.Items) != 1 {
		panic("wired_network with id `" + id + "` could not be found")
	}

	return result.Items[0]
}
