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

// checks if the RBDClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RBDClient{}

// RBDClient RBDClient defines the rbd client +X:model:generate;
type RBDClient struct {
	AdminIp *string `json:"admin_ip,omitempty"`
	Cluster *Cluster `json:"cluster,omitempty"`
	CpuArch *string `json:"cpu_arch,omitempty"`
	CpuCoreNum *int64 `json:"cpu_core_num,omitempty"`
	CpuModel *string `json:"cpu_model,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Id *int64 `json:"id,omitempty"`
	MemoryKbyte *int64 `json:"memory_kbyte,omitempty"`
	Name *string `json:"name,omitempty"`
	Os *string `json:"os,omitempty"`
	Port *int32 `json:"port,omitempty"`
	PublicIp *string `json:"public_ip,omitempty"`
	PublicNetwork *string `json:"public_network,omitempty"`
	Status *string `json:"status,omitempty"`
	Token *string `json:"token,omitempty"`
	Type *string `json:"type,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	Version *string `json:"version,omitempty"`
	VolumeNum *int64 `json:"volume_num,omitempty"`
	XdcStatus *string `json:"xdc_status,omitempty"`
}

// NewRBDClient instantiates a new RBDClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRBDClient() *RBDClient {
	this := RBDClient{}
	return &this
}

// NewRBDClientWithDefaults instantiates a new RBDClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRBDClientWithDefaults() *RBDClient {
	this := RBDClient{}
	return &this
}

// GetAdminIp returns the AdminIp field value if set, zero value otherwise.
func (o *RBDClient) GetAdminIp() string {
	if o == nil || IsNil(o.AdminIp) {
		var ret string
		return ret
	}
	return *o.AdminIp
}

// GetAdminIpOk returns a tuple with the AdminIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetAdminIpOk() (*string, bool) {
	if o == nil || IsNil(o.AdminIp) {
		return nil, false
	}
	return o.AdminIp, true
}

// HasAdminIp returns a boolean if a field has been set.
func (o *RBDClient) HasAdminIp() bool {
	if o != nil && !IsNil(o.AdminIp) {
		return true
	}

	return false
}

// SetAdminIp gets a reference to the given string and assigns it to the AdminIp field.
func (o *RBDClient) SetAdminIp(v string) {
	o.AdminIp = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *RBDClient) GetCluster() Cluster {
	if o == nil || IsNil(o.Cluster) {
		var ret Cluster
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetClusterOk() (*Cluster, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *RBDClient) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given Cluster and assigns it to the Cluster field.
func (o *RBDClient) SetCluster(v Cluster) {
	o.Cluster = &v
}

// GetCpuArch returns the CpuArch field value if set, zero value otherwise.
func (o *RBDClient) GetCpuArch() string {
	if o == nil || IsNil(o.CpuArch) {
		var ret string
		return ret
	}
	return *o.CpuArch
}

// GetCpuArchOk returns a tuple with the CpuArch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetCpuArchOk() (*string, bool) {
	if o == nil || IsNil(o.CpuArch) {
		return nil, false
	}
	return o.CpuArch, true
}

// HasCpuArch returns a boolean if a field has been set.
func (o *RBDClient) HasCpuArch() bool {
	if o != nil && !IsNil(o.CpuArch) {
		return true
	}

	return false
}

// SetCpuArch gets a reference to the given string and assigns it to the CpuArch field.
func (o *RBDClient) SetCpuArch(v string) {
	o.CpuArch = &v
}

// GetCpuCoreNum returns the CpuCoreNum field value if set, zero value otherwise.
func (o *RBDClient) GetCpuCoreNum() int64 {
	if o == nil || IsNil(o.CpuCoreNum) {
		var ret int64
		return ret
	}
	return *o.CpuCoreNum
}

// GetCpuCoreNumOk returns a tuple with the CpuCoreNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetCpuCoreNumOk() (*int64, bool) {
	if o == nil || IsNil(o.CpuCoreNum) {
		return nil, false
	}
	return o.CpuCoreNum, true
}

// HasCpuCoreNum returns a boolean if a field has been set.
func (o *RBDClient) HasCpuCoreNum() bool {
	if o != nil && !IsNil(o.CpuCoreNum) {
		return true
	}

	return false
}

// SetCpuCoreNum gets a reference to the given int64 and assigns it to the CpuCoreNum field.
func (o *RBDClient) SetCpuCoreNum(v int64) {
	o.CpuCoreNum = &v
}

// GetCpuModel returns the CpuModel field value if set, zero value otherwise.
func (o *RBDClient) GetCpuModel() string {
	if o == nil || IsNil(o.CpuModel) {
		var ret string
		return ret
	}
	return *o.CpuModel
}

// GetCpuModelOk returns a tuple with the CpuModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetCpuModelOk() (*string, bool) {
	if o == nil || IsNil(o.CpuModel) {
		return nil, false
	}
	return o.CpuModel, true
}

// HasCpuModel returns a boolean if a field has been set.
func (o *RBDClient) HasCpuModel() bool {
	if o != nil && !IsNil(o.CpuModel) {
		return true
	}

	return false
}

// SetCpuModel gets a reference to the given string and assigns it to the CpuModel field.
func (o *RBDClient) SetCpuModel(v string) {
	o.CpuModel = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *RBDClient) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *RBDClient) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *RBDClient) SetCreate(v time.Time) {
	o.Create = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RBDClient) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RBDClient) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RBDClient) SetId(v int64) {
	o.Id = &v
}

// GetMemoryKbyte returns the MemoryKbyte field value if set, zero value otherwise.
func (o *RBDClient) GetMemoryKbyte() int64 {
	if o == nil || IsNil(o.MemoryKbyte) {
		var ret int64
		return ret
	}
	return *o.MemoryKbyte
}

