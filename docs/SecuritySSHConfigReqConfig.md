# SecuritySSHConfigReqConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedPassword** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PermitRootLogin** | Pointer to **bool** |  | [optional] 
**Port** | **int64** |  | 
**PubKeyAuthentication** | Pointer to **bool** |  | [optional] 
**RsaKeyId** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewSecuritySSHConfigReqConfig

`func NewSecuritySSHConfigReqConfig(port int64, ) *SecuritySSHConfigReqConfig`

NewSecuritySSHConfigReqConfig instantiates a new SecuritySSHConfigReqConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySSHConfigReqConfigWithDefaults

`func NewSecuritySSHConfigReqConfigWithDefaults() *SecuritySSHConfigReqConfig`

NewSecuritySSHConfigReqConfigWithDefaults instantiates a new SecuritySSHConfigReqConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedPassword

`func (o *SecuritySSHConfigReqConfig) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *SecuritySSHConfigReqConfig) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *SecuritySSHConfigReqConfig) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *SecuritySSHConfigReqConfig) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetPassword

`func (o *SecuritySSHConfigReqConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SecuritySSHConfigReqConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SecuritySSHConfigReqConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SecuritySSHConfigReqConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPermitRootLogin

`func (o *SecuritySSHConfigReqConfig) GetPermitRootLogin() bool`

GetPermitRootLogin returns the PermitRootLogin field if non-nil, zero value otherwise.

### GetPermitRootLoginOk

`func (o *SecuritySSHConfigReqConfig) GetPermitRootLoginOk() (*bool, bool)`

GetPermitRootLoginOk returns a tuple with the PermitRootLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitRootLogin

`func (o *SecuritySSHConfigReqConfig) SetPermitRootLogin(v bool)`

SetPermitRootLogin sets PermitRootLogin field to given value.

### HasPermitRootLogin

`func (o *SecuritySSHConfigReqConfig) HasPermitRootLogin() bool`

HasPermitRootLogin returns a boolean if a field has been set.

### GetPort

`func (o *SecuritySSHConfigReqConfig) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SecuritySSHConfigReqConfig) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SecuritySSHConfigReqConfig) SetPort(v int64)`

SetPort sets Port field to given value.


### GetPubKeyAuthentication

`func (o *SecuritySSHConfigReqConfig) GetPubKeyAuthentication() bool`

GetPubKeyAuthentication returns the PubKeyAuthentication field if non-nil, zero value otherwise.

### GetPubKeyAuthenticationOk

`func (o *SecuritySSHConfigReqConfig) GetPubKeyAuthenticationOk() (*bool, bool)`

GetPubKeyAuthenticationOk returns a tuple with the PubKeyAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKeyAuthentication

`func (o *SecuritySSHConfigReqConfig) SetPubKeyAuthentication(v bool)`

SetPubKeyAuthentication sets PubKeyAuthentication field to given value.

### HasPubKeyAuthentication

`func (o *SecuritySSHConfigReqConfig) HasPubKeyAuthentication() bool`

HasPubKeyAuthentication returns a boolean if a field has been set.

### GetRsaKeyId

`func (o *SecuritySSHConfigReqConfig) GetRsaKeyId() string`

GetRsaKeyId returns the RsaKeyId field if non-nil, zero value otherwise.

### GetRsaKeyIdOk

`func (o *SecuritySSHConfigReqConfig) GetRsaKeyIdOk() (*string, bool)`

GetRsaKeyIdOk returns a tuple with the RsaKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaKeyId

`func (o *SecuritySSHConfigReqConfig) SetRsaKeyId(v string)`

SetRsaKeyId sets RsaKeyId field to given value.

### HasRsaKeyId

`func (o *SecuritySSHConfigReqConfig) HasRsaKeyId() bool`

HasRsaKeyId returns a boolean if a field has been set.

### GetTimeout

`func (o *SecuritySSHConfigReqConfig) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SecuritySSHConfigReqConfig) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SecuritySSHConfigReqConfig) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SecuritySSHConfigReqConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUsername

`func (o *SecuritySSHConfigReqConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SecuritySSHConfigReqConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SecuritySSHConfigReqConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SecuritySSHConfigReqConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


