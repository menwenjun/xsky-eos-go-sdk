# NgObjectStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**CnameEnabled** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DomainNames** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**VirtualBucketDump** | Pointer to **bool** |  | [optional] 

## Methods

### NewNgObjectStorage

`func NewNgObjectStorage() *NgObjectStorage`

NewNgObjectStorage instantiates a new NgObjectStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgObjectStorageWithDefaults

`func NewNgObjectStorageWithDefaults() *NgObjectStorage`

NewNgObjectStorageWithDefaults instantiates a new NgObjectStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *NgObjectStorage) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *NgObjectStorage) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *NgObjectStorage) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *NgObjectStorage) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *NgObjectStorage) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *NgObjectStorage) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *NgObjectStorage) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *NgObjectStorage) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCnameEnabled

`func (o *NgObjectStorage) GetCnameEnabled() bool`

GetCnameEnabled returns the CnameEnabled field if non-nil, zero value otherwise.

### GetCnameEnabledOk

`func (o *NgObjectStorage) GetCnameEnabledOk() (*bool, bool)`

GetCnameEnabledOk returns a tuple with the CnameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameEnabled

`func (o *NgObjectStorage) SetCnameEnabled(v bool)`

SetCnameEnabled sets CnameEnabled field to given value.

### HasCnameEnabled

`func (o *NgObjectStorage) HasCnameEnabled() bool`

HasCnameEnabled returns a boolean if a field has been set.

### GetCreate

`func (o *NgObjectStorage) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *NgObjectStorage) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *NgObjectStorage) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *NgObjectStorage) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDomainNames

`func (o *NgObjectStorage) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *NgObjectStorage) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *NgObjectStorage) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *NgObjectStorage) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetId

`func (o *NgObjectStorage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NgObjectStorage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NgObjectStorage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NgObjectStorage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NgObjectStorage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NgObjectStorage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NgObjectStorage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NgObjectStorage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *NgObjectStorage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NgObjectStorage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NgObjectStorage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NgObjectStorage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVirtualBucketDump

`func (o *NgObjectStorage) GetVirtualBucketDump() bool`

GetVirtualBucketDump returns the VirtualBucketDump field if non-nil, zero value otherwise.

### GetVirtualBucketDumpOk

`func (o *NgObjectStorage) GetVirtualBucketDumpOk() (*bool, bool)`

GetVirtualBucketDumpOk returns a tuple with the VirtualBucketDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualBucketDump

`func (o *NgObjectStorage) SetVirtualBucketDump(v bool)`

SetVirtualBucketDump sets VirtualBucketDump field to given value.

### HasVirtualBucketDump

`func (o *NgObjectStorage) HasVirtualBucketDump() bool`

HasVirtualBucketDump returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


