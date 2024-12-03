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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ServiceTestGroupAssignmentPostServiceTest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTestGroupAssignmentPostServiceTest{}

// ServiceTestGroupAssignmentPostServiceTest struct for ServiceTestGroupAssignmentPostServiceTest
type ServiceTestGroupAssignmentPostServiceTest struct {
	Id string `json:"id"`
}

type _ServiceTestGroupAssignmentPostServiceTest ServiceTestGroupAssignmentPostServiceTest

// NewServiceTestGroupAssignmentPostServiceTest instantiates a new ServiceTestGroupAssignmentPostServiceTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTestGroupAssignmentPostServiceTest(
	id string,
) *ServiceTestGroupAssignmentPostServiceTest {
	this := ServiceTestGroupAssignmentPostServiceTest{}
	this.Id = id
	return &this
}

// NewServiceTestGroupAssignmentPostServiceTestWithDefaults instantiates a new ServiceTestGroupAssignmentPostServiceTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTestGroupAssignmentPostServiceTestWithDefaults() *ServiceTestGroupAssignmentPostServiceTest {
	this := ServiceTestGroupAssignmentPostServiceTest{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceTestGroupAssignmentPostServiceTest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentPostServiceTest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceTestGroupAssignmentPostServiceTest) SetId(v string) {
	o.Id = v
}

func (o ServiceTestGroupAssignmentPostServiceTest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTestGroupAssignmentPostServiceTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ServiceTestGroupAssignmentPostServiceTest) UnmarshalJSON(data []byte) (err error) {
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

	varServiceTestGroupAssignmentPostServiceTest := _ServiceTestGroupAssignmentPostServiceTest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceTestGroupAssignmentPostServiceTest)

	if err != nil {
		return err
	}

	*o = ServiceTestGroupAssignmentPostServiceTest(varServiceTestGroupAssignmentPostServiceTest)

	return err
}

type NullableServiceTestGroupAssignmentPostServiceTest struct {
	value *ServiceTestGroupAssignmentPostServiceTest
	isSet bool
}

func (v NullableServiceTestGroupAssignmentPostServiceTest) Get() *ServiceTestGroupAssignmentPostServiceTest {
	return v.value
}

func (v *NullableServiceTestGroupAssignmentPostServiceTest) Set(
	val *ServiceTestGroupAssignmentPostServiceTest,
) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTestGroupAssignmentPostServiceTest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTestGroupAssignmentPostServiceTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTestGroupAssignmentPostServiceTest(
	val *ServiceTestGroupAssignmentPostServiceTest,
) *NullableServiceTestGroupAssignmentPostServiceTest {
	return &NullableServiceTestGroupAssignmentPostServiceTest{value: val, isSet: true}
}

func (v NullableServiceTestGroupAssignmentPostServiceTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTestGroupAssignmentPostServiceTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}