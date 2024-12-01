# SensorGroupAssignmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Group** | [**Group**](Group.md) |  | 
**Sensor** | [**Sensor**](Sensor.md) |  | 
**Type** | **string** |  | 

## Methods

### NewSensorGroupAssignmentResponse

`func NewSensorGroupAssignmentResponse(id string, group Group, sensor Sensor, type_ string, ) *SensorGroupAssignmentResponse`

NewSensorGroupAssignmentResponse instantiates a new SensorGroupAssignmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorGroupAssignmentResponseWithDefaults

`func NewSensorGroupAssignmentResponseWithDefaults() *SensorGroupAssignmentResponse`

NewSensorGroupAssignmentResponseWithDefaults instantiates a new SensorGroupAssignmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SensorGroupAssignmentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SensorGroupAssignmentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SensorGroupAssignmentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *SensorGroupAssignmentResponse) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SensorGroupAssignmentResponse) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SensorGroupAssignmentResponse) SetGroup(v Group)`

SetGroup sets Group field to given value.


### GetSensor

`func (o *SensorGroupAssignmentResponse) GetSensor() Sensor`

GetSensor returns the Sensor field if non-nil, zero value otherwise.

### GetSensorOk

`func (o *SensorGroupAssignmentResponse) GetSensorOk() (*Sensor, bool)`

GetSensorOk returns a tuple with the Sensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensor

`func (o *SensorGroupAssignmentResponse) SetSensor(v Sensor)`

SetSensor sets Sensor field to given value.


### GetType

`func (o *SensorGroupAssignmentResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SensorGroupAssignmentResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SensorGroupAssignmentResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


