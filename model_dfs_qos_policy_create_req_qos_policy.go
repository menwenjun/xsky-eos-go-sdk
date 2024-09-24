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

// checks if the DfsQosPolicyCreateReqQosPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsQosPolicyCreateReqQosPolicy{}

// DfsQosPolicyCreateReqQosPolicy struct for DfsQosPolicyCreateReqQosPolicy
type DfsQosPolicyCreateReqQosPolicy struct {
	// description if dfs qos policy
	Description *string `json:"description,omitempty"`
	// name of dfs qos policy
	Name string `json:"name"`
	// read bandwidth limitation
	ReadBandwidth *int64 `json:"read_bandwidth,omitempty"`
	// read objects limitation
	ReadObject *int64 `json:"read_object,omitempty"`
	// total bandwidth limitation
	TotalBandwidth *int64 `json:"total_bandwidth,omitempty"`
	// total objects limitation
	TotalObject *int64 `json:"total_object,omitempty"`
	// write bandwidth limitation
	WriteBandwidth *int64 `json:"write_bandwidth,omitempty"`
	// write objects limitation
	WriteObject *int64 `json:"write_object,omitempty"`
}

type _DfsQosPolicyCreateReqQosPolicy DfsQosPolicyCreateReqQosPolicy

// NewDfsQosPolicyCreateReqQosPolicy instantiates a new DfsQosPolicyCreateReqQosPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsQosPolicyCreateReqQosPolicy(name string) *DfsQosPolicyCreateReqQosPolicy {
	this := DfsQosPolicyCreateReqQosPolicy{}
	this.Name = name
	return &this
}

// NewDfsQosPolicyCreateReqQosPolicyWithDefaults instantiates a new DfsQosPolicyCreateReqQosPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsQosPolicyCreateReqQosPolicyWithDefaults() *DfsQosPolicyCreateReqQosPolicy {
	this := DfsQosPolicyCreateReqQosPolicy{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DfsQosPolicyCreateReqQosPolicy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DfsQosPolicyCreateReqQosPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *DfsQosPolicyCreateReqQosPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DfsQosPolicyCreateReqQosPolicy) SetName(v string) {
	o.Name = v
}

// GetReadBandwidth returns the ReadBandwidth field value if set, zero value otherwise.
func (o *DfsQosPolicyCreateReqQosPolicy) GetReadBandwidth() int64 {
	if o == nil || IsNil(o.ReadBandwidth) {
		var ret int64
		return ret
	}
	return *o.ReadBandwidth
}

// GetReadBandwidthOk returns a tuple with the ReadBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) GetReadBandwidthOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadBandwidth) {
		return nil, false
	}
	return o.ReadBandwidth, true
}

// HasReadBandwidth returns a boolean if a field has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) HasReadBandwidth() bool {
	if o != nil && !IsNil(o.ReadBandwidth) {
		return true
	}

	return false
}

// SetReadBandwidth gets a reference to the given int64 and assigns it to the ReadBandwidth field.
func (o *DfsQosPolicyCreateReqQosPolicy) SetReadBandwidth(v int64) {
	o.ReadBandwidth = &v
}

// GetReadObject returns the ReadObject field value if set, zero value otherwise.
func (o *DfsQosPolicyCreateReqQosPolicy) GetReadObject() int64 {
	if o == nil || IsNil(o.ReadObject) {
		var ret int64
		return ret
	}
	return *o.ReadObject
}

// GetReadObjectOk returns a tuple with the ReadObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) GetReadObjectOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadObject) {
		return nil, false
	}
	return o.ReadObject, true
}

// HasReadObject returns a boolean if a field has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) HasReadObject() bool {
	if o != nil && !IsNil(o.ReadObject) {
		return true
	}

	return false
}

// SetReadObject gets a reference to the given int64 and assigns it to the ReadObject field.
func (o *DfsQosPolicyCreateReqQosPolicy) SetReadObject(v int64) {
	o.ReadObject = &v
}

// GetTotalBandwidth returns the TotalBandwidth field value if set, zero value otherwise.
func (o *DfsQosPolicyCreateReqQosPolicy) GetTotalBandwidth() int64 {
	if o == nil || IsNil(o.TotalBandwidth) {
		var ret int64
		return ret
	}
	return *o.TotalBandwidth
}

// GetTotalBandwidthOk returns a tuple with the TotalBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) GetTotalBandwidthOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalBandwidth) {
		return nil, false
	}
	return o.TotalBandwidth, true
}

// HasTotalBandwidth returns a boolean if a field has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) HasTotalBandwidth() bool {
	if o != nil && !IsNil(o.TotalBandwidth) {
		return true
	}

	return false
}

// SetTotalBandwidth gets a reference to the given int64 and assigns it to the TotalBandwidth field.
func (o *DfsQosPolicyCreateReqQosPolicy) SetTotalBandwidth(v int64) {
	o.TotalBandwidth = &v
}

