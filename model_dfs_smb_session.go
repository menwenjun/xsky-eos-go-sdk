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

// checks if the DfsSMBSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsSMBSession{}

// DfsSMBSession DfsSMBSession defines model of dfs smb session
type DfsSMBSession struct {
	ClientIp *string `json:"client_ip,omitempty"`
	ClientPort *int64 `json:"client_port,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	ConnectedAt *time.Time `json:"connected_at,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DfsGateway *DfsGatewayNestview `json:"dfs_gateway,omitempty"`
	DfsSmbShare *DfsSMBShareNestview `json:"dfs_smb_share,omitempty"`
	Group *string `json:"group,omitempty"`
	Id *int64 `json:"id,omitempty"`
	ProtocolVersion *string `json:"protocol_version,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewDfsSMBSession instantiates a new DfsSMBSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsSMBSession() *DfsSMBSession {
	this := DfsSMBSession{}
	return &this
}

// NewDfsSMBSessionWithDefaults instantiates a new DfsSMBSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsSMBSessionWithDefaults() *DfsSMBSession {
	this := DfsSMBSession{}
	return &this
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *DfsSMBSession) GetClientIp() string {
	if o == nil || IsNil(o.ClientIp) {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetClientIpOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIp) {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *DfsSMBSession) HasClientIp() bool {
	if o != nil && !IsNil(o.ClientIp) {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *DfsSMBSession) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetClientPort returns the ClientPort field value if set, zero value otherwise.
func (o *DfsSMBSession) GetClientPort() int64 {
	if o == nil || IsNil(o.ClientPort) {
		var ret int64
		return ret
	}
	return *o.ClientPort
}

// GetClientPortOk returns a tuple with the ClientPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetClientPortOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientPort) {
		return nil, false
	}
	return o.ClientPort, true
}

// HasClientPort returns a boolean if a field has been set.
func (o *DfsSMBSession) HasClientPort() bool {
	if o != nil && !IsNil(o.ClientPort) {
		return true
	}

	return false
}

// SetClientPort gets a reference to the given int64 and assigns it to the ClientPort field.
func (o *DfsSMBSession) SetClientPort(v int64) {
	o.ClientPort = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DfsSMBSession) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DfsSMBSession) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DfsSMBSession) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetConnectedAt returns the ConnectedAt field value if set, zero value otherwise.
func (o *DfsSMBSession) GetConnectedAt() time.Time {
	if o == nil || IsNil(o.ConnectedAt) {
		var ret time.Time
		return ret
	}
	return *o.ConnectedAt
}

// GetConnectedAtOk returns a tuple with the ConnectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetConnectedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ConnectedAt) {
		return nil, false
	}
	return o.ConnectedAt, true
}

// HasConnectedAt returns a boolean if a field has been set.
func (o *DfsSMBSession) HasConnectedAt() bool {
	if o != nil && !IsNil(o.ConnectedAt) {
		return true
	}

	return false
}

// SetConnectedAt gets a reference to the given time.Time and assigns it to the ConnectedAt field.
func (o *DfsSMBSession) SetConnectedAt(v time.Time) {
	o.ConnectedAt = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsSMBSession) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsSMBSession) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsSMBSession) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDfsGateway returns the DfsGateway field value if set, zero value otherwise.
func (o *DfsSMBSession) GetDfsGateway() DfsGatewayNestview {
	if o == nil || IsNil(o.DfsGateway) {
		var ret DfsGatewayNestview
		return ret
	}
	return *o.DfsGateway
}

// GetDfsGatewayOk returns a tuple with the DfsGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetDfsGatewayOk() (*DfsGatewayNestview, bool) {
	if o == nil || IsNil(o.DfsGateway) {
		return nil, false
	}
	return o.DfsGateway, true
}

// HasDfsGateway returns a boolean if a field has been set.
func (o *DfsSMBSession) HasDfsGateway() bool {
	if o != nil && !IsNil(o.DfsGateway) {
		return true
	}

	return false
}

// SetDfsGateway gets a reference to the given DfsGatewayNestview and assigns it to the DfsGateway field.
func (o *DfsSMBSession) SetDfsGateway(v DfsGatewayNestview) {
	o.DfsGateway = &v
}

// GetDfsSmbShare returns the DfsSmbShare field value if set, zero value otherwise.
func (o *DfsSMBSession) GetDfsSmbShare() DfsSMBShareNestview {
	if o == nil || IsNil(o.DfsSmbShare) {
		var ret DfsSMBShareNestview
		return ret
	}
	return *o.DfsSmbShare
}

// GetDfsSmbShareOk returns a tuple with the DfsSmbShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetDfsSmbShareOk() (*DfsSMBShareNestview, bool) {
	if o == nil || IsNil(o.DfsSmbShare) {
		return nil, false
	}
	return o.DfsSmbShare, true
}

// HasDfsSmbShare returns a boolean if a field has been set.
func (o *DfsSMBSession) HasDfsSmbShare() bool {
	if o != nil && !IsNil(o.DfsSmbShare) {
		return true
	}

	return false
}

// SetDfsSmbShare gets a reference to the given DfsSMBShareNestview and assigns it to the DfsSmbShare field.
func (o *DfsSMBSession) SetDfsSmbShare(v DfsSMBShareNestview) {
	o.DfsSmbShare = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *DfsSMBSession) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *DfsSMBSession) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *DfsSMBSession) SetGroup(v string) {
	o.Group = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsSMBSession) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsSMBSession) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsSMBSession) SetId(v int64) {
	o.Id = &v
}

// GetProtocolVersion returns the ProtocolVersion field value if set, zero value otherwise.
func (o *DfsSMBSession) GetProtocolVersion() string {
	if o == nil || IsNil(o.ProtocolVersion) {
		var ret string
		return ret
	}
	return *o.ProtocolVersion
}

// GetProtocolVersionOk returns a tuple with the ProtocolVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetProtocolVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolVersion) {
		return nil, false
	}
	return o.ProtocolVersion, true
}

// HasProtocolVersion returns a boolean if a field has been set.
func (o *DfsSMBSession) HasProtocolVersion() bool {
	if o != nil && !IsNil(o.ProtocolVersion) {
		return true
	}

	return false
}

// SetProtocolVersion gets a reference to the given string and assigns it to the ProtocolVersion field.
func (o *DfsSMBSession) SetProtocolVersion(v string) {
	o.ProtocolVersion = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DfsSMBSession) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DfsSMBSession) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DfsSMBSession) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DfsSMBSession) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsSMBSession) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DfsSMBSession) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DfsSMBSession) SetUsername(v string) {
	o.Username = &v
}

func (o DfsSMBSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsSMBSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientIp) {
		toSerialize["client_ip"] = o.ClientIp
	}
	if !IsNil(o.ClientPort) {
		toSerialize["client_port"] = o.ClientPort
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.ConnectedAt) {
		toSerialize["connected_at"] = o.ConnectedAt
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DfsGateway) {
		toSerialize["dfs_gateway"] = o.DfsGateway
	}
	if !IsNil(o.DfsSmbShare) {
		toSerialize["dfs_smb_share"] = o.DfsSmbShare
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProtocolVersion) {
		toSerialize["protocol_version"] = o.ProtocolVersion
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableDfsSMBSession struct {
	value *DfsSMBSession
	isSet bool
}

func (v NullableDfsSMBSession) Get() *DfsSMBSession {
	return v.value
}

func (v *NullableDfsSMBSession) Set(val *DfsSMBSession) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsSMBSession) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsSMBSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsSMBSession(val *DfsSMBSession) *NullableDfsSMBSession {
	return &NullableDfsSMBSession{value: val, isSet: true}
}

func (v NullableDfsSMBSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsSMBSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


