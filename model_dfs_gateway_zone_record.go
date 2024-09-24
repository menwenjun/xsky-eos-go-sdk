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

// checks if the DfsGatewayZoneRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsGatewayZoneRecord{}

// DfsGatewayZoneRecord DfsGatewayZoneRecord combine DfsGatewayZone and DfsGatewayZoneStat as API response
type DfsGatewayZoneRecord struct {
	ActionStatus *string `json:"action_status,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Default *bool `json:"default,omitempty"`
	DfsGatewayGroup *DfsGatewayGroupNestview `json:"dfs_gateway_group,omitempty"`
	DfsHdfsNum *int64 `json:"dfs_hdfs_num,omitempty"`
	DfsHdfses []DfsHdfs `json:"dfs_hdfses,omitempty"`
	DnsDomainNames []DNSDomainName `json:"dns_domain_names,omitempty"`
	GatewayNum *int64 `json:"gateway_num,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	VipBalanceStatus *string `json:"vip_balance_status,omitempty"`
	// gateways for VIP addresses in policy routing
	VipGateways []string `json:"vip_gateways,omitempty"`
	VipNum *int64 `json:"vip_num,omitempty"`
	Samples []DfsGatewayZoneStat `json:"samples,omitempty"`
}

// NewDfsGatewayZoneRecord instantiates a new DfsGatewayZoneRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsGatewayZoneRecord() *DfsGatewayZoneRecord {
	this := DfsGatewayZoneRecord{}
	return &this
}

// NewDfsGatewayZoneRecordWithDefaults instantiates a new DfsGatewayZoneRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsGatewayZoneRecordWithDefaults() *DfsGatewayZoneRecord {
	this := DfsGatewayZoneRecord{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *DfsGatewayZoneRecord) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DfsGatewayZoneRecord) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsGatewayZoneRecord) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *DfsGatewayZoneRecord) SetDefault(v bool) {
	o.Default = &v
}

// GetDfsGatewayGroup returns the DfsGatewayGroup field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetDfsGatewayGroup() DfsGatewayGroupNestview {
	if o == nil || IsNil(o.DfsGatewayGroup) {
		var ret DfsGatewayGroupNestview
		return ret
	}
	return *o.DfsGatewayGroup
}

// GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool) {
	if o == nil || IsNil(o.DfsGatewayGroup) {
		return nil, false
	}
	return o.DfsGatewayGroup, true
}

// HasDfsGatewayGroup returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasDfsGatewayGroup() bool {
	if o != nil && !IsNil(o.DfsGatewayGroup) {
		return true
	}

	return false
}

// SetDfsGatewayGroup gets a reference to the given DfsGatewayGroupNestview and assigns it to the DfsGatewayGroup field.
func (o *DfsGatewayZoneRecord) SetDfsGatewayGroup(v DfsGatewayGroupNestview) {
	o.DfsGatewayGroup = &v
}

// GetDfsHdfsNum returns the DfsHdfsNum field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetDfsHdfsNum() int64 {
	if o == nil || IsNil(o.DfsHdfsNum) {
		var ret int64
		return ret
	}
	return *o.DfsHdfsNum
}

// GetDfsHdfsNumOk returns a tuple with the DfsHdfsNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetDfsHdfsNumOk() (*int64, bool) {
	if o == nil || IsNil(o.DfsHdfsNum) {
		return nil, false
	}
	return o.DfsHdfsNum, true
}

// HasDfsHdfsNum returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasDfsHdfsNum() bool {
	if o != nil && !IsNil(o.DfsHdfsNum) {
		return true
	}

	return false
}

// SetDfsHdfsNum gets a reference to the given int64 and assigns it to the DfsHdfsNum field.
func (o *DfsGatewayZoneRecord) SetDfsHdfsNum(v int64) {
	o.DfsHdfsNum = &v
}

// GetDfsHdfses returns the DfsHdfses field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetDfsHdfses() []DfsHdfs {
	if o == nil || IsNil(o.DfsHdfses) {
		var ret []DfsHdfs
		return ret
	}
	return o.DfsHdfses
}

// GetDfsHdfsesOk returns a tuple with the DfsHdfses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetDfsHdfsesOk() ([]DfsHdfs, bool) {
	if o == nil || IsNil(o.DfsHdfses) {
		return nil, false
	}
	return o.DfsHdfses, true
}

// HasDfsHdfses returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasDfsHdfses() bool {
	if o != nil && !IsNil(o.DfsHdfses) {
		return true
	}

	return false
}

// SetDfsHdfses gets a reference to the given []DfsHdfs and assigns it to the DfsHdfses field.
func (o *DfsGatewayZoneRecord) SetDfsHdfses(v []DfsHdfs) {
	o.DfsHdfses = v
}

// GetDnsDomainNames returns the DnsDomainNames field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetDnsDomainNames() []DNSDomainName {
	if o == nil || IsNil(o.DnsDomainNames) {
		var ret []DNSDomainName
		return ret
	}
	return o.DnsDomainNames
}

// GetDnsDomainNamesOk returns a tuple with the DnsDomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetDnsDomainNamesOk() ([]DNSDomainName, bool) {
	if o == nil || IsNil(o.DnsDomainNames) {
		return nil, false
	}
	return o.DnsDomainNames, true
}

// HasDnsDomainNames returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasDnsDomainNames() bool {
	if o != nil && !IsNil(o.DnsDomainNames) {
		return true
	}

	return false
}

// SetDnsDomainNames gets a reference to the given []DNSDomainName and assigns it to the DnsDomainNames field.
func (o *DfsGatewayZoneRecord) SetDnsDomainNames(v []DNSDomainName) {
	o.DnsDomainNames = v
}

