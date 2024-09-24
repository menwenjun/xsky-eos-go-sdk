# DfsS3BucketCreateReqBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllUserPermission** | Pointer to **string** |  | [optional] 
**AuthUserPermission** | Pointer to **string** |  | [optional] 
**DataVerify** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** | description of bucket | [optional] 
**EnableEtag** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**OwnerId** | **int64** |  | 
**OwnerPermission** | Pointer to **string** |  | [optional] 
**Path** | **string** |  | 
**RootfsId** | **int64** |  | 

## Methods

### NewDfsS3BucketCreateReqBucket

`func NewDfsS3BucketCreateReqBucket(name string, ownerId int64, path string, rootfsId int64, ) *DfsS3BucketCreateReqBucket`

NewDfsS3BucketCreateReqBucket instantiates a new DfsS3BucketCreateReqBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsS3BucketCreateReqBucketWithDefaults

`func NewDfsS3BucketCreateReqBucketWithDefaults() *DfsS3BucketCreateReqBucket`

NewDfsS3BucketCreateReqBucketWithDefaults instantiates a new DfsS3BucketCreateReqBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllUserPermission

`func (o *DfsS3BucketCreateReqBucket) GetAllUserPermission() string`

GetAllUserPermission returns the AllUserPermission field if non-nil, zero value otherwise.

### GetAllUserPermissionOk

`func (o *DfsS3BucketCreateReqBucket) GetAllUserPermissionOk() (*string, bool)`

GetAllUserPermissionOk returns a tuple with the AllUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUserPermission

`func (o *DfsS3BucketCreateReqBucket) SetAllUserPermission(v string)`

SetAllUserPermission sets AllUserPermission field to given value.

### HasAllUserPermission

`func (o *DfsS3BucketCreateReqBucket) HasAllUserPermission() bool`

HasAllUserPermission returns a boolean if a field has been set.

### GetAuthUserPermission

`func (o *DfsS3BucketCreateReqBucket) GetAuthUserPermission() string`

GetAuthUserPermission returns the AuthUserPermission field if non-nil, zero value otherwise.

### GetAuthUserPermissionOk

`func (o *DfsS3BucketCreateReqBucket) GetAuthUserPermissionOk() (*string, bool)`

GetAuthUserPermissionOk returns a tuple with the AuthUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserPermission

`func (o *DfsS3BucketCreateReqBucket) SetAuthUserPermission(v string)`

SetAuthUserPermission sets AuthUserPermission field to given value.

### HasAuthUserPermission

`func (o *DfsS3BucketCreateReqBucket) HasAuthUserPermission() bool`

HasAuthUserPermission returns a boolean if a field has been set.

### GetDataVerify

`func (o *DfsS3BucketCreateReqBucket) GetDataVerify() bool`

GetDataVerify returns the DataVerify field if non-nil, zero value otherwise.

### GetDataVerifyOk

`func (o *DfsS3BucketCreateReqBucket) GetDataVerifyOk() (*bool, bool)`

GetDataVerifyOk returns a tuple with the DataVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVerify

`func (o *DfsS3BucketCreateReqBucket) SetDataVerify(v bool)`

SetDataVerify sets DataVerify field to given value.

### HasDataVerify

`func (o *DfsS3BucketCreateReqBucket) HasDataVerify() bool`

HasDataVerify returns a boolean if a field has been set.

### GetDescription

`func (o *DfsS3BucketCreateReqBucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsS3BucketCreateReqBucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsS3BucketCreateReqBucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsS3BucketCreateReqBucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableEtag

`func (o *DfsS3BucketCreateReqBucket) GetEnableEtag() bool`

GetEnableEtag returns the EnableEtag field if non-nil, zero value otherwise.

### GetEnableEtagOk

`func (o *DfsS3BucketCreateReqBucket) GetEnableEtagOk() (*bool, bool)`

GetEnableEtagOk returns a tuple with the EnableEtag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEtag

`func (o *DfsS3BucketCreateReqBucket) SetEnableEtag(v bool)`

SetEnableEtag sets EnableEtag field to given value.

### HasEnableEtag

`func (o *DfsS3BucketCreateReqBucket) HasEnableEtag() bool`

HasEnableEtag returns a boolean if a field has been set.

### GetName

`func (o *DfsS3BucketCreateReqBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsS3BucketCreateReqBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsS3BucketCreateReqBucket) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *DfsS3BucketCreateReqBucket) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DfsS3BucketCreateReqBucket) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DfsS3BucketCreateReqBucket) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerPermission

`func (o *DfsS3BucketCreateReqBucket) GetOwnerPermission() string`

GetOwnerPermission returns the OwnerPermission field if non-nil, zero value otherwise.

### GetOwnerPermissionOk

`func (o *DfsS3BucketCreateReqBucket) GetOwnerPermissionOk() (*string, bool)`

GetOwnerPermissionOk returns a tuple with the OwnerPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPermission

`func (o *DfsS3BucketCreateReqBucket) SetOwnerPermission(v string)`

SetOwnerPermission sets OwnerPermission field to given value.

### HasOwnerPermission

`func (o *DfsS3BucketCreateReqBucket) HasOwnerPermission() bool`

HasOwnerPermission returns a boolean if a field has been set.

### GetPath

`func (o *DfsS3BucketCreateReqBucket) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsS3BucketCreateReqBucket) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsS3BucketCreateReqBucket) SetPath(v string)`

SetPath sets Path field to given value.


### GetRootfsId

`func (o *DfsS3BucketCreateReqBucket) GetRootfsId() int64`

GetRootfsId returns the RootfsId field if non-nil, zero value otherwise.

### GetRootfsIdOk

`func (o *DfsS3BucketCreateReqBucket) GetRootfsIdOk() (*int64, bool)`

GetRootfsIdOk returns a tuple with the RootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfsId

`func (o *DfsS3BucketCreateReqBucket) SetRootfsId(v int64)`

SetRootfsId sets RootfsId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


