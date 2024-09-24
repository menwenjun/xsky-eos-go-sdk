# EmailConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**SmtpAuthEnabled** | Pointer to **bool** |  | [optional] 
**SmtpEmailFrom** | Pointer to **string** |  | [optional] 
**SmtpEnableSsl** | Pointer to **bool** |  | [optional] 
**SmtpHost** | Pointer to **string** |  | [optional] 
**SmtpPassword** | Pointer to **string** |  | [optional] 
**SmtpPort** | Pointer to **int64** |  | [optional] 
**SmtpRsaKeyId** | Pointer to **string** |  | [optional] 
**SmtpSkipTlsCertVerify** | Pointer to **bool** |  | [optional] 
**SmtpSslMode** | Pointer to **string** |  | [optional] 
**SmtpUser** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailConfig

`func NewEmailConfig() *EmailConfig`

NewEmailConfig instantiates a new EmailConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConfigWithDefaults

`func NewEmailConfigWithDefaults() *EmailConfig`

NewEmailConfigWithDefaults instantiates a new EmailConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *EmailConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EmailConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EmailConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EmailConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSmtpAuthEnabled

`func (o *EmailConfig) GetSmtpAuthEnabled() bool`

GetSmtpAuthEnabled returns the SmtpAuthEnabled field if non-nil, zero value otherwise.

### GetSmtpAuthEnabledOk

`func (o *EmailConfig) GetSmtpAuthEnabledOk() (*bool, bool)`

GetSmtpAuthEnabledOk returns a tuple with the SmtpAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpAuthEnabled

`func (o *EmailConfig) SetSmtpAuthEnabled(v bool)`

SetSmtpAuthEnabled sets SmtpAuthEnabled field to given value.

### HasSmtpAuthEnabled

`func (o *EmailConfig) HasSmtpAuthEnabled() bool`

HasSmtpAuthEnabled returns a boolean if a field has been set.

### GetSmtpEmailFrom

`func (o *EmailConfig) GetSmtpEmailFrom() string`

GetSmtpEmailFrom returns the SmtpEmailFrom field if non-nil, zero value otherwise.

### GetSmtpEmailFromOk

`func (o *EmailConfig) GetSmtpEmailFromOk() (*string, bool)`

GetSmtpEmailFromOk returns a tuple with the SmtpEmailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpEmailFrom

`func (o *EmailConfig) SetSmtpEmailFrom(v string)`

SetSmtpEmailFrom sets SmtpEmailFrom field to given value.

### HasSmtpEmailFrom

`func (o *EmailConfig) HasSmtpEmailFrom() bool`

HasSmtpEmailFrom returns a boolean if a field has been set.

### GetSmtpEnableSsl

`func (o *EmailConfig) GetSmtpEnableSsl() bool`

GetSmtpEnableSsl returns the SmtpEnableSsl field if non-nil, zero value otherwise.

### GetSmtpEnableSslOk

`func (o *EmailConfig) GetSmtpEnableSslOk() (*bool, bool)`

GetSmtpEnableSslOk returns a tuple with the SmtpEnableSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpEnableSsl

`func (o *EmailConfig) SetSmtpEnableSsl(v bool)`

SetSmtpEnableSsl sets SmtpEnableSsl field to given value.

### HasSmtpEnableSsl

`func (o *EmailConfig) HasSmtpEnableSsl() bool`

HasSmtpEnableSsl returns a boolean if a field has been set.

### GetSmtpHost

`func (o *EmailConfig) GetSmtpHost() string`

GetSmtpHost returns the SmtpHost field if non-nil, zero value otherwise.

### GetSmtpHostOk

`func (o *EmailConfig) GetSmtpHostOk() (*string, bool)`

GetSmtpHostOk returns a tuple with the SmtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpHost

`func (o *EmailConfig) SetSmtpHost(v string)`

SetSmtpHost sets SmtpHost field to given value.

### HasSmtpHost

`func (o *EmailConfig) HasSmtpHost() bool`

HasSmtpHost returns a boolean if a field has been set.

