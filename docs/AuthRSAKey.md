# AuthRSAKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthRSAKey

`func NewAuthRSAKey() *AuthRSAKey`

NewAuthRSAKey instantiates a new AuthRSAKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRSAKeyWithDefaults

`func NewAuthRSAKeyWithDefaults() *AuthRSAKey`

NewAuthRSAKeyWithDefaults instantiates a new AuthRSAKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *AuthRSAKey) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *AuthRSAKey) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *AuthRSAKey) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *AuthRSAKey) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetId

`func (o *AuthRSAKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthRSAKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthRSAKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthRSAKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublicKey

`func (o *AuthRSAKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *AuthRSAKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *AuthRSAKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *AuthRSAKey) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


