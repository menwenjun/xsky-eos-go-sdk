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

// checks if the DfsHdfsProxyUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsHdfsProxyUser{}

// DfsHdfsProxyUser DfsHdfsProxyUser define dfs hdfs proxy user list +X:model:generate;etcd_lock;
type DfsHdfsProxyUser struct {
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DfsHdfs *DfsHdfsNestview `json:"dfs_hdfs,omitempty"`
	Host *string `json:"host,omitempty"`
	Id *int64 `json:"id,omitempty"`
	ProxyUserName *string `json:"proxy_user_name,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	User *string `json:"user,omitempty"`
	UserGroup *string `json:"user_group,omitempty"`
}

// NewDfsHdfsProxyUser instantiates a new DfsHdfsProxyUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsHdfsProxyUser() *DfsHdfsProxyUser {
	this := DfsHdfsProxyUser{}
	return &this
}

// NewDfsHdfsProxyUserWithDefaults instantiates a new DfsHdfsProxyUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsHdfsProxyUserWithDefaults() *DfsHdfsProxyUser {
	this := DfsHdfsProxyUser{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DfsHdfsProxyUser) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsProxyUser) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DfsHdfsProxyUser) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DfsHdfsProxyUser) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsHdfsProxyUser) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsProxyUser) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsHdfsProxyUser) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsHdfsProxyUser) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDfsHdfs returns the DfsHdfs field value if set, zero value otherwise.
func (o *DfsHdfsProxyUser) GetDfsHdfs() DfsHdfsNestview {
	if o == nil || IsNil(o.DfsHdfs) {
		var ret DfsHdfsNestview
		return ret
	}
	return *o.DfsHdfs
}

// GetDfsHdfsOk returns a tuple with the DfsHdfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsProxyUser) GetDfsHdfsOk() (*DfsHdfsNestview, bool) {
	if o == nil || IsNil(o.DfsHdfs) {
		return nil, false
	}
	return o.DfsHdfs, true
}

// HasDfsHdfs returns a boolean if a field has been set.
func (o *DfsHdfsProxyUser) HasDfsHdfs() bool {
	if o != nil && !IsNil(o.DfsHdfs) {
		return true
	}

	return false
}

// SetDfsHdfs gets a reference to the given DfsHdfsNestview and assigns it to the DfsHdfs field.
func (o *DfsHdfsProxyUser) SetDfsHdfs(v DfsHdfsNestview) {
	o.DfsHdfs = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *DfsHdfsProxyUser) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsProxyUser) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *DfsHdfsProxyUser) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *DfsHdfsProxyUser) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsHdfsProxyUser) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsProxyUser) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsHdfsProxyUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsHdfsProxyUser) SetId(v int64) {
	o.Id = &v
}

// GetProxyUserName returns the ProxyUserName field value if set, zero value otherwise.
func (o *DfsHdfsProxyUser) GetProxyUserName() string {
	if o == nil || IsNil(o.ProxyUserName) {
		var ret string
		return ret
	}
	return *o.ProxyUserName
}

// GetProxyUserNameOk returns a tuple with the ProxyUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsProxyUser) GetProxyUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyUserName) {
		return nil, false
	}
	return o.ProxyUserName, true
}

// HasProxyUserName returns a boolean if a field has been set.
func (o *DfsHdfsProxyUser) HasProxyUserName() bool {
	if o != nil && !IsNil(o.ProxyUserName) {
		return true
	}

	return false
}

// SetProxyUserName gets a reference to the given string and assigns it to the ProxyUserName field.
func (o *DfsHdfsProxyUser) SetProxyUserName(v string) {
	o.ProxyUserName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DfsHdfsProxyUser) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsProxyUser) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DfsHdfsProxyUser) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DfsHdfsProxyUser) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DfsHdfsProxyUser) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsProxyUser) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DfsHdfsProxyUser) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DfsHdfsProxyUser) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *DfsHdfsProxyUser) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsProxyUser) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *DfsHdfsProxyUser) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *DfsHdfsProxyUser) SetUser(v string) {
	o.User = &v
}

// GetUserGroup returns the UserGroup field value if set, zero value otherwise.
func (o *DfsHdfsProxyUser) GetUserGroup() string {
	if o == nil || IsNil(o.UserGroup) {
		var ret string
		return ret
	}
	return *o.UserGroup
}

// GetUserGroupOk returns a tuple with the UserGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsProxyUser) GetUserGroupOk() (*string, bool) {
	if o == nil || IsNil(o.UserGroup) {
		return nil, false
	}
	return o.UserGroup, true
}

// HasUserGroup returns a boolean if a field has been set.
func (o *DfsHdfsProxyUser) HasUserGroup() bool {
	if o != nil && !IsNil(o.UserGroup) {
		return true
	}

	return false
}

// SetUserGroup gets a reference to the given string and assigns it to the UserGroup field.
func (o *DfsHdfsProxyUser) SetUserGroup(v string) {
	o.UserGroup = &v
}

func (o DfsHdfsProxyUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsHdfsProxyUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DfsHdfs) {
		toSerialize["dfs_hdfs"] = o.DfsHdfs
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProxyUserName) {
		toSerialize["proxy_user_name"] = o.ProxyUserName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.UserGroup) {
		toSerialize["user_group"] = o.UserGroup
	}
	return toSerialize, nil
}

type NullableDfsHdfsProxyUser struct {
	value *DfsHdfsProxyUser
	isSet bool
}

func (v NullableDfsHdfsProxyUser) Get() *DfsHdfsProxyUser {
	return v.value
}

func (v *NullableDfsHdfsProxyUser) Set(val *DfsHdfsProxyUser) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsHdfsProxyUser) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsHdfsProxyUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsHdfsProxyUser(val *DfsHdfsProxyUser) *NullableDfsHdfsProxyUser {
	return &NullableDfsHdfsProxyUser{value: val, isSet: true}
}

func (v NullableDfsHdfsProxyUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsHdfsProxyUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


