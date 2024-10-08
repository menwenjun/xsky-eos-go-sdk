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

// checks if the DfsAuditLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsAuditLog{}

// DfsAuditLog DfsAuditLog defines model of dfs audit log +X:model:generate;order_by=-ID
type DfsAuditLog struct {
	Actions []string `json:"actions,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Id *int64 `json:"id,omitempty"`
	LogPath *string `json:"log_path,omitempty"`
	Rootfs *DfsRootfsNestview `json:"rootfs,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewDfsAuditLog instantiates a new DfsAuditLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsAuditLog() *DfsAuditLog {
	this := DfsAuditLog{}
	return &this
}

// NewDfsAuditLogWithDefaults instantiates a new DfsAuditLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsAuditLogWithDefaults() *DfsAuditLog {
	this := DfsAuditLog{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *DfsAuditLog) GetActions() []string {
	if o == nil || IsNil(o.Actions) {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsAuditLog) GetActionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *DfsAuditLog) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *DfsAuditLog) SetActions(v []string) {
	o.Actions = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DfsAuditLog) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsAuditLog) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DfsAuditLog) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DfsAuditLog) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsAuditLog) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsAuditLog) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsAuditLog) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsAuditLog) SetCreate(v time.Time) {
	o.Create = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsAuditLog) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsAuditLog) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsAuditLog) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsAuditLog) SetId(v int64) {
	o.Id = &v
}

// GetLogPath returns the LogPath field value if set, zero value otherwise.
func (o *DfsAuditLog) GetLogPath() string {
	if o == nil || IsNil(o.LogPath) {
		var ret string
		return ret
	}
	return *o.LogPath
}

// GetLogPathOk returns a tuple with the LogPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsAuditLog) GetLogPathOk() (*string, bool) {
	if o == nil || IsNil(o.LogPath) {
		return nil, false
	}
	return o.LogPath, true
}

// HasLogPath returns a boolean if a field has been set.
func (o *DfsAuditLog) HasLogPath() bool {
	if o != nil && !IsNil(o.LogPath) {
		return true
	}

	return false
}

// SetLogPath gets a reference to the given string and assigns it to the LogPath field.
func (o *DfsAuditLog) SetLogPath(v string) {
	o.LogPath = &v
}

// GetRootfs returns the Rootfs field value if set, zero value otherwise.
func (o *DfsAuditLog) GetRootfs() DfsRootfsNestview {
	if o == nil || IsNil(o.Rootfs) {
		var ret DfsRootfsNestview
		return ret
	}
	return *o.Rootfs
}

// GetRootfsOk returns a tuple with the Rootfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsAuditLog) GetRootfsOk() (*DfsRootfsNestview, bool) {
	if o == nil || IsNil(o.Rootfs) {
		return nil, false
	}
	return o.Rootfs, true
}

// HasRootfs returns a boolean if a field has been set.
func (o *DfsAuditLog) HasRootfs() bool {
	if o != nil && !IsNil(o.Rootfs) {
		return true
	}

	return false
}

// SetRootfs gets a reference to the given DfsRootfsNestview and assigns it to the Rootfs field.
func (o *DfsAuditLog) SetRootfs(v DfsRootfsNestview) {
	o.Rootfs = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DfsAuditLog) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsAuditLog) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DfsAuditLog) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *DfsAuditLog) SetSize(v int64) {
	o.Size = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DfsAuditLog) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsAuditLog) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DfsAuditLog) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DfsAuditLog) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DfsAuditLog) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsAuditLog) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DfsAuditLog) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DfsAuditLog) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o DfsAuditLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsAuditLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LogPath) {
		toSerialize["log_path"] = o.LogPath
	}
	if !IsNil(o.Rootfs) {
		toSerialize["rootfs"] = o.Rootfs
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableDfsAuditLog struct {
	value *DfsAuditLog
	isSet bool
}

func (v NullableDfsAuditLog) Get() *DfsAuditLog {
	return v.value
}

func (v *NullableDfsAuditLog) Set(val *DfsAuditLog) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsAuditLog) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsAuditLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsAuditLog(val *DfsAuditLog) *NullableDfsAuditLog {
	return &NullableDfsAuditLog{value: val, isSet: true}
}

func (v NullableDfsAuditLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsAuditLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


