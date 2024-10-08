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

// checks if the OSReplicationPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSReplicationPath{}

// OSReplicationPath OSReplicationPath defines models of os replication path.
type OSReplicationPath struct {
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	OsZoneUuids []string `json:"os_zone_uuids,omitempty"`
	OsZones []ObjectStorageZoneNestview `json:"os_zones,omitempty"`
	ReplicationUuid *string `json:"replication_uuid,omitempty"`
	Status *string `json:"status,omitempty"`
	Suspended *bool `json:"suspended,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewOSReplicationPath instantiates a new OSReplicationPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSReplicationPath() *OSReplicationPath {
	this := OSReplicationPath{}
	return &this
}

// NewOSReplicationPathWithDefaults instantiates a new OSReplicationPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSReplicationPathWithDefaults() *OSReplicationPath {
	this := OSReplicationPath{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *OSReplicationPath) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationPath) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *OSReplicationPath) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *OSReplicationPath) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *OSReplicationPath) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationPath) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *OSReplicationPath) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *OSReplicationPath) SetCreate(v time.Time) {
	o.Create = &v
}

// GetOsZoneUuids returns the OsZoneUuids field value if set, zero value otherwise.
func (o *OSReplicationPath) GetOsZoneUuids() []string {
	if o == nil || IsNil(o.OsZoneUuids) {
		var ret []string
		return ret
	}
	return o.OsZoneUuids
}

// GetOsZoneUuidsOk returns a tuple with the OsZoneUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationPath) GetOsZoneUuidsOk() ([]string, bool) {
	if o == nil || IsNil(o.OsZoneUuids) {
		return nil, false
	}
	return o.OsZoneUuids, true
}

// HasOsZoneUuids returns a boolean if a field has been set.
func (o *OSReplicationPath) HasOsZoneUuids() bool {
	if o != nil && !IsNil(o.OsZoneUuids) {
		return true
	}

	return false
}

// SetOsZoneUuids gets a reference to the given []string and assigns it to the OsZoneUuids field.
func (o *OSReplicationPath) SetOsZoneUuids(v []string) {
	o.OsZoneUuids = v
}

// GetOsZones returns the OsZones field value if set, zero value otherwise.
func (o *OSReplicationPath) GetOsZones() []ObjectStorageZoneNestview {
	if o == nil || IsNil(o.OsZones) {
		var ret []ObjectStorageZoneNestview
		return ret
	}
	return o.OsZones
}

// GetOsZonesOk returns a tuple with the OsZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationPath) GetOsZonesOk() ([]ObjectStorageZoneNestview, bool) {
	if o == nil || IsNil(o.OsZones) {
		return nil, false
	}
	return o.OsZones, true
}

// HasOsZones returns a boolean if a field has been set.
func (o *OSReplicationPath) HasOsZones() bool {
	if o != nil && !IsNil(o.OsZones) {
		return true
	}

	return false
}

// SetOsZones gets a reference to the given []ObjectStorageZoneNestview and assigns it to the OsZones field.
func (o *OSReplicationPath) SetOsZones(v []ObjectStorageZoneNestview) {
	o.OsZones = v
}

// GetReplicationUuid returns the ReplicationUuid field value if set, zero value otherwise.
func (o *OSReplicationPath) GetReplicationUuid() string {
	if o == nil || IsNil(o.ReplicationUuid) {
		var ret string
		return ret
	}
	return *o.ReplicationUuid
}

// GetReplicationUuidOk returns a tuple with the ReplicationUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationPath) GetReplicationUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationUuid) {
		return nil, false
	}
	return o.ReplicationUuid, true
}

// HasReplicationUuid returns a boolean if a field has been set.
func (o *OSReplicationPath) HasReplicationUuid() bool {
	if o != nil && !IsNil(o.ReplicationUuid) {
		return true
	}

	return false
}

// SetReplicationUuid gets a reference to the given string and assigns it to the ReplicationUuid field.
func (o *OSReplicationPath) SetReplicationUuid(v string) {
	o.ReplicationUuid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OSReplicationPath) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationPath) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OSReplicationPath) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OSReplicationPath) SetStatus(v string) {
	o.Status = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *OSReplicationPath) GetSuspended() bool {
	if o == nil || IsNil(o.Suspended) {
		var ret bool
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationPath) GetSuspendedOk() (*bool, bool) {
	if o == nil || IsNil(o.Suspended) {
		return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *OSReplicationPath) HasSuspended() bool {
	if o != nil && !IsNil(o.Suspended) {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given bool and assigns it to the Suspended field.
func (o *OSReplicationPath) SetSuspended(v bool) {
	o.Suspended = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *OSReplicationPath) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationPath) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *OSReplicationPath) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *OSReplicationPath) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *OSReplicationPath) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSReplicationPath) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *OSReplicationPath) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *OSReplicationPath) SetUuid(v string) {
	o.Uuid = &v
}

func (o OSReplicationPath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSReplicationPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.OsZoneUuids) {
		toSerialize["os_zone_uuids"] = o.OsZoneUuids
	}
	if !IsNil(o.OsZones) {
		toSerialize["os_zones"] = o.OsZones
	}
	if !IsNil(o.ReplicationUuid) {
		toSerialize["replication_uuid"] = o.ReplicationUuid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Suspended) {
		toSerialize["suspended"] = o.Suspended
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableOSReplicationPath struct {
	value *OSReplicationPath
	isSet bool
}

func (v NullableOSReplicationPath) Get() *OSReplicationPath {
	return v.value
}

func (v *NullableOSReplicationPath) Set(val *OSReplicationPath) {
	v.value = val
	v.isSet = true
}

func (v NullableOSReplicationPath) IsSet() bool {
	return v.isSet
}

func (v *NullableOSReplicationPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSReplicationPath(val *OSReplicationPath) *NullableOSReplicationPath {
	return &NullableOSReplicationPath{value: val, isSet: true}
}

func (v NullableOSReplicationPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSReplicationPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


