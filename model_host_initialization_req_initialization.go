/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HostInitializationReqInitialization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostInitializationReqInitialization{}

// HostInitializationReqInitialization struct for HostInitializationReqInitialization
type HostInitializationReqInitialization struct {
	AdminIps []string `json:"admin_ips"`
	DisableFirewalld *bool `json:"disable_firewalld,omitempty"`
	HostnamePrefix *string `json:"hostname_prefix,omitempty"`
	HostnameSuffixLen *int64 `json:"hostname_suffix_len,omitempty"`
	RsaKeyId *string `json:"rsa_key_id,omitempty"`
	SetChrony *bool `json:"set_chrony,omitempty"`
	SetHostname *bool `json:"set_hostname,omitempty"`
	SshPassword *string `json:"ssh_password,omitempty"`
	SshUser string `json:"ssh_user"`
}

type _HostInitializationReqInitialization HostInitializationReqInitialization

// NewHostInitializationReqInitialization instantiates a new HostInitializationReqInitialization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostInitializationReqInitialization(adminIps []string, sshUser string) *HostInitializationReqInitialization {
	this := HostInitializationReqInitialization{}
	this.AdminIps = adminIps
	this.SshUser = sshUser
	return &this
}

// NewHostInitializationReqInitializationWithDefaults instantiates a new HostInitializationReqInitialization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostInitializationReqInitializationWithDefaults() *HostInitializationReqInitialization {
	this := HostInitializationReqInitialization{}
	return &this
}

// GetAdminIps returns the AdminIps field value
func (o *HostInitializationReqInitialization) GetAdminIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AdminIps
}

// GetAdminIpsOk returns a tuple with the AdminIps field value
// and a boolean to check if the value has been set.
func (o *HostInitializationReqInitialization) GetAdminIpsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdminIps, true
}

// SetAdminIps sets field value
func (o *HostInitializationReqInitialization) SetAdminIps(v []string) {
	o.AdminIps = v
}

// GetDisableFirewalld returns the DisableFirewalld field value if set, zero value otherwise.
func (o *HostInitializationReqInitialization) GetDisableFirewalld() bool {
	if o == nil || IsNil(o.DisableFirewalld) {
		var ret bool
		return ret
	}
	return *o.DisableFirewalld
}

// GetDisableFirewalldOk returns a tuple with the DisableFirewalld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitializationReqInitialization) GetDisableFirewalldOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableFirewalld) {
		return nil, false
	}
	return o.DisableFirewalld, true
}

// HasDisableFirewalld returns a boolean if a field has been set.
func (o *HostInitializationReqInitialization) HasDisableFirewalld() bool {
	if o != nil && !IsNil(o.DisableFirewalld) {
		return true
	}

	return false
}

// SetDisableFirewalld gets a reference to the given bool and assigns it to the DisableFirewalld field.
func (o *HostInitializationReqInitialization) SetDisableFirewalld(v bool) {
	o.DisableFirewalld = &v
}

// GetHostnamePrefix returns the HostnamePrefix field value if set, zero value otherwise.
func (o *HostInitializationReqInitialization) GetHostnamePrefix() string {
	if o == nil || IsNil(o.HostnamePrefix) {
		var ret string
		return ret
	}
	return *o.HostnamePrefix
}

// GetHostnamePrefixOk returns a tuple with the HostnamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitializationReqInitialization) GetHostnamePrefixOk() (*string, bool) {
	if o == nil || IsNil(o.HostnamePrefix) {
		return nil, false
	}
	return o.HostnamePrefix, true
}

// HasHostnamePrefix returns a boolean if a field has been set.
func (o *HostInitializationReqInitialization) HasHostnamePrefix() bool {
	if o != nil && !IsNil(o.HostnamePrefix) {
		return true
	}

	return false
}

// SetHostnamePrefix gets a reference to the given string and assigns it to the HostnamePrefix field.
func (o *HostInitializationReqInitialization) SetHostnamePrefix(v string) {
	o.HostnamePrefix = &v
}

// GetHostnameSuffixLen returns the HostnameSuffixLen field value if set, zero value otherwise.
func (o *HostInitializationReqInitialization) GetHostnameSuffixLen() int64 {
	if o == nil || IsNil(o.HostnameSuffixLen) {
		var ret int64
		return ret
	}
	return *o.HostnameSuffixLen
}

// GetHostnameSuffixLenOk returns a tuple with the HostnameSuffixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitializationReqInitialization) GetHostnameSuffixLenOk() (*int64, bool) {
	if o == nil || IsNil(o.HostnameSuffixLen) {
		return nil, false
	}
	return o.HostnameSuffixLen, true
}

// HasHostnameSuffixLen returns a boolean if a field has been set.
func (o *HostInitializationReqInitialization) HasHostnameSuffixLen() bool {
	if o != nil && !IsNil(o.HostnameSuffixLen) {
		return true
	}

	return false
}

// SetHostnameSuffixLen gets a reference to the given int64 and assigns it to the HostnameSuffixLen field.
func (o *HostInitializationReqInitialization) SetHostnameSuffixLen(v int64) {
	o.HostnameSuffixLen = &v
}

// GetRsaKeyId returns the RsaKeyId field value if set, zero value otherwise.
func (o *HostInitializationReqInitialization) GetRsaKeyId() string {
	if o == nil || IsNil(o.RsaKeyId) {
		var ret string
		return ret
	}
	return *o.RsaKeyId
}

// GetRsaKeyIdOk returns a tuple with the RsaKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitializationReqInitialization) GetRsaKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.RsaKeyId) {
		return nil, false
	}
	return o.RsaKeyId, true
}

