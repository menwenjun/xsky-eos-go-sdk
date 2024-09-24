# OsdGroup

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

## Methods

### NewOsdGroup

`func NewOsdGroup() *OsdGroup`

NewOsdGroup instantiates a new OsdGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsdGroupWithDefaults

`func NewOsdGroupWithDefaults() *OsdGroup`

NewOsdGroupWithDefaults instantiates a new OsdGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *OsdGroup) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *OsdGroup) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *OsdGroup) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *OsdGroup) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *OsdGroup) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OsdGroup) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OsdGroup) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OsdGroup) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OsdGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OsdGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OsdGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OsdGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDeviceType

`func (o *OsdGroup) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *OsdGroup) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *OsdGroup) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *OsdGroup) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceTypeCheckDisabled

`func (o *OsdGroup) GetDeviceTypeCheckDisabled() bool`

GetDeviceTypeCheckDisabled returns the DeviceTypeCheckDisabled field if non-nil, zero value otherwise.

### GetDeviceTypeCheckDisabledOk

`func (o *OsdGroup) GetDeviceTypeCheckDisabledOk() (*bool, bool)`

GetDeviceTypeCheckDisabledOk returns a tuple with the DeviceTypeCheckDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeCheckDisabled

`func (o *OsdGroup) SetDeviceTypeCheckDisabled(v bool)`

SetDeviceTypeCheckDisabled sets DeviceTypeCheckDisabled field to given value.

### HasDeviceTypeCheckDisabled

`func (o *OsdGroup) HasDeviceTypeCheckDisabled() bool`

HasDeviceTypeCheckDisabled returns a boolean if a field has been set.

### GetFailureDomainType

`func (o *OsdGroup) GetFailureDomainType() string`

GetFailureDomainType returns the FailureDomainType field if non-nil, zero value otherwise.

### GetFailureDomainTypeOk

`func (o *OsdGroup) GetFailureDomainTypeOk() (*string, bool)`

GetFailureDomainTypeOk returns a tuple with the FailureDomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDomainType

`func (o *OsdGroup) SetFailureDomainType(v string)`

SetFailureDomainType sets FailureDomainType field to given value.

### HasFailureDomainType

`func (o *OsdGroup) HasFailureDomainType() bool`

HasFailureDomainType returns a boolean if a field has been set.

### GetId

`func (o *OsdGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OsdGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OsdGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OsdGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumaBindPolicy

`func (o *OsdGroup) GetNumaBindPolicy() string`

GetNumaBindPolicy returns the NumaBindPolicy field if non-nil, zero value otherwise.

### GetNumaBindPolicyOk

`func (o *OsdGroup) GetNumaBindPolicyOk() (*string, bool)`

GetNumaBindPolicyOk returns a tuple with the NumaBindPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaBindPolicy

`func (o *OsdGroup) SetNumaBindPolicy(v string)`

SetNumaBindPolicy sets NumaBindPolicy field to given value.

### HasNumaBindPolicy

`func (o *OsdGroup) HasNumaBindPolicy() bool`

HasNumaBindPolicy returns a boolean if a field has been set.

### GetNumaEnabled

`func (o *OsdGroup) GetNumaEnabled() bool`

GetNumaEnabled returns the NumaEnabled field if non-nil, zero value otherwise.

### GetNumaEnabledOk

`func (o *OsdGroup) GetNumaEnabledOk() (*bool, bool)`

GetNumaEnabledOk returns a tuple with the NumaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaEnabled

`func (o *OsdGroup) SetNumaEnabled(v bool)`

SetNumaEnabled sets NumaEnabled field to given value.

### HasNumaEnabled

`func (o *OsdGroup) HasNumaEnabled() bool`

HasNumaEnabled returns a boolean if a field has been set.

### GetOsdAsyncRecoveryMaxUpdates

`func (o *OsdGroup) GetOsdAsyncRecoveryMaxUpdates() int64`

GetOsdAsyncRecoveryMaxUpdates returns the OsdAsyncRecoveryMaxUpdates field if non-nil, zero value otherwise.

### GetOsdAsyncRecoveryMaxUpdatesOk

`func (o *OsdGroup) GetOsdAsyncRecoveryMaxUpdatesOk() (*int64, bool)`

GetOsdAsyncRecoveryMaxUpdatesOk returns a tuple with the OsdAsyncRecoveryMaxUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdAsyncRecoveryMaxUpdates

`func (o *OsdGroup) SetOsdAsyncRecoveryMaxUpdates(v int64)`

SetOsdAsyncRecoveryMaxUpdates sets OsdAsyncRecoveryMaxUpdates field to given value.

### HasOsdAsyncRecoveryMaxUpdates

`func (o *OsdGroup) HasOsdAsyncRecoveryMaxUpdates() bool`

HasOsdAsyncRecoveryMaxUpdates returns a boolean if a field has been set.

### GetOsdFullRatio

`func (o *OsdGroup) GetOsdFullRatio() float64`

GetOsdFullRatio returns the OsdFullRatio field if non-nil, zero value otherwise.

### GetOsdFullRatioOk

`func (o *OsdGroup) GetOsdFullRatioOk() (*float64, bool)`

GetOsdFullRatioOk returns a tuple with the OsdFullRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdFullRatio

`func (o *OsdGroup) SetOsdFullRatio(v float64)`

SetOsdFullRatio sets OsdFullRatio field to given value.

### HasOsdFullRatio

`func (o *OsdGroup) HasOsdFullRatio() bool`

HasOsdFullRatio returns a boolean if a field has been set.

### GetOsdNum

`func (o *OsdGroup) GetOsdNum() int64`

GetOsdNum returns the OsdNum field if non-nil, zero value otherwise.

### GetOsdNumOk

`func (o *OsdGroup) GetOsdNumOk() (*int64, bool)`

GetOsdNumOk returns a tuple with the OsdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdNum

`func (o *OsdGroup) SetOsdNum(v int64)`

SetOsdNum sets OsdNum field to given value.

### HasOsdNum

`func (o *OsdGroup) HasOsdNum() bool`

HasOsdNum returns a boolean if a field has been set.

### GetPools

`func (o *OsdGroup) GetPools() []PoolNestview`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *OsdGroup) GetPoolsOk() (*[]PoolNestview, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *OsdGroup) SetPools(v []PoolNestview)`

SetPools sets Pools field to given value.

### HasPools

`func (o *OsdGroup) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetQos

`func (o *OsdGroup) GetQos() OsdQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *OsdGroup) GetQosOk() (*OsdQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *OsdGroup) SetQos(v OsdQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *OsdGroup) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetStatus

`func (o *OsdGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OsdGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OsdGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OsdGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *OsdGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OsdGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OsdGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OsdGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


