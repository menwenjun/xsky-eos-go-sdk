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

// checks if the DfsSnapShotsOverviewPageResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsSnapShotsOverviewPageResp{}

// DfsSnapShotsOverviewPageResp struct for DfsSnapShotsOverviewPageResp
type DfsSnapShotsOverviewPageResp struct {
	DfsSnapshotsOverviewPage DfsSnapshotsOverviewPage `json:"dfs_snapshots_overview_page"`
}

type _DfsSnapShotsOverviewPageResp DfsSnapShotsOverviewPageResp

// NewDfsSnapShotsOverviewPageResp instantiates a new DfsSnapShotsOverviewPageResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsSnapShotsOverviewPageResp(dfsSnapshotsOverviewPage DfsSnapshotsOverviewPage) *DfsSnapShotsOverviewPageResp {
	this := DfsSnapShotsOverviewPageResp{}
	this.DfsSnapshotsOverviewPage = dfsSnapshotsOverviewPage
	return &this
}

// NewDfsSnapShotsOverviewPageRespWithDefaults instantiates a new DfsSnapShotsOverviewPageResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsSnapShotsOverviewPageRespWithDefaults() *DfsSnapShotsOverviewPageResp {
	this := DfsSnapShotsOverviewPageResp{}
	return &this
}

// GetDfsSnapshotsOverviewPage returns the DfsSnapshotsOverviewPage field value
func (o *DfsSnapShotsOverviewPageResp) GetDfsSnapshotsOverviewPage() DfsSnapshotsOverviewPage {
	if o == nil {
		var ret DfsSnapshotsOverviewPage
		return ret
	}

	return o.DfsSnapshotsOverviewPage
}

// GetDfsSnapshotsOverviewPageOk returns a tuple with the DfsSnapshotsOverviewPage field value
// and a boolean to check if the value has been set.
func (o *DfsSnapShotsOverviewPageResp) GetDfsSnapshotsOverviewPageOk() (*DfsSnapshotsOverviewPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsSnapshotsOverviewPage, true
}

// SetDfsSnapshotsOverviewPage sets field value
func (o *DfsSnapShotsOverviewPageResp) SetDfsSnapshotsOverviewPage(v DfsSnapshotsOverviewPage) {
	o.DfsSnapshotsOverviewPage = v
}

func (o DfsSnapShotsOverviewPageResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsSnapShotsOverviewPageResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dfs_snapshots_overview_page"] = o.DfsSnapshotsOverviewPage
	return toSerialize, nil
}

func (o *DfsSnapShotsOverviewPageResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dfs_snapshots_overview_page",
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

	varDfsSnapShotsOverviewPageResp := _DfsSnapShotsOverviewPageResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsSnapShotsOverviewPageResp)

	if err != nil {
		return err
	}

	*o = DfsSnapShotsOverviewPageResp(varDfsSnapShotsOverviewPageResp)

	return err
}

type NullableDfsSnapShotsOverviewPageResp struct {
	value *DfsSnapShotsOverviewPageResp
	isSet bool
}

func (v NullableDfsSnapShotsOverviewPageResp) Get() *DfsSnapShotsOverviewPageResp {
	return v.value
}

func (v *NullableDfsSnapShotsOverviewPageResp) Set(val *DfsSnapShotsOverviewPageResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsSnapShotsOverviewPageResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsSnapShotsOverviewPageResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsSnapShotsOverviewPageResp(val *DfsSnapShotsOverviewPageResp) *NullableDfsSnapShotsOverviewPageResp {
	return &NullableDfsSnapShotsOverviewPageResp{value: val, isSet: true}
}

func (v NullableDfsSnapShotsOverviewPageResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsSnapShotsOverviewPageResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


