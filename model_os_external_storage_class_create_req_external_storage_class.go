/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OSExternalStorageClassCreateReqExternalStorageClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSExternalStorageClassCreateReqExternalStorageClass{}

// OSExternalStorageClassCreateReqExternalStorageClass struct for OSExternalStorageClassCreateReqExternalStorageClass
type OSExternalStorageClassCreateReqExternalStorageClass struct {
	ClassId string `json:"class_id"`
	Description *string `json:"description,omitempty"`
	ExternalStoragePlatforms []OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt `json:"external_storage_platforms"`
	Name string `json:"name"`
	OsPolicyId int64 `json:"os_policy_id"`
	Platform string `json:"platform"`
	PlatformType string `json:"platform_type"`
	PrefixEnabled *bool `json:"prefix_enabled,omitempty"`
	// SliceStringField defines slice string field
	StoragePattern []string `json:"storage_pattern,omitempty"`
}

type _OSExternalStorageClassCreateReqExternalStorageClass OSExternalStorageClassCreateReqExternalStorageClass

// NewOSExternalStorageClassCreateReqExternalStorageClass instantiates a new OSExternalStorageClassCreateReqExternalStorageClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSExternalStorageClassCreateReqExternalStorageClass(classId string, externalStoragePlatforms []OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt, name string, osPolicyId int64, platform string, platformType string) *OSExternalStorageClassCreateReqExternalStorageClass {
	this := OSExternalStorageClassCreateReqExternalStorageClass{}
	this.ClassId = classId
	this.ExternalStoragePlatforms = externalStoragePlatforms
	this.Name = name
	this.OsPolicyId = osPolicyId
	this.Platform = platform
	this.PlatformType = platformType
	return &this
}

// NewOSExternalStorageClassCreateReqExternalStorageClassWithDefaults instantiates a new OSExternalStorageClassCreateReqExternalStorageClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSExternalStorageClassCreateReqExternalStorageClassWithDefaults() *OSExternalStorageClassCreateReqExternalStorageClass {
	this := OSExternalStorageClassCreateReqExternalStorageClass{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetClassId(v string) {
	o.ClassId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetDescription(v string) {
	o.Description = &v
}

// GetExternalStoragePlatforms returns the ExternalStoragePlatforms field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetExternalStoragePlatforms() []OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt {
	if o == nil {
		var ret []OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt
		return ret
	}

	return o.ExternalStoragePlatforms
}

// GetExternalStoragePlatformsOk returns a tuple with the ExternalStoragePlatforms field value
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetExternalStoragePlatformsOk() ([]OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalStoragePlatforms, true
}

// SetExternalStoragePlatforms sets field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetExternalStoragePlatforms(v []OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) {
	o.ExternalStoragePlatforms = v
}

// GetName returns the Name field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetName(v string) {
	o.Name = v
}

// GetOsPolicyId returns the OsPolicyId field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetOsPolicyId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OsPolicyId
}

// GetOsPolicyIdOk returns a tuple with the OsPolicyId field value
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetOsPolicyIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsPolicyId, true
}

// SetOsPolicyId sets field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetOsPolicyId(v int64) {
	o.OsPolicyId = v
}

// GetPlatform returns the Platform field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetPlatform(v string) {
	o.Platform = v
}

// GetPlatformType returns the PlatformType field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPlatformType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPlatformTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlatformType, true
}

// SetPlatformType sets field value
func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetPlatformType(v string) {
	o.PlatformType = v
}

// GetPrefixEnabled returns the PrefixEnabled field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPrefixEnabled() bool {
	if o == nil || IsNil(o.PrefixEnabled) {
		var ret bool
		return ret
	}
	return *o.PrefixEnabled
}

// GetPrefixEnabledOk returns a tuple with the PrefixEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPrefixEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PrefixEnabled) {
		return nil, false
	}
	return o.PrefixEnabled, true
}

// HasPrefixEnabled returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) HasPrefixEnabled() bool {
	if o != nil && !IsNil(o.PrefixEnabled) {
		return true
	}

	return false
}

// SetPrefixEnabled gets a reference to the given bool and assigns it to the PrefixEnabled field.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetPrefixEnabled(v bool) {
	o.PrefixEnabled = &v
}

// GetStoragePattern returns the StoragePattern field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetStoragePattern() []string {
	if o == nil || IsNil(o.StoragePattern) {
		var ret []string
		return ret
	}
	return o.StoragePattern
}

// GetStoragePatternOk returns a tuple with the StoragePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetStoragePatternOk() ([]string, bool) {
	if o == nil || IsNil(o.StoragePattern) {
		return nil, false
	}
	return o.StoragePattern, true
}

// HasStoragePattern returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) HasStoragePattern() bool {
	if o != nil && !IsNil(o.StoragePattern) {
		return true
	}

	return false
}

// SetStoragePattern gets a reference to the given []string and assigns it to the StoragePattern field.
func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetStoragePattern(v []string) {
	o.StoragePattern = v
}

func (o OSExternalStorageClassCreateReqExternalStorageClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSExternalStorageClassCreateReqExternalStorageClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["class_id"] = o.ClassId
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["external_storage_platforms"] = o.ExternalStoragePlatforms
	toSerialize["name"] = o.Name
	toSerialize["os_policy_id"] = o.OsPolicyId
	toSerialize["platform"] = o.Platform
	toSerialize["platform_type"] = o.PlatformType
	if !IsNil(o.PrefixEnabled) {
		toSerialize["prefix_enabled"] = o.PrefixEnabled
	}
	if !IsNil(o.StoragePattern) {
		toSerialize["storage_pattern"] = o.StoragePattern
	}
	return toSerialize, nil
}

func (o *OSExternalStorageClassCreateReqExternalStorageClass) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"class_id",
		"external_storage_platforms",
		"name",
		"os_policy_id",
		"platform",
		"platform_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOSExternalStorageClassCreateReqExternalStorageClass := _OSExternalStorageClassCreateReqExternalStorageClass{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOSExternalStorageClassCreateReqExternalStorageClass)

	if err != nil {
		return err
	}

	*o = OSExternalStorageClassCreateReqExternalStorageClass(varOSExternalStorageClassCreateReqExternalStorageClass)

	return err
}

type NullableOSExternalStorageClassCreateReqExternalStorageClass struct {
	value *OSExternalStorageClassCreateReqExternalStorageClass
	isSet bool
}

func (v NullableOSExternalStorageClassCreateReqExternalStorageClass) Get() *OSExternalStorageClassCreateReqExternalStorageClass {
	return v.value
}

func (v *NullableOSExternalStorageClassCreateReqExternalStorageClass) Set(val *OSExternalStorageClassCreateReqExternalStorageClass) {
	v.value = val
	v.isSet = true
}

func (v NullableOSExternalStorageClassCreateReqExternalStorageClass) IsSet() bool {
	return v.isSet
}

func (v *NullableOSExternalStorageClassCreateReqExternalStorageClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSExternalStorageClassCreateReqExternalStorageClass(val *OSExternalStorageClassCreateReqExternalStorageClass) *NullableOSExternalStorageClassCreateReqExternalStorageClass {
	return &NullableOSExternalStorageClassCreateReqExternalStorageClass{value: val, isSet: true}
}

func (v NullableOSExternalStorageClassCreateReqExternalStorageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSExternalStorageClassCreateReqExternalStorageClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


