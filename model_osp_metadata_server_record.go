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

// checks if the OspMetadataServerRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspMetadataServerRecord{}

// OspMetadataServerRecord OspMetadataServerRecord combines OspMetadataServer and OspMetadataServerStat as api response
type OspMetadataServerRecord struct {
	Class *string `json:"class,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Datacenter *PlacementNodeNestview `json:"datacenter,omitempty"`
	Host *HostNestview `json:"host,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Ip *string `json:"ip,omitempty"`
	MetadataCluster *OspMetadataClusterNestview `json:"metadata_cluster,omitempty"`
	Osd *OsdNestview `json:"osd,omitempty"`
	Partition *PartitionNestview `json:"partition,omitempty"`
	Port *int64 `json:"port,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Version *int64 `json:"version,omitempty"`
	Samples []OspMetadataServerStat `json:"samples,omitempty"`
}

// NewOspMetadataServerRecord instantiates a new OspMetadataServerRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspMetadataServerRecord() *OspMetadataServerRecord {
	this := OspMetadataServerRecord{}
	return &this
}

// NewOspMetadataServerRecordWithDefaults instantiates a new OspMetadataServerRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspMetadataServerRecordWithDefaults() *OspMetadataServerRecord {
	this := OspMetadataServerRecord{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *OspMetadataServerRecord) SetClass(v string) {
	o.Class = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *OspMetadataServerRecord) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *OspMetadataServerRecord) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetDatacenter() PlacementNodeNestview {
	if o == nil || IsNil(o.Datacenter) {
		var ret PlacementNodeNestview
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetDatacenterOk() (*PlacementNodeNestview, bool) {
	if o == nil || IsNil(o.Datacenter) {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasDatacenter() bool {
	if o != nil && !IsNil(o.Datacenter) {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given PlacementNodeNestview and assigns it to the Datacenter field.
func (o *OspMetadataServerRecord) SetDatacenter(v PlacementNodeNestview) {
	o.Datacenter = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetHost() HostNestview {
	if o == nil || IsNil(o.Host) {
		var ret HostNestview
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetHostOk() (*HostNestview, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostNestview and assigns it to the Host field.
func (o *OspMetadataServerRecord) SetHost(v HostNestview) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OspMetadataServerRecord) SetId(v int64) {
	o.Id = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *OspMetadataServerRecord) SetIp(v string) {
	o.Ip = &v
}

// GetMetadataCluster returns the MetadataCluster field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetMetadataCluster() OspMetadataClusterNestview {
	if o == nil || IsNil(o.MetadataCluster) {
		var ret OspMetadataClusterNestview
		return ret
	}
	return *o.MetadataCluster
}

// GetMetadataClusterOk returns a tuple with the MetadataCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetMetadataClusterOk() (*OspMetadataClusterNestview, bool) {
	if o == nil || IsNil(o.MetadataCluster) {
		return nil, false
	}
	return o.MetadataCluster, true
}

// HasMetadataCluster returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasMetadataCluster() bool {
	if o != nil && !IsNil(o.MetadataCluster) {
		return true
	}

	return false
}

// SetMetadataCluster gets a reference to the given OspMetadataClusterNestview and assigns it to the MetadataCluster field.
func (o *OspMetadataServerRecord) SetMetadataCluster(v OspMetadataClusterNestview) {
	o.MetadataCluster = &v
}

// GetOsd returns the Osd field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetOsd() OsdNestview {
	if o == nil || IsNil(o.Osd) {
		var ret OsdNestview
		return ret
	}
	return *o.Osd
}

// GetOsdOk returns a tuple with the Osd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetOsdOk() (*OsdNestview, bool) {
	if o == nil || IsNil(o.Osd) {
		return nil, false
	}
	return o.Osd, true
}

// HasOsd returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasOsd() bool {
	if o != nil && !IsNil(o.Osd) {
		return true
	}

	return false
}

// SetOsd gets a reference to the given OsdNestview and assigns it to the Osd field.
func (o *OspMetadataServerRecord) SetOsd(v OsdNestview) {
	o.Osd = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetPartition() PartitionNestview {
	if o == nil || IsNil(o.Partition) {
		var ret PartitionNestview
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetPartitionOk() (*PartitionNestview, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given PartitionNestview and assigns it to the Partition field.
func (o *OspMetadataServerRecord) SetPartition(v PartitionNestview) {
	o.Partition = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *OspMetadataServerRecord) SetPort(v int64) {
	o.Port = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *OspMetadataServerRecord) SetRoles(v []string) {
	o.Roles = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OspMetadataServerRecord) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *OspMetadataServerRecord) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *OspMetadataServerRecord) SetVersion(v int64) {
	o.Version = &v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *OspMetadataServerRecord) GetSamples() []OspMetadataServerStat {
	if o == nil || IsNil(o.Samples) {
		var ret []OspMetadataServerStat
		return ret
	}
	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspMetadataServerRecord) GetSamplesOk() ([]OspMetadataServerStat, bool) {
	if o == nil || IsNil(o.Samples) {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *OspMetadataServerRecord) HasSamples() bool {
	if o != nil && !IsNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given []OspMetadataServerStat and assigns it to the Samples field.
func (o *OspMetadataServerRecord) SetSamples(v []OspMetadataServerStat) {
	o.Samples = v
}

func (o OspMetadataServerRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspMetadataServerRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Datacenter) {
		toSerialize["datacenter"] = o.Datacenter
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.MetadataCluster) {
		toSerialize["metadata_cluster"] = o.MetadataCluster
	}
	if !IsNil(o.Osd) {
		toSerialize["osd"] = o.Osd
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	return toSerialize, nil
}

type NullableOspMetadataServerRecord struct {
	value *OspMetadataServerRecord
	isSet bool
}

func (v NullableOspMetadataServerRecord) Get() *OspMetadataServerRecord {
	return v.value
}

func (v *NullableOspMetadataServerRecord) Set(val *OspMetadataServerRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableOspMetadataServerRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableOspMetadataServerRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspMetadataServerRecord(val *OspMetadataServerRecord) *NullableOspMetadataServerRecord {
	return &NullableOspMetadataServerRecord{value: val, isSet: true}
}

func (v NullableOspMetadataServerRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspMetadataServerRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


