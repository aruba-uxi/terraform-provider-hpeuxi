/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
Configuration Api

Nice description goes here

API version: 5.17.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SensorGroupAssignmentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SensorGroupAssignmentsResponse{}

// SensorGroupAssignmentsResponse struct for SensorGroupAssignmentsResponse
type SensorGroupAssignmentsResponse struct {
	Items []SensorGroupAssignmentsItem `json:"items"`
	Count int32                        `json:"count"`
	Next  NullableString               `json:"next"`
}

type _SensorGroupAssignmentsResponse SensorGroupAssignmentsResponse

// NewSensorGroupAssignmentsResponse instantiates a new SensorGroupAssignmentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensorGroupAssignmentsResponse(
	items []SensorGroupAssignmentsItem,
	count int32,
	next NullableString,
) *SensorGroupAssignmentsResponse {
	this := SensorGroupAssignmentsResponse{}
	this.Items = items
	this.Count = count
	this.Next = next
	return &this
}

// NewSensorGroupAssignmentsResponseWithDefaults instantiates a new SensorGroupAssignmentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorGroupAssignmentsResponseWithDefaults() *SensorGroupAssignmentsResponse {
	this := SensorGroupAssignmentsResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *SensorGroupAssignmentsResponse) GetItems() []SensorGroupAssignmentsItem {
	if o == nil {
		var ret []SensorGroupAssignmentsItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentsResponse) GetItemsOk() ([]SensorGroupAssignmentsItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SensorGroupAssignmentsResponse) SetItems(v []SensorGroupAssignmentsItem) {
	o.Items = v
}

// GetCount returns the Count field value
func (o *SensorGroupAssignmentsResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentsResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *SensorGroupAssignmentsResponse) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SensorGroupAssignmentsResponse) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorGroupAssignmentsResponse) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *SensorGroupAssignmentsResponse) SetNext(v string) {
	o.Next.Set(&v)
}

func (o SensorGroupAssignmentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SensorGroupAssignmentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["count"] = o.Count
	toSerialize["next"] = o.Next.Get()
	return toSerialize, nil
}

func (o *SensorGroupAssignmentsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varSensorGroupAssignmentsResponse := _SensorGroupAssignmentsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSensorGroupAssignmentsResponse)

	if err != nil {
		return err
	}

	*o = SensorGroupAssignmentsResponse(varSensorGroupAssignmentsResponse)

	return err
}

type NullableSensorGroupAssignmentsResponse struct {
	value *SensorGroupAssignmentsResponse
	isSet bool
}

func (v NullableSensorGroupAssignmentsResponse) Get() *SensorGroupAssignmentsResponse {
	return v.value
}

func (v *NullableSensorGroupAssignmentsResponse) Set(val *SensorGroupAssignmentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorGroupAssignmentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorGroupAssignmentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorGroupAssignmentsResponse(
	val *SensorGroupAssignmentsResponse,
) *NullableSensorGroupAssignmentsResponse {
	return &NullableSensorGroupAssignmentsResponse{value: val, isSet: true}
}

func (v NullableSensorGroupAssignmentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorGroupAssignmentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
