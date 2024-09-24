/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ObjectStorageLifecycle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStorageLifecycle{}

// ObjectStorageLifecycle ObjectStorageLifecycle is the model of object_storage_lifecycle
type ObjectStorageLifecycle struct {
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Rules []map[string]interface{} `json:"rules,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewObjectStorageLifecycle instantiates a new ObjectStorageLifecycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageLifecycle() *ObjectStorageLifecycle {
	this := ObjectStorageLifecycle{}
	return &this
}

// NewObjectStorageLifecycleWithDefaults instantiates a new ObjectStorageLifecycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageLifecycleWithDefaults() *ObjectStorageLifecycle {
	this := ObjectStorageLifecycle{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ObjectStorageLifecycle) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageLifecycle) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ObjectStorageLifecycle) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *ObjectStorageLifecycle) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObjectStorageLifecycle) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageLifecycle) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectStorageLifecycle) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ObjectStorageLifecycle) SetId(v int64) {
	o.Id = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ObjectStorageLifecycle) GetRules() []map[string]interface{} {
	if o == nil || IsNil(o.Rules) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageLifecycle) GetRulesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ObjectStorageLifecycle) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []map[string]interface{} and assigns it to the Rules field.
func (o *ObjectStorageLifecycle) SetRules(v []map[string]interface{}) {
	o.Rules = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ObjectStorageLifecycle) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStorageLifecycle) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ObjectStorageLifecycle) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ObjectStorageLifecycle) SetStatus(v string) {
	o.Status = &v
}

func (o ObjectStorageLifecycle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStorageLifecycle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableObjectStorageLifecycle struct {
	value *ObjectStorageLifecycle
	isSet bool
}

func (v NullableObjectStorageLifecycle) Get() *ObjectStorageLifecycle {
	return v.value
}

func (v *NullableObjectStorageLifecycle) Set(val *ObjectStorageLifecycle) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageLifecycle) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageLifecycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageLifecycle(val *ObjectStorageLifecycle) *NullableObjectStorageLifecycle {
	return &NullableObjectStorageLifecycle{value: val, isSet: true}
}

func (v NullableObjectStorageLifecycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageLifecycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


