# DfsS3BucketUpdateReqBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllUserPermission** | Pointer to **string** |  | [optional] 
**AuthUserPermission** | Pointer to **string** |  | [optional] 
**DataVerify** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** | description of bucket | [optional] 
**EnableEtag** | Pointer to **bool** |  | [optional] 
**OwnerId** | Pointer to **int64** |  | [optional] 
**OwnerPermission** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsS3BucketUpdateReqBucket

`func NewDfsS3BucketUpdateReqBucket() *DfsS3BucketUpdateReqBucket`

NewDfsS3BucketUpdateReqBucket instantiates a new DfsS3BucketUpdateReqBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsS3BucketUpdateReqBucketWithDefaults

`func NewDfsS3BucketUpdateReqBucketWithDefaults() *DfsS3BucketUpdateReqBucket`

NewDfsS3BucketUpdateReqBucketWithDefaults instantiates a new DfsS3BucketUpdateReqBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllUserPermission

`func (o *DfsS3BucketUpdateReqBucket) GetAllUserPermission() string`

GetAllUserPermission returns the AllUserPermission field if non-nil, zero value otherwise.

### GetAllUserPermissionOk

`func (o *DfsS3BucketUpdateReqBucket) GetAllUserPermissionOk() (*string, bool)`

GetAllUserPermissionOk returns a tuple with the AllUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUserPermission

`func (o *DfsS3BucketUpdateReqBucket) SetAllUserPermission(v string)`

SetAllUserPermission sets AllUserPermission field to given value.

### HasAllUserPermission

`func (o *DfsS3BucketUpdateReqBucket) HasAllUserPermission() bool`

HasAllUserPermission returns a boolean if a field has been set.

### GetAuthUserPermission

`func (o *DfsS3BucketUpdateReqBucket) GetAuthUserPermission() string`

GetAuthUserPermission returns the AuthUserPermission field if non-nil, zero value otherwise.

### GetAuthUserPermissionOk

`func (o *DfsS3BucketUpdateReqBucket) GetAuthUserPermissionOk() (*string, bool)`

GetAuthUserPermissionOk returns a tuple with the AuthUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserPermission

`func (o *DfsS3BucketUpdateReqBucket) SetAuthUserPermission(v string)`

SetAuthUserPermission sets AuthUserPermission field to given value.

### HasAuthUserPermission

`func (o *DfsS3BucketUpdateReqBucket) HasAuthUserPermission() bool`

HasAuthUserPermission returns a boolean if a field has been set.

### GetDataVerify

`func (o *DfsS3BucketUpdateReqBucket) GetDataVerify() bool`

GetDataVerify returns the DataVerify field if non-nil, zero value otherwise.

### GetDataVerifyOk

`func (o *DfsS3BucketUpdateReqBucket) GetDataVerifyOk() (*bool, bool)`

GetDataVerifyOk returns a tuple with the DataVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVerify

`func (o *DfsS3BucketUpdateReqBucket) SetDataVerify(v bool)`

SetDataVerify sets DataVerify field to given value.

### HasDataVerify

`func (o *DfsS3BucketUpdateReqBucket) HasDataVerify() bool`

HasDataVerify returns a boolean if a field has been set.

### GetDescription

`func (o *DfsS3BucketUpdateReqBucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsS3BucketUpdateReqBucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsS3BucketUpdateReqBucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsS3BucketUpdateReqBucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableEtag

`func (o *DfsS3BucketUpdateReqBucket) GetEnableEtag() bool`

GetEnableEtag returns the EnableEtag field if non-nil, zero value otherwise.

### GetEnableEtagOk

`func (o *DfsS3BucketUpdateReqBucket) GetEnableEtagOk() (*bool, bool)`

GetEnableEtagOk returns a tuple with the EnableEtag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEtag

`func (o *DfsS3BucketUpdateReqBucket) SetEnableEtag(v bool)`

SetEnableEtag sets EnableEtag field to given value.

### HasEnableEtag

`func (o *DfsS3BucketUpdateReqBucket) HasEnableEtag() bool`

HasEnableEtag returns a boolean if a field has been set.

### GetOwnerId

`func (o *DfsS3BucketUpdateReqBucket) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DfsS3BucketUpdateReqBucket) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DfsS3BucketUpdateReqBucket) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *DfsS3BucketUpdateReqBucket) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerPermission

`func (o *DfsS3BucketUpdateReqBucket) GetOwnerPermission() string`

GetOwnerPermission returns the OwnerPermission field if non-nil, zero value otherwise.

### GetOwnerPermissionOk

`func (o *DfsS3BucketUpdateReqBucket) GetOwnerPermissionOk() (*string, bool)`

GetOwnerPermissionOk returns a tuple with the OwnerPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPermission

`func (o *DfsS3BucketUpdateReqBucket) SetOwnerPermission(v string)`

SetOwnerPermission sets OwnerPermission field to given value.

### HasOwnerPermission

`func (o *DfsS3BucketUpdateReqBucket) HasOwnerPermission() bool`

HasOwnerPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


