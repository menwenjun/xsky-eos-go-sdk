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

// checks if the EmailGroupRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailGroupRecord{}

// EmailGroupRecord EmailGroupRecord defines email group record
type EmailGroupRecord struct {
	Create *time.Time `json:"create,omitempty"`
	Emails []string `json:"emails,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewEmailGroupRecord instantiates a new EmailGroupRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailGroupRecord() *EmailGroupRecord {
	this := EmailGroupRecord{}
	return &this
}

// NewEmailGroupRecordWithDefaults instantiates a new EmailGroupRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailGroupRecordWithDefaults() *EmailGroupRecord {
	this := EmailGroupRecord{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *EmailGroupRecord) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailGroupRecord) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *EmailGroupRecord) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *EmailGroupRecord) SetCreate(v time.Time) {
	o.Create = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *EmailGroupRecord) GetEmails() []string {
	if o == nil || IsNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailGroupRecord) GetEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *EmailGroupRecord) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *EmailGroupRecord) SetEmails(v []string) {
	o.Emails = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailGroupRecord) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailGroupRecord) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailGroupRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *EmailGroupRecord) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmailGroupRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailGroupRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmailGroupRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EmailGroupRecord) SetName(v string) {
	o.Name = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *EmailGroupRecord) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailGroupRecord) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *EmailGroupRecord) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *EmailGroupRecord) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o EmailGroupRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailGroupRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableEmailGroupRecord struct {
	value *EmailGroupRecord
	isSet bool
}

func (v NullableEmailGroupRecord) Get() *EmailGroupRecord {
	return v.value
}

func (v *NullableEmailGroupRecord) Set(val *EmailGroupRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailGroupRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailGroupRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailGroupRecord(val *EmailGroupRecord) *NullableEmailGroupRecord {
	return &NullableEmailGroupRecord{value: val, isSet: true}
}

func (v NullableEmailGroupRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailGroupRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


