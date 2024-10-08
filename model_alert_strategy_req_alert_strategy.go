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

// checks if the AlertStrategyReqAlertStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertStrategyReqAlertStrategy{}

// AlertStrategyReqAlertStrategy struct for AlertStrategyReqAlertStrategy
type AlertStrategyReqAlertStrategy struct {
	AlertContactIds []int64 `json:"alert_contact_ids,omitempty"`
	AlertRuleIds []int64 `json:"alert_rule_ids,omitempty"`
	Description *string `json:"description,omitempty"`
	Enable *bool `json:"enable,omitempty"`
	Name *string `json:"name,omitempty"`
	NotifyMethods *string `json:"notify_methods,omitempty"`
}

// NewAlertStrategyReqAlertStrategy instantiates a new AlertStrategyReqAlertStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertStrategyReqAlertStrategy() *AlertStrategyReqAlertStrategy {
	this := AlertStrategyReqAlertStrategy{}
	return &this
}

// NewAlertStrategyReqAlertStrategyWithDefaults instantiates a new AlertStrategyReqAlertStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertStrategyReqAlertStrategyWithDefaults() *AlertStrategyReqAlertStrategy {
	this := AlertStrategyReqAlertStrategy{}
	return &this
}

// GetAlertContactIds returns the AlertContactIds field value if set, zero value otherwise.
func (o *AlertStrategyReqAlertStrategy) GetAlertContactIds() []int64 {
	if o == nil || IsNil(o.AlertContactIds) {
		var ret []int64
		return ret
	}
	return o.AlertContactIds
}

// GetAlertContactIdsOk returns a tuple with the AlertContactIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategyReqAlertStrategy) GetAlertContactIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AlertContactIds) {
		return nil, false
	}
	return o.AlertContactIds, true
}

// HasAlertContactIds returns a boolean if a field has been set.
func (o *AlertStrategyReqAlertStrategy) HasAlertContactIds() bool {
	if o != nil && !IsNil(o.AlertContactIds) {
		return true
	}

	return false
}

// SetAlertContactIds gets a reference to the given []int64 and assigns it to the AlertContactIds field.
func (o *AlertStrategyReqAlertStrategy) SetAlertContactIds(v []int64) {
	o.AlertContactIds = v
}

// GetAlertRuleIds returns the AlertRuleIds field value if set, zero value otherwise.
func (o *AlertStrategyReqAlertStrategy) GetAlertRuleIds() []int64 {
	if o == nil || IsNil(o.AlertRuleIds) {
		var ret []int64
		return ret
	}
	return o.AlertRuleIds
}

// GetAlertRuleIdsOk returns a tuple with the AlertRuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategyReqAlertStrategy) GetAlertRuleIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AlertRuleIds) {
		return nil, false
	}
	return o.AlertRuleIds, true
}

// HasAlertRuleIds returns a boolean if a field has been set.
func (o *AlertStrategyReqAlertStrategy) HasAlertRuleIds() bool {
	if o != nil && !IsNil(o.AlertRuleIds) {
		return true
	}

	return false
}

// SetAlertRuleIds gets a reference to the given []int64 and assigns it to the AlertRuleIds field.
func (o *AlertStrategyReqAlertStrategy) SetAlertRuleIds(v []int64) {
	o.AlertRuleIds = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertStrategyReqAlertStrategy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategyReqAlertStrategy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertStrategyReqAlertStrategy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertStrategyReqAlertStrategy) SetDescription(v string) {
	o.Description = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *AlertStrategyReqAlertStrategy) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategyReqAlertStrategy) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *AlertStrategyReqAlertStrategy) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *AlertStrategyReqAlertStrategy) SetEnable(v bool) {
	o.Enable = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertStrategyReqAlertStrategy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategyReqAlertStrategy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertStrategyReqAlertStrategy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertStrategyReqAlertStrategy) SetName(v string) {
	o.Name = &v
}

// GetNotifyMethods returns the NotifyMethods field value if set, zero value otherwise.
func (o *AlertStrategyReqAlertStrategy) GetNotifyMethods() string {
	if o == nil || IsNil(o.NotifyMethods) {
		var ret string
		return ret
	}
	return *o.NotifyMethods
}

// GetNotifyMethodsOk returns a tuple with the NotifyMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertStrategyReqAlertStrategy) GetNotifyMethodsOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyMethods) {
		return nil, false
	}
	return o.NotifyMethods, true
}

// HasNotifyMethods returns a boolean if a field has been set.
func (o *AlertStrategyReqAlertStrategy) HasNotifyMethods() bool {
	if o != nil && !IsNil(o.NotifyMethods) {
		return true
	}

	return false
}

// SetNotifyMethods gets a reference to the given string and assigns it to the NotifyMethods field.
func (o *AlertStrategyReqAlertStrategy) SetNotifyMethods(v string) {
	o.NotifyMethods = &v
}

func (o AlertStrategyReqAlertStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertStrategyReqAlertStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertContactIds) {
		toSerialize["alert_contact_ids"] = o.AlertContactIds
	}
	if !IsNil(o.AlertRuleIds) {
		toSerialize["alert_rule_ids"] = o.AlertRuleIds
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NotifyMethods) {
		toSerialize["notify_methods"] = o.NotifyMethods
	}
	return toSerialize, nil
}

type NullableAlertStrategyReqAlertStrategy struct {
	value *AlertStrategyReqAlertStrategy
	isSet bool
}

func (v NullableAlertStrategyReqAlertStrategy) Get() *AlertStrategyReqAlertStrategy {
	return v.value
}

func (v *NullableAlertStrategyReqAlertStrategy) Set(val *AlertStrategyReqAlertStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertStrategyReqAlertStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertStrategyReqAlertStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertStrategyReqAlertStrategy(val *AlertStrategyReqAlertStrategy) *NullableAlertStrategyReqAlertStrategy {
	return &NullableAlertStrategyReqAlertStrategy{value: val, isSet: true}
}

func (v NullableAlertStrategyReqAlertStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertStrategyReqAlertStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


