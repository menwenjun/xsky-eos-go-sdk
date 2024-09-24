/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AuthSecurityPolicyUpdateReqAuthSecurityPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthSecurityPolicyUpdateReqAuthSecurityPolicy{}

// AuthSecurityPolicyUpdateReqAuthSecurityPolicy struct for AuthSecurityPolicyUpdateReqAuthSecurityPolicy
type AuthSecurityPolicyUpdateReqAuthSecurityPolicy struct {
	Enabled *bool `json:"enabled,omitempty"`
	EnhancedPassword *bool `json:"enhanced_password,omitempty"`
	ExpirationTime *int64 `json:"expiration_time,omitempty"`
	FreezeDuration *int64 `json:"freeze_duration,omitempty"`
	MaxAttempts *int64 `json:"max_attempts,omitempty"`
	OneSessionOnly *bool `json:"one_session_only,omitempty"`
	PasswordExpirationWarning *bool `json:"password_expiration_warning,omitempty"`
	PasswordLifetime *int64 `json:"password_lifetime,omitempty"`
	TwoFactor *bool `json:"two_factor,omitempty"`
	VerificationCode *string `json:"verification_code,omitempty"`
}

// NewAuthSecurityPolicyUpdateReqAuthSecurityPolicy instantiates a new AuthSecurityPolicyUpdateReqAuthSecurityPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthSecurityPolicyUpdateReqAuthSecurityPolicy() *AuthSecurityPolicyUpdateReqAuthSecurityPolicy {
	this := AuthSecurityPolicyUpdateReqAuthSecurityPolicy{}
	return &this
}

// NewAuthSecurityPolicyUpdateReqAuthSecurityPolicyWithDefaults instantiates a new AuthSecurityPolicyUpdateReqAuthSecurityPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthSecurityPolicyUpdateReqAuthSecurityPolicyWithDefaults() *AuthSecurityPolicyUpdateReqAuthSecurityPolicy {
	this := AuthSecurityPolicyUpdateReqAuthSecurityPolicy{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEnhancedPassword returns the EnhancedPassword field value if set, zero value otherwise.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetEnhancedPassword() bool {
	if o == nil || IsNil(o.EnhancedPassword) {
		var ret bool
		return ret
	}
	return *o.EnhancedPassword
}

// GetEnhancedPasswordOk returns a tuple with the EnhancedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetEnhancedPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.EnhancedPassword) {
		return nil, false
	}
	return o.EnhancedPassword, true
}

// HasEnhancedPassword returns a boolean if a field has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasEnhancedPassword() bool {
	if o != nil && !IsNil(o.EnhancedPassword) {
		return true
	}

	return false
}

// SetEnhancedPassword gets a reference to the given bool and assigns it to the EnhancedPassword field.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetEnhancedPassword(v bool) {
	o.EnhancedPassword = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetExpirationTime() int64 {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret int64
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetExpirationTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given int64 and assigns it to the ExpirationTime field.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetExpirationTime(v int64) {
	o.ExpirationTime = &v
}

// GetFreezeDuration returns the FreezeDuration field value if set, zero value otherwise.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetFreezeDuration() int64 {
	if o == nil || IsNil(o.FreezeDuration) {
		var ret int64
		return ret
	}
	return *o.FreezeDuration
}

// GetFreezeDurationOk returns a tuple with the FreezeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetFreezeDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.FreezeDuration) {
		return nil, false
	}
	return o.FreezeDuration, true
}

// HasFreezeDuration returns a boolean if a field has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasFreezeDuration() bool {
	if o != nil && !IsNil(o.FreezeDuration) {
		return true
	}

	return false
}

// SetFreezeDuration gets a reference to the given int64 and assigns it to the FreezeDuration field.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetFreezeDuration(v int64) {
	o.FreezeDuration = &v
}

// GetMaxAttempts returns the MaxAttempts field value if set, zero value otherwise.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetMaxAttempts() int64 {
	if o == nil || IsNil(o.MaxAttempts) {
		var ret int64
		return ret
	}
	return *o.MaxAttempts
}

// GetMaxAttemptsOk returns a tuple with the MaxAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetMaxAttemptsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxAttempts) {
		return nil, false
	}
	return o.MaxAttempts, true
}

// HasMaxAttempts returns a boolean if a field has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasMaxAttempts() bool {
	if o != nil && !IsNil(o.MaxAttempts) {
		return true
	}

	return false
}

// SetMaxAttempts gets a reference to the given int64 and assigns it to the MaxAttempts field.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetMaxAttempts(v int64) {
	o.MaxAttempts = &v
}

// GetOneSessionOnly returns the OneSessionOnly field value if set, zero value otherwise.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetOneSessionOnly() bool {
	if o == nil || IsNil(o.OneSessionOnly) {
		var ret bool
		return ret
	}
	return *o.OneSessionOnly
}

// GetOneSessionOnlyOk returns a tuple with the OneSessionOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetOneSessionOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.OneSessionOnly) {
		return nil, false
	}
	return o.OneSessionOnly, true
}

// HasOneSessionOnly returns a boolean if a field has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasOneSessionOnly() bool {
	if o != nil && !IsNil(o.OneSessionOnly) {
		return true
	}

	return false
}

// SetOneSessionOnly gets a reference to the given bool and assigns it to the OneSessionOnly field.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetOneSessionOnly(v bool) {
	o.OneSessionOnly = &v
}

// GetPasswordExpirationWarning returns the PasswordExpirationWarning field value if set, zero value otherwise.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetPasswordExpirationWarning() bool {
	if o == nil || IsNil(o.PasswordExpirationWarning) {
		var ret bool
		return ret
	}
	return *o.PasswordExpirationWarning
}

