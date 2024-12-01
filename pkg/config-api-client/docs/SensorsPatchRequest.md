# SensorsPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AddressNote** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**PcapMode** | Pointer to [**PcapMode**](PcapMode.md) |  | [optional] 

## Methods

### NewSensorsPatchRequest

`func NewSensorsPatchRequest() *SensorsPatchRequest`

NewSensorsPatchRequest instantiates a new SensorsPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorsPatchRequestWithDefaults

`func NewSensorsPatchRequestWithDefaults() *SensorsPatchRequest`

NewSensorsPatchRequestWithDefaults instantiates a new SensorsPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SensorsPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SensorsPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SensorsPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SensorsPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddressNote

`func (o *SensorsPatchRequest) GetAddressNote() string`

GetAddressNote returns the AddressNote field if non-nil, zero value otherwise.

### GetAddressNoteOk

`func (o *SensorsPatchRequest) GetAddressNoteOk() (*string, bool)`

GetAddressNoteOk returns a tuple with the AddressNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressNote

`func (o *SensorsPatchRequest) SetAddressNote(v string)`

SetAddressNote sets AddressNote field to given value.

### HasAddressNote

`func (o *SensorsPatchRequest) HasAddressNote() bool`

HasAddressNote returns a boolean if a field has been set.

### GetNotes

`func (o *SensorsPatchRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SensorsPatchRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SensorsPatchRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SensorsPatchRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPcapMode

`func (o *SensorsPatchRequest) GetPcapMode() PcapMode`

GetPcapMode returns the PcapMode field if non-nil, zero value otherwise.

### GetPcapModeOk

`func (o *SensorsPatchRequest) GetPcapModeOk() (*PcapMode, bool)`

GetPcapModeOk returns a tuple with the PcapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapMode

`func (o *SensorsPatchRequest) SetPcapMode(v PcapMode)`

SetPcapMode sets PcapMode field to given value.

### HasPcapMode

`func (o *SensorsPatchRequest) HasPcapMode() bool`

HasPcapMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


