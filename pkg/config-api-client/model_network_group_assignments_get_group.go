/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
Configuration Api

Nice description goes here

API version: 5.18.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NetworkGroupAssignmentsGetGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkGroupAssignmentsGetGroup{}

// NetworkGroupAssignmentsGetGroup struct for NetworkGroupAssignmentsGetGroup
type NetworkGroupAssignmentsGetGroup struct {
	Id string `json:"id"`
}

type _NetworkGroupAssignmentsGetGroup NetworkGroupAssignmentsGetGroup

// NewNetworkGroupAssignmentsGetGroup instantiates a new NetworkGroupAssignmentsGetGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkGroupAssignmentsGetGroup(id string) *NetworkGroupAssignmentsGetGroup {
	this := NetworkGroupAssignmentsGetGroup{}
	this.Id = id
	return &this
}

// NewNetworkGroupAssignmentsGetGroupWithDefaults instantiates a new NetworkGroupAssignmentsGetGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkGroupAssignmentsGetGroupWithDefaults() *NetworkGroupAssignmentsGetGroup {
	this := NetworkGroupAssignmentsGetGroup{}
	return &this
}

// GetId returns the Id field value
func (o *NetworkGroupAssignmentsGetGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsGetGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworkGroupAssignmentsGetGroup) SetId(v string) {
	o.Id = v
}

func (o NetworkGroupAssignmentsGetGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkGroupAssignmentsGetGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *NetworkGroupAssignmentsGetGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varNetworkGroupAssignmentsGetGroup := _NetworkGroupAssignmentsGetGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkGroupAssignmentsGetGroup)

	if err != nil {
		return err
	}

	*o = NetworkGroupAssignmentsGetGroup(varNetworkGroupAssignmentsGetGroup)

	return err
}

type NullableNetworkGroupAssignmentsGetGroup struct {
	value *NetworkGroupAssignmentsGetGroup
	isSet bool
}

func (v NullableNetworkGroupAssignmentsGetGroup) Get() *NetworkGroupAssignmentsGetGroup {
	return v.value
}

func (v *NullableNetworkGroupAssignmentsGetGroup) Set(val *NetworkGroupAssignmentsGetGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkGroupAssignmentsGetGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkGroupAssignmentsGetGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkGroupAssignmentsGetGroup(
	val *NetworkGroupAssignmentsGetGroup,
) *NullableNetworkGroupAssignmentsGetGroup {
	return &NullableNetworkGroupAssignmentsGetGroup{value: val, isSet: true}
}

func (v NullableNetworkGroupAssignmentsGetGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkGroupAssignmentsGetGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
