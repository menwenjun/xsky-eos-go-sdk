# DfsQosPolicyUpdateReqQosPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description if dfs qos policy | [optional] 
**Name** | **string** | name of dfs qos policy | 
**ReadBandwidth** | Pointer to **int64** | read bandwidth limitation | [optional] 
**ReadObject** | Pointer to **int64** | read objects limitation | [optional] 
**TotalBandwidth** | Pointer to **int64** | total bandwidth limitation | [optional] 
**TotalObject** | Pointer to **int64** | total objects limitation | [optional] 
**WriteBandwidth** | Pointer to **int64** | write bandwidth limitation | [optional] 
**WriteObject** | Pointer to **int64** | write objects limitation | [optional] 

## Methods

### NewDfsQosPolicyUpdateReqQosPolicy

`func NewDfsQosPolicyUpdateReqQosPolicy(name string, ) *DfsQosPolicyUpdateReqQosPolicy`

NewDfsQosPolicyUpdateReqQosPolicy instantiates a new DfsQosPolicyUpdateReqQosPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQosPolicyUpdateReqQosPolicyWithDefaults

`func NewDfsQosPolicyUpdateReqQosPolicyWithDefaults() *DfsQosPolicyUpdateReqQosPolicy`

NewDfsQosPolicyUpdateReqQosPolicyWithDefaults instantiates a new DfsQosPolicyUpdateReqQosPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsQosPolicyUpdateReqQosPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsQosPolicyUpdateReqQosPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsQosPolicyUpdateReqQosPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetReadBandwidth

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetReadBandwidth() int64`

GetReadBandwidth returns the ReadBandwidth field if non-nil, zero value otherwise.

### GetReadBandwidthOk

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetReadBandwidthOk() (*int64, bool)`

GetReadBandwidthOk returns a tuple with the ReadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidth

`func (o *DfsQosPolicyUpdateReqQosPolicy) SetReadBandwidth(v int64)`

SetReadBandwidth sets ReadBandwidth field to given value.

### HasReadBandwidth

`func (o *DfsQosPolicyUpdateReqQosPolicy) HasReadBandwidth() bool`

HasReadBandwidth returns a boolean if a field has been set.

### GetReadObject

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetReadObject() int64`

GetReadObject returns the ReadObject field if non-nil, zero value otherwise.

### GetReadObjectOk

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetReadObjectOk() (*int64, bool)`

GetReadObjectOk returns a tuple with the ReadObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadObject

`func (o *DfsQosPolicyUpdateReqQosPolicy) SetReadObject(v int64)`

SetReadObject sets ReadObject field to given value.

### HasReadObject

`func (o *DfsQosPolicyUpdateReqQosPolicy) HasReadObject() bool`

HasReadObject returns a boolean if a field has been set.

### GetTotalBandwidth

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetTotalBandwidth() int64`

GetTotalBandwidth returns the TotalBandwidth field if non-nil, zero value otherwise.

### GetTotalBandwidthOk

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetTotalBandwidthOk() (*int64, bool)`

GetTotalBandwidthOk returns a tuple with the TotalBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBandwidth

`func (o *DfsQosPolicyUpdateReqQosPolicy) SetTotalBandwidth(v int64)`

SetTotalBandwidth sets TotalBandwidth field to given value.

### HasTotalBandwidth

`func (o *DfsQosPolicyUpdateReqQosPolicy) HasTotalBandwidth() bool`

HasTotalBandwidth returns a boolean if a field has been set.

### GetTotalObject

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetTotalObject() int64`

GetTotalObject returns the TotalObject field if non-nil, zero value otherwise.

### GetTotalObjectOk

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetTotalObjectOk() (*int64, bool)`

GetTotalObjectOk returns a tuple with the TotalObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObject

`func (o *DfsQosPolicyUpdateReqQosPolicy) SetTotalObject(v int64)`

SetTotalObject sets TotalObject field to given value.

### HasTotalObject

`func (o *DfsQosPolicyUpdateReqQosPolicy) HasTotalObject() bool`

HasTotalObject returns a boolean if a field has been set.

### GetWriteBandwidth

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetWriteBandwidth() int64`

GetWriteBandwidth returns the WriteBandwidth field if non-nil, zero value otherwise.

### GetWriteBandwidthOk

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetWriteBandwidthOk() (*int64, bool)`

GetWriteBandwidthOk returns a tuple with the WriteBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidth

`func (o *DfsQosPolicyUpdateReqQosPolicy) SetWriteBandwidth(v int64)`

SetWriteBandwidth sets WriteBandwidth field to given value.

### HasWriteBandwidth

`func (o *DfsQosPolicyUpdateReqQosPolicy) HasWriteBandwidth() bool`

HasWriteBandwidth returns a boolean if a field has been set.

### GetWriteObject

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetWriteObject() int64`

GetWriteObject returns the WriteObject field if non-nil, zero value otherwise.

### GetWriteObjectOk

`func (o *DfsQosPolicyUpdateReqQosPolicy) GetWriteObjectOk() (*int64, bool)`

GetWriteObjectOk returns a tuple with the WriteObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteObject

`func (o *DfsQosPolicyUpdateReqQosPolicy) SetWriteObject(v int64)`

SetWriteObject sets WriteObject field to given value.

### HasWriteObject

`func (o *DfsQosPolicyUpdateReqQosPolicy) HasWriteObject() bool`

HasWriteObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


