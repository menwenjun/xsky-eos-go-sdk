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

// checks if the AlertWechatConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertWechatConfig{}

// AlertWechatConfig AlertWechatConfig define alert wechat config
type AlertWechatConfig struct {
	AgentId *string `json:"agent_id,omitempty"`
	ApiSecret *string `json:"api_secret,omitempty"`
	ApiUrl *string `json:"api_url,omitempty"`
	CorpId *string `json:"corp_id,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewAlertWechatConfig instantiates a new AlertWechatConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertWechatConfig() *AlertWechatConfig {
	this := AlertWechatConfig{}
	return &this
}

// NewAlertWechatConfigWithDefaults instantiates a new AlertWechatConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertWechatConfigWithDefaults() *AlertWechatConfig {
	this := AlertWechatConfig{}
	return &this
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *AlertWechatConfig) GetAgentId() string {
	if o == nil || IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertWechatConfig) GetAgentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *AlertWechatConfig) HasAgentId() bool {
	if o != nil && !IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *AlertWechatConfig) SetAgentId(v string) {
	o.AgentId = &v
}

// GetApiSecret returns the ApiSecret field value if set, zero value otherwise.
func (o *AlertWechatConfig) GetApiSecret() string {
	if o == nil || IsNil(o.ApiSecret) {
		var ret string
		return ret
	}
	return *o.ApiSecret
}

// GetApiSecretOk returns a tuple with the ApiSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertWechatConfig) GetApiSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ApiSecret) {
		return nil, false
	}
	return o.ApiSecret, true
}

// HasApiSecret returns a boolean if a field has been set.
func (o *AlertWechatConfig) HasApiSecret() bool {
	if o != nil && !IsNil(o.ApiSecret) {
		return true
	}

	return false
}

// SetApiSecret gets a reference to the given string and assigns it to the ApiSecret field.
func (o *AlertWechatConfig) SetApiSecret(v string) {
	o.ApiSecret = &v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *AlertWechatConfig) GetApiUrl() string {
	if o == nil || IsNil(o.ApiUrl) {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertWechatConfig) GetApiUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ApiUrl) {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *AlertWechatConfig) HasApiUrl() bool {
	if o != nil && !IsNil(o.ApiUrl) {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *AlertWechatConfig) SetApiUrl(v string) {
	o.ApiUrl = &v
}

// GetCorpId returns the CorpId field value if set, zero value otherwise.
func (o *AlertWechatConfig) GetCorpId() string {
	if o == nil || IsNil(o.CorpId) {
		var ret string
		return ret
	}
	return *o.CorpId
}

// GetCorpIdOk returns a tuple with the CorpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertWechatConfig) GetCorpIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorpId) {
		return nil, false
	}
	return o.CorpId, true
}

// HasCorpId returns a boolean if a field has been set.
func (o *AlertWechatConfig) HasCorpId() bool {
	if o != nil && !IsNil(o.CorpId) {
		return true
	}

	return false
}

// SetCorpId gets a reference to the given string and assigns it to the CorpId field.
func (o *AlertWechatConfig) SetCorpId(v string) {
	o.CorpId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AlertWechatConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertWechatConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AlertWechatConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AlertWechatConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o AlertWechatConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertWechatConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentId) {
		toSerialize["agent_id"] = o.AgentId
	}
	if !IsNil(o.ApiSecret) {
		toSerialize["api_secret"] = o.ApiSecret
	}
	if !IsNil(o.ApiUrl) {
		toSerialize["api_url"] = o.ApiUrl
	}
	if !IsNil(o.CorpId) {
		toSerialize["corp_id"] = o.CorpId
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableAlertWechatConfig struct {
	value *AlertWechatConfig
	isSet bool
}

func (v NullableAlertWechatConfig) Get() *AlertWechatConfig {
	return v.value
}

func (v *NullableAlertWechatConfig) Set(val *AlertWechatConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertWechatConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertWechatConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertWechatConfig(val *AlertWechatConfig) *NullableAlertWechatConfig {
	return &NullableAlertWechatConfig{value: val, isSet: true}
}

func (v NullableAlertWechatConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertWechatConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


