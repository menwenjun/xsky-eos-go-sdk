# OSCustomLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to [**ObjectStorageBucketNestview**](ObjectStorageBucketNestview.md) |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOSCustomLabel

`func NewOSCustomLabel() *OSCustomLabel`

NewOSCustomLabel instantiates a new OSCustomLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSCustomLabelWithDefaults

`func NewOSCustomLabelWithDefaults() *OSCustomLabel`

NewOSCustomLabelWithDefaults instantiates a new OSCustomLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *OSCustomLabel) GetBucket() ObjectStorageBucketNestview`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *OSCustomLabel) GetBucketOk() (*ObjectStorageBucketNestview, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *OSCustomLabel) SetBucket(v ObjectStorageBucketNestview)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *OSCustomLabel) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCategory

`func (o *OSCustomLabel) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OSCustomLabel) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OSCustomLabel) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OSCustomLabel) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCluster

`func (o *OSCustomLabel) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSCustomLabel) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSCustomLabel) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSCustomLabel) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSCustomLabel) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSCustomLabel) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSCustomLabel) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSCustomLabel) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *OSCustomLabel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSCustomLabel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSCustomLabel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSCustomLabel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OSCustomLabel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSCustomLabel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSCustomLabel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OSCustomLabel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OSCustomLabel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSCustomLabel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSCustomLabel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSCustomLabel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *OSCustomLabel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OSCustomLabel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OSCustomLabel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OSCustomLabel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *OSCustomLabel) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSCustomLabel) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSCustomLabel) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSCustomLabel) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


