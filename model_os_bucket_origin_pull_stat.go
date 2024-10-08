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

// checks if the OSBucketOriginPullStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSBucketOriginPullStat{}

// OSBucketOriginPullStat OSBucketOriginPullStat define the origin pull statistics
type OSBucketOriginPullStat struct {
	Create *time.Time `json:"create,omitempty"`
	ExpiredObjects *int64 `json:"expired_objects,omitempty"`
	ExpiredSize *int64 `json:"expired_size,omitempty"`
	FailureOps *int64 `json:"failure_ops,omitempty"`
	RxObjects *int64 `json:"rx_objects,omitempty"`
	RxSize *int64 `json:"rx_size,omitempty"`
	StorageClasses *map[string]OSOriginPullStorageClassStat `json:"storage_classes,omitempty"`
	SuccessOps *int64 `json:"success_ops,omitempty"`
	TotalOps *int64 `json:"total_ops,omitempty"`
	TotalRxObjects *int64 `json:"total_rx_objects,omitempty"`
	TotalRxSize *int64 `json:"total_rx_size,omitempty"`
}

// NewOSBucketOriginPullStat instantiates a new OSBucketOriginPullStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSBucketOriginPullStat() *OSBucketOriginPullStat {
	this := OSBucketOriginPullStat{}
	return &this
}

// NewOSBucketOriginPullStatWithDefaults instantiates a new OSBucketOriginPullStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSBucketOriginPullStatWithDefaults() *OSBucketOriginPullStat {
	this := OSBucketOriginPullStat{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *OSBucketOriginPullStat) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketOriginPullStat) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *OSBucketOriginPullStat) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *OSBucketOriginPullStat) SetCreate(v time.Time) {
	o.Create = &v
}

// GetExpiredObjects returns the ExpiredObjects field value if set, zero value otherwise.
func (o *OSBucketOriginPullStat) GetExpiredObjects() int64 {
	if o == nil || IsNil(o.ExpiredObjects) {
		var ret int64
		return ret
	}
	return *o.ExpiredObjects
}

// GetExpiredObjectsOk returns a tuple with the ExpiredObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketOriginPullStat) GetExpiredObjectsOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiredObjects) {
		return nil, false
	}
	return o.ExpiredObjects, true
}

// HasExpiredObjects returns a boolean if a field has been set.
func (o *OSBucketOriginPullStat) HasExpiredObjects() bool {
	if o != nil && !IsNil(o.ExpiredObjects) {
		return true
	}

	return false
}

// SetExpiredObjects gets a reference to the given int64 and assigns it to the ExpiredObjects field.
func (o *OSBucketOriginPullStat) SetExpiredObjects(v int64) {
	o.ExpiredObjects = &v
}

// GetExpiredSize returns the ExpiredSize field value if set, zero value otherwise.
func (o *OSBucketOriginPullStat) GetExpiredSize() int64 {
	if o == nil || IsNil(o.ExpiredSize) {
		var ret int64
		return ret
	}
	return *o.ExpiredSize
}

// GetExpiredSizeOk returns a tuple with the ExpiredSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketOriginPullStat) GetExpiredSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiredSize) {
		return nil, false
	}
	return o.ExpiredSize, true
}

// HasExpiredSize returns a boolean if a field has been set.
func (o *OSBucketOriginPullStat) HasExpiredSize() bool {
	if o != nil && !IsNil(o.ExpiredSize) {
		return true
	}

	return false
}

// SetExpiredSize gets a reference to the given int64 and assigns it to the ExpiredSize field.
func (o *OSBucketOriginPullStat) SetExpiredSize(v int64) {
	o.ExpiredSize = &v
}

// GetFailureOps returns the FailureOps field value if set, zero value otherwise.
func (o *OSBucketOriginPullStat) GetFailureOps() int64 {
	if o == nil || IsNil(o.FailureOps) {
		var ret int64
		return ret
	}
	return *o.FailureOps
}

// GetFailureOpsOk returns a tuple with the FailureOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketOriginPullStat) GetFailureOpsOk() (*int64, bool) {
	if o == nil || IsNil(o.FailureOps) {
		return nil, false
	}
	return o.FailureOps, true
}

// HasFailureOps returns a boolean if a field has been set.
func (o *OSBucketOriginPullStat) HasFailureOps() bool {
	if o != nil && !IsNil(o.FailureOps) {
		return true
	}

	return false
}

