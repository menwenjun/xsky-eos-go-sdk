# ObjectStorageBucketNFSClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to [**ObjectStorageBucketNestview**](ObjectStorageBucketNestview.md) |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewObjectStorageBucketNFSClient

`func NewObjectStorageBucketNFSClient() *ObjectStorageBucketNFSClient`

NewObjectStorageBucketNFSClient instantiates a new ObjectStorageBucketNFSClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketNFSClientWithDefaults

`func NewObjectStorageBucketNFSClientWithDefaults() *ObjectStorageBucketNFSClient`

NewObjectStorageBucketNFSClientWithDefaults instantiates a new ObjectStorageBucketNFSClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *ObjectStorageBucketNFSClient) GetBucket() ObjectStorageBucketNestview`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ObjectStorageBucketNFSClient) GetBucketOk() (*ObjectStorageBucketNestview, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ObjectStorageBucketNFSClient) SetBucket(v ObjectStorageBucketNestview)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ObjectStorageBucketNFSClient) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetClient

`func (o *ObjectStorageBucketNFSClient) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ObjectStorageBucketNFSClient) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ObjectStorageBucketNFSClient) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *ObjectStorageBucketNFSClient) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCluster

`func (o *ObjectStorageBucketNFSClient) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageBucketNFSClient) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageBucketNFSClient) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageBucketNFSClient) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageBucketNFSClient) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageBucketNFSClient) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageBucketNFSClient) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageBucketNFSClient) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageBucketNFSClient) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageBucketNFSClient) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageBucketNFSClient) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageBucketNFSClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermission

`func (o *ObjectStorageBucketNFSClient) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ObjectStorageBucketNFSClient) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ObjectStorageBucketNFSClient) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ObjectStorageBucketNFSClient) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageBucketNFSClient) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageBucketNFSClient) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageBucketNFSClient) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageBucketNFSClient) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *ObjectStorageBucketNFSClient) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ObjectStorageBucketNFSClient) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ObjectStorageBucketNFSClient) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ObjectStorageBucketNFSClient) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


