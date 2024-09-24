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

// checks if the S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt{}

// S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt struct for S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt
type S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt struct {
	// endpoints of service
	Endpoints []string `json:"endpoints"`
	// zone id of load balancers
	OspZoneId int64 `json:"osp_zone_id"`
}

type _S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt

// NewS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt instantiates a new S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt(endpoints []string, ospZoneId int64) *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt {
	this := S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt{}
	this.Endpoints = endpoints
	this.OspZoneId = ospZoneId
	return &this
}

// NewS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsEltWithDefaults instantiates a new S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsEltWithDefaults() *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt {
	this := S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt{}
	return &this
}

// GetEndpoints returns the Endpoints field value
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) GetEndpoints() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value
// and a boolean to check if the value has been set.
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) GetEndpointsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// SetEndpoints sets field value
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) SetEndpoints(v []string) {
	o.Endpoints = v
}

// GetOspZoneId returns the OspZoneId field value
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) GetOspZoneId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OspZoneId
}

// GetOspZoneIdOk returns a tuple with the OspZoneId field value
// and a boolean to check if the value has been set.
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) GetOspZoneIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OspZoneId, true
}

// SetOspZoneId sets field value
func (o *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) SetOspZoneId(v int64) {
	o.OspZoneId = v
}

func (o S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endpoints"] = o.Endpoints
	toSerialize["osp_zone_id"] = o.OspZoneId
	return toSerialize, nil
}

func (o *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"endpoints",
		"osp_zone_id",
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

	varS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt := _S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt)

	if err != nil {
		return err
	}

	*o = S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt(varS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt)

	return err
}

type NullableS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt struct {
	value *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt
	isSet bool
}

func (v NullableS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) Get() *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt {
	return v.value
}

func (v *NullableS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) Set(val *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) {
	v.value = val
	v.isSet = true
}

func (v NullableS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) IsSet() bool {
	return v.isSet
}

func (v *NullableS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt(val *S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) *NullableS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt {
	return &NullableS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt{value: val, isSet: true}
}

func (v NullableS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


