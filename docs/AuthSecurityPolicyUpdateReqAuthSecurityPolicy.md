# AuthSecurityPolicyUpdateReqAuthSecurityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**EnhancedPassword** | Pointer to **bool** |  | [optional] 
**ExpirationTime** | Pointer to **int64** |  | [optional] 
**FreezeDuration** | Pointer to **int64** |  | [optional] 
**MaxAttempts** | Pointer to **int64** |  | [optional] 
**OneSessionOnly** | Pointer to **bool** |  | [optional] 
**PasswordExpirationWarning** | Pointer to **bool** |  | [optional] 
**PasswordLifetime** | Pointer to **int64** |  | [optional] 
**TwoFactor** | Pointer to **bool** |  | [optional] 
**VerificationCode** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthSecurityPolicyUpdateReqAuthSecurityPolicy

`func NewAuthSecurityPolicyUpdateReqAuthSecurityPolicy() *AuthSecurityPolicyUpdateReqAuthSecurityPolicy`

NewAuthSecurityPolicyUpdateReqAuthSecurityPolicy instantiates a new AuthSecurityPolicyUpdateReqAuthSecurityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSecurityPolicyUpdateReqAuthSecurityPolicyWithDefaults

`func NewAuthSecurityPolicyUpdateReqAuthSecurityPolicyWithDefaults() *AuthSecurityPolicyUpdateReqAuthSecurityPolicy`

NewAuthSecurityPolicyUpdateReqAuthSecurityPolicyWithDefaults instantiates a new AuthSecurityPolicyUpdateReqAuthSecurityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnhancedPassword

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetEnhancedPassword() bool`

GetEnhancedPassword returns the EnhancedPassword field if non-nil, zero value otherwise.

### GetEnhancedPasswordOk

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetEnhancedPasswordOk() (*bool, bool)`

GetEnhancedPasswordOk returns a tuple with the EnhancedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedPassword

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetEnhancedPassword(v bool)`

SetEnhancedPassword sets EnhancedPassword field to given value.

### HasEnhancedPassword

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasEnhancedPassword() bool`

HasEnhancedPassword returns a boolean if a field has been set.

### GetExpirationTime

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetFreezeDuration

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetFreezeDuration() int64`

GetFreezeDuration returns the FreezeDuration field if non-nil, zero value otherwise.

### GetFreezeDurationOk

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetFreezeDurationOk() (*int64, bool)`

GetFreezeDurationOk returns a tuple with the FreezeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeDuration

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetFreezeDuration(v int64)`

SetFreezeDuration sets FreezeDuration field to given value.

### HasFreezeDuration

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasFreezeDuration() bool`

HasFreezeDuration returns a boolean if a field has been set.

### GetMaxAttempts

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetMaxAttempts() int64`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetMaxAttemptsOk() (*int64, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetMaxAttempts(v int64)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetOneSessionOnly

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetOneSessionOnly() bool`

GetOneSessionOnly returns the OneSessionOnly field if non-nil, zero value otherwise.

### GetOneSessionOnlyOk

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetOneSessionOnlyOk() (*bool, bool)`

GetOneSessionOnlyOk returns a tuple with the OneSessionOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneSessionOnly

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetOneSessionOnly(v bool)`

SetOneSessionOnly sets OneSessionOnly field to given value.

### HasOneSessionOnly

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasOneSessionOnly() bool`

HasOneSessionOnly returns a boolean if a field has been set.

### GetPasswordExpirationWarning

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetPasswordExpirationWarning() bool`

GetPasswordExpirationWarning returns the PasswordExpirationWarning field if non-nil, zero value otherwise.

### GetPasswordExpirationWarningOk

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetPasswordExpirationWarningOk() (*bool, bool)`

GetPasswordExpirationWarningOk returns a tuple with the PasswordExpirationWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationWarning

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetPasswordExpirationWarning(v bool)`

SetPasswordExpirationWarning sets PasswordExpirationWarning field to given value.

### HasPasswordExpirationWarning

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasPasswordExpirationWarning() bool`

HasPasswordExpirationWarning returns a boolean if a field has been set.

### GetPasswordLifetime

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetPasswordLifetime() int64`

GetPasswordLifetime returns the PasswordLifetime field if non-nil, zero value otherwise.

### GetPasswordLifetimeOk

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetPasswordLifetimeOk() (*int64, bool)`

GetPasswordLifetimeOk returns a tuple with the PasswordLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLifetime

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetPasswordLifetime(v int64)`

SetPasswordLifetime sets PasswordLifetime field to given value.

### HasPasswordLifetime

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasPasswordLifetime() bool`

HasPasswordLifetime returns a boolean if a field has been set.

### GetTwoFactor

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetTwoFactor() bool`

GetTwoFactor returns the TwoFactor field if non-nil, zero value otherwise.

### GetTwoFactorOk

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetTwoFactorOk() (*bool, bool)`

GetTwoFactorOk returns a tuple with the TwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactor

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetTwoFactor(v bool)`

SetTwoFactor sets TwoFactor field to given value.

### HasTwoFactor

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasTwoFactor() bool`

HasTwoFactor returns a boolean if a field has been set.

### GetVerificationCode

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.

### HasVerificationCode

`func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasVerificationCode() bool`

HasVerificationCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


