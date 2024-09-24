# OSZoneLock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**ExpiredTime** | Pointer to **time.Time** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewOSZoneLock

`func NewOSZoneLock() *OSZoneLock`

NewOSZoneLock instantiates a new OSZoneLock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSZoneLockWithDefaults

`func NewOSZoneLockWithDefaults() *OSZoneLock`

NewOSZoneLockWithDefaults instantiates a new OSZoneLock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *OSZoneLock) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSZoneLock) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSZoneLock) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSZoneLock) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSZoneLock) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSZoneLock) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSZoneLock) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSZoneLock) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetExpiredTime

`func (o *OSZoneLock) GetExpiredTime() time.Time`

GetExpiredTime returns the ExpiredTime field if non-nil, zero value otherwise.

### GetExpiredTimeOk

`func (o *OSZoneLock) GetExpiredTimeOk() (*time.Time, bool)`

GetExpiredTimeOk returns a tuple with the ExpiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTime

`func (o *OSZoneLock) SetExpiredTime(v time.Time)`

SetExpiredTime sets ExpiredTime field to given value.

### HasExpiredTime

`func (o *OSZoneLock) HasExpiredTime() bool`

HasExpiredTime returns a boolean if a field has been set.

### GetKey

`func (o *OSZoneLock) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OSZoneLock) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OSZoneLock) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OSZoneLock) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetResourceType

`func (o *OSZoneLock) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *OSZoneLock) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *OSZoneLock) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *OSZoneLock) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetUpdate

`func (o *OSZoneLock) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSZoneLock) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSZoneLock) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSZoneLock) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *OSZoneLock) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OSZoneLock) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OSZoneLock) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OSZoneLock) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetValue

`func (o *OSZoneLock) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OSZoneLock) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OSZoneLock) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OSZoneLock) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


