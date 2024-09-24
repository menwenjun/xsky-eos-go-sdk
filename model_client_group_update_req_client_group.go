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

// checks if the ClientGroupUpdateReqClientGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientGroupUpdateReqClientGroup{}

// ClientGroupUpdateReqClientGroup struct for ClientGroupUpdateReqClientGroup
type ClientGroupUpdateReqClientGroup struct {
	Clients []ClientGroupUpdateReqClientGroupClientsElt `json:"clients,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewClientGroupUpdateReqClientGroup instantiates a new ClientGroupUpdateReqClientGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientGroupUpdateReqClientGroup() *ClientGroupUpdateReqClientGroup {
	this := ClientGroupUpdateReqClientGroup{}
	return &this
}

// NewClientGroupUpdateReqClientGroupWithDefaults instantiates a new ClientGroupUpdateReqClientGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientGroupUpdateReqClientGroupWithDefaults() *ClientGroupUpdateReqClientGroup {
	this := ClientGroupUpdateReqClientGroup{}
	return &this
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *ClientGroupUpdateReqClientGroup) GetClients() []ClientGroupUpdateReqClientGroupClientsElt {
	if o == nil || IsNil(o.Clients) {
		var ret []ClientGroupUpdateReqClientGroupClientsElt
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroupUpdateReqClientGroup) GetClientsOk() ([]ClientGroupUpdateReqClientGroupClientsElt, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *ClientGroupUpdateReqClientGroup) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []ClientGroupUpdateReqClientGroupClientsElt and assigns it to the Clients field.
func (o *ClientGroupUpdateReqClientGroup) SetClients(v []ClientGroupUpdateReqClientGroupClientsElt) {
	o.Clients = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClientGroupUpdateReqClientGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroupUpdateReqClientGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClientGroupUpdateReqClientGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClientGroupUpdateReqClientGroup) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClientGroupUpdateReqClientGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroupUpdateReqClientGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClientGroupUpdateReqClientGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClientGroupUpdateReqClientGroup) SetName(v string) {
	o.Name = &v
}

func (o ClientGroupUpdateReqClientGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientGroupUpdateReqClientGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableClientGroupUpdateReqClientGroup struct {
	value *ClientGroupUpdateReqClientGroup
	isSet bool
}

func (v NullableClientGroupUpdateReqClientGroup) Get() *ClientGroupUpdateReqClientGroup {
	return v.value
}

func (v *NullableClientGroupUpdateReqClientGroup) Set(val *ClientGroupUpdateReqClientGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableClientGroupUpdateReqClientGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableClientGroupUpdateReqClientGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientGroupUpdateReqClientGroup(val *ClientGroupUpdateReqClientGroup) *NullableClientGroupUpdateReqClientGroup {
	return &NullableClientGroupUpdateReqClientGroup{value: val, isSet: true}
}

func (v NullableClientGroupUpdateReqClientGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientGroupUpdateReqClientGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


