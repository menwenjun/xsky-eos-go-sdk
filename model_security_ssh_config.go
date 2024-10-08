/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the SecuritySSHConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecuritySSHConfig{}

// SecuritySSHConfig SecuritySSHConfig is the model of ssh config
type SecuritySSHConfig struct {
	ActionStatus *string `json:"action_status,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Message *string `json:"message,omitempty"`
	PermitRootLogin *bool `json:"permit_root_login,omitempty"`
	Port *int64 `json:"port,omitempty"`
	PubKeyAuthentication *bool `json:"pub_key_authentication,omitempty"`
	Status *string `json:"status,omitempty"`
	Timeout *int64 `json:"timeout,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewSecuritySSHConfig instantiates a new SecuritySSHConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecuritySSHConfig() *SecuritySSHConfig {
	this := SecuritySSHConfig{}
	return &this
}

// NewSecuritySSHConfigWithDefaults instantiates a new SecuritySSHConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecuritySSHConfigWithDefaults() *SecuritySSHConfig {
	this := SecuritySSHConfig{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *SecuritySSHConfig) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritySSHConfig) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *SecuritySSHConfig) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *SecuritySSHConfig) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *SecuritySSHConfig) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritySSHConfig) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *SecuritySSHConfig) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *SecuritySSHConfig) SetCreate(v time.Time) {
	o.Create = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecuritySSHConfig) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritySSHConfig) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecuritySSHConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SecuritySSHConfig) SetId(v int64) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SecuritySSHConfig) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritySSHConfig) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SecuritySSHConfig) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SecuritySSHConfig) SetMessage(v string) {
	o.Message = &v
}

// GetPermitRootLogin returns the PermitRootLogin field value if set, zero value otherwise.
func (o *SecuritySSHConfig) GetPermitRootLogin() bool {
	if o == nil || IsNil(o.PermitRootLogin) {
		var ret bool
		return ret
	}
	return *o.PermitRootLogin
}

// GetPermitRootLoginOk returns a tuple with the PermitRootLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritySSHConfig) GetPermitRootLoginOk() (*bool, bool) {
	if o == nil || IsNil(o.PermitRootLogin) {
		return nil, false
	}
	return o.PermitRootLogin, true
}

// HasPermitRootLogin returns a boolean if a field has been set.
func (o *SecuritySSHConfig) HasPermitRootLogin() bool {
	if o != nil && !IsNil(o.PermitRootLogin) {
		return true
	}

	return false
}

// SetPermitRootLogin gets a reference to the given bool and assigns it to the PermitRootLogin field.
func (o *SecuritySSHConfig) SetPermitRootLogin(v bool) {
	o.PermitRootLogin = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SecuritySSHConfig) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritySSHConfig) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SecuritySSHConfig) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *SecuritySSHConfig) SetPort(v int64) {
	o.Port = &v
}

// GetPubKeyAuthentication returns the PubKeyAuthentication field value if set, zero value otherwise.
func (o *SecuritySSHConfig) GetPubKeyAuthentication() bool {
	if o == nil || IsNil(o.PubKeyAuthentication) {
		var ret bool
		return ret
	}
	return *o.PubKeyAuthentication
}

// GetPubKeyAuthenticationOk returns a tuple with the PubKeyAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritySSHConfig) GetPubKeyAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.PubKeyAuthentication) {
		return nil, false
	}
	return o.PubKeyAuthentication, true
}

// HasPubKeyAuthentication returns a boolean if a field has been set.
func (o *SecuritySSHConfig) HasPubKeyAuthentication() bool {
	if o != nil && !IsNil(o.PubKeyAuthentication) {
		return true
	}

	return false
}

// SetPubKeyAuthentication gets a reference to the given bool and assigns it to the PubKeyAuthentication field.
func (o *SecuritySSHConfig) SetPubKeyAuthentication(v bool) {
	o.PubKeyAuthentication = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecuritySSHConfig) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritySSHConfig) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecuritySSHConfig) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SecuritySSHConfig) SetStatus(v string) {
	o.Status = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *SecuritySSHConfig) GetTimeout() int64 {
	if o == nil || IsNil(o.Timeout) {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritySSHConfig) GetTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *SecuritySSHConfig) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *SecuritySSHConfig) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *SecuritySSHConfig) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritySSHConfig) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *SecuritySSHConfig) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *SecuritySSHConfig) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SecuritySSHConfig) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritySSHConfig) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SecuritySSHConfig) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SecuritySSHConfig) SetUsername(v string) {
	o.Username = &v
}

func (o SecuritySSHConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecuritySSHConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.PermitRootLogin) {
		toSerialize["permit_root_login"] = o.PermitRootLogin
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.PubKeyAuthentication) {
		toSerialize["pub_key_authentication"] = o.PubKeyAuthentication
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableSecuritySSHConfig struct {
	value *SecuritySSHConfig
	isSet bool
}

func (v NullableSecuritySSHConfig) Get() *SecuritySSHConfig {
	return v.value
}

func (v *NullableSecuritySSHConfig) Set(val *SecuritySSHConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSecuritySSHConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSecuritySSHConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecuritySSHConfig(val *SecuritySSHConfig) *NullableSecuritySSHConfig {
	return &NullableSecuritySSHConfig{value: val, isSet: true}
}

func (v NullableSecuritySSHConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecuritySSHConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


