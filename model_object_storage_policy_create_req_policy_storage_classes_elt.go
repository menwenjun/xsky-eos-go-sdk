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

// checks if the ObjectStoragePolicyCreateReqPolicyStorageClassesElt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStoragePolicyCreateReqPolicyStorageClassesElt{}

// ObjectStoragePolicyCreateReqPolicyStorageClassesElt struct for ObjectStoragePolicyCreateReqPolicyStorageClassesElt
type ObjectStoragePolicyCreateReqPolicyStorageClassesElt struct {
	ActivePoolIds []int64 `json:"active_pool_ids"`
	Class *int32 `json:"class,omitempty"`
	ClassId *string `json:"class_id,omitempty"`
	Description *string `json:"description,omitempty"`
	InactivePoolIds []int64 `json:"inactive_pool_ids,omitempty"`
	Name string `json:"name"`
}

type _ObjectStoragePolicyCreateReqPolicyStorageClassesElt ObjectStoragePolicyCreateReqPolicyStorageClassesElt

// NewObjectStoragePolicyCreateReqPolicyStorageClassesElt instantiates a new ObjectStoragePolicyCreateReqPolicyStorageClassesElt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStoragePolicyCreateReqPolicyStorageClassesElt(activePoolIds []int64, name string) *ObjectStoragePolicyCreateReqPolicyStorageClassesElt {
	this := ObjectStoragePolicyCreateReqPolicyStorageClassesElt{}
	this.ActivePoolIds = activePoolIds
	this.Name = name
	return &this
}

// NewObjectStoragePolicyCreateReqPolicyStorageClassesEltWithDefaults instantiates a new ObjectStoragePolicyCreateReqPolicyStorageClassesElt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStoragePolicyCreateReqPolicyStorageClassesEltWithDefaults() *ObjectStoragePolicyCreateReqPolicyStorageClassesElt {
	this := ObjectStoragePolicyCreateReqPolicyStorageClassesElt{}
	return &this
}

// GetActivePoolIds returns the ActivePoolIds field value
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetActivePoolIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.ActivePoolIds
}

// GetActivePoolIdsOk returns a tuple with the ActivePoolIds field value
// and a boolean to check if the value has been set.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetActivePoolIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivePoolIds, true
}

// SetActivePoolIds sets field value
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetActivePoolIds(v []int64) {
	o.ActivePoolIds = v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetClass() int32 {
	if o == nil || IsNil(o.Class) {
		var ret int32
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetClassOk() (*int32, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given int32 and assigns it to the Class field.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetClass(v int32) {
	o.Class = &v
}

// GetClassId returns the ClassId field value if set, zero value otherwise.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetClassId() string {
	if o == nil || IsNil(o.ClassId) {
		var ret string
		return ret
	}
	return *o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClassId) {
		return nil, false
	}
	return o.ClassId, true
}

// HasClassId returns a boolean if a field has been set.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) HasClassId() bool {
	if o != nil && !IsNil(o.ClassId) {
		return true
	}

	return false
}

// SetClassId gets a reference to the given string and assigns it to the ClassId field.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetClassId(v string) {
	o.ClassId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetDescription(v string) {
	o.Description = &v
}

// GetInactivePoolIds returns the InactivePoolIds field value if set, zero value otherwise.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetInactivePoolIds() []int64 {
	if o == nil || IsNil(o.InactivePoolIds) {
		var ret []int64
		return ret
	}
	return o.InactivePoolIds
}

// GetInactivePoolIdsOk returns a tuple with the InactivePoolIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetInactivePoolIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.InactivePoolIds) {
		return nil, false
	}
	return o.InactivePoolIds, true
}

// HasInactivePoolIds returns a boolean if a field has been set.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) HasInactivePoolIds() bool {
	if o != nil && !IsNil(o.InactivePoolIds) {
		return true
	}

	return false
}

// SetInactivePoolIds gets a reference to the given []int64 and assigns it to the InactivePoolIds field.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetInactivePoolIds(v []int64) {
	o.InactivePoolIds = v
}

// GetName returns the Name field value
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetName(v string) {
	o.Name = v
}

func (o ObjectStoragePolicyCreateReqPolicyStorageClassesElt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectStoragePolicyCreateReqPolicyStorageClassesElt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active_pool_ids"] = o.ActivePoolIds
	if !IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !IsNil(o.ClassId) {
		toSerialize["class_id"] = o.ClassId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InactivePoolIds) {
		toSerialize["inactive_pool_ids"] = o.InactivePoolIds
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active_pool_ids",
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

	varObjectStoragePolicyCreateReqPolicyStorageClassesElt := _ObjectStoragePolicyCreateReqPolicyStorageClassesElt{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectStoragePolicyCreateReqPolicyStorageClassesElt)

	if err != nil {
		return err
	}

	*o = ObjectStoragePolicyCreateReqPolicyStorageClassesElt(varObjectStoragePolicyCreateReqPolicyStorageClassesElt)

	return err
}

type NullableObjectStoragePolicyCreateReqPolicyStorageClassesElt struct {
	value *ObjectStoragePolicyCreateReqPolicyStorageClassesElt
	isSet bool
}

func (v NullableObjectStoragePolicyCreateReqPolicyStorageClassesElt) Get() *ObjectStoragePolicyCreateReqPolicyStorageClassesElt {
	return v.value
}

func (v *NullableObjectStoragePolicyCreateReqPolicyStorageClassesElt) Set(val *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStoragePolicyCreateReqPolicyStorageClassesElt) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStoragePolicyCreateReqPolicyStorageClassesElt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStoragePolicyCreateReqPolicyStorageClassesElt(val *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) *NullableObjectStoragePolicyCreateReqPolicyStorageClassesElt {
	return &NullableObjectStoragePolicyCreateReqPolicyStorageClassesElt{value: val, isSet: true}
}

func (v NullableObjectStoragePolicyCreateReqPolicyStorageClassesElt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStoragePolicyCreateReqPolicyStorageClassesElt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


