/*
Configuration Api

Nice description goes here

API version: 2.9.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"encoding/json"
	"fmt"
)

// IssueSubject the model 'IssueSubject'
type IssueSubject string

// List of IssueSubject
const (
	FIELD           IssueSubject = "field"
	HEADER          IssueSubject = "header"
	QUERY_PARAMETER IssueSubject = "query.parameter"
)

// All allowed values of IssueSubject enum
var AllowedIssueSubjectEnumValues = []IssueSubject{
	"field",
	"header",
	"query.parameter",
}

func (v *IssueSubject) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IssueSubject(value)
	for _, existing := range AllowedIssueSubjectEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IssueSubject", value)
}

// NewIssueSubjectFromValue returns a pointer to a valid IssueSubject
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIssueSubjectFromValue(v string) (*IssueSubject, error) {
	ev := IssueSubject(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IssueSubject: valid values are %v", v, AllowedIssueSubjectEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IssueSubject) IsValid() bool {
	for _, existing := range AllowedIssueSubjectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueSubject value
func (v IssueSubject) Ptr() *IssueSubject {
	return &v
}

type NullableIssueSubject struct {
	value *IssueSubject
	isSet bool
}

func (v NullableIssueSubject) Get() *IssueSubject {
	return v.value
}

func (v *NullableIssueSubject) Set(val *IssueSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueSubject(val *IssueSubject) *NullableIssueSubject {
	return &NullableIssueSubject{value: val, isSet: true}
}

func (v NullableIssueSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
