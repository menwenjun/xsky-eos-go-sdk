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

// checks if the NFSGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NFSGateway{}

// NFSGateway NFSGateway defines nfs gateway
type NFSGateway struct {
	ActionStatus *string `json:"action_status,omitempty"`
	BucketNum *int64 `json:"bucket_num,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Description *string `json:"description,omitempty"`
	GatewayIp *string `json:"gateway_ip,omitempty"`
	GatewayName *string `json:"gateway_name,omitempty"`
	Host *HostNestview `json:"host,omitempty"`
	Id *int64 `json:"id,omitempty"`
	MountClients *string `json:"mount_clients,omitempty"`
	MountNum *int64 `json:"mount_num,omitempty"`
	Name *string `json:"name,omitempty"`
	OspZoneId *int64 `json:"osp_zone_id,omitempty"`
	PlatformIp *string `json:"platform_ip,omitempty"`
	Port *int64 `json:"port,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewNFSGateway instantiates a new NFSGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFSGateway() *NFSGateway {
	this := NFSGateway{}
	return &this
}

// NewNFSGatewayWithDefaults instantiates a new NFSGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFSGatewayWithDefaults() *NFSGateway {
	this := NFSGateway{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *NFSGateway) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *NFSGateway) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *NFSGateway) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetBucketNum returns the BucketNum field value if set, zero value otherwise.
func (o *NFSGateway) GetBucketNum() int64 {
	if o == nil || IsNil(o.BucketNum) {
		var ret int64
		return ret
	}
	return *o.BucketNum
}

// GetBucketNumOk returns a tuple with the BucketNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetBucketNumOk() (*int64, bool) {
	if o == nil || IsNil(o.BucketNum) {
		return nil, false
	}
	return o.BucketNum, true
}

// HasBucketNum returns a boolean if a field has been set.
func (o *NFSGateway) HasBucketNum() bool {
	if o != nil && !IsNil(o.BucketNum) {
		return true
	}

	return false
}

// SetBucketNum gets a reference to the given int64 and assigns it to the BucketNum field.
func (o *NFSGateway) SetBucketNum(v int64) {
	o.BucketNum = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *NFSGateway) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *NFSGateway) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *NFSGateway) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *NFSGateway) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *NFSGateway) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *NFSGateway) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NFSGateway) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NFSGateway) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NFSGateway) SetDescription(v string) {
	o.Description = &v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *NFSGateway) GetGatewayIp() string {
	if o == nil || IsNil(o.GatewayIp) {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetGatewayIpOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayIp) {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *NFSGateway) HasGatewayIp() bool {
	if o != nil && !IsNil(o.GatewayIp) {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *NFSGateway) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetGatewayName returns the GatewayName field value if set, zero value otherwise.
func (o *NFSGateway) GetGatewayName() string {
	if o == nil || IsNil(o.GatewayName) {
		var ret string
		return ret
	}
	return *o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetGatewayNameOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayName) {
		return nil, false
	}
	return o.GatewayName, true
}

// HasGatewayName returns a boolean if a field has been set.
func (o *NFSGateway) HasGatewayName() bool {
	if o != nil && !IsNil(o.GatewayName) {
		return true
	}

	return false
}

// SetGatewayName gets a reference to the given string and assigns it to the GatewayName field.
func (o *NFSGateway) SetGatewayName(v string) {
	o.GatewayName = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NFSGateway) GetHost() HostNestview {
	if o == nil || IsNil(o.Host) {
		var ret HostNestview
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetHostOk() (*HostNestview, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NFSGateway) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostNestview and assigns it to the Host field.
func (o *NFSGateway) SetHost(v HostNestview) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NFSGateway) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NFSGateway) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *NFSGateway) SetId(v int64) {
	o.Id = &v
}

// GetMountClients returns the MountClients field value if set, zero value otherwise.
func (o *NFSGateway) GetMountClients() string {
	if o == nil || IsNil(o.MountClients) {
		var ret string
		return ret
	}
	return *o.MountClients
}

// GetMountClientsOk returns a tuple with the MountClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetMountClientsOk() (*string, bool) {
	if o == nil || IsNil(o.MountClients) {
		return nil, false
	}
	return o.MountClients, true
}

