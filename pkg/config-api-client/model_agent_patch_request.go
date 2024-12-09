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
	"encoding/json"
)

// checks if the AgentPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentPatchRequest{}

// AgentPatchRequest Request body for patching an agent.
type AgentPatchRequest struct {
	Name     *string        `json:"name,omitempty"`
	Notes    *string        `json:"notes,omitempty"`
	PcapMode *AgentPcapMode `json:"pcapMode,omitempty"`
}

// NewAgentPatchRequest instantiates a new AgentPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentPatchRequest() *AgentPatchRequest {
	this := AgentPatchRequest{}
	return &this
}

// NewAgentPatchRequestWithDefaults instantiates a new AgentPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentPatchRequestWithDefaults() *AgentPatchRequest {
	this := AgentPatchRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AgentPatchRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPatchRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AgentPatchRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AgentPatchRequest) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AgentPatchRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPatchRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AgentPatchRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *AgentPatchRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetPcapMode returns the PcapMode field value if set, zero value otherwise.
func (o *AgentPatchRequest) GetPcapMode() AgentPcapMode {
	if o == nil || IsNil(o.PcapMode) {
		var ret AgentPcapMode
		return ret
	}
	return *o.PcapMode
}

// GetPcapModeOk returns a tuple with the PcapMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentPatchRequest) GetPcapModeOk() (*AgentPcapMode, bool) {
	if o == nil || IsNil(o.PcapMode) {
		return nil, false
	}
	return o.PcapMode, true
}

// HasPcapMode returns a boolean if a field has been set.
func (o *AgentPatchRequest) HasPcapMode() bool {
	if o != nil && !IsNil(o.PcapMode) {
		return true
	}

	return false
}

// SetPcapMode gets a reference to the given AgentPcapMode and assigns it to the PcapMode field.
func (o *AgentPatchRequest) SetPcapMode(v AgentPcapMode) {
	o.PcapMode = &v
}

func (o AgentPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.PcapMode) {
		toSerialize["pcapMode"] = o.PcapMode
	}
	return toSerialize, nil
}

type NullableAgentPatchRequest struct {
	value *AgentPatchRequest
	isSet bool
}

func (v NullableAgentPatchRequest) Get() *AgentPatchRequest {
	return v.value
}

func (v *NullableAgentPatchRequest) Set(val *AgentPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentPatchRequest(val *AgentPatchRequest) *NullableAgentPatchRequest {
	return &NullableAgentPatchRequest{value: val, isSet: true}
}

func (v NullableAgentPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
