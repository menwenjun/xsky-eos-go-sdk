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

// checks if the PlacementNodeNestviewParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlacementNodeNestviewParent{}

// PlacementNodeNestviewParent parent of placement node
type PlacementNodeNestviewParent struct {
	Cluster *Cluster `json:"cluster,omitempty"`
	// time of creating placement node
	Create *time.Time `json:"create,omitempty"`
	// id of placement node
	Id *int64 `json:"id,omitempty"`
	// name of placement node
	Name *string `json:"name,omitempty"`
	// old id of placement node
	OrigId *int64 `json:"orig_id,omitempty"`
	Parent *PlacementNode `json:"parent,omitempty"`
	Properties *PlacementNodeProperties `json:"properties,omitempty"`
	// type of placement node
	Type *string `json:"type,omitempty"`
	// time of updating placement node
	Update *time.Time `json:"update,omitempty"`
}

// NewPlacementNodeNestviewParent instantiates a new PlacementNodeNestviewParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlacementNodeNestviewParent() *PlacementNodeNestviewParent {
	this := PlacementNodeNestviewParent{}
	return &this
}

// NewPlacementNodeNestviewParentWithDefaults instantiates a new PlacementNodeNestviewParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlacementNodeNestviewParentWithDefaults() *PlacementNodeNestviewParent {
	this := PlacementNodeNestviewParent{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *PlacementNodeNestviewParent) GetCluster() Cluster {
	if o == nil || IsNil(o.Cluster) {
		var ret Cluster
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementNodeNestviewParent) GetClusterOk() (*Cluster, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *PlacementNodeNestviewParent) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given Cluster and assigns it to the Cluster field.
func (o *PlacementNodeNestviewParent) SetCluster(v Cluster) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *PlacementNodeNestviewParent) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementNodeNestviewParent) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *PlacementNodeNestviewParent) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *PlacementNodeNestviewParent) SetCreate(v time.Time) {
	o.Create = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlacementNodeNestviewParent) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementNodeNestviewParent) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlacementNodeNestviewParent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PlacementNodeNestviewParent) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlacementNodeNestviewParent) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementNodeNestviewParent) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlacementNodeNestviewParent) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlacementNodeNestviewParent) SetName(v string) {
	o.Name = &v
}

// GetOrigId returns the OrigId field value if set, zero value otherwise.
func (o *PlacementNodeNestviewParent) GetOrigId() int64 {
	if o == nil || IsNil(o.OrigId) {
		var ret int64
		return ret
	}
	return *o.OrigId
}

// GetOrigIdOk returns a tuple with the OrigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementNodeNestviewParent) GetOrigIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrigId) {
		return nil, false
	}
	return o.OrigId, true
}

// HasOrigId returns a boolean if a field has been set.
func (o *PlacementNodeNestviewParent) HasOrigId() bool {
	if o != nil && !IsNil(o.OrigId) {
		return true
	}

	return false
}

// SetOrigId gets a reference to the given int64 and assigns it to the OrigId field.
func (o *PlacementNodeNestviewParent) SetOrigId(v int64) {
	o.OrigId = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *PlacementNodeNestviewParent) GetParent() PlacementNode {
	if o == nil || IsNil(o.Parent) {
		var ret PlacementNode
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementNodeNestviewParent) GetParentOk() (*PlacementNode, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *PlacementNodeNestviewParent) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given PlacementNode and assigns it to the Parent field.
func (o *PlacementNodeNestviewParent) SetParent(v PlacementNode) {
	o.Parent = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *PlacementNodeNestviewParent) GetProperties() PlacementNodeProperties {
	if o == nil || IsNil(o.Properties) {
		var ret PlacementNodeProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementNodeNestviewParent) GetPropertiesOk() (*PlacementNodeProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *PlacementNodeNestviewParent) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given PlacementNodeProperties and assigns it to the Properties field.
func (o *PlacementNodeNestviewParent) SetProperties(v PlacementNodeProperties) {
	o.Properties = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlacementNodeNestviewParent) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementNodeNestviewParent) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlacementNodeNestviewParent) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlacementNodeNestviewParent) SetType(v string) {
	o.Type = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *PlacementNodeNestviewParent) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlacementNodeNestviewParent) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *PlacementNodeNestviewParent) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *PlacementNodeNestviewParent) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o PlacementNodeNestviewParent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlacementNodeNestviewParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrigId) {
		toSerialize["orig_id"] = o.OrigId
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullablePlacementNodeNestviewParent struct {
	value *PlacementNodeNestviewParent
	isSet bool
}

func (v NullablePlacementNodeNestviewParent) Get() *PlacementNodeNestviewParent {
	return v.value
}

func (v *NullablePlacementNodeNestviewParent) Set(val *PlacementNodeNestviewParent) {
	v.value = val
	v.isSet = true
}

func (v NullablePlacementNodeNestviewParent) IsSet() bool {
	return v.isSet
}

func (v *NullablePlacementNodeNestviewParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlacementNodeNestviewParent(val *PlacementNodeNestviewParent) *NullablePlacementNodeNestviewParent {
	return &NullablePlacementNodeNestviewParent{value: val, isSet: true}
}

func (v NullablePlacementNodeNestviewParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlacementNodeNestviewParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


