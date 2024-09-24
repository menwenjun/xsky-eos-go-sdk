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

// checks if the DpBlockBackupPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DpBlockBackupPolicy{}

// DpBlockBackupPolicy DpBlockBackupPolicy protects a block volume by backup snapshot to remote site
type DpBlockBackupPolicy struct {
	BlockVolumeNum *int64 `json:"block_volume_num,omitempty"`
	Cluster *Cluster `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DpGateway *DpGatewayNestview `json:"dp_gateway,omitempty"`
	DpSite *DpSiteNestview `json:"dp_site,omitempty"`
	Encrypted *bool `json:"encrypted,omitempty"`
	FullCronExpr *string `json:"full_cron_expr,omitempty"`
	Id *int64 `json:"id,omitempty"`
	IncCronExpr *string `json:"inc_cron_expr,omitempty"`
	Name *string `json:"name,omitempty"`
	RetainedMax *int64 `json:"retained_max,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewDpBlockBackupPolicy instantiates a new DpBlockBackupPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDpBlockBackupPolicy() *DpBlockBackupPolicy {
	this := DpBlockBackupPolicy{}
	return &this
}

// NewDpBlockBackupPolicyWithDefaults instantiates a new DpBlockBackupPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDpBlockBackupPolicyWithDefaults() *DpBlockBackupPolicy {
	this := DpBlockBackupPolicy{}
	return &this
}

// GetBlockVolumeNum returns the BlockVolumeNum field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetBlockVolumeNum() int64 {
	if o == nil || IsNil(o.BlockVolumeNum) {
		var ret int64
		return ret
	}
	return *o.BlockVolumeNum
}

// GetBlockVolumeNumOk returns a tuple with the BlockVolumeNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetBlockVolumeNumOk() (*int64, bool) {
	if o == nil || IsNil(o.BlockVolumeNum) {
		return nil, false
	}
	return o.BlockVolumeNum, true
}

// HasBlockVolumeNum returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasBlockVolumeNum() bool {
	if o != nil && !IsNil(o.BlockVolumeNum) {
		return true
	}

	return false
}

// SetBlockVolumeNum gets a reference to the given int64 and assigns it to the BlockVolumeNum field.
func (o *DpBlockBackupPolicy) SetBlockVolumeNum(v int64) {
	o.BlockVolumeNum = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetCluster() Cluster {
	if o == nil || IsNil(o.Cluster) {
		var ret Cluster
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetClusterOk() (*Cluster, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given Cluster and assigns it to the Cluster field.
func (o *DpBlockBackupPolicy) SetCluster(v Cluster) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DpBlockBackupPolicy) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDpGateway returns the DpGateway field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetDpGateway() DpGatewayNestview {
	if o == nil || IsNil(o.DpGateway) {
		var ret DpGatewayNestview
		return ret
	}
	return *o.DpGateway
}

// GetDpGatewayOk returns a tuple with the DpGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetDpGatewayOk() (*DpGatewayNestview, bool) {
	if o == nil || IsNil(o.DpGateway) {
		return nil, false
	}
	return o.DpGateway, true
}

// HasDpGateway returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasDpGateway() bool {
	if o != nil && !IsNil(o.DpGateway) {
		return true
	}

	return false
}

// SetDpGateway gets a reference to the given DpGatewayNestview and assigns it to the DpGateway field.
func (o *DpBlockBackupPolicy) SetDpGateway(v DpGatewayNestview) {
	o.DpGateway = &v
}

// GetDpSite returns the DpSite field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetDpSite() DpSiteNestview {
	if o == nil || IsNil(o.DpSite) {
		var ret DpSiteNestview
		return ret
	}
	return *o.DpSite
}

// GetDpSiteOk returns a tuple with the DpSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetDpSiteOk() (*DpSiteNestview, bool) {
	if o == nil || IsNil(o.DpSite) {
		return nil, false
	}
	return o.DpSite, true
}

// HasDpSite returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasDpSite() bool {
	if o != nil && !IsNil(o.DpSite) {
		return true
	}

	return false
}