### GetSmtpPassword

`func (o *EmailConfig) GetSmtpPassword() string`

GetSmtpPassword returns the SmtpPassword field if non-nil, zero value otherwise.

### GetSmtpPasswordOk

`func (o *EmailConfig) GetSmtpPasswordOk() (*string, bool)`

GetSmtpPasswordOk returns a tuple with the SmtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPassword

`func (o *EmailConfig) SetSmtpPassword(v string)`

SetSmtpPassword sets SmtpPassword field to given value.

### HasSmtpPassword

`func (o *EmailConfig) HasSmtpPassword() bool`

HasSmtpPassword returns a boolean if a field has been set.

### GetSmtpPort

`func (o *EmailConfig) GetSmtpPort() int64`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *EmailConfig) GetSmtpPortOk() (*int64, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *EmailConfig) SetSmtpPort(v int64)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *EmailConfig) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpRsaKeyId

`func (o *EmailConfig) GetSmtpRsaKeyId() string`

GetSmtpRsaKeyId returns the SmtpRsaKeyId field if non-nil, zero value otherwise.

### GetSmtpRsaKeyIdOk

`func (o *EmailConfig) GetSmtpRsaKeyIdOk() (*string, bool)`

GetSmtpRsaKeyIdOk returns a tuple with the SmtpRsaKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpRsaKeyId

`func (o *EmailConfig) SetSmtpRsaKeyId(v string)`

SetSmtpRsaKeyId sets SmtpRsaKeyId field to given value.

### HasSmtpRsaKeyId

`func (o *EmailConfig) HasSmtpRsaKeyId() bool`

HasSmtpRsaKeyId returns a boolean if a field has been set.

### GetSmtpSkipTlsCertVerify

`func (o *EmailConfig) GetSmtpSkipTlsCertVerify() bool`

GetSmtpSkipTlsCertVerify returns the SmtpSkipTlsCertVerify field if non-nil, zero value otherwise.

### GetSmtpSkipTlsCertVerifyOk

`func (o *EmailConfig) GetSmtpSkipTlsCertVerifyOk() (*bool, bool)`

GetSmtpSkipTlsCertVerifyOk returns a tuple with the SmtpSkipTlsCertVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSkipTlsCertVerify

`func (o *EmailConfig) SetSmtpSkipTlsCertVerify(v bool)`

SetSmtpSkipTlsCertVerify sets SmtpSkipTlsCertVerify field to given value.

### HasSmtpSkipTlsCertVerify

`func (o *EmailConfig) HasSmtpSkipTlsCertVerify() bool`

HasSmtpSkipTlsCertVerify returns a boolean if a field has been set.

### GetSmtpSslMode

`func (o *EmailConfig) GetSmtpSslMode() string`

GetSmtpSslMode returns the SmtpSslMode field if non-nil, zero value otherwise.

### GetSmtpSslModeOk

`func (o *EmailConfig) GetSmtpSslModeOk() (*string, bool)`

GetSmtpSslModeOk returns a tuple with the SmtpSslMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSslMode

`func (o *EmailConfig) SetSmtpSslMode(v string)`

SetSmtpSslMode sets SmtpSslMode field to given value.

### HasSmtpSslMode

`func (o *EmailConfig) HasSmtpSslMode() bool`

HasSmtpSslMode returns a boolean if a field has been set.

### GetSmtpUser

`func (o *EmailConfig) GetSmtpUser() string`

GetSmtpUser returns the SmtpUser field if non-nil, zero value otherwise.

### GetSmtpUserOk

`func (o *EmailConfig) GetSmtpUserOk() (*string, bool)`

GetSmtpUserOk returns a tuple with the SmtpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpUser

`func (o *EmailConfig) SetSmtpUser(v string)`

SetSmtpUser sets SmtpUser field to given value.

### HasSmtpUser

`func (o *EmailConfig) HasSmtpUser() bool`

HasSmtpUser returns a boolean if a field has been set.

### GetSubject

`func (o *EmailConfig) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailConfig) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailConfig) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailConfig) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


