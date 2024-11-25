/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
Configuration Api

Nice description goes here

API version: 5.17.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentItem{}

// AgentItem struct for AgentItem
type AgentItem struct {
	Id                 string         `json:"id"`
	Serial             string         `json:"serial"`
	Name               string         `json:"name"`
	ModelNumber        NullableString `json:"modelNumber"`
	WifiMacAddress     NullableString `json:"wifiMacAddress"`
	EthernetMacAddress NullableString `json:"ethernetMacAddress"`
	Notes              NullableString `json:"notes"`
	PcapMode           NullableString `json:"pcapMode"`
	Type               string         `json:"type"`
}

type _AgentItem AgentItem

// NewAgentItem instantiates a new AgentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentItem(
	id string,
	serial string,
	name string,
	modelNumber NullableString,
	wifiMacAddress NullableString,
	ethernetMacAddress NullableString,
	notes NullableString,
	pcapMode NullableString,
	type_ string,
) *AgentItem {
	this := AgentItem{}
	this.Id = id
	this.Serial = serial
	this.Name = name
	this.ModelNumber = modelNumber
	this.WifiMacAddress = wifiMacAddress
	this.EthernetMacAddress = ethernetMacAddress
	this.Notes = notes
	this.PcapMode = pcapMode
	this.Type = type_
	return &this
}

// NewAgentItemWithDefaults instantiates a new AgentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentItemWithDefaults() *AgentItem {
	this := AgentItem{}
	return &this
}

// GetId returns the Id field value
func (o *AgentItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentItem) SetId(v string) {
	o.Id = v
}

// GetSerial returns the Serial field value
func (o *AgentItem) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *AgentItem) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *AgentItem) SetSerial(v string) {
	o.Serial = v
}

// GetName returns the Name field value
func (o *AgentItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AgentItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AgentItem) SetName(v string) {
	o.Name = v
}

// GetModelNumber returns the ModelNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AgentItem) GetModelNumber() string {
	if o == nil || o.ModelNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.ModelNumber.Get()
}

// GetModelNumberOk returns a tuple with the ModelNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentItem) GetModelNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModelNumber.Get(), o.ModelNumber.IsSet()
}

// SetModelNumber sets field value
func (o *AgentItem) SetModelNumber(v string) {
	o.ModelNumber.Set(&v)
}

// GetWifiMacAddress returns the WifiMacAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AgentItem) GetWifiMacAddress() string {
	if o == nil || o.WifiMacAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.WifiMacAddress.Get()
}

// GetWifiMacAddressOk returns a tuple with the WifiMacAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentItem) GetWifiMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WifiMacAddress.Get(), o.WifiMacAddress.IsSet()
}

// SetWifiMacAddress sets field value
func (o *AgentItem) SetWifiMacAddress(v string) {
	o.WifiMacAddress.Set(&v)
}

// GetEthernetMacAddress returns the EthernetMacAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AgentItem) GetEthernetMacAddress() string {
	if o == nil || o.EthernetMacAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.EthernetMacAddress.Get()
}

// GetEthernetMacAddressOk returns a tuple with the EthernetMacAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentItem) GetEthernetMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EthernetMacAddress.Get(), o.EthernetMacAddress.IsSet()
}

// SetEthernetMacAddress sets field value
func (o *AgentItem) SetEthernetMacAddress(v string) {
	o.EthernetMacAddress.Set(&v)
}

// GetNotes returns the Notes field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AgentItem) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}

	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentItem) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// SetNotes sets field value
func (o *AgentItem) SetNotes(v string) {
	o.Notes.Set(&v)
}

// GetPcapMode returns the PcapMode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AgentItem) GetPcapMode() string {
	if o == nil || o.PcapMode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PcapMode.Get()
}

// GetPcapModeOk returns a tuple with the PcapMode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentItem) GetPcapModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PcapMode.Get(), o.PcapMode.IsSet()
}

// SetPcapMode sets field value
func (o *AgentItem) SetPcapMode(v string) {
	o.PcapMode.Set(&v)
}

// GetType returns the Type field value
func (o *AgentItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AgentItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AgentItem) SetType(v string) {
	o.Type = v
}

func (o AgentItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["serial"] = o.Serial
	toSerialize["name"] = o.Name
	toSerialize["modelNumber"] = o.ModelNumber.Get()
	toSerialize["wifiMacAddress"] = o.WifiMacAddress.Get()
	toSerialize["ethernetMacAddress"] = o.EthernetMacAddress.Get()
	toSerialize["notes"] = o.Notes.Get()
	toSerialize["pcapMode"] = o.PcapMode.Get()
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *AgentItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"serial",
		"name",
		"modelNumber",
		"wifiMacAddress",
		"ethernetMacAddress",
		"notes",
		"pcapMode",
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

	varAgentItem := _AgentItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentItem)

	if err != nil {
		return err
	}

	*o = AgentItem(varAgentItem)

	return err
}

type NullableAgentItem struct {
	value *AgentItem
	isSet bool
}

func (v NullableAgentItem) Get() *AgentItem {
	return v.value
}

func (v *NullableAgentItem) Set(val *AgentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentItem(val *AgentItem) *NullableAgentItem {
	return &NullableAgentItem{value: val, isSet: true}
}

func (v NullableAgentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
