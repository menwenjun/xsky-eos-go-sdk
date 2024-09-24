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

// checks if the FSKerberosNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FSKerberosNestview{}

// FSKerberosNestview struct for FSKerberosNestview
type FSKerberosNestview struct {
	ActionStatus *string `json:"action_status,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Realm *string `json:"realm,omitempty"`
}

// NewFSKerberosNestview instantiates a new FSKerberosNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFSKerberosNestview() *FSKerberosNestview {
	this := FSKerberosNestview{}
	return &this
}

// NewFSKerberosNestviewWithDefaults instantiates a new FSKerberosNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFSKerberosNestviewWithDefaults() *FSKerberosNestview {
	this := FSKerberosNestview{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *FSKerberosNestview) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSKerberosNestview) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *FSKerberosNestview) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *FSKerberosNestview) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FSKerberosNestview) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSKerberosNestview) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FSKerberosNestview) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FSKerberosNestview) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FSKerberosNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSKerberosNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FSKerberosNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FSKerberosNestview) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FSKerberosNestview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSKerberosNestview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FSKerberosNestview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FSKerberosNestview) SetName(v string) {
	o.Name = &v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *FSKerberosNestview) GetRealm() string {
	if o == nil || IsNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSKerberosNestview) GetRealmOk() (*string, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *FSKerberosNestview) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *FSKerberosNestview) SetRealm(v string) {
	o.Realm = &v
}

func (o FSKerberosNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FSKerberosNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	return toSerialize, nil
}

type NullableFSKerberosNestview struct {
	value *FSKerberosNestview
	isSet bool
}

func (v NullableFSKerberosNestview) Get() *FSKerberosNestview {
	return v.value
}

func (v *NullableFSKerberosNestview) Set(val *FSKerberosNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableFSKerberosNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableFSKerberosNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFSKerberosNestview(val *FSKerberosNestview) *NullableFSKerberosNestview {
	return &NullableFSKerberosNestview{value: val, isSet: true}
}

func (v NullableFSKerberosNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFSKerberosNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


