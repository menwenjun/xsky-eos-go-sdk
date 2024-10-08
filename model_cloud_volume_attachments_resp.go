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

// checks if the CloudVolumeAttachmentsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudVolumeAttachmentsResp{}

// CloudVolumeAttachmentsResp struct for CloudVolumeAttachmentsResp
type CloudVolumeAttachmentsResp struct {
	// cloud volume attachments
	CloudVolumeAttachments []CloudVolumeAttachment `json:"cloud_volume_attachments,omitempty"`
}

// NewCloudVolumeAttachmentsResp instantiates a new CloudVolumeAttachmentsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudVolumeAttachmentsResp() *CloudVolumeAttachmentsResp {
	this := CloudVolumeAttachmentsResp{}
	return &this
}

// NewCloudVolumeAttachmentsRespWithDefaults instantiates a new CloudVolumeAttachmentsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudVolumeAttachmentsRespWithDefaults() *CloudVolumeAttachmentsResp {
	this := CloudVolumeAttachmentsResp{}
	return &this
}

// GetCloudVolumeAttachments returns the CloudVolumeAttachments field value if set, zero value otherwise.
func (o *CloudVolumeAttachmentsResp) GetCloudVolumeAttachments() []CloudVolumeAttachment {
	if o == nil || IsNil(o.CloudVolumeAttachments) {
		var ret []CloudVolumeAttachment
		return ret
	}
	return o.CloudVolumeAttachments
}

// GetCloudVolumeAttachmentsOk returns a tuple with the CloudVolumeAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachmentsResp) GetCloudVolumeAttachmentsOk() ([]CloudVolumeAttachment, bool) {
	if o == nil || IsNil(o.CloudVolumeAttachments) {
		return nil, false
	}
	return o.CloudVolumeAttachments, true
}

// HasCloudVolumeAttachments returns a boolean if a field has been set.
func (o *CloudVolumeAttachmentsResp) HasCloudVolumeAttachments() bool {
	if o != nil && !IsNil(o.CloudVolumeAttachments) {
		return true
	}

	return false
}

// SetCloudVolumeAttachments gets a reference to the given []CloudVolumeAttachment and assigns it to the CloudVolumeAttachments field.
func (o *CloudVolumeAttachmentsResp) SetCloudVolumeAttachments(v []CloudVolumeAttachment) {
	o.CloudVolumeAttachments = v
}

func (o CloudVolumeAttachmentsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudVolumeAttachmentsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudVolumeAttachments) {
		toSerialize["cloud_volume_attachments"] = o.CloudVolumeAttachments
	}
	return toSerialize, nil
}

type NullableCloudVolumeAttachmentsResp struct {
	value *CloudVolumeAttachmentsResp
	isSet bool
}

func (v NullableCloudVolumeAttachmentsResp) Get() *CloudVolumeAttachmentsResp {
	return v.value
}

func (v *NullableCloudVolumeAttachmentsResp) Set(val *CloudVolumeAttachmentsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudVolumeAttachmentsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudVolumeAttachmentsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudVolumeAttachmentsResp(val *CloudVolumeAttachmentsResp) *NullableCloudVolumeAttachmentsResp {
	return &NullableCloudVolumeAttachmentsResp{value: val, isSet: true}
}

func (v NullableCloudVolumeAttachmentsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudVolumeAttachmentsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


