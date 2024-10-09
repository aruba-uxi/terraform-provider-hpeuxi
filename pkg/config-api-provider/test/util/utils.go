package util

import (
	"encoding/json"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/h2non/gock"
)

func GenerateSensorResponseModel(uid string, postfix string) resources.SensorResponseModel {
	return resources.SensorResponseModel{
		UID:                uid,
		Serial:             "serial" + postfix,
		Name:               "name" + postfix,
		ModelNumber:        "model_number" + postfix,
		WifiMacAddress:     "wifi_mac_address" + postfix,
		EthernetMacAddress: "ethernet_mac_address" + postfix,
		AddressNote:        "address_note" + postfix,
		Longitude:          "longitude" + postfix,
		Latitude:           "latitude" + postfix,
		Notes:              "notes" + postfix,
		PCapMode:           "light" + postfix,
	}
}

func GenerateAgentResponseModel(uid string, postfix string) resources.AgentResponseModel {
	return resources.AgentResponseModel{
		UID:                uid,
		Serial:             "serial" + postfix,
		Name:               "name" + postfix,
		ModelNumber:        "model_number" + postfix,
		WifiMacAddress:     "wifi_mac_address" + postfix,
		EthernetMacAddress: "ethernet_mac_address" + postfix,
		Notes:              "notes" + postfix,
		PCapMode:           "light" + postfix,
	}
}

func GenerateGroupResponseGetModel(uid string, nonReplacementFieldPostfix string, replacementFieldPostfix string) config_api_client.GroupsGetItem {
	parentId := "parent_uid" + replacementFieldPostfix

	return config_api_client.GroupsGetItem{
		Id:     uid,
		Name:   "name" + nonReplacementFieldPostfix,
		Parent: *config_api_client.NewNullableParent(&config_api_client.Parent{Id: parentId}),
		Path:   parentId + "." + uid,
	}
}

func GenerateGroupResponseModel(uid string, nonReplacementFieldPostfix string, replacementFieldPostfix string) map[string]interface{} {
	parentId := "parent_uid" + replacementFieldPostfix

	return map[string]interface{}{
		"id":     uid,
		"name":   "name" + nonReplacementFieldPostfix,
		"parent": map[string]string{"id": parentId},
		"path":   parentId + "." + uid,
	}
}

func GenerateGroupResponsePostModel(uid string, nonReplacementFieldPostfix string, replacementFieldPostfix string) config_api_client.GroupsPostResponse {
	parentId := "parent_uid" + replacementFieldPostfix

	return config_api_client.GroupsPostResponse{
		Id:     uid,
		Name:   "name" + nonReplacementFieldPostfix,
		Parent: *config_api_client.NewParent(parentId),
		Path:   parentId + "." + uid,
	}
}

func GeneratePaginatedResponse(items []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"items": items,
		"next":  nil,
		"count": len(items),
	}
}

func GenerateServiceTestResponseModel(uid string, postfix string) resources.ServiceTestResponseModel {
	return resources.ServiceTestResponseModel{
		Uid:       uid,
		Category:  "external" + postfix,
		Title:     "title" + postfix,
		Target:    "target" + postfix,
		Template:  "template" + postfix,
		IsEnabled: true,
	}
}

func GenerateWiredNetworkResponse(uid string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":                   uid,
		"name":                 "alias" + postfix,
		"createdAt":            "2024-09-11T12:00:00.000Z",
		"updatedAt":            "2024-09-11T12:00:00.000Z",
		"ipVersion":            "ip_version" + postfix,
		"security":             "security" + postfix,
		"dnsLookupDomain":      "dns_lookup_domain" + postfix,
		"disableEdns":          false,
		"useDns64":             false,
		"externalConnectivity": false,
		"vLanId":               123,
	}
}

func GenerateWirelessNetworkResponse(uid string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":                   uid,
		"ssid":                 "ssid" + postfix,
		"createdAt":            "2024-09-11T12:00:00.000Z",
		"updatedAt":            "2024-09-11T12:00:00.000Z",
		"name":                 "alias" + postfix,
		"ipVersion":            "ip_version" + postfix,
		"security":             "security" + postfix,
		"hidden":               false,
		"bandLocking":          "band_locking" + postfix,
		"dnsLookupDomain":      "dns_lookup_domain" + postfix,
		"disableEdns":          false,
		"useDns64":             false,
		"externalConnectivity": false,
	}
}

