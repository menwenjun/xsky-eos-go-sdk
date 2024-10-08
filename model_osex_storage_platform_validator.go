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

// checks if the OSExStoragePlatformValidator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSExStoragePlatformValidator{}

// OSExStoragePlatformValidator OSExStoragePlatformValidator is the model of os external storage platform valdator +X:model:generate;check_get +X:benchmark:
type OSExStoragePlatformValidator struct {
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Id *int64 `json:"id,omitempty"`
	PlatformInfo map[string]interface{} `json:"platform_info,omitempty"`
	Results []map[string]interface{} `json:"results,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewOSExStoragePlatformValidator instantiates a new OSExStoragePlatformValidator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSExStoragePlatformValidator() *OSExStoragePlatformValidator {
	this := OSExStoragePlatformValidator{}
	return &this
}

// NewOSExStoragePlatformValidatorWithDefaults instantiates a new OSExStoragePlatformValidator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSExStoragePlatformValidatorWithDefaults() *OSExStoragePlatformValidator {
	this := OSExStoragePlatformValidator{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *OSExStoragePlatformValidator) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExStoragePlatformValidator) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *OSExStoragePlatformValidator) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *OSExStoragePlatformValidator) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *OSExStoragePlatformValidator) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExStoragePlatformValidator) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *OSExStoragePlatformValidator) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *OSExStoragePlatformValidator) SetCreate(v time.Time) {
	o.Create = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OSExStoragePlatformValidator) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExStoragePlatformValidator) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OSExStoragePlatformValidator) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OSExStoragePlatformValidator) SetId(v int64) {
	o.Id = &v
}

// GetPlatformInfo returns the PlatformInfo field value if set, zero value otherwise.
func (o *OSExStoragePlatformValidator) GetPlatformInfo() map[string]interface{} {
	if o == nil || IsNil(o.PlatformInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.PlatformInfo
}

// GetPlatformInfoOk returns a tuple with the PlatformInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExStoragePlatformValidator) GetPlatformInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PlatformInfo) {
		return map[string]interface{}{}, false
	}
	return o.PlatformInfo, true
}

// HasPlatformInfo returns a boolean if a field has been set.
func (o *OSExStoragePlatformValidator) HasPlatformInfo() bool {
	if o != nil && !IsNil(o.PlatformInfo) {
		return true
	}

	return false
}

// SetPlatformInfo gets a reference to the given map[string]interface{} and assigns it to the PlatformInfo field.
func (o *OSExStoragePlatformValidator) SetPlatformInfo(v map[string]interface{}) {
	o.PlatformInfo = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *OSExStoragePlatformValidator) GetResults() []map[string]interface{} {
	if o == nil || IsNil(o.Results) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExStoragePlatformValidator) GetResultsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *OSExStoragePlatformValidator) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []map[string]interface{} and assigns it to the Results field.
func (o *OSExStoragePlatformValidator) SetResults(v []map[string]interface{}) {
	o.Results = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OSExStoragePlatformValidator) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSExStoragePlatformValidator) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OSExStoragePlatformValidator) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OSExStoragePlatformValidator) SetStatus(v string) {
	o.Status = &v
}

func (o OSExStoragePlatformValidator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSExStoragePlatformValidator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PlatformInfo) {
		toSerialize["platform_info"] = o.PlatformInfo
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableOSExStoragePlatformValidator struct {
	value *OSExStoragePlatformValidator
	isSet bool
}

func (v NullableOSExStoragePlatformValidator) Get() *OSExStoragePlatformValidator {
	return v.value
}

func (v *NullableOSExStoragePlatformValidator) Set(val *OSExStoragePlatformValidator) {
	v.value = val
	v.isSet = true
}

func (v NullableOSExStoragePlatformValidator) IsSet() bool {
	return v.isSet
}

func (v *NullableOSExStoragePlatformValidator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSExStoragePlatformValidator(val *OSExStoragePlatformValidator) *NullableOSExStoragePlatformValidator {
	return &NullableOSExStoragePlatformValidator{value: val, isSet: true}
}

func (v NullableOSExStoragePlatformValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSExStoragePlatformValidator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


