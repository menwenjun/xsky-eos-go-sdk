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

// checks if the DfsPathNestviewDfsPathPerformance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsPathNestviewDfsPathPerformance{}

// DfsPathNestviewDfsPathPerformance struct for DfsPathNestviewDfsPathPerformance
type DfsPathNestviewDfsPathPerformance struct {
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Path *DfsPathNestview `json:"path,omitempty"`
	Rootfs *DfsRootfsNestview `json:"rootfs,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewDfsPathNestviewDfsPathPerformance instantiates a new DfsPathNestviewDfsPathPerformance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsPathNestviewDfsPathPerformance() *DfsPathNestviewDfsPathPerformance {
	this := DfsPathNestviewDfsPathPerformance{}
	return &this
}

// NewDfsPathNestviewDfsPathPerformanceWithDefaults instantiates a new DfsPathNestviewDfsPathPerformance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsPathNestviewDfsPathPerformanceWithDefaults() *DfsPathNestviewDfsPathPerformance {
	this := DfsPathNestviewDfsPathPerformance{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DfsPathNestviewDfsPathPerformance) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathNestviewDfsPathPerformance) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DfsPathNestviewDfsPathPerformance) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DfsPathNestviewDfsPathPerformance) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsPathNestviewDfsPathPerformance) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathNestviewDfsPathPerformance) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsPathNestviewDfsPathPerformance) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsPathNestviewDfsPathPerformance) SetCreate(v time.Time) {
	o.Create = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsPathNestviewDfsPathPerformance) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathNestviewDfsPathPerformance) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsPathNestviewDfsPathPerformance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsPathNestviewDfsPathPerformance) SetId(v int64) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DfsPathNestviewDfsPathPerformance) GetPath() DfsPathNestview {
	if o == nil || IsNil(o.Path) {
		var ret DfsPathNestview
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathNestviewDfsPathPerformance) GetPathOk() (*DfsPathNestview, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DfsPathNestviewDfsPathPerformance) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given DfsPathNestview and assigns it to the Path field.
func (o *DfsPathNestviewDfsPathPerformance) SetPath(v DfsPathNestview) {
	o.Path = &v
}

// GetRootfs returns the Rootfs field value if set, zero value otherwise.
func (o *DfsPathNestviewDfsPathPerformance) GetRootfs() DfsRootfsNestview {
	if o == nil || IsNil(o.Rootfs) {
		var ret DfsRootfsNestview
		return ret
	}
	return *o.Rootfs
}

// GetRootfsOk returns a tuple with the Rootfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathNestviewDfsPathPerformance) GetRootfsOk() (*DfsRootfsNestview, bool) {
	if o == nil || IsNil(o.Rootfs) {
		return nil, false
	}
	return o.Rootfs, true
}

// HasRootfs returns a boolean if a field has been set.
func (o *DfsPathNestviewDfsPathPerformance) HasRootfs() bool {
	if o != nil && !IsNil(o.Rootfs) {
		return true
	}

	return false
}

// SetRootfs gets a reference to the given DfsRootfsNestview and assigns it to the Rootfs field.
func (o *DfsPathNestviewDfsPathPerformance) SetRootfs(v DfsRootfsNestview) {
	o.Rootfs = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DfsPathNestviewDfsPathPerformance) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathNestviewDfsPathPerformance) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DfsPathNestviewDfsPathPerformance) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DfsPathNestviewDfsPathPerformance) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DfsPathNestviewDfsPathPerformance) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsPathNestviewDfsPathPerformance) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DfsPathNestviewDfsPathPerformance) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DfsPathNestviewDfsPathPerformance) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o DfsPathNestviewDfsPathPerformance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsPathNestviewDfsPathPerformance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Rootfs) {
		toSerialize["rootfs"] = o.Rootfs
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableDfsPathNestviewDfsPathPerformance struct {
	value *DfsPathNestviewDfsPathPerformance
	isSet bool
}

func (v NullableDfsPathNestviewDfsPathPerformance) Get() *DfsPathNestviewDfsPathPerformance {
	return v.value
}

func (v *NullableDfsPathNestviewDfsPathPerformance) Set(val *DfsPathNestviewDfsPathPerformance) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsPathNestviewDfsPathPerformance) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsPathNestviewDfsPathPerformance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsPathNestviewDfsPathPerformance(val *DfsPathNestviewDfsPathPerformance) *NullableDfsPathNestviewDfsPathPerformance {
	return &NullableDfsPathNestviewDfsPathPerformance{value: val, isSet: true}
}

func (v NullableDfsPathNestviewDfsPathPerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsPathNestviewDfsPathPerformance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


