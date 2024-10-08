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

// checks if the DiskNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskNestview{}

// DiskNestview struct for DiskNestview
type DiskNestview struct {
	Device *string `json:"device,omitempty"`
	DiskType *string `json:"disk_type,omitempty"`
	Id *int64 `json:"id,omitempty"`
	SlotId *string `json:"slot_id,omitempty"`
	SsdLifeLeft *int64 `json:"ssd_life_left,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewDiskNestview instantiates a new DiskNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskNestview() *DiskNestview {
	this := DiskNestview{}
	return &this
}

// NewDiskNestviewWithDefaults instantiates a new DiskNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskNestviewWithDefaults() *DiskNestview {
	this := DiskNestview{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *DiskNestview) GetDevice() string {
	if o == nil || IsNil(o.Device) {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskNestview) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *DiskNestview) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *DiskNestview) SetDevice(v string) {
	o.Device = &v
}

// GetDiskType returns the DiskType field value if set, zero value otherwise.
func (o *DiskNestview) GetDiskType() string {
	if o == nil || IsNil(o.DiskType) {
		var ret string
		return ret
	}
	return *o.DiskType
}

// GetDiskTypeOk returns a tuple with the DiskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskNestview) GetDiskTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DiskType) {
		return nil, false
	}
	return o.DiskType, true
}

// HasDiskType returns a boolean if a field has been set.
func (o *DiskNestview) HasDiskType() bool {
	if o != nil && !IsNil(o.DiskType) {
		return true
	}

	return false
}

// SetDiskType gets a reference to the given string and assigns it to the DiskType field.
func (o *DiskNestview) SetDiskType(v string) {
	o.DiskType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiskNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiskNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DiskNestview) SetId(v int64) {
	o.Id = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *DiskNestview) GetSlotId() string {
	if o == nil || IsNil(o.SlotId) {
		var ret string
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskNestview) GetSlotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SlotId) {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *DiskNestview) HasSlotId() bool {
	if o != nil && !IsNil(o.SlotId) {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given string and assigns it to the SlotId field.
func (o *DiskNestview) SetSlotId(v string) {
	o.SlotId = &v
}

// GetSsdLifeLeft returns the SsdLifeLeft field value if set, zero value otherwise.
func (o *DiskNestview) GetSsdLifeLeft() int64 {
	if o == nil || IsNil(o.SsdLifeLeft) {
		var ret int64
		return ret
	}
	return *o.SsdLifeLeft
}

// GetSsdLifeLeftOk returns a tuple with the SsdLifeLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskNestview) GetSsdLifeLeftOk() (*int64, bool) {
	if o == nil || IsNil(o.SsdLifeLeft) {
		return nil, false
	}
	return o.SsdLifeLeft, true
}

// HasSsdLifeLeft returns a boolean if a field has been set.
func (o *DiskNestview) HasSsdLifeLeft() bool {
	if o != nil && !IsNil(o.SsdLifeLeft) {
		return true
	}

	return false
}

// SetSsdLifeLeft gets a reference to the given int64 and assigns it to the SsdLifeLeft field.
func (o *DiskNestview) SetSsdLifeLeft(v int64) {
	o.SsdLifeLeft = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DiskNestview) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskNestview) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DiskNestview) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DiskNestview) SetStatus(v string) {
	o.Status = &v
}

func (o DiskNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.DiskType) {
		toSerialize["disk_type"] = o.DiskType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SlotId) {
		toSerialize["slot_id"] = o.SlotId
	}
	if !IsNil(o.SsdLifeLeft) {
		toSerialize["ssd_life_left"] = o.SsdLifeLeft
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableDiskNestview struct {
	value *DiskNestview
	isSet bool
}

func (v NullableDiskNestview) Get() *DiskNestview {
	return v.value
}

func (v *NullableDiskNestview) Set(val *DiskNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskNestview(val *DiskNestview) *NullableDiskNestview {
	return &NullableDiskNestview{value: val, isSet: true}
}

func (v NullableDiskNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


