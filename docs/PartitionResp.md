# PartitionResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | Pointer to [**PartitionRecord**](PartitionRecord.md) |  | [optional] 

## Methods

### NewPartitionResp

`func NewPartitionResp() *PartitionResp`

NewPartitionResp instantiates a new PartitionResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionRespWithDefaults

`func NewPartitionRespWithDefaults() *PartitionResp`

NewPartitionRespWithDefaults instantiates a new PartitionResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *PartitionResp) GetPartition() PartitionRecord`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *PartitionResp) GetPartitionOk() (*PartitionRecord, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *PartitionResp) SetPartition(v PartitionRecord)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *PartitionResp) HasPartition() bool`

HasPartition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


