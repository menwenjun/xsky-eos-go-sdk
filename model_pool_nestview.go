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

// checks if the PoolNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolNestview{}

// PoolNestview struct for PoolNestview
type PoolNestview struct {
	ActionStatus *string `json:"action_status,omitempty"`
	EncryptEnabled *bool `json:"encrypt_enabled,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	NumaApplyPolicy *string `json:"numa_apply_policy,omitempty"`
	NumaBindPolicy *string `json:"numa_bind_policy,omitempty"`
	NumaEnabled *bool `json:"numa_enabled,omitempty"`
	PoolType *string `json:"pool_type,omitempty"`
	Status *string `json:"status,omitempty"`
	Stretched *bool `json:"stretched,omitempty"`
}

// NewPoolNestview instantiates a new PoolNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolNestview() *PoolNestview {
	this := PoolNestview{}
	return &this
}

// NewPoolNestviewWithDefaults instantiates a new PoolNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolNestviewWithDefaults() *PoolNestview {
	this := PoolNestview{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *PoolNestview) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolNestview) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *PoolNestview) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *PoolNestview) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetEncryptEnabled returns the EncryptEnabled field value if set, zero value otherwise.
func (o *PoolNestview) GetEncryptEnabled() bool {
	if o == nil || IsNil(o.EncryptEnabled) {
		var ret bool
		return ret
	}
	return *o.EncryptEnabled
}

// GetEncryptEnabledOk returns a tuple with the EncryptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolNestview) GetEncryptEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptEnabled) {
		return nil, false
	}
	return o.EncryptEnabled, true
}

// HasEncryptEnabled returns a boolean if a field has been set.
func (o *PoolNestview) HasEncryptEnabled() bool {
	if o != nil && !IsNil(o.EncryptEnabled) {
		return true
	}

	return false
}

// SetEncryptEnabled gets a reference to the given bool and assigns it to the EncryptEnabled field.
func (o *PoolNestview) SetEncryptEnabled(v bool) {
	o.EncryptEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PoolNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PoolNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PoolNestview) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PoolNestview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolNestview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PoolNestview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PoolNestview) SetName(v string) {
	o.Name = &v
}

// GetNumaApplyPolicy returns the NumaApplyPolicy field value if set, zero value otherwise.
func (o *PoolNestview) GetNumaApplyPolicy() string {
	if o == nil || IsNil(o.NumaApplyPolicy) {
		var ret string
		return ret
	}
	return *o.NumaApplyPolicy
}

// GetNumaApplyPolicyOk returns a tuple with the NumaApplyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolNestview) GetNumaApplyPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.NumaApplyPolicy) {
		return nil, false
	}
	return o.NumaApplyPolicy, true
}

// HasNumaApplyPolicy returns a boolean if a field has been set.
func (o *PoolNestview) HasNumaApplyPolicy() bool {
	if o != nil && !IsNil(o.NumaApplyPolicy) {
		return true
	}

	return false
}

// SetNumaApplyPolicy gets a reference to the given string and assigns it to the NumaApplyPolicy field.
func (o *PoolNestview) SetNumaApplyPolicy(v string) {
	o.NumaApplyPolicy = &v
}

// GetNumaBindPolicy returns the NumaBindPolicy field value if set, zero value otherwise.
func (o *PoolNestview) GetNumaBindPolicy() string {
	if o == nil || IsNil(o.NumaBindPolicy) {
		var ret string
		return ret
	}
	return *o.NumaBindPolicy
}

// GetNumaBindPolicyOk returns a tuple with the NumaBindPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolNestview) GetNumaBindPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.NumaBindPolicy) {
		return nil, false
	}
	return o.NumaBindPolicy, true
}

// HasNumaBindPolicy returns a boolean if a field has been set.
func (o *PoolNestview) HasNumaBindPolicy() bool {
	if o != nil && !IsNil(o.NumaBindPolicy) {
		return true
	}

	return false
}

// SetNumaBindPolicy gets a reference to the given string and assigns it to the NumaBindPolicy field.
func (o *PoolNestview) SetNumaBindPolicy(v string) {
	o.NumaBindPolicy = &v
}

// GetNumaEnabled returns the NumaEnabled field value if set, zero value otherwise.
func (o *PoolNestview) GetNumaEnabled() bool {
	if o == nil || IsNil(o.NumaEnabled) {
		var ret bool
		return ret
	}
	return *o.NumaEnabled
}

// GetNumaEnabledOk returns a tuple with the NumaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolNestview) GetNumaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NumaEnabled) {
		return nil, false
	}
	return o.NumaEnabled, true
}

// HasNumaEnabled returns a boolean if a field has been set.
func (o *PoolNestview) HasNumaEnabled() bool {
	if o != nil && !IsNil(o.NumaEnabled) {
		return true
	}

	return false
}

// SetNumaEnabled gets a reference to the given bool and assigns it to the NumaEnabled field.
func (o *PoolNestview) SetNumaEnabled(v bool) {
	o.NumaEnabled = &v
}

// GetPoolType returns the PoolType field value if set, zero value otherwise.
func (o *PoolNestview) GetPoolType() string {
	if o == nil || IsNil(o.PoolType) {
		var ret string
		return ret
	}
	return *o.PoolType
}

// GetPoolTypeOk returns a tuple with the PoolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolNestview) GetPoolTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PoolType) {
		return nil, false
	}
	return o.PoolType, true
}

// HasPoolType returns a boolean if a field has been set.
func (o *PoolNestview) HasPoolType() bool {
	if o != nil && !IsNil(o.PoolType) {
		return true
	}

	return false
}

// SetPoolType gets a reference to the given string and assigns it to the PoolType field.
func (o *PoolNestview) SetPoolType(v string) {
	o.PoolType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PoolNestview) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolNestview) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PoolNestview) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PoolNestview) SetStatus(v string) {
	o.Status = &v
}

// GetStretched returns the Stretched field value if set, zero value otherwise.
func (o *PoolNestview) GetStretched() bool {
	if o == nil || IsNil(o.Stretched) {
		var ret bool
		return ret
	}
	return *o.Stretched
}

// GetStretchedOk returns a tuple with the Stretched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolNestview) GetStretchedOk() (*bool, bool) {
	if o == nil || IsNil(o.Stretched) {
		return nil, false
	}
	return o.Stretched, true
}

// HasStretched returns a boolean if a field has been set.
func (o *PoolNestview) HasStretched() bool {
	if o != nil && !IsNil(o.Stretched) {
		return true
	}

	return false
}

// SetStretched gets a reference to the given bool and assigns it to the Stretched field.
func (o *PoolNestview) SetStretched(v bool) {
	o.Stretched = &v
}

func (o PoolNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.EncryptEnabled) {
		toSerialize["encrypt_enabled"] = o.EncryptEnabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumaApplyPolicy) {
		toSerialize["numa_apply_policy"] = o.NumaApplyPolicy
	}
	if !IsNil(o.NumaBindPolicy) {
		toSerialize["numa_bind_policy"] = o.NumaBindPolicy
	}
	if !IsNil(o.NumaEnabled) {
		toSerialize["numa_enabled"] = o.NumaEnabled
	}
	if !IsNil(o.PoolType) {
		toSerialize["pool_type"] = o.PoolType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Stretched) {
		toSerialize["stretched"] = o.Stretched
	}
	return toSerialize, nil
}

type NullablePoolNestview struct {
	value *PoolNestview
	isSet bool
}

func (v NullablePoolNestview) Get() *PoolNestview {
	return v.value
}

func (v *NullablePoolNestview) Set(val *PoolNestview) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolNestview) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolNestview(val *PoolNestview) *NullablePoolNestview {
	return &NullablePoolNestview{value: val, isSet: true}
}

func (v NullablePoolNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


