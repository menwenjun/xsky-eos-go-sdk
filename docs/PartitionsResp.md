# PartitionsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partitions** | Pointer to [**[]PartitionRecord**](PartitionRecord.md) |  | [optional] 

## Methods

### NewPartitionsResp

`func NewPartitionsResp() *PartitionsResp`

NewPartitionsResp instantiates a new PartitionsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionsRespWithDefaults

`func NewPartitionsRespWithDefaults() *PartitionsResp`

NewPartitionsRespWithDefaults instantiates a new PartitionsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitions

`func (o *PartitionsResp) GetPartitions() []PartitionRecord`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *PartitionsResp) GetPartitionsOk() (*[]PartitionRecord, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *PartitionsResp) SetPartitions(v []PartitionRecord)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *PartitionsResp) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


