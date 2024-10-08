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

// checks if the DfsQosPolicyUpdateReqQosPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsQosPolicyUpdateReqQosPolicy{}

// DfsQosPolicyUpdateReqQosPolicy struct for DfsQosPolicyUpdateReqQosPolicy
type DfsQosPolicyUpdateReqQosPolicy struct {
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

type _DfsQosPolicyUpdateReqQosPolicy DfsQosPolicyUpdateReqQosPolicy

// NewDfsQosPolicyUpdateReqQosPolicy instantiates a new DfsQosPolicyUpdateReqQosPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsQosPolicyUpdateReqQosPolicy(name string) *DfsQosPolicyUpdateReqQosPolicy {
	this := DfsQosPolicyUpdateReqQosPolicy{}
	this.Name = name
	return &this
}

// NewDfsQosPolicyUpdateReqQosPolicyWithDefaults instantiates a new DfsQosPolicyUpdateReqQosPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsQosPolicyUpdateReqQosPolicyWithDefaults() *DfsQosPolicyUpdateReqQosPolicy {
	this := DfsQosPolicyUpdateReqQosPolicy{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DfsQosPolicyUpdateReqQosPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *DfsQosPolicyUpdateReqQosPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DfsQosPolicyUpdateReqQosPolicy) SetName(v string) {
	o.Name = v
}

// GetReadBandwidth returns the ReadBandwidth field value if set, zero value otherwise.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetReadBandwidth() int64 {
	if o == nil || IsNil(o.ReadBandwidth) {
		var ret int64
		return ret
	}
	return *o.ReadBandwidth
}

// GetReadBandwidthOk returns a tuple with the ReadBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetReadBandwidthOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadBandwidth) {
		return nil, false
	}
	return o.ReadBandwidth, true
}

// HasReadBandwidth returns a boolean if a field has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) HasReadBandwidth() bool {
	if o != nil && !IsNil(o.ReadBandwidth) {
		return true
	}

	return false
}

// SetReadBandwidth gets a reference to the given int64 and assigns it to the ReadBandwidth field.
func (o *DfsQosPolicyUpdateReqQosPolicy) SetReadBandwidth(v int64) {
	o.ReadBandwidth = &v
}

// GetReadObject returns the ReadObject field value if set, zero value otherwise.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetReadObject() int64 {
	if o == nil || IsNil(o.ReadObject) {
		var ret int64
		return ret
	}
	return *o.ReadObject
}

// GetReadObjectOk returns a tuple with the ReadObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetReadObjectOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadObject) {
		return nil, false
	}
	return o.ReadObject, true
}

// HasReadObject returns a boolean if a field has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) HasReadObject() bool {
	if o != nil && !IsNil(o.ReadObject) {
		return true
	}

	return false
}

// SetReadObject gets a reference to the given int64 and assigns it to the ReadObject field.
func (o *DfsQosPolicyUpdateReqQosPolicy) SetReadObject(v int64) {
	o.ReadObject = &v
}

// GetTotalBandwidth returns the TotalBandwidth field value if set, zero value otherwise.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetTotalBandwidth() int64 {
	if o == nil || IsNil(o.TotalBandwidth) {
		var ret int64
		return ret
	}
	return *o.TotalBandwidth
}

// GetTotalBandwidthOk returns a tuple with the TotalBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetTotalBandwidthOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalBandwidth) {
		return nil, false
	}
	return o.TotalBandwidth, true
}

// HasTotalBandwidth returns a boolean if a field has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) HasTotalBandwidth() bool {
	if o != nil && !IsNil(o.TotalBandwidth) {
		return true
	}

	return false
}

// SetTotalBandwidth gets a reference to the given int64 and assigns it to the TotalBandwidth field.
func (o *DfsQosPolicyUpdateReqQosPolicy) SetTotalBandwidth(v int64) {
	o.TotalBandwidth = &v
}

// GetTotalObject returns the TotalObject field value if set, zero value otherwise.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetTotalObject() int64 {
	if o == nil || IsNil(o.TotalObject) {
		var ret int64
		return ret
	}
	return *o.TotalObject
}

