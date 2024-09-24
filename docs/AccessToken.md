# AccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Used** | Pointer to **bool** |  | [optional] 
**User** | Pointer to [**UserNestview**](UserNestview.md) |  | [optional] 

## Methods

### NewAccessToken

`func NewAccessToken() *AccessToken`

NewAccessToken instantiates a new AccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenWithDefaults

`func NewAccessTokenWithDefaults() *AccessToken`

NewAccessTokenWithDefaults instantiates a new AccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *AccessToken) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AccessToken) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AccessToken) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AccessToken) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *AccessToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AccessToken) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessToken) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessToken) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AccessToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccessToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *AccessToken) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccessToken) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccessToken) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AccessToken) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUpdate

`func (o *AccessToken) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AccessToken) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AccessToken) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AccessToken) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsed

`func (o *AccessToken) GetUsed() bool`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *AccessToken) GetUsedOk() (*bool, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *AccessToken) SetUsed(v bool)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *AccessToken) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUser

`func (o *AccessToken) GetUser() UserNestview`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AccessToken) GetUserOk() (*UserNestview, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AccessToken) SetUser(v UserNestview)`

SetUser sets User field to given value.

### HasUser

`func (o *AccessToken) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


