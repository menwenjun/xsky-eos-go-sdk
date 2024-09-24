# OspMetadataClusterStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**MovingDataSizeByte** | Pointer to **int64** |  | [optional] 
**ReadLatencyUs** | Pointer to **float64** |  | [optional] 
**ReadOps** | Pointer to **float64** |  | [optional] 
**TotalByte** | Pointer to **int64** |  | [optional] 
**TotalKvSizeByte** | Pointer to **int64** |  | [optional] 
**UsedByte** | Pointer to **int64** |  | [optional] 
**WriteLatencyUs** | Pointer to **float64** |  | [optional] 
**WriteOps** | Pointer to **float64** |  | [optional] 

## Methods

### NewOspMetadataClusterStat

`func NewOspMetadataClusterStat() *OspMetadataClusterStat`

NewOspMetadataClusterStat instantiates a new OspMetadataClusterStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspMetadataClusterStatWithDefaults

`func NewOspMetadataClusterStatWithDefaults() *OspMetadataClusterStat`

NewOspMetadataClusterStatWithDefaults instantiates a new OspMetadataClusterStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *OspMetadataClusterStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OspMetadataClusterStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OspMetadataClusterStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OspMetadataClusterStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetMovingDataSizeByte

`func (o *OspMetadataClusterStat) GetMovingDataSizeByte() int64`

GetMovingDataSizeByte returns the MovingDataSizeByte field if non-nil, zero value otherwise.

### GetMovingDataSizeByteOk

`func (o *OspMetadataClusterStat) GetMovingDataSizeByteOk() (*int64, bool)`

GetMovingDataSizeByteOk returns a tuple with the MovingDataSizeByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovingDataSizeByte

`func (o *OspMetadataClusterStat) SetMovingDataSizeByte(v int64)`

SetMovingDataSizeByte sets MovingDataSizeByte field to given value.

### HasMovingDataSizeByte

`func (o *OspMetadataClusterStat) HasMovingDataSizeByte() bool`

HasMovingDataSizeByte returns a boolean if a field has been set.

### GetReadLatencyUs

`func (o *OspMetadataClusterStat) GetReadLatencyUs() float64`

GetReadLatencyUs returns the ReadLatencyUs field if non-nil, zero value otherwise.

### GetReadLatencyUsOk

`func (o *OspMetadataClusterStat) GetReadLatencyUsOk() (*float64, bool)`

GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLatencyUs

`func (o *OspMetadataClusterStat) SetReadLatencyUs(v float64)`

SetReadLatencyUs sets ReadLatencyUs field to given value.

### HasReadLatencyUs

`func (o *OspMetadataClusterStat) HasReadLatencyUs() bool`

HasReadLatencyUs returns a boolean if a field has been set.

### GetReadOps

`func (o *OspMetadataClusterStat) GetReadOps() float64`

GetReadOps returns the ReadOps field if non-nil, zero value otherwise.

### GetReadOpsOk

`func (o *OspMetadataClusterStat) GetReadOpsOk() (*float64, bool)`

GetReadOpsOk returns a tuple with the ReadOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOps

`func (o *OspMetadataClusterStat) SetReadOps(v float64)`

SetReadOps sets ReadOps field to given value.

### HasReadOps

`func (o *OspMetadataClusterStat) HasReadOps() bool`

HasReadOps returns a boolean if a field has been set.

### GetTotalByte

`func (o *OspMetadataClusterStat) GetTotalByte() int64`

GetTotalByte returns the TotalByte field if non-nil, zero value otherwise.

### GetTotalByteOk

`func (o *OspMetadataClusterStat) GetTotalByteOk() (*int64, bool)`

GetTotalByteOk returns a tuple with the TotalByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalByte

`func (o *OspMetadataClusterStat) SetTotalByte(v int64)`

SetTotalByte sets TotalByte field to given value.

### HasTotalByte

`func (o *OspMetadataClusterStat) HasTotalByte() bool`

HasTotalByte returns a boolean if a field has been set.

### GetTotalKvSizeByte

`func (o *OspMetadataClusterStat) GetTotalKvSizeByte() int64`

GetTotalKvSizeByte returns the TotalKvSizeByte field if non-nil, zero value otherwise.

### GetTotalKvSizeByteOk

`func (o *OspMetadataClusterStat) GetTotalKvSizeByteOk() (*int64, bool)`

GetTotalKvSizeByteOk returns a tuple with the TotalKvSizeByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKvSizeByte

`func (o *OspMetadataClusterStat) SetTotalKvSizeByte(v int64)`

SetTotalKvSizeByte sets TotalKvSizeByte field to given value.

### HasTotalKvSizeByte

`func (o *OspMetadataClusterStat) HasTotalKvSizeByte() bool`

HasTotalKvSizeByte returns a boolean if a field has been set.

### GetUsedByte

`func (o *OspMetadataClusterStat) GetUsedByte() int64`

GetUsedByte returns the UsedByte field if non-nil, zero value otherwise.

### GetUsedByteOk

`func (o *OspMetadataClusterStat) GetUsedByteOk() (*int64, bool)`

GetUsedByteOk returns a tuple with the UsedByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedByte

`func (o *OspMetadataClusterStat) SetUsedByte(v int64)`

SetUsedByte sets UsedByte field to given value.

### HasUsedByte

`func (o *OspMetadataClusterStat) HasUsedByte() bool`

HasUsedByte returns a boolean if a field has been set.

### GetWriteLatencyUs

`func (o *OspMetadataClusterStat) GetWriteLatencyUs() float64`

GetWriteLatencyUs returns the WriteLatencyUs field if non-nil, zero value otherwise.

### GetWriteLatencyUsOk

`func (o *OspMetadataClusterStat) GetWriteLatencyUsOk() (*float64, bool)`

GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLatencyUs

`func (o *OspMetadataClusterStat) SetWriteLatencyUs(v float64)`

SetWriteLatencyUs sets WriteLatencyUs field to given value.

### HasWriteLatencyUs

`func (o *OspMetadataClusterStat) HasWriteLatencyUs() bool`

HasWriteLatencyUs returns a boolean if a field has been set.

### GetWriteOps

`func (o *OspMetadataClusterStat) GetWriteOps() float64`

GetWriteOps returns the WriteOps field if non-nil, zero value otherwise.

### GetWriteOpsOk

`func (o *OspMetadataClusterStat) GetWriteOpsOk() (*float64, bool)`

GetWriteOpsOk returns a tuple with the WriteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteOps

`func (o *OspMetadataClusterStat) SetWriteOps(v float64)`

SetWriteOps sets WriteOps field to given value.

### HasWriteOps

`func (o *OspMetadataClusterStat) HasWriteOps() bool`

HasWriteOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


