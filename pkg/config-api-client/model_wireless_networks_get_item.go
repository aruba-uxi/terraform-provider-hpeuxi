/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
HPE Aruba Networking UXI Configuration

This document outlines the API contracts for HPE Aruba Networking UXI.

API version: 5.19.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the WirelessNetworksGetItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WirelessNetworksGetItem{}

// WirelessNetworksGetItem struct for WirelessNetworksGetItem
type WirelessNetworksGetItem struct {
	Id                   string         `json:"id"`
	Name                 string         `json:"name"`
	Ssid                 string         `json:"ssid"`
	Security             NullableString `json:"security"`
	IpVersion            IpVersion      `json:"ipVersion"`
	CreatedAt            time.Time      `json:"createdAt"`
	UpdatedAt            time.Time      `json:"updatedAt"`
	Hidden               bool           `json:"hidden"`
	BandLocking          string         `json:"bandLocking"`
	DnsLookupDomain      NullableString `json:"dnsLookupDomain"`
	DisableEdns          bool           `json:"disableEdns"`
	UseDns64             bool           `json:"useDns64"`
	ExternalConnectivity bool           `json:"externalConnectivity"`
	Type                 string         `json:"type"`
}

type _WirelessNetworksGetItem WirelessNetworksGetItem

// NewWirelessNetworksGetItem instantiates a new WirelessNetworksGetItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWirelessNetworksGetItem(
	id string,
	name string,
	ssid string,
	security NullableString,
	ipVersion IpVersion,
	createdAt time.Time,
	updatedAt time.Time,
	hidden bool,
	bandLocking string,
	dnsLookupDomain NullableString,
	disableEdns bool,
	useDns64 bool,
	externalConnectivity bool,
	type_ string,
) *WirelessNetworksGetItem {
	this := WirelessNetworksGetItem{}
	this.Id = id
	this.Name = name
	this.Ssid = ssid
	this.Security = security
	this.IpVersion = ipVersion
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Hidden = hidden
	this.BandLocking = bandLocking
	this.DnsLookupDomain = dnsLookupDomain
	this.DisableEdns = disableEdns
	this.UseDns64 = useDns64
	this.ExternalConnectivity = externalConnectivity
	this.Type = type_
	return &this
}

// NewWirelessNetworksGetItemWithDefaults instantiates a new WirelessNetworksGetItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWirelessNetworksGetItemWithDefaults() *WirelessNetworksGetItem {
	this := WirelessNetworksGetItem{}
	return &this
}

// GetId returns the Id field value
func (o *WirelessNetworksGetItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WirelessNetworksGetItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *WirelessNetworksGetItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WirelessNetworksGetItem) SetName(v string) {
	o.Name = v
}

// GetSsid returns the Ssid field value
func (o *WirelessNetworksGetItem) GetSsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *WirelessNetworksGetItem) SetSsid(v string) {
	o.Ssid = v
}

// GetSecurity returns the Security field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WirelessNetworksGetItem) GetSecurity() string {
	if o == nil || o.Security.Get() == nil {
		var ret string
		return ret
	}

	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessNetworksGetItem) GetSecurityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// SetSecurity sets field value
func (o *WirelessNetworksGetItem) SetSecurity(v string) {
	o.Security.Set(&v)
}

// GetIpVersion returns the IpVersion field value
func (o *WirelessNetworksGetItem) GetIpVersion() IpVersion {
	if o == nil {
		var ret IpVersion
		return ret
	}

	return o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetIpVersionOk() (*IpVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpVersion, true
}

// SetIpVersion sets field value
func (o *WirelessNetworksGetItem) SetIpVersion(v IpVersion) {
	o.IpVersion = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WirelessNetworksGetItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WirelessNetworksGetItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WirelessNetworksGetItem) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WirelessNetworksGetItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetHidden returns the Hidden field value
func (o *WirelessNetworksGetItem) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *WirelessNetworksGetItem) SetHidden(v bool) {
	o.Hidden = v
}

// GetBandLocking returns the BandLocking field value
func (o *WirelessNetworksGetItem) GetBandLocking() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BandLocking
}

