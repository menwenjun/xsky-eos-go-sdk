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

// checks if the DfsGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsGateway{}

// DfsGateway DfsGateway defines model of dfs gateway
type DfsGateway struct {
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	ConnNum *int64 `json:"conn_num,omitempty"`
	// container resource limit
	Cpus *int64 `json:"cpus,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	CtdbServiceStatus *string `json:"ctdb_service_status,omitempty"`
	DfsGatewayGroup *DfsGatewayGroupNestview `json:"dfs_gateway_group,omitempty"`
	Host *HostNestview `json:"host,omitempty"`
	Id *int64 `json:"id,omitempty"`
	MemoryKbyte *int64 `json:"memory_kbyte,omitempty"`
	S3Port *int64 `json:"s3_port,omitempty"`
	S3Scheme *string `json:"s3_scheme,omitempty"`
	SslCertificate *SSLCertificateNestview `json:"ssl_certificate,omitempty"`
	Status *string `json:"status,omitempty"`
	Update *time.Time `json:"update,omitempty"`
}

// NewDfsGateway instantiates a new DfsGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsGateway() *DfsGateway {
	this := DfsGateway{}
	return &this
}

// NewDfsGatewayWithDefaults instantiates a new DfsGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsGatewayWithDefaults() *DfsGateway {
	this := DfsGateway{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DfsGateway) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DfsGateway) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DfsGateway) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetConnNum returns the ConnNum field value if set, zero value otherwise.
func (o *DfsGateway) GetConnNum() int64 {
	if o == nil || IsNil(o.ConnNum) {
		var ret int64
		return ret
	}
	return *o.ConnNum
}

// GetConnNumOk returns a tuple with the ConnNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetConnNumOk() (*int64, bool) {
	if o == nil || IsNil(o.ConnNum) {
		return nil, false
	}
	return o.ConnNum, true
}

// HasConnNum returns a boolean if a field has been set.
func (o *DfsGateway) HasConnNum() bool {
	if o != nil && !IsNil(o.ConnNum) {
		return true
	}

	return false
}

// SetConnNum gets a reference to the given int64 and assigns it to the ConnNum field.
func (o *DfsGateway) SetConnNum(v int64) {
	o.ConnNum = &v
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *DfsGateway) GetCpus() int64 {
	if o == nil || IsNil(o.Cpus) {
		var ret int64
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetCpusOk() (*int64, bool) {
	if o == nil || IsNil(o.Cpus) {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *DfsGateway) HasCpus() bool {
	if o != nil && !IsNil(o.Cpus) {
		return true
	}

	return false
}

// SetCpus gets a reference to the given int64 and assigns it to the Cpus field.
func (o *DfsGateway) SetCpus(v int64) {
	o.Cpus = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsGateway) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsGateway) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsGateway) SetCreate(v time.Time) {
	o.Create = &v
}

// GetCtdbServiceStatus returns the CtdbServiceStatus field value if set, zero value otherwise.
func (o *DfsGateway) GetCtdbServiceStatus() string {
	if o == nil || IsNil(o.CtdbServiceStatus) {
		var ret string
		return ret
	}
	return *o.CtdbServiceStatus
}

// GetCtdbServiceStatusOk returns a tuple with the CtdbServiceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetCtdbServiceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CtdbServiceStatus) {
		return nil, false
	}
	return o.CtdbServiceStatus, true
}

// HasCtdbServiceStatus returns a boolean if a field has been set.
func (o *DfsGateway) HasCtdbServiceStatus() bool {
	if o != nil && !IsNil(o.CtdbServiceStatus) {
		return true
	}

	return false
}

// SetCtdbServiceStatus gets a reference to the given string and assigns it to the CtdbServiceStatus field.
func (o *DfsGateway) SetCtdbServiceStatus(v string) {
	o.CtdbServiceStatus = &v
}

// GetDfsGatewayGroup returns the DfsGatewayGroup field value if set, zero value otherwise.
func (o *DfsGateway) GetDfsGatewayGroup() DfsGatewayGroupNestview {
	if o == nil || IsNil(o.DfsGatewayGroup) {
		var ret DfsGatewayGroupNestview
		return ret
	}
	return *o.DfsGatewayGroup
}

// GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool) {
	if o == nil || IsNil(o.DfsGatewayGroup) {
		return nil, false
	}
	return o.DfsGatewayGroup, true
}

// HasDfsGatewayGroup returns a boolean if a field has been set.
func (o *DfsGateway) HasDfsGatewayGroup() bool {
	if o != nil && !IsNil(o.DfsGatewayGroup) {
		return true
	}

	return false
}

// SetDfsGatewayGroup gets a reference to the given DfsGatewayGroupNestview and assigns it to the DfsGatewayGroup field.
func (o *DfsGateway) SetDfsGatewayGroup(v DfsGatewayGroupNestview) {
	o.DfsGatewayGroup = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *DfsGateway) GetHost() HostNestview {
	if o == nil || IsNil(o.Host) {
		var ret HostNestview
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetHostOk() (*HostNestview, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *DfsGateway) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostNestview and assigns it to the Host field.
func (o *DfsGateway) SetHost(v HostNestview) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsGateway) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsGateway) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsGateway) SetId(v int64) {
	o.Id = &v
}

// GetMemoryKbyte returns the MemoryKbyte field value if set, zero value otherwise.
func (o *DfsGateway) GetMemoryKbyte() int64 {
	if o == nil || IsNil(o.MemoryKbyte) {
		var ret int64
		return ret
	}
	return *o.MemoryKbyte
}

// GetMemoryKbyteOk returns a tuple with the MemoryKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetMemoryKbyteOk() (*int64, bool) {
	if o == nil || IsNil(o.MemoryKbyte) {
		return nil, false
	}
	return o.MemoryKbyte, true
}

// HasMemoryKbyte returns a boolean if a field has been set.
func (o *DfsGateway) HasMemoryKbyte() bool {
	if o != nil && !IsNil(o.MemoryKbyte) {
		return true
	}

	return false
}

// SetMemoryKbyte gets a reference to the given int64 and assigns it to the MemoryKbyte field.
func (o *DfsGateway) SetMemoryKbyte(v int64) {
	o.MemoryKbyte = &v
}

// GetS3Port returns the S3Port field value if set, zero value otherwise.
func (o *DfsGateway) GetS3Port() int64 {
	if o == nil || IsNil(o.S3Port) {
		var ret int64
		return ret
	}
	return *o.S3Port
}

// GetS3PortOk returns a tuple with the S3Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetS3PortOk() (*int64, bool) {
	if o == nil || IsNil(o.S3Port) {
		return nil, false
	}
	return o.S3Port, true
}

// HasS3Port returns a boolean if a field has been set.
func (o *DfsGateway) HasS3Port() bool {
	if o != nil && !IsNil(o.S3Port) {
		return true
	}

	return false
}

// SetS3Port gets a reference to the given int64 and assigns it to the S3Port field.
func (o *DfsGateway) SetS3Port(v int64) {
	o.S3Port = &v
}

// GetS3Scheme returns the S3Scheme field value if set, zero value otherwise.
func (o *DfsGateway) GetS3Scheme() string {
	if o == nil || IsNil(o.S3Scheme) {
		var ret string
		return ret
	}
	return *o.S3Scheme
}

// GetS3SchemeOk returns a tuple with the S3Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetS3SchemeOk() (*string, bool) {
	if o == nil || IsNil(o.S3Scheme) {
		return nil, false
	}
	return o.S3Scheme, true
}

// HasS3Scheme returns a boolean if a field has been set.
func (o *DfsGateway) HasS3Scheme() bool {
	if o != nil && !IsNil(o.S3Scheme) {
		return true
	}

	return false
}

// SetS3Scheme gets a reference to the given string and assigns it to the S3Scheme field.
func (o *DfsGateway) SetS3Scheme(v string) {
	o.S3Scheme = &v
}

// GetSslCertificate returns the SslCertificate field value if set, zero value otherwise.
func (o *DfsGateway) GetSslCertificate() SSLCertificateNestview {
	if o == nil || IsNil(o.SslCertificate) {
		var ret SSLCertificateNestview
		return ret
	}
	return *o.SslCertificate
}

// GetSslCertificateOk returns a tuple with the SslCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetSslCertificateOk() (*SSLCertificateNestview, bool) {
	if o == nil || IsNil(o.SslCertificate) {
		return nil, false
	}
	return o.SslCertificate, true
}

// HasSslCertificate returns a boolean if a field has been set.
func (o *DfsGateway) HasSslCertificate() bool {
	if o != nil && !IsNil(o.SslCertificate) {
		return true
	}

	return false
}

// SetSslCertificate gets a reference to the given SSLCertificateNestview and assigns it to the SslCertificate field.
func (o *DfsGateway) SetSslCertificate(v SSLCertificateNestview) {
	o.SslCertificate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DfsGateway) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DfsGateway) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DfsGateway) SetStatus(v string) {
	o.Status = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DfsGateway) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGateway) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DfsGateway) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DfsGateway) SetUpdate(v time.Time) {
	o.Update = &v
}

func (o DfsGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.ConnNum) {
		toSerialize["conn_num"] = o.ConnNum
	}
	if !IsNil(o.Cpus) {
		toSerialize["cpus"] = o.Cpus
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.CtdbServiceStatus) {
		toSerialize["ctdb_service_status"] = o.CtdbServiceStatus
	}
	if !IsNil(o.DfsGatewayGroup) {
		toSerialize["dfs_gateway_group"] = o.DfsGatewayGroup
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MemoryKbyte) {
		toSerialize["memory_kbyte"] = o.MemoryKbyte
	}
	if !IsNil(o.S3Port) {
		toSerialize["s3_port"] = o.S3Port
	}
	if !IsNil(o.S3Scheme) {
		toSerialize["s3_scheme"] = o.S3Scheme
	}
	if !IsNil(o.SslCertificate) {
		toSerialize["ssl_certificate"] = o.SslCertificate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	return toSerialize, nil
}

type NullableDfsGateway struct {
	value *DfsGateway
	isSet bool
}

func (v NullableDfsGateway) Get() *DfsGateway {
	return v.value
}

func (v *NullableDfsGateway) Set(val *DfsGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsGateway(val *DfsGateway) *NullableDfsGateway {
	return &NullableDfsGateway{value: val, isSet: true}
}

func (v NullableDfsGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


