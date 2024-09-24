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

// checks if the TargetRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetRecord{}

// TargetRecord struct for TargetRecord
type TargetRecord struct {
	AccessPath *AccessPathNestview `json:"access_path,omitempty"`
	Board *int64 `json:"board,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DcName *string `json:"dc_name,omitempty"`
	GatewayIps *string `json:"gateway_ips,omitempty"`
	Host *HostNestview `json:"host,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Iqn *string `json:"iqn,omitempty"`
	Logout *bool `json:"logout,omitempty"`
	Port *int64 `json:"port,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	LunsInfo []LunInfo `json:"luns_info,omitempty"`
}

// NewTargetRecord instantiates a new TargetRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetRecord() *TargetRecord {
	this := TargetRecord{}
	return &this
}

// NewTargetRecordWithDefaults instantiates a new TargetRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetRecordWithDefaults() *TargetRecord {
	this := TargetRecord{}
	return &this
}

// GetAccessPath returns the AccessPath field value if set, zero value otherwise.
func (o *TargetRecord) GetAccessPath() AccessPathNestview {
	if o == nil || IsNil(o.AccessPath) {
		var ret AccessPathNestview
		return ret
	}
	return *o.AccessPath
}

// GetAccessPathOk returns a tuple with the AccessPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetAccessPathOk() (*AccessPathNestview, bool) {
	if o == nil || IsNil(o.AccessPath) {
		return nil, false
	}
	return o.AccessPath, true
}

// HasAccessPath returns a boolean if a field has been set.
func (o *TargetRecord) HasAccessPath() bool {
	if o != nil && !IsNil(o.AccessPath) {
		return true
	}

	return false
}

// SetAccessPath gets a reference to the given AccessPathNestview and assigns it to the AccessPath field.
func (o *TargetRecord) SetAccessPath(v AccessPathNestview) {
	o.AccessPath = &v
}

// GetBoard returns the Board field value if set, zero value otherwise.
func (o *TargetRecord) GetBoard() int64 {
	if o == nil || IsNil(o.Board) {
		var ret int64
		return ret
	}
	return *o.Board
}

// GetBoardOk returns a tuple with the Board field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetBoardOk() (*int64, bool) {
	if o == nil || IsNil(o.Board) {
		return nil, false
	}
	return o.Board, true
}

// HasBoard returns a boolean if a field has been set.
func (o *TargetRecord) HasBoard() bool {
	if o != nil && !IsNil(o.Board) {
		return true
	}

	return false
}

// SetBoard gets a reference to the given int64 and assigns it to the Board field.
func (o *TargetRecord) SetBoard(v int64) {
	o.Board = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *TargetRecord) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *TargetRecord) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *TargetRecord) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *TargetRecord) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *TargetRecord) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *TargetRecord) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDcName returns the DcName field value if set, zero value otherwise.
func (o *TargetRecord) GetDcName() string {
	if o == nil || IsNil(o.DcName) {
		var ret string
		return ret
	}
	return *o.DcName
}

// GetDcNameOk returns a tuple with the DcName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetDcNameOk() (*string, bool) {
	if o == nil || IsNil(o.DcName) {
		return nil, false
	}
	return o.DcName, true
}

// HasDcName returns a boolean if a field has been set.
func (o *TargetRecord) HasDcName() bool {
	if o != nil && !IsNil(o.DcName) {
		return true
	}

	return false
}

// SetDcName gets a reference to the given string and assigns it to the DcName field.
func (o *TargetRecord) SetDcName(v string) {
	o.DcName = &v
}

// GetGatewayIps returns the GatewayIps field value if set, zero value otherwise.
func (o *TargetRecord) GetGatewayIps() string {
	if o == nil || IsNil(o.GatewayIps) {
		var ret string
		return ret
	}
	return *o.GatewayIps
}

// GetGatewayIpsOk returns a tuple with the GatewayIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetGatewayIpsOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayIps) {
		return nil, false
	}
	return o.GatewayIps, true
}

// HasGatewayIps returns a boolean if a field has been set.
func (o *TargetRecord) HasGatewayIps() bool {
	if o != nil && !IsNil(o.GatewayIps) {
		return true
	}

	return false
}

// SetGatewayIps gets a reference to the given string and assigns it to the GatewayIps field.
func (o *TargetRecord) SetGatewayIps(v string) {
	o.GatewayIps = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *TargetRecord) GetHost() HostNestview {
	if o == nil || IsNil(o.Host) {
		var ret HostNestview
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetHostOk() (*HostNestview, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *TargetRecord) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostNestview and assigns it to the Host field.
func (o *TargetRecord) SetHost(v HostNestview) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TargetRecord) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TargetRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TargetRecord) SetId(v int64) {
	o.Id = &v
}

// GetIqn returns the Iqn field value if set, zero value otherwise.
func (o *TargetRecord) GetIqn() string {
	if o == nil || IsNil(o.Iqn) {
		var ret string
		return ret
	}
	return *o.Iqn
}

// GetIqnOk returns a tuple with the Iqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetIqnOk() (*string, bool) {
	if o == nil || IsNil(o.Iqn) {
		return nil, false
	}
	return o.Iqn, true
}

// HasIqn returns a boolean if a field has been set.
func (o *TargetRecord) HasIqn() bool {
	if o != nil && !IsNil(o.Iqn) {
		return true
	}

	return false
}

// SetIqn gets a reference to the given string and assigns it to the Iqn field.
func (o *TargetRecord) SetIqn(v string) {
	o.Iqn = &v
}

// GetLogout returns the Logout field value if set, zero value otherwise.
func (o *TargetRecord) GetLogout() bool {
	if o == nil || IsNil(o.Logout) {
		var ret bool
		return ret
	}
	return *o.Logout
}

// GetLogoutOk returns a tuple with the Logout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetLogoutOk() (*bool, bool) {
	if o == nil || IsNil(o.Logout) {
		return nil, false
	}
	return o.Logout, true
}

// HasLogout returns a boolean if a field has been set.
func (o *TargetRecord) HasLogout() bool {
	if o != nil && !IsNil(o.Logout) {
		return true
	}

	return false
}

// SetLogout gets a reference to the given bool and assigns it to the Logout field.
func (o *TargetRecord) SetLogout(v bool) {
	o.Logout = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *TargetRecord) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *TargetRecord) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *TargetRecord) SetPort(v int64) {
	o.Port = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TargetRecord) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TargetRecord) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TargetRecord) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *TargetRecord) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *TargetRecord) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *TargetRecord) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetLunsInfo returns the LunsInfo field value if set, zero value otherwise.
func (o *TargetRecord) GetLunsInfo() []LunInfo {
	if o == nil || IsNil(o.LunsInfo) {
		var ret []LunInfo
		return ret
	}
	return o.LunsInfo
}

// GetLunsInfoOk returns a tuple with the LunsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetRecord) GetLunsInfoOk() ([]LunInfo, bool) {
	if o == nil || IsNil(o.LunsInfo) {
		return nil, false
	}
	return o.LunsInfo, true
}

// HasLunsInfo returns a boolean if a field has been set.
func (o *TargetRecord) HasLunsInfo() bool {
	if o != nil && !IsNil(o.LunsInfo) {
		return true
	}

	return false
}

// SetLunsInfo gets a reference to the given []LunInfo and assigns it to the LunsInfo field.
func (o *TargetRecord) SetLunsInfo(v []LunInfo) {
	o.LunsInfo = v
}

func (o TargetRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessPath) {
		toSerialize["access_path"] = o.AccessPath
	}
	if !IsNil(o.Board) {
		toSerialize["board"] = o.Board
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DcName) {
		toSerialize["dc_name"] = o.DcName
	}
	if !IsNil(o.GatewayIps) {
		toSerialize["gateway_ips"] = o.GatewayIps
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Iqn) {
		toSerialize["iqn"] = o.Iqn
	}
	if !IsNil(o.Logout) {
		toSerialize["logout"] = o.Logout
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.LunsInfo) {
		toSerialize["luns_info"] = o.LunsInfo
	}
	return toSerialize, nil
}

type NullableTargetRecord struct {
	value *TargetRecord
	isSet bool
}

func (v NullableTargetRecord) Get() *TargetRecord {
	return v.value
}

func (v *NullableTargetRecord) Set(val *TargetRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetRecord(val *TargetRecord) *NullableTargetRecord {
	return &NullableTargetRecord{value: val, isSet: true}
}

func (v NullableTargetRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


