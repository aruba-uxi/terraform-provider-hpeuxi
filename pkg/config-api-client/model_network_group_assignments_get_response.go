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

// checks if the NetworkGroupAssignmentsGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkGroupAssignmentsGetResponse{}

// NetworkGroupAssignmentsGetResponse struct for NetworkGroupAssignmentsGetResponse
type NetworkGroupAssignmentsGetResponse struct {
	Items []NetworkGroupAssignmentsGetItem `json:"items"`
	Count int32                            `json:"count"`
	Next  NullableString                   `json:"next"`
}

type _NetworkGroupAssignmentsGetResponse NetworkGroupAssignmentsGetResponse

// NewNetworkGroupAssignmentsGetResponse instantiates a new NetworkGroupAssignmentsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkGroupAssignmentsGetResponse(
	items []NetworkGroupAssignmentsGetItem,
	count int32,
	next NullableString,
) *NetworkGroupAssignmentsGetResponse {
	this := NetworkGroupAssignmentsGetResponse{}
	this.Items = items
	this.Count = count
	this.Next = next
	return &this
}

// NewNetworkGroupAssignmentsGetResponseWithDefaults instantiates a new NetworkGroupAssignmentsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkGroupAssignmentsGetResponseWithDefaults() *NetworkGroupAssignmentsGetResponse {
	this := NetworkGroupAssignmentsGetResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *NetworkGroupAssignmentsGetResponse) GetItems() []NetworkGroupAssignmentsGetItem {
	if o == nil {
		var ret []NetworkGroupAssignmentsGetItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsGetResponse) GetItemsOk() ([]NetworkGroupAssignmentsGetItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *NetworkGroupAssignmentsGetResponse) SetItems(v []NetworkGroupAssignmentsGetItem) {
	o.Items = v
}

// GetCount returns the Count field value
func (o *NetworkGroupAssignmentsGetResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsGetResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *NetworkGroupAssignmentsGetResponse) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkGroupAssignmentsGetResponse) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkGroupAssignmentsGetResponse) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *NetworkGroupAssignmentsGetResponse) SetNext(v string) {
	o.Next.Set(&v)
}

func (o NetworkGroupAssignmentsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkGroupAssignmentsGetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["count"] = o.Count
	toSerialize["next"] = o.Next.Get()
	return toSerialize, nil
}

func (o *NetworkGroupAssignmentsGetResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"count",
		"next",
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

	varNetworkGroupAssignmentsGetResponse := _NetworkGroupAssignmentsGetResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkGroupAssignmentsGetResponse)

	if err != nil {
		return err
	}

	*o = NetworkGroupAssignmentsGetResponse(varNetworkGroupAssignmentsGetResponse)

	return err
}

type NullableNetworkGroupAssignmentsGetResponse struct {
	value *NetworkGroupAssignmentsGetResponse
	isSet bool
}

func (v NullableNetworkGroupAssignmentsGetResponse) Get() *NetworkGroupAssignmentsGetResponse {
	return v.value
}

func (v *NullableNetworkGroupAssignmentsGetResponse) Set(val *NetworkGroupAssignmentsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkGroupAssignmentsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkGroupAssignmentsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkGroupAssignmentsGetResponse(
	val *NetworkGroupAssignmentsGetResponse,
) *NullableNetworkGroupAssignmentsGetResponse {
	return &NullableNetworkGroupAssignmentsGetResponse{value: val, isSet: true}
}

func (v NullableNetworkGroupAssignmentsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkGroupAssignmentsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
