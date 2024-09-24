# UserUpdateReqUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | email of user | 
**Enabled** | **bool** | enable or disable the user | 
**Name** | **string** | name of user | 
**Roles** | Pointer to **[]string** | roles of user | [optional] 

## Methods

### NewUserUpdateReqUser

`func NewUserUpdateReqUser(email string, enabled bool, name string, ) *UserUpdateReqUser`

NewUserUpdateReqUser instantiates a new UserUpdateReqUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateReqUserWithDefaults

`func NewUserUpdateReqUserWithDefaults() *UserUpdateReqUser`

NewUserUpdateReqUserWithDefaults instantiates a new UserUpdateReqUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserUpdateReqUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserUpdateReqUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserUpdateReqUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnabled

`func (o *UserUpdateReqUser) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserUpdateReqUser) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserUpdateReqUser) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetName

`func (o *UserUpdateReqUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserUpdateReqUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserUpdateReqUser) SetName(v string)`

SetName sets Name field to given value.


### GetRoles

`func (o *UserUpdateReqUser) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserUpdateReqUser) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserUpdateReqUser) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserUpdateReqUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