// GetBandLockingOk returns a tuple with the BandLocking field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetBandLockingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BandLocking, true
}

// SetBandLocking sets field value
func (o *WirelessNetworksGetItem) SetBandLocking(v string) {
	o.BandLocking = v
}

// GetDnsLookupDomain returns the DnsLookupDomain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WirelessNetworksGetItem) GetDnsLookupDomain() string {
	if o == nil || o.DnsLookupDomain.Get() == nil {
		var ret string
		return ret
	}

	return *o.DnsLookupDomain.Get()
}

// GetDnsLookupDomainOk returns a tuple with the DnsLookupDomain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessNetworksGetItem) GetDnsLookupDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsLookupDomain.Get(), o.DnsLookupDomain.IsSet()
}

// SetDnsLookupDomain sets field value
func (o *WirelessNetworksGetItem) SetDnsLookupDomain(v string) {
	o.DnsLookupDomain.Set(&v)
}

// GetDisableEdns returns the DisableEdns field value
func (o *WirelessNetworksGetItem) GetDisableEdns() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableEdns
}

// GetDisableEdnsOk returns a tuple with the DisableEdns field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetDisableEdnsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableEdns, true
}

// SetDisableEdns sets field value
func (o *WirelessNetworksGetItem) SetDisableEdns(v bool) {
	o.DisableEdns = v
}

// GetUseDns64 returns the UseDns64 field value
func (o *WirelessNetworksGetItem) GetUseDns64() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseDns64
}

// GetUseDns64Ok returns a tuple with the UseDns64 field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetUseDns64Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseDns64, true
}

// SetUseDns64 sets field value
func (o *WirelessNetworksGetItem) SetUseDns64(v bool) {
	o.UseDns64 = v
}

// GetExternalConnectivity returns the ExternalConnectivity field value
func (o *WirelessNetworksGetItem) GetExternalConnectivity() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExternalConnectivity
}

// GetExternalConnectivityOk returns a tuple with the ExternalConnectivity field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetExternalConnectivityOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalConnectivity, true
}

// SetExternalConnectivity sets field value
func (o *WirelessNetworksGetItem) SetExternalConnectivity(v bool) {
	o.ExternalConnectivity = v
}

// GetType returns the Type field value
func (o *WirelessNetworksGetItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WirelessNetworksGetItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WirelessNetworksGetItem) SetType(v string) {
	o.Type = v
}

func (o WirelessNetworksGetItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WirelessNetworksGetItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["ssid"] = o.Ssid
	toSerialize["security"] = o.Security.Get()
	toSerialize["ipVersion"] = o.IpVersion
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["hidden"] = o.Hidden
	toSerialize["bandLocking"] = o.BandLocking
	toSerialize["dnsLookupDomain"] = o.DnsLookupDomain.Get()
	toSerialize["disableEdns"] = o.DisableEdns
	toSerialize["useDns64"] = o.UseDns64
	toSerialize["externalConnectivity"] = o.ExternalConnectivity
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *WirelessNetworksGetItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"ssid",
		"security",
		"ipVersion",
		"createdAt",
		"updatedAt",
		"hidden",
		"bandLocking",
		"dnsLookupDomain",
		"disableEdns",
		"useDns64",
		"externalConnectivity",
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

	varWirelessNetworksGetItem := _WirelessNetworksGetItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWirelessNetworksGetItem)

	if err != nil {
		return err
	}

	*o = WirelessNetworksGetItem(varWirelessNetworksGetItem)

	return err
}

type NullableWirelessNetworksGetItem struct {
	value *WirelessNetworksGetItem
	isSet bool
}

func (v NullableWirelessNetworksGetItem) Get() *WirelessNetworksGetItem {
	return v.value
}

func (v *NullableWirelessNetworksGetItem) Set(val *WirelessNetworksGetItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessNetworksGetItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessNetworksGetItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessNetworksGetItem(
	val *WirelessNetworksGetItem,
) *NullableWirelessNetworksGetItem {
	return &NullableWirelessNetworksGetItem{value: val, isSet: true}
}

func (v NullableWirelessNetworksGetItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessNetworksGetItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
