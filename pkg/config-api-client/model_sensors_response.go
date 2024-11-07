/*
Configuration Api

Nice description goes here

API version: 5.4.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SensorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SensorsResponse{}

// SensorsResponse struct for SensorsResponse
type SensorsResponse struct {
	Items []SensorItem   `json:"items"`
	Count int32          `json:"count"`
	Next  NullableString `json:"next"`
}

type _SensorsResponse SensorsResponse

// NewSensorsResponse instantiates a new SensorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensorsResponse(items []SensorItem, count int32, next NullableString) *SensorsResponse {
	this := SensorsResponse{}
	this.Items = items
	this.Count = count
	this.Next = next
	return &this
}

// NewSensorsResponseWithDefaults instantiates a new SensorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorsResponseWithDefaults() *SensorsResponse {
	this := SensorsResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *SensorsResponse) GetItems() []SensorItem {
	if o == nil {
		var ret []SensorItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SensorsResponse) GetItemsOk() ([]SensorItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SensorsResponse) SetItems(v []SensorItem) {
	o.Items = v
}

// GetCount returns the Count field value
func (o *SensorsResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SensorsResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *SensorsResponse) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SensorsResponse) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorsResponse) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *SensorsResponse) SetNext(v string) {
	o.Next.Set(&v)
}

func (o SensorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SensorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["count"] = o.Count
	toSerialize["next"] = o.Next.Get()
	return toSerialize, nil
}

func (o *SensorsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varSensorsResponse := _SensorsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSensorsResponse)

	if err != nil {
		return err
	}

	*o = SensorsResponse(varSensorsResponse)

	return err
}

type NullableSensorsResponse struct {
	value *SensorsResponse
	isSet bool
}

func (v NullableSensorsResponse) Get() *SensorsResponse {
	return v.value
}

func (v *NullableSensorsResponse) Set(val *SensorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorsResponse(val *SensorsResponse) *NullableSensorsResponse {
	return &NullableSensorsResponse{value: val, isSet: true}
}

func (v NullableSensorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
