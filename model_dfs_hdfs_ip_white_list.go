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

// checks if the DfsHdfsIPWhiteList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsHdfsIPWhiteList{}

// DfsHdfsIPWhiteList DfsHdfsIPWhiteList define dfs hdfs IP white list +X:model:generate;order_by=-ID
type DfsHdfsIPWhiteList struct {
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DfsHdfs *DfsHdfsNestview `json:"dfs_hdfs,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Ips *string `json:"ips,omitempty"`
	Permission *string `json:"permission,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewDfsHdfsIPWhiteList instantiates a new DfsHdfsIPWhiteList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsHdfsIPWhiteList() *DfsHdfsIPWhiteList {
	this := DfsHdfsIPWhiteList{}
	return &this
}

// NewDfsHdfsIPWhiteListWithDefaults instantiates a new DfsHdfsIPWhiteList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsHdfsIPWhiteListWithDefaults() *DfsHdfsIPWhiteList {
	this := DfsHdfsIPWhiteList{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DfsHdfsIPWhiteList) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsIPWhiteList) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DfsHdfsIPWhiteList) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DfsHdfsIPWhiteList) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsHdfsIPWhiteList) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsIPWhiteList) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsHdfsIPWhiteList) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsHdfsIPWhiteList) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDfsHdfs returns the DfsHdfs field value if set, zero value otherwise.
func (o *DfsHdfsIPWhiteList) GetDfsHdfs() DfsHdfsNestview {
	if o == nil || IsNil(o.DfsHdfs) {
		var ret DfsHdfsNestview
		return ret
	}
	return *o.DfsHdfs
}

// GetDfsHdfsOk returns a tuple with the DfsHdfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsIPWhiteList) GetDfsHdfsOk() (*DfsHdfsNestview, bool) {
	if o == nil || IsNil(o.DfsHdfs) {
		return nil, false
	}
	return o.DfsHdfs, true
}

// HasDfsHdfs returns a boolean if a field has been set.
func (o *DfsHdfsIPWhiteList) HasDfsHdfs() bool {
	if o != nil && !IsNil(o.DfsHdfs) {
		return true
	}

	return false
}

// SetDfsHdfs gets a reference to the given DfsHdfsNestview and assigns it to the DfsHdfs field.
func (o *DfsHdfsIPWhiteList) SetDfsHdfs(v DfsHdfsNestview) {
	o.DfsHdfs = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsHdfsIPWhiteList) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsIPWhiteList) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsHdfsIPWhiteList) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsHdfsIPWhiteList) SetId(v int64) {
	o.Id = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *DfsHdfsIPWhiteList) GetIps() string {
	if o == nil || IsNil(o.Ips) {
		var ret string
		return ret
	}
	return *o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsIPWhiteList) GetIpsOk() (*string, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *DfsHdfsIPWhiteList) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given string and assigns it to the Ips field.
func (o *DfsHdfsIPWhiteList) SetIps(v string) {
	o.Ips = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *DfsHdfsIPWhiteList) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsIPWhiteList) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *DfsHdfsIPWhiteList) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *DfsHdfsIPWhiteList) SetPermission(v string) {
	o.Permission = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DfsHdfsIPWhiteList) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsIPWhiteList) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DfsHdfsIPWhiteList) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DfsHdfsIPWhiteList) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DfsHdfsIPWhiteList) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsHdfsIPWhiteList) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DfsHdfsIPWhiteList) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DfsHdfsIPWhiteList) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o DfsHdfsIPWhiteList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsHdfsIPWhiteList) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableDfsHdfsIPWhiteList struct {
	value *DfsHdfsIPWhiteList
	isSet bool
}

func (v NullableDfsHdfsIPWhiteList) Get() *DfsHdfsIPWhiteList {
	return v.value
}

func (v *NullableDfsHdfsIPWhiteList) Set(val *DfsHdfsIPWhiteList) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsHdfsIPWhiteList) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsHdfsIPWhiteList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsHdfsIPWhiteList(val *DfsHdfsIPWhiteList) *NullableDfsHdfsIPWhiteList {
	return &NullableDfsHdfsIPWhiteList{value: val, isSet: true}
}

func (v NullableDfsHdfsIPWhiteList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsHdfsIPWhiteList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


