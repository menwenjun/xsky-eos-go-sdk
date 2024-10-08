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

// checks if the OSStorageClassStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSStorageClassStat{}

// OSStorageClassStat OSStorageClassStat define the os storage class or os external storage class statistics of a user or bucket
type OSStorageClassStat struct {
	AllocatedObjects *int64 `json:"allocated_objects,omitempty"`
	AllocatedSize *int64 `json:"allocated_size,omitempty"`
	CacheAllocatedObjects *int64 `json:"cache_allocated_objects,omitempty"`
	CacheAllocatedSize *int64 `json:"cache_allocated_size,omitempty"`
	// ClassName used in GetObjectStorageUserSamples
	ClassName *string `json:"class_name,omitempty"`
}

// NewOSStorageClassStat instantiates a new OSStorageClassStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSStorageClassStat() *OSStorageClassStat {
	this := OSStorageClassStat{}
	return &this
}

// NewOSStorageClassStatWithDefaults instantiates a new OSStorageClassStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSStorageClassStatWithDefaults() *OSStorageClassStat {
	this := OSStorageClassStat{}
	return &this
}

// GetAllocatedObjects returns the AllocatedObjects field value if set, zero value otherwise.
func (o *OSStorageClassStat) GetAllocatedObjects() int64 {
	if o == nil || IsNil(o.AllocatedObjects) {
		var ret int64
		return ret
	}
	return *o.AllocatedObjects
}

// GetAllocatedObjectsOk returns a tuple with the AllocatedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSStorageClassStat) GetAllocatedObjectsOk() (*int64, bool) {
	if o == nil || IsNil(o.AllocatedObjects) {
		return nil, false
	}
	return o.AllocatedObjects, true
}

// HasAllocatedObjects returns a boolean if a field has been set.
func (o *OSStorageClassStat) HasAllocatedObjects() bool {
	if o != nil && !IsNil(o.AllocatedObjects) {
		return true
	}

	return false
}

// SetAllocatedObjects gets a reference to the given int64 and assigns it to the AllocatedObjects field.
func (o *OSStorageClassStat) SetAllocatedObjects(v int64) {
	o.AllocatedObjects = &v
}

// GetAllocatedSize returns the AllocatedSize field value if set, zero value otherwise.
func (o *OSStorageClassStat) GetAllocatedSize() int64 {
	if o == nil || IsNil(o.AllocatedSize) {
		var ret int64
		return ret
	}
	return *o.AllocatedSize
}

// GetAllocatedSizeOk returns a tuple with the AllocatedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSStorageClassStat) GetAllocatedSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.AllocatedSize) {
		return nil, false
	}
	return o.AllocatedSize, true
}

// HasAllocatedSize returns a boolean if a field has been set.
func (o *OSStorageClassStat) HasAllocatedSize() bool {
	if o != nil && !IsNil(o.AllocatedSize) {
		return true
	}

	return false
}

// SetAllocatedSize gets a reference to the given int64 and assigns it to the AllocatedSize field.
func (o *OSStorageClassStat) SetAllocatedSize(v int64) {
	o.AllocatedSize = &v
}

// GetCacheAllocatedObjects returns the CacheAllocatedObjects field value if set, zero value otherwise.
func (o *OSStorageClassStat) GetCacheAllocatedObjects() int64 {
	if o == nil || IsNil(o.CacheAllocatedObjects) {
		var ret int64
		return ret
	}
	return *o.CacheAllocatedObjects
}

// GetCacheAllocatedObjectsOk returns a tuple with the CacheAllocatedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSStorageClassStat) GetCacheAllocatedObjectsOk() (*int64, bool) {
	if o == nil || IsNil(o.CacheAllocatedObjects) {
		return nil, false
	}
	return o.CacheAllocatedObjects, true
}

// HasCacheAllocatedObjects returns a boolean if a field has been set.
func (o *OSStorageClassStat) HasCacheAllocatedObjects() bool {
	if o != nil && !IsNil(o.CacheAllocatedObjects) {
		return true
	}

	return false
}

// SetCacheAllocatedObjects gets a reference to the given int64 and assigns it to the CacheAllocatedObjects field.
func (o *OSStorageClassStat) SetCacheAllocatedObjects(v int64) {
	o.CacheAllocatedObjects = &v
}

// GetCacheAllocatedSize returns the CacheAllocatedSize field value if set, zero value otherwise.
func (o *OSStorageClassStat) GetCacheAllocatedSize() int64 {
	if o == nil || IsNil(o.CacheAllocatedSize) {
		var ret int64
		return ret
	}
	return *o.CacheAllocatedSize
}

// GetCacheAllocatedSizeOk returns a tuple with the CacheAllocatedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSStorageClassStat) GetCacheAllocatedSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.CacheAllocatedSize) {
		return nil, false
	}
	return o.CacheAllocatedSize, true
}

// HasCacheAllocatedSize returns a boolean if a field has been set.
func (o *OSStorageClassStat) HasCacheAllocatedSize() bool {
	if o != nil && !IsNil(o.CacheAllocatedSize) {
		return true
	}

	return false
}

// SetCacheAllocatedSize gets a reference to the given int64 and assigns it to the CacheAllocatedSize field.
func (o *OSStorageClassStat) SetCacheAllocatedSize(v int64) {
	o.CacheAllocatedSize = &v
}

// GetClassName returns the ClassName field value if set, zero value otherwise.
func (o *OSStorageClassStat) GetClassName() string {
	if o == nil || IsNil(o.ClassName) {
		var ret string
		return ret
	}
	return *o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSStorageClassStat) GetClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClassName) {
		return nil, false
	}
	return o.ClassName, true
}

// HasClassName returns a boolean if a field has been set.
func (o *OSStorageClassStat) HasClassName() bool {
	if o != nil && !IsNil(o.ClassName) {
		return true
	}

	return false
}

// SetClassName gets a reference to the given string and assigns it to the ClassName field.
func (o *OSStorageClassStat) SetClassName(v string) {
	o.ClassName = &v
}

func (o OSStorageClassStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSStorageClassStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllocatedObjects) {
		toSerialize["allocated_objects"] = o.AllocatedObjects
	}
	if !IsNil(o.AllocatedSize) {
		toSerialize["allocated_size"] = o.AllocatedSize
	}
	if !IsNil(o.CacheAllocatedObjects) {
		toSerialize["cache_allocated_objects"] = o.CacheAllocatedObjects
	}
	if !IsNil(o.CacheAllocatedSize) {
		toSerialize["cache_allocated_size"] = o.CacheAllocatedSize
	}
	if !IsNil(o.ClassName) {
		toSerialize["class_name"] = o.ClassName
	}
	return toSerialize, nil
}

type NullableOSStorageClassStat struct {
	value *OSStorageClassStat
	isSet bool
}

func (v NullableOSStorageClassStat) Get() *OSStorageClassStat {
	return v.value
}

func (v *NullableOSStorageClassStat) Set(val *OSStorageClassStat) {
	v.value = val
	v.isSet = true
}

func (v NullableOSStorageClassStat) IsSet() bool {
	return v.isSet
}

func (v *NullableOSStorageClassStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSStorageClassStat(val *OSStorageClassStat) *NullableOSStorageClassStat {
	return &NullableOSStorageClassStat{value: val, isSet: true}
}

func (v NullableOSStorageClassStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSStorageClassStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


