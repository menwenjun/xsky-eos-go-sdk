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

// checks if the OspGatewayRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OspGatewayRecord{}

// OspGatewayRecord OspGatewayRecord combine OspGateway and OspGatewayStat as API response
type OspGatewayRecord struct {
	ActionStatus *string `json:"action_status,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Description *string `json:"description,omitempty"`
	Etag *string `json:"etag,omitempty"`
	GatewayId *int64 `json:"gateway_id,omitempty"`
	GatewayIp *string `json:"gateway_ip,omitempty"`
	GatewayName *string `json:"gateway_name,omitempty"`
	Host *HostNestview `json:"host,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	PlatformIp *string `json:"platform_ip,omitempty"`
	Role *string `json:"role,omitempty"`
	S3Port *int64 `json:"s3_port,omitempty"`
	Status *string `json:"status,omitempty"`
	XmsPort *int64 `json:"xms_port,omitempty"`
	ZoneId *int64 `json:"zone_id,omitempty"`
	Samples []OspGatewayStat `json:"samples,omitempty"`
}

// NewOspGatewayRecord instantiates a new OspGatewayRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOspGatewayRecord() *OspGatewayRecord {
	this := OspGatewayRecord{}
	return &this
}

// NewOspGatewayRecordWithDefaults instantiates a new OspGatewayRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOspGatewayRecordWithDefaults() *OspGatewayRecord {
	this := OspGatewayRecord{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *OspGatewayRecord) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *OspGatewayRecord) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *OspGatewayRecord) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OspGatewayRecord) SetDescription(v string) {
	o.Description = &v
}

// GetEtag returns the Etag field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetEtag() string {
	if o == nil || IsNil(o.Etag) {
		var ret string
		return ret
	}
	return *o.Etag
}

// GetEtagOk returns a tuple with the Etag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetEtagOk() (*string, bool) {
	if o == nil || IsNil(o.Etag) {
		return nil, false
	}
	return o.Etag, true
}

// HasEtag returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasEtag() bool {
	if o != nil && !IsNil(o.Etag) {
		return true
	}

	return false
}

// SetEtag gets a reference to the given string and assigns it to the Etag field.
func (o *OspGatewayRecord) SetEtag(v string) {
	o.Etag = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *OspGatewayRecord) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetGatewayIp() string {
	if o == nil || IsNil(o.GatewayIp) {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetGatewayIpOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayIp) {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasGatewayIp() bool {
	if o != nil && !IsNil(o.GatewayIp) {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *OspGatewayRecord) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetGatewayName returns the GatewayName field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetGatewayName() string {
	if o == nil || IsNil(o.GatewayName) {
		var ret string
		return ret
	}
	return *o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetGatewayNameOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayName) {
		return nil, false
	}
	return o.GatewayName, true
}

// HasGatewayName returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasGatewayName() bool {
	if o != nil && !IsNil(o.GatewayName) {
		return true
	}

	return false
}

// SetGatewayName gets a reference to the given string and assigns it to the GatewayName field.
func (o *OspGatewayRecord) SetGatewayName(v string) {
	o.GatewayName = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetHost() HostNestview {
	if o == nil || IsNil(o.Host) {
		var ret HostNestview
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetHostOk() (*HostNestview, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostNestview and assigns it to the Host field.
func (o *OspGatewayRecord) SetHost(v HostNestview) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OspGatewayRecord) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OspGatewayRecord) SetName(v string) {
	o.Name = &v
}

// GetPlatformIp returns the PlatformIp field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetPlatformIp() string {
	if o == nil || IsNil(o.PlatformIp) {
		var ret string
		return ret
	}
	return *o.PlatformIp
}

// GetPlatformIpOk returns a tuple with the PlatformIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetPlatformIpOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformIp) {
		return nil, false
	}
	return o.PlatformIp, true
}

// HasPlatformIp returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasPlatformIp() bool {
	if o != nil && !IsNil(o.PlatformIp) {
		return true
	}

	return false
}

