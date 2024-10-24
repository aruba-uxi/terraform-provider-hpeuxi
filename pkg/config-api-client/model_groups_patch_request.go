/*
Configuration Api

Nice description goes here

API version: 2.9.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GroupsPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsPatchRequest{}

// GroupsPatchRequest struct for GroupsPatchRequest
type GroupsPatchRequest struct {
	Name string `json:"name"`
}

type _GroupsPatchRequest GroupsPatchRequest

// NewGroupsPatchRequest instantiates a new GroupsPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsPatchRequest(name string) *GroupsPatchRequest {
	this := GroupsPatchRequest{}
	this.Name = name
	return &this
}

// NewGroupsPatchRequestWithDefaults instantiates a new GroupsPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsPatchRequestWithDefaults() *GroupsPatchRequest {
	this := GroupsPatchRequest{}
	return &this
}

// GetName returns the Name field value
func (o *GroupsPatchRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupsPatchRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupsPatchRequest) SetName(v string) {
	o.Name = v
}

func (o GroupsPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GroupsPatchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varGroupsPatchRequest := _GroupsPatchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupsPatchRequest)

	if err != nil {
		return err
	}

	*o = GroupsPatchRequest(varGroupsPatchRequest)

	return err
}

type NullableGroupsPatchRequest struct {
	value *GroupsPatchRequest
	isSet bool
}

func (v NullableGroupsPatchRequest) Get() *GroupsPatchRequest {
	return v.value
}

func (v *NullableGroupsPatchRequest) Set(val *GroupsPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsPatchRequest(val *GroupsPatchRequest) *NullableGroupsPatchRequest {
	return &NullableGroupsPatchRequest{value: val, isSet: true}
}

func (v NullableGroupsPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
