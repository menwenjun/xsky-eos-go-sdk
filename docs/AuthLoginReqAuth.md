# AuthLoginReqAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedPassword** | Pointer to **string** | encrypted password for auth | [optional] 
**Extra** | Pointer to [**LoginSSOExtra**](LoginSSOExtra.md) |  | [optional] 
**Name** | Pointer to **string** | user name or email for auth | [optional] 
**Password** | Pointer to **string** | password for auth | [optional] 
**Platform** | Pointer to **string** | uuid of the identity platform | [optional] 
**RsaKeyId** | Pointer to **string** | rsa key id | [optional] 
**Sign** | Pointer to **string** | signature of login request | [optional] 
**Timestamp** | Pointer to **string** | timestamp of platform token | [optional] 
**Token** | Pointer to **string** | token of the identity platform | [optional] 
**VerificationCode** | Pointer to **string** | verification code | [optional] 

## Methods

### NewAuthLoginReqAuth

`func NewAuthLoginReqAuth() *AuthLoginReqAuth`

NewAuthLoginReqAuth instantiates a new AuthLoginReqAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLoginReqAuthWithDefaults

`func NewAuthLoginReqAuthWithDefaults() *AuthLoginReqAuth`

NewAuthLoginReqAuthWithDefaults instantiates a new AuthLoginReqAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedPassword

`func (o *AuthLoginReqAuth) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *AuthLoginReqAuth) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *AuthLoginReqAuth) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *AuthLoginReqAuth) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetExtra

`func (o *AuthLoginReqAuth) GetExtra() LoginSSOExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *AuthLoginReqAuth) GetExtraOk() (*LoginSSOExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *AuthLoginReqAuth) SetExtra(v LoginSSOExtra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *AuthLoginReqAuth) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetName

`func (o *AuthLoginReqAuth) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthLoginReqAuth) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthLoginReqAuth) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthLoginReqAuth) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *AuthLoginReqAuth) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthLoginReqAuth) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthLoginReqAuth) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthLoginReqAuth) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPlatform

`func (o *AuthLoginReqAuth) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AuthLoginReqAuth) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AuthLoginReqAuth) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AuthLoginReqAuth) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRsaKeyId

`func (o *AuthLoginReqAuth) GetRsaKeyId() string`

GetRsaKeyId returns the RsaKeyId field if non-nil, zero value otherwise.

### GetRsaKeyIdOk

`func (o *AuthLoginReqAuth) GetRsaKeyIdOk() (*string, bool)`

GetRsaKeyIdOk returns a tuple with the RsaKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaKeyId

`func (o *AuthLoginReqAuth) SetRsaKeyId(v string)`

SetRsaKeyId sets RsaKeyId field to given value.

### HasRsaKeyId

`func (o *AuthLoginReqAuth) HasRsaKeyId() bool`

HasRsaKeyId returns a boolean if a field has been set.

### GetSign

`func (o *AuthLoginReqAuth) GetSign() string`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *AuthLoginReqAuth) GetSignOk() (*string, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *AuthLoginReqAuth) SetSign(v string)`

SetSign sets Sign field to given value.

### HasSign

`func (o *AuthLoginReqAuth) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetTimestamp

`func (o *AuthLoginReqAuth) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuthLoginReqAuth) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuthLoginReqAuth) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AuthLoginReqAuth) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetToken

`func (o *AuthLoginReqAuth) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthLoginReqAuth) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthLoginReqAuth) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthLoginReqAuth) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetVerificationCode

`func (o *AuthLoginReqAuth) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *AuthLoginReqAuth) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *AuthLoginReqAuth) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.

### HasVerificationCode

`func (o *AuthLoginReqAuth) HasVerificationCode() bool`

HasVerificationCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


