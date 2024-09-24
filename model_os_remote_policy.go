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

// checks if the OSRemotePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSRemotePolicy{}

// OSRemotePolicy OSRemotePolicy defines model of cached remote os policy.
type OSRemotePolicy struct {
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Default *bool `json:"default,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	ZoneUuid *string `json:"zone_uuid,omitempty"`
}

// NewOSRemotePolicy instantiates a new OSRemotePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSRemotePolicy() *OSRemotePolicy {
	this := OSRemotePolicy{}
	return &this
}

// NewOSRemotePolicyWithDefaults instantiates a new OSRemotePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSRemotePolicyWithDefaults() *OSRemotePolicy {
	this := OSRemotePolicy{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *OSRemotePolicy) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSRemotePolicy) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *OSRemotePolicy) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *OSRemotePolicy) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *OSRemotePolicy) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSRemotePolicy) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *OSRemotePolicy) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *OSRemotePolicy) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *OSRemotePolicy) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSRemotePolicy) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *OSRemotePolicy) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *OSRemotePolicy) SetDefault(v bool) {
	o.Default = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OSRemotePolicy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSRemotePolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OSRemotePolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OSRemotePolicy) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OSRemotePolicy) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSRemotePolicy) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OSRemotePolicy) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OSRemotePolicy) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *OSRemotePolicy) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSRemotePolicy) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *OSRemotePolicy) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *OSRemotePolicy) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *OSRemotePolicy) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSRemotePolicy) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *OSRemotePolicy) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *OSRemotePolicy) SetUuid(v string) {
	o.Uuid = &v
}

// GetZoneUuid returns the ZoneUuid field value if set, zero value otherwise.
func (o *OSRemotePolicy) GetZoneUuid() string {
	if o == nil || IsNil(o.ZoneUuid) {
		var ret string
		return ret
	}
	return *o.ZoneUuid
}

// GetZoneUuidOk returns a tuple with the ZoneUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSRemotePolicy) GetZoneUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneUuid) {
		return nil, false
	}
	return o.ZoneUuid, true
}

// HasZoneUuid returns a boolean if a field has been set.
func (o *OSRemotePolicy) HasZoneUuid() bool {
	if o != nil && !IsNil(o.ZoneUuid) {
		return true
	}

	return false
}

// SetZoneUuid gets a reference to the given string and assigns it to the ZoneUuid field.
func (o *OSRemotePolicy) SetZoneUuid(v string) {
	o.ZoneUuid = &v
}

func (o OSRemotePolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSRemotePolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.ZoneUuid) {
		toSerialize["zone_uuid"] = o.ZoneUuid
	}
	return toSerialize, nil
}

type NullableOSRemotePolicy struct {
	value *OSRemotePolicy
	isSet bool
}

func (v NullableOSRemotePolicy) Get() *OSRemotePolicy {
	return v.value
}

func (v *NullableOSRemotePolicy) Set(val *OSRemotePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableOSRemotePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableOSRemotePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSRemotePolicy(val *OSRemotePolicy) *NullableOSRemotePolicy {
	return &NullableOSRemotePolicy{value: val, isSet: true}
}

func (v NullableOSRemotePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSRemotePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