// SetPlatformIp gets a reference to the given string and assigns it to the PlatformIp field.
func (o *OspGatewayRecord) SetPlatformIp(v string) {
	o.PlatformIp = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OspGatewayRecord) SetRole(v string) {
	o.Role = &v
}

// GetS3Port returns the S3Port field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetS3Port() int64 {
	if o == nil || IsNil(o.S3Port) {
		var ret int64
		return ret
	}
	return *o.S3Port
}

// GetS3PortOk returns a tuple with the S3Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetS3PortOk() (*int64, bool) {
	if o == nil || IsNil(o.S3Port) {
		return nil, false
	}
	return o.S3Port, true
}

// HasS3Port returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasS3Port() bool {
	if o != nil && !IsNil(o.S3Port) {
		return true
	}

	return false
}

// SetS3Port gets a reference to the given int64 and assigns it to the S3Port field.
func (o *OspGatewayRecord) SetS3Port(v int64) {
	o.S3Port = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OspGatewayRecord) SetStatus(v string) {
	o.Status = &v
}

// GetXmsPort returns the XmsPort field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetXmsPort() int64 {
	if o == nil || IsNil(o.XmsPort) {
		var ret int64
		return ret
	}
	return *o.XmsPort
}

// GetXmsPortOk returns a tuple with the XmsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetXmsPortOk() (*int64, bool) {
	if o == nil || IsNil(o.XmsPort) {
		return nil, false
	}
	return o.XmsPort, true
}

// HasXmsPort returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasXmsPort() bool {
	if o != nil && !IsNil(o.XmsPort) {
		return true
	}

	return false
}

// SetXmsPort gets a reference to the given int64 and assigns it to the XmsPort field.
func (o *OspGatewayRecord) SetXmsPort(v int64) {
	o.XmsPort = &v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetZoneId() int64 {
	if o == nil || IsNil(o.ZoneId) {
		var ret int64
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetZoneIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ZoneId) {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasZoneId() bool {
	if o != nil && !IsNil(o.ZoneId) {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given int64 and assigns it to the ZoneId field.
func (o *OspGatewayRecord) SetZoneId(v int64) {
	o.ZoneId = &v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *OspGatewayRecord) GetSamples() []OspGatewayStat {
	if o == nil || IsNil(o.Samples) {
		var ret []OspGatewayStat
		return ret
	}
	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OspGatewayRecord) GetSamplesOk() ([]OspGatewayStat, bool) {
	if o == nil || IsNil(o.Samples) {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *OspGatewayRecord) HasSamples() bool {
	if o != nil && !IsNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given []OspGatewayStat and assigns it to the Samples field.
func (o *OspGatewayRecord) SetSamples(v []OspGatewayStat) {
	o.Samples = v
}

func (o OspGatewayRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OspGatewayRecord) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Etag) {
		toSerialize["etag"] = o.Etag
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gateway_id"] = o.GatewayId
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PlatformIp) {
		toSerialize["platform_ip"] = o.PlatformIp
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.S3Port) {
		toSerialize["s3_port"] = o.S3Port
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.XmsPort) {
		toSerialize["xms_port"] = o.XmsPort
	}
	if !IsNil(o.ZoneId) {
		toSerialize["zone_id"] = o.ZoneId
	}
	if !IsNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	return toSerialize, nil
}

type NullableOspGatewayRecord struct {
	value *OspGatewayRecord
	isSet bool
}

func (v NullableOspGatewayRecord) Get() *OspGatewayRecord {
	return v.value
}

func (v *NullableOspGatewayRecord) Set(val *OspGatewayRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableOspGatewayRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableOspGatewayRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspGatewayRecord(val *OspGatewayRecord) *NullableOspGatewayRecord {
	return &NullableOspGatewayRecord{value: val, isSet: true}
}

func (v NullableOspGatewayRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspGatewayRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


