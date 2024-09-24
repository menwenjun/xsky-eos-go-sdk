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

// checks if the ClientGroupCreateReqClientGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientGroupCreateReqClientGroup{}

// ClientGroupCreateReqClientGroup struct for ClientGroupCreateReqClientGroup
type ClientGroupCreateReqClientGroup struct {
	Chap *bool `json:"chap,omitempty"`
	Clients []ClientGroupCreateReqClientGroupClientsElt `json:"clients,omitempty"`
	Description *string `json:"description,omitempty"`
	Iname *string `json:"iname,omitempty"`
	Isecret *string `json:"isecret,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Uid *string `json:"uid,omitempty"`
}

// NewClientGroupCreateReqClientGroup instantiates a new ClientGroupCreateReqClientGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientGroupCreateReqClientGroup() *ClientGroupCreateReqClientGroup {
	this := ClientGroupCreateReqClientGroup{}
	return &this
}

// NewClientGroupCreateReqClientGroupWithDefaults instantiates a new ClientGroupCreateReqClientGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientGroupCreateReqClientGroupWithDefaults() *ClientGroupCreateReqClientGroup {
	this := ClientGroupCreateReqClientGroup{}
	return &this
}

// GetChap returns the Chap field value if set, zero value otherwise.
func (o *ClientGroupCreateReqClientGroup) GetChap() bool {
	if o == nil || IsNil(o.Chap) {
		var ret bool
		return ret
	}
	return *o.Chap
}

// GetChapOk returns a tuple with the Chap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroupCreateReqClientGroup) GetChapOk() (*bool, bool) {
	if o == nil || IsNil(o.Chap) {
		return nil, false
	}
	return o.Chap, true
}

// HasChap returns a boolean if a field has been set.
func (o *ClientGroupCreateReqClientGroup) HasChap() bool {
	if o != nil && !IsNil(o.Chap) {
		return true
	}

	return false
}

// SetChap gets a reference to the given bool and assigns it to the Chap field.
func (o *ClientGroupCreateReqClientGroup) SetChap(v bool) {
	o.Chap = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *ClientGroupCreateReqClientGroup) GetClients() []ClientGroupCreateReqClientGroupClientsElt {
	if o == nil || IsNil(o.Clients) {
		var ret []ClientGroupCreateReqClientGroupClientsElt
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroupCreateReqClientGroup) GetClientsOk() ([]ClientGroupCreateReqClientGroupClientsElt, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *ClientGroupCreateReqClientGroup) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []ClientGroupCreateReqClientGroupClientsElt and assigns it to the Clients field.
func (o *ClientGroupCreateReqClientGroup) SetClients(v []ClientGroupCreateReqClientGroupClientsElt) {
	o.Clients = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClientGroupCreateReqClientGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroupCreateReqClientGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClientGroupCreateReqClientGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClientGroupCreateReqClientGroup) SetDescription(v string) {
	o.Description = &v
}

// GetIname returns the Iname field value if set, zero value otherwise.
func (o *ClientGroupCreateReqClientGroup) GetIname() string {
	if o == nil || IsNil(o.Iname) {
		var ret string
		return ret
	}
	return *o.Iname
}

// GetInameOk returns a tuple with the Iname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroupCreateReqClientGroup) GetInameOk() (*string, bool) {
	if o == nil || IsNil(o.Iname) {
		return nil, false
	}
	return o.Iname, true
}

// HasIname returns a boolean if a field has been set.
func (o *ClientGroupCreateReqClientGroup) HasIname() bool {
	if o != nil && !IsNil(o.Iname) {
		return true
	}

	return false
}

// SetIname gets a reference to the given string and assigns it to the Iname field.
func (o *ClientGroupCreateReqClientGroup) SetIname(v string) {
	o.Iname = &v
}

// GetIsecret returns the Isecret field value if set, zero value otherwise.
func (o *ClientGroupCreateReqClientGroup) GetIsecret() string {
	if o == nil || IsNil(o.Isecret) {
		var ret string
		return ret
	}
	return *o.Isecret
}

// GetIsecretOk returns a tuple with the Isecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroupCreateReqClientGroup) GetIsecretOk() (*string, bool) {
	if o == nil || IsNil(o.Isecret) {
		return nil, false
	}
	return o.Isecret, true
}

// HasIsecret returns a boolean if a field has been set.
func (o *ClientGroupCreateReqClientGroup) HasIsecret() bool {
	if o != nil && !IsNil(o.Isecret) {
		return true
	}

	return false
}

// SetIsecret gets a reference to the given string and assigns it to the Isecret field.
func (o *ClientGroupCreateReqClientGroup) SetIsecret(v string) {
	o.Isecret = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClientGroupCreateReqClientGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroupCreateReqClientGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClientGroupCreateReqClientGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClientGroupCreateReqClientGroup) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClientGroupCreateReqClientGroup) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroupCreateReqClientGroup) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClientGroupCreateReqClientGroup) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClientGroupCreateReqClientGroup) SetType(v string) {
	o.Type = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ClientGroupCreateReqClientGroup) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGroupCreateReqClientGroup) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ClientGroupCreateReqClientGroup) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *ClientGroupCreateReqClientGroup) SetUid(v string) {
	o.Uid = &v
}

func (o ClientGroupCreateReqClientGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientGroupCreateReqClientGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Chap) {
		toSerialize["chap"] = o.Chap
	}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Iname) {
		toSerialize["iname"] = o.Iname
	}
	if !IsNil(o.Isecret) {
		toSerialize["isecret"] = o.Isecret
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableClientGroupCreateReqClientGroup struct {
	value *ClientGroupCreateReqClientGroup
	isSet bool
}

func (v NullableClientGroupCreateReqClientGroup) Get() *ClientGroupCreateReqClientGroup {
	return v.value
}

func (v *NullableClientGroupCreateReqClientGroup) Set(val *ClientGroupCreateReqClientGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableClientGroupCreateReqClientGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableClientGroupCreateReqClientGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientGroupCreateReqClientGroup(val *ClientGroupCreateReqClientGroup) *NullableClientGroupCreateReqClientGroup {
	return &NullableClientGroupCreateReqClientGroup{value: val, isSet: true}
}

func (v NullableClientGroupCreateReqClientGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientGroupCreateReqClientGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


