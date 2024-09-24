# FCPortsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FcPorts** | [**[]FCPortRecord**](FCPortRecord.md) | fc port | 

## Methods

### NewFCPortsResp

`func NewFCPortsResp(fcPorts []FCPortRecord, ) *FCPortsResp`

NewFCPortsResp instantiates a new FCPortsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFCPortsRespWithDefaults

`func NewFCPortsRespWithDefaults() *FCPortsResp`

NewFCPortsRespWithDefaults instantiates a new FCPortsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFcPorts

`func (o *FCPortsResp) GetFcPorts() []FCPortRecord`

GetFcPorts returns the FcPorts field if non-nil, zero value otherwise.

### GetFcPortsOk

`func (o *FCPortsResp) GetFcPortsOk() (*[]FCPortRecord, bool)`

GetFcPortsOk returns a tuple with the FcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPorts

`func (o *FCPortsResp) SetFcPorts(v []FCPortRecord)`

SetFcPorts sets FcPorts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


