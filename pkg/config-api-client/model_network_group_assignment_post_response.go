/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
HPE Aruba Networking UXI Configuration

This document outlines the API contracts for HPE Aruba Networking UXI.

API version: 5.21.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NetworkGroupAssignmentPostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkGroupAssignmentPostResponse{}

// NetworkGroupAssignmentPostResponse struct for NetworkGroupAssignmentPostResponse
type NetworkGroupAssignmentPostResponse struct {
	Id      string                            `json:"id"`
	Group   NetworkGroupAssignmentPostGroup   `json:"group"`
	Network NetworkGroupAssignmentPostNetwork `json:"network"`
	Type    string                            `json:"type"`
}

type _NetworkGroupAssignmentPostResponse NetworkGroupAssignmentPostResponse

// NewNetworkGroupAssignmentPostResponse instantiates a new NetworkGroupAssignmentPostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkGroupAssignmentPostResponse(
	id string,
	group NetworkGroupAssignmentPostGroup,
	network NetworkGroupAssignmentPostNetwork,
	type_ string,
) *NetworkGroupAssignmentPostResponse {
	this := NetworkGroupAssignmentPostResponse{}
	this.Id = id
	this.Group = group
	this.Network = network
	this.Type = type_
	return &this
}

// NewNetworkGroupAssignmentPostResponseWithDefaults instantiates a new NetworkGroupAssignmentPostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkGroupAssignmentPostResponseWithDefaults() *NetworkGroupAssignmentPostResponse {
	this := NetworkGroupAssignmentPostResponse{}
	return &this
}

// GetId returns the Id field value
func (o *NetworkGroupAssignmentPostResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentPostResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworkGroupAssignmentPostResponse) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *NetworkGroupAssignmentPostResponse) GetGroup() NetworkGroupAssignmentPostGroup {
	if o == nil {
		var ret NetworkGroupAssignmentPostGroup
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentPostResponse) GetGroupOk() (*NetworkGroupAssignmentPostGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *NetworkGroupAssignmentPostResponse) SetGroup(v NetworkGroupAssignmentPostGroup) {
	o.Group = v
}

// GetNetwork returns the Network field value
func (o *NetworkGroupAssignmentPostResponse) GetNetwork() NetworkGroupAssignmentPostNetwork {
	if o == nil {
		var ret NetworkGroupAssignmentPostNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentPostResponse) GetNetworkOk() (*NetworkGroupAssignmentPostNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *NetworkGroupAssignmentPostResponse) SetNetwork(v NetworkGroupAssignmentPostNetwork) {
	o.Network = v
}

// GetType returns the Type field value
func (o *NetworkGroupAssignmentPostResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentPostResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworkGroupAssignmentPostResponse) SetType(v string) {
	o.Type = v
}

func (o NetworkGroupAssignmentPostResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkGroupAssignmentPostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["group"] = o.Group
	toSerialize["network"] = o.Network
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *NetworkGroupAssignmentPostResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"group",
		"network",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNetworkGroupAssignmentPostResponse := _NetworkGroupAssignmentPostResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkGroupAssignmentPostResponse)

	if err != nil {
		return err
	}

	*o = NetworkGroupAssignmentPostResponse(varNetworkGroupAssignmentPostResponse)

	return err
}

type NullableNetworkGroupAssignmentPostResponse struct {
	value *NetworkGroupAssignmentPostResponse
	isSet bool
}

func (v NullableNetworkGroupAssignmentPostResponse) Get() *NetworkGroupAssignmentPostResponse {
	return v.value
}

func (v *NullableNetworkGroupAssignmentPostResponse) Set(val *NetworkGroupAssignmentPostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkGroupAssignmentPostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkGroupAssignmentPostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkGroupAssignmentPostResponse(
	val *NetworkGroupAssignmentPostResponse,
) *NullableNetworkGroupAssignmentPostResponse {
	return &NullableNetworkGroupAssignmentPostResponse{value: val, isSet: true}
}

func (v NullableNetworkGroupAssignmentPostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkGroupAssignmentPostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
