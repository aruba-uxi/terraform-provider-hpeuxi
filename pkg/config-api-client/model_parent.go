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

// checks if the Parent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Parent{}

// Parent struct for Parent
type Parent struct {
	Id string `json:"id"`
}

type _Parent Parent

// NewParent instantiates a new Parent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParent(id string) *Parent {
	this := Parent{}
	this.Id = id
	return &this
}

// NewParentWithDefaults instantiates a new Parent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParentWithDefaults() *Parent {
	this := Parent{}
	return &this
}

// GetId returns the Id field value
func (o *Parent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Parent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Parent) SetId(v string) {
	o.Id = v
}

func (o Parent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Parent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *Parent) UnmarshalJSON(data []byte) (err error) {
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

	varParent := _Parent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParent)

	if err != nil {
		return err
	}

	*o = Parent(varParent)

	return err
}

type NullableParent struct {
	value *Parent
	isSet bool
}

func (v NullableParent) Get() *Parent {
	return v.value
}

func (v *NullableParent) Set(val *Parent) {
	v.value = val
	v.isSet = true
}

func (v NullableParent) IsSet() bool {
	return v.isSet
}

func (v *NullableParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParent(val *Parent) *NullableParent {
	return &NullableParent{value: val, isSet: true}
}

func (v NullableParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
