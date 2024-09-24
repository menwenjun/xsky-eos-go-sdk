# ObjectStoragePolicyCreateReqPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CachePoolId** | Pointer to **int64** |  | [optional] 
**Compress** | Pointer to **bool** |  | [optional] 
**Crypto** | Pointer to **bool** |  | [optional] 
**Deduplication** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IndexPoolId** | **int64** |  | 
**Name** | **string** |  | 
**Shared** | Pointer to **bool** |  | [optional] 
**StorageClasses** | [**[]ObjectStoragePolicyCreateReqPolicyStorageClassesElt**](ObjectStoragePolicyCreateReqPolicyStorageClassesElt.md) |  | 

## Methods

### NewObjectStoragePolicyCreateReqPolicy

`func NewObjectStoragePolicyCreateReqPolicy(indexPoolId int64, name string, storageClasses []ObjectStoragePolicyCreateReqPolicyStorageClassesElt, ) *ObjectStoragePolicyCreateReqPolicy`

NewObjectStoragePolicyCreateReqPolicy instantiates a new ObjectStoragePolicyCreateReqPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStoragePolicyCreateReqPolicyWithDefaults

`func NewObjectStoragePolicyCreateReqPolicyWithDefaults() *ObjectStoragePolicyCreateReqPolicy`

NewObjectStoragePolicyCreateReqPolicyWithDefaults instantiates a new ObjectStoragePolicyCreateReqPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCachePoolId

`func (o *ObjectStoragePolicyCreateReqPolicy) GetCachePoolId() int64`

GetCachePoolId returns the CachePoolId field if non-nil, zero value otherwise.

### GetCachePoolIdOk

`func (o *ObjectStoragePolicyCreateReqPolicy) GetCachePoolIdOk() (*int64, bool)`

GetCachePoolIdOk returns a tuple with the CachePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePoolId

`func (o *ObjectStoragePolicyCreateReqPolicy) SetCachePoolId(v int64)`

SetCachePoolId sets CachePoolId field to given value.

### HasCachePoolId

`func (o *ObjectStoragePolicyCreateReqPolicy) HasCachePoolId() bool`

HasCachePoolId returns a boolean if a field has been set.

### GetCompress

`func (o *ObjectStoragePolicyCreateReqPolicy) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *ObjectStoragePolicyCreateReqPolicy) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *ObjectStoragePolicyCreateReqPolicy) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *ObjectStoragePolicyCreateReqPolicy) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetCrypto

`func (o *ObjectStoragePolicyCreateReqPolicy) GetCrypto() bool`

GetCrypto returns the Crypto field if non-nil, zero value otherwise.

### GetCryptoOk

`func (o *ObjectStoragePolicyCreateReqPolicy) GetCryptoOk() (*bool, bool)`

GetCryptoOk returns a tuple with the Crypto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrypto

`func (o *ObjectStoragePolicyCreateReqPolicy) SetCrypto(v bool)`

SetCrypto sets Crypto field to given value.

### HasCrypto

`func (o *ObjectStoragePolicyCreateReqPolicy) HasCrypto() bool`

HasCrypto returns a boolean if a field has been set.

### GetDeduplication

`func (o *ObjectStoragePolicyCreateReqPolicy) GetDeduplication() bool`

GetDeduplication returns the Deduplication field if non-nil, zero value otherwise.

### GetDeduplicationOk

`func (o *ObjectStoragePolicyCreateReqPolicy) GetDeduplicationOk() (*bool, bool)`

GetDeduplicationOk returns a tuple with the Deduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplication

`func (o *ObjectStoragePolicyCreateReqPolicy) SetDeduplication(v bool)`

SetDeduplication sets Deduplication field to given value.

### HasDeduplication

`func (o *ObjectStoragePolicyCreateReqPolicy) HasDeduplication() bool`

HasDeduplication returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStoragePolicyCreateReqPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStoragePolicyCreateReqPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStoragePolicyCreateReqPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStoragePolicyCreateReqPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIndexPoolId

`func (o *ObjectStoragePolicyCreateReqPolicy) GetIndexPoolId() int64`

GetIndexPoolId returns the IndexPoolId field if non-nil, zero value otherwise.

### GetIndexPoolIdOk

`func (o *ObjectStoragePolicyCreateReqPolicy) GetIndexPoolIdOk() (*int64, bool)`

GetIndexPoolIdOk returns a tuple with the IndexPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPoolId

`func (o *ObjectStoragePolicyCreateReqPolicy) SetIndexPoolId(v int64)`

SetIndexPoolId sets IndexPoolId field to given value.


### GetName

`func (o *ObjectStoragePolicyCreateReqPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStoragePolicyCreateReqPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStoragePolicyCreateReqPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetShared

`func (o *ObjectStoragePolicyCreateReqPolicy) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *ObjectStoragePolicyCreateReqPolicy) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *ObjectStoragePolicyCreateReqPolicy) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *ObjectStoragePolicyCreateReqPolicy) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetStorageClasses

`func (o *ObjectStoragePolicyCreateReqPolicy) GetStorageClasses() []ObjectStoragePolicyCreateReqPolicyStorageClassesElt`

GetStorageClasses returns the StorageClasses field if non-nil, zero value otherwise.

### GetStorageClassesOk

`func (o *ObjectStoragePolicyCreateReqPolicy) GetStorageClassesOk() (*[]ObjectStoragePolicyCreateReqPolicyStorageClassesElt, bool)`

GetStorageClassesOk returns a tuple with the StorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClasses

`func (o *ObjectStoragePolicyCreateReqPolicy) SetStorageClasses(v []ObjectStoragePolicyCreateReqPolicyStorageClassesElt)`

SetStorageClasses sets StorageClasses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


