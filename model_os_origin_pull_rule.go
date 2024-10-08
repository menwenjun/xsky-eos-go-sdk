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

// checks if the OSOriginPullRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSOriginPullRule{}

// OSOriginPullRule OSOriginPullRule is the model of os_origin_pull_rule
type OSOriginPullRule struct {
	BucketId *int64 `json:"bucket_id,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Connected *bool `json:"connected,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	EscapeToSlash *bool `json:"escape_to_slash,omitempty"`
	Id *int64 `json:"id,omitempty"`
	ModeType *string `json:"mode_type,omitempty"`
	OriginConf *OriginConf `json:"origin_conf,omitempty"`
	OriginInfo *OriginInfo `json:"origin_info,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
	S3LoadBalancerGroup *S3LoadBalancerGroup `json:"s3_load_balancer_group,omitempty"`
	Status *string `json:"status,omitempty"`
	Suffix *string `json:"suffix,omitempty"`
}

// NewOSOriginPullRule instantiates a new OSOriginPullRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSOriginPullRule() *OSOriginPullRule {
	this := OSOriginPullRule{}
	return &this
}

// NewOSOriginPullRuleWithDefaults instantiates a new OSOriginPullRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSOriginPullRuleWithDefaults() *OSOriginPullRule {
	this := OSOriginPullRule{}
	return &this
}

// GetBucketId returns the BucketId field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetBucketId() int64 {
	if o == nil || IsNil(o.BucketId) {
		var ret int64
		return ret
	}
	return *o.BucketId
}

// GetBucketIdOk returns a tuple with the BucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetBucketIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BucketId) {
		return nil, false
	}
	return o.BucketId, true
}

// HasBucketId returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasBucketId() bool {
	if o != nil && !IsNil(o.BucketId) {
		return true
	}

	return false
}

// SetBucketId gets a reference to the given int64 and assigns it to the BucketId field.
func (o *OSOriginPullRule) SetBucketId(v int64) {
	o.BucketId = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *OSOriginPullRule) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetConnected() bool {
	if o == nil || IsNil(o.Connected) {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Connected) {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasConnected() bool {
	if o != nil && !IsNil(o.Connected) {
		return true
	}

	return false
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *OSOriginPullRule) SetConnected(v bool) {
	o.Connected = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *OSOriginPullRule) SetCreate(v time.Time) {
	o.Create = &v
}

// GetEscapeToSlash returns the EscapeToSlash field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetEscapeToSlash() bool {
	if o == nil || IsNil(o.EscapeToSlash) {
		var ret bool
		return ret
	}
	return *o.EscapeToSlash
}

// GetEscapeToSlashOk returns a tuple with the EscapeToSlash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetEscapeToSlashOk() (*bool, bool) {
	if o == nil || IsNil(o.EscapeToSlash) {
		return nil, false
	}
	return o.EscapeToSlash, true
}

// HasEscapeToSlash returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasEscapeToSlash() bool {
	if o != nil && !IsNil(o.EscapeToSlash) {
		return true
	}

	return false
}

// SetEscapeToSlash gets a reference to the given bool and assigns it to the EscapeToSlash field.
func (o *OSOriginPullRule) SetEscapeToSlash(v bool) {
	o.EscapeToSlash = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OSOriginPullRule) SetId(v int64) {
	o.Id = &v
}

// GetModeType returns the ModeType field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetModeType() string {
	if o == nil || IsNil(o.ModeType) {
		var ret string
		return ret
	}
	return *o.ModeType
}

// GetModeTypeOk returns a tuple with the ModeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetModeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ModeType) {
		return nil, false
	}
	return o.ModeType, true
}

// HasModeType returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasModeType() bool {
	if o != nil && !IsNil(o.ModeType) {
		return true
	}

	return false
}

// SetModeType gets a reference to the given string and assigns it to the ModeType field.
func (o *OSOriginPullRule) SetModeType(v string) {
	o.ModeType = &v
}

// GetOriginConf returns the OriginConf field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetOriginConf() OriginConf {
	if o == nil || IsNil(o.OriginConf) {
		var ret OriginConf
		return ret
	}
	return *o.OriginConf
}

// GetOriginConfOk returns a tuple with the OriginConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetOriginConfOk() (*OriginConf, bool) {
	if o == nil || IsNil(o.OriginConf) {
		return nil, false
	}
	return o.OriginConf, true
}

// HasOriginConf returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasOriginConf() bool {
	if o != nil && !IsNil(o.OriginConf) {
		return true
	}

	return false
}

// SetOriginConf gets a reference to the given OriginConf and assigns it to the OriginConf field.
func (o *OSOriginPullRule) SetOriginConf(v OriginConf) {
	o.OriginConf = &v
}

// GetOriginInfo returns the OriginInfo field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetOriginInfo() OriginInfo {
	if o == nil || IsNil(o.OriginInfo) {
		var ret OriginInfo
		return ret
	}
	return *o.OriginInfo
}

// GetOriginInfoOk returns a tuple with the OriginInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetOriginInfoOk() (*OriginInfo, bool) {
	if o == nil || IsNil(o.OriginInfo) {
		return nil, false
	}
	return o.OriginInfo, true
}

// HasOriginInfo returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasOriginInfo() bool {
	if o != nil && !IsNil(o.OriginInfo) {
		return true
	}

	return false
}

// SetOriginInfo gets a reference to the given OriginInfo and assigns it to the OriginInfo field.
func (o *OSOriginPullRule) SetOriginInfo(v OriginInfo) {
	o.OriginInfo = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *OSOriginPullRule) SetPrefix(v string) {
	o.Prefix = &v
}

// GetS3LoadBalancerGroup returns the S3LoadBalancerGroup field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetS3LoadBalancerGroup() S3LoadBalancerGroup {
	if o == nil || IsNil(o.S3LoadBalancerGroup) {
		var ret S3LoadBalancerGroup
		return ret
	}
	return *o.S3LoadBalancerGroup
}

// GetS3LoadBalancerGroupOk returns a tuple with the S3LoadBalancerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetS3LoadBalancerGroupOk() (*S3LoadBalancerGroup, bool) {
	if o == nil || IsNil(o.S3LoadBalancerGroup) {
		return nil, false
	}
	return o.S3LoadBalancerGroup, true
}

// HasS3LoadBalancerGroup returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasS3LoadBalancerGroup() bool {
	if o != nil && !IsNil(o.S3LoadBalancerGroup) {
		return true
	}

	return false
}

// SetS3LoadBalancerGroup gets a reference to the given S3LoadBalancerGroup and assigns it to the S3LoadBalancerGroup field.
func (o *OSOriginPullRule) SetS3LoadBalancerGroup(v S3LoadBalancerGroup) {
	o.S3LoadBalancerGroup = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OSOriginPullRule) SetStatus(v string) {
	o.Status = &v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise.
func (o *OSOriginPullRule) GetSuffix() string {
	if o == nil || IsNil(o.Suffix) {
		var ret string
		return ret
	}
	return *o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSOriginPullRule) GetSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.Suffix) {
		return nil, false
	}
	return o.Suffix, true
}

// HasSuffix returns a boolean if a field has been set.
func (o *OSOriginPullRule) HasSuffix() bool {
	if o != nil && !IsNil(o.Suffix) {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given string and assigns it to the Suffix field.
func (o *OSOriginPullRule) SetSuffix(v string) {
	o.Suffix = &v
}

func (o OSOriginPullRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSOriginPullRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BucketId) {
		toSerialize["bucket_id"] = o.BucketId
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Connected) {
		toSerialize["connected"] = o.Connected
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.EscapeToSlash) {
		toSerialize["escape_to_slash"] = o.EscapeToSlash
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ModeType) {
		toSerialize["mode_type"] = o.ModeType
	}
	if !IsNil(o.OriginConf) {
		toSerialize["origin_conf"] = o.OriginConf
	}
	if !IsNil(o.OriginInfo) {
		toSerialize["origin_info"] = o.OriginInfo
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.S3LoadBalancerGroup) {
		toSerialize["s3_load_balancer_group"] = o.S3LoadBalancerGroup
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Suffix) {
		toSerialize["suffix"] = o.Suffix
	}
	return toSerialize, nil
}

type NullableOSOriginPullRule struct {
	value *OSOriginPullRule
	isSet bool
}

func (v NullableOSOriginPullRule) Get() *OSOriginPullRule {
	return v.value
}

func (v *NullableOSOriginPullRule) Set(val *OSOriginPullRule) {
	v.value = val
	v.isSet = true
}

func (v NullableOSOriginPullRule) IsSet() bool {
	return v.isSet
}

func (v *NullableOSOriginPullRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSOriginPullRule(val *OSOriginPullRule) *NullableOSOriginPullRule {
	return &NullableOSOriginPullRule{value: val, isSet: true}
}

func (v NullableOSOriginPullRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSOriginPullRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


