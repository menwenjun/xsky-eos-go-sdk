# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** | time of creating user | [optional] 
**Email** | Pointer to **string** | email of user | [optional] 
**Enabled** | Pointer to **bool** | enable the user or not | [optional] 
**ExternalId** | Pointer to **string** | external id of user | [optional] 
**ForcePasswordChange** | Pointer to **bool** | force change password on first login | [optional] 
**Id** | Pointer to **int64** | id of user | [optional] 
**IdentityPlatform** | Pointer to [**IdentityPlatformNestview**](IdentityPlatformNestview.md) |  | [optional] 
**Name** | Pointer to **string** | name of user | [optional] 
**PasswordLastUpdate** | Pointer to **time.Time** | time of last password update | [optional] 
**Roles** | Pointer to **[]string** | roles of user | [optional] 
**SkipGuide** | Pointer to **bool** | skip Guide | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *User) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *User) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *User) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *User) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *User) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *User) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *User) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *User) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalId

`func (o *User) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *User) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *User) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *User) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetForcePasswordChange

`func (o *User) GetForcePasswordChange() bool`

GetForcePasswordChange returns the ForcePasswordChange field if non-nil, zero value otherwise.

### GetForcePasswordChangeOk

`func (o *User) GetForcePasswordChangeOk() (*bool, bool)`

GetForcePasswordChangeOk returns a tuple with the ForcePasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePasswordChange

`func (o *User) SetForcePasswordChange(v bool)`

SetForcePasswordChange sets ForcePasswordChange field to given value.

### HasForcePasswordChange

`func (o *User) HasForcePasswordChange() bool`

HasForcePasswordChange returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityPlatform

`func (o *User) GetIdentityPlatform() IdentityPlatformNestview`

GetIdentityPlatform returns the IdentityPlatform field if non-nil, zero value otherwise.

### GetIdentityPlatformOk

`func (o *User) GetIdentityPlatformOk() (*IdentityPlatformNestview, bool)`

GetIdentityPlatformOk returns a tuple with the IdentityPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityPlatform

`func (o *User) SetIdentityPlatform(v IdentityPlatformNestview)`

SetIdentityPlatform sets IdentityPlatform field to given value.

### HasIdentityPlatform

`func (o *User) HasIdentityPlatform() bool`

HasIdentityPlatform returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPasswordLastUpdate

`func (o *User) GetPasswordLastUpdate() time.Time`

GetPasswordLastUpdate returns the PasswordLastUpdate field if non-nil, zero value otherwise.

### GetPasswordLastUpdateOk

`func (o *User) GetPasswordLastUpdateOk() (*time.Time, bool)`

GetPasswordLastUpdateOk returns a tuple with the PasswordLastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastUpdate

`func (o *User) SetPasswordLastUpdate(v time.Time)`

SetPasswordLastUpdate sets PasswordLastUpdate field to given value.

### HasPasswordLastUpdate

`func (o *User) HasPasswordLastUpdate() bool`

HasPasswordLastUpdate returns a boolean if a field has been set.

### GetRoles

`func (o *User) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *User) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *User) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSkipGuide

`func (o *User) GetSkipGuide() bool`

GetSkipGuide returns the SkipGuide field if non-nil, zero value otherwise.

### GetSkipGuideOk

`func (o *User) GetSkipGuideOk() (*bool, bool)`

GetSkipGuideOk returns a tuple with the SkipGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipGuide

`func (o *User) SetSkipGuide(v bool)`

SetSkipGuide sets SkipGuide field to given value.

### HasSkipGuide

`func (o *User) HasSkipGuide() bool`

HasSkipGuide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


