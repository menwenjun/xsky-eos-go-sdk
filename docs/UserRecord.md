# UserRecord

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

### NewUserRecord

`func NewUserRecord() *UserRecord`

NewUserRecord instantiates a new UserRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRecordWithDefaults

`func NewUserRecordWithDefaults() *UserRecord`

NewUserRecordWithDefaults instantiates a new UserRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *UserRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *UserRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *UserRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *UserRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEmail

`func (o *UserRecord) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRecord) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRecord) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserRecord) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *UserRecord) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserRecord) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserRecord) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserRecord) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalId

`func (o *UserRecord) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UserRecord) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UserRecord) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UserRecord) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetForcePasswordChange

`func (o *UserRecord) GetForcePasswordChange() bool`

GetForcePasswordChange returns the ForcePasswordChange field if non-nil, zero value otherwise.

### GetForcePasswordChangeOk

`func (o *UserRecord) GetForcePasswordChangeOk() (*bool, bool)`

GetForcePasswordChangeOk returns a tuple with the ForcePasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePasswordChange

`func (o *UserRecord) SetForcePasswordChange(v bool)`

SetForcePasswordChange sets ForcePasswordChange field to given value.

### HasForcePasswordChange

`func (o *UserRecord) HasForcePasswordChange() bool`

HasForcePasswordChange returns a boolean if a field has been set.

### GetId

`func (o *UserRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityPlatform

`func (o *UserRecord) GetIdentityPlatform() IdentityPlatformNestview`

GetIdentityPlatform returns the IdentityPlatform field if non-nil, zero value otherwise.

### GetIdentityPlatformOk

`func (o *UserRecord) GetIdentityPlatformOk() (*IdentityPlatformNestview, bool)`

GetIdentityPlatformOk returns a tuple with the IdentityPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityPlatform

`func (o *UserRecord) SetIdentityPlatform(v IdentityPlatformNestview)`

SetIdentityPlatform sets IdentityPlatform field to given value.

### HasIdentityPlatform

`func (o *UserRecord) HasIdentityPlatform() bool`

HasIdentityPlatform returns a boolean if a field has been set.

### GetName

`func (o *UserRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPasswordLastUpdate

`func (o *UserRecord) GetPasswordLastUpdate() time.Time`

GetPasswordLastUpdate returns the PasswordLastUpdate field if non-nil, zero value otherwise.

### GetPasswordLastUpdateOk

`func (o *UserRecord) GetPasswordLastUpdateOk() (*time.Time, bool)`

GetPasswordLastUpdateOk returns a tuple with the PasswordLastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastUpdate

`func (o *UserRecord) SetPasswordLastUpdate(v time.Time)`

SetPasswordLastUpdate sets PasswordLastUpdate field to given value.

### HasPasswordLastUpdate

`func (o *UserRecord) HasPasswordLastUpdate() bool`

HasPasswordLastUpdate returns a boolean if a field has been set.

### GetRoles

`func (o *UserRecord) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserRecord) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserRecord) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserRecord) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSkipGuide

`func (o *UserRecord) GetSkipGuide() bool`

GetSkipGuide returns the SkipGuide field if non-nil, zero value otherwise.

### GetSkipGuideOk

`func (o *UserRecord) GetSkipGuideOk() (*bool, bool)`

GetSkipGuideOk returns a tuple with the SkipGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipGuide

`func (o *UserRecord) SetSkipGuide(v bool)`

SetSkipGuide sets SkipGuide field to given value.

### HasSkipGuide

`func (o *UserRecord) HasSkipGuide() bool`

HasSkipGuide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


