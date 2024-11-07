/*
Configuration Api

Nice description goes here

API version: 5.4.0
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

// checks if the WiredNetworksItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WiredNetworksItem{}

// WiredNetworksItem struct for WiredNetworksItem
type WiredNetworksItem struct {
	Id                   string         `json:"id"`
	Name                 string         `json:"name"`
	IpVersion            string         `json:"ipVersion"`
	CreatedAt            time.Time      `json:"createdAt"`
	UpdatedAt            time.Time      `json:"updatedAt"`
	Security             NullableString `json:"security"`
	DnsLookupDomain      NullableString `json:"dnsLookupDomain"`
	DisableEdns          bool           `json:"disableEdns"`
	UseDns64             bool           `json:"useDns64"`
	ExternalConnectivity bool           `json:"externalConnectivity"`
	VLanId               NullableInt32  `json:"vLanId"`
	Type                 string         `json:"type"`
}

type _WiredNetworksItem WiredNetworksItem

// NewWiredNetworksItem instantiates a new WiredNetworksItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWiredNetworksItem(
	id string,
	name string,
	ipVersion string,
	createdAt time.Time,
	updatedAt time.Time,
	security NullableString,
	dnsLookupDomain NullableString,
	disableEdns bool,
	useDns64 bool,
	externalConnectivity bool,
	vLanId NullableInt32,
	type_ string,
) *WiredNetworksItem {
	this := WiredNetworksItem{}
	this.Id = id
	this.Name = name
	this.IpVersion = ipVersion
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Security = security
	this.DnsLookupDomain = dnsLookupDomain
	this.DisableEdns = disableEdns
	this.UseDns64 = useDns64
	this.ExternalConnectivity = externalConnectivity
	this.VLanId = vLanId
	this.Type = type_
	return &this
}

// NewWiredNetworksItemWithDefaults instantiates a new WiredNetworksItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWiredNetworksItemWithDefaults() *WiredNetworksItem {
	this := WiredNetworksItem{}
	return &this
}

// GetId returns the Id field value
func (o *WiredNetworksItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WiredNetworksItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *WiredNetworksItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WiredNetworksItem) SetName(v string) {
	o.Name = v
}

// GetIpVersion returns the IpVersion field value
func (o *WiredNetworksItem) GetIpVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetIpVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpVersion, true
}

// SetIpVersion sets field value
func (o *WiredNetworksItem) SetIpVersion(v string) {
	o.IpVersion = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WiredNetworksItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WiredNetworksItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WiredNetworksItem) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WiredNetworksItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetSecurity returns the Security field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WiredNetworksItem) GetSecurity() string {
	if o == nil || o.Security.Get() == nil {
		var ret string
		return ret
	}

	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WiredNetworksItem) GetSecurityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// SetSecurity sets field value
func (o *WiredNetworksItem) SetSecurity(v string) {
	o.Security.Set(&v)
}

// GetDnsLookupDomain returns the DnsLookupDomain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WiredNetworksItem) GetDnsLookupDomain() string {
	if o == nil || o.DnsLookupDomain.Get() == nil {
		var ret string
		return ret
	}

	return *o.DnsLookupDomain.Get()
}

// GetDnsLookupDomainOk returns a tuple with the DnsLookupDomain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WiredNetworksItem) GetDnsLookupDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsLookupDomain.Get(), o.DnsLookupDomain.IsSet()
}

// SetDnsLookupDomain sets field value
func (o *WiredNetworksItem) SetDnsLookupDomain(v string) {
	o.DnsLookupDomain.Set(&v)
}

// GetDisableEdns returns the DisableEdns field value
func (o *WiredNetworksItem) GetDisableEdns() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableEdns
}

// GetDisableEdnsOk returns a tuple with the DisableEdns field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetDisableEdnsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableEdns, true
}

// SetDisableEdns sets field value
func (o *WiredNetworksItem) SetDisableEdns(v bool) {
	o.DisableEdns = v
}

// GetUseDns64 returns the UseDns64 field value
func (o *WiredNetworksItem) GetUseDns64() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseDns64
}

// GetUseDns64Ok returns a tuple with the UseDns64 field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetUseDns64Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseDns64, true
}

// SetUseDns64 sets field value
func (o *WiredNetworksItem) SetUseDns64(v bool) {
	o.UseDns64 = v
}

// GetExternalConnectivity returns the ExternalConnectivity field value
func (o *WiredNetworksItem) GetExternalConnectivity() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExternalConnectivity
}

// GetExternalConnectivityOk returns a tuple with the ExternalConnectivity field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetExternalConnectivityOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalConnectivity, true
}

// SetExternalConnectivity sets field value
func (o *WiredNetworksItem) SetExternalConnectivity(v bool) {
	o.ExternalConnectivity = v
}

// GetVLanId returns the VLanId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *WiredNetworksItem) GetVLanId() int32 {
	if o == nil || o.VLanId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.VLanId.Get()
}

// GetVLanIdOk returns a tuple with the VLanId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WiredNetworksItem) GetVLanIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VLanId.Get(), o.VLanId.IsSet()
}

// SetVLanId sets field value
func (o *WiredNetworksItem) SetVLanId(v int32) {
	o.VLanId.Set(&v)
}

// GetType returns the Type field value
func (o *WiredNetworksItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WiredNetworksItem) SetType(v string) {
	o.Type = v
}

func (o WiredNetworksItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WiredNetworksItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["ipVersion"] = o.IpVersion
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["security"] = o.Security.Get()
	toSerialize["dnsLookupDomain"] = o.DnsLookupDomain.Get()
	toSerialize["disableEdns"] = o.DisableEdns
	toSerialize["useDns64"] = o.UseDns64
	toSerialize["externalConnectivity"] = o.ExternalConnectivity
	toSerialize["vLanId"] = o.VLanId.Get()
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *WiredNetworksItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"ipVersion",
		"createdAt",
		"updatedAt",
		"security",
		"dnsLookupDomain",
		"disableEdns",
		"useDns64",
		"externalConnectivity",
		"vLanId",
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

	varWiredNetworksItem := _WiredNetworksItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWiredNetworksItem)

	if err != nil {
		return err
	}

	*o = WiredNetworksItem(varWiredNetworksItem)

	return err
}

type NullableWiredNetworksItem struct {
	value *WiredNetworksItem
	isSet bool
}

func (v NullableWiredNetworksItem) Get() *WiredNetworksItem {
	return v.value
}

func (v *NullableWiredNetworksItem) Set(val *WiredNetworksItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWiredNetworksItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWiredNetworksItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWiredNetworksItem(val *WiredNetworksItem) *NullableWiredNetworksItem {
	return &NullableWiredNetworksItem{value: val, isSet: true}
}

func (v NullableWiredNetworksItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWiredNetworksItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
