# OSStorageClassStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedObjects** | Pointer to **int64** |  | [optional] 
**AllocatedSize** | Pointer to **int64** |  | [optional] 
**CacheAllocatedObjects** | Pointer to **int64** |  | [optional] 
**CacheAllocatedSize** | Pointer to **int64** |  | [optional] 
**ClassName** | Pointer to **string** | ClassName used in GetObjectStorageUserSamples | [optional] 

## Methods

### NewOSStorageClassStat

`func NewOSStorageClassStat() *OSStorageClassStat`

NewOSStorageClassStat instantiates a new OSStorageClassStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSStorageClassStatWithDefaults

`func NewOSStorageClassStatWithDefaults() *OSStorageClassStat`

NewOSStorageClassStatWithDefaults instantiates a new OSStorageClassStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedObjects

`func (o *OSStorageClassStat) GetAllocatedObjects() int64`

GetAllocatedObjects returns the AllocatedObjects field if non-nil, zero value otherwise.

### GetAllocatedObjectsOk

`func (o *OSStorageClassStat) GetAllocatedObjectsOk() (*int64, bool)`

GetAllocatedObjectsOk returns a tuple with the AllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedObjects

`func (o *OSStorageClassStat) SetAllocatedObjects(v int64)`

SetAllocatedObjects sets AllocatedObjects field to given value.

### HasAllocatedObjects

`func (o *OSStorageClassStat) HasAllocatedObjects() bool`

HasAllocatedObjects returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *OSStorageClassStat) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *OSStorageClassStat) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *OSStorageClassStat) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *OSStorageClassStat) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetCacheAllocatedObjects

`func (o *OSStorageClassStat) GetCacheAllocatedObjects() int64`

GetCacheAllocatedObjects returns the CacheAllocatedObjects field if non-nil, zero value otherwise.

### GetCacheAllocatedObjectsOk

`func (o *OSStorageClassStat) GetCacheAllocatedObjectsOk() (*int64, bool)`

GetCacheAllocatedObjectsOk returns a tuple with the CacheAllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAllocatedObjects

`func (o *OSStorageClassStat) SetCacheAllocatedObjects(v int64)`

SetCacheAllocatedObjects sets CacheAllocatedObjects field to given value.

### HasCacheAllocatedObjects

`func (o *OSStorageClassStat) HasCacheAllocatedObjects() bool`

HasCacheAllocatedObjects returns a boolean if a field has been set.

### GetCacheAllocatedSize

`func (o *OSStorageClassStat) GetCacheAllocatedSize() int64`

GetCacheAllocatedSize returns the CacheAllocatedSize field if non-nil, zero value otherwise.

### GetCacheAllocatedSizeOk

`func (o *OSStorageClassStat) GetCacheAllocatedSizeOk() (*int64, bool)`

GetCacheAllocatedSizeOk returns a tuple with the CacheAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAllocatedSize

`func (o *OSStorageClassStat) SetCacheAllocatedSize(v int64)`

SetCacheAllocatedSize sets CacheAllocatedSize field to given value.

### HasCacheAllocatedSize

`func (o *OSStorageClassStat) HasCacheAllocatedSize() bool`

HasCacheAllocatedSize returns a boolean if a field has been set.

### GetClassName

`func (o *OSStorageClassStat) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *OSStorageClassStat) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *OSStorageClassStat) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *OSStorageClassStat) HasClassName() bool`

HasClassName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


