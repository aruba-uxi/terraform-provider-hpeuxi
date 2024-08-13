/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the StatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusResponse{}

// StatusResponse struct for StatusResponse
type StatusResponse struct {
	Name string `json:"name"`
	Version string `json:"version"`
}

type _StatusResponse StatusResponse

// NewStatusResponse instantiates a new StatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponse(name string, version string) *StatusResponse {
	this := StatusResponse{}
	this.Name = name
	this.Version = version
	return &this
}

// NewStatusResponseWithDefaults instantiates a new StatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseWithDefaults() *StatusResponse {
	this := StatusResponse{}
	return &this
}

// GetName returns the Name field value
func (o *StatusResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StatusResponse) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *StatusResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *StatusResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *StatusResponse) SetVersion(v string) {
	o.Version = v
}

func (o StatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *StatusResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"version",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varStatusResponse := _StatusResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStatusResponse)

	if err != nil {
		return err
	}

	*o = StatusResponse(varStatusResponse)

	return err
}

type NullableStatusResponse struct {
	value *StatusResponse
	isSet bool
}

func (v NullableStatusResponse) Get() *StatusResponse {
	return v.value
}

func (v *NullableStatusResponse) Set(val *StatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponse(val *StatusResponse) *NullableStatusResponse {
	return &NullableStatusResponse{value: val, isSet: true}
}

func (v NullableStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

