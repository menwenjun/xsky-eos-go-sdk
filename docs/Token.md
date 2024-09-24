# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *Token) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Token) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Token) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Token) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetExpires

`func (o *Token) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Token) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Token) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Token) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetUser

`func (o *Token) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Token) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Token) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Token) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUuid

`func (o *Token) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Token) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Token) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Token) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetValid

`func (o *Token) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *Token) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *Token) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *Token) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


