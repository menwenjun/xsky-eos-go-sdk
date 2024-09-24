# ObjectStoragePolicyStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageClasses** | Pointer to [**map[string]OSStorageClassStat**](OSStorageClassStat.md) |  | [optional] 

## Methods

### NewObjectStoragePolicyStat

`func NewObjectStoragePolicyStat() *ObjectStoragePolicyStat`

NewObjectStoragePolicyStat instantiates a new ObjectStoragePolicyStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStoragePolicyStatWithDefaults

`func NewObjectStoragePolicyStatWithDefaults() *ObjectStoragePolicyStat`

NewObjectStoragePolicyStatWithDefaults instantiates a new ObjectStoragePolicyStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageClasses

`func (o *ObjectStoragePolicyStat) GetStorageClasses() map[string]OSStorageClassStat`

GetStorageClasses returns the StorageClasses field if non-nil, zero value otherwise.

### GetStorageClassesOk

`func (o *ObjectStoragePolicyStat) GetStorageClassesOk() (*map[string]OSStorageClassStat, bool)`

GetStorageClassesOk returns a tuple with the StorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClasses

`func (o *ObjectStoragePolicyStat) SetStorageClasses(v map[string]OSStorageClassStat)`

SetStorageClasses sets StorageClasses field to given value.

### HasStorageClasses

`func (o *ObjectStoragePolicyStat) HasStorageClasses() bool`

HasStorageClasses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


