# ObjectStoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketNum** | Pointer to **int64** |  | [optional] 
**CachePool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Compress** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Crypto** | Pointer to **bool** |  | [optional] 
**Deduplication** | Pointer to **bool** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IndexPool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectSizeThreshold** | Pointer to **int64** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewObjectStoragePolicy

`func NewObjectStoragePolicy() *ObjectStoragePolicy`

NewObjectStoragePolicy instantiates a new ObjectStoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStoragePolicyWithDefaults

`func NewObjectStoragePolicyWithDefaults() *ObjectStoragePolicy`

NewObjectStoragePolicyWithDefaults instantiates a new ObjectStoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketNum

`func (o *ObjectStoragePolicy) GetBucketNum() int64`

GetBucketNum returns the BucketNum field if non-nil, zero value otherwise.

### GetBucketNumOk

`func (o *ObjectStoragePolicy) GetBucketNumOk() (*int64, bool)`

GetBucketNumOk returns a tuple with the BucketNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketNum

`func (o *ObjectStoragePolicy) SetBucketNum(v int64)`

SetBucketNum sets BucketNum field to given value.

### HasBucketNum

`func (o *ObjectStoragePolicy) HasBucketNum() bool`

HasBucketNum returns a boolean if a field has been set.

### GetCachePool

`func (o *ObjectStoragePolicy) GetCachePool() PoolNestview`

GetCachePool returns the CachePool field if non-nil, zero value otherwise.

### GetCachePoolOk

`func (o *ObjectStoragePolicy) GetCachePoolOk() (*PoolNestview, bool)`

GetCachePoolOk returns a tuple with the CachePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePool

`func (o *ObjectStoragePolicy) SetCachePool(v PoolNestview)`

SetCachePool sets CachePool field to given value.

### HasCachePool

`func (o *ObjectStoragePolicy) HasCachePool() bool`

HasCachePool returns a boolean if a field has been set.

### GetCluster

`func (o *ObjectStoragePolicy) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStoragePolicy) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStoragePolicy) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStoragePolicy) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCompress

`func (o *ObjectStoragePolicy) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *ObjectStoragePolicy) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *ObjectStoragePolicy) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *ObjectStoragePolicy) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStoragePolicy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStoragePolicy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStoragePolicy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStoragePolicy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCrypto

`func (o *ObjectStoragePolicy) GetCrypto() bool`

GetCrypto returns the Crypto field if non-nil, zero value otherwise.

### GetCryptoOk

`func (o *ObjectStoragePolicy) GetCryptoOk() (*bool, bool)`

GetCryptoOk returns a tuple with the Crypto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrypto

`func (o *ObjectStoragePolicy) SetCrypto(v bool)`

SetCrypto sets Crypto field to given value.

### HasCrypto

`func (o *ObjectStoragePolicy) HasCrypto() bool`

HasCrypto returns a boolean if a field has been set.

### GetDeduplication

`func (o *ObjectStoragePolicy) GetDeduplication() bool`

GetDeduplication returns the Deduplication field if non-nil, zero value otherwise.

### GetDeduplicationOk

`func (o *ObjectStoragePolicy) GetDeduplicationOk() (*bool, bool)`

GetDeduplicationOk returns a tuple with the Deduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplication

`func (o *ObjectStoragePolicy) SetDeduplication(v bool)`

SetDeduplication sets Deduplication field to given value.

### HasDeduplication

`func (o *ObjectStoragePolicy) HasDeduplication() bool`

HasDeduplication returns a boolean if a field has been set.

### GetDefault

`func (o *ObjectStoragePolicy) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ObjectStoragePolicy) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ObjectStoragePolicy) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ObjectStoragePolicy) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStoragePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStoragePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStoragePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStoragePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ObjectStoragePolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStoragePolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStoragePolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStoragePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndexPool

`func (o *ObjectStoragePolicy) GetIndexPool() PoolNestview`

GetIndexPool returns the IndexPool field if non-nil, zero value otherwise.

### GetIndexPoolOk

`func (o *ObjectStoragePolicy) GetIndexPoolOk() (*PoolNestview, bool)`

GetIndexPoolOk returns a tuple with the IndexPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPool

`func (o *ObjectStoragePolicy) SetIndexPool(v PoolNestview)`

SetIndexPool sets IndexPool field to given value.

### HasIndexPool

`func (o *ObjectStoragePolicy) HasIndexPool() bool`

HasIndexPool returns a boolean if a field has been set.

### GetName

`func (o *ObjectStoragePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStoragePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStoragePolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectStoragePolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectSizeThreshold

`func (o *ObjectStoragePolicy) GetObjectSizeThreshold() int64`

GetObjectSizeThreshold returns the ObjectSizeThreshold field if non-nil, zero value otherwise.

### GetObjectSizeThresholdOk

`func (o *ObjectStoragePolicy) GetObjectSizeThresholdOk() (*int64, bool)`

GetObjectSizeThresholdOk returns a tuple with the ObjectSizeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSizeThreshold

`func (o *ObjectStoragePolicy) SetObjectSizeThreshold(v int64)`

SetObjectSizeThreshold sets ObjectSizeThreshold field to given value.

### HasObjectSizeThreshold

`func (o *ObjectStoragePolicy) HasObjectSizeThreshold() bool`

HasObjectSizeThreshold returns a boolean if a field has been set.

### GetPolicyName

`func (o *ObjectStoragePolicy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ObjectStoragePolicy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ObjectStoragePolicy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ObjectStoragePolicy) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetShared

`func (o *ObjectStoragePolicy) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *ObjectStoragePolicy) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *ObjectStoragePolicy) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *ObjectStoragePolicy) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStoragePolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStoragePolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStoragePolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStoragePolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *ObjectStoragePolicy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ObjectStoragePolicy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ObjectStoragePolicy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ObjectStoragePolicy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


