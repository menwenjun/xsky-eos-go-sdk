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

// checks if the DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy{}

// DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy struct for DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy
type DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy struct {
	// cron expression
	CronExpr string `json:"cron_expr"`
	// dp gateway
	DpGatewayId int64 `json:"dp_gateway_id"`
	// dp site
	DpSiteId int64 `json:"dp_site_id"`
	// name
	Name string `json:"name"`
	// retained max number of snapshots
	RetainedMax *int64 `json:"retained_max,omitempty"`
	// timeout
	Timeout *int64 `json:"timeout,omitempty"`
}

type _DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy

// NewDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy instantiates a new DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy(cronExpr string, dpGatewayId int64, dpSiteId int64, name string) *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy {
	this := DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy{}
	this.CronExpr = cronExpr
	this.DpGatewayId = dpGatewayId
	this.DpSiteId = dpSiteId
	this.Name = name
	return &this
}

// NewDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicyWithDefaults instantiates a new DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicyWithDefaults() *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy {
	this := DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy{}
	return &this
}

// GetCronExpr returns the CronExpr field value
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetCronExpr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CronExpr
}

// GetCronExprOk returns a tuple with the CronExpr field value
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetCronExprOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CronExpr, true
}

// SetCronExpr sets field value
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) SetCronExpr(v string) {
	o.CronExpr = v
}

// GetDpGatewayId returns the DpGatewayId field value
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetDpGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DpGatewayId
}

// GetDpGatewayIdOk returns a tuple with the DpGatewayId field value
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetDpGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DpGatewayId, true
}

// SetDpGatewayId sets field value
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) SetDpGatewayId(v int64) {
	o.DpGatewayId = v
}

// GetDpSiteId returns the DpSiteId field value
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetDpSiteId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DpSiteId
}

// GetDpSiteIdOk returns a tuple with the DpSiteId field value
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetDpSiteIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DpSiteId, true
}

// SetDpSiteId sets field value
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) SetDpSiteId(v int64) {
	o.DpSiteId = v
}

// GetName returns the Name field value
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) SetName(v string) {
	o.Name = v
}

// GetRetainedMax returns the RetainedMax field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetRetainedMax() int64 {
	if o == nil || IsNil(o.RetainedMax) {
		var ret int64
		return ret
	}
	return *o.RetainedMax
}

// GetRetainedMaxOk returns a tuple with the RetainedMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetRetainedMaxOk() (*int64, bool) {
	if o == nil || IsNil(o.RetainedMax) {
		return nil, false
	}
	return o.RetainedMax, true
}

// HasRetainedMax returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) HasRetainedMax() bool {
	if o != nil && !IsNil(o.RetainedMax) {
		return true
	}

	return false
}

// SetRetainedMax gets a reference to the given int64 and assigns it to the RetainedMax field.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) SetRetainedMax(v int64) {
	o.RetainedMax = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetTimeout() int64 {
	if o == nil || IsNil(o.Timeout) {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) GetTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cron_expr"] = o.CronExpr
	toSerialize["dp_gateway_id"] = o.DpGatewayId
	toSerialize["dp_site_id"] = o.DpSiteId
	toSerialize["name"] = o.Name
	if !IsNil(o.RetainedMax) {
		toSerialize["retained_max"] = o.RetainedMax
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}

func (o *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cron_expr",
		"dp_gateway_id",
		"dp_site_id",
		"name",
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

	varDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy := _DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy)

	if err != nil {
		return err
	}

	*o = DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy(varDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy)

	return err
}

type NullableDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy struct {
	value *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy
	isSet bool
}

func (v NullableDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) Get() *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy {
	return v.value
}

func (v *NullableDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) Set(val *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy(val *DpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) *NullableDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy {
	return &NullableDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy{value: val, isSet: true}
}

func (v NullableDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpVolumeGroupSnapshotReplicationPolicyCreateReqPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