// GetTotalObjectOk returns a tuple with the TotalObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetTotalObjectOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalObject) {
		return nil, false
	}
	return o.TotalObject, true
}

// HasTotalObject returns a boolean if a field has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) HasTotalObject() bool {
	if o != nil && !IsNil(o.TotalObject) {
		return true
	}

	return false
}

// SetTotalObject gets a reference to the given int64 and assigns it to the TotalObject field.
func (o *DfsQosPolicyUpdateReqQosPolicy) SetTotalObject(v int64) {
	o.TotalObject = &v
}

// GetWriteBandwidth returns the WriteBandwidth field value if set, zero value otherwise.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetWriteBandwidth() int64 {
	if o == nil || IsNil(o.WriteBandwidth) {
		var ret int64
		return ret
	}
	return *o.WriteBandwidth
}

// GetWriteBandwidthOk returns a tuple with the WriteBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetWriteBandwidthOk() (*int64, bool) {
	if o == nil || IsNil(o.WriteBandwidth) {
		return nil, false
	}
	return o.WriteBandwidth, true
}

// HasWriteBandwidth returns a boolean if a field has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) HasWriteBandwidth() bool {
	if o != nil && !IsNil(o.WriteBandwidth) {
		return true
	}

	return false
}

// SetWriteBandwidth gets a reference to the given int64 and assigns it to the WriteBandwidth field.
func (o *DfsQosPolicyUpdateReqQosPolicy) SetWriteBandwidth(v int64) {
	o.WriteBandwidth = &v
}

// GetWriteObject returns the WriteObject field value if set, zero value otherwise.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetWriteObject() int64 {
	if o == nil || IsNil(o.WriteObject) {
		var ret int64
		return ret
	}
	return *o.WriteObject
}

// GetWriteObjectOk returns a tuple with the WriteObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) GetWriteObjectOk() (*int64, bool) {
	if o == nil || IsNil(o.WriteObject) {
		return nil, false
	}
	return o.WriteObject, true
}

// HasWriteObject returns a boolean if a field has been set.
func (o *DfsQosPolicyUpdateReqQosPolicy) HasWriteObject() bool {
	if o != nil && !IsNil(o.WriteObject) {
		return true
	}

	return false
}

// SetWriteObject gets a reference to the given int64 and assigns it to the WriteObject field.
func (o *DfsQosPolicyUpdateReqQosPolicy) SetWriteObject(v int64) {
	o.WriteObject = &v
}

func (o DfsQosPolicyUpdateReqQosPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsQosPolicyUpdateReqQosPolicy) ToMap() (map[string]interface{}, error) {
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

func (o *DfsQosPolicyUpdateReqQosPolicy) UnmarshalJSON(data []byte) (err error) {
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

	varDfsQosPolicyUpdateReqQosPolicy := _DfsQosPolicyUpdateReqQosPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDfsQosPolicyUpdateReqQosPolicy)

	if err != nil {
		return err
	}

	*o = DfsQosPolicyUpdateReqQosPolicy(varDfsQosPolicyUpdateReqQosPolicy)

	return err
}

type NullableDfsQosPolicyUpdateReqQosPolicy struct {
	value *DfsQosPolicyUpdateReqQosPolicy
	isSet bool
}

func (v NullableDfsQosPolicyUpdateReqQosPolicy) Get() *DfsQosPolicyUpdateReqQosPolicy {
	return v.value
}

func (v *NullableDfsQosPolicyUpdateReqQosPolicy) Set(val *DfsQosPolicyUpdateReqQosPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsQosPolicyUpdateReqQosPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsQosPolicyUpdateReqQosPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsQosPolicyUpdateReqQosPolicy(val *DfsQosPolicyUpdateReqQosPolicy) *NullableDfsQosPolicyUpdateReqQosPolicy {
	return &NullableDfsQosPolicyUpdateReqQosPolicy{value: val, isSet: true}
}

func (v NullableDfsQosPolicyUpdateReqQosPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsQosPolicyUpdateReqQosPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


