# OSSample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteObjCategory** | Pointer to [**OSBucketUsageCategory**](OSBucketUsageCategory.md) |  | [optional] 
**GetObjCategory** | Pointer to [**OSBucketUsageCategory**](OSBucketUsageCategory.md) |  | [optional] 
**ListBucketCategory** | Pointer to [**OSBucketUsageCategory**](OSBucketUsageCategory.md) |  | [optional] 
**Objects** | Pointer to **int64** |  | [optional] 
**PostObjCategory** | Pointer to [**OSBucketUsageCategory**](OSBucketUsageCategory.md) |  | [optional] 
**PutObjCategory** | Pointer to [**OSBucketUsageCategory**](OSBucketUsageCategory.md) |  | [optional] 
**UsedCapacityBytes** | Pointer to **map[string]int64** |  | [optional] 

## Methods

### NewOSSample

`func NewOSSample() *OSSample`

NewOSSample instantiates a new OSSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSSampleWithDefaults

`func NewOSSampleWithDefaults() *OSSample`

NewOSSampleWithDefaults instantiates a new OSSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteObjCategory

`func (o *OSSample) GetDeleteObjCategory() OSBucketUsageCategory`

GetDeleteObjCategory returns the DeleteObjCategory field if non-nil, zero value otherwise.

### GetDeleteObjCategoryOk

`func (o *OSSample) GetDeleteObjCategoryOk() (*OSBucketUsageCategory, bool)`

GetDeleteObjCategoryOk returns a tuple with the DeleteObjCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteObjCategory

`func (o *OSSample) SetDeleteObjCategory(v OSBucketUsageCategory)`

SetDeleteObjCategory sets DeleteObjCategory field to given value.

### HasDeleteObjCategory

`func (o *OSSample) HasDeleteObjCategory() bool`

HasDeleteObjCategory returns a boolean if a field has been set.

### GetGetObjCategory

`func (o *OSSample) GetGetObjCategory() OSBucketUsageCategory`

GetGetObjCategory returns the GetObjCategory field if non-nil, zero value otherwise.

### GetGetObjCategoryOk

`func (o *OSSample) GetGetObjCategoryOk() (*OSBucketUsageCategory, bool)`

GetGetObjCategoryOk returns a tuple with the GetObjCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetObjCategory

`func (o *OSSample) SetGetObjCategory(v OSBucketUsageCategory)`

SetGetObjCategory sets GetObjCategory field to given value.

### HasGetObjCategory

`func (o *OSSample) HasGetObjCategory() bool`

HasGetObjCategory returns a boolean if a field has been set.

### GetListBucketCategory

`func (o *OSSample) GetListBucketCategory() OSBucketUsageCategory`

GetListBucketCategory returns the ListBucketCategory field if non-nil, zero value otherwise.

### GetListBucketCategoryOk

`func (o *OSSample) GetListBucketCategoryOk() (*OSBucketUsageCategory, bool)`

GetListBucketCategoryOk returns a tuple with the ListBucketCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListBucketCategory

`func (o *OSSample) SetListBucketCategory(v OSBucketUsageCategory)`

SetListBucketCategory sets ListBucketCategory field to given value.

### HasListBucketCategory

`func (o *OSSample) HasListBucketCategory() bool`

HasListBucketCategory returns a boolean if a field has been set.

### GetObjects

`func (o *OSSample) GetObjects() int64`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *OSSample) GetObjectsOk() (*int64, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *OSSample) SetObjects(v int64)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *OSSample) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetPostObjCategory

`func (o *OSSample) GetPostObjCategory() OSBucketUsageCategory`

GetPostObjCategory returns the PostObjCategory field if non-nil, zero value otherwise.

### GetPostObjCategoryOk

`func (o *OSSample) GetPostObjCategoryOk() (*OSBucketUsageCategory, bool)`

GetPostObjCategoryOk returns a tuple with the PostObjCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostObjCategory

`func (o *OSSample) SetPostObjCategory(v OSBucketUsageCategory)`

SetPostObjCategory sets PostObjCategory field to given value.

### HasPostObjCategory

`func (o *OSSample) HasPostObjCategory() bool`

HasPostObjCategory returns a boolean if a field has been set.

### GetPutObjCategory

`func (o *OSSample) GetPutObjCategory() OSBucketUsageCategory`

GetPutObjCategory returns the PutObjCategory field if non-nil, zero value otherwise.

### GetPutObjCategoryOk

`func (o *OSSample) GetPutObjCategoryOk() (*OSBucketUsageCategory, bool)`

GetPutObjCategoryOk returns a tuple with the PutObjCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutObjCategory

`func (o *OSSample) SetPutObjCategory(v OSBucketUsageCategory)`

SetPutObjCategory sets PutObjCategory field to given value.

### HasPutObjCategory

`func (o *OSSample) HasPutObjCategory() bool`

HasPutObjCategory returns a boolean if a field has been set.

### GetUsedCapacityBytes

`func (o *OSSample) GetUsedCapacityBytes() map[string]int64`

GetUsedCapacityBytes returns the UsedCapacityBytes field if non-nil, zero value otherwise.

### GetUsedCapacityBytesOk

`func (o *OSSample) GetUsedCapacityBytesOk() (*map[string]int64, bool)`

GetUsedCapacityBytesOk returns a tuple with the UsedCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCapacityBytes

`func (o *OSSample) SetUsedCapacityBytes(v map[string]int64)`

SetUsedCapacityBytes sets UsedCapacityBytes field to given value.

### HasUsedCapacityBytes

`func (o *OSSample) HasUsedCapacityBytes() bool`

HasUsedCapacityBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