// SetDpSite gets a reference to the given DpSiteNestview and assigns it to the DpSite field.
func (o *DpBlockBackupPolicy) SetDpSite(v DpSiteNestview) {
	o.DpSite = &v
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetEncrypted() bool {
	if o == nil || IsNil(o.Encrypted) {
		var ret bool
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.Encrypted) {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasEncrypted() bool {
	if o != nil && !IsNil(o.Encrypted) {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given bool and assigns it to the Encrypted field.
func (o *DpBlockBackupPolicy) SetEncrypted(v bool) {
	o.Encrypted = &v
}

// GetFullCronExpr returns the FullCronExpr field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetFullCronExpr() string {
	if o == nil || IsNil(o.FullCronExpr) {
		var ret string
		return ret
	}
	return *o.FullCronExpr
}

// GetFullCronExprOk returns a tuple with the FullCronExpr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetFullCronExprOk() (*string, bool) {
	if o == nil || IsNil(o.FullCronExpr) {
		return nil, false
	}
	return o.FullCronExpr, true
}

// HasFullCronExpr returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasFullCronExpr() bool {
	if o != nil && !IsNil(o.FullCronExpr) {
		return true
	}

	return false
}

// SetFullCronExpr gets a reference to the given string and assigns it to the FullCronExpr field.
func (o *DpBlockBackupPolicy) SetFullCronExpr(v string) {
	o.FullCronExpr = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DpBlockBackupPolicy) SetId(v int64) {
	o.Id = &v
}

// GetIncCronExpr returns the IncCronExpr field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetIncCronExpr() string {
	if o == nil || IsNil(o.IncCronExpr) {
		var ret string
		return ret
	}
	return *o.IncCronExpr
}

// GetIncCronExprOk returns a tuple with the IncCronExpr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetIncCronExprOk() (*string, bool) {
	if o == nil || IsNil(o.IncCronExpr) {
		return nil, false
	}
	return o.IncCronExpr, true
}

// HasIncCronExpr returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasIncCronExpr() bool {
	if o != nil && !IsNil(o.IncCronExpr) {
		return true
	}

	return false
}

// SetIncCronExpr gets a reference to the given string and assigns it to the IncCronExpr field.
func (o *DpBlockBackupPolicy) SetIncCronExpr(v string) {
	o.IncCronExpr = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DpBlockBackupPolicy) SetName(v string) {
	o.Name = &v
}

// GetRetainedMax returns the RetainedMax field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetRetainedMax() int64 {
	if o == nil || IsNil(o.RetainedMax) {
		var ret int64
		return ret
	}
	return *o.RetainedMax
}

// GetRetainedMaxOk returns a tuple with the RetainedMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetRetainedMaxOk() (*int64, bool) {
	if o == nil || IsNil(o.RetainedMax) {
		return nil, false
	}
	return o.RetainedMax, true
}

// HasRetainedMax returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasRetainedMax() bool {
	if o != nil && !IsNil(o.RetainedMax) {
		return true
	}

	return false
}

// SetRetainedMax gets a reference to the given int64 and assigns it to the RetainedMax field.
func (o *DpBlockBackupPolicy) SetRetainedMax(v int64) {
	o.RetainedMax = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DpBlockBackupPolicy) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DpBlockBackupPolicy) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DpBlockBackupPolicy) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DpBlockBackupPolicy) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DpBlockBackupPolicy) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o DpBlockBackupPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DpBlockBackupPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockVolumeNum) {
		toSerialize["block_volume_num"] = o.BlockVolumeNum
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DpGateway) {
		toSerialize["dp_gateway"] = o.DpGateway
	}
	if !IsNil(o.DpSite) {
		toSerialize["dp_site"] = o.DpSite
	}
	if !IsNil(o.Encrypted) {
		toSerialize["encrypted"] = o.Encrypted
	}
	if !IsNil(o.FullCronExpr) {
		toSerialize["full_cron_expr"] = o.FullCronExpr
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IncCronExpr) {
		toSerialize["inc_cron_expr"] = o.IncCronExpr
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RetainedMax) {
		toSerialize["retained_max"] = o.RetainedMax
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableDpBlockBackupPolicy struct {
	value *DpBlockBackupPolicy
	isSet bool
}

func (v NullableDpBlockBackupPolicy) Get() *DpBlockBackupPolicy {
	return v.value
}

func (v *NullableDpBlockBackupPolicy) Set(val *DpBlockBackupPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDpBlockBackupPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDpBlockBackupPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDpBlockBackupPolicy(val *DpBlockBackupPolicy) *NullableDpBlockBackupPolicy {
	return &NullableDpBlockBackupPolicy{value: val, isSet: true}
}

func (v NullableDpBlockBackupPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDpBlockBackupPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


