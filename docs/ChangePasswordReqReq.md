# ChangePasswordReqReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedOriginalPassword** | Pointer to **string** | encrypted original password for auth | [optional] 
**EncryptedPassword** | Pointer to **string** | encrypted password for auth | [optional] 
**OriginalPassword** | **string** | original password of user | 
**Password** | **string** | new password of user | 
**RsaKeyId** | Pointer to **string** | rsa key id | [optional] 

## Methods

### NewChangePasswordReqReq

`func NewChangePasswordReqReq(originalPassword string, password string, ) *ChangePasswordReqReq`

NewChangePasswordReqReq instantiates a new ChangePasswordReqReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePasswordReqReqWithDefaults

`func NewChangePasswordReqReqWithDefaults() *ChangePasswordReqReq`

NewChangePasswordReqReqWithDefaults instantiates a new ChangePasswordReqReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedOriginalPassword

`func (o *ChangePasswordReqReq) GetEncryptedOriginalPassword() string`

GetEncryptedOriginalPassword returns the EncryptedOriginalPassword field if non-nil, zero value otherwise.

### GetEncryptedOriginalPasswordOk

`func (o *ChangePasswordReqReq) GetEncryptedOriginalPasswordOk() (*string, bool)`

GetEncryptedOriginalPasswordOk returns a tuple with the EncryptedOriginalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedOriginalPassword

`func (o *ChangePasswordReqReq) SetEncryptedOriginalPassword(v string)`

SetEncryptedOriginalPassword sets EncryptedOriginalPassword field to given value.

### HasEncryptedOriginalPassword

`func (o *ChangePasswordReqReq) HasEncryptedOriginalPassword() bool`

HasEncryptedOriginalPassword returns a boolean if a field has been set.

### GetEncryptedPassword

`func (o *ChangePasswordReqReq) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *ChangePasswordReqReq) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *ChangePasswordReqReq) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *ChangePasswordReqReq) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetOriginalPassword

`func (o *ChangePasswordReqReq) GetOriginalPassword() string`

GetOriginalPassword returns the OriginalPassword field if non-nil, zero value otherwise.

### GetOriginalPasswordOk

`func (o *ChangePasswordReqReq) GetOriginalPasswordOk() (*string, bool)`

GetOriginalPasswordOk returns a tuple with the OriginalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPassword

`func (o *ChangePasswordReqReq) SetOriginalPassword(v string)`

SetOriginalPassword sets OriginalPassword field to given value.


### GetPassword

`func (o *ChangePasswordReqReq) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ChangePasswordReqReq) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ChangePasswordReqReq) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRsaKeyId

`func (o *ChangePasswordReqReq) GetRsaKeyId() string`

GetRsaKeyId returns the RsaKeyId field if non-nil, zero value otherwise.

### GetRsaKeyIdOk

`func (o *ChangePasswordReqReq) GetRsaKeyIdOk() (*string, bool)`

GetRsaKeyIdOk returns a tuple with the RsaKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaKeyId

`func (o *ChangePasswordReqReq) SetRsaKeyId(v string)`

SetRsaKeyId sets RsaKeyId field to given value.

### HasRsaKeyId

`func (o *ChangePasswordReqReq) HasRsaKeyId() bool`

HasRsaKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


