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
	"encoding/json"
	"fmt"
)

// AgentPcapMode the model 'AgentPcapMode'
type AgentPcapMode string

// List of AgentPcapMode
const (
	AGENTPCAPMODE_LIGHT AgentPcapMode = "light"
	AGENTPCAPMODE_FULL  AgentPcapMode = "full"
	AGENTPCAPMODE_OFF   AgentPcapMode = "off"
)

// All allowed values of AgentPcapMode enum
var AllowedAgentPcapModeEnumValues = []AgentPcapMode{
	"light",
	"full",
	"off",
}

func (v *AgentPcapMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgentPcapMode(value)
	for _, existing := range AllowedAgentPcapModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AgentPcapMode", value)
}

// NewAgentPcapModeFromValue returns a pointer to a valid AgentPcapMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentPcapModeFromValue(v string) (*AgentPcapMode, error) {
	ev := AgentPcapMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentPcapMode: valid values are %v", v, AllowedAgentPcapModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentPcapMode) IsValid() bool {
	for _, existing := range AllowedAgentPcapModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentPcapMode value
func (v AgentPcapMode) Ptr() *AgentPcapMode {
	return &v
}

type NullableAgentPcapMode struct {
	value *AgentPcapMode
	isSet bool
}

func (v NullableAgentPcapMode) Get() *AgentPcapMode {
	return v.value
}

func (v *NullableAgentPcapMode) Set(val *AgentPcapMode) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentPcapMode) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentPcapMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentPcapMode(val *AgentPcapMode) *NullableAgentPcapMode {
	return &NullableAgentPcapMode{value: val, isSet: true}
}

func (v NullableAgentPcapMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentPcapMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}