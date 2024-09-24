# NFSGatewayBucketMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to [**ObjectStorageBucketNestview**](ObjectStorageBucketNestview.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Key** | Pointer to [**ObjectStorageKeyNestview**](ObjectStorageKeyNestview.md) |  | [optional] 
**MountClients** | Pointer to **string** |  | [optional] 
**MountNum** | Pointer to **int64** |  | [optional] 
**NfsGateway** | Pointer to [**NFSGatewayNestview**](NFSGatewayNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNFSGatewayBucketMap

`func NewNFSGatewayBucketMap() *NFSGatewayBucketMap`

NewNFSGatewayBucketMap instantiates a new NFSGatewayBucketMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFSGatewayBucketMapWithDefaults

`func NewNFSGatewayBucketMapWithDefaults() *NFSGatewayBucketMap`

NewNFSGatewayBucketMapWithDefaults instantiates a new NFSGatewayBucketMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *NFSGatewayBucketMap) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *NFSGatewayBucketMap) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *NFSGatewayBucketMap) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *NFSGatewayBucketMap) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetBucket

`func (o *NFSGatewayBucketMap) GetBucket() ObjectStorageBucketNestview`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *NFSGatewayBucketMap) GetBucketOk() (*ObjectStorageBucketNestview, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *NFSGatewayBucketMap) SetBucket(v ObjectStorageBucketNestview)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *NFSGatewayBucketMap) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCluster

`func (o *NFSGatewayBucketMap) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *NFSGatewayBucketMap) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *NFSGatewayBucketMap) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *NFSGatewayBucketMap) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *NFSGatewayBucketMap) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *NFSGatewayBucketMap) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *NFSGatewayBucketMap) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *NFSGatewayBucketMap) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *NFSGatewayBucketMap) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NFSGatewayBucketMap) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NFSGatewayBucketMap) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NFSGatewayBucketMap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *NFSGatewayBucketMap) GetKey() ObjectStorageKeyNestview`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NFSGatewayBucketMap) GetKeyOk() (*ObjectStorageKeyNestview, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NFSGatewayBucketMap) SetKey(v ObjectStorageKeyNestview)`

SetKey sets Key field to given value.

### HasKey

`func (o *NFSGatewayBucketMap) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMountClients

`func (o *NFSGatewayBucketMap) GetMountClients() string`

GetMountClients returns the MountClients field if non-nil, zero value otherwise.

### GetMountClientsOk

`func (o *NFSGatewayBucketMap) GetMountClientsOk() (*string, bool)`

GetMountClientsOk returns a tuple with the MountClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountClients

`func (o *NFSGatewayBucketMap) SetMountClients(v string)`

SetMountClients sets MountClients field to given value.

### HasMountClients

`func (o *NFSGatewayBucketMap) HasMountClients() bool`

HasMountClients returns a boolean if a field has been set.

### GetMountNum

`func (o *NFSGatewayBucketMap) GetMountNum() int64`

GetMountNum returns the MountNum field if non-nil, zero value otherwise.

### GetMountNumOk

`func (o *NFSGatewayBucketMap) GetMountNumOk() (*int64, bool)`

GetMountNumOk returns a tuple with the MountNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountNum

`func (o *NFSGatewayBucketMap) SetMountNum(v int64)`

SetMountNum sets MountNum field to given value.

### HasMountNum

`func (o *NFSGatewayBucketMap) HasMountNum() bool`

HasMountNum returns a boolean if a field has been set.

### GetNfsGateway

`func (o *NFSGatewayBucketMap) GetNfsGateway() NFSGatewayNestview`

GetNfsGateway returns the NfsGateway field if non-nil, zero value otherwise.

### GetNfsGatewayOk

`func (o *NFSGatewayBucketMap) GetNfsGatewayOk() (*NFSGatewayNestview, bool)`

GetNfsGatewayOk returns a tuple with the NfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsGateway

`func (o *NFSGatewayBucketMap) SetNfsGateway(v NFSGatewayNestview)`

SetNfsGateway sets NfsGateway field to given value.

### HasNfsGateway

`func (o *NFSGatewayBucketMap) HasNfsGateway() bool`

HasNfsGateway returns a boolean if a field has been set.

### GetStatus

`func (o *NFSGatewayBucketMap) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NFSGatewayBucketMap) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NFSGatewayBucketMap) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NFSGatewayBucketMap) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *NFSGatewayBucketMap) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *NFSGatewayBucketMap) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *NFSGatewayBucketMap) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *NFSGatewayBucketMap) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


