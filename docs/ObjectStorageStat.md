# ObjectStorageStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedObjects** | Pointer to **int64** |  | [optional] 
**AllocatedSize** | Pointer to **int64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**ExternalAllocatedObjects** | Pointer to **int64** |  | [optional] 
**ExternalAllocatedSize** | Pointer to **int64** |  | [optional] 
**LocalAllocatedObjects** | Pointer to **int64** |  | [optional] 
**LocalAllocatedSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewObjectStorageStat

`func NewObjectStorageStat() *ObjectStorageStat`

NewObjectStorageStat instantiates a new ObjectStorageStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageStatWithDefaults

`func NewObjectStorageStatWithDefaults() *ObjectStorageStat`

NewObjectStorageStatWithDefaults instantiates a new ObjectStorageStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedObjects

`func (o *ObjectStorageStat) GetAllocatedObjects() int64`

GetAllocatedObjects returns the AllocatedObjects field if non-nil, zero value otherwise.

### GetAllocatedObjectsOk

`func (o *ObjectStorageStat) GetAllocatedObjectsOk() (*int64, bool)`

GetAllocatedObjectsOk returns a tuple with the AllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedObjects

`func (o *ObjectStorageStat) SetAllocatedObjects(v int64)`

SetAllocatedObjects sets AllocatedObjects field to given value.

### HasAllocatedObjects

`func (o *ObjectStorageStat) HasAllocatedObjects() bool`

HasAllocatedObjects returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *ObjectStorageStat) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *ObjectStorageStat) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *ObjectStorageStat) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *ObjectStorageStat) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetExternalAllocatedObjects

`func (o *ObjectStorageStat) GetExternalAllocatedObjects() int64`

GetExternalAllocatedObjects returns the ExternalAllocatedObjects field if non-nil, zero value otherwise.

### GetExternalAllocatedObjectsOk

`func (o *ObjectStorageStat) GetExternalAllocatedObjectsOk() (*int64, bool)`

GetExternalAllocatedObjectsOk returns a tuple with the ExternalAllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAllocatedObjects

`func (o *ObjectStorageStat) SetExternalAllocatedObjects(v int64)`

SetExternalAllocatedObjects sets ExternalAllocatedObjects field to given value.

### HasExternalAllocatedObjects

`func (o *ObjectStorageStat) HasExternalAllocatedObjects() bool`

HasExternalAllocatedObjects returns a boolean if a field has been set.

### GetExternalAllocatedSize

`func (o *ObjectStorageStat) GetExternalAllocatedSize() int64`

GetExternalAllocatedSize returns the ExternalAllocatedSize field if non-nil, zero value otherwise.

### GetExternalAllocatedSizeOk

`func (o *ObjectStorageStat) GetExternalAllocatedSizeOk() (*int64, bool)`

GetExternalAllocatedSizeOk returns a tuple with the ExternalAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAllocatedSize

`func (o *ObjectStorageStat) SetExternalAllocatedSize(v int64)`

SetExternalAllocatedSize sets ExternalAllocatedSize field to given value.

### HasExternalAllocatedSize

`func (o *ObjectStorageStat) HasExternalAllocatedSize() bool`

HasExternalAllocatedSize returns a boolean if a field has been set.

### GetLocalAllocatedObjects

`func (o *ObjectStorageStat) GetLocalAllocatedObjects() int64`

GetLocalAllocatedObjects returns the LocalAllocatedObjects field if non-nil, zero value otherwise.

### GetLocalAllocatedObjectsOk

`func (o *ObjectStorageStat) GetLocalAllocatedObjectsOk() (*int64, bool)`

GetLocalAllocatedObjectsOk returns a tuple with the LocalAllocatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAllocatedObjects

`func (o *ObjectStorageStat) SetLocalAllocatedObjects(v int64)`

SetLocalAllocatedObjects sets LocalAllocatedObjects field to given value.

### HasLocalAllocatedObjects

`func (o *ObjectStorageStat) HasLocalAllocatedObjects() bool`

HasLocalAllocatedObjects returns a boolean if a field has been set.

### GetLocalAllocatedSize

`func (o *ObjectStorageStat) GetLocalAllocatedSize() int64`

GetLocalAllocatedSize returns the LocalAllocatedSize field if non-nil, zero value otherwise.

### GetLocalAllocatedSizeOk

`func (o *ObjectStorageStat) GetLocalAllocatedSizeOk() (*int64, bool)`

GetLocalAllocatedSizeOk returns a tuple with the LocalAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAllocatedSize

`func (o *ObjectStorageStat) SetLocalAllocatedSize(v int64)`

SetLocalAllocatedSize sets LocalAllocatedSize field to given value.

### HasLocalAllocatedSize

`func (o *ObjectStorageStat) HasLocalAllocatedSize() bool`

HasLocalAllocatedSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


