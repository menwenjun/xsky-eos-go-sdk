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

// checks if the FSLdapCreateReqInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FSLdapCreateReqInfo{}

// FSLdapCreateReqInfo struct for FSLdapCreateReqInfo
type FSLdapCreateReqInfo struct {
	// ldap admin dn
	AdminDn *string `json:"admin_dn,omitempty"`
	// timeout for connection
	ConnectionTimeout *int64 `json:"connection_timeout,omitempty"`
	// dns server ip
	DnsIp *string `json:"dns_ip,omitempty"`
	// group suffix
	GroupSuffix *string `json:"group_suffix,omitempty"`
	// ip of server
	Ip string `json:"ip"`
	// ips of standby servers
	Ips []string `json:"ips,omitempty"`
	// name of ldap
	Name string `json:"name"`
	// bind password
	Password *string `json:"password,omitempty"`
	// ldap service port
	Port int64 `json:"port"`
	// ldap suffix
	Suffix string `json:"suffix"`
	// timeout for searching
	Timeout *int64 `json:"timeout,omitempty"`
	// user suffix
	UserSuffix *string `json:"user_suffix,omitempty"`
}

type _FSLdapCreateReqInfo FSLdapCreateReqInfo

// NewFSLdapCreateReqInfo instantiates a new FSLdapCreateReqInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFSLdapCreateReqInfo(ip string, name string, port int64, suffix string) *FSLdapCreateReqInfo {
	this := FSLdapCreateReqInfo{}
	this.Ip = ip
	this.Name = name
	this.Port = port
	this.Suffix = suffix
	return &this
}

// NewFSLdapCreateReqInfoWithDefaults instantiates a new FSLdapCreateReqInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFSLdapCreateReqInfoWithDefaults() *FSLdapCreateReqInfo {
	this := FSLdapCreateReqInfo{}
	return &this
}

// GetAdminDn returns the AdminDn field value if set, zero value otherwise.
func (o *FSLdapCreateReqInfo) GetAdminDn() string {
	if o == nil || IsNil(o.AdminDn) {
		var ret string
		return ret
	}
	return *o.AdminDn
}

// GetAdminDnOk returns a tuple with the AdminDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetAdminDnOk() (*string, bool) {
	if o == nil || IsNil(o.AdminDn) {
		return nil, false
	}
	return o.AdminDn, true
}

// HasAdminDn returns a boolean if a field has been set.
func (o *FSLdapCreateReqInfo) HasAdminDn() bool {
	if o != nil && !IsNil(o.AdminDn) {
		return true
	}

	return false
}

// SetAdminDn gets a reference to the given string and assigns it to the AdminDn field.
func (o *FSLdapCreateReqInfo) SetAdminDn(v string) {
	o.AdminDn = &v
}

// GetConnectionTimeout returns the ConnectionTimeout field value if set, zero value otherwise.
func (o *FSLdapCreateReqInfo) GetConnectionTimeout() int64 {
	if o == nil || IsNil(o.ConnectionTimeout) {
		var ret int64
		return ret
	}
	return *o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetConnectionTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ConnectionTimeout) {
		return nil, false
	}
	return o.ConnectionTimeout, true
}

// HasConnectionTimeout returns a boolean if a field has been set.
func (o *FSLdapCreateReqInfo) HasConnectionTimeout() bool {
	if o != nil && !IsNil(o.ConnectionTimeout) {
		return true
	}

	return false
}

// SetConnectionTimeout gets a reference to the given int64 and assigns it to the ConnectionTimeout field.
func (o *FSLdapCreateReqInfo) SetConnectionTimeout(v int64) {
	o.ConnectionTimeout = &v
}

// GetDnsIp returns the DnsIp field value if set, zero value otherwise.
func (o *FSLdapCreateReqInfo) GetDnsIp() string {
	if o == nil || IsNil(o.DnsIp) {
		var ret string
		return ret
	}
	return *o.DnsIp
}

// GetDnsIpOk returns a tuple with the DnsIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetDnsIpOk() (*string, bool) {
	if o == nil || IsNil(o.DnsIp) {
		return nil, false
	}
	return o.DnsIp, true
}

// HasDnsIp returns a boolean if a field has been set.
func (o *FSLdapCreateReqInfo) HasDnsIp() bool {
	if o != nil && !IsNil(o.DnsIp) {
		return true
	}

	return false
}

// SetDnsIp gets a reference to the given string and assigns it to the DnsIp field.
func (o *FSLdapCreateReqInfo) SetDnsIp(v string) {
	o.DnsIp = &v
}

// GetGroupSuffix returns the GroupSuffix field value if set, zero value otherwise.
func (o *FSLdapCreateReqInfo) GetGroupSuffix() string {
	if o == nil || IsNil(o.GroupSuffix) {
		var ret string
		return ret
	}
	return *o.GroupSuffix
}

