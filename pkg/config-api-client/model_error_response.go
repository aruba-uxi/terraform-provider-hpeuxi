/*
Copyright 2025 Hewlett Packard Enterprise Development LP.
*/

/*
HPE Aruba Networking UXI Configuration

This document outlines the API contracts for HPE Aruba Networking UXI.

API version: 6.3.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponse{}

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	DebugId        string        `json:"debugId"`
	ErrorCode      string        `json:"errorCode"`
	HttpStatusCode int32         `json:"httpStatusCode"`
	Message        string        `json:"message"`
	ErrorDetails   []ErrorDetail `json:"errorDetails,omitempty"`
}

type _ErrorResponse ErrorResponse

// NewErrorResponse instantiates a new ErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponse(
	debugId string,
	errorCode string,
	httpStatusCode int32,
	message string,
) *ErrorResponse {
	this := ErrorResponse{}
	this.DebugId = debugId
	this.ErrorCode = errorCode
	this.HttpStatusCode = httpStatusCode
	this.Message = message
	return &this
}

// NewErrorResponseWithDefaults instantiates a new ErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// GetDebugId returns the DebugId field value
func (o *ErrorResponse) GetDebugId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DebugId
}

// GetDebugIdOk returns a tuple with the DebugId field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetDebugIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DebugId, true
}

// SetDebugId sets field value
func (o *ErrorResponse) SetDebugId(v string) {
	o.DebugId = v
}

// GetErrorCode returns the ErrorCode field value
func (o *ErrorResponse) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ErrorResponse) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetHttpStatusCode returns the HttpStatusCode field value
func (o *ErrorResponse) GetHttpStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetHttpStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpStatusCode, true
}

// SetHttpStatusCode sets field value
func (o *ErrorResponse) SetHttpStatusCode(v int32) {
	o.HttpStatusCode = v
}

// GetMessage returns the Message field value
func (o *ErrorResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorResponse) SetMessage(v string) {
	o.Message = v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *ErrorResponse) GetErrorDetails() []ErrorDetail {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret []ErrorDetail
		return ret
	}
	return o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrorDetailsOk() ([]ErrorDetail, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *ErrorResponse) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given []ErrorDetail and assigns it to the ErrorDetails field.
func (o *ErrorResponse) SetErrorDetails(v []ErrorDetail) {
	o.ErrorDetails = v
}

func (o ErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["debugId"] = o.DebugId
	toSerialize["errorCode"] = o.ErrorCode
	toSerialize["httpStatusCode"] = o.HttpStatusCode
	toSerialize["message"] = o.Message
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	return toSerialize, nil
}

func (o *ErrorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"debugId",
		"errorCode",
		"httpStatusCode",
		"message",
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

	varErrorResponse := _ErrorResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorResponse)

	if err != nil {
		return err
	}

	*o = ErrorResponse(varErrorResponse)

	return err
}

type NullableErrorResponse struct {
	value *ErrorResponse
	isSet bool
}

func (v NullableErrorResponse) Get() *ErrorResponse {
	return v.value
}

func (v *NullableErrorResponse) Set(val *ErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponse(val *ErrorResponse) *NullableErrorResponse {
	return &NullableErrorResponse{value: val, isSet: true}
}

func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