// SetFailureOps gets a reference to the given int64 and assigns it to the FailureOps field.
func (o *OSBucketOriginPullStat) SetFailureOps(v int64) {
	o.FailureOps = &v
}

// GetRxObjects returns the RxObjects field value if set, zero value otherwise.
func (o *OSBucketOriginPullStat) GetRxObjects() int64 {
	if o == nil || IsNil(o.RxObjects) {
		var ret int64
		return ret
	}
	return *o.RxObjects
}

// GetRxObjectsOk returns a tuple with the RxObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketOriginPullStat) GetRxObjectsOk() (*int64, bool) {
	if o == nil || IsNil(o.RxObjects) {
		return nil, false
	}
	return o.RxObjects, true
}

// HasRxObjects returns a boolean if a field has been set.
func (o *OSBucketOriginPullStat) HasRxObjects() bool {
	if o != nil && !IsNil(o.RxObjects) {
		return true
	}

	return false
}

// SetRxObjects gets a reference to the given int64 and assigns it to the RxObjects field.
func (o *OSBucketOriginPullStat) SetRxObjects(v int64) {
	o.RxObjects = &v
}

// GetRxSize returns the RxSize field value if set, zero value otherwise.
func (o *OSBucketOriginPullStat) GetRxSize() int64 {
	if o == nil || IsNil(o.RxSize) {
		var ret int64
		return ret
	}
	return *o.RxSize
}

// GetRxSizeOk returns a tuple with the RxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketOriginPullStat) GetRxSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.RxSize) {
		return nil, false
	}
	return o.RxSize, true
}

// HasRxSize returns a boolean if a field has been set.
func (o *OSBucketOriginPullStat) HasRxSize() bool {
	if o != nil && !IsNil(o.RxSize) {
		return true
	}

	return false
}

// SetRxSize gets a reference to the given int64 and assigns it to the RxSize field.
func (o *OSBucketOriginPullStat) SetRxSize(v int64) {
	o.RxSize = &v
}

// GetStorageClasses returns the StorageClasses field value if set, zero value otherwise.
func (o *OSBucketOriginPullStat) GetStorageClasses() map[string]OSOriginPullStorageClassStat {
	if o == nil || IsNil(o.StorageClasses) {
		var ret map[string]OSOriginPullStorageClassStat
		return ret
	}
	return *o.StorageClasses
}

// GetStorageClassesOk returns a tuple with the StorageClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketOriginPullStat) GetStorageClassesOk() (*map[string]OSOriginPullStorageClassStat, bool) {
	if o == nil || IsNil(o.StorageClasses) {
		return nil, false
	}
	return o.StorageClasses, true
}

// HasStorageClasses returns a boolean if a field has been set.
func (o *OSBucketOriginPullStat) HasStorageClasses() bool {
	if o != nil && !IsNil(o.StorageClasses) {
		return true
	}

	return false
}

// SetStorageClasses gets a reference to the given map[string]OSOriginPullStorageClassStat and assigns it to the StorageClasses field.
func (o *OSBucketOriginPullStat) SetStorageClasses(v map[string]OSOriginPullStorageClassStat) {
	o.StorageClasses = &v
}

// GetSuccessOps returns the SuccessOps field value if set, zero value otherwise.
func (o *OSBucketOriginPullStat) GetSuccessOps() int64 {
	if o == nil || IsNil(o.SuccessOps) {
		var ret int64
		return ret
	}
	return *o.SuccessOps
}

// GetSuccessOpsOk returns a tuple with the SuccessOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketOriginPullStat) GetSuccessOpsOk() (*int64, bool) {
	if o == nil || IsNil(o.SuccessOps) {
		return nil, false
	}
	return o.SuccessOps, true
}

// HasSuccessOps returns a boolean if a field has been set.
func (o *OSBucketOriginPullStat) HasSuccessOps() bool {
	if o != nil && !IsNil(o.SuccessOps) {
		return true
	}

	return false
}

// SetSuccessOps gets a reference to the given int64 and assigns it to the SuccessOps field.
func (o *OSBucketOriginPullStat) SetSuccessOps(v int64) {
	o.SuccessOps = &v
}

// GetTotalOps returns the TotalOps field value if set, zero value otherwise.
func (o *OSBucketOriginPullStat) GetTotalOps() int64 {
	if o == nil || IsNil(o.TotalOps) {
		var ret int64
		return ret
	}
	return *o.TotalOps
}

// GetTotalOpsOk returns a tuple with the TotalOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketOriginPullStat) GetTotalOpsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalOps) {
		return nil, false
	}
	return o.TotalOps, true
}

