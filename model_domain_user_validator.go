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

// checks if the DomainUserValidator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainUserValidator{}

// DomainUserValidator DomainUserValidator defines the validator for domain users
type DomainUserValidator struct {
	Create *time.Time `json:"create,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Members map[string]interface{} `json:"members,omitempty"`
	Report map[string]interface{} `json:"report,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewDomainUserValidator instantiates a new DomainUserValidator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainUserValidator() *DomainUserValidator {
	this := DomainUserValidator{}
	return &this
}

// NewDomainUserValidatorWithDefaults instantiates a new DomainUserValidator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainUserValidatorWithDefaults() *DomainUserValidator {
	this := DomainUserValidator{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DomainUserValidator) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainUserValidator) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DomainUserValidator) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DomainUserValidator) SetCreate(v time.Time) {
	o.Create = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DomainUserValidator) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainUserValidator) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DomainUserValidator) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DomainUserValidator) SetId(v int64) {
	o.Id = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *DomainUserValidator) GetMembers() map[string]interface{} {
	if o == nil || IsNil(o.Members) {
		var ret map[string]interface{}
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainUserValidator) GetMembersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Members) {
		return map[string]interface{}{}, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *DomainUserValidator) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given map[string]interface{} and assigns it to the Members field.
func (o *DomainUserValidator) SetMembers(v map[string]interface{}) {
	o.Members = v
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *DomainUserValidator) GetReport() map[string]interface{} {
	if o == nil || IsNil(o.Report) {
		var ret map[string]interface{}
		return ret
	}
	return o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainUserValidator) GetReportOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Report) {
		return map[string]interface{}{}, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *DomainUserValidator) HasReport() bool {
	if o != nil && !IsNil(o.Report) {
		return true
	}

	return false
}

// SetReport gets a reference to the given map[string]interface{} and assigns it to the Report field.
func (o *DomainUserValidator) SetReport(v map[string]interface{}) {
	o.Report = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DomainUserValidator) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainUserValidator) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DomainUserValidator) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DomainUserValidator) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DomainUserValidator) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainUserValidator) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DomainUserValidator) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DomainUserValidator) SetType(v string) {
	o.Type = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DomainUserValidator) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainUserValidator) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DomainUserValidator) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DomainUserValidator) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o DomainUserValidator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainUserValidator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.Report) {
		toSerialize["report"] = o.Report
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableDomainUserValidator struct {
	value *DomainUserValidator
	isSet bool
}

func (v NullableDomainUserValidator) Get() *DomainUserValidator {
	return v.value
}

func (v *NullableDomainUserValidator) Set(val *DomainUserValidator) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainUserValidator) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainUserValidator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainUserValidator(val *DomainUserValidator) *NullableDomainUserValidator {
	return &NullableDomainUserValidator{value: val, isSet: true}
}

func (v NullableDomainUserValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainUserValidator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


