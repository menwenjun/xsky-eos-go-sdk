# DfsQosPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DirNum** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReadBandwidth** | Pointer to **int64** |  | [optional] 
**ReadObject** | Pointer to **int64** |  | [optional] 
**TotalBandwidth** | Pointer to **int64** |  | [optional] 
**TotalObject** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**WriteBandwidth** | Pointer to **int64** |  | [optional] 
**WriteObject** | Pointer to **int64** |  | [optional] 

## Methods

### NewDfsQosPolicy

`func NewDfsQosPolicy() *DfsQosPolicy`

NewDfsQosPolicy instantiates a new DfsQosPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQosPolicyWithDefaults

`func NewDfsQosPolicyWithDefaults() *DfsQosPolicy`

NewDfsQosPolicyWithDefaults instantiates a new DfsQosPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsQosPolicy) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsQosPolicy) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsQosPolicy) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsQosPolicy) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DfsQosPolicy) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsQosPolicy) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsQosPolicy) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsQosPolicy) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsQosPolicy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsQosPolicy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsQosPolicy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsQosPolicy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DfsQosPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsQosPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsQosPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsQosPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirNum

`func (o *DfsQosPolicy) GetDirNum() int64`

GetDirNum returns the DirNum field if non-nil, zero value otherwise.

### GetDirNumOk

`func (o *DfsQosPolicy) GetDirNumOk() (*int64, bool)`

GetDirNumOk returns a tuple with the DirNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirNum

`func (o *DfsQosPolicy) SetDirNum(v int64)`

SetDirNum sets DirNum field to given value.

### HasDirNum

`func (o *DfsQosPolicy) HasDirNum() bool`

HasDirNum returns a boolean if a field has been set.

### GetId

`func (o *DfsQosPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsQosPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsQosPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsQosPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DfsQosPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsQosPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsQosPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsQosPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadBandwidth

`func (o *DfsQosPolicy) GetReadBandwidth() int64`

GetReadBandwidth returns the ReadBandwidth field if non-nil, zero value otherwise.

### GetReadBandwidthOk

`func (o *DfsQosPolicy) GetReadBandwidthOk() (*int64, bool)`

GetReadBandwidthOk returns a tuple with the ReadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidth

`func (o *DfsQosPolicy) SetReadBandwidth(v int64)`

SetReadBandwidth sets ReadBandwidth field to given value.

### HasReadBandwidth

`func (o *DfsQosPolicy) HasReadBandwidth() bool`

HasReadBandwidth returns a boolean if a field has been set.

### GetReadObject

`func (o *DfsQosPolicy) GetReadObject() int64`

GetReadObject returns the ReadObject field if non-nil, zero value otherwise.

### GetReadObjectOk

`func (o *DfsQosPolicy) GetReadObjectOk() (*int64, bool)`

GetReadObjectOk returns a tuple with the ReadObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadObject

`func (o *DfsQosPolicy) SetReadObject(v int64)`

SetReadObject sets ReadObject field to given value.

### HasReadObject

`func (o *DfsQosPolicy) HasReadObject() bool`

HasReadObject returns a boolean if a field has been set.

### GetTotalBandwidth

`func (o *DfsQosPolicy) GetTotalBandwidth() int64`

GetTotalBandwidth returns the TotalBandwidth field if non-nil, zero value otherwise.

### GetTotalBandwidthOk

`func (o *DfsQosPolicy) GetTotalBandwidthOk() (*int64, bool)`

GetTotalBandwidthOk returns a tuple with the TotalBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBandwidth

`func (o *DfsQosPolicy) SetTotalBandwidth(v int64)`

SetTotalBandwidth sets TotalBandwidth field to given value.

### HasTotalBandwidth

`func (o *DfsQosPolicy) HasTotalBandwidth() bool`

HasTotalBandwidth returns a boolean if a field has been set.

### GetTotalObject

`func (o *DfsQosPolicy) GetTotalObject() int64`

GetTotalObject returns the TotalObject field if non-nil, zero value otherwise.

### GetTotalObjectOk

`func (o *DfsQosPolicy) GetTotalObjectOk() (*int64, bool)`

GetTotalObjectOk returns a tuple with the TotalObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObject

`func (o *DfsQosPolicy) SetTotalObject(v int64)`

SetTotalObject sets TotalObject field to given value.

### HasTotalObject

`func (o *DfsQosPolicy) HasTotalObject() bool`

HasTotalObject returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsQosPolicy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsQosPolicy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsQosPolicy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsQosPolicy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetWriteBandwidth

`func (o *DfsQosPolicy) GetWriteBandwidth() int64`

GetWriteBandwidth returns the WriteBandwidth field if non-nil, zero value otherwise.

### GetWriteBandwidthOk

`func (o *DfsQosPolicy) GetWriteBandwidthOk() (*int64, bool)`

GetWriteBandwidthOk returns a tuple with the WriteBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidth

`func (o *DfsQosPolicy) SetWriteBandwidth(v int64)`

SetWriteBandwidth sets WriteBandwidth field to given value.

### HasWriteBandwidth

`func (o *DfsQosPolicy) HasWriteBandwidth() bool`

HasWriteBandwidth returns a boolean if a field has been set.

### GetWriteObject

`func (o *DfsQosPolicy) GetWriteObject() int64`

GetWriteObject returns the WriteObject field if non-nil, zero value otherwise.

### GetWriteObjectOk

`func (o *DfsQosPolicy) GetWriteObjectOk() (*int64, bool)`

GetWriteObjectOk returns a tuple with the WriteObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteObject

`func (o *DfsQosPolicy) SetWriteObject(v int64)`

SetWriteObject sets WriteObject field to given value.

### HasWriteObject

`func (o *DfsQosPolicy) HasWriteObject() bool`

HasWriteObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