// GetGatewayNum returns the GatewayNum field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetGatewayNum() int64 {
	if o == nil || IsNil(o.GatewayNum) {
		var ret int64
		return ret
	}
	return *o.GatewayNum
}

// GetGatewayNumOk returns a tuple with the GatewayNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetGatewayNumOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayNum) {
		return nil, false
	}
	return o.GatewayNum, true
}

// HasGatewayNum returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasGatewayNum() bool {
	if o != nil && !IsNil(o.GatewayNum) {
		return true
	}

	return false
}

// SetGatewayNum gets a reference to the given int64 and assigns it to the GatewayNum field.
func (o *DfsGatewayZoneRecord) SetGatewayNum(v int64) {
	o.GatewayNum = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsGatewayZoneRecord) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DfsGatewayZoneRecord) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DfsGatewayZoneRecord) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DfsGatewayZoneRecord) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetVipBalanceStatus returns the VipBalanceStatus field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetVipBalanceStatus() string {
	if o == nil || IsNil(o.VipBalanceStatus) {
		var ret string
		return ret
	}
	return *o.VipBalanceStatus
}

// GetVipBalanceStatusOk returns a tuple with the VipBalanceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetVipBalanceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VipBalanceStatus) {
		return nil, false
	}
	return o.VipBalanceStatus, true
}

// HasVipBalanceStatus returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasVipBalanceStatus() bool {
	if o != nil && !IsNil(o.VipBalanceStatus) {
		return true
	}

	return false
}

// SetVipBalanceStatus gets a reference to the given string and assigns it to the VipBalanceStatus field.
func (o *DfsGatewayZoneRecord) SetVipBalanceStatus(v string) {
	o.VipBalanceStatus = &v
}

// GetVipGateways returns the VipGateways field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetVipGateways() []string {
	if o == nil || IsNil(o.VipGateways) {
		var ret []string
		return ret
	}
	return o.VipGateways
}

// GetVipGatewaysOk returns a tuple with the VipGateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetVipGatewaysOk() ([]string, bool) {
	if o == nil || IsNil(o.VipGateways) {
		return nil, false
	}
	return o.VipGateways, true
}

// HasVipGateways returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasVipGateways() bool {
	if o != nil && !IsNil(o.VipGateways) {
		return true
	}

	return false
}

// SetVipGateways gets a reference to the given []string and assigns it to the VipGateways field.
func (o *DfsGatewayZoneRecord) SetVipGateways(v []string) {
	o.VipGateways = v
}

// GetVipNum returns the VipNum field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetVipNum() int64 {
	if o == nil || IsNil(o.VipNum) {
		var ret int64
		return ret
	}
	return *o.VipNum
}

// GetVipNumOk returns a tuple with the VipNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetVipNumOk() (*int64, bool) {
	if o == nil || IsNil(o.VipNum) {
		return nil, false
	}
	return o.VipNum, true
}

// HasVipNum returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasVipNum() bool {
	if o != nil && !IsNil(o.VipNum) {
		return true
	}

	return false
}

// SetVipNum gets a reference to the given int64 and assigns it to the VipNum field.
func (o *DfsGatewayZoneRecord) SetVipNum(v int64) {
	o.VipNum = &v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *DfsGatewayZoneRecord) GetSamples() []DfsGatewayZoneStat {
	if o == nil || IsNil(o.Samples) {
		var ret []DfsGatewayZoneStat
		return ret
	}
	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayZoneRecord) GetSamplesOk() ([]DfsGatewayZoneStat, bool) {
	if o == nil || IsNil(o.Samples) {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *DfsGatewayZoneRecord) HasSamples() bool {
	if o != nil && !IsNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given []DfsGatewayZoneStat and assigns it to the Samples field.
func (o *DfsGatewayZoneRecord) SetSamples(v []DfsGatewayZoneStat) {
	o.Samples = v
}

func (o DfsGatewayZoneRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsGatewayZoneRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.DfsGatewayGroup) {
		toSerialize["dfs_gateway_group"] = o.DfsGatewayGroup
	}
	if !IsNil(o.DfsHdfsNum) {
		toSerialize["dfs_hdfs_num"] = o.DfsHdfsNum
	}
	if !IsNil(o.DfsHdfses) {
		toSerialize["dfs_hdfses"] = o.DfsHdfses
	}
	if !IsNil(o.DnsDomainNames) {
		toSerialize["dns_domain_names"] = o.DnsDomainNames
	}
	if !IsNil(o.GatewayNum) {
		toSerialize["gateway_num"] = o.GatewayNum
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.VipBalanceStatus) {
		toSerialize["vip_balance_status"] = o.VipBalanceStatus
	}
	if !IsNil(o.VipGateways) {
		toSerialize["vip_gateways"] = o.VipGateways
	}
	if !IsNil(o.VipNum) {
		toSerialize["vip_num"] = o.VipNum
	}
	if !IsNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	return toSerialize, nil
}

type NullableDfsGatewayZoneRecord struct {
	value *DfsGatewayZoneRecord
	isSet bool
}

func (v NullableDfsGatewayZoneRecord) Get() *DfsGatewayZoneRecord {
	return v.value
}

func (v *NullableDfsGatewayZoneRecord) Set(val *DfsGatewayZoneRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsGatewayZoneRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsGatewayZoneRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsGatewayZoneRecord(val *DfsGatewayZoneRecord) *NullableDfsGatewayZoneRecord {
	return &NullableDfsGatewayZoneRecord{value: val, isSet: true}
}

func (v NullableDfsGatewayZoneRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsGatewayZoneRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


