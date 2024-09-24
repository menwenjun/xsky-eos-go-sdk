# DpBlockReplicationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DpSite** | Pointer to [**DpSiteNestview**](DpSiteNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimeoutSeconds** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpBlockReplicationPolicy

`func NewDpBlockReplicationPolicy() *DpBlockReplicationPolicy`

NewDpBlockReplicationPolicy instantiates a new DpBlockReplicationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockReplicationPolicyWithDefaults

`func NewDpBlockReplicationPolicyWithDefaults() *DpBlockReplicationPolicy`

NewDpBlockReplicationPolicyWithDefaults instantiates a new DpBlockReplicationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DpBlockReplicationPolicy) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpBlockReplicationPolicy) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpBlockReplicationPolicy) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpBlockReplicationPolicy) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DpBlockReplicationPolicy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpBlockReplicationPolicy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpBlockReplicationPolicy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpBlockReplicationPolicy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDpSite

`func (o *DpBlockReplicationPolicy) GetDpSite() DpSiteNestview`

GetDpSite returns the DpSite field if non-nil, zero value otherwise.

### GetDpSiteOk

`func (o *DpBlockReplicationPolicy) GetDpSiteOk() (*DpSiteNestview, bool)`

GetDpSiteOk returns a tuple with the DpSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpSite

`func (o *DpBlockReplicationPolicy) SetDpSite(v DpSiteNestview)`

SetDpSite sets DpSite field to given value.

### HasDpSite

`func (o *DpBlockReplicationPolicy) HasDpSite() bool`

HasDpSite returns a boolean if a field has been set.

### GetId

`func (o *DpBlockReplicationPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpBlockReplicationPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpBlockReplicationPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpBlockReplicationPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DpBlockReplicationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockReplicationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockReplicationPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpBlockReplicationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DpBlockReplicationPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpBlockReplicationPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpBlockReplicationPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpBlockReplicationPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *DpBlockReplicationPolicy) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *DpBlockReplicationPolicy) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *DpBlockReplicationPolicy) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *DpBlockReplicationPolicy) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetUpdate

`func (o *DpBlockReplicationPolicy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpBlockReplicationPolicy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpBlockReplicationPolicy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpBlockReplicationPolicy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


