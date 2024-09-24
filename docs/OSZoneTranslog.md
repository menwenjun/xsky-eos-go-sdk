# OSZoneTranslog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**ResourceUuid** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewOSZoneTranslog

`func NewOSZoneTranslog() *OSZoneTranslog`

NewOSZoneTranslog instantiates a new OSZoneTranslog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSZoneTranslogWithDefaults

`func NewOSZoneTranslogWithDefaults() *OSZoneTranslog`

NewOSZoneTranslogWithDefaults instantiates a new OSZoneTranslog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OSZoneTranslog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OSZoneTranslog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OSZoneTranslog) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *OSZoneTranslog) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCluster

`func (o *OSZoneTranslog) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSZoneTranslog) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSZoneTranslog) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSZoneTranslog) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSZoneTranslog) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSZoneTranslog) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSZoneTranslog) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSZoneTranslog) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetResourceType

`func (o *OSZoneTranslog) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *OSZoneTranslog) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *OSZoneTranslog) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *OSZoneTranslog) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceUuid

`func (o *OSZoneTranslog) GetResourceUuid() string`

GetResourceUuid returns the ResourceUuid field if non-nil, zero value otherwise.

### GetResourceUuidOk

`func (o *OSZoneTranslog) GetResourceUuidOk() (*string, bool)`

GetResourceUuidOk returns a tuple with the ResourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUuid

`func (o *OSZoneTranslog) SetResourceUuid(v string)`

SetResourceUuid sets ResourceUuid field to given value.

### HasResourceUuid

`func (o *OSZoneTranslog) HasResourceUuid() bool`

HasResourceUuid returns a boolean if a field has been set.

### GetUpdate

`func (o *OSZoneTranslog) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSZoneTranslog) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSZoneTranslog) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSZoneTranslog) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *OSZoneTranslog) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OSZoneTranslog) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OSZoneTranslog) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OSZoneTranslog) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


