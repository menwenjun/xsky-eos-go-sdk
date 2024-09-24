# DfsQosStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**ReadBandwidth** | Pointer to **int64** |  | [optional] 
**ReadObject** | Pointer to **int64** |  | [optional] 
**TotalBandwidth** | Pointer to **int64** |  | [optional] 
**TotalObject** | Pointer to **int64** |  | [optional] 
**WriteBandwidth** | Pointer to **int64** |  | [optional] 
**WriteObject** | Pointer to **int64** |  | [optional] 

## Methods

### NewDfsQosStat

`func NewDfsQosStat() *DfsQosStat`

NewDfsQosStat instantiates a new DfsQosStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQosStatWithDefaults

`func NewDfsQosStatWithDefaults() *DfsQosStat`

NewDfsQosStatWithDefaults instantiates a new DfsQosStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DfsQosStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsQosStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsQosStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsQosStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetReadBandwidth

`func (o *DfsQosStat) GetReadBandwidth() int64`

GetReadBandwidth returns the ReadBandwidth field if non-nil, zero value otherwise.

### GetReadBandwidthOk

`func (o *DfsQosStat) GetReadBandwidthOk() (*int64, bool)`

GetReadBandwidthOk returns a tuple with the ReadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidth

`func (o *DfsQosStat) SetReadBandwidth(v int64)`

SetReadBandwidth sets ReadBandwidth field to given value.

### HasReadBandwidth

`func (o *DfsQosStat) HasReadBandwidth() bool`

HasReadBandwidth returns a boolean if a field has been set.

### GetReadObject

`func (o *DfsQosStat) GetReadObject() int64`

GetReadObject returns the ReadObject field if non-nil, zero value otherwise.

### GetReadObjectOk

`func (o *DfsQosStat) GetReadObjectOk() (*int64, bool)`

GetReadObjectOk returns a tuple with the ReadObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadObject

`func (o *DfsQosStat) SetReadObject(v int64)`

SetReadObject sets ReadObject field to given value.

### HasReadObject

`func (o *DfsQosStat) HasReadObject() bool`

HasReadObject returns a boolean if a field has been set.

### GetTotalBandwidth

`func (o *DfsQosStat) GetTotalBandwidth() int64`

GetTotalBandwidth returns the TotalBandwidth field if non-nil, zero value otherwise.

### GetTotalBandwidthOk

`func (o *DfsQosStat) GetTotalBandwidthOk() (*int64, bool)`

GetTotalBandwidthOk returns a tuple with the TotalBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBandwidth

`func (o *DfsQosStat) SetTotalBandwidth(v int64)`

SetTotalBandwidth sets TotalBandwidth field to given value.

### HasTotalBandwidth

`func (o *DfsQosStat) HasTotalBandwidth() bool`

HasTotalBandwidth returns a boolean if a field has been set.

### GetTotalObject

`func (o *DfsQosStat) GetTotalObject() int64`

GetTotalObject returns the TotalObject field if non-nil, zero value otherwise.

### GetTotalObjectOk

`func (o *DfsQosStat) GetTotalObjectOk() (*int64, bool)`

GetTotalObjectOk returns a tuple with the TotalObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObject

`func (o *DfsQosStat) SetTotalObject(v int64)`

SetTotalObject sets TotalObject field to given value.

### HasTotalObject

`func (o *DfsQosStat) HasTotalObject() bool`

HasTotalObject returns a boolean if a field has been set.

### GetWriteBandwidth

`func (o *DfsQosStat) GetWriteBandwidth() int64`

GetWriteBandwidth returns the WriteBandwidth field if non-nil, zero value otherwise.

### GetWriteBandwidthOk

`func (o *DfsQosStat) GetWriteBandwidthOk() (*int64, bool)`

GetWriteBandwidthOk returns a tuple with the WriteBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidth

`func (o *DfsQosStat) SetWriteBandwidth(v int64)`

SetWriteBandwidth sets WriteBandwidth field to given value.

### HasWriteBandwidth

`func (o *DfsQosStat) HasWriteBandwidth() bool`

HasWriteBandwidth returns a boolean if a field has been set.

### GetWriteObject

`func (o *DfsQosStat) GetWriteObject() int64`

GetWriteObject returns the WriteObject field if non-nil, zero value otherwise.

### GetWriteObjectOk

`func (o *DfsQosStat) GetWriteObjectOk() (*int64, bool)`

GetWriteObjectOk returns a tuple with the WriteObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteObject

`func (o *DfsQosStat) SetWriteObject(v int64)`

SetWriteObject sets WriteObject field to given value.

### HasWriteObject

`func (o *DfsQosStat) HasWriteObject() bool`

HasWriteObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