// GetPasswordExpirationWarningOk returns a tuple with the PasswordExpirationWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetPasswordExpirationWarningOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordExpirationWarning) {
		return nil, false
	}
	return o.PasswordExpirationWarning, true
}

// HasPasswordExpirationWarning returns a boolean if a field has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasPasswordExpirationWarning() bool {
	if o != nil && !IsNil(o.PasswordExpirationWarning) {
		return true
	}

	return false
}

// SetPasswordExpirationWarning gets a reference to the given bool and assigns it to the PasswordExpirationWarning field.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetPasswordExpirationWarning(v bool) {
	o.PasswordExpirationWarning = &v
}

// GetPasswordLifetime returns the PasswordLifetime field value if set, zero value otherwise.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetPasswordLifetime() int64 {
	if o == nil || IsNil(o.PasswordLifetime) {
		var ret int64
		return ret
	}
	return *o.PasswordLifetime
}

// GetPasswordLifetimeOk returns a tuple with the PasswordLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetPasswordLifetimeOk() (*int64, bool) {
	if o == nil || IsNil(o.PasswordLifetime) {
		return nil, false
	}
	return o.PasswordLifetime, true
}

// HasPasswordLifetime returns a boolean if a field has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasPasswordLifetime() bool {
	if o != nil && !IsNil(o.PasswordLifetime) {
		return true
	}

	return false
}

// SetPasswordLifetime gets a reference to the given int64 and assigns it to the PasswordLifetime field.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetPasswordLifetime(v int64) {
	o.PasswordLifetime = &v
}

// GetTwoFactor returns the TwoFactor field value if set, zero value otherwise.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetTwoFactor() bool {
	if o == nil || IsNil(o.TwoFactor) {
		var ret bool
		return ret
	}
	return *o.TwoFactor
}

// GetTwoFactorOk returns a tuple with the TwoFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetTwoFactorOk() (*bool, bool) {
	if o == nil || IsNil(o.TwoFactor) {
		return nil, false
	}
	return o.TwoFactor, true
}

// HasTwoFactor returns a boolean if a field has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasTwoFactor() bool {
	if o != nil && !IsNil(o.TwoFactor) {
		return true
	}

	return false
}

// SetTwoFactor gets a reference to the given bool and assigns it to the TwoFactor field.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetTwoFactor(v bool) {
	o.TwoFactor = &v
}

// GetVerificationCode returns the VerificationCode field value if set, zero value otherwise.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetVerificationCode() string {
	if o == nil || IsNil(o.VerificationCode) {
		var ret string
		return ret
	}
	return *o.VerificationCode
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) GetVerificationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationCode) {
		return nil, false
	}
	return o.VerificationCode, true
}

// HasVerificationCode returns a boolean if a field has been set.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) HasVerificationCode() bool {
	if o != nil && !IsNil(o.VerificationCode) {
		return true
	}

	return false
}

// SetVerificationCode gets a reference to the given string and assigns it to the VerificationCode field.
func (o *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) SetVerificationCode(v string) {
	o.VerificationCode = &v
}

func (o AuthSecurityPolicyUpdateReqAuthSecurityPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthSecurityPolicyUpdateReqAuthSecurityPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.EnhancedPassword) {
		toSerialize["enhanced_password"] = o.EnhancedPassword
	}
	if !IsNil(o.ExpirationTime) {
		toSerialize["expiration_time"] = o.ExpirationTime
	}
	if !IsNil(o.FreezeDuration) {
		toSerialize["freeze_duration"] = o.FreezeDuration
	}
	if !IsNil(o.MaxAttempts) {
		toSerialize["max_attempts"] = o.MaxAttempts
	}
	if !IsNil(o.OneSessionOnly) {
		toSerialize["one_session_only"] = o.OneSessionOnly
	}
	if !IsNil(o.PasswordExpirationWarning) {
		toSerialize["password_expiration_warning"] = o.PasswordExpirationWarning
	}
	if !IsNil(o.PasswordLifetime) {
		toSerialize["password_lifetime"] = o.PasswordLifetime
	}
	if !IsNil(o.TwoFactor) {
		toSerialize["two_factor"] = o.TwoFactor
	}
	if !IsNil(o.VerificationCode) {
		toSerialize["verification_code"] = o.VerificationCode
	}
	return toSerialize, nil
}

type NullableAuthSecurityPolicyUpdateReqAuthSecurityPolicy struct {
	value *AuthSecurityPolicyUpdateReqAuthSecurityPolicy
	isSet bool
}

func (v NullableAuthSecurityPolicyUpdateReqAuthSecurityPolicy) Get() *AuthSecurityPolicyUpdateReqAuthSecurityPolicy {
	return v.value
}

func (v *NullableAuthSecurityPolicyUpdateReqAuthSecurityPolicy) Set(val *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthSecurityPolicyUpdateReqAuthSecurityPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthSecurityPolicyUpdateReqAuthSecurityPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthSecurityPolicyUpdateReqAuthSecurityPolicy(val *AuthSecurityPolicyUpdateReqAuthSecurityPolicy) *NullableAuthSecurityPolicyUpdateReqAuthSecurityPolicy {
	return &NullableAuthSecurityPolicyUpdateReqAuthSecurityPolicy{value: val, isSet: true}
}

func (v NullableAuthSecurityPolicyUpdateReqAuthSecurityPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthSecurityPolicyUpdateReqAuthSecurityPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


