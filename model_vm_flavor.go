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

// checks if the VMFlavor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMFlavor{}

// VMFlavor VMFlavor defines flavor of vm
type VMFlavor struct {
	Cpu *int64 `json:"cpu,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Id *int64 `json:"id,omitempty"`
	MemoryKbyte *int64 `json:"memory_kbyte,omitempty"`
	Name *string `json:"name,omitempty"`
	RootDiskSize *int64 `json:"root_disk_size,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewVMFlavor instantiates a new VMFlavor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMFlavor() *VMFlavor {
	this := VMFlavor{}
	return &this
}

// NewVMFlavorWithDefaults instantiates a new VMFlavor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMFlavorWithDefaults() *VMFlavor {
	this := VMFlavor{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *VMFlavor) GetCpu() int64 {
	if o == nil || IsNil(o.Cpu) {
		var ret int64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMFlavor) GetCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *VMFlavor) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int64 and assigns it to the Cpu field.
func (o *VMFlavor) SetCpu(v int64) {
	o.Cpu = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *VMFlavor) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMFlavor) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *VMFlavor) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *VMFlavor) SetCreate(v time.Time) {
	o.Create = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VMFlavor) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMFlavor) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VMFlavor) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *VMFlavor) SetId(v int64) {
	o.Id = &v
}

// GetMemoryKbyte returns the MemoryKbyte field value if set, zero value otherwise.
func (o *VMFlavor) GetMemoryKbyte() int64 {
	if o == nil || IsNil(o.MemoryKbyte) {
		var ret int64
		return ret
	}
	return *o.MemoryKbyte
}

// GetMemoryKbyteOk returns a tuple with the MemoryKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMFlavor) GetMemoryKbyteOk() (*int64, bool) {
	if o == nil || IsNil(o.MemoryKbyte) {
		return nil, false
	}
	return o.MemoryKbyte, true
}

// HasMemoryKbyte returns a boolean if a field has been set.
func (o *VMFlavor) HasMemoryKbyte() bool {
	if o != nil && !IsNil(o.MemoryKbyte) {
		return true
	}

	return false
}

// SetMemoryKbyte gets a reference to the given int64 and assigns it to the MemoryKbyte field.
func (o *VMFlavor) SetMemoryKbyte(v int64) {
	o.MemoryKbyte = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VMFlavor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMFlavor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VMFlavor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VMFlavor) SetName(v string) {
	o.Name = &v
}

// GetRootDiskSize returns the RootDiskSize field value if set, zero value otherwise.
func (o *VMFlavor) GetRootDiskSize() int64 {
	if o == nil || IsNil(o.RootDiskSize) {
		var ret int64
		return ret
	}
	return *o.RootDiskSize
}

// GetRootDiskSizeOk returns a tuple with the RootDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMFlavor) GetRootDiskSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.RootDiskSize) {
		return nil, false
	}
	return o.RootDiskSize, true
}

// HasRootDiskSize returns a boolean if a field has been set.
func (o *VMFlavor) HasRootDiskSize() bool {
	if o != nil && !IsNil(o.RootDiskSize) {
		return true
	}

	return false
}

// SetRootDiskSize gets a reference to the given int64 and assigns it to the RootDiskSize field.
func (o *VMFlavor) SetRootDiskSize(v int64) {
	o.RootDiskSize = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *VMFlavor) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMFlavor) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *VMFlavor) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *VMFlavor) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o VMFlavor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMFlavor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MemoryKbyte) {
		toSerialize["memory_kbyte"] = o.MemoryKbyte
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RootDiskSize) {
		toSerialize["root_disk_size"] = o.RootDiskSize
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableVMFlavor struct {
	value *VMFlavor
	isSet bool
}

func (v NullableVMFlavor) Get() *VMFlavor {
	return v.value
}

func (v *NullableVMFlavor) Set(val *VMFlavor) {
	v.value = val
	v.isSet = true
}

func (v NullableVMFlavor) IsSet() bool {
	return v.isSet
}

func (v *NullableVMFlavor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMFlavor(val *VMFlavor) *NullableVMFlavor {
	return &NullableVMFlavor{value: val, isSet: true}
}

func (v NullableVMFlavor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMFlavor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


