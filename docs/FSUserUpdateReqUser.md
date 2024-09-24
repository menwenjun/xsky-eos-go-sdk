# FSUserUpdateReqUser

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
**Email** | Pointer to **string** | email of user | [optional] 
**Name** | Pointer to **string** | name of user (deprecated, fs user name cannot be changed) | [optional] 
**Password** | Pointer to **string** | password of user | [optional] 

## Methods

### NewFSUserUpdateReqUser

`func NewFSUserUpdateReqUser() *FSUserUpdateReqUser`

NewFSUserUpdateReqUser instantiates a new FSUserUpdateReqUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSUserUpdateReqUserWithDefaults

`func NewFSUserUpdateReqUserWithDefaults() *FSUserUpdateReqUser`

NewFSUserUpdateReqUserWithDefaults instantiates a new FSUserUpdateReqUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketLimit

`func (o *FSUserUpdateReqUser) GetBucketLimit() int64`

GetBucketLimit returns the BucketLimit field if non-nil, zero value otherwise.

### GetBucketLimitOk

`func (o *FSUserUpdateReqUser) GetBucketLimitOk() (*int64, bool)`

GetBucketLimitOk returns a tuple with the BucketLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketLimit

`func (o *FSUserUpdateReqUser) SetBucketLimit(v int64)`

SetBucketLimit sets BucketLimit field to given value.

### HasBucketLimit

`func (o *FSUserUpdateReqUser) HasBucketLimit() bool`

HasBucketLimit returns a boolean if a field has been set.

### GetBucketPath

`func (o *FSUserUpdateReqUser) GetBucketPath() string`

GetBucketPath returns the BucketPath field if non-nil, zero value otherwise.

### GetBucketPathOk

`func (o *FSUserUpdateReqUser) GetBucketPathOk() (*string, bool)`

GetBucketPathOk returns a tuple with the BucketPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPath

`func (o *FSUserUpdateReqUser) SetBucketPath(v string)`

SetBucketPath sets BucketPath field to given value.

### HasBucketPath

`func (o *FSUserUpdateReqUser) HasBucketPath() bool`

HasBucketPath returns a boolean if a field has been set.

### GetBucketPermission

`func (o *FSUserUpdateReqUser) GetBucketPermission() string`

GetBucketPermission returns the BucketPermission field if non-nil, zero value otherwise.

### GetBucketPermissionOk

`func (o *FSUserUpdateReqUser) GetBucketPermissionOk() (*string, bool)`

GetBucketPermissionOk returns a tuple with the BucketPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPermission

`func (o *FSUserUpdateReqUser) SetBucketPermission(v string)`

SetBucketPermission sets BucketPermission field to given value.

### HasBucketPermission

`func (o *FSUserUpdateReqUser) HasBucketPermission() bool`

HasBucketPermission returns a boolean if a field has been set.

### GetGatewayGroupId

`func (o *FSUserUpdateReqUser) GetGatewayGroupId() int64`

GetGatewayGroupId returns the GatewayGroupId field if non-nil, zero value otherwise.

### GetGatewayGroupIdOk

`func (o *FSUserUpdateReqUser) GetGatewayGroupIdOk() (*int64, bool)`

GetGatewayGroupIdOk returns a tuple with the GatewayGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayGroupId

`func (o *FSUserUpdateReqUser) SetGatewayGroupId(v int64)`

SetGatewayGroupId sets GatewayGroupId field to given value.

### HasGatewayGroupId

`func (o *FSUserUpdateReqUser) HasGatewayGroupId() bool`

HasGatewayGroupId returns a boolean if a field has been set.

### GetS3Enabled

`func (o *FSUserUpdateReqUser) GetS3Enabled() bool`

GetS3Enabled returns the S3Enabled field if non-nil, zero value otherwise.

### GetS3EnabledOk

`func (o *FSUserUpdateReqUser) GetS3EnabledOk() (*bool, bool)`

GetS3EnabledOk returns a tuple with the S3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Enabled

`func (o *FSUserUpdateReqUser) SetS3Enabled(v bool)`

SetS3Enabled sets S3Enabled field to given value.

### HasS3Enabled

`func (o *FSUserUpdateReqUser) HasS3Enabled() bool`

HasS3Enabled returns a boolean if a field has been set.

### GetS3Keys

`func (o *FSUserUpdateReqUser) GetS3Keys() []S3Key`

GetS3Keys returns the S3Keys field if non-nil, zero value otherwise.

### GetS3KeysOk

`func (o *FSUserUpdateReqUser) GetS3KeysOk() (*[]S3Key, bool)`

GetS3KeysOk returns a tuple with the S3Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Keys

`func (o *FSUserUpdateReqUser) SetS3Keys(v []S3Key)`

SetS3Keys sets S3Keys field to given value.

### HasS3Keys

`func (o *FSUserUpdateReqUser) HasS3Keys() bool`

HasS3Keys returns a boolean if a field has been set.

### GetDescription

`func (o *FSUserUpdateReqUser) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FSUserUpdateReqUser) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FSUserUpdateReqUser) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FSUserUpdateReqUser) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *FSUserUpdateReqUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FSUserUpdateReqUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FSUserUpdateReqUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FSUserUpdateReqUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *FSUserUpdateReqUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FSUserUpdateReqUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FSUserUpdateReqUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FSUserUpdateReqUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *FSUserUpdateReqUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FSUserUpdateReqUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FSUserUpdateReqUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FSUserUpdateReqUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


