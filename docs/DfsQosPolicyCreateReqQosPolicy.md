# DfsQosPolicyCreateReqQosPolicy

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

### NewDfsQosPolicyCreateReqQosPolicy

`func NewDfsQosPolicyCreateReqQosPolicy(name string, ) *DfsQosPolicyCreateReqQosPolicy`

NewDfsQosPolicyCreateReqQosPolicy instantiates a new DfsQosPolicyCreateReqQosPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQosPolicyCreateReqQosPolicyWithDefaults

`func NewDfsQosPolicyCreateReqQosPolicyWithDefaults() *DfsQosPolicyCreateReqQosPolicy`

NewDfsQosPolicyCreateReqQosPolicyWithDefaults instantiates a new DfsQosPolicyCreateReqQosPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DfsQosPolicyCreateReqQosPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsQosPolicyCreateReqQosPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsQosPolicyCreateReqQosPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsQosPolicyCreateReqQosPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DfsQosPolicyCreateReqQosPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsQosPolicyCreateReqQosPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsQosPolicyCreateReqQosPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetReadBandwidth

`func (o *DfsQosPolicyCreateReqQosPolicy) GetReadBandwidth() int64`

GetReadBandwidth returns the ReadBandwidth field if non-nil, zero value otherwise.

### GetReadBandwidthOk

`func (o *DfsQosPolicyCreateReqQosPolicy) GetReadBandwidthOk() (*int64, bool)`

GetReadBandwidthOk returns a tuple with the ReadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidth

`func (o *DfsQosPolicyCreateReqQosPolicy) SetReadBandwidth(v int64)`

SetReadBandwidth sets ReadBandwidth field to given value.

### HasReadBandwidth

`func (o *DfsQosPolicyCreateReqQosPolicy) HasReadBandwidth() bool`

HasReadBandwidth returns a boolean if a field has been set.

### GetReadObject

`func (o *DfsQosPolicyCreateReqQosPolicy) GetReadObject() int64`

GetReadObject returns the ReadObject field if non-nil, zero value otherwise.

### GetReadObjectOk

`func (o *DfsQosPolicyCreateReqQosPolicy) GetReadObjectOk() (*int64, bool)`

GetReadObjectOk returns a tuple with the ReadObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadObject

`func (o *DfsQosPolicyCreateReqQosPolicy) SetReadObject(v int64)`

SetReadObject sets ReadObject field to given value.

### HasReadObject

`func (o *DfsQosPolicyCreateReqQosPolicy) HasReadObject() bool`

HasReadObject returns a boolean if a field has been set.

### GetTotalBandwidth

`func (o *DfsQosPolicyCreateReqQosPolicy) GetTotalBandwidth() int64`

GetTotalBandwidth returns the TotalBandwidth field if non-nil, zero value otherwise.

### GetTotalBandwidthOk

`func (o *DfsQosPolicyCreateReqQosPolicy) GetTotalBandwidthOk() (*int64, bool)`

GetTotalBandwidthOk returns a tuple with the TotalBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBandwidth

`func (o *DfsQosPolicyCreateReqQosPolicy) SetTotalBandwidth(v int64)`

SetTotalBandwidth sets TotalBandwidth field to given value.

### HasTotalBandwidth

`func (o *DfsQosPolicyCreateReqQosPolicy) HasTotalBandwidth() bool`

HasTotalBandwidth returns a boolean if a field has been set.

### GetTotalObject

`func (o *DfsQosPolicyCreateReqQosPolicy) GetTotalObject() int64`

GetTotalObject returns the TotalObject field if non-nil, zero value otherwise.

### GetTotalObjectOk

`func (o *DfsQosPolicyCreateReqQosPolicy) GetTotalObjectOk() (*int64, bool)`

GetTotalObjectOk returns a tuple with the TotalObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObject

`func (o *DfsQosPolicyCreateReqQosPolicy) SetTotalObject(v int64)`

SetTotalObject sets TotalObject field to given value.

### HasTotalObject

`func (o *DfsQosPolicyCreateReqQosPolicy) HasTotalObject() bool`

HasTotalObject returns a boolean if a field has been set.

### GetWriteBandwidth

`func (o *DfsQosPolicyCreateReqQosPolicy) GetWriteBandwidth() int64`

GetWriteBandwidth returns the WriteBandwidth field if non-nil, zero value otherwise.

### GetWriteBandwidthOk

`func (o *DfsQosPolicyCreateReqQosPolicy) GetWriteBandwidthOk() (*int64, bool)`

GetWriteBandwidthOk returns a tuple with the WriteBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidth

`func (o *DfsQosPolicyCreateReqQosPolicy) SetWriteBandwidth(v int64)`

SetWriteBandwidth sets WriteBandwidth field to given value.

### HasWriteBandwidth

`func (o *DfsQosPolicyCreateReqQosPolicy) HasWriteBandwidth() bool`

HasWriteBandwidth returns a boolean if a field has been set.

### GetWriteObject

`func (o *DfsQosPolicyCreateReqQosPolicy) GetWriteObject() int64`

GetWriteObject returns the WriteObject field if non-nil, zero value otherwise.

### GetWriteObjectOk

`func (o *DfsQosPolicyCreateReqQosPolicy) GetWriteObjectOk() (*int64, bool)`

GetWriteObjectOk returns a tuple with the WriteObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteObject

`func (o *DfsQosPolicyCreateReqQosPolicy) SetWriteObject(v int64)`

SetWriteObject sets WriteObject field to given value.

### HasWriteObject

`func (o *DfsQosPolicyCreateReqQosPolicy) HasWriteObject() bool`

HasWriteObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


