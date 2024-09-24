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

// checks if the FSLdap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FSLdap{}

// FSLdap FSLdap defines model of file storage ldap
type FSLdap struct {
	ActionStatus *string `json:"action_status,omitempty"`
	AdminDn *string `json:"admin_dn,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	ConnectionTimeout *int64 `json:"connection_timeout,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DfsGatewayGroupNum *int64 `json:"dfs_gateway_group_num,omitempty"`
	DnsIp *string `json:"dns_ip,omitempty"`
	GroupSuffix *string `json:"group_suffix,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Ips []string `json:"ips,omitempty"`
	Name *string `json:"name,omitempty"`
	NetbiosName *string `json:"netbios_name,omitempty"`
	Port *int64 `json:"port,omitempty"`
	SambaSid *string `json:"samba_sid,omitempty"`
	Status *string `json:"status,omitempty"`
	Suffix *string `json:"suffix,omitempty"`
	Timeout *int64 `json:"timeout,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	UserSuffix *string `json:"user_suffix,omitempty"`
	VerifyTime *time.Time `json:"verify_time,omitempty"`
}

// NewFSLdap instantiates a new FSLdap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFSLdap() *FSLdap {
	this := FSLdap{}
	return &this
}

// NewFSLdapWithDefaults instantiates a new FSLdap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFSLdapWithDefaults() *FSLdap {
	this := FSLdap{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *FSLdap) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *FSLdap) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *FSLdap) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetAdminDn returns the AdminDn field value if set, zero value otherwise.
func (o *FSLdap) GetAdminDn() string {
	if o == nil || IsNil(o.AdminDn) {
		var ret string
		return ret
	}
	return *o.AdminDn
}

// GetAdminDnOk returns a tuple with the AdminDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetAdminDnOk() (*string, bool) {
	if o == nil || IsNil(o.AdminDn) {
		return nil, false
	}
	return o.AdminDn, true
}

// HasAdminDn returns a boolean if a field has been set.
func (o *FSLdap) HasAdminDn() bool {
	if o != nil && !IsNil(o.AdminDn) {
		return true
	}

	return false
}

// SetAdminDn gets a reference to the given string and assigns it to the AdminDn field.
func (o *FSLdap) SetAdminDn(v string) {
	o.AdminDn = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *FSLdap) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *FSLdap) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *FSLdap) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetConnectionTimeout returns the ConnectionTimeout field value if set, zero value otherwise.
func (o *FSLdap) GetConnectionTimeout() int64 {
	if o == nil || IsNil(o.ConnectionTimeout) {
		var ret int64
		return ret
	}
	return *o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetConnectionTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ConnectionTimeout) {
		return nil, false
	}
	return o.ConnectionTimeout, true
}

// HasConnectionTimeout returns a boolean if a field has been set.
func (o *FSLdap) HasConnectionTimeout() bool {
	if o != nil && !IsNil(o.ConnectionTimeout) {
		return true
	}

	return false
}

// SetConnectionTimeout gets a reference to the given int64 and assigns it to the ConnectionTimeout field.
func (o *FSLdap) SetConnectionTimeout(v int64) {
	o.ConnectionTimeout = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *FSLdap) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *FSLdap) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *FSLdap) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDfsGatewayGroupNum returns the DfsGatewayGroupNum field value if set, zero value otherwise.
func (o *FSLdap) GetDfsGatewayGroupNum() int64 {
	if o == nil || IsNil(o.DfsGatewayGroupNum) {
		var ret int64
		return ret
	}
	return *o.DfsGatewayGroupNum
}

// GetDfsGatewayGroupNumOk returns a tuple with the DfsGatewayGroupNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetDfsGatewayGroupNumOk() (*int64, bool) {
	if o == nil || IsNil(o.DfsGatewayGroupNum) {
		return nil, false
	}
	return o.DfsGatewayGroupNum, true
}

// HasDfsGatewayGroupNum returns a boolean if a field has been set.
func (o *FSLdap) HasDfsGatewayGroupNum() bool {
	if o != nil && !IsNil(o.DfsGatewayGroupNum) {
		return true
	}

	return false
}

