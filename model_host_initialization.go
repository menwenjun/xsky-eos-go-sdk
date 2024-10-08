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

// checks if the HostInitialization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostInitialization{}

// HostInitialization HostInitialization is the model of host initialization +X:model:etcd_lock;
type HostInitialization struct {
	AdminIps []string `json:"admin_ips,omitempty"`
	Cluster *Cluster `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DisableFirewalld *bool `json:"disable_firewalld,omitempty"`
	HostnamePrefix *string `json:"hostname_prefix,omitempty"`
	HostnameSuffixLen *int64 `json:"hostname_suffix_len,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Message *string `json:"message,omitempty"`
	SetChrony *bool `json:"set_chrony,omitempty"`
	SetHostname *bool `json:"set_hostname,omitempty"`
	SshUser *string `json:"ssh_user,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewHostInitialization instantiates a new HostInitialization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostInitialization() *HostInitialization {
	this := HostInitialization{}
	return &this
}

// NewHostInitializationWithDefaults instantiates a new HostInitialization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostInitializationWithDefaults() *HostInitialization {
	this := HostInitialization{}
	return &this
}

// GetAdminIps returns the AdminIps field value if set, zero value otherwise.
func (o *HostInitialization) GetAdminIps() []string {
	if o == nil || IsNil(o.AdminIps) {
		var ret []string
		return ret
	}
	return o.AdminIps
}

// GetAdminIpsOk returns a tuple with the AdminIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetAdminIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdminIps) {
		return nil, false
	}
	return o.AdminIps, true
}

// HasAdminIps returns a boolean if a field has been set.
func (o *HostInitialization) HasAdminIps() bool {
	if o != nil && !IsNil(o.AdminIps) {
		return true
	}

	return false
}

// SetAdminIps gets a reference to the given []string and assigns it to the AdminIps field.
func (o *HostInitialization) SetAdminIps(v []string) {
	o.AdminIps = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HostInitialization) GetCluster() Cluster {
	if o == nil || IsNil(o.Cluster) {
		var ret Cluster
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetClusterOk() (*Cluster, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HostInitialization) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given Cluster and assigns it to the Cluster field.
func (o *HostInitialization) SetCluster(v Cluster) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *HostInitialization) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *HostInitialization) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *HostInitialization) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDisableFirewalld returns the DisableFirewalld field value if set, zero value otherwise.
func (o *HostInitialization) GetDisableFirewalld() bool {
	if o == nil || IsNil(o.DisableFirewalld) {
		var ret bool
		return ret
	}
	return *o.DisableFirewalld
}

// GetDisableFirewalldOk returns a tuple with the DisableFirewalld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetDisableFirewalldOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableFirewalld) {
		return nil, false
	}
	return o.DisableFirewalld, true
}

// HasDisableFirewalld returns a boolean if a field has been set.
func (o *HostInitialization) HasDisableFirewalld() bool {
	if o != nil && !IsNil(o.DisableFirewalld) {
		return true
	}

	return false
}

// SetDisableFirewalld gets a reference to the given bool and assigns it to the DisableFirewalld field.
func (o *HostInitialization) SetDisableFirewalld(v bool) {
	o.DisableFirewalld = &v
}

// GetHostnamePrefix returns the HostnamePrefix field value if set, zero value otherwise.
func (o *HostInitialization) GetHostnamePrefix() string {
	if o == nil || IsNil(o.HostnamePrefix) {
		var ret string
		return ret
	}
	return *o.HostnamePrefix
}

// GetHostnamePrefixOk returns a tuple with the HostnamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetHostnamePrefixOk() (*string, bool) {
	if o == nil || IsNil(o.HostnamePrefix) {
		return nil, false
	}
	return o.HostnamePrefix, true
}

// HasHostnamePrefix returns a boolean if a field has been set.
func (o *HostInitialization) HasHostnamePrefix() bool {
	if o != nil && !IsNil(o.HostnamePrefix) {
		return true
	}

	return false
}

// SetHostnamePrefix gets a reference to the given string and assigns it to the HostnamePrefix field.
func (o *HostInitialization) SetHostnamePrefix(v string) {
	o.HostnamePrefix = &v
}

// GetHostnameSuffixLen returns the HostnameSuffixLen field value if set, zero value otherwise.
func (o *HostInitialization) GetHostnameSuffixLen() int64 {
	if o == nil || IsNil(o.HostnameSuffixLen) {
		var ret int64
		return ret
	}
	return *o.HostnameSuffixLen
}

// GetHostnameSuffixLenOk returns a tuple with the HostnameSuffixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetHostnameSuffixLenOk() (*int64, bool) {
	if o == nil || IsNil(o.HostnameSuffixLen) {
		return nil, false
	}
	return o.HostnameSuffixLen, true
}

// HasHostnameSuffixLen returns a boolean if a field has been set.
func (o *HostInitialization) HasHostnameSuffixLen() bool {
	if o != nil && !IsNil(o.HostnameSuffixLen) {
		return true
	}

	return false
}

// SetHostnameSuffixLen gets a reference to the given int64 and assigns it to the HostnameSuffixLen field.
func (o *HostInitialization) SetHostnameSuffixLen(v int64) {
	o.HostnameSuffixLen = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HostInitialization) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HostInitialization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *HostInitialization) SetId(v int64) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HostInitialization) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HostInitialization) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HostInitialization) SetMessage(v string) {
	o.Message = &v
}

// GetSetChrony returns the SetChrony field value if set, zero value otherwise.
func (o *HostInitialization) GetSetChrony() bool {
	if o == nil || IsNil(o.SetChrony) {
		var ret bool
		return ret
	}
	return *o.SetChrony
}

// GetSetChronyOk returns a tuple with the SetChrony field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetSetChronyOk() (*bool, bool) {
	if o == nil || IsNil(o.SetChrony) {
		return nil, false
	}
	return o.SetChrony, true
}

// HasSetChrony returns a boolean if a field has been set.
func (o *HostInitialization) HasSetChrony() bool {
	if o != nil && !IsNil(o.SetChrony) {
		return true
	}

	return false
}

// SetSetChrony gets a reference to the given bool and assigns it to the SetChrony field.
func (o *HostInitialization) SetSetChrony(v bool) {
	o.SetChrony = &v
}

// GetSetHostname returns the SetHostname field value if set, zero value otherwise.
func (o *HostInitialization) GetSetHostname() bool {
	if o == nil || IsNil(o.SetHostname) {
		var ret bool
		return ret
	}
	return *o.SetHostname
}

// GetSetHostnameOk returns a tuple with the SetHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetSetHostnameOk() (*bool, bool) {
	if o == nil || IsNil(o.SetHostname) {
		return nil, false
	}
	return o.SetHostname, true
}

// HasSetHostname returns a boolean if a field has been set.
func (o *HostInitialization) HasSetHostname() bool {
	if o != nil && !IsNil(o.SetHostname) {
		return true
	}

	return false
}

// SetSetHostname gets a reference to the given bool and assigns it to the SetHostname field.
func (o *HostInitialization) SetSetHostname(v bool) {
	o.SetHostname = &v
}

// GetSshUser returns the SshUser field value if set, zero value otherwise.
func (o *HostInitialization) GetSshUser() string {
	if o == nil || IsNil(o.SshUser) {
		var ret string
		return ret
	}
	return *o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetSshUserOk() (*string, bool) {
	if o == nil || IsNil(o.SshUser) {
		return nil, false
	}
	return o.SshUser, true
}

// HasSshUser returns a boolean if a field has been set.
func (o *HostInitialization) HasSshUser() bool {
	if o != nil && !IsNil(o.SshUser) {
		return true
	}

	return false
}

// SetSshUser gets a reference to the given string and assigns it to the SshUser field.
func (o *HostInitialization) SetSshUser(v string) {
	o.SshUser = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HostInitialization) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostInitialization) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HostInitialization) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HostInitialization) SetStatus(v string) {
	o.Status = &v
}

func (o HostInitialization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostInitialization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminIps) {
		toSerialize["admin_ips"] = o.AdminIps
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DisableFirewalld) {
		toSerialize["disable_firewalld"] = o.DisableFirewalld
	}
	if !IsNil(o.HostnamePrefix) {
		toSerialize["hostname_prefix"] = o.HostnamePrefix
	}
	if !IsNil(o.HostnameSuffixLen) {
		toSerialize["hostname_suffix_len"] = o.HostnameSuffixLen
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.SetChrony) {
		toSerialize["set_chrony"] = o.SetChrony
	}
	if !IsNil(o.SetHostname) {
		toSerialize["set_hostname"] = o.SetHostname
	}
	if !IsNil(o.SshUser) {
		toSerialize["ssh_user"] = o.SshUser
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableHostInitialization struct {
	value *HostInitialization
	isSet bool
}

func (v NullableHostInitialization) Get() *HostInitialization {
	return v.value
}

func (v *NullableHostInitialization) Set(val *HostInitialization) {
	v.value = val
	v.isSet = true
}

func (v NullableHostInitialization) IsSet() bool {
	return v.isSet
}

func (v *NullableHostInitialization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostInitialization(val *HostInitialization) *NullableHostInitialization {
	return &NullableHostInitialization{value: val, isSet: true}
}

func (v NullableHostInitialization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostInitialization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


