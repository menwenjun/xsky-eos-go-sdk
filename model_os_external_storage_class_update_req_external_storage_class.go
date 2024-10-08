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

// checks if the OSExternalStorageClassUpdateReqExternalStorageClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSExternalStorageClassUpdateReqExternalStorageClass{}

// OSExternalStorageClassUpdateReqExternalStorageClass struct for OSExternalStorageClassUpdateReqExternalStorageClass
type OSExternalStorageClassUpdateReqExternalStorageClass struct {
	Description *string `json:"description,omitempty"`
	ExternalStoragePlatforms []OSExternalStoragePlatformInfo `json:"external_storage_platforms,omitempty"`
	Name *string `json:"name,omitempty"`
	PrefixEnabled *bool `json:"prefix_enabled,omitempty"`
	// SliceStringField defines slice string field
	StoragePattern []string `json:"storage_pattern,omitempty"`
}

// NewOSExternalStorageClassUpdateReqExternalStorageClass instantiates a new OSExternalStorageClassUpdateReqExternalStorageClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSExternalStorageClassUpdateReqExternalStorageClass() *OSExternalStorageClassUpdateReqExternalStorageClass {
	this := OSExternalStorageClassUpdateReqExternalStorageClass{}
	return &this
}

// NewOSExternalStorageClassUpdateReqExternalStorageClassWithDefaults instantiates a new OSExternalStorageClassUpdateReqExternalStorageClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSExternalStorageClassUpdateReqExternalStorageClassWithDefaults() *OSExternalStorageClassUpdateReqExternalStorageClass {
	this := OSExternalStorageClassUpdateReqExternalStorageClass{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) SetDescription(v string) {
	o.Description = &v
}

// GetExternalStoragePlatforms returns the ExternalStoragePlatforms field value if set, zero value otherwise.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) GetExternalStoragePlatforms() []OSExternalStoragePlatformInfo {
	if o == nil || IsNil(o.ExternalStoragePlatforms) {
		var ret []OSExternalStoragePlatformInfo
		return ret
	}
	return o.ExternalStoragePlatforms
}

// GetExternalStoragePlatformsOk returns a tuple with the ExternalStoragePlatforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) GetExternalStoragePlatformsOk() ([]OSExternalStoragePlatformInfo, bool) {
	if o == nil || IsNil(o.ExternalStoragePlatforms) {
		return nil, false
	}
	return o.ExternalStoragePlatforms, true
}

// HasExternalStoragePlatforms returns a boolean if a field has been set.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) HasExternalStoragePlatforms() bool {
	if o != nil && !IsNil(o.ExternalStoragePlatforms) {
		return true
	}

	return false
}

// SetExternalStoragePlatforms gets a reference to the given []OSExternalStoragePlatformInfo and assigns it to the ExternalStoragePlatforms field.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) SetExternalStoragePlatforms(v []OSExternalStoragePlatformInfo) {
	o.ExternalStoragePlatforms = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) SetName(v string) {
	o.Name = &v
}

// GetPrefixEnabled returns the PrefixEnabled field value if set, zero value otherwise.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) GetPrefixEnabled() bool {
	if o == nil || IsNil(o.PrefixEnabled) {
		var ret bool
		return ret
	}
	return *o.PrefixEnabled
}

// GetPrefixEnabledOk returns a tuple with the PrefixEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) GetPrefixEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PrefixEnabled) {
		return nil, false
	}
	return o.PrefixEnabled, true
}

// HasPrefixEnabled returns a boolean if a field has been set.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) HasPrefixEnabled() bool {
	if o != nil && !IsNil(o.PrefixEnabled) {
		return true
	}

	return false
}

// SetPrefixEnabled gets a reference to the given bool and assigns it to the PrefixEnabled field.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) SetPrefixEnabled(v bool) {
	o.PrefixEnabled = &v
}

// GetStoragePattern returns the StoragePattern field value if set, zero value otherwise.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) GetStoragePattern() []string {
	if o == nil || IsNil(o.StoragePattern) {
		var ret []string
		return ret
	}
	return o.StoragePattern
}

// GetStoragePatternOk returns a tuple with the StoragePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) GetStoragePatternOk() ([]string, bool) {
	if o == nil || IsNil(o.StoragePattern) {
		return nil, false
	}
	return o.StoragePattern, true
}

// HasStoragePattern returns a boolean if a field has been set.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) HasStoragePattern() bool {
	if o != nil && !IsNil(o.StoragePattern) {
		return true
	}

	return false
}

// SetStoragePattern gets a reference to the given []string and assigns it to the StoragePattern field.
func (o *OSExternalStorageClassUpdateReqExternalStorageClass) SetStoragePattern(v []string) {
	o.StoragePattern = v
}

func (o OSExternalStorageClassUpdateReqExternalStorageClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSExternalStorageClassUpdateReqExternalStorageClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalStoragePlatforms) {
		toSerialize["external_storage_platforms"] = o.ExternalStoragePlatforms
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PrefixEnabled) {
		toSerialize["prefix_enabled"] = o.PrefixEnabled
	}
	if !IsNil(o.StoragePattern) {
		toSerialize["storage_pattern"] = o.StoragePattern
	}
	return toSerialize, nil
}

type NullableOSExternalStorageClassUpdateReqExternalStorageClass struct {
	value *OSExternalStorageClassUpdateReqExternalStorageClass
	isSet bool
}

func (v NullableOSExternalStorageClassUpdateReqExternalStorageClass) Get() *OSExternalStorageClassUpdateReqExternalStorageClass {
	return v.value
}

func (v *NullableOSExternalStorageClassUpdateReqExternalStorageClass) Set(val *OSExternalStorageClassUpdateReqExternalStorageClass) {
	v.value = val
	v.isSet = true
}

func (v NullableOSExternalStorageClassUpdateReqExternalStorageClass) IsSet() bool {
	return v.isSet
}

func (v *NullableOSExternalStorageClassUpdateReqExternalStorageClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSExternalStorageClassUpdateReqExternalStorageClass(val *OSExternalStorageClassUpdateReqExternalStorageClass) *NullableOSExternalStorageClassUpdateReqExternalStorageClass {
	return &NullableOSExternalStorageClassUpdateReqExternalStorageClass{value: val, isSet: true}
}

func (v NullableOSExternalStorageClassUpdateReqExternalStorageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSExternalStorageClassUpdateReqExternalStorageClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


