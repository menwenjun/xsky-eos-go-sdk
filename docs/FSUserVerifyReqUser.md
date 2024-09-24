# FSUserVerifyReqUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedPassword** | Pointer to **string** | encrypted password for user | [optional] 
**Name** | **string** | user name or email for user | 
**Password** | Pointer to **string** | password for user | [optional] 
**RsaKeyId** | Pointer to **string** | rsa key id | [optional] 

## Methods

### NewFSUserVerifyReqUser

`func NewFSUserVerifyReqUser(name string, ) *FSUserVerifyReqUser`

NewFSUserVerifyReqUser instantiates a new FSUserVerifyReqUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSUserVerifyReqUserWithDefaults

`func NewFSUserVerifyReqUserWithDefaults() *FSUserVerifyReqUser`

NewFSUserVerifyReqUserWithDefaults instantiates a new FSUserVerifyReqUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedPassword

`func (o *FSUserVerifyReqUser) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *FSUserVerifyReqUser) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *FSUserVerifyReqUser) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *FSUserVerifyReqUser) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetName

`func (o *FSUserVerifyReqUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FSUserVerifyReqUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FSUserVerifyReqUser) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *FSUserVerifyReqUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FSUserVerifyReqUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FSUserVerifyReqUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FSUserVerifyReqUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRsaKeyId

`func (o *FSUserVerifyReqUser) GetRsaKeyId() string`

GetRsaKeyId returns the RsaKeyId field if non-nil, zero value otherwise.

### GetRsaKeyIdOk

`func (o *FSUserVerifyReqUser) GetRsaKeyIdOk() (*string, bool)`

GetRsaKeyIdOk returns a tuple with the RsaKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaKeyId

`func (o *FSUserVerifyReqUser) SetRsaKeyId(v string)`

SetRsaKeyId sets RsaKeyId field to given value.

### HasRsaKeyId

`func (o *FSUserVerifyReqUser) HasRsaKeyId() bool`

HasRsaKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


