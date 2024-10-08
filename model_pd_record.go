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

// checks if the PdRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PdRecord{}

// PdRecord PdRecord defines the response content of protection domain related API
type PdRecord struct {
	// time of creating protection domain
	Create *time.Time `json:"create,omitempty"`
	// description of protection domain
	Description *string `json:"description,omitempty"`
	// id of protection domain
	Id *int64 `json:"id,omitempty"`
	// name of protection domain
	Name *string `json:"name,omitempty"`
	PlacementNode *PlacementNodeNestview `json:"placement_node,omitempty"`
	// status protection domain
	Status *string `json:"status,omitempty"`
	// time of updating protection domain
	Update *time.Time `json:"update,omitempty"`
}

// NewPdRecord instantiates a new PdRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPdRecord() *PdRecord {
	this := PdRecord{}
	return &this
}

// NewPdRecordWithDefaults instantiates a new PdRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPdRecordWithDefaults() *PdRecord {
	this := PdRecord{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *PdRecord) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdRecord) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *PdRecord) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *PdRecord) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PdRecord) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdRecord) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PdRecord) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PdRecord) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PdRecord) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdRecord) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PdRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PdRecord) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PdRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PdRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PdRecord) SetName(v string) {
	o.Name = &v
}

// GetPlacementNode returns the PlacementNode field value if set, zero value otherwise.
func (o *PdRecord) GetPlacementNode() PlacementNodeNestview {
	if o == nil || IsNil(o.PlacementNode) {
		var ret PlacementNodeNestview
		return ret
	}
	return *o.PlacementNode
}

// GetPlacementNodeOk returns a tuple with the PlacementNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdRecord) GetPlacementNodeOk() (*PlacementNodeNestview, bool) {
	if o == nil || IsNil(o.PlacementNode) {
		return nil, false
	}
	return o.PlacementNode, true
}

// HasPlacementNode returns a boolean if a field has been set.
func (o *PdRecord) HasPlacementNode() bool {
	if o != nil && !IsNil(o.PlacementNode) {
		return true
	}

	return false
}

// SetPlacementNode gets a reference to the given PlacementNodeNestview and assigns it to the PlacementNode field.
func (o *PdRecord) SetPlacementNode(v PlacementNodeNestview) {
	o.PlacementNode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PdRecord) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdRecord) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PdRecord) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PdRecord) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *PdRecord) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdRecord) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *PdRecord) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *PdRecord) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o PdRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PdRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
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
	if !IsNil(o.PlacementNode) {
		toSerialize["placement_node"] = o.PlacementNode
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullablePdRecord struct {
	value *PdRecord
	isSet bool
}

func (v NullablePdRecord) Get() *PdRecord {
	return v.value
}

func (v *NullablePdRecord) Set(val *PdRecord) {
	v.value = val
	v.isSet = true
}

func (v NullablePdRecord) IsSet() bool {
	return v.isSet
}

func (v *NullablePdRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePdRecord(val *PdRecord) *NullablePdRecord {
	return &NullablePdRecord{value: val, isSet: true}
}

func (v NullablePdRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePdRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


