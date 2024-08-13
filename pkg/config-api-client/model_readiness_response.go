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

// checks if the ReadinessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadinessResponse{}

// ReadinessResponse struct for ReadinessResponse
type ReadinessResponse struct {
	Data map[string]string `json:"data"`
	Status string `json:"status"`
}

type _ReadinessResponse ReadinessResponse

// NewReadinessResponse instantiates a new ReadinessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadinessResponse(data map[string]string, status string) *ReadinessResponse {
	this := ReadinessResponse{}
	this.Data = data
	this.Status = status
	return &this
}

// NewReadinessResponseWithDefaults instantiates a new ReadinessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadinessResponseWithDefaults() *ReadinessResponse {
	this := ReadinessResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ReadinessResponse) GetData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReadinessResponse) GetDataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ReadinessResponse) SetData(v map[string]string) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *ReadinessResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReadinessResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReadinessResponse) SetStatus(v string) {
	o.Status = v
}

func (o ReadinessResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadinessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ReadinessResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"status",
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

	varReadinessResponse := _ReadinessResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReadinessResponse)

	if err != nil {
		return err
	}

	*o = ReadinessResponse(varReadinessResponse)

	return err
}

type NullableReadinessResponse struct {
	value *ReadinessResponse
	isSet bool
}

func (v NullableReadinessResponse) Get() *ReadinessResponse {
	return v.value
}

func (v *NullableReadinessResponse) Set(val *ReadinessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadinessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadinessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadinessResponse(val *ReadinessResponse) *NullableReadinessResponse {
	return &NullableReadinessResponse{value: val, isSet: true}
}

func (v NullableReadinessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadinessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

