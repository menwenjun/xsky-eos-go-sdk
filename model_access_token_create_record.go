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

// checks if the AccessTokenCreateRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenCreateRecord{}

// AccessTokenCreateRecord AccessTokenCreateRecord defines access token record for creating api response
type AccessTokenCreateRecord struct {
	Create *time.Time `json:"create,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Used *bool `json:"used,omitempty"`
	User *UserNestview `json:"user,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewAccessTokenCreateRecord instantiates a new AccessTokenCreateRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenCreateRecord() *AccessTokenCreateRecord {
	this := AccessTokenCreateRecord{}
	return &this
}

// NewAccessTokenCreateRecordWithDefaults instantiates a new AccessTokenCreateRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenCreateRecordWithDefaults() *AccessTokenCreateRecord {
	this := AccessTokenCreateRecord{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *AccessTokenCreateRecord) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenCreateRecord) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *AccessTokenCreateRecord) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *AccessTokenCreateRecord) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccessTokenCreateRecord) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenCreateRecord) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessTokenCreateRecord) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccessTokenCreateRecord) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessTokenCreateRecord) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenCreateRecord) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessTokenCreateRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AccessTokenCreateRecord) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessTokenCreateRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenCreateRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessTokenCreateRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessTokenCreateRecord) SetName(v string) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *AccessTokenCreateRecord) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenCreateRecord) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *AccessTokenCreateRecord) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *AccessTokenCreateRecord) SetRoles(v []string) {
	o.Roles = v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *AccessTokenCreateRecord) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenCreateRecord) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *AccessTokenCreateRecord) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *AccessTokenCreateRecord) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *AccessTokenCreateRecord) GetUsed() bool {
	if o == nil || IsNil(o.Used) {
		var ret bool
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenCreateRecord) GetUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *AccessTokenCreateRecord) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given bool and assigns it to the Used field.
func (o *AccessTokenCreateRecord) SetUsed(v bool) {
	o.Used = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AccessTokenCreateRecord) GetUser() UserNestview {
	if o == nil || IsNil(o.User) {
		var ret UserNestview
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenCreateRecord) GetUserOk() (*UserNestview, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AccessTokenCreateRecord) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserNestview and assigns it to the User field.
func (o *AccessTokenCreateRecord) SetUser(v UserNestview) {
	o.User = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *AccessTokenCreateRecord) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenCreateRecord) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *AccessTokenCreateRecord) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *AccessTokenCreateRecord) SetUuid(v string) {
	o.Uuid = &v
}

func (o AccessTokenCreateRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenCreateRecord) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableAccessTokenCreateRecord struct {
	value *AccessTokenCreateRecord
	isSet bool
}

func (v NullableAccessTokenCreateRecord) Get() *AccessTokenCreateRecord {
	return v.value
}

func (v *NullableAccessTokenCreateRecord) Set(val *AccessTokenCreateRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenCreateRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenCreateRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenCreateRecord(val *AccessTokenCreateRecord) *NullableAccessTokenCreateRecord {
	return &NullableAccessTokenCreateRecord{value: val, isSet: true}
}

func (v NullableAccessTokenCreateRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenCreateRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


