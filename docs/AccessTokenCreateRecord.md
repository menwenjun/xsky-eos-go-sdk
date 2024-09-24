# AccessTokenCreateRecord

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
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessTokenCreateRecord

`func NewAccessTokenCreateRecord() *AccessTokenCreateRecord`

NewAccessTokenCreateRecord instantiates a new AccessTokenCreateRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenCreateRecordWithDefaults

`func NewAccessTokenCreateRecordWithDefaults() *AccessTokenCreateRecord`

NewAccessTokenCreateRecordWithDefaults instantiates a new AccessTokenCreateRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *AccessTokenCreateRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AccessTokenCreateRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AccessTokenCreateRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AccessTokenCreateRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *AccessTokenCreateRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessTokenCreateRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessTokenCreateRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessTokenCreateRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AccessTokenCreateRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessTokenCreateRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessTokenCreateRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AccessTokenCreateRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccessTokenCreateRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessTokenCreateRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessTokenCreateRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessTokenCreateRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *AccessTokenCreateRecord) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccessTokenCreateRecord) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccessTokenCreateRecord) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AccessTokenCreateRecord) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUpdate

`func (o *AccessTokenCreateRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AccessTokenCreateRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AccessTokenCreateRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AccessTokenCreateRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsed

`func (o *AccessTokenCreateRecord) GetUsed() bool`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *AccessTokenCreateRecord) GetUsedOk() (*bool, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *AccessTokenCreateRecord) SetUsed(v bool)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *AccessTokenCreateRecord) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUser

`func (o *AccessTokenCreateRecord) GetUser() UserNestview`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AccessTokenCreateRecord) GetUserOk() (*UserNestview, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AccessTokenCreateRecord) SetUser(v UserNestview)`

SetUser sets User field to given value.

### HasUser

`func (o *AccessTokenCreateRecord) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUuid

`func (o *AccessTokenCreateRecord) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AccessTokenCreateRecord) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AccessTokenCreateRecord) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AccessTokenCreateRecord) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


