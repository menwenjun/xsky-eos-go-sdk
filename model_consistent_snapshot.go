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

// checks if the ConsistentSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsistentSnapshot{}

// ConsistentSnapshot ConsistentSnapshot defines the model of consistent snapshot
type ConsistentSnapshot struct {
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *int64 `json:"id,omitempty"`
	SnapName *string `json:"snap_name,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewConsistentSnapshot instantiates a new ConsistentSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsistentSnapshot() *ConsistentSnapshot {
	this := ConsistentSnapshot{}
	return &this
}

// NewConsistentSnapshotWithDefaults instantiates a new ConsistentSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsistentSnapshotWithDefaults() *ConsistentSnapshot {
	this := ConsistentSnapshot{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ConsistentSnapshot) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistentSnapshot) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ConsistentSnapshot) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *ConsistentSnapshot) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ConsistentSnapshot) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistentSnapshot) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ConsistentSnapshot) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ConsistentSnapshot) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConsistentSnapshot) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistentSnapshot) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConsistentSnapshot) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConsistentSnapshot) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConsistentSnapshot) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistentSnapshot) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConsistentSnapshot) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ConsistentSnapshot) SetId(v int64) {
	o.Id = &v
}

// GetSnapName returns the SnapName field value if set, zero value otherwise.
func (o *ConsistentSnapshot) GetSnapName() string {
	if o == nil || IsNil(o.SnapName) {
		var ret string
		return ret
	}
	return *o.SnapName
}

// GetSnapNameOk returns a tuple with the SnapName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistentSnapshot) GetSnapNameOk() (*string, bool) {
	if o == nil || IsNil(o.SnapName) {
		return nil, false
	}
	return o.SnapName, true
}

// HasSnapName returns a boolean if a field has been set.
func (o *ConsistentSnapshot) HasSnapName() bool {
	if o != nil && !IsNil(o.SnapName) {
		return true
	}

	return false
}

// SetSnapName gets a reference to the given string and assigns it to the SnapName field.
func (o *ConsistentSnapshot) SetSnapName(v string) {
	o.SnapName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConsistentSnapshot) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistentSnapshot) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConsistentSnapshot) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConsistentSnapshot) SetStatus(v string) {
	o.Status = &v
}

func (o ConsistentSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsistentSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SnapName) {
		toSerialize["snap_name"] = o.SnapName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableConsistentSnapshot struct {
	value *ConsistentSnapshot
	isSet bool
}

func (v NullableConsistentSnapshot) Get() *ConsistentSnapshot {
	return v.value
}

func (v *NullableConsistentSnapshot) Set(val *ConsistentSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableConsistentSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableConsistentSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsistentSnapshot(val *ConsistentSnapshot) *NullableConsistentSnapshot {
	return &NullableConsistentSnapshot{value: val, isSet: true}
}

func (v NullableConsistentSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsistentSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


