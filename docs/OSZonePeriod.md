# OSZonePeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MasterOsZoneUuid** | Pointer to **string** |  | [optional] 
**NextMasterOsZoneUuid** | Pointer to **string** |  | [optional] 
**PeriodId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewOSZonePeriod

`func NewOSZonePeriod() *OSZonePeriod`

NewOSZonePeriod instantiates a new OSZonePeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSZonePeriodWithDefaults

`func NewOSZonePeriodWithDefaults() *OSZonePeriod`

NewOSZonePeriodWithDefaults instantiates a new OSZonePeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OSZonePeriod) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OSZonePeriod) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OSZonePeriod) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *OSZonePeriod) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCluster

`func (o *OSZonePeriod) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSZonePeriod) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSZonePeriod) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSZonePeriod) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSZonePeriod) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSZonePeriod) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSZonePeriod) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSZonePeriod) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *OSZonePeriod) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSZonePeriod) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSZonePeriod) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSZonePeriod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMasterOsZoneUuid

`func (o *OSZonePeriod) GetMasterOsZoneUuid() string`

GetMasterOsZoneUuid returns the MasterOsZoneUuid field if non-nil, zero value otherwise.

### GetMasterOsZoneUuidOk

`func (o *OSZonePeriod) GetMasterOsZoneUuidOk() (*string, bool)`

GetMasterOsZoneUuidOk returns a tuple with the MasterOsZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOsZoneUuid

`func (o *OSZonePeriod) SetMasterOsZoneUuid(v string)`

SetMasterOsZoneUuid sets MasterOsZoneUuid field to given value.

### HasMasterOsZoneUuid

`func (o *OSZonePeriod) HasMasterOsZoneUuid() bool`

HasMasterOsZoneUuid returns a boolean if a field has been set.

### GetNextMasterOsZoneUuid

`func (o *OSZonePeriod) GetNextMasterOsZoneUuid() string`

GetNextMasterOsZoneUuid returns the NextMasterOsZoneUuid field if non-nil, zero value otherwise.

### GetNextMasterOsZoneUuidOk

`func (o *OSZonePeriod) GetNextMasterOsZoneUuidOk() (*string, bool)`

GetNextMasterOsZoneUuidOk returns a tuple with the NextMasterOsZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMasterOsZoneUuid

`func (o *OSZonePeriod) SetNextMasterOsZoneUuid(v string)`

SetNextMasterOsZoneUuid sets NextMasterOsZoneUuid field to given value.

### HasNextMasterOsZoneUuid

`func (o *OSZonePeriod) HasNextMasterOsZoneUuid() bool`

HasNextMasterOsZoneUuid returns a boolean if a field has been set.

### GetPeriodId

`func (o *OSZonePeriod) GetPeriodId() string`

GetPeriodId returns the PeriodId field if non-nil, zero value otherwise.

### GetPeriodIdOk

`func (o *OSZonePeriod) GetPeriodIdOk() (*string, bool)`

GetPeriodIdOk returns a tuple with the PeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodId

`func (o *OSZonePeriod) SetPeriodId(v string)`

SetPeriodId sets PeriodId field to given value.

### HasPeriodId

`func (o *OSZonePeriod) HasPeriodId() bool`

HasPeriodId returns a boolean if a field has been set.

### GetStatus

`func (o *OSZonePeriod) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSZonePeriod) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSZonePeriod) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSZonePeriod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *OSZonePeriod) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSZonePeriod) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSZonePeriod) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSZonePeriod) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *OSZonePeriod) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OSZonePeriod) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OSZonePeriod) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OSZonePeriod) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