// GetTotalObject returns the TotalObject field value if set, zero value otherwise.
func (o *DfsQosPolicyCreateReqQosPolicy) GetTotalObject() int64 {
	if o == nil || IsNil(o.TotalObject) {
		var ret int64
		return ret
	}
	return *o.TotalObject
}

// GetTotalObjectOk returns a tuple with the TotalObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) GetTotalObjectOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalObject) {
		return nil, false
	}
	return o.TotalObject, true
}

// HasTotalObject returns a boolean if a field has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) HasTotalObject() bool {
	if o != nil && !IsNil(o.TotalObject) {
		return true
	}

	return false
}

// SetTotalObject gets a reference to the given int64 and assigns it to the TotalObject field.
func (o *DfsQosPolicyCreateReqQosPolicy) SetTotalObject(v int64) {
	o.TotalObject = &v
}

// GetWriteBandwidth returns the WriteBandwidth field value if set, zero value otherwise.
func (o *DfsQosPolicyCreateReqQosPolicy) GetWriteBandwidth() int64 {
	if o == nil || IsNil(o.WriteBandwidth) {
		var ret int64
		return ret
	}
	return *o.WriteBandwidth
}

// GetWriteBandwidthOk returns a tuple with the WriteBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) GetWriteBandwidthOk() (*int64, bool) {
	if o == nil || IsNil(o.WriteBandwidth) {
		return nil, false
	}
	return o.WriteBandwidth, true
}

// HasWriteBandwidth returns a boolean if a field has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) HasWriteBandwidth() bool {
	if o != nil && !IsNil(o.WriteBandwidth) {
		return true
	}

	return false
}

// SetWriteBandwidth gets a reference to the given int64 and assigns it to the WriteBandwidth field.
func (o *DfsQosPolicyCreateReqQosPolicy) SetWriteBandwidth(v int64) {
	o.WriteBandwidth = &v
}

// GetWriteObject returns the WriteObject field value if set, zero value otherwise.
func (o *DfsQosPolicyCreateReqQosPolicy) GetWriteObject() int64 {
	if o == nil || IsNil(o.WriteObject) {
		var ret int64
		return ret
	}
	return *o.WriteObject
}

// GetWriteObjectOk returns a tuple with the WriteObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) GetWriteObjectOk() (*int64, bool) {
	if o == nil || IsNil(o.WriteObject) {
		return nil, false
	}
	return o.WriteObject, true
}

// HasWriteObject returns a boolean if a field has been set.
func (o *DfsQosPolicyCreateReqQosPolicy) HasWriteObject() bool {
	if o != nil && !IsNil(o.WriteObject) {
		return true
	}

	return false
}

// SetWriteObject gets a reference to the given int64 and assigns it to the WriteObject field.
func (o *DfsQosPolicyCreateReqQosPolicy) SetWriteObject(v int64) {
	o.WriteObject = &v
}

func (o DfsQosPolicyCreateReqQosPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsQosPolicyCreateReqQosPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ReadBandwidth) {
		toSerialize["read_bandwidth"] = o.ReadBandwidth
	}
	if !IsNil(o.ReadObject) {
		toSerialize["read_object"] = o.ReadObject
	}
	if !IsNil(o.TotalBandwidth) {
		toSerialize["total_bandwidth"] = o.TotalBandwidth
	}
	if !IsNil(o.TotalObject) {
		toSerialize["total_object"] = o.TotalObject
	}
	if !IsNil(o.WriteBandwidth) {
		toSerialize["write_bandwidth"] = o.WriteBandwidth
	}
	if !IsNil(o.WriteObject) {
		toSerialize["write_object"] = o.WriteObject
	}
	return toSerialize, nil
}

func (o *DfsQosPolicyCreateReqQosPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varDfsQosPolicyCreateReqQosPolicy := _DfsQosPolicyCreateReqQosPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsQosPolicyCreateReqQosPolicy)

	if err != nil {
		return err
	}

	*o = DfsQosPolicyCreateReqQosPolicy(varDfsQosPolicyCreateReqQosPolicy)

	return err
}

type NullableDfsQosPolicyCreateReqQosPolicy struct {
	value *DfsQosPolicyCreateReqQosPolicy
	isSet bool
}

func (v NullableDfsQosPolicyCreateReqQosPolicy) Get() *DfsQosPolicyCreateReqQosPolicy {
	return v.value
}

func (v *NullableDfsQosPolicyCreateReqQosPolicy) Set(val *DfsQosPolicyCreateReqQosPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsQosPolicyCreateReqQosPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsQosPolicyCreateReqQosPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsQosPolicyCreateReqQosPolicy(val *DfsQosPolicyCreateReqQosPolicy) *NullableDfsQosPolicyCreateReqQosPolicy {
	return &NullableDfsQosPolicyCreateReqQosPolicy{value: val, isSet: true}
}

func (v NullableDfsQosPolicyCreateReqQosPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsQosPolicyCreateReqQosPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