func GenerateSensorGroupAssignmentPostResponse(uid string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":     uid,
		"group":  map[string]string{"id": "group_uid" + postfix},
		"sensor": map[string]string{"id": "sensor_uid" + postfix},
	}
}

func GenerateSensorGroupAssignmentGetResponse(uid string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":     uid,
		"group":  map[string]string{"id": "group_uid" + postfix},
		"sensor": map[string]string{"id": "sensor_uid" + postfix},
	}
}

func GenerateAgentGroupAssignmentResponse(uid string, postfix string) resources.AgentGroupAssignmentResponseModel {
	return resources.AgentGroupAssignmentResponseModel{
		UID:      uid,
		GroupUID: "group_uid" + postfix,
		AgentUID: "agent_uid" + postfix,
	}
}

func GenerateNetworkGroupAssignmentResponse(uid string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":      uid,
		"group":   map[string]string{"id": "group_uid" + postfix},
		"network": map[string]string{"id": "network_uid" + postfix},
	}
}

func GenerateNetworkGroupAssignmentGetResponse(uid string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":      uid,
		"group":   map[string]string{"id": "group_uid" + postfix},
		"network": map[string]string{"id": "network_uid" + postfix},
	}
}

func GenerateServiceTestGroupAssignmentResponse(uid string, postfix string) resources.ServiceTestGroupAssignmentResponseModel {
	return resources.ServiceTestGroupAssignmentResponseModel{
		UID:            uid,
		GroupUID:       "group_uid" + postfix,
		ServiceTestUID: "service_test_uid" + postfix,
	}
}

// Converts a struct to a map while maintaining the json alias as keys
func StructToMap(obj interface{}) map[string]interface{} {
	data, _ := json.Marshal(obj) // Convert to a json string

	newMap := map[string]interface{}{}

	_ = json.Unmarshal(data, &newMap) // Convert to a map
	return newMap
}

func MockOAuth() *gock.Response {
	return gock.New("https://test.sso.common.cloud.hpe.com").
		Post("/as/token.oauth2").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").
		Persist().
		Reply(200).
		JSON(map[string]interface{}{
			"access_token": "mock_token",
			"token_type":   "bearer",
			"expires_in":   3600,
		})

}

func MockPostGroup(response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Post("/uxi/v1alpha1/groups").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(200).
		JSON(response)
}

func MockGetGroup(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/uxi/v1alpha1/groups").
		MatchHeader("Authorization", "mock_token").
		MatchParam("uid", uid).
		Times(times).
		Reply(200).
		JSON(response)
}

func MockUpdateGroup(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Patch("/uxi/v1alpha1/groups/"+uid).
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(200).
		JSON(response)
}

func MockGetWiredNetwork(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/uxi/v1alpha1/wired-networks").
		MatchHeader("Authorization", "mock_token").
		MatchParam("uid", uid).
		Times(times).
		Reply(200).
		JSON(response)
}

func MockGetWirelessNetwork(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/uxi/v1alpha1/wireless-networks").
		MatchHeader("Authorization", "mock_token").
		MatchParam("uid", uid).
		Times(times).
		Reply(200).
		JSON(response)
}

func MockGetSensorGroupAssignment(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/uxi/v1alpha1/sensor-group-assignments").
		MatchHeader("Authorization", "mock_token").
		MatchParam("uid", uid).
		Times(times).
		Reply(200).
		JSON(response)
}

func MockPostSensorGroupAssignment(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Post("/uxi/v1alpha1/sensor-group-assignments").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(200).
		JSON(response)
}

func MockDeleteSensorGroupAssignment(uid string, times int) {
	gock.New("https://test.api.capenetworks.com").
		Delete("/uxi/v1alpha1/sensor-group-assignments/"+uid).
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(204)
}

func MockGetNetworkGroupAssignment(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/uxi/v1alpha1/network-group-assignments").
		MatchHeader("Authorization", "mock_token").
		MatchParam("uid", uid).
		Times(times).
		Reply(200).
		JSON(response)
}

func MockPostNetworkGroupAssignment(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Post("/uxi/v1alpha1/network-group-assignments").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(200).
		JSON(response)
}

func MockDeleteNetworkGroupAssignment(uid string, times int) {
	gock.New("https://test.api.capenetworks.com").
		Delete("/uxi/v1alpha1/network-group-assignments/"+uid).
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(204)
}

var RateLimitingHeaders = map[string]string{
	"X-RateLimit-Limit":     "100",
	"X-RateLimit-Remaining": "0",
	"X-RateLimit-Reset":     "1",
}
