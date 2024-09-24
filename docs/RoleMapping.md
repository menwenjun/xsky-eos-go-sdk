# RoleMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** | time of creating identity platform | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** | id of role mapping | [optional] 
**IdentityPlatform** | Pointer to [**IdentityPlatformNestview**](IdentityPlatformNestview.md) |  | [optional] 
**Role** | Pointer to **string** | mapped role of sds platform | [optional] 
**Update** | Pointer to **time.Time** | time of updating identity platform | [optional] 
**Value** | Pointer to **string** | roles of external identity platform | [optional] 

## Methods

### NewRoleMapping

`func NewRoleMapping() *RoleMapping`

NewRoleMapping instantiates a new RoleMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMappingWithDefaults

`func NewRoleMappingWithDefaults() *RoleMapping`

NewRoleMappingWithDefaults instantiates a new RoleMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *RoleMapping) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *RoleMapping) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *RoleMapping) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *RoleMapping) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDefault

`func (o *RoleMapping) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RoleMapping) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RoleMapping) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RoleMapping) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetId

`func (o *RoleMapping) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleMapping) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleMapping) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RoleMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityPlatform

`func (o *RoleMapping) GetIdentityPlatform() IdentityPlatformNestview`

GetIdentityPlatform returns the IdentityPlatform field if non-nil, zero value otherwise.

### GetIdentityPlatformOk

`func (o *RoleMapping) GetIdentityPlatformOk() (*IdentityPlatformNestview, bool)`

GetIdentityPlatformOk returns a tuple with the IdentityPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityPlatform

`func (o *RoleMapping) SetIdentityPlatform(v IdentityPlatformNestview)`

SetIdentityPlatform sets IdentityPlatform field to given value.

### HasIdentityPlatform

`func (o *RoleMapping) HasIdentityPlatform() bool`

HasIdentityPlatform returns a boolean if a field has been set.

### GetRole

`func (o *RoleMapping) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleMapping) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleMapping) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleMapping) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUpdate

`func (o *RoleMapping) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *RoleMapping) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *RoleMapping) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *RoleMapping) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetValue

`func (o *RoleMapping) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RoleMapping) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RoleMapping) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RoleMapping) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


