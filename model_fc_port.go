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

// checks if the FCPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FCPort{}

// FCPort FCPort defines model of FC port @grpc-models-proto +X:model:etcd_lock;
type FCPort struct {
	ActionStatus *string `json:"action_status,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	ConnOptMode *string `json:"conn_opt_mode,omitempty"`
	ConnType *string `json:"conn_type,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DataRateMode *string `json:"data_rate_mode,omitempty"`
	Health *string `json:"health,omitempty"`
	Host *HostNestview `json:"host,omitempty"`
	Id *int64 `json:"id,omitempty"`
	LinkSpeed *string `json:"link_speed,omitempty"`
	LinkState *string `json:"link_state,omitempty"`
	MaxSpeed *string `json:"max_speed,omitempty"`
	PciAddress *string `json:"pci_address,omitempty"`
	PortId *string `json:"port_id,omitempty"`
	RecvBytes *int64 `json:"recv_bytes,omitempty"`
	RecvFrames *int64 `json:"recv_frames,omitempty"`
	RgHost *int64 `json:"rg_host,omitempty"`
	SendBytes *int64 `json:"send_bytes,omitempty"`
	SendFrames *int64 `json:"send_frames,omitempty"`
	Start *time.Time `json:"start,omitempty"`
	SupportedSpeeds []string `json:"supported_speeds,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Wwpn *string `json:"wwpn,omitempty"`
}

// NewFCPort instantiates a new FCPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFCPort() *FCPort {
	this := FCPort{}
	return &this
}

// NewFCPortWithDefaults instantiates a new FCPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFCPortWithDefaults() *FCPort {
	this := FCPort{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *FCPort) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *FCPort) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *FCPort) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *FCPort) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *FCPort) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *FCPort) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetConnOptMode returns the ConnOptMode field value if set, zero value otherwise.
func (o *FCPort) GetConnOptMode() string {
	if o == nil || IsNil(o.ConnOptMode) {
		var ret string
		return ret
	}
	return *o.ConnOptMode
}

// GetConnOptModeOk returns a tuple with the ConnOptMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetConnOptModeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnOptMode) {
		return nil, false
	}
	return o.ConnOptMode, true
}

// HasConnOptMode returns a boolean if a field has been set.
func (o *FCPort) HasConnOptMode() bool {
	if o != nil && !IsNil(o.ConnOptMode) {
		return true
	}

	return false
}

// SetConnOptMode gets a reference to the given string and assigns it to the ConnOptMode field.
func (o *FCPort) SetConnOptMode(v string) {
	o.ConnOptMode = &v
}

// GetConnType returns the ConnType field value if set, zero value otherwise.
func (o *FCPort) GetConnType() string {
	if o == nil || IsNil(o.ConnType) {
		var ret string
		return ret
	}
	return *o.ConnType
}

// GetConnTypeOk returns a tuple with the ConnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetConnTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnType) {
		return nil, false
	}
	return o.ConnType, true
}

// HasConnType returns a boolean if a field has been set.
func (o *FCPort) HasConnType() bool {
	if o != nil && !IsNil(o.ConnType) {
		return true
	}

	return false
}

// SetConnType gets a reference to the given string and assigns it to the ConnType field.
func (o *FCPort) SetConnType(v string) {
	o.ConnType = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *FCPort) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *FCPort) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *FCPort) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDataRateMode returns the DataRateMode field value if set, zero value otherwise.
func (o *FCPort) GetDataRateMode() string {
	if o == nil || IsNil(o.DataRateMode) {
		var ret string
		return ret
	}
	return *o.DataRateMode
}

// GetDataRateModeOk returns a tuple with the DataRateMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetDataRateModeOk() (*string, bool) {
	if o == nil || IsNil(o.DataRateMode) {
		return nil, false
	}
	return o.DataRateMode, true
}

// HasDataRateMode returns a boolean if a field has been set.
func (o *FCPort) HasDataRateMode() bool {
	if o != nil && !IsNil(o.DataRateMode) {
		return true
	}

	return false
}

// SetDataRateMode gets a reference to the given string and assigns it to the DataRateMode field.
func (o *FCPort) SetDataRateMode(v string) {
	o.DataRateMode = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *FCPort) GetHealth() string {
	if o == nil || IsNil(o.Health) {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetHealthOk() (*string, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *FCPort) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *FCPort) SetHealth(v string) {
	o.Health = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *FCPort) GetHost() HostNestview {
	if o == nil || IsNil(o.Host) {
		var ret HostNestview
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetHostOk() (*HostNestview, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *FCPort) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostNestview and assigns it to the Host field.
func (o *FCPort) SetHost(v HostNestview) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FCPort) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FCPort) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FCPort) SetId(v int64) {
	o.Id = &v
}

// GetLinkSpeed returns the LinkSpeed field value if set, zero value otherwise.
func (o *FCPort) GetLinkSpeed() string {
	if o == nil || IsNil(o.LinkSpeed) {
		var ret string
		return ret
	}
	return *o.LinkSpeed
}

// GetLinkSpeedOk returns a tuple with the LinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetLinkSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.LinkSpeed) {
		return nil, false
	}
	return o.LinkSpeed, true
}

// HasLinkSpeed returns a boolean if a field has been set.
func (o *FCPort) HasLinkSpeed() bool {
	if o != nil && !IsNil(o.LinkSpeed) {
		return true
	}

	return false
}

// SetLinkSpeed gets a reference to the given string and assigns it to the LinkSpeed field.
func (o *FCPort) SetLinkSpeed(v string) {
	o.LinkSpeed = &v
}

// GetLinkState returns the LinkState field value if set, zero value otherwise.
func (o *FCPort) GetLinkState() string {
	if o == nil || IsNil(o.LinkState) {
		var ret string
		return ret
	}
	return *o.LinkState
}

// GetLinkStateOk returns a tuple with the LinkState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetLinkStateOk() (*string, bool) {
	if o == nil || IsNil(o.LinkState) {
		return nil, false
	}
	return o.LinkState, true
}

// HasLinkState returns a boolean if a field has been set.
func (o *FCPort) HasLinkState() bool {
	if o != nil && !IsNil(o.LinkState) {
		return true
	}

	return false
}

// SetLinkState gets a reference to the given string and assigns it to the LinkState field.
func (o *FCPort) SetLinkState(v string) {
	o.LinkState = &v
}

// GetMaxSpeed returns the MaxSpeed field value if set, zero value otherwise.
func (o *FCPort) GetMaxSpeed() string {
	if o == nil || IsNil(o.MaxSpeed) {
		var ret string
		return ret
	}
	return *o.MaxSpeed
}

// GetMaxSpeedOk returns a tuple with the MaxSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetMaxSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.MaxSpeed) {
		return nil, false
	}
	return o.MaxSpeed, true
}

// HasMaxSpeed returns a boolean if a field has been set.
func (o *FCPort) HasMaxSpeed() bool {
	if o != nil && !IsNil(o.MaxSpeed) {
		return true
	}

	return false
}

// SetMaxSpeed gets a reference to the given string and assigns it to the MaxSpeed field.
func (o *FCPort) SetMaxSpeed(v string) {
	o.MaxSpeed = &v
}

// GetPciAddress returns the PciAddress field value if set, zero value otherwise.
func (o *FCPort) GetPciAddress() string {
	if o == nil || IsNil(o.PciAddress) {
		var ret string
		return ret
	}
	return *o.PciAddress
}

// GetPciAddressOk returns a tuple with the PciAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetPciAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PciAddress) {
		return nil, false
	}
	return o.PciAddress, true
}

// HasPciAddress returns a boolean if a field has been set.
func (o *FCPort) HasPciAddress() bool {
	if o != nil && !IsNil(o.PciAddress) {
		return true
	}

	return false
}

// SetPciAddress gets a reference to the given string and assigns it to the PciAddress field.
func (o *FCPort) SetPciAddress(v string) {
	o.PciAddress = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *FCPort) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *FCPort) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *FCPort) SetPortId(v string) {
	o.PortId = &v
}

// GetRecvBytes returns the RecvBytes field value if set, zero value otherwise.
func (o *FCPort) GetRecvBytes() int64 {
	if o == nil || IsNil(o.RecvBytes) {
		var ret int64
		return ret
	}
	return *o.RecvBytes
}

// GetRecvBytesOk returns a tuple with the RecvBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetRecvBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.RecvBytes) {
		return nil, false
	}
	return o.RecvBytes, true
}

// HasRecvBytes returns a boolean if a field has been set.
func (o *FCPort) HasRecvBytes() bool {
	if o != nil && !IsNil(o.RecvBytes) {
		return true
	}

	return false
}

// SetRecvBytes gets a reference to the given int64 and assigns it to the RecvBytes field.
func (o *FCPort) SetRecvBytes(v int64) {
	o.RecvBytes = &v
}

// GetRecvFrames returns the RecvFrames field value if set, zero value otherwise.
func (o *FCPort) GetRecvFrames() int64 {
	if o == nil || IsNil(o.RecvFrames) {
		var ret int64
		return ret
	}
	return *o.RecvFrames
}

// GetRecvFramesOk returns a tuple with the RecvFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetRecvFramesOk() (*int64, bool) {
	if o == nil || IsNil(o.RecvFrames) {
		return nil, false
	}
	return o.RecvFrames, true
}

// HasRecvFrames returns a boolean if a field has been set.
func (o *FCPort) HasRecvFrames() bool {
	if o != nil && !IsNil(o.RecvFrames) {
		return true
	}

	return false
}

// SetRecvFrames gets a reference to the given int64 and assigns it to the RecvFrames field.
func (o *FCPort) SetRecvFrames(v int64) {
	o.RecvFrames = &v
}

// GetRgHost returns the RgHost field value if set, zero value otherwise.
func (o *FCPort) GetRgHost() int64 {
	if o == nil || IsNil(o.RgHost) {
		var ret int64
		return ret
	}
	return *o.RgHost
}

// GetRgHostOk returns a tuple with the RgHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetRgHostOk() (*int64, bool) {
	if o == nil || IsNil(o.RgHost) {
		return nil, false
	}
	return o.RgHost, true
}

// HasRgHost returns a boolean if a field has been set.
func (o *FCPort) HasRgHost() bool {
	if o != nil && !IsNil(o.RgHost) {
		return true
	}

	return false
}

// SetRgHost gets a reference to the given int64 and assigns it to the RgHost field.
func (o *FCPort) SetRgHost(v int64) {
	o.RgHost = &v
}

// GetSendBytes returns the SendBytes field value if set, zero value otherwise.
func (o *FCPort) GetSendBytes() int64 {
	if o == nil || IsNil(o.SendBytes) {
		var ret int64
		return ret
	}
	return *o.SendBytes
}

// GetSendBytesOk returns a tuple with the SendBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetSendBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.SendBytes) {
		return nil, false
	}
	return o.SendBytes, true
}

// HasSendBytes returns a boolean if a field has been set.
func (o *FCPort) HasSendBytes() bool {
	if o != nil && !IsNil(o.SendBytes) {
		return true
	}

	return false
}

// SetSendBytes gets a reference to the given int64 and assigns it to the SendBytes field.
func (o *FCPort) SetSendBytes(v int64) {
	o.SendBytes = &v
}

// GetSendFrames returns the SendFrames field value if set, zero value otherwise.
func (o *FCPort) GetSendFrames() int64 {
	if o == nil || IsNil(o.SendFrames) {
		var ret int64
		return ret
	}
	return *o.SendFrames
}

// GetSendFramesOk returns a tuple with the SendFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetSendFramesOk() (*int64, bool) {
	if o == nil || IsNil(o.SendFrames) {
		return nil, false
	}
	return o.SendFrames, true
}

// HasSendFrames returns a boolean if a field has been set.
func (o *FCPort) HasSendFrames() bool {
	if o != nil && !IsNil(o.SendFrames) {
		return true
	}

	return false
}

// SetSendFrames gets a reference to the given int64 and assigns it to the SendFrames field.
func (o *FCPort) SetSendFrames(v int64) {
	o.SendFrames = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *FCPort) GetStart() time.Time {
	if o == nil || IsNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *FCPort) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *FCPort) SetStart(v time.Time) {
	o.Start = &v
}

// GetSupportedSpeeds returns the SupportedSpeeds field value if set, zero value otherwise.
func (o *FCPort) GetSupportedSpeeds() []string {
	if o == nil || IsNil(o.SupportedSpeeds) {
		var ret []string
		return ret
	}
	return o.SupportedSpeeds
}

// GetSupportedSpeedsOk returns a tuple with the SupportedSpeeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetSupportedSpeedsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedSpeeds) {
		return nil, false
	}
	return o.SupportedSpeeds, true
}

// HasSupportedSpeeds returns a boolean if a field has been set.
func (o *FCPort) HasSupportedSpeeds() bool {
	if o != nil && !IsNil(o.SupportedSpeeds) {
		return true
	}

	return false
}

// SetSupportedSpeeds gets a reference to the given []string and assigns it to the SupportedSpeeds field.
func (o *FCPort) SetSupportedSpeeds(v []string) {
	o.SupportedSpeeds = v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *FCPort) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *FCPort) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *FCPort) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetWwpn returns the Wwpn field value if set, zero value otherwise.
func (o *FCPort) GetWwpn() string {
	if o == nil || IsNil(o.Wwpn) {
		var ret string
		return ret
	}
	return *o.Wwpn
}

// GetWwpnOk returns a tuple with the Wwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCPort) GetWwpnOk() (*string, bool) {
	if o == nil || IsNil(o.Wwpn) {
		return nil, false
	}
	return o.Wwpn, true
}

// HasWwpn returns a boolean if a field has been set.
func (o *FCPort) HasWwpn() bool {
	if o != nil && !IsNil(o.Wwpn) {
		return true
	}

	return false
}

// SetWwpn gets a reference to the given string and assigns it to the Wwpn field.
func (o *FCPort) SetWwpn(v string) {
	o.Wwpn = &v
}

func (o FCPort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FCPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.ConnOptMode) {
		toSerialize["conn_opt_mode"] = o.ConnOptMode
	}
	if !IsNil(o.ConnType) {
		toSerialize["conn_type"] = o.ConnType
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DataRateMode) {
		toSerialize["data_rate_mode"] = o.DataRateMode
	}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LinkSpeed) {
		toSerialize["link_speed"] = o.LinkSpeed
	}
	if !IsNil(o.LinkState) {
		toSerialize["link_state"] = o.LinkState
	}
	if !IsNil(o.MaxSpeed) {
		toSerialize["max_speed"] = o.MaxSpeed
	}
	if !IsNil(o.PciAddress) {
		toSerialize["pci_address"] = o.PciAddress
	}
	if !IsNil(o.PortId) {
		toSerialize["port_id"] = o.PortId
	}
	if !IsNil(o.RecvBytes) {
		toSerialize["recv_bytes"] = o.RecvBytes
	}
	if !IsNil(o.RecvFrames) {
		toSerialize["recv_frames"] = o.RecvFrames
	}
	if !IsNil(o.RgHost) {
		toSerialize["rg_host"] = o.RgHost
	}
	if !IsNil(o.SendBytes) {
		toSerialize["send_bytes"] = o.SendBytes
	}
	if !IsNil(o.SendFrames) {
		toSerialize["send_frames"] = o.SendFrames
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.SupportedSpeeds) {
		toSerialize["supported_speeds"] = o.SupportedSpeeds
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Wwpn) {
		toSerialize["wwpn"] = o.Wwpn
	}
	return toSerialize, nil
}

type NullableFCPort struct {
	value *FCPort
	isSet bool
}

func (v NullableFCPort) Get() *FCPort {
	return v.value
}

func (v *NullableFCPort) Set(val *FCPort) {
	v.value = val
	v.isSet = true
}

func (v NullableFCPort) IsSet() bool {
	return v.isSet
}

func (v *NullableFCPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFCPort(val *FCPort) *NullableFCPort {
	return &NullableFCPort{value: val, isSet: true}
}

func (v NullableFCPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFCPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


