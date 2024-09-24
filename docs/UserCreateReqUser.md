# UserCreateReqUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | email of user | 
**Enabled** | **bool** | enable or disable the user | 
**EncryptedPassword** | Pointer to **string** | encrypted password for auth | [optional] 
**ForcePasswordChange** | Pointer to **bool** | force to change password on first login | [optional] 
**Name** | **string** | name of user | 
**Password** | **string** | password of user | 
**Roles** | Pointer to **[]string** | roles of user | [optional] 
**RsaKeyId** | Pointer to **string** | rsa key id | [optional] 

## Methods

### NewUserCreateReqUser

`func NewUserCreateReqUser(email string, enabled bool, name string, password string, ) *UserCreateReqUser`

NewUserCreateReqUser instantiates a new UserCreateReqUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateReqUserWithDefaults

`func NewUserCreateReqUserWithDefaults() *UserCreateReqUser`

NewUserCreateReqUserWithDefaults instantiates a new UserCreateReqUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserCreateReqUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreateReqUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreateReqUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnabled

`func (o *UserCreateReqUser) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserCreateReqUser) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserCreateReqUser) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEncryptedPassword

`func (o *UserCreateReqUser) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *UserCreateReqUser) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *UserCreateReqUser) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *UserCreateReqUser) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetForcePasswordChange

`func (o *UserCreateReqUser) GetForcePasswordChange() bool`

GetForcePasswordChange returns the ForcePasswordChange field if non-nil, zero value otherwise.

### GetForcePasswordChangeOk

`func (o *UserCreateReqUser) GetForcePasswordChangeOk() (*bool, bool)`

GetForcePasswordChangeOk returns a tuple with the ForcePasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePasswordChange

`func (o *UserCreateReqUser) SetForcePasswordChange(v bool)`

SetForcePasswordChange sets ForcePasswordChange field to given value.

### HasForcePasswordChange

`func (o *UserCreateReqUser) HasForcePasswordChange() bool`

HasForcePasswordChange returns a boolean if a field has been set.

### GetName

`func (o *UserCreateReqUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserCreateReqUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserCreateReqUser) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *UserCreateReqUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreateReqUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreateReqUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRoles

`func (o *UserCreateReqUser) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserCreateReqUser) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserCreateReqUser) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserCreateReqUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRsaKeyId

`func (o *UserCreateReqUser) GetRsaKeyId() string`

GetRsaKeyId returns the RsaKeyId field if non-nil, zero value otherwise.

### GetRsaKeyIdOk

`func (o *UserCreateReqUser) GetRsaKeyIdOk() (*string, bool)`

GetRsaKeyIdOk returns a tuple with the RsaKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaKeyId

`func (o *UserCreateReqUser) SetRsaKeyId(v string)`

SetRsaKeyId sets RsaKeyId field to given value.

### HasRsaKeyId

`func (o *UserCreateReqUser) HasRsaKeyId() bool`

HasRsaKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


