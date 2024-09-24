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

// checks if the OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt{}

// OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt struct for OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt
type OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt struct {
	AccessKey *string `json:"access_key,omitempty"`
	Authentication *string `json:"authentication,omitempty"`
	BluRayId *int64 `json:"blu_ray_id,omitempty"`
	Buckets []OSExternalStorageBucketInfo `json:"buckets,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	HostStyle *string `json:"host_style,omitempty"`
	Name *string `json:"name,omitempty"`
	Region *string `json:"region,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
	Weight *int64 `json:"weight,omitempty"`
}

// NewOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt instantiates a new OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt() *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt {
	this := OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt{}
	return &this
}

// NewOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsEltWithDefaults instantiates a new OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsEltWithDefaults() *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt {
	this := OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetAuthentication() string {
	if o == nil || IsNil(o.Authentication) {
		var ret string
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetAuthenticationOk() (*string, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given string and assigns it to the Authentication field.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) SetAuthentication(v string) {
	o.Authentication = &v
}

// GetBluRayId returns the BluRayId field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetBluRayId() int64 {
	if o == nil || IsNil(o.BluRayId) {
		var ret int64
		return ret
	}
	return *o.BluRayId
}

// GetBluRayIdOk returns a tuple with the BluRayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetBluRayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BluRayId) {
		return nil, false
	}
	return o.BluRayId, true
}

// HasBluRayId returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) HasBluRayId() bool {
	if o != nil && !IsNil(o.BluRayId) {
		return true
	}

	return false
}

// SetBluRayId gets a reference to the given int64 and assigns it to the BluRayId field.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) SetBluRayId(v int64) {
	o.BluRayId = &v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetBuckets() []OSExternalStorageBucketInfo {
	if o == nil || IsNil(o.Buckets) {
		var ret []OSExternalStorageBucketInfo
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetBucketsOk() ([]OSExternalStorageBucketInfo, bool) {
	if o == nil || IsNil(o.Buckets) {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) HasBuckets() bool {
	if o != nil && !IsNil(o.Buckets) {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []OSExternalStorageBucketInfo and assigns it to the Buckets field.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) SetBuckets(v []OSExternalStorageBucketInfo) {
	o.Buckets = v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetHostStyle returns the HostStyle field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetHostStyle() string {
	if o == nil || IsNil(o.HostStyle) {
		var ret string
		return ret
	}
	return *o.HostStyle
}

// GetHostStyleOk returns a tuple with the HostStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetHostStyleOk() (*string, bool) {
	if o == nil || IsNil(o.HostStyle) {
		return nil, false
	}
	return o.HostStyle, true
}

// HasHostStyle returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) HasHostStyle() bool {
	if o != nil && !IsNil(o.HostStyle) {
		return true
	}

	return false
}

// SetHostStyle gets a reference to the given string and assigns it to the HostStyle field.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) SetHostStyle(v string) {
	o.HostStyle = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) SetRegion(v string) {
	o.Region = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetWeight() int64 {
	if o == nil || IsNil(o.Weight) {
		var ret int64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) GetWeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int64 and assigns it to the Weight field.
func (o *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) SetWeight(v int64) {
	o.Weight = &v
}

func (o OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) ToMap() (map[string]interface{}, error) {
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

type NullableOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt struct {
	value *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt
	isSet bool
}

func (v NullableOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) Get() *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt {
	return v.value
}

func (v *NullableOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) Set(val *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) {
	v.value = val
	v.isSet = true
}

func (v NullableOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) IsSet() bool {
	return v.isSet
}

func (v *NullableOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt(val *OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) *NullableOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt {
	return &NullableOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt{value: val, isSet: true}
}

func (v NullableOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


