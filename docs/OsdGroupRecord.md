# OsdGroupRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**DeviceTypeCheckDisabled** | Pointer to **bool** |  | [optional] 
**FailureDomainType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**NumaBindPolicy** | Pointer to **string** |  | [optional] 
**NumaEnabled** | Pointer to **bool** |  | [optional] 
**OsdAsyncRecoveryMaxUpdates** | Pointer to **int64** |  | [optional] 
**OsdFullRatio** | Pointer to **float64** |  | [optional] 
**OsdNum** | Pointer to **int64** |  | [optional] 
**Pools** | Pointer to [**[]PoolNestview**](PoolNestview.md) |  | [optional] 
**Qos** | Pointer to [**OsdQos**](OsdQos.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Samples** | Pointer to [**[]OsdGroupStat**](OsdGroupStat.md) |  | [optional] 

## Methods

### NewOsdGroupRecord

`func NewOsdGroupRecord() *OsdGroupRecord`

NewOsdGroupRecord instantiates a new OsdGroupRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsdGroupRecordWithDefaults

`func NewOsdGroupRecordWithDefaults() *OsdGroupRecord`

NewOsdGroupRecordWithDefaults instantiates a new OsdGroupRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *OsdGroupRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *OsdGroupRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *OsdGroupRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *OsdGroupRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *OsdGroupRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OsdGroupRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OsdGroupRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OsdGroupRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OsdGroupRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OsdGroupRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OsdGroupRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OsdGroupRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDeviceType

`func (o *OsdGroupRecord) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *OsdGroupRecord) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *OsdGroupRecord) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *OsdGroupRecord) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceTypeCheckDisabled

`func (o *OsdGroupRecord) GetDeviceTypeCheckDisabled() bool`

GetDeviceTypeCheckDisabled returns the DeviceTypeCheckDisabled field if non-nil, zero value otherwise.

### GetDeviceTypeCheckDisabledOk

`func (o *OsdGroupRecord) GetDeviceTypeCheckDisabledOk() (*bool, bool)`

GetDeviceTypeCheckDisabledOk returns a tuple with the DeviceTypeCheckDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeCheckDisabled

`func (o *OsdGroupRecord) SetDeviceTypeCheckDisabled(v bool)`

SetDeviceTypeCheckDisabled sets DeviceTypeCheckDisabled field to given value.

### HasDeviceTypeCheckDisabled

`func (o *OsdGroupRecord) HasDeviceTypeCheckDisabled() bool`

HasDeviceTypeCheckDisabled returns a boolean if a field has been set.

### GetFailureDomainType

`func (o *OsdGroupRecord) GetFailureDomainType() string`

GetFailureDomainType returns the FailureDomainType field if non-nil, zero value otherwise.

### GetFailureDomainTypeOk

`func (o *OsdGroupRecord) GetFailureDomainTypeOk() (*string, bool)`

GetFailureDomainTypeOk returns a tuple with the FailureDomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDomainType

`func (o *OsdGroupRecord) SetFailureDomainType(v string)`

SetFailureDomainType sets FailureDomainType field to given value.

### HasFailureDomainType

`func (o *OsdGroupRecord) HasFailureDomainType() bool`

HasFailureDomainType returns a boolean if a field has been set.

### GetId

`func (o *OsdGroupRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OsdGroupRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OsdGroupRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OsdGroupRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumaBindPolicy

`func (o *OsdGroupRecord) GetNumaBindPolicy() string`

GetNumaBindPolicy returns the NumaBindPolicy field if non-nil, zero value otherwise.

### GetNumaBindPolicyOk

`func (o *OsdGroupRecord) GetNumaBindPolicyOk() (*string, bool)`

GetNumaBindPolicyOk returns a tuple with the NumaBindPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaBindPolicy

`func (o *OsdGroupRecord) SetNumaBindPolicy(v string)`

SetNumaBindPolicy sets NumaBindPolicy field to given value.

### HasNumaBindPolicy

`func (o *OsdGroupRecord) HasNumaBindPolicy() bool`

HasNumaBindPolicy returns a boolean if a field has been set.

### GetNumaEnabled

`func (o *OsdGroupRecord) GetNumaEnabled() bool`

GetNumaEnabled returns the NumaEnabled field if non-nil, zero value otherwise.

### GetNumaEnabledOk

`func (o *OsdGroupRecord) GetNumaEnabledOk() (*bool, bool)`

GetNumaEnabledOk returns a tuple with the NumaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaEnabled

`func (o *OsdGroupRecord) SetNumaEnabled(v bool)`

SetNumaEnabled sets NumaEnabled field to given value.

### HasNumaEnabled

`func (o *OsdGroupRecord) HasNumaEnabled() bool`

HasNumaEnabled returns a boolean if a field has been set.

### GetOsdAsyncRecoveryMaxUpdates

`func (o *OsdGroupRecord) GetOsdAsyncRecoveryMaxUpdates() int64`

GetOsdAsyncRecoveryMaxUpdates returns the OsdAsyncRecoveryMaxUpdates field if non-nil, zero value otherwise.

### GetOsdAsyncRecoveryMaxUpdatesOk

`func (o *OsdGroupRecord) GetOsdAsyncRecoveryMaxUpdatesOk() (*int64, bool)`

GetOsdAsyncRecoveryMaxUpdatesOk returns a tuple with the OsdAsyncRecoveryMaxUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdAsyncRecoveryMaxUpdates

`func (o *OsdGroupRecord) SetOsdAsyncRecoveryMaxUpdates(v int64)`

SetOsdAsyncRecoveryMaxUpdates sets OsdAsyncRecoveryMaxUpdates field to given value.

### HasOsdAsyncRecoveryMaxUpdates

`func (o *OsdGroupRecord) HasOsdAsyncRecoveryMaxUpdates() bool`

HasOsdAsyncRecoveryMaxUpdates returns a boolean if a field has been set.

### GetOsdFullRatio

`func (o *OsdGroupRecord) GetOsdFullRatio() float64`

GetOsdFullRatio returns the OsdFullRatio field if non-nil, zero value otherwise.

### GetOsdFullRatioOk

`func (o *OsdGroupRecord) GetOsdFullRatioOk() (*float64, bool)`

GetOsdFullRatioOk returns a tuple with the OsdFullRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdFullRatio

`func (o *OsdGroupRecord) SetOsdFullRatio(v float64)`

SetOsdFullRatio sets OsdFullRatio field to given value.

### HasOsdFullRatio

`func (o *OsdGroupRecord) HasOsdFullRatio() bool`

HasOsdFullRatio returns a boolean if a field has been set.

### GetOsdNum

`func (o *OsdGroupRecord) GetOsdNum() int64`

GetOsdNum returns the OsdNum field if non-nil, zero value otherwise.

### GetOsdNumOk

`func (o *OsdGroupRecord) GetOsdNumOk() (*int64, bool)`

GetOsdNumOk returns a tuple with the OsdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdNum

`func (o *OsdGroupRecord) SetOsdNum(v int64)`

SetOsdNum sets OsdNum field to given value.

### HasOsdNum

`func (o *OsdGroupRecord) HasOsdNum() bool`

HasOsdNum returns a boolean if a field has been set.

### GetPools

`func (o *OsdGroupRecord) GetPools() []PoolNestview`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *OsdGroupRecord) GetPoolsOk() (*[]PoolNestview, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *OsdGroupRecord) SetPools(v []PoolNestview)`

SetPools sets Pools field to given value.

### HasPools

`func (o *OsdGroupRecord) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetQos

`func (o *OsdGroupRecord) GetQos() OsdQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *OsdGroupRecord) GetQosOk() (*OsdQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *OsdGroupRecord) SetQos(v OsdQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *OsdGroupRecord) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetStatus

`func (o *OsdGroupRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OsdGroupRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OsdGroupRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OsdGroupRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *OsdGroupRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OsdGroupRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OsdGroupRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OsdGroupRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *OsdGroupRecord) GetSamples() []OsdGroupStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *OsdGroupRecord) GetSamplesOk() (*[]OsdGroupStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *OsdGroupRecord) SetSamples(v []OsdGroupStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *OsdGroupRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


