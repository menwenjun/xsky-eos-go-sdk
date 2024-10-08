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

// checks if the FSActiveDirectoryNestview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FSActiveDirectoryNestview{}

// FSActiveDirectoryNestview struct for FSActiveDirectoryNestview
type FSActiveDirectoryNestview struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Realm *string `json:"realm,omitempty"`
}

// NewFSActiveDirectoryNestview instantiates a new FSActiveDirectoryNestview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFSActiveDirectoryNestview() *FSActiveDirectoryNestview {
	this := FSActiveDirectoryNestview{}
	return &this
}

// NewFSActiveDirectoryNestviewWithDefaults instantiates a new FSActiveDirectoryNestview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFSActiveDirectoryNestviewWithDefaults() *FSActiveDirectoryNestview {
	this := FSActiveDirectoryNestview{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FSActiveDirectoryNestview) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSActiveDirectoryNestview) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FSActiveDirectoryNestview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FSActiveDirectoryNestview) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FSActiveDirectoryNestview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSActiveDirectoryNestview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FSActiveDirectoryNestview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FSActiveDirectoryNestview) SetName(v string) {
	o.Name = &v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *FSActiveDirectoryNestview) GetRealm() string {
	if o == nil || IsNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSActiveDirectoryNestview) GetRealmOk() (*string, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *FSActiveDirectoryNestview) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *FSActiveDirectoryNestview) SetRealm(v string) {
	o.Realm = &v
}

func (o FSActiveDirectoryNestview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FSActiveDirectoryNestview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableFSActiveDirectoryNestview struct {
	value *FSActiveDirectoryNestview
	isSet bool
}

func (v NullableFSActiveDirectoryNestview) Get() *FSActiveDirectoryNestview {
	return v.value
}

func (v *NullableFSActiveDirectoryNestview) Set(val *FSActiveDirectoryNestview) {
	v.value = val
	v.isSet = true
}

func (v NullableFSActiveDirectoryNestview) IsSet() bool {
	return v.isSet
}

func (v *NullableFSActiveDirectoryNestview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFSActiveDirectoryNestview(val *FSActiveDirectoryNestview) *NullableFSActiveDirectoryNestview {
	return &NullableFSActiveDirectoryNestview{value: val, isSet: true}
}

func (v NullableFSActiveDirectoryNestview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFSActiveDirectoryNestview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


