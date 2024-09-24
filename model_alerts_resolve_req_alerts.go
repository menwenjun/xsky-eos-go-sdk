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

// checks if the AlertsResolveReqAlerts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertsResolveReqAlerts{}

// AlertsResolveReqAlerts struct for AlertsResolveReqAlerts
type AlertsResolveReqAlerts struct {
	Group *string `json:"group,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
}

// NewAlertsResolveReqAlerts instantiates a new AlertsResolveReqAlerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsResolveReqAlerts() *AlertsResolveReqAlerts {
	this := AlertsResolveReqAlerts{}
	return &this
}

// NewAlertsResolveReqAlertsWithDefaults instantiates a new AlertsResolveReqAlerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsResolveReqAlertsWithDefaults() *AlertsResolveReqAlerts {
	this := AlertsResolveReqAlerts{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *AlertsResolveReqAlerts) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsResolveReqAlerts) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *AlertsResolveReqAlerts) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *AlertsResolveReqAlerts) SetGroup(v string) {
	o.Group = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *AlertsResolveReqAlerts) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsResolveReqAlerts) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *AlertsResolveReqAlerts) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *AlertsResolveReqAlerts) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o AlertsResolveReqAlerts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertsResolveReqAlerts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resource_type"] = o.ResourceType
	}
	return toSerialize, nil
}

type NullableAlertsResolveReqAlerts struct {
	value *AlertsResolveReqAlerts
	isSet bool
}

func (v NullableAlertsResolveReqAlerts) Get() *AlertsResolveReqAlerts {
	return v.value
}

func (v *NullableAlertsResolveReqAlerts) Set(val *AlertsResolveReqAlerts) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsResolveReqAlerts) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsResolveReqAlerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsResolveReqAlerts(val *AlertsResolveReqAlerts) *NullableAlertsResolveReqAlerts {
	return &NullableAlertsResolveReqAlerts{value: val, isSet: true}
}

func (v NullableAlertsResolveReqAlerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsResolveReqAlerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


