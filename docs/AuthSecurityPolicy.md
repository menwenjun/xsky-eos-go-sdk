# AuthSecurityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**EnhancedPassword** | Pointer to **bool** |  | [optional] 
**ExpirationTime** | Pointer to **int64** |  | [optional] 
**FreezeDuration** | Pointer to **int64** | unit: second | [optional] 
**MaxAttempts** | Pointer to **int64** |  | [optional] 
**OneSessionOnly** | Pointer to **bool** |  | [optional] 
**PasswordExpirationWarning** | Pointer to **bool** |  | [optional] 
**PasswordLifetime** | Pointer to **int64** | unit: second | [optional] 
**TwoFactor** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthSecurityPolicy

`func NewAuthSecurityPolicy() *AuthSecurityPolicy`

NewAuthSecurityPolicy instantiates a new AuthSecurityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSecurityPolicyWithDefaults

`func NewAuthSecurityPolicyWithDefaults() *AuthSecurityPolicy`

NewAuthSecurityPolicyWithDefaults instantiates a new AuthSecurityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AuthSecurityPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthSecurityPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthSecurityPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthSecurityPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnhancedPassword

`func (o *AuthSecurityPolicy) GetEnhancedPassword() bool`

GetEnhancedPassword returns the EnhancedPassword field if non-nil, zero value otherwise.

### GetEnhancedPasswordOk

`func (o *AuthSecurityPolicy) GetEnhancedPasswordOk() (*bool, bool)`

GetEnhancedPasswordOk returns a tuple with the EnhancedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedPassword

`func (o *AuthSecurityPolicy) SetEnhancedPassword(v bool)`

SetEnhancedPassword sets EnhancedPassword field to given value.

### HasEnhancedPassword

`func (o *AuthSecurityPolicy) HasEnhancedPassword() bool`

HasEnhancedPassword returns a boolean if a field has been set.

### GetExpirationTime

`func (o *AuthSecurityPolicy) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *AuthSecurityPolicy) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *AuthSecurityPolicy) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *AuthSecurityPolicy) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetFreezeDuration

`func (o *AuthSecurityPolicy) GetFreezeDuration() int64`

GetFreezeDuration returns the FreezeDuration field if non-nil, zero value otherwise.

### GetFreezeDurationOk

`func (o *AuthSecurityPolicy) GetFreezeDurationOk() (*int64, bool)`

GetFreezeDurationOk returns a tuple with the FreezeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeDuration

`func (o *AuthSecurityPolicy) SetFreezeDuration(v int64)`

SetFreezeDuration sets FreezeDuration field to given value.

### HasFreezeDuration

`func (o *AuthSecurityPolicy) HasFreezeDuration() bool`

HasFreezeDuration returns a boolean if a field has been set.

### GetMaxAttempts

`func (o *AuthSecurityPolicy) GetMaxAttempts() int64`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *AuthSecurityPolicy) GetMaxAttemptsOk() (*int64, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *AuthSecurityPolicy) SetMaxAttempts(v int64)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *AuthSecurityPolicy) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetOneSessionOnly

`func (o *AuthSecurityPolicy) GetOneSessionOnly() bool`

GetOneSessionOnly returns the OneSessionOnly field if non-nil, zero value otherwise.

### GetOneSessionOnlyOk

`func (o *AuthSecurityPolicy) GetOneSessionOnlyOk() (*bool, bool)`

GetOneSessionOnlyOk returns a tuple with the OneSessionOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneSessionOnly

`func (o *AuthSecurityPolicy) SetOneSessionOnly(v bool)`

SetOneSessionOnly sets OneSessionOnly field to given value.

### HasOneSessionOnly

`func (o *AuthSecurityPolicy) HasOneSessionOnly() bool`

HasOneSessionOnly returns a boolean if a field has been set.

### GetPasswordExpirationWarning

`func (o *AuthSecurityPolicy) GetPasswordExpirationWarning() bool`

GetPasswordExpirationWarning returns the PasswordExpirationWarning field if non-nil, zero value otherwise.

### GetPasswordExpirationWarningOk

`func (o *AuthSecurityPolicy) GetPasswordExpirationWarningOk() (*bool, bool)`

GetPasswordExpirationWarningOk returns a tuple with the PasswordExpirationWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationWarning

`func (o *AuthSecurityPolicy) SetPasswordExpirationWarning(v bool)`

SetPasswordExpirationWarning sets PasswordExpirationWarning field to given value.

### HasPasswordExpirationWarning

`func (o *AuthSecurityPolicy) HasPasswordExpirationWarning() bool`

HasPasswordExpirationWarning returns a boolean if a field has been set.

### GetPasswordLifetime

`func (o *AuthSecurityPolicy) GetPasswordLifetime() int64`

GetPasswordLifetime returns the PasswordLifetime field if non-nil, zero value otherwise.

### GetPasswordLifetimeOk

`func (o *AuthSecurityPolicy) GetPasswordLifetimeOk() (*int64, bool)`

GetPasswordLifetimeOk returns a tuple with the PasswordLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLifetime

`func (o *AuthSecurityPolicy) SetPasswordLifetime(v int64)`

SetPasswordLifetime sets PasswordLifetime field to given value.

### HasPasswordLifetime

`func (o *AuthSecurityPolicy) HasPasswordLifetime() bool`

HasPasswordLifetime returns a boolean if a field has been set.

### GetTwoFactor

`func (o *AuthSecurityPolicy) GetTwoFactor() bool`

GetTwoFactor returns the TwoFactor field if non-nil, zero value otherwise.

### GetTwoFactorOk

`func (o *AuthSecurityPolicy) GetTwoFactorOk() (*bool, bool)`

GetTwoFactorOk returns a tuple with the TwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactor

`func (o *AuthSecurityPolicy) SetTwoFactor(v bool)`

SetTwoFactor sets TwoFactor field to given value.

### HasTwoFactor

`func (o *AuthSecurityPolicy) HasTwoFactor() bool`

HasTwoFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