// SetDfsGatewayGroupNum gets a reference to the given int64 and assigns it to the DfsGatewayGroupNum field.
func (o *FSLdap) SetDfsGatewayGroupNum(v int64) {
	o.DfsGatewayGroupNum = &v
}

// GetDnsIp returns the DnsIp field value if set, zero value otherwise.
func (o *FSLdap) GetDnsIp() string {
	if o == nil || IsNil(o.DnsIp) {
		var ret string
		return ret
	}
	return *o.DnsIp
}

// GetDnsIpOk returns a tuple with the DnsIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetDnsIpOk() (*string, bool) {
	if o == nil || IsNil(o.DnsIp) {
		return nil, false
	}
	return o.DnsIp, true
}

// HasDnsIp returns a boolean if a field has been set.
func (o *FSLdap) HasDnsIp() bool {
	if o != nil && !IsNil(o.DnsIp) {
		return true
	}

	return false
}

// SetDnsIp gets a reference to the given string and assigns it to the DnsIp field.
func (o *FSLdap) SetDnsIp(v string) {
	o.DnsIp = &v
}

// GetGroupSuffix returns the GroupSuffix field value if set, zero value otherwise.
func (o *FSLdap) GetGroupSuffix() string {
	if o == nil || IsNil(o.GroupSuffix) {
		var ret string
		return ret
	}
	return *o.GroupSuffix
}

// GetGroupSuffixOk returns a tuple with the GroupSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetGroupSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.GroupSuffix) {
		return nil, false
	}
	return o.GroupSuffix, true
}

// HasGroupSuffix returns a boolean if a field has been set.
func (o *FSLdap) HasGroupSuffix() bool {
	if o != nil && !IsNil(o.GroupSuffix) {
		return true
	}

	return false
}

// SetGroupSuffix gets a reference to the given string and assigns it to the GroupSuffix field.
func (o *FSLdap) SetGroupSuffix(v string) {
	o.GroupSuffix = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FSLdap) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FSLdap) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FSLdap) SetId(v int64) {
	o.Id = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *FSLdap) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *FSLdap) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *FSLdap) SetIp(v string) {
	o.Ip = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *FSLdap) GetIps() []string {
	if o == nil || IsNil(o.Ips) {
		var ret []string
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *FSLdap) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *FSLdap) SetIps(v []string) {
	o.Ips = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FSLdap) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FSLdap) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FSLdap) SetName(v string) {
	o.Name = &v
}

// GetNetbiosName returns the NetbiosName field value if set, zero value otherwise.
func (o *FSLdap) GetNetbiosName() string {
	if o == nil || IsNil(o.NetbiosName) {
		var ret string
		return ret
	}
	return *o.NetbiosName
}

// GetNetbiosNameOk returns a tuple with the NetbiosName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetNetbiosNameOk() (*string, bool) {
	if o == nil || IsNil(o.NetbiosName) {
		return nil, false
	}
	return o.NetbiosName, true
}

// HasNetbiosName returns a boolean if a field has been set.
func (o *FSLdap) HasNetbiosName() bool {
	if o != nil && !IsNil(o.NetbiosName) {
		return true
	}

	return false
}

// SetNetbiosName gets a reference to the given string and assigns it to the NetbiosName field.
func (o *FSLdap) SetNetbiosName(v string) {
	o.NetbiosName = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *FSLdap) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *FSLdap) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *FSLdap) SetPort(v int64) {
	o.Port = &v
}

// GetSambaSid returns the SambaSid field value if set, zero value otherwise.
func (o *FSLdap) GetSambaSid() string {
	if o == nil || IsNil(o.SambaSid) {
		var ret string
		return ret
	}
	return *o.SambaSid
}

// GetSambaSidOk returns a tuple with the SambaSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetSambaSidOk() (*string, bool) {
	if o == nil || IsNil(o.SambaSid) {
		return nil, false
	}
	return o.SambaSid, true
}

// HasSambaSid returns a boolean if a field has been set.
func (o *FSLdap) HasSambaSid() bool {
	if o != nil && !IsNil(o.SambaSid) {
		return true
	}

	return false
}