// HasMountClients returns a boolean if a field has been set.
func (o *NFSGateway) HasMountClients() bool {
	if o != nil && !IsNil(o.MountClients) {
		return true
	}

	return false
}

// SetMountClients gets a reference to the given string and assigns it to the MountClients field.
func (o *NFSGateway) SetMountClients(v string) {
	o.MountClients = &v
}

// GetMountNum returns the MountNum field value if set, zero value otherwise.
func (o *NFSGateway) GetMountNum() int64 {
	if o == nil || IsNil(o.MountNum) {
		var ret int64
		return ret
	}
	return *o.MountNum
}

// GetMountNumOk returns a tuple with the MountNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetMountNumOk() (*int64, bool) {
	if o == nil || IsNil(o.MountNum) {
		return nil, false
	}
	return o.MountNum, true
}

// HasMountNum returns a boolean if a field has been set.
func (o *NFSGateway) HasMountNum() bool {
	if o != nil && !IsNil(o.MountNum) {
		return true
	}

	return false
}

// SetMountNum gets a reference to the given int64 and assigns it to the MountNum field.
func (o *NFSGateway) SetMountNum(v int64) {
	o.MountNum = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NFSGateway) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NFSGateway) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NFSGateway) SetName(v string) {
	o.Name = &v
}

// GetOspZoneId returns the OspZoneId field value if set, zero value otherwise.
func (o *NFSGateway) GetOspZoneId() int64 {
	if o == nil || IsNil(o.OspZoneId) {
		var ret int64
		return ret
	}
	return *o.OspZoneId
}

// GetOspZoneIdOk returns a tuple with the OspZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetOspZoneIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OspZoneId) {
		return nil, false
	}
	return o.OspZoneId, true
}

// HasOspZoneId returns a boolean if a field has been set.
func (o *NFSGateway) HasOspZoneId() bool {
	if o != nil && !IsNil(o.OspZoneId) {
		return true
	}

	return false
}

// SetOspZoneId gets a reference to the given int64 and assigns it to the OspZoneId field.
func (o *NFSGateway) SetOspZoneId(v int64) {
	o.OspZoneId = &v
}

// GetPlatformIp returns the PlatformIp field value if set, zero value otherwise.
func (o *NFSGateway) GetPlatformIp() string {
	if o == nil || IsNil(o.PlatformIp) {
		var ret string
		return ret
	}
	return *o.PlatformIp
}

// GetPlatformIpOk returns a tuple with the PlatformIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetPlatformIpOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformIp) {
		return nil, false
	}
	return o.PlatformIp, true
}

// HasPlatformIp returns a boolean if a field has been set.
func (o *NFSGateway) HasPlatformIp() bool {
	if o != nil && !IsNil(o.PlatformIp) {
		return true
	}

	return false
}

// SetPlatformIp gets a reference to the given string and assigns it to the PlatformIp field.
func (o *NFSGateway) SetPlatformIp(v string) {
	o.PlatformIp = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NFSGateway) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NFSGateway) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *NFSGateway) SetPort(v int64) {
	o.Port = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NFSGateway) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NFSGateway) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NFSGateway) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *NFSGateway) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFSGateway) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *NFSGateway) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *NFSGateway) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o NFSGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NFSGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.BucketNum) {
		toSerialize["bucket_num"] = o.BucketNum
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GatewayIp) {
		toSerialize["gateway_ip"] = o.GatewayIp
	}
	if !IsNil(o.GatewayName) {
		toSerialize["gateway_name"] = o.GatewayName
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MountClients) {
		toSerialize["mount_clients"] = o.MountClients
	}
	if !IsNil(o.MountNum) {
		toSerialize["mount_num"] = o.MountNum
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OspZoneId) {
		toSerialize["osp_zone_id"] = o.OspZoneId
	}
	if !IsNil(o.PlatformIp) {
		toSerialize["platform_ip"] = o.PlatformIp
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
	return toSerialize, nil
}

type NullableNFSGateway struct {
	value *NFSGateway
	isSet bool
}

func (v NullableNFSGateway) Get() *NFSGateway {
	return v.value
}

func (v *NullableNFSGateway) Set(val *NFSGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableNFSGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableNFSGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFSGateway(val *NFSGateway) *NullableNFSGateway {
	return &NullableNFSGateway{value: val, isSet: true}
}

func (v NullableNFSGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFSGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


