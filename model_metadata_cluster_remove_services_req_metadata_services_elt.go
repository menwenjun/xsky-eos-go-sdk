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

// checks if the MetadataClusterRemoveServicesReqMetadataServicesElt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataClusterRemoveServicesReqMetadataServicesElt{}

// MetadataClusterRemoveServicesReqMetadataServicesElt struct for MetadataClusterRemoveServicesReqMetadataServicesElt
type MetadataClusterRemoveServicesReqMetadataServicesElt struct {
	// MetadataServiceID is the metadata service id
	MetadataServiceId *int64 `json:"metadata_service_id,omitempty"`
}

// NewMetadataClusterRemoveServicesReqMetadataServicesElt instantiates a new MetadataClusterRemoveServicesReqMetadataServicesElt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataClusterRemoveServicesReqMetadataServicesElt() *MetadataClusterRemoveServicesReqMetadataServicesElt {
	this := MetadataClusterRemoveServicesReqMetadataServicesElt{}
	return &this
}

// NewMetadataClusterRemoveServicesReqMetadataServicesEltWithDefaults instantiates a new MetadataClusterRemoveServicesReqMetadataServicesElt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataClusterRemoveServicesReqMetadataServicesEltWithDefaults() *MetadataClusterRemoveServicesReqMetadataServicesElt {
	this := MetadataClusterRemoveServicesReqMetadataServicesElt{}
	return &this
}

// GetMetadataServiceId returns the MetadataServiceId field value if set, zero value otherwise.
func (o *MetadataClusterRemoveServicesReqMetadataServicesElt) GetMetadataServiceId() int64 {
	if o == nil || IsNil(o.MetadataServiceId) {
		var ret int64
		return ret
	}
	return *o.MetadataServiceId
}

// GetMetadataServiceIdOk returns a tuple with the MetadataServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataClusterRemoveServicesReqMetadataServicesElt) GetMetadataServiceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MetadataServiceId) {
		return nil, false
	}
	return o.MetadataServiceId, true
}

// HasMetadataServiceId returns a boolean if a field has been set.
func (o *MetadataClusterRemoveServicesReqMetadataServicesElt) HasMetadataServiceId() bool {
	if o != nil && !IsNil(o.MetadataServiceId) {
		return true
	}

	return false
}

// SetMetadataServiceId gets a reference to the given int64 and assigns it to the MetadataServiceId field.
func (o *MetadataClusterRemoveServicesReqMetadataServicesElt) SetMetadataServiceId(v int64) {
	o.MetadataServiceId = &v
}

func (o MetadataClusterRemoveServicesReqMetadataServicesElt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataClusterRemoveServicesReqMetadataServicesElt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetadataServiceId) {
		toSerialize["metadata_service_id"] = o.MetadataServiceId
	}
	return toSerialize, nil
}

type NullableMetadataClusterRemoveServicesReqMetadataServicesElt struct {
	value *MetadataClusterRemoveServicesReqMetadataServicesElt
	isSet bool
}

func (v NullableMetadataClusterRemoveServicesReqMetadataServicesElt) Get() *MetadataClusterRemoveServicesReqMetadataServicesElt {
	return v.value
}

func (v *NullableMetadataClusterRemoveServicesReqMetadataServicesElt) Set(val *MetadataClusterRemoveServicesReqMetadataServicesElt) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataClusterRemoveServicesReqMetadataServicesElt) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataClusterRemoveServicesReqMetadataServicesElt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataClusterRemoveServicesReqMetadataServicesElt(val *MetadataClusterRemoveServicesReqMetadataServicesElt) *NullableMetadataClusterRemoveServicesReqMetadataServicesElt {
	return &NullableMetadataClusterRemoveServicesReqMetadataServicesElt{value: val, isSet: true}
}

func (v NullableMetadataClusterRemoveServicesReqMetadataServicesElt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataClusterRemoveServicesReqMetadataServicesElt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


