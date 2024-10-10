/*
Configuration Api

Nice description goes here

API version: 2.0.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Sensor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sensor{}

// Sensor struct for Sensor
type Sensor struct {
	Id string `json:"id"`
}

type _Sensor Sensor

// NewSensor instantiates a new Sensor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensor(id string) *Sensor {
	this := Sensor{}
	this.Id = id
	return &this
}

// NewSensorWithDefaults instantiates a new Sensor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorWithDefaults() *Sensor {
	this := Sensor{}
	return &this
}

// GetId returns the Id field value
func (o *Sensor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Sensor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Sensor) SetId(v string) {
	o.Id = v
}

func (o Sensor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sensor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *Sensor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varSensor := _Sensor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSensor)

	if err != nil {
		return err
	}

	*o = Sensor(varSensor)

	return err
}

type NullableSensor struct {
	value *Sensor
	isSet bool
}

func (v NullableSensor) Get() *Sensor {
	return v.value
}

func (v *NullableSensor) Set(val *Sensor) {
	v.value = val
	v.isSet = true
}

func (v NullableSensor) IsSet() bool {
	return v.isSet
}

func (v *NullableSensor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensor(val *Sensor) *NullableSensor {
	return &NullableSensor{value: val, isSet: true}
}

func (v NullableSensor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
