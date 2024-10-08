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

// checks if the MetadataClusterCreateReqMetadataCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataClusterCreateReqMetadataCluster{}

// MetadataClusterCreateReqMetadataCluster struct for MetadataClusterCreateReqMetadataCluster
type MetadataClusterCreateReqMetadataCluster struct {
	// MetadataServices is the disk info of metadata cluster
	MetadataServices []MetadataClusterCreateReqMetadataClusterMetadataServicesElt `json:"metadata_services,omitempty"`
	// MinDiskInDC is the minimum disk number in data center
	MinDiskInDc *int64 `json:"min_disk_in_dc,omitempty"`
	// Name is the name of metadata cluster
	Name *string `json:"name,omitempty"`
	// PrimaryPlacementNodeID is the primary placement node id of metadata cluster
	PrimaryPlacementNodeId *int64 `json:"primary_placement_node_id,omitempty"`
}

// NewMetadataClusterCreateReqMetadataCluster instantiates a new MetadataClusterCreateReqMetadataCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataClusterCreateReqMetadataCluster() *MetadataClusterCreateReqMetadataCluster {
	this := MetadataClusterCreateReqMetadataCluster{}
	return &this
}

// NewMetadataClusterCreateReqMetadataClusterWithDefaults instantiates a new MetadataClusterCreateReqMetadataCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataClusterCreateReqMetadataClusterWithDefaults() *MetadataClusterCreateReqMetadataCluster {
	this := MetadataClusterCreateReqMetadataCluster{}
	return &this
}

// GetMetadataServices returns the MetadataServices field value if set, zero value otherwise.
func (o *MetadataClusterCreateReqMetadataCluster) GetMetadataServices() []MetadataClusterCreateReqMetadataClusterMetadataServicesElt {
	if o == nil || IsNil(o.MetadataServices) {
		var ret []MetadataClusterCreateReqMetadataClusterMetadataServicesElt
		return ret
	}
	return o.MetadataServices
}

// GetMetadataServicesOk returns a tuple with the MetadataServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataClusterCreateReqMetadataCluster) GetMetadataServicesOk() ([]MetadataClusterCreateReqMetadataClusterMetadataServicesElt, bool) {
	if o == nil || IsNil(o.MetadataServices) {
		return nil, false
	}
	return o.MetadataServices, true
}

// HasMetadataServices returns a boolean if a field has been set.
func (o *MetadataClusterCreateReqMetadataCluster) HasMetadataServices() bool {
	if o != nil && !IsNil(o.MetadataServices) {
		return true
	}

	return false
}

// SetMetadataServices gets a reference to the given []MetadataClusterCreateReqMetadataClusterMetadataServicesElt and assigns it to the MetadataServices field.
func (o *MetadataClusterCreateReqMetadataCluster) SetMetadataServices(v []MetadataClusterCreateReqMetadataClusterMetadataServicesElt) {
	o.MetadataServices = v
}

// GetMinDiskInDc returns the MinDiskInDc field value if set, zero value otherwise.
func (o *MetadataClusterCreateReqMetadataCluster) GetMinDiskInDc() int64 {
	if o == nil || IsNil(o.MinDiskInDc) {
		var ret int64
		return ret
	}
	return *o.MinDiskInDc
}

// GetMinDiskInDcOk returns a tuple with the MinDiskInDc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataClusterCreateReqMetadataCluster) GetMinDiskInDcOk() (*int64, bool) {
	if o == nil || IsNil(o.MinDiskInDc) {
		return nil, false
	}
	return o.MinDiskInDc, true
}

// HasMinDiskInDc returns a boolean if a field has been set.
func (o *MetadataClusterCreateReqMetadataCluster) HasMinDiskInDc() bool {
	if o != nil && !IsNil(o.MinDiskInDc) {
		return true
	}

	return false
}

// SetMinDiskInDc gets a reference to the given int64 and assigns it to the MinDiskInDc field.
func (o *MetadataClusterCreateReqMetadataCluster) SetMinDiskInDc(v int64) {
	o.MinDiskInDc = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetadataClusterCreateReqMetadataCluster) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataClusterCreateReqMetadataCluster) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetadataClusterCreateReqMetadataCluster) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetadataClusterCreateReqMetadataCluster) SetName(v string) {
	o.Name = &v
}

// GetPrimaryPlacementNodeId returns the PrimaryPlacementNodeId field value if set, zero value otherwise.
func (o *MetadataClusterCreateReqMetadataCluster) GetPrimaryPlacementNodeId() int64 {
	if o == nil || IsNil(o.PrimaryPlacementNodeId) {
		var ret int64
		return ret
	}
	return *o.PrimaryPlacementNodeId
}

// GetPrimaryPlacementNodeIdOk returns a tuple with the PrimaryPlacementNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataClusterCreateReqMetadataCluster) GetPrimaryPlacementNodeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PrimaryPlacementNodeId) {
		return nil, false
	}
	return o.PrimaryPlacementNodeId, true
}

// HasPrimaryPlacementNodeId returns a boolean if a field has been set.
func (o *MetadataClusterCreateReqMetadataCluster) HasPrimaryPlacementNodeId() bool {
	if o != nil && !IsNil(o.PrimaryPlacementNodeId) {
		return true
	}

	return false
}

// SetPrimaryPlacementNodeId gets a reference to the given int64 and assigns it to the PrimaryPlacementNodeId field.
func (o *MetadataClusterCreateReqMetadataCluster) SetPrimaryPlacementNodeId(v int64) {
	o.PrimaryPlacementNodeId = &v
}

func (o MetadataClusterCreateReqMetadataCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataClusterCreateReqMetadataCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetadataServices) {
		toSerialize["metadata_services"] = o.MetadataServices
	}
	if !IsNil(o.MinDiskInDc) {
		toSerialize["min_disk_in_dc"] = o.MinDiskInDc
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PrimaryPlacementNodeId) {
		toSerialize["primary_placement_node_id"] = o.PrimaryPlacementNodeId
	}
	return toSerialize, nil
}

type NullableMetadataClusterCreateReqMetadataCluster struct {
	value *MetadataClusterCreateReqMetadataCluster
	isSet bool
}

func (v NullableMetadataClusterCreateReqMetadataCluster) Get() *MetadataClusterCreateReqMetadataCluster {
	return v.value
}

func (v *NullableMetadataClusterCreateReqMetadataCluster) Set(val *MetadataClusterCreateReqMetadataCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataClusterCreateReqMetadataCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataClusterCreateReqMetadataCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataClusterCreateReqMetadataCluster(val *MetadataClusterCreateReqMetadataCluster) *NullableMetadataClusterCreateReqMetadataCluster {
	return &NullableMetadataClusterCreateReqMetadataCluster{value: val, isSet: true}
}

func (v NullableMetadataClusterCreateReqMetadataCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataClusterCreateReqMetadataCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


