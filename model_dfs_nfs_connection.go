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

// checks if the DfsNFSConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsNFSConnection{}

// DfsNFSConnection DfsNFSConnection defines model of dfs nfs connection
type DfsNFSConnection struct {
	ClientIp *string `json:"client_ip,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DfsGateway *DfsGatewayNestview `json:"dfs_gateway,omitempty"`
	DfsNfsShare *DfsNFSShareNestview `json:"dfs_nfs_share,omitempty"`
	Id *int64 `json:"id,omitempty"`
	NfsVersion *string `json:"nfs_version,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewDfsNFSConnection instantiates a new DfsNFSConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsNFSConnection() *DfsNFSConnection {
	this := DfsNFSConnection{}
	return &this
}

// NewDfsNFSConnectionWithDefaults instantiates a new DfsNFSConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsNFSConnectionWithDefaults() *DfsNFSConnection {
	this := DfsNFSConnection{}
	return &this
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *DfsNFSConnection) GetClientIp() string {
	if o == nil || IsNil(o.ClientIp) {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsNFSConnection) GetClientIpOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIp) {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *DfsNFSConnection) HasClientIp() bool {
	if o != nil && !IsNil(o.ClientIp) {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *DfsNFSConnection) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DfsNFSConnection) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsNFSConnection) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DfsNFSConnection) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DfsNFSConnection) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsNFSConnection) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsNFSConnection) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsNFSConnection) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsNFSConnection) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDfsGateway returns the DfsGateway field value if set, zero value otherwise.
func (o *DfsNFSConnection) GetDfsGateway() DfsGatewayNestview {
	if o == nil || IsNil(o.DfsGateway) {
		var ret DfsGatewayNestview
		return ret
	}
	return *o.DfsGateway
}

// GetDfsGatewayOk returns a tuple with the DfsGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsNFSConnection) GetDfsGatewayOk() (*DfsGatewayNestview, bool) {
	if o == nil || IsNil(o.DfsGateway) {
		return nil, false
	}
	return o.DfsGateway, true
}

// HasDfsGateway returns a boolean if a field has been set.
func (o *DfsNFSConnection) HasDfsGateway() bool {
	if o != nil && !IsNil(o.DfsGateway) {
		return true
	}

	return false
}

// SetDfsGateway gets a reference to the given DfsGatewayNestview and assigns it to the DfsGateway field.
func (o *DfsNFSConnection) SetDfsGateway(v DfsGatewayNestview) {
	o.DfsGateway = &v
}

// GetDfsNfsShare returns the DfsNfsShare field value if set, zero value otherwise.
func (o *DfsNFSConnection) GetDfsNfsShare() DfsNFSShareNestview {
	if o == nil || IsNil(o.DfsNfsShare) {
		var ret DfsNFSShareNestview
		return ret
	}
	return *o.DfsNfsShare
}

// GetDfsNfsShareOk returns a tuple with the DfsNfsShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsNFSConnection) GetDfsNfsShareOk() (*DfsNFSShareNestview, bool) {
	if o == nil || IsNil(o.DfsNfsShare) {
		return nil, false
	}
	return o.DfsNfsShare, true
}

// HasDfsNfsShare returns a boolean if a field has been set.
func (o *DfsNFSConnection) HasDfsNfsShare() bool {
	if o != nil && !IsNil(o.DfsNfsShare) {
		return true
	}

	return false
}

// SetDfsNfsShare gets a reference to the given DfsNFSShareNestview and assigns it to the DfsNfsShare field.
func (o *DfsNFSConnection) SetDfsNfsShare(v DfsNFSShareNestview) {
	o.DfsNfsShare = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsNFSConnection) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsNFSConnection) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsNFSConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsNFSConnection) SetId(v int64) {
	o.Id = &v
}

// GetNfsVersion returns the NfsVersion field value if set, zero value otherwise.
func (o *DfsNFSConnection) GetNfsVersion() string {
	if o == nil || IsNil(o.NfsVersion) {
		var ret string
		return ret
	}
	return *o.NfsVersion
}

// GetNfsVersionOk returns a tuple with the NfsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsNFSConnection) GetNfsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.NfsVersion) {
		return nil, false
	}
	return o.NfsVersion, true
}

// HasNfsVersion returns a boolean if a field has been set.
func (o *DfsNFSConnection) HasNfsVersion() bool {
	if o != nil && !IsNil(o.NfsVersion) {
		return true
	}

	return false
}

// SetNfsVersion gets a reference to the given string and assigns it to the NfsVersion field.
func (o *DfsNFSConnection) SetNfsVersion(v string) {
	o.NfsVersion = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DfsNFSConnection) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsNFSConnection) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DfsNFSConnection) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DfsNFSConnection) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o DfsNFSConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsNFSConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientIp) {
		toSerialize["client_ip"] = o.ClientIp
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.DfsGateway) {
		toSerialize["dfs_gateway"] = o.DfsGateway
	}
	if !IsNil(o.DfsNfsShare) {
		toSerialize["dfs_nfs_share"] = o.DfsNfsShare
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NfsVersion) {
		toSerialize["nfs_version"] = o.NfsVersion
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableDfsNFSConnection struct {
	value *DfsNFSConnection
	isSet bool
}

func (v NullableDfsNFSConnection) Get() *DfsNFSConnection {
	return v.value
}

func (v *NullableDfsNFSConnection) Set(val *DfsNFSConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsNFSConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsNFSConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsNFSConnection(val *DfsNFSConnection) *NullableDfsNFSConnection {
	return &NullableDfsNFSConnection{value: val, isSet: true}
}

func (v NullableDfsNFSConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsNFSConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


