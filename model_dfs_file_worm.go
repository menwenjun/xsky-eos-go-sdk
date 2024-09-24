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

// checks if the DfsFileWorm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsFileWorm{}

// DfsFileWorm struct for DfsFileWorm
type DfsFileWorm struct {
	// automatic lock time of file
	AutoLockTime *time.Time `json:"auto_lock_time,omitempty"`
	DfsWorm *DfsWorm `json:"dfs_worm,omitempty"`
	// worm expire time of file
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	// worm state of file
	State *string `json:"state,omitempty"`
}

// NewDfsFileWorm instantiates a new DfsFileWorm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsFileWorm() *DfsFileWorm {
	this := DfsFileWorm{}
	return &this
}

// NewDfsFileWormWithDefaults instantiates a new DfsFileWorm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsFileWormWithDefaults() *DfsFileWorm {
	this := DfsFileWorm{}
	return &this
}

// GetAutoLockTime returns the AutoLockTime field value if set, zero value otherwise.
func (o *DfsFileWorm) GetAutoLockTime() time.Time {
	if o == nil || IsNil(o.AutoLockTime) {
		var ret time.Time
		return ret
	}
	return *o.AutoLockTime
}

// GetAutoLockTimeOk returns a tuple with the AutoLockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFileWorm) GetAutoLockTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AutoLockTime) {
		return nil, false
	}
	return o.AutoLockTime, true
}

// HasAutoLockTime returns a boolean if a field has been set.
func (o *DfsFileWorm) HasAutoLockTime() bool {
	if o != nil && !IsNil(o.AutoLockTime) {
		return true
	}

	return false
}

// SetAutoLockTime gets a reference to the given time.Time and assigns it to the AutoLockTime field.
func (o *DfsFileWorm) SetAutoLockTime(v time.Time) {
	o.AutoLockTime = &v
}

// GetDfsWorm returns the DfsWorm field value if set, zero value otherwise.
func (o *DfsFileWorm) GetDfsWorm() DfsWorm {
	if o == nil || IsNil(o.DfsWorm) {
		var ret DfsWorm
		return ret
	}
	return *o.DfsWorm
}

// GetDfsWormOk returns a tuple with the DfsWorm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFileWorm) GetDfsWormOk() (*DfsWorm, bool) {
	if o == nil || IsNil(o.DfsWorm) {
		return nil, false
	}
	return o.DfsWorm, true
}

// HasDfsWorm returns a boolean if a field has been set.
func (o *DfsFileWorm) HasDfsWorm() bool {
	if o != nil && !IsNil(o.DfsWorm) {
		return true
	}

	return false
}

// SetDfsWorm gets a reference to the given DfsWorm and assigns it to the DfsWorm field.
func (o *DfsFileWorm) SetDfsWorm(v DfsWorm) {
	o.DfsWorm = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *DfsFileWorm) GetExpireTime() time.Time {
	if o == nil || IsNil(o.ExpireTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFileWorm) GetExpireTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *DfsFileWorm) HasExpireTime() bool {
	if o != nil && !IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given time.Time and assigns it to the ExpireTime field.
func (o *DfsFileWorm) SetExpireTime(v time.Time) {
	o.ExpireTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DfsFileWorm) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsFileWorm) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DfsFileWorm) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DfsFileWorm) SetState(v string) {
	o.State = &v
}

func (o DfsFileWorm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsFileWorm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoLockTime) {
		toSerialize["auto_lock_time"] = o.AutoLockTime
	}
	if !IsNil(o.DfsWorm) {
		toSerialize["dfs_worm"] = o.DfsWorm
	}
	if !IsNil(o.ExpireTime) {
		toSerialize["expire_time"] = o.ExpireTime
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableDfsFileWorm struct {
	value *DfsFileWorm
	isSet bool
}

func (v NullableDfsFileWorm) Get() *DfsFileWorm {
	return v.value
}

func (v *NullableDfsFileWorm) Set(val *DfsFileWorm) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsFileWorm) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsFileWorm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsFileWorm(val *DfsFileWorm) *NullableDfsFileWorm {
	return &NullableDfsFileWorm{value: val, isSet: true}
}

func (v NullableDfsFileWorm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsFileWorm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