// SetSambaSid gets a reference to the given string and assigns it to the SambaSid field.
func (o *FSLdap) SetSambaSid(v string) {
	o.SambaSid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FSLdap) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FSLdap) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FSLdap) SetStatus(v string) {
	o.Status = &v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise.
func (o *FSLdap) GetSuffix() string {
	if o == nil || IsNil(o.Suffix) {
		var ret string
		return ret
	}
	return *o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.Suffix) {
		return nil, false
	}
	return o.Suffix, true
}

// HasSuffix returns a boolean if a field has been set.
func (o *FSLdap) HasSuffix() bool {
	if o != nil && !IsNil(o.Suffix) {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given string and assigns it to the Suffix field.
func (o *FSLdap) SetSuffix(v string) {
	o.Suffix = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *FSLdap) GetTimeout() int64 {
	if o == nil || IsNil(o.Timeout) {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *FSLdap) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *FSLdap) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *FSLdap) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *FSLdap) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *FSLdap) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUserSuffix returns the UserSuffix field value if set, zero value otherwise.
func (o *FSLdap) GetUserSuffix() string {
	if o == nil || IsNil(o.UserSuffix) {
		var ret string
		return ret
	}
	return *o.UserSuffix
}

// GetUserSuffixOk returns a tuple with the UserSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetUserSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.UserSuffix) {
		return nil, false
	}
	return o.UserSuffix, true
}

// HasUserSuffix returns a boolean if a field has been set.
func (o *FSLdap) HasUserSuffix() bool {
	if o != nil && !IsNil(o.UserSuffix) {
		return true
	}

	return false
}

// SetUserSuffix gets a reference to the given string and assigns it to the UserSuffix field.
func (o *FSLdap) SetUserSuffix(v string) {
	o.UserSuffix = &v
}

// GetVerifyTime returns the VerifyTime field value if set, zero value otherwise.
func (o *FSLdap) GetVerifyTime() time.Time {
	if o == nil || IsNil(o.VerifyTime) {
		var ret time.Time
		return ret
	}
	return *o.VerifyTime
}

// GetVerifyTimeOk returns a tuple with the VerifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSLdap) GetVerifyTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VerifyTime) {
		return nil, false
	}
	return o.VerifyTime, true
}

// HasVerifyTime returns a boolean if a field has been set.
func (o *FSLdap) HasVerifyTime() bool {
	if o != nil && !IsNil(o.VerifyTime) {
		return true
	}

	return false
}

// SetVerifyTime gets a reference to the given time.Time and assigns it to the VerifyTime field.
func (o *FSLdap) SetVerifyTime(v time.Time) {
	o.VerifyTime = &v
}

func (o FSLdap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FSLdap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.AdminDn) {
		toSerialize["admin_dn"] = o.AdminDn
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.ConnectionTimeout) {
		toSerialize["connection_timeout"] = o.ConnectionTimeout
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DfsGatewayGroupNum) {
		toSerialize["dfs_gateway_group_num"] = o.DfsGatewayGroupNum
	}
	if !IsNil(o.DnsIp) {
		toSerialize["dns_ip"] = o.DnsIp
	}
	if !IsNil(o.GroupSuffix) {
		toSerialize["group_suffix"] = o.GroupSuffix
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NetbiosName) {
		toSerialize["netbios_name"] = o.NetbiosName
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.SambaSid) {
		toSerialize["samba_sid"] = o.SambaSid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Suffix) {
		toSerialize["suffix"] = o.Suffix
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.UserSuffix) {
		toSerialize["user_suffix"] = o.UserSuffix
	}
	if !IsNil(o.VerifyTime) {
		toSerialize["verify_time"] = o.VerifyTime
	}
	return toSerialize, nil
}

type NullableFSLdap struct {
	value *FSLdap
	isSet bool
}

func (v NullableFSLdap) Get() *FSLdap {
	return v.value
}

func (v *NullableFSLdap) Set(val *FSLdap) {
	v.value = val
	v.isSet = true
}

func (v NullableFSLdap) IsSet() bool {
	return v.isSet
}

func (v *NullableFSLdap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFSLdap(val *FSLdap) *NullableFSLdap {
	return &NullableFSLdap{value: val, isSet: true}
}

func (v NullableFSLdap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFSLdap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


