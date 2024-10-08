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

// checks if the DfsStorageClassUpdateReqDfsStorageClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsStorageClassUpdateReqDfsStorageClass{}

// DfsStorageClassUpdateReqDfsStorageClass struct for DfsStorageClassUpdateReqDfsStorageClass
type DfsStorageClassUpdateReqDfsStorageClass struct {
	// description of custom class
	Description *string `json:"description,omitempty"`
	// name of custom class
	Name *string `json:"name,omitempty"`
	// active pool policy array of custom class
	PoolPolicies []PoolPolicy `json:"pool_policies,omitempty"`
	// write policy of custom class
	WritePolicy *string `json:"write_policy,omitempty"`
}

// NewDfsStorageClassUpdateReqDfsStorageClass instantiates a new DfsStorageClassUpdateReqDfsStorageClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsStorageClassUpdateReqDfsStorageClass() *DfsStorageClassUpdateReqDfsStorageClass {
	this := DfsStorageClassUpdateReqDfsStorageClass{}
	return &this
}

// NewDfsStorageClassUpdateReqDfsStorageClassWithDefaults instantiates a new DfsStorageClassUpdateReqDfsStorageClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsStorageClassUpdateReqDfsStorageClassWithDefaults() *DfsStorageClassUpdateReqDfsStorageClass {
	this := DfsStorageClassUpdateReqDfsStorageClass{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DfsStorageClassUpdateReqDfsStorageClass) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsStorageClassUpdateReqDfsStorageClass) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DfsStorageClassUpdateReqDfsStorageClass) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DfsStorageClassUpdateReqDfsStorageClass) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DfsStorageClassUpdateReqDfsStorageClass) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsStorageClassUpdateReqDfsStorageClass) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DfsStorageClassUpdateReqDfsStorageClass) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DfsStorageClassUpdateReqDfsStorageClass) SetName(v string) {
	o.Name = &v
}

// GetPoolPolicies returns the PoolPolicies field value if set, zero value otherwise.
func (o *DfsStorageClassUpdateReqDfsStorageClass) GetPoolPolicies() []PoolPolicy {
	if o == nil || IsNil(o.PoolPolicies) {
		var ret []PoolPolicy
		return ret
	}
	return o.PoolPolicies
}

// GetPoolPoliciesOk returns a tuple with the PoolPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsStorageClassUpdateReqDfsStorageClass) GetPoolPoliciesOk() ([]PoolPolicy, bool) {
	if o == nil || IsNil(o.PoolPolicies) {
		return nil, false
	}
	return o.PoolPolicies, true
}

// HasPoolPolicies returns a boolean if a field has been set.
func (o *DfsStorageClassUpdateReqDfsStorageClass) HasPoolPolicies() bool {
	if o != nil && !IsNil(o.PoolPolicies) {
		return true
	}

	return false
}

// SetPoolPolicies gets a reference to the given []PoolPolicy and assigns it to the PoolPolicies field.
func (o *DfsStorageClassUpdateReqDfsStorageClass) SetPoolPolicies(v []PoolPolicy) {
	o.PoolPolicies = v
}

// GetWritePolicy returns the WritePolicy field value if set, zero value otherwise.
func (o *DfsStorageClassUpdateReqDfsStorageClass) GetWritePolicy() string {
	if o == nil || IsNil(o.WritePolicy) {
		var ret string
		return ret
	}
	return *o.WritePolicy
}

// GetWritePolicyOk returns a tuple with the WritePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsStorageClassUpdateReqDfsStorageClass) GetWritePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.WritePolicy) {
		return nil, false
	}
	return o.WritePolicy, true
}

// HasWritePolicy returns a boolean if a field has been set.
func (o *DfsStorageClassUpdateReqDfsStorageClass) HasWritePolicy() bool {
	if o != nil && !IsNil(o.WritePolicy) {
		return true
	}

	return false
}

// SetWritePolicy gets a reference to the given string and assigns it to the WritePolicy field.
func (o *DfsStorageClassUpdateReqDfsStorageClass) SetWritePolicy(v string) {
	o.WritePolicy = &v
}

func (o DfsStorageClassUpdateReqDfsStorageClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsStorageClassUpdateReqDfsStorageClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PoolPolicies) {
		toSerialize["pool_policies"] = o.PoolPolicies
	}
	if !IsNil(o.WritePolicy) {
		toSerialize["write_policy"] = o.WritePolicy
	}
	return toSerialize, nil
}

type NullableDfsStorageClassUpdateReqDfsStorageClass struct {
	value *DfsStorageClassUpdateReqDfsStorageClass
	isSet bool
}

func (v NullableDfsStorageClassUpdateReqDfsStorageClass) Get() *DfsStorageClassUpdateReqDfsStorageClass {
	return v.value
}

func (v *NullableDfsStorageClassUpdateReqDfsStorageClass) Set(val *DfsStorageClassUpdateReqDfsStorageClass) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsStorageClassUpdateReqDfsStorageClass) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsStorageClassUpdateReqDfsStorageClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsStorageClassUpdateReqDfsStorageClass(val *DfsStorageClassUpdateReqDfsStorageClass) *NullableDfsStorageClassUpdateReqDfsStorageClass {
	return &NullableDfsStorageClassUpdateReqDfsStorageClass{value: val, isSet: true}
}

func (v NullableDfsStorageClassUpdateReqDfsStorageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsStorageClassUpdateReqDfsStorageClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


