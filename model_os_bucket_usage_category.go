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

// checks if the OSBucketUsageCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSBucketUsageCategory{}

// OSBucketUsageCategory OSBucketUsageCategory contain usage info
type OSBucketUsageCategory struct {
	DownloadBytes *int64 `json:"download_bytes,omitempty"`
	Ops *int64 `json:"ops,omitempty"`
	UploadBytes *int64 `json:"upload_bytes,omitempty"`
}

// NewOSBucketUsageCategory instantiates a new OSBucketUsageCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSBucketUsageCategory() *OSBucketUsageCategory {
	this := OSBucketUsageCategory{}
	return &this
}

// NewOSBucketUsageCategoryWithDefaults instantiates a new OSBucketUsageCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSBucketUsageCategoryWithDefaults() *OSBucketUsageCategory {
	this := OSBucketUsageCategory{}
	return &this
}

// GetDownloadBytes returns the DownloadBytes field value if set, zero value otherwise.
func (o *OSBucketUsageCategory) GetDownloadBytes() int64 {
	if o == nil || IsNil(o.DownloadBytes) {
		var ret int64
		return ret
	}
	return *o.DownloadBytes
}

// GetDownloadBytesOk returns a tuple with the DownloadBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketUsageCategory) GetDownloadBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DownloadBytes) {
		return nil, false
	}
	return o.DownloadBytes, true
}

// HasDownloadBytes returns a boolean if a field has been set.
func (o *OSBucketUsageCategory) HasDownloadBytes() bool {
	if o != nil && !IsNil(o.DownloadBytes) {
		return true
	}

	return false
}

// SetDownloadBytes gets a reference to the given int64 and assigns it to the DownloadBytes field.
func (o *OSBucketUsageCategory) SetDownloadBytes(v int64) {
	o.DownloadBytes = &v
}

// GetOps returns the Ops field value if set, zero value otherwise.
func (o *OSBucketUsageCategory) GetOps() int64 {
	if o == nil || IsNil(o.Ops) {
		var ret int64
		return ret
	}
	return *o.Ops
}

// GetOpsOk returns a tuple with the Ops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketUsageCategory) GetOpsOk() (*int64, bool) {
	if o == nil || IsNil(o.Ops) {
		return nil, false
	}
	return o.Ops, true
}

// HasOps returns a boolean if a field has been set.
func (o *OSBucketUsageCategory) HasOps() bool {
	if o != nil && !IsNil(o.Ops) {
		return true
	}

	return false
}

// SetOps gets a reference to the given int64 and assigns it to the Ops field.
func (o *OSBucketUsageCategory) SetOps(v int64) {
	o.Ops = &v
}

// GetUploadBytes returns the UploadBytes field value if set, zero value otherwise.
func (o *OSBucketUsageCategory) GetUploadBytes() int64 {
	if o == nil || IsNil(o.UploadBytes) {
		var ret int64
		return ret
	}
	return *o.UploadBytes
}

// GetUploadBytesOk returns a tuple with the UploadBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSBucketUsageCategory) GetUploadBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.UploadBytes) {
		return nil, false
	}
	return o.UploadBytes, true
}

// HasUploadBytes returns a boolean if a field has been set.
func (o *OSBucketUsageCategory) HasUploadBytes() bool {
	if o != nil && !IsNil(o.UploadBytes) {
		return true
	}

	return false
}

// SetUploadBytes gets a reference to the given int64 and assigns it to the UploadBytes field.
func (o *OSBucketUsageCategory) SetUploadBytes(v int64) {
	o.UploadBytes = &v
}

func (o OSBucketUsageCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSBucketUsageCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownloadBytes) {
		toSerialize["download_bytes"] = o.DownloadBytes
	}
	if !IsNil(o.Ops) {
		toSerialize["ops"] = o.Ops
	}
	if !IsNil(o.UploadBytes) {
		toSerialize["upload_bytes"] = o.UploadBytes
	}
	return toSerialize, nil
}

type NullableOSBucketUsageCategory struct {
	value *OSBucketUsageCategory
	isSet bool
}

func (v NullableOSBucketUsageCategory) Get() *OSBucketUsageCategory {
	return v.value
}

func (v *NullableOSBucketUsageCategory) Set(val *OSBucketUsageCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableOSBucketUsageCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableOSBucketUsageCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSBucketUsageCategory(val *OSBucketUsageCategory) *NullableOSBucketUsageCategory {
	return &NullableOSBucketUsageCategory{value: val, isSet: true}
}

func (v NullableOSBucketUsageCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSBucketUsageCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


