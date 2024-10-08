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

// checks if the ClusterOverviewResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterOverviewResp{}

// ClusterOverviewResp struct for ClusterOverviewResp
type ClusterOverviewResp struct {
	ClusterStats ClusterOverview `json:"cluster_stats"`
}

type _ClusterOverviewResp ClusterOverviewResp

// NewClusterOverviewResp instantiates a new ClusterOverviewResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterOverviewResp(clusterStats ClusterOverview) *ClusterOverviewResp {
	this := ClusterOverviewResp{}
	this.ClusterStats = clusterStats
	return &this
}

// NewClusterOverviewRespWithDefaults instantiates a new ClusterOverviewResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterOverviewRespWithDefaults() *ClusterOverviewResp {
	this := ClusterOverviewResp{}
	return &this
}

// GetClusterStats returns the ClusterStats field value
func (o *ClusterOverviewResp) GetClusterStats() ClusterOverview {
	if o == nil {
		var ret ClusterOverview
		return ret
	}

	return o.ClusterStats
}

// GetClusterStatsOk returns a tuple with the ClusterStats field value
// and a boolean to check if the value has been set.
func (o *ClusterOverviewResp) GetClusterStatsOk() (*ClusterOverview, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterStats, true
}

// SetClusterStats sets field value
func (o *ClusterOverviewResp) SetClusterStats(v ClusterOverview) {
	o.ClusterStats = v
}

func (o ClusterOverviewResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterOverviewResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster_stats"] = o.ClusterStats
	return toSerialize, nil
}

func (o *ClusterOverviewResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cluster_stats",
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

	varClusterOverviewResp := _ClusterOverviewResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClusterOverviewResp)

	if err != nil {
		return err
	}

	*o = ClusterOverviewResp(varClusterOverviewResp)

	return err
}

type NullableClusterOverviewResp struct {
	value *ClusterOverviewResp
	isSet bool
}

func (v NullableClusterOverviewResp) Get() *ClusterOverviewResp {
	return v.value
}

func (v *NullableClusterOverviewResp) Set(val *ClusterOverviewResp) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterOverviewResp) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterOverviewResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterOverviewResp(val *ClusterOverviewResp) *NullableClusterOverviewResp {
	return &NullableClusterOverviewResp{value: val, isSet: true}
}

func (v NullableClusterOverviewResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterOverviewResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


