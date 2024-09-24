# FSUserCreateReqUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketLimit** | Pointer to **int64** |  | [optional] 
**BucketPath** | Pointer to **string** |  | [optional] 
**BucketPermission** | Pointer to **string** |  | [optional] 
**GatewayGroupId** | Pointer to **int64** |  | [optional] 
**S3Enabled** | Pointer to **bool** |  | [optional] 
**S3Keys** | Pointer to [**[]S3Key**](S3Key.md) |  | [optional] 
**Description** | Pointer to **string** | description of file storage user | [optional] 
**Email** | Pointer to **string** | email of file storage user | [optional] 
**Name** | Pointer to **string** | name of file storage user | [optional] 
**Password** | Pointer to **string** | password of file storage user | [optional] 
**PrimaryGroupId** | Pointer to **int64** | primary group of user | [optional] 
**UserGroupIds** | Pointer to **[]int64** | secondary group of user | [optional] 
**UserId** | Pointer to **int64** | uid of file storage user | [optional] 

## Methods

### NewFSUserCreateReqUser

`func NewFSUserCreateReqUser() *FSUserCreateReqUser`

NewFSUserCreateReqUser instantiates a new FSUserCreateReqUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSUserCreateReqUserWithDefaults

`func NewFSUserCreateReqUserWithDefaults() *FSUserCreateReqUser`

NewFSUserCreateReqUserWithDefaults instantiates a new FSUserCreateReqUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketLimit

`func (o *FSUserCreateReqUser) GetBucketLimit() int64`

GetBucketLimit returns the BucketLimit field if non-nil, zero value otherwise.

### GetBucketLimitOk

`func (o *FSUserCreateReqUser) GetBucketLimitOk() (*int64, bool)`

GetBucketLimitOk returns a tuple with the BucketLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketLimit

`func (o *FSUserCreateReqUser) SetBucketLimit(v int64)`

SetBucketLimit sets BucketLimit field to given value.

### HasBucketLimit

`func (o *FSUserCreateReqUser) HasBucketLimit() bool`

HasBucketLimit returns a boolean if a field has been set.

### GetBucketPath

`func (o *FSUserCreateReqUser) GetBucketPath() string`

GetBucketPath returns the BucketPath field if non-nil, zero value otherwise.

### GetBucketPathOk

`func (o *FSUserCreateReqUser) GetBucketPathOk() (*string, bool)`

GetBucketPathOk returns a tuple with the BucketPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPath

`func (o *FSUserCreateReqUser) SetBucketPath(v string)`

SetBucketPath sets BucketPath field to given value.

### HasBucketPath

`func (o *FSUserCreateReqUser) HasBucketPath() bool`

HasBucketPath returns a boolean if a field has been set.

### GetBucketPermission

`func (o *FSUserCreateReqUser) GetBucketPermission() string`

GetBucketPermission returns the BucketPermission field if non-nil, zero value otherwise.

### GetBucketPermissionOk

`func (o *FSUserCreateReqUser) GetBucketPermissionOk() (*string, bool)`

GetBucketPermissionOk returns a tuple with the BucketPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPermission

`func (o *FSUserCreateReqUser) SetBucketPermission(v string)`

SetBucketPermission sets BucketPermission field to given value.

### HasBucketPermission

`func (o *FSUserCreateReqUser) HasBucketPermission() bool`

HasBucketPermission returns a boolean if a field has been set.

### GetGatewayGroupId

`func (o *FSUserCreateReqUser) GetGatewayGroupId() int64`

GetGatewayGroupId returns the GatewayGroupId field if non-nil, zero value otherwise.

### GetGatewayGroupIdOk

`func (o *FSUserCreateReqUser) GetGatewayGroupIdOk() (*int64, bool)`

GetGatewayGroupIdOk returns a tuple with the GatewayGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayGroupId

`func (o *FSUserCreateReqUser) SetGatewayGroupId(v int64)`

SetGatewayGroupId sets GatewayGroupId field to given value.

### HasGatewayGroupId

`func (o *FSUserCreateReqUser) HasGatewayGroupId() bool`

HasGatewayGroupId returns a boolean if a field has been set.

### GetS3Enabled

`func (o *FSUserCreateReqUser) GetS3Enabled() bool`

GetS3Enabled returns the S3Enabled field if non-nil, zero value otherwise.

### GetS3EnabledOk

`func (o *FSUserCreateReqUser) GetS3EnabledOk() (*bool, bool)`

GetS3EnabledOk returns a tuple with the S3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Enabled

`func (o *FSUserCreateReqUser) SetS3Enabled(v bool)`

SetS3Enabled sets S3Enabled field to given value.

### HasS3Enabled

`func (o *FSUserCreateReqUser) HasS3Enabled() bool`

HasS3Enabled returns a boolean if a field has been set.

### GetS3Keys

`func (o *FSUserCreateReqUser) GetS3Keys() []S3Key`

GetS3Keys returns the S3Keys field if non-nil, zero value otherwise.

### GetS3KeysOk

`func (o *FSUserCreateReqUser) GetS3KeysOk() (*[]S3Key, bool)`

GetS3KeysOk returns a tuple with the S3Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Keys

`func (o *FSUserCreateReqUser) SetS3Keys(v []S3Key)`

SetS3Keys sets S3Keys field to given value.

### HasS3Keys

`func (o *FSUserCreateReqUser) HasS3Keys() bool`

HasS3Keys returns a boolean if a field has been set.

### GetDescription

`func (o *FSUserCreateReqUser) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FSUserCreateReqUser) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FSUserCreateReqUser) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FSUserCreateReqUser) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *FSUserCreateReqUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FSUserCreateReqUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FSUserCreateReqUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FSUserCreateReqUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *FSUserCreateReqUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FSUserCreateReqUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FSUserCreateReqUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FSUserCreateReqUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *FSUserCreateReqUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FSUserCreateReqUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FSUserCreateReqUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FSUserCreateReqUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPrimaryGroupId

`func (o *FSUserCreateReqUser) GetPrimaryGroupId() int64`

GetPrimaryGroupId returns the PrimaryGroupId field if non-nil, zero value otherwise.

### GetPrimaryGroupIdOk

`func (o *FSUserCreateReqUser) GetPrimaryGroupIdOk() (*int64, bool)`

GetPrimaryGroupIdOk returns a tuple with the PrimaryGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGroupId

`func (o *FSUserCreateReqUser) SetPrimaryGroupId(v int64)`

SetPrimaryGroupId sets PrimaryGroupId field to given value.

### HasPrimaryGroupId

`func (o *FSUserCreateReqUser) HasPrimaryGroupId() bool`

HasPrimaryGroupId returns a boolean if a field has been set.

### GetUserGroupIds

`func (o *FSUserCreateReqUser) GetUserGroupIds() []int64`

GetUserGroupIds returns the UserGroupIds field if non-nil, zero value otherwise.

### GetUserGroupIdsOk

`func (o *FSUserCreateReqUser) GetUserGroupIdsOk() (*[]int64, bool)`

GetUserGroupIdsOk returns a tuple with the UserGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupIds

`func (o *FSUserCreateReqUser) SetUserGroupIds(v []int64)`

SetUserGroupIds sets UserGroupIds field to given value.

### HasUserGroupIds

`func (o *FSUserCreateReqUser) HasUserGroupIds() bool`

HasUserGroupIds returns a boolean if a field has been set.

### GetUserId

`func (o *FSUserCreateReqUser) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FSUserCreateReqUser) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FSUserCreateReqUser) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FSUserCreateReqUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


