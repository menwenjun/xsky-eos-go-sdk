# OSRemotePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ZoneUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewOSRemotePolicy

`func NewOSRemotePolicy() *OSRemotePolicy`

NewOSRemotePolicy instantiates a new OSRemotePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSRemotePolicyWithDefaults

`func NewOSRemotePolicyWithDefaults() *OSRemotePolicy`

NewOSRemotePolicyWithDefaults instantiates a new OSRemotePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *OSRemotePolicy) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSRemotePolicy) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSRemotePolicy) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSRemotePolicy) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSRemotePolicy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSRemotePolicy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSRemotePolicy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSRemotePolicy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDefault

`func (o *OSRemotePolicy) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *OSRemotePolicy) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *OSRemotePolicy) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *OSRemotePolicy) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetName

`func (o *OSRemotePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSRemotePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSRemotePolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OSRemotePolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OSRemotePolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSRemotePolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSRemotePolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSRemotePolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *OSRemotePolicy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSRemotePolicy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSRemotePolicy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSRemotePolicy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *OSRemotePolicy) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OSRemotePolicy) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OSRemotePolicy) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OSRemotePolicy) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetZoneUuid

`func (o *OSRemotePolicy) GetZoneUuid() string`

GetZoneUuid returns the ZoneUuid field if non-nil, zero value otherwise.

### GetZoneUuidOk

`func (o *OSRemotePolicy) GetZoneUuidOk() (*string, bool)`

GetZoneUuidOk returns a tuple with the ZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUuid

`func (o *OSRemotePolicy) SetZoneUuid(v string)`

SetZoneUuid sets ZoneUuid field to given value.

### HasZoneUuid

`func (o *OSRemotePolicy) HasZoneUuid() bool`

HasZoneUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