// GetGroupSuffixOk returns a tuple with the GroupSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetGroupSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.GroupSuffix) {
		return nil, false
	}
	return o.GroupSuffix, true
}

// HasGroupSuffix returns a boolean if a field has been set.
func (o *FSLdapCreateReqInfo) HasGroupSuffix() bool {
	if o != nil && !IsNil(o.GroupSuffix) {
		return true
	}

	return false
}

// SetGroupSuffix gets a reference to the given string and assigns it to the GroupSuffix field.
func (o *FSLdapCreateReqInfo) SetGroupSuffix(v string) {
	o.GroupSuffix = &v
}

// GetIp returns the Ip field value
func (o *FSLdapCreateReqInfo) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *FSLdapCreateReqInfo) SetIp(v string) {
	o.Ip = v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *FSLdapCreateReqInfo) GetIps() []string {
	if o == nil || IsNil(o.Ips) {
		var ret []string
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *FSLdapCreateReqInfo) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *FSLdapCreateReqInfo) SetIps(v []string) {
	o.Ips = v
}

// GetName returns the Name field value
func (o *FSLdapCreateReqInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FSLdapCreateReqInfo) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *FSLdapCreateReqInfo) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *FSLdapCreateReqInfo) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *FSLdapCreateReqInfo) SetPassword(v string) {
	o.Password = &v
}

// GetPort returns the Port field value
func (o *FSLdapCreateReqInfo) GetPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *FSLdapCreateReqInfo) SetPort(v int64) {
	o.Port = v
}

// GetSuffix returns the Suffix field value
func (o *FSLdapCreateReqInfo) GetSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suffix, true
}

// SetSuffix sets field value
func (o *FSLdapCreateReqInfo) SetSuffix(v string) {
	o.Suffix = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *FSLdapCreateReqInfo) GetTimeout() int64 {
	if o == nil || IsNil(o.Timeout) {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *FSLdapCreateReqInfo) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *FSLdapCreateReqInfo) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetUserSuffix returns the UserSuffix field value if set, zero value otherwise.
func (o *FSLdapCreateReqInfo) GetUserSuffix() string {
	if o == nil || IsNil(o.UserSuffix) {
		var ret string
		return ret
	}
	return *o.UserSuffix
}

// GetUserSuffixOk returns a tuple with the UserSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdapCreateReqInfo) GetUserSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.UserSuffix) {
		return nil, false
	}
	return o.UserSuffix, true
}

// HasUserSuffix returns a boolean if a field has been set.
func (o *FSLdapCreateReqInfo) HasUserSuffix() bool {
	if o != nil && !IsNil(o.UserSuffix) {
		return true
	}

	return false
}

// SetUserSuffix gets a reference to the given string and assigns it to the UserSuffix field.
func (o *FSLdapCreateReqInfo) SetUserSuffix(v string) {
	o.UserSuffix = &v
}

func (o FSLdapCreateReqInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FSLdapCreateReqInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminDn) {
		toSerialize["admin_dn"] = o.AdminDn
	}
	if !IsNil(o.ConnectionTimeout) {
		toSerialize["connection_timeout"] = o.ConnectionTimeout
	}
	if !IsNil(o.DnsIp) {
		toSerialize["dns_ip"] = o.DnsIp
	}
	if !IsNil(o.GroupSuffix) {
		toSerialize["group_suffix"] = o.GroupSuffix
	}
	toSerialize["ip"] = o.Ip
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["port"] = o.Port
	toSerialize["suffix"] = o.Suffix
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.UserSuffix) {
		toSerialize["user_suffix"] = o.UserSuffix
	}
	return toSerialize, nil
}

func (o *FSLdapCreateReqInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip",
		"name",
		"port",
		"suffix",
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

	varFSLdapCreateReqInfo := _FSLdapCreateReqInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFSLdapCreateReqInfo)

	if err != nil {
		return err
	}

	*o = FSLdapCreateReqInfo(varFSLdapCreateReqInfo)

	return err
}

type NullableFSLdapCreateReqInfo struct {
	value *FSLdapCreateReqInfo
	isSet bool
}

func (v NullableFSLdapCreateReqInfo) Get() *FSLdapCreateReqInfo {
	return v.value
}

func (v *NullableFSLdapCreateReqInfo) Set(val *FSLdapCreateReqInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFSLdapCreateReqInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFSLdapCreateReqInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFSLdapCreateReqInfo(val *FSLdapCreateReqInfo) *NullableFSLdapCreateReqInfo {
	return &NullableFSLdapCreateReqInfo{value: val, isSet: true}
}

func (v NullableFSLdapCreateReqInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFSLdapCreateReqInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


