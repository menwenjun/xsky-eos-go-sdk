# TokenCreateReqAuthIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to [**AuthPasswordReq**](AuthPasswordReq.md) |  | [optional] 
**Token** | Pointer to [**AuthTokenReq**](AuthTokenReq.md) |  | [optional] 

## Methods

### NewTokenCreateReqAuthIdentity

`func NewTokenCreateReqAuthIdentity() *TokenCreateReqAuthIdentity`

NewTokenCreateReqAuthIdentity instantiates a new TokenCreateReqAuthIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreateReqAuthIdentityWithDefaults

`func NewTokenCreateReqAuthIdentityWithDefaults() *TokenCreateReqAuthIdentity`

NewTokenCreateReqAuthIdentityWithDefaults instantiates a new TokenCreateReqAuthIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *TokenCreateReqAuthIdentity) GetPassword() AuthPasswordReq`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TokenCreateReqAuthIdentity) GetPasswordOk() (*AuthPasswordReq, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TokenCreateReqAuthIdentity) SetPassword(v AuthPasswordReq)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TokenCreateReqAuthIdentity) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *TokenCreateReqAuthIdentity) GetToken() AuthTokenReq`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenCreateReqAuthIdentity) GetTokenOk() (*AuthTokenReq, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenCreateReqAuthIdentity) SetToken(v AuthTokenReq)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenCreateReqAuthIdentity) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