// HasRsaKeyId returns a boolean if a field has been set.
func (o *HostInitializationReqInitialization) HasRsaKeyId() bool {
	if o != nil && !IsNil(o.RsaKeyId) {
		return true
	}

	return false
}

// SetRsaKeyId gets a reference to the given string and assigns it to the RsaKeyId field.
func (o *HostInitializationReqInitialization) SetRsaKeyId(v string) {
	o.RsaKeyId = &v
}

// GetSetChrony returns the SetChrony field value if set, zero value otherwise.
func (o *HostInitializationReqInitialization) GetSetChrony() bool {
	if o == nil || IsNil(o.SetChrony) {
		var ret bool
		return ret
	}
	return *o.SetChrony
}

// GetSetChronyOk returns a tuple with the SetChrony field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitializationReqInitialization) GetSetChronyOk() (*bool, bool) {
	if o == nil || IsNil(o.SetChrony) {
		return nil, false
	}
	return o.SetChrony, true
}

// HasSetChrony returns a boolean if a field has been set.
func (o *HostInitializationReqInitialization) HasSetChrony() bool {
	if o != nil && !IsNil(o.SetChrony) {
		return true
	}

	return false
}

// SetSetChrony gets a reference to the given bool and assigns it to the SetChrony field.
func (o *HostInitializationReqInitialization) SetSetChrony(v bool) {
	o.SetChrony = &v
}

// GetSetHostname returns the SetHostname field value if set, zero value otherwise.
func (o *HostInitializationReqInitialization) GetSetHostname() bool {
	if o == nil || IsNil(o.SetHostname) {
		var ret bool
		return ret
	}
	return *o.SetHostname
}

// GetSetHostnameOk returns a tuple with the SetHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitializationReqInitialization) GetSetHostnameOk() (*bool, bool) {
	if o == nil || IsNil(o.SetHostname) {
		return nil, false
	}
	return o.SetHostname, true
}

// HasSetHostname returns a boolean if a field has been set.
func (o *HostInitializationReqInitialization) HasSetHostname() bool {
	if o != nil && !IsNil(o.SetHostname) {
		return true
	}

	return false
}

// SetSetHostname gets a reference to the given bool and assigns it to the SetHostname field.
func (o *HostInitializationReqInitialization) SetSetHostname(v bool) {
	o.SetHostname = &v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *HostInitializationReqInitialization) GetSshPassword() string {
	if o == nil || IsNil(o.SshPassword) {
		var ret string
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitializationReqInitialization) GetSshPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.SshPassword) {
		return nil, false
	}
	return o.SshPassword, true
}

// HasSshPassword returns a boolean if a field has been set.
func (o *HostInitializationReqInitialization) HasSshPassword() bool {
	if o != nil && !IsNil(o.SshPassword) {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given string and assigns it to the SshPassword field.
func (o *HostInitializationReqInitialization) SetSshPassword(v string) {
	o.SshPassword = &v
}

// GetSshUser returns the SshUser field value
func (o *HostInitializationReqInitialization) GetSshUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value
// and a boolean to check if the value has been set.
func (o *HostInitializationReqInitialization) GetSshUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SshUser, true
}

// SetSshUser sets field value
func (o *HostInitializationReqInitialization) SetSshUser(v string) {
	o.SshUser = v
}

func (o HostInitializationReqInitialization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostInitializationReqInitialization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["admin_ips"] = o.AdminIps
	if !IsNil(o.DisableFirewalld) {
		toSerialize["disable_firewalld"] = o.DisableFirewalld
	}
	if !IsNil(o.HostnamePrefix) {
		toSerialize["hostname_prefix"] = o.HostnamePrefix
	}
	if !IsNil(o.HostnameSuffixLen) {
		toSerialize["hostname_suffix_len"] = o.HostnameSuffixLen
	}
	if !IsNil(o.RsaKeyId) {
		toSerialize["rsa_key_id"] = o.RsaKeyId
	}
	if !IsNil(o.SetChrony) {
		toSerialize["set_chrony"] = o.SetChrony
	}
	if !IsNil(o.SetHostname) {
		toSerialize["set_hostname"] = o.SetHostname
	}
	if !IsNil(o.SshPassword) {
		toSerialize["ssh_password"] = o.SshPassword
	}
	toSerialize["ssh_user"] = o.SshUser
	return toSerialize, nil
}

func (o *HostInitializationReqInitialization) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"admin_ips",
		"ssh_user",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHostInitializationReqInitialization := _HostInitializationReqInitialization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHostInitializationReqInitialization)

	if err != nil {
		return err
	}

	*o = HostInitializationReqInitialization(varHostInitializationReqInitialization)

	return err
}

type NullableHostInitializationReqInitialization struct {
	value *HostInitializationReqInitialization
	isSet bool
}

func (v NullableHostInitializationReqInitialization) Get() *HostInitializationReqInitialization {
	return v.value
}

func (v *NullableHostInitializationReqInitialization) Set(val *HostInitializationReqInitialization) {
	v.value = val
	v.isSet = true
}

func (v NullableHostInitializationReqInitialization) IsSet() bool {
	return v.isSet
}

func (v *NullableHostInitializationReqInitialization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostInitializationReqInitialization(val *HostInitializationReqInitialization) *NullableHostInitializationReqInitialization {
	return &NullableHostInitializationReqInitialization{value: val, isSet: true}
}

func (v NullableHostInitializationReqInitialization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostInitializationReqInitialization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


