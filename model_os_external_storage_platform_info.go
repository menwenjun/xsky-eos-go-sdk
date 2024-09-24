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

// checks if the OSExternalStoragePlatformInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSExternalStoragePlatformInfo{}

// OSExternalStoragePlatformInfo OSExternalStoragePlatformInfo used for update os external storage class
type OSExternalStoragePlatformInfo struct {
	AccessKey *string `json:"access_key,omitempty"`
	Authentication *string `json:"authentication,omitempty"`
	BluRayId *int64 `json:"blu_ray_id,omitempty"`
	Buckets []OSExternalStorageBucketInfo `json:"buckets,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	HostStyle *string `json:"host_style,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Region *string `json:"region,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
	Weight *int64 `json:"weight,omitempty"`
}

// NewOSExternalStoragePlatformInfo instantiates a new OSExternalStoragePlatformInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSExternalStoragePlatformInfo() *OSExternalStoragePlatformInfo {
	this := OSExternalStoragePlatformInfo{}
	return &this
}

// NewOSExternalStoragePlatformInfoWithDefaults instantiates a new OSExternalStoragePlatformInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSExternalStoragePlatformInfoWithDefaults() *OSExternalStoragePlatformInfo {
	this := OSExternalStoragePlatformInfo{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *OSExternalStoragePlatformInfo) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStoragePlatformInfo) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *OSExternalStoragePlatformInfo) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *OSExternalStoragePlatformInfo) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *OSExternalStoragePlatformInfo) GetAuthentication() string {
	if o == nil || IsNil(o.Authentication) {
		var ret string
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStoragePlatformInfo) GetAuthenticationOk() (*string, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *OSExternalStoragePlatformInfo) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given string and assigns it to the Authentication field.
func (o *OSExternalStoragePlatformInfo) SetAuthentication(v string) {
	o.Authentication = &v
}

// GetBluRayId returns the BluRayId field value if set, zero value otherwise.
func (o *OSExternalStoragePlatformInfo) GetBluRayId() int64 {
	if o == nil || IsNil(o.BluRayId) {
		var ret int64
		return ret
	}
	return *o.BluRayId
}

// GetBluRayIdOk returns a tuple with the BluRayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStoragePlatformInfo) GetBluRayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BluRayId) {
		return nil, false
	}
	return o.BluRayId, true
}

// HasBluRayId returns a boolean if a field has been set.
func (o *OSExternalStoragePlatformInfo) HasBluRayId() bool {
	if o != nil && !IsNil(o.BluRayId) {
		return true
	}

	return false
}

// SetBluRayId gets a reference to the given int64 and assigns it to the BluRayId field.
func (o *OSExternalStoragePlatformInfo) SetBluRayId(v int64) {
	o.BluRayId = &v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *OSExternalStoragePlatformInfo) GetBuckets() []OSExternalStorageBucketInfo {
	if o == nil || IsNil(o.Buckets) {
		var ret []OSExternalStorageBucketInfo
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStoragePlatformInfo) GetBucketsOk() ([]OSExternalStorageBucketInfo, bool) {
	if o == nil || IsNil(o.Buckets) {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *OSExternalStoragePlatformInfo) HasBuckets() bool {
	if o != nil && !IsNil(o.Buckets) {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []OSExternalStorageBucketInfo and assigns it to the Buckets field.
func (o *OSExternalStoragePlatformInfo) SetBuckets(v []OSExternalStorageBucketInfo) {
	o.Buckets = v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *OSExternalStoragePlatformInfo) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStoragePlatformInfo) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *OSExternalStoragePlatformInfo) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *OSExternalStoragePlatformInfo) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetHostStyle returns the HostStyle field value if set, zero value otherwise.
func (o *OSExternalStoragePlatformInfo) GetHostStyle() string {
	if o == nil || IsNil(o.HostStyle) {
		var ret string
		return ret
	}
	return *o.HostStyle
}

// GetHostStyleOk returns a tuple with the HostStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStoragePlatformInfo) GetHostStyleOk() (*string, bool) {
	if o == nil || IsNil(o.HostStyle) {
		return nil, false
	}
	return o.HostStyle, true
}

// HasHostStyle returns a boolean if a field has been set.
func (o *OSExternalStoragePlatformInfo) HasHostStyle() bool {
	if o != nil && !IsNil(o.HostStyle) {
		return true
	}

	return false
}

// SetHostStyle gets a reference to the given string and assigns it to the HostStyle field.
func (o *OSExternalStoragePlatformInfo) SetHostStyle(v string) {
	o.HostStyle = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OSExternalStoragePlatformInfo) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStoragePlatformInfo) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OSExternalStoragePlatformInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OSExternalStoragePlatformInfo) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OSExternalStoragePlatformInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStoragePlatformInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OSExternalStoragePlatformInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OSExternalStoragePlatformInfo) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *OSExternalStoragePlatformInfo) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStoragePlatformInfo) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *OSExternalStoragePlatformInfo) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *OSExternalStoragePlatformInfo) SetRegion(v string) {
	o.Region = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *OSExternalStoragePlatformInfo) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStoragePlatformInfo) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *OSExternalStoragePlatformInfo) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *OSExternalStoragePlatformInfo) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *OSExternalStoragePlatformInfo) GetWeight() int64 {
	if o == nil || IsNil(o.Weight) {
		var ret int64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStoragePlatformInfo) GetWeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *OSExternalStoragePlatformInfo) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int64 and assigns it to the Weight field.
func (o *OSExternalStoragePlatformInfo) SetWeight(v int64) {
	o.Weight = &v
}

func (o OSExternalStoragePlatformInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSExternalStoragePlatformInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessKey) {
		toSerialize["access_key"] = o.AccessKey
	}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if !IsNil(o.BluRayId) {
		toSerialize["blu_ray_id"] = o.BluRayId
	}
	if !IsNil(o.Buckets) {
		toSerialize["buckets"] = o.Buckets
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.HostStyle) {
		toSerialize["host_style"] = o.HostStyle
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.SecretKey) {
		toSerialize["secret_key"] = o.SecretKey
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableOSExternalStoragePlatformInfo struct {
	value *OSExternalStoragePlatformInfo
	isSet bool
}

func (v NullableOSExternalStoragePlatformInfo) Get() *OSExternalStoragePlatformInfo {
	return v.value
}

func (v *NullableOSExternalStoragePlatformInfo) Set(val *OSExternalStoragePlatformInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOSExternalStoragePlatformInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOSExternalStoragePlatformInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSExternalStoragePlatformInfo(val *OSExternalStoragePlatformInfo) *NullableOSExternalStoragePlatformInfo {
	return &NullableOSExternalStoragePlatformInfo{value: val, isSet: true}
}

func (v NullableOSExternalStoragePlatformInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSExternalStoragePlatformInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


