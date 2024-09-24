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

// checks if the BatchReqAPIRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchReqAPIRequest{}

// BatchReqAPIRequest struct for BatchReqAPIRequest
type BatchReqAPIRequest struct {
	// API request specification.
	Body map[string]interface{} `json:"body,omitempty"`
	// API request method.
	Method string `json:"method"`
	// API request path and query.
	PathAndParams string `json:"path_and_params"`
}

type _BatchReqAPIRequest BatchReqAPIRequest

// NewBatchReqAPIRequest instantiates a new BatchReqAPIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchReqAPIRequest(method string, pathAndParams string) *BatchReqAPIRequest {
	this := BatchReqAPIRequest{}
	this.Method = method
	this.PathAndParams = pathAndParams
	return &this
}

// NewBatchReqAPIRequestWithDefaults instantiates a new BatchReqAPIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchReqAPIRequestWithDefaults() *BatchReqAPIRequest {
	this := BatchReqAPIRequest{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BatchReqAPIRequest) GetBody() map[string]interface{} {
	if o == nil || IsNil(o.Body) {
		var ret map[string]interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchReqAPIRequest) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Body) {
		return map[string]interface{}{}, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BatchReqAPIRequest) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given map[string]interface{} and assigns it to the Body field.
func (o *BatchReqAPIRequest) SetBody(v map[string]interface{}) {
	o.Body = v
}

// GetMethod returns the Method field value
func (o *BatchReqAPIRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *BatchReqAPIRequest) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *BatchReqAPIRequest) SetMethod(v string) {
	o.Method = v
}

// GetPathAndParams returns the PathAndParams field value
func (o *BatchReqAPIRequest) GetPathAndParams() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PathAndParams
}

// GetPathAndParamsOk returns a tuple with the PathAndParams field value
// and a boolean to check if the value has been set.
func (o *BatchReqAPIRequest) GetPathAndParamsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PathAndParams, true
}

// SetPathAndParams sets field value
func (o *BatchReqAPIRequest) SetPathAndParams(v string) {
	o.PathAndParams = v
}

func (o BatchReqAPIRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchReqAPIRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	toSerialize["method"] = o.Method
	toSerialize["path_and_params"] = o.PathAndParams
	return toSerialize, nil
}

func (o *BatchReqAPIRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"path_and_params",
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

	varBatchReqAPIRequest := _BatchReqAPIRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchReqAPIRequest)

	if err != nil {
		return err
	}

	*o = BatchReqAPIRequest(varBatchReqAPIRequest)

	return err
}

type NullableBatchReqAPIRequest struct {
	value *BatchReqAPIRequest
	isSet bool
}

func (v NullableBatchReqAPIRequest) Get() *BatchReqAPIRequest {
	return v.value
}

func (v *NullableBatchReqAPIRequest) Set(val *BatchReqAPIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchReqAPIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchReqAPIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchReqAPIRequest(val *BatchReqAPIRequest) *NullableBatchReqAPIRequest {
	return &NullableBatchReqAPIRequest{value: val, isSet: true}
}

func (v NullableBatchReqAPIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchReqAPIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