// HasTotalOps returns a boolean if a field has been set.
func (o *OSBucketOriginPullStat) HasTotalOps() bool {
	if o != nil && !IsNil(o.TotalOps) {
		return true
	}

	return false
}

// SetTotalOps gets a reference to the given int64 and assigns it to the TotalOps field.
func (o *OSBucketOriginPullStat) SetTotalOps(v int64) {
	o.TotalOps = &v
}

// GetTotalRxObjects returns the TotalRxObjects field value if set, zero value otherwise.
func (o *OSBucketOriginPullStat) GetTotalRxObjects() int64 {
	if o == nil || IsNil(o.TotalRxObjects) {
		var ret int64
		return ret
	}
	return *o.TotalRxObjects
}

// GetTotalRxObjectsOk returns a tuple with the TotalRxObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketOriginPullStat) GetTotalRxObjectsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalRxObjects) {
		return nil, false
	}
	return o.TotalRxObjects, true
}

// HasTotalRxObjects returns a boolean if a field has been set.
func (o *OSBucketOriginPullStat) HasTotalRxObjects() bool {
	if o != nil && !IsNil(o.TotalRxObjects) {
		return true
	}

	return false
}

// SetTotalRxObjects gets a reference to the given int64 and assigns it to the TotalRxObjects field.
func (o *OSBucketOriginPullStat) SetTotalRxObjects(v int64) {
	o.TotalRxObjects = &v
}

// GetTotalRxSize returns the TotalRxSize field value if set, zero value otherwise.
func (o *OSBucketOriginPullStat) GetTotalRxSize() int64 {
	if o == nil || IsNil(o.TotalRxSize) {
		var ret int64
		return ret
	}
	return *o.TotalRxSize
}

// GetTotalRxSizeOk returns a tuple with the TotalRxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketOriginPullStat) GetTotalRxSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalRxSize) {
		return nil, false
	}
	return o.TotalRxSize, true
}

// HasTotalRxSize returns a boolean if a field has been set.
func (o *OSBucketOriginPullStat) HasTotalRxSize() bool {
	if o != nil && !IsNil(o.TotalRxSize) {
		return true
	}

	return false
}

// SetTotalRxSize gets a reference to the given int64 and assigns it to the TotalRxSize field.
func (o *OSBucketOriginPullStat) SetTotalRxSize(v int64) {
	o.TotalRxSize = &v
}

func (o OSBucketOriginPullStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSBucketOriginPullStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.ExpiredObjects) {
		toSerialize["expired_objects"] = o.ExpiredObjects
	}
	if !IsNil(o.ExpiredSize) {
		toSerialize["expired_size"] = o.ExpiredSize
	}
	if !IsNil(o.FailureOps) {
		toSerialize["failure_ops"] = o.FailureOps
	}
	if !IsNil(o.RxObjects) {
		toSerialize["rx_objects"] = o.RxObjects
	}
	if !IsNil(o.RxSize) {
		toSerialize["rx_size"] = o.RxSize
	}
	if !IsNil(o.StorageClasses) {
		toSerialize["storage_classes"] = o.StorageClasses
	}
	if !IsNil(o.SuccessOps) {
		toSerialize["success_ops"] = o.SuccessOps
	}
	if !IsNil(o.TotalOps) {
		toSerialize["total_ops"] = o.TotalOps
	}
	if !IsNil(o.TotalRxObjects) {
		toSerialize["total_rx_objects"] = o.TotalRxObjects
	}
	if !IsNil(o.TotalRxSize) {
		toSerialize["total_rx_size"] = o.TotalRxSize
	}
	return toSerialize, nil
}

type NullableOSBucketOriginPullStat struct {
	value *OSBucketOriginPullStat
	isSet bool
}

func (v NullableOSBucketOriginPullStat) Get() *OSBucketOriginPullStat {
	return v.value
}

func (v *NullableOSBucketOriginPullStat) Set(val *OSBucketOriginPullStat) {
	v.value = val
	v.isSet = true
}

func (v NullableOSBucketOriginPullStat) IsSet() bool {
	return v.isSet
}

func (v *NullableOSBucketOriginPullStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSBucketOriginPullStat(val *OSBucketOriginPullStat) *NullableOSBucketOriginPullStat {
	return &NullableOSBucketOriginPullStat{value: val, isSet: true}
}

func (v NullableOSBucketOriginPullStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSBucketOriginPullStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


