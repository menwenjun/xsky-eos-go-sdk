# TokenRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTokenRecord

`func NewTokenRecord() *TokenRecord`

NewTokenRecord instantiates a new TokenRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRecordWithDefaults

`func NewTokenRecordWithDefaults() *TokenRecord`

NewTokenRecordWithDefaults instantiates a new TokenRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *TokenRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *TokenRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *TokenRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *TokenRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetExpires

`func (o *TokenRecord) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenRecord) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenRecord) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TokenRecord) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetUser

`func (o *TokenRecord) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TokenRecord) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TokenRecord) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *TokenRecord) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUuid

`func (o *TokenRecord) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TokenRecord) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TokenRecord) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TokenRecord) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetValid

`func (o *TokenRecord) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *TokenRecord) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *TokenRecord) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *TokenRecord) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetRoles

`func (o *TokenRecord) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *TokenRecord) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *TokenRecord) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *TokenRecord) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