// GetMemoryKbyteOk returns a tuple with the MemoryKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetMemoryKbyteOk() (*int64, bool) {
	if o == nil || IsNil(o.MemoryKbyte) {
		return nil, false
	}
	return o.MemoryKbyte, true
}

// HasMemoryKbyte returns a boolean if a field has been set.
func (o *RBDClient) HasMemoryKbyte() bool {
	if o != nil && !IsNil(o.MemoryKbyte) {
		return true
	}

	return false
}

// SetMemoryKbyte gets a reference to the given int64 and assigns it to the MemoryKbyte field.
func (o *RBDClient) SetMemoryKbyte(v int64) {
	o.MemoryKbyte = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RBDClient) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RBDClient) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RBDClient) SetName(v string) {
	o.Name = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *RBDClient) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *RBDClient) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *RBDClient) SetOs(v string) {
	o.Os = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RBDClient) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RBDClient) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *RBDClient) SetPort(v int32) {
	o.Port = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *RBDClient) GetPublicIp() string {
	if o == nil || IsNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetPublicIpOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *RBDClient) HasPublicIp() bool {
	if o != nil && !IsNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *RBDClient) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPublicNetwork returns the PublicNetwork field value if set, zero value otherwise.
func (o *RBDClient) GetPublicNetwork() string {
	if o == nil || IsNil(o.PublicNetwork) {
		var ret string
		return ret
	}
	return *o.PublicNetwork
}

// GetPublicNetworkOk returns a tuple with the PublicNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetPublicNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.PublicNetwork) {
		return nil, false
	}
	return o.PublicNetwork, true
}

// HasPublicNetwork returns a boolean if a field has been set.
func (o *RBDClient) HasPublicNetwork() bool {
	if o != nil && !IsNil(o.PublicNetwork) {
		return true
	}

	return false
}

// SetPublicNetwork gets a reference to the given string and assigns it to the PublicNetwork field.
func (o *RBDClient) SetPublicNetwork(v string) {
	o.PublicNetwork = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RBDClient) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RBDClient) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RBDClient) SetStatus(v string) {
	o.Status = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RBDClient) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RBDClient) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RBDClient) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RBDClient) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RBDClient) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RBDClient) SetType(v string) {
	o.Type = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *RBDClient) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *RBDClient) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *RBDClient) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *RBDClient) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *RBDClient) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *RBDClient) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RBDClient) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RBDClient) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RBDClient) SetVersion(v string) {
	o.Version = &v
}

// GetVolumeNum returns the VolumeNum field value if set, zero value otherwise.
func (o *RBDClient) GetVolumeNum() int64 {
	if o == nil || IsNil(o.VolumeNum) {
		var ret int64
		return ret
	}
	return *o.VolumeNum
}

// GetVolumeNumOk returns a tuple with the VolumeNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetVolumeNumOk() (*int64, bool) {
	if o == nil || IsNil(o.VolumeNum) {
		return nil, false
	}
	return o.VolumeNum, true
}

// HasVolumeNum returns a boolean if a field has been set.
func (o *RBDClient) HasVolumeNum() bool {
	if o != nil && !IsNil(o.VolumeNum) {
		return true
	}

	return false
}

// SetVolumeNum gets a reference to the given int64 and assigns it to the VolumeNum field.
func (o *RBDClient) SetVolumeNum(v int64) {
	o.VolumeNum = &v
}

// GetXdcStatus returns the XdcStatus field value if set, zero value otherwise.
func (o *RBDClient) GetXdcStatus() string {
	if o == nil || IsNil(o.XdcStatus) {
		var ret string
		return ret
	}
	return *o.XdcStatus
}

// GetXdcStatusOk returns a tuple with the XdcStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RBDClient) GetXdcStatusOk() (*string, bool) {
	if o == nil || IsNil(o.XdcStatus) {
		return nil, false
	}
	return o.XdcStatus, true
}

// HasXdcStatus returns a boolean if a field has been set.
func (o *RBDClient) HasXdcStatus() bool {
	if o != nil && !IsNil(o.XdcStatus) {
		return true
	}

	return false
}

// SetXdcStatus gets a reference to the given string and assigns it to the XdcStatus field.
func (o *RBDClient) SetXdcStatus(v string) {
	o.XdcStatus = &v
}

func (o RBDClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RBDClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminIp) {
		toSerialize["admin_ip"] = o.AdminIp
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.CpuArch) {
		toSerialize["cpu_arch"] = o.CpuArch
	}
	if !IsNil(o.CpuCoreNum) {
		toSerialize["cpu_core_num"] = o.CpuCoreNum
	}
	if !IsNil(o.CpuModel) {
		toSerialize["cpu_model"] = o.CpuModel
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MemoryKbyte) {
		toSerialize["memory_kbyte"] = o.MemoryKbyte
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.PublicIp) {
		toSerialize["public_ip"] = o.PublicIp
	}
	if !IsNil(o.PublicNetwork) {
		toSerialize["public_network"] = o.PublicNetwork
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.VolumeNum) {
		toSerialize["volume_num"] = o.VolumeNum
	}
	if !IsNil(o.XdcStatus) {
		toSerialize["xdc_status"] = o.XdcStatus
	}
	return toSerialize, nil
}

type NullableRBDClient struct {
	value *RBDClient
	isSet bool
}

func (v NullableRBDClient) Get() *RBDClient {
	return v.value
}

func (v *NullableRBDClient) Set(val *RBDClient) {
	v.value = val
	v.isSet = true
}

func (v NullableRBDClient) IsSet() bool {
	return v.isSet
}

func (v *NullableRBDClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRBDClient(val *RBDClient) *NullableRBDClient {
	return &NullableRBDClient{value: val, isSet: true}
}

func (v NullableRBDClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRBDClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


