/*
Configuration Api

Nice description goes here

API version: 5.13.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentGroupAssignmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentGroupAssignmentResponse{}

// AgentGroupAssignmentResponse struct for AgentGroupAssignmentResponse
type AgentGroupAssignmentResponse struct {
	Id    string `json:"id"`
	Group Group  `json:"group"`
	Agent Agent  `json:"agent"`
	Type  string `json:"type"`
}

type _AgentGroupAssignmentResponse AgentGroupAssignmentResponse

// NewAgentGroupAssignmentResponse instantiates a new AgentGroupAssignmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentGroupAssignmentResponse(
	id string,
	group Group,
	agent Agent,
	type_ string,
) *AgentGroupAssignmentResponse {
	this := AgentGroupAssignmentResponse{}
	this.Id = id
	this.Group = group
	this.Agent = agent
	this.Type = type_
	return &this
}

// NewAgentGroupAssignmentResponseWithDefaults instantiates a new AgentGroupAssignmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentGroupAssignmentResponseWithDefaults() *AgentGroupAssignmentResponse {
	this := AgentGroupAssignmentResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AgentGroupAssignmentResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentGroupAssignmentResponse) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *AgentGroupAssignmentResponse) GetGroup() Group {
	if o == nil {
		var ret Group
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentResponse) GetGroupOk() (*Group, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *AgentGroupAssignmentResponse) SetGroup(v Group) {
	o.Group = v
}

// GetAgent returns the Agent field value
func (o *AgentGroupAssignmentResponse) GetAgent() Agent {
	if o == nil {
		var ret Agent
		return ret
	}

	return o.Agent
}

// GetAgentOk returns a tuple with the Agent field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentResponse) GetAgentOk() (*Agent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent, true
}

// SetAgent sets field value
func (o *AgentGroupAssignmentResponse) SetAgent(v Agent) {
	o.Agent = v
}

// GetType returns the Type field value
func (o *AgentGroupAssignmentResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AgentGroupAssignmentResponse) SetType(v string) {
	o.Type = v
}

func (o AgentGroupAssignmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentGroupAssignmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["group"] = o.Group
	toSerialize["agent"] = o.Agent
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *AgentGroupAssignmentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"group",
		"agent",
		"type",
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

	varAgentGroupAssignmentResponse := _AgentGroupAssignmentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentGroupAssignmentResponse)

	if err != nil {
		return err
	}

	*o = AgentGroupAssignmentResponse(varAgentGroupAssignmentResponse)

	return err
}

type NullableAgentGroupAssignmentResponse struct {
	value *AgentGroupAssignmentResponse
	isSet bool
}

func (v NullableAgentGroupAssignmentResponse) Get() *AgentGroupAssignmentResponse {
	return v.value
}

func (v *NullableAgentGroupAssignmentResponse) Set(val *AgentGroupAssignmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentGroupAssignmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentGroupAssignmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentGroupAssignmentResponse(
	val *AgentGroupAssignmentResponse,
) *NullableAgentGroupAssignmentResponse {
	return &NullableAgentGroupAssignmentResponse{value: val, isSet: true}
}

func (v NullableAgentGroupAssignmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentGroupAssignmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
