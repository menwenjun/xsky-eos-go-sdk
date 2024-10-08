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

// checks if the AlertStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertStrategy{}

// AlertStrategy AlertStrategy defines strategy for alert
type AlertStrategy struct {
	ActionStatus *string `json:"action_status,omitempty"`
	AlertContacts []AlertContactNestview `json:"alert_contacts,omitempty"`
	AlertRules []AlertRuleNestview `json:"alert_rules,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Description *string `json:"description,omitempty"`
	Enable *bool `json:"enable,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	NotifyMethods *string `json:"notify_methods,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewAlertStrategy instantiates a new AlertStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertStrategy() *AlertStrategy {
	this := AlertStrategy{}
	return &this
}

// NewAlertStrategyWithDefaults instantiates a new AlertStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertStrategyWithDefaults() *AlertStrategy {
	this := AlertStrategy{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *AlertStrategy) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *AlertStrategy) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *AlertStrategy) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetAlertContacts returns the AlertContacts field value if set, zero value otherwise.
func (o *AlertStrategy) GetAlertContacts() []AlertContactNestview {
	if o == nil || IsNil(o.AlertContacts) {
		var ret []AlertContactNestview
		return ret
	}
	return o.AlertContacts
}

// GetAlertContactsOk returns a tuple with the AlertContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetAlertContactsOk() ([]AlertContactNestview, bool) {
	if o == nil || IsNil(o.AlertContacts) {
		return nil, false
	}
	return o.AlertContacts, true
}

// HasAlertContacts returns a boolean if a field has been set.
func (o *AlertStrategy) HasAlertContacts() bool {
	if o != nil && !IsNil(o.AlertContacts) {
		return true
	}

	return false
}

// SetAlertContacts gets a reference to the given []AlertContactNestview and assigns it to the AlertContacts field.
func (o *AlertStrategy) SetAlertContacts(v []AlertContactNestview) {
	o.AlertContacts = v
}

// GetAlertRules returns the AlertRules field value if set, zero value otherwise.
func (o *AlertStrategy) GetAlertRules() []AlertRuleNestview {
	if o == nil || IsNil(o.AlertRules) {
		var ret []AlertRuleNestview
		return ret
	}
	return o.AlertRules
}

// GetAlertRulesOk returns a tuple with the AlertRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetAlertRulesOk() ([]AlertRuleNestview, bool) {
	if o == nil || IsNil(o.AlertRules) {
		return nil, false
	}
	return o.AlertRules, true
}

// HasAlertRules returns a boolean if a field has been set.
func (o *AlertStrategy) HasAlertRules() bool {
	if o != nil && !IsNil(o.AlertRules) {
		return true
	}

	return false
}

// SetAlertRules gets a reference to the given []AlertRuleNestview and assigns it to the AlertRules field.
func (o *AlertStrategy) SetAlertRules(v []AlertRuleNestview) {
	o.AlertRules = v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *AlertStrategy) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *AlertStrategy) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *AlertStrategy) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertStrategy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertStrategy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertStrategy) SetDescription(v string) {
	o.Description = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *AlertStrategy) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *AlertStrategy) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *AlertStrategy) SetEnable(v bool) {
	o.Enable = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertStrategy) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertStrategy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AlertStrategy) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertStrategy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertStrategy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertStrategy) SetName(v string) {
	o.Name = &v
}

// GetNotifyMethods returns the NotifyMethods field value if set, zero value otherwise.
func (o *AlertStrategy) GetNotifyMethods() string {
	if o == nil || IsNil(o.NotifyMethods) {
		var ret string
		return ret
	}
	return *o.NotifyMethods
}

// GetNotifyMethodsOk returns a tuple with the NotifyMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetNotifyMethodsOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyMethods) {
		return nil, false
	}
	return o.NotifyMethods, true
}

// HasNotifyMethods returns a boolean if a field has been set.
func (o *AlertStrategy) HasNotifyMethods() bool {
	if o != nil && !IsNil(o.NotifyMethods) {
		return true
	}

	return false
}

// SetNotifyMethods gets a reference to the given string and assigns it to the NotifyMethods field.
func (o *AlertStrategy) SetNotifyMethods(v string) {
	o.NotifyMethods = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlertStrategy) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlertStrategy) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlertStrategy) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *AlertStrategy) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategy) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *AlertStrategy) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *AlertStrategy) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o AlertStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.AlertContacts) {
		toSerialize["alert_contacts"] = o.AlertContacts
	}
	if !IsNil(o.AlertRules) {
		toSerialize["alert_rules"] = o.AlertRules
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NotifyMethods) {
		toSerialize["notify_methods"] = o.NotifyMethods
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableAlertStrategy struct {
	value *AlertStrategy
	isSet bool
}

func (v NullableAlertStrategy) Get() *AlertStrategy {
	return v.value
}

func (v *NullableAlertStrategy) Set(val *AlertStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertStrategy(val *AlertStrategy) *NullableAlertStrategy {
	return &NullableAlertStrategy{value: val, isSet: true}
}

func (v NullableAlertStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


