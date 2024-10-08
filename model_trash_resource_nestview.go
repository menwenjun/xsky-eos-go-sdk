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

// checks if the TrashResourceNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrashResourceNestview{}

// TrashResourceNestview struct for TrashResourceNestview
type TrashResourceNestview struct {
	Create *time.Time `json:"create,omitempty"`
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	Id *int64 `json:"id,omitempty"`
	ResourceId *int64 `json:"resource_id,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
}

// NewTrashResourceNestview instantiates a new TrashResourceNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrashResourceNestview() *TrashResourceNestview {
	this := TrashResourceNestview{}
	return &this
}

// NewTrashResourceNestviewWithDefaults instantiates a new TrashResourceNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrashResourceNestviewWithDefaults() *TrashResourceNestview {
	this := TrashResourceNestview{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *TrashResourceNestview) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrashResourceNestview) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *TrashResourceNestview) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *TrashResourceNestview) SetCreate(v time.Time) {
	o.Create = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *TrashResourceNestview) GetExpireTime() time.Time {
	if o == nil || IsNil(o.ExpireTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrashResourceNestview) GetExpireTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *TrashResourceNestview) HasExpireTime() bool {
	if o != nil && !IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given time.Time and assigns it to the ExpireTime field.
func (o *TrashResourceNestview) SetExpireTime(v time.Time) {
	o.ExpireTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrashResourceNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrashResourceNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrashResourceNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TrashResourceNestview) SetId(v int64) {
	o.Id = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *TrashResourceNestview) GetResourceId() int64 {
	if o == nil || IsNil(o.ResourceId) {
		var ret int64
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrashResourceNestview) GetResourceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *TrashResourceNestview) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given int64 and assigns it to the ResourceId field.
func (o *TrashResourceNestview) SetResourceId(v int64) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *TrashResourceNestview) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrashResourceNestview) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *TrashResourceNestview) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *TrashResourceNestview) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o TrashResourceNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrashResourceNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.ExpireTime) {
		toSerialize["expire_time"] = o.ExpireTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resource_id"] = o.ResourceId
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resource_type"] = o.ResourceType
	}
	return toSerialize, nil
}

type NullableTrashResourceNestview struct {
	value *TrashResourceNestview
	isSet bool
}

func (v NullableTrashResourceNestview) Get() *TrashResourceNestview {
	return v.value
}

func (v *NullableTrashResourceNestview) Set(val *TrashResourceNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableTrashResourceNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableTrashResourceNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrashResourceNestview(val *TrashResourceNestview) *NullableTrashResourceNestview {
	return &NullableTrashResourceNestview{value: val, isSet: true}
}

func (v NullableTrashResourceNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrashResourceNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


