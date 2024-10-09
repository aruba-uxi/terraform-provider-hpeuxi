/*
Configuration Api

Nice description goes here

API version: 1.22.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SensorGroupAssignmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SensorGroupAssignmentResponse{}

// SensorGroupAssignmentResponse struct for SensorGroupAssignmentResponse
type SensorGroupAssignmentResponse struct {
	Id     string  `json:"id"`
	Group  Group   `json:"group"`
	Sensor Sensor  `json:"sensor"`
	Type   *string `json:"type,omitempty"`
}

type _SensorGroupAssignmentResponse SensorGroupAssignmentResponse

// NewSensorGroupAssignmentResponse instantiates a new SensorGroupAssignmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensorGroupAssignmentResponse(id string, group Group, sensor Sensor) *SensorGroupAssignmentResponse {
	this := SensorGroupAssignmentResponse{}
	this.Id = id
	this.Group = group
	this.Sensor = sensor
	var type_ string = "uxi/sensor-group-assignment"
	this.Type = &type_
	return &this
}

// NewSensorGroupAssignmentResponseWithDefaults instantiates a new SensorGroupAssignmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorGroupAssignmentResponseWithDefaults() *SensorGroupAssignmentResponse {
	this := SensorGroupAssignmentResponse{}
	var type_ string = "uxi/sensor-group-assignment"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value
func (o *SensorGroupAssignmentResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SensorGroupAssignmentResponse) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *SensorGroupAssignmentResponse) GetGroup() Group {
	if o == nil {
		var ret Group
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentResponse) GetGroupOk() (*Group, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SensorGroupAssignmentResponse) SetGroup(v Group) {
	o.Group = v
}

// GetSensor returns the Sensor field value
func (o *SensorGroupAssignmentResponse) GetSensor() Sensor {
	if o == nil {
		var ret Sensor
		return ret
	}

	return o.Sensor
}

// GetSensorOk returns a tuple with the Sensor field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentResponse) GetSensorOk() (*Sensor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sensor, true
}

// SetSensor sets field value
func (o *SensorGroupAssignmentResponse) SetSensor(v Sensor) {
	o.Sensor = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SensorGroupAssignmentResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SensorGroupAssignmentResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SensorGroupAssignmentResponse) SetType(v string) {
	o.Type = &v
}

func (o SensorGroupAssignmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SensorGroupAssignmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["group"] = o.Group
	toSerialize["sensor"] = o.Sensor
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *SensorGroupAssignmentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"group",
		"sensor",
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

	varSensorGroupAssignmentResponse := _SensorGroupAssignmentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSensorGroupAssignmentResponse)

	if err != nil {
		return err
	}

	*o = SensorGroupAssignmentResponse(varSensorGroupAssignmentResponse)

	return err
}

type NullableSensorGroupAssignmentResponse struct {
	value *SensorGroupAssignmentResponse
	isSet bool
}

func (v NullableSensorGroupAssignmentResponse) Get() *SensorGroupAssignmentResponse {
	return v.value
}

func (v *NullableSensorGroupAssignmentResponse) Set(val *SensorGroupAssignmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorGroupAssignmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorGroupAssignmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorGroupAssignmentResponse(val *SensorGroupAssignmentResponse) *NullableSensorGroupAssignmentResponse {
	return &NullableSensorGroupAssignmentResponse{value: val, isSet: true}
}

func (v NullableSensorGroupAssignmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorGroupAssignmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
