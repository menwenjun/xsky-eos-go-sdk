# OSSearchEngineStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageMetaQueryLatency** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**MetaQueryPm** | Pointer to **float64** |  | [optional] 
**SyncBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**SyncObjectsPm** | Pointer to **float64** |  | [optional] 
**TotalMetaNum** | Pointer to **int64** |  | [optional] 
**TotalMetaSizeBytes** | Pointer to **int64** |  | [optional] 

## Methods

### NewOSSearchEngineStat

`func NewOSSearchEngineStat() *OSSearchEngineStat`

NewOSSearchEngineStat instantiates a new OSSearchEngineStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSSearchEngineStatWithDefaults

`func NewOSSearchEngineStatWithDefaults() *OSSearchEngineStat`

NewOSSearchEngineStatWithDefaults instantiates a new OSSearchEngineStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageMetaQueryLatency

`func (o *OSSearchEngineStat) GetAverageMetaQueryLatency() float64`

GetAverageMetaQueryLatency returns the AverageMetaQueryLatency field if non-nil, zero value otherwise.

### GetAverageMetaQueryLatencyOk

`func (o *OSSearchEngineStat) GetAverageMetaQueryLatencyOk() (*float64, bool)`

GetAverageMetaQueryLatencyOk returns a tuple with the AverageMetaQueryLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageMetaQueryLatency

`func (o *OSSearchEngineStat) SetAverageMetaQueryLatency(v float64)`

SetAverageMetaQueryLatency sets AverageMetaQueryLatency field to given value.

### HasAverageMetaQueryLatency

`func (o *OSSearchEngineStat) HasAverageMetaQueryLatency() bool`

HasAverageMetaQueryLatency returns a boolean if a field has been set.

### GetCreate

`func (o *OSSearchEngineStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSSearchEngineStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSSearchEngineStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSSearchEngineStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetMetaQueryPm

`func (o *OSSearchEngineStat) GetMetaQueryPm() float64`

GetMetaQueryPm returns the MetaQueryPm field if non-nil, zero value otherwise.

### GetMetaQueryPmOk

`func (o *OSSearchEngineStat) GetMetaQueryPmOk() (*float64, bool)`

GetMetaQueryPmOk returns a tuple with the MetaQueryPm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaQueryPm

`func (o *OSSearchEngineStat) SetMetaQueryPm(v float64)`

SetMetaQueryPm sets MetaQueryPm field to given value.

### HasMetaQueryPm

`func (o *OSSearchEngineStat) HasMetaQueryPm() bool`

HasMetaQueryPm returns a boolean if a field has been set.

### GetSyncBandwidthKbyte

`func (o *OSSearchEngineStat) GetSyncBandwidthKbyte() float64`

GetSyncBandwidthKbyte returns the SyncBandwidthKbyte field if non-nil, zero value otherwise.

### GetSyncBandwidthKbyteOk

`func (o *OSSearchEngineStat) GetSyncBandwidthKbyteOk() (*float64, bool)`

GetSyncBandwidthKbyteOk returns a tuple with the SyncBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncBandwidthKbyte

`func (o *OSSearchEngineStat) SetSyncBandwidthKbyte(v float64)`

SetSyncBandwidthKbyte sets SyncBandwidthKbyte field to given value.

### HasSyncBandwidthKbyte

`func (o *OSSearchEngineStat) HasSyncBandwidthKbyte() bool`

HasSyncBandwidthKbyte returns a boolean if a field has been set.

### GetSyncObjectsPm

`func (o *OSSearchEngineStat) GetSyncObjectsPm() float64`

GetSyncObjectsPm returns the SyncObjectsPm field if non-nil, zero value otherwise.

### GetSyncObjectsPmOk

`func (o *OSSearchEngineStat) GetSyncObjectsPmOk() (*float64, bool)`

GetSyncObjectsPmOk returns a tuple with the SyncObjectsPm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncObjectsPm

`func (o *OSSearchEngineStat) SetSyncObjectsPm(v float64)`

SetSyncObjectsPm sets SyncObjectsPm field to given value.

### HasSyncObjectsPm

`func (o *OSSearchEngineStat) HasSyncObjectsPm() bool`

HasSyncObjectsPm returns a boolean if a field has been set.

### GetTotalMetaNum

`func (o *OSSearchEngineStat) GetTotalMetaNum() int64`

GetTotalMetaNum returns the TotalMetaNum field if non-nil, zero value otherwise.

### GetTotalMetaNumOk

`func (o *OSSearchEngineStat) GetTotalMetaNumOk() (*int64, bool)`

GetTotalMetaNumOk returns a tuple with the TotalMetaNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMetaNum

`func (o *OSSearchEngineStat) SetTotalMetaNum(v int64)`

SetTotalMetaNum sets TotalMetaNum field to given value.

### HasTotalMetaNum

`func (o *OSSearchEngineStat) HasTotalMetaNum() bool`

HasTotalMetaNum returns a boolean if a field has been set.

### GetTotalMetaSizeBytes

`func (o *OSSearchEngineStat) GetTotalMetaSizeBytes() int64`

GetTotalMetaSizeBytes returns the TotalMetaSizeBytes field if non-nil, zero value otherwise.

### GetTotalMetaSizeBytesOk

`func (o *OSSearchEngineStat) GetTotalMetaSizeBytesOk() (*int64, bool)`

GetTotalMetaSizeBytesOk returns a tuple with the TotalMetaSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMetaSizeBytes

`func (o *OSSearchEngineStat) SetTotalMetaSizeBytes(v int64)`

SetTotalMetaSizeBytes sets TotalMetaSizeBytes field to given value.

### HasTotalMetaSizeBytes

`func (o *OSSearchEngineStat) HasTotalMetaSizeBytes() bool`

HasTotalMetaSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


