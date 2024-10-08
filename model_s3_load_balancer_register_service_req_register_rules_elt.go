/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the S3LoadBalancerRegisterServiceReqRegisterRulesElt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3LoadBalancerRegisterServiceReqRegisterRulesElt{}

// S3LoadBalancerRegisterServiceReqRegisterRulesElt struct for S3LoadBalancerRegisterServiceReqRegisterRulesElt
type S3LoadBalancerRegisterServiceReqRegisterRulesElt struct {
	// recognition rules of service
	RecognitionRules []string `json:"recognition_rules"`
	// type of service
	ServiceType *string `json:"service_type,omitempty"`
	// target endpoints of service
	TargetEndpoints []S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt `json:"target_endpoints"`
}

type _S3LoadBalancerRegisterServiceReqRegisterRulesElt S3LoadBalancerRegisterServiceReqRegisterRulesElt

// NewS3LoadBalancerRegisterServiceReqRegisterRulesElt instantiates a new S3LoadBalancerRegisterServiceReqRegisterRulesElt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3LoadBalancerRegisterServiceReqRegisterRulesElt(recognitionRules []string, targetEndpoints []S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) *S3LoadBalancerRegisterServiceReqRegisterRulesElt {
	this := S3LoadBalancerRegisterServiceReqRegisterRulesElt{}
	this.RecognitionRules = recognitionRules
	this.TargetEndpoints = targetEndpoints
	return &this
}

// NewS3LoadBalancerRegisterServiceReqRegisterRulesEltWithDefaults instantiates a new S3LoadBalancerRegisterServiceReqRegisterRulesElt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3LoadBalancerRegisterServiceReqRegisterRulesEltWithDefaults() *S3LoadBalancerRegisterServiceReqRegisterRulesElt {
	this := S3LoadBalancerRegisterServiceReqRegisterRulesElt{}
	return &this
}

// GetRecognitionRules returns the RecognitionRules field value
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetRecognitionRules() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RecognitionRules
}

// GetRecognitionRulesOk returns a tuple with the RecognitionRules field value
// and a boolean to check if the value has been set.
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetRecognitionRulesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecognitionRules, true
}

// SetRecognitionRules sets field value
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) SetRecognitionRules(v []string) {
	o.RecognitionRules = v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetServiceType() string {
	if o == nil || IsNil(o.ServiceType) {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetTargetEndpoints returns the TargetEndpoints field value
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetTargetEndpoints() []S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt {
	if o == nil {
		var ret []S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt
		return ret
	}

	return o.TargetEndpoints
}

// GetTargetEndpointsOk returns a tuple with the TargetEndpoints field value
// and a boolean to check if the value has been set.
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetTargetEndpointsOk() ([]S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetEndpoints, true
}

// SetTargetEndpoints sets field value
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) SetTargetEndpoints(v []S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) {
	o.TargetEndpoints = v
}

func (o S3LoadBalancerRegisterServiceReqRegisterRulesElt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3LoadBalancerRegisterServiceReqRegisterRulesElt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recognition_rules"] = o.RecognitionRules
	if !IsNil(o.ServiceType) {
		toSerialize["service_type"] = o.ServiceType
	}
	toSerialize["target_endpoints"] = o.TargetEndpoints
	return toSerialize, nil
}

func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recognition_rules",
		"target_endpoints",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varS3LoadBalancerRegisterServiceReqRegisterRulesElt := _S3LoadBalancerRegisterServiceReqRegisterRulesElt{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varS3LoadBalancerRegisterServiceReqRegisterRulesElt)

	if err != nil {
		return err
	}

	*o = S3LoadBalancerRegisterServiceReqRegisterRulesElt(varS3LoadBalancerRegisterServiceReqRegisterRulesElt)

	return err
}

type NullableS3LoadBalancerRegisterServiceReqRegisterRulesElt struct {
	value *S3LoadBalancerRegisterServiceReqRegisterRulesElt
	isSet bool
}

func (v NullableS3LoadBalancerRegisterServiceReqRegisterRulesElt) Get() *S3LoadBalancerRegisterServiceReqRegisterRulesElt {
	return v.value
}

func (v *NullableS3LoadBalancerRegisterServiceReqRegisterRulesElt) Set(val *S3LoadBalancerRegisterServiceReqRegisterRulesElt) {
	v.value = val
	v.isSet = true
}

func (v NullableS3LoadBalancerRegisterServiceReqRegisterRulesElt) IsSet() bool {
	return v.isSet
}

func (v *NullableS3LoadBalancerRegisterServiceReqRegisterRulesElt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3LoadBalancerRegisterServiceReqRegisterRulesElt(val *S3LoadBalancerRegisterServiceReqRegisterRulesElt) *NullableS3LoadBalancerRegisterServiceReqRegisterRulesElt {
	return &NullableS3LoadBalancerRegisterServiceReqRegisterRulesElt{value: val, isSet: true}
}

func (v NullableS3LoadBalancerRegisterServiceReqRegisterRulesElt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3LoadBalancerRegisterServiceReqRegisterRulesElt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


