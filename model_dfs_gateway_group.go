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

// checks if the DfsGatewayGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsGatewayGroup{}

// DfsGatewayGroup DfsGatewayGroup defines model of dfs gateway group
type DfsGatewayGroup struct {
	ActionStatus *string `json:"action_status,omitempty"`
	Ad *FSActiveDirectoryNestview `json:"ad,omitempty"`
	BucketNum *int64 `json:"bucket_num,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	// container resource limit
	Cpus *int64 `json:"cpus,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Description *string `json:"description,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	GatewayNum *int64 `json:"gateway_num,omitempty"`
	HdfsNum *int64 `json:"hdfs_num,omitempty"`
	// hdfs domain auth type
	HdfsSecurities []string `json:"hdfs_securities,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Ldap *FSLdapNestview `json:"ldap,omitempty"`
	MemoryKbyte *int64 `json:"memory_kbyte,omitempty"`
	Name *string `json:"name,omitempty"`
	// NFS attributes
	NfsVersions []string `json:"nfs_versions,omitempty"`
	// FTP attributes
	Port *int64 `json:"port,omitempty"`
	ProtocolAuthType *map[string][]string `json:"protocol_auth_type,omitempty"`
	RootfsNum *int64 `json:"rootfs_num,omitempty"`
	S3Port *int64 `json:"s3_port,omitempty"`
	S3Scheme *string `json:"s3_scheme,omitempty"`
	S3UserNum *int64 `json:"s3_user_num,omitempty"`
	Securities []string `json:"securities,omitempty"`
	ShareNums *map[string]int64 `json:"share_nums,omitempty"`
	Shared *bool `json:"shared,omitempty"`
	Smb1Enabled *bool `json:"smb1_enabled,omitempty"`
	SmbBrowseable *bool `json:"smb_browseable,omitempty"`
	SmbHomes *bool `json:"smb_homes,omitempty"`
	SmbPorts []int64 `json:"smb_ports,omitempty"`
	// SMB attributes
	SmbType *string `json:"smb_type,omitempty"`
	SslCertificate *SSLCertificateNestview `json:"ssl_certificate,omitempty"`
	Status *string `json:"status,omitempty"`
	Types []string `json:"types,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	ZoneNum *int64 `json:"zone_num,omitempty"`
}

// NewDfsGatewayGroup instantiates a new DfsGatewayGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsGatewayGroup() *DfsGatewayGroup {
	this := DfsGatewayGroup{}
	return &this
}

// NewDfsGatewayGroupWithDefaults instantiates a new DfsGatewayGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsGatewayGroupWithDefaults() *DfsGatewayGroup {
	this := DfsGatewayGroup{}
	return &this
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *DfsGatewayGroup) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetAd returns the Ad field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetAd() FSActiveDirectoryNestview {
	if o == nil || IsNil(o.Ad) {
		var ret FSActiveDirectoryNestview
		return ret
	}
	return *o.Ad
}

// GetAdOk returns a tuple with the Ad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetAdOk() (*FSActiveDirectoryNestview, bool) {
	if o == nil || IsNil(o.Ad) {
		return nil, false
	}
	return o.Ad, true
}

// HasAd returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasAd() bool {
	if o != nil && !IsNil(o.Ad) {
		return true
	}

	return false
}

// SetAd gets a reference to the given FSActiveDirectoryNestview and assigns it to the Ad field.
func (o *DfsGatewayGroup) SetAd(v FSActiveDirectoryNestview) {
	o.Ad = &v
}

// GetBucketNum returns the BucketNum field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetBucketNum() int64 {
	if o == nil || IsNil(o.BucketNum) {
		var ret int64
		return ret
	}
	return *o.BucketNum
}

// GetBucketNumOk returns a tuple with the BucketNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetBucketNumOk() (*int64, bool) {
	if o == nil || IsNil(o.BucketNum) {
		return nil, false
	}
	return o.BucketNum, true
}

// HasBucketNum returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasBucketNum() bool {
	if o != nil && !IsNil(o.BucketNum) {
		return true
	}

	return false
}

// SetBucketNum gets a reference to the given int64 and assigns it to the BucketNum field.
func (o *DfsGatewayGroup) SetBucketNum(v int64) {
	o.BucketNum = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DfsGatewayGroup) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetCpus() int64 {
	if o == nil || IsNil(o.Cpus) {
		var ret int64
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetCpusOk() (*int64, bool) {
	if o == nil || IsNil(o.Cpus) {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasCpus() bool {
	if o != nil && !IsNil(o.Cpus) {
		return true
	}

	return false
}

// SetCpus gets a reference to the given int64 and assigns it to the Cpus field.
func (o *DfsGatewayGroup) SetCpus(v int64) {
	o.Cpus = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsGatewayGroup) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DfsGatewayGroup) SetDescription(v string) {
	o.Description = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetEncoding() string {
	if o == nil || IsNil(o.Encoding) {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetEncodingOk() (*string, bool) {
	if o == nil || IsNil(o.Encoding) {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasEncoding() bool {
	if o != nil && !IsNil(o.Encoding) {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *DfsGatewayGroup) SetEncoding(v string) {
	o.Encoding = &v
}

// GetGatewayNum returns the GatewayNum field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetGatewayNum() int64 {
	if o == nil || IsNil(o.GatewayNum) {
		var ret int64
		return ret
	}
	return *o.GatewayNum
}

// GetGatewayNumOk returns a tuple with the GatewayNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetGatewayNumOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayNum) {
		return nil, false
	}
	return o.GatewayNum, true
}

// HasGatewayNum returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasGatewayNum() bool {
	if o != nil && !IsNil(o.GatewayNum) {
		return true
	}

	return false
}

// SetGatewayNum gets a reference to the given int64 and assigns it to the GatewayNum field.
func (o *DfsGatewayGroup) SetGatewayNum(v int64) {
	o.GatewayNum = &v
}

// GetHdfsNum returns the HdfsNum field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetHdfsNum() int64 {
	if o == nil || IsNil(o.HdfsNum) {
		var ret int64
		return ret
	}
	return *o.HdfsNum
}

// GetHdfsNumOk returns a tuple with the HdfsNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetHdfsNumOk() (*int64, bool) {
	if o == nil || IsNil(o.HdfsNum) {
		return nil, false
	}
	return o.HdfsNum, true
}

// HasHdfsNum returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasHdfsNum() bool {
	if o != nil && !IsNil(o.HdfsNum) {
		return true
	}

	return false
}

// SetHdfsNum gets a reference to the given int64 and assigns it to the HdfsNum field.
func (o *DfsGatewayGroup) SetHdfsNum(v int64) {
	o.HdfsNum = &v
}

// GetHdfsSecurities returns the HdfsSecurities field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetHdfsSecurities() []string {
	if o == nil || IsNil(o.HdfsSecurities) {
		var ret []string
		return ret
	}
	return o.HdfsSecurities
}

// GetHdfsSecuritiesOk returns a tuple with the HdfsSecurities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetHdfsSecuritiesOk() ([]string, bool) {
	if o == nil || IsNil(o.HdfsSecurities) {
		return nil, false
	}
	return o.HdfsSecurities, true
}

// HasHdfsSecurities returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasHdfsSecurities() bool {
	if o != nil && !IsNil(o.HdfsSecurities) {
		return true
	}

	return false
}

// SetHdfsSecurities gets a reference to the given []string and assigns it to the HdfsSecurities field.
func (o *DfsGatewayGroup) SetHdfsSecurities(v []string) {
	o.HdfsSecurities = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsGatewayGroup) SetId(v int64) {
	o.Id = &v
}

// GetLdap returns the Ldap field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetLdap() FSLdapNestview {
	if o == nil || IsNil(o.Ldap) {
		var ret FSLdapNestview
		return ret
	}
	return *o.Ldap
}

// GetLdapOk returns a tuple with the Ldap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetLdapOk() (*FSLdapNestview, bool) {
	if o == nil || IsNil(o.Ldap) {
		return nil, false
	}
	return o.Ldap, true
}

// HasLdap returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasLdap() bool {
	if o != nil && !IsNil(o.Ldap) {
		return true
	}

	return false
}

// SetLdap gets a reference to the given FSLdapNestview and assigns it to the Ldap field.
func (o *DfsGatewayGroup) SetLdap(v FSLdapNestview) {
	o.Ldap = &v
}

// GetMemoryKbyte returns the MemoryKbyte field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetMemoryKbyte() int64 {
	if o == nil || IsNil(o.MemoryKbyte) {
		var ret int64
		return ret
	}
	return *o.MemoryKbyte
}

// GetMemoryKbyteOk returns a tuple with the MemoryKbyte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetMemoryKbyteOk() (*int64, bool) {
	if o == nil || IsNil(o.MemoryKbyte) {
		return nil, false
	}
	return o.MemoryKbyte, true
}

// HasMemoryKbyte returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasMemoryKbyte() bool {
	if o != nil && !IsNil(o.MemoryKbyte) {
		return true
	}

	return false
}

// SetMemoryKbyte gets a reference to the given int64 and assigns it to the MemoryKbyte field.
func (o *DfsGatewayGroup) SetMemoryKbyte(v int64) {
	o.MemoryKbyte = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DfsGatewayGroup) SetName(v string) {
	o.Name = &v
}

// GetNfsVersions returns the NfsVersions field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetNfsVersions() []string {
	if o == nil || IsNil(o.NfsVersions) {
		var ret []string
		return ret
	}
	return o.NfsVersions
}

// GetNfsVersionsOk returns a tuple with the NfsVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetNfsVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.NfsVersions) {
		return nil, false
	}
	return o.NfsVersions, true
}

// HasNfsVersions returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasNfsVersions() bool {
	if o != nil && !IsNil(o.NfsVersions) {
		return true
	}

	return false
}

// SetNfsVersions gets a reference to the given []string and assigns it to the NfsVersions field.
func (o *DfsGatewayGroup) SetNfsVersions(v []string) {
	o.NfsVersions = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *DfsGatewayGroup) SetPort(v int64) {
	o.Port = &v
}

// GetProtocolAuthType returns the ProtocolAuthType field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetProtocolAuthType() map[string][]string {
	if o == nil || IsNil(o.ProtocolAuthType) {
		var ret map[string][]string
		return ret
	}
	return *o.ProtocolAuthType
}

// GetProtocolAuthTypeOk returns a tuple with the ProtocolAuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetProtocolAuthTypeOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.ProtocolAuthType) {
		return nil, false
	}
	return o.ProtocolAuthType, true
}

// HasProtocolAuthType returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasProtocolAuthType() bool {
	if o != nil && !IsNil(o.ProtocolAuthType) {
		return true
	}

	return false
}

// SetProtocolAuthType gets a reference to the given map[string][]string and assigns it to the ProtocolAuthType field.
func (o *DfsGatewayGroup) SetProtocolAuthType(v map[string][]string) {
	o.ProtocolAuthType = &v
}

// GetRootfsNum returns the RootfsNum field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetRootfsNum() int64 {
	if o == nil || IsNil(o.RootfsNum) {
		var ret int64
		return ret
	}
	return *o.RootfsNum
}

// GetRootfsNumOk returns a tuple with the RootfsNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetRootfsNumOk() (*int64, bool) {
	if o == nil || IsNil(o.RootfsNum) {
		return nil, false
	}
	return o.RootfsNum, true
}

// HasRootfsNum returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasRootfsNum() bool {
	if o != nil && !IsNil(o.RootfsNum) {
		return true
	}

	return false
}

// SetRootfsNum gets a reference to the given int64 and assigns it to the RootfsNum field.
func (o *DfsGatewayGroup) SetRootfsNum(v int64) {
	o.RootfsNum = &v
}

// GetS3Port returns the S3Port field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetS3Port() int64 {
	if o == nil || IsNil(o.S3Port) {
		var ret int64
		return ret
	}
	return *o.S3Port
}

// GetS3PortOk returns a tuple with the S3Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetS3PortOk() (*int64, bool) {
	if o == nil || IsNil(o.S3Port) {
		return nil, false
	}
	return o.S3Port, true
}

// HasS3Port returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasS3Port() bool {
	if o != nil && !IsNil(o.S3Port) {
		return true
	}

	return false
}

// SetS3Port gets a reference to the given int64 and assigns it to the S3Port field.
func (o *DfsGatewayGroup) SetS3Port(v int64) {
	o.S3Port = &v
}

// GetS3Scheme returns the S3Scheme field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetS3Scheme() string {
	if o == nil || IsNil(o.S3Scheme) {
		var ret string
		return ret
	}
	return *o.S3Scheme
}

// GetS3SchemeOk returns a tuple with the S3Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetS3SchemeOk() (*string, bool) {
	if o == nil || IsNil(o.S3Scheme) {
		return nil, false
	}
	return o.S3Scheme, true
}

// HasS3Scheme returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasS3Scheme() bool {
	if o != nil && !IsNil(o.S3Scheme) {
		return true
	}

	return false
}

// SetS3Scheme gets a reference to the given string and assigns it to the S3Scheme field.
func (o *DfsGatewayGroup) SetS3Scheme(v string) {
	o.S3Scheme = &v
}

// GetS3UserNum returns the S3UserNum field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetS3UserNum() int64 {
	if o == nil || IsNil(o.S3UserNum) {
		var ret int64
		return ret
	}
	return *o.S3UserNum
}

// GetS3UserNumOk returns a tuple with the S3UserNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetS3UserNumOk() (*int64, bool) {
	if o == nil || IsNil(o.S3UserNum) {
		return nil, false
	}
	return o.S3UserNum, true
}

// HasS3UserNum returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasS3UserNum() bool {
	if o != nil && !IsNil(o.S3UserNum) {
		return true
	}

	return false
}

// SetS3UserNum gets a reference to the given int64 and assigns it to the S3UserNum field.
func (o *DfsGatewayGroup) SetS3UserNum(v int64) {
	o.S3UserNum = &v
}

// GetSecurities returns the Securities field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetSecurities() []string {
	if o == nil || IsNil(o.Securities) {
		var ret []string
		return ret
	}
	return o.Securities
}

// GetSecuritiesOk returns a tuple with the Securities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetSecuritiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Securities) {
		return nil, false
	}
	return o.Securities, true
}

// HasSecurities returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasSecurities() bool {
	if o != nil && !IsNil(o.Securities) {
		return true
	}

	return false
}

// SetSecurities gets a reference to the given []string and assigns it to the Securities field.
func (o *DfsGatewayGroup) SetSecurities(v []string) {
	o.Securities = v
}

// GetShareNums returns the ShareNums field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetShareNums() map[string]int64 {
	if o == nil || IsNil(o.ShareNums) {
		var ret map[string]int64
		return ret
	}
	return *o.ShareNums
}

// GetShareNumsOk returns a tuple with the ShareNums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetShareNumsOk() (*map[string]int64, bool) {
	if o == nil || IsNil(o.ShareNums) {
		return nil, false
	}
	return o.ShareNums, true
}

// HasShareNums returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasShareNums() bool {
	if o != nil && !IsNil(o.ShareNums) {
		return true
	}

	return false
}

// SetShareNums gets a reference to the given map[string]int64 and assigns it to the ShareNums field.
func (o *DfsGatewayGroup) SetShareNums(v map[string]int64) {
	o.ShareNums = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *DfsGatewayGroup) SetShared(v bool) {
	o.Shared = &v
}

// GetSmb1Enabled returns the Smb1Enabled field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetSmb1Enabled() bool {
	if o == nil || IsNil(o.Smb1Enabled) {
		var ret bool
		return ret
	}
	return *o.Smb1Enabled
}

// GetSmb1EnabledOk returns a tuple with the Smb1Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetSmb1EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Smb1Enabled) {
		return nil, false
	}
	return o.Smb1Enabled, true
}

// HasSmb1Enabled returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasSmb1Enabled() bool {
	if o != nil && !IsNil(o.Smb1Enabled) {
		return true
	}

	return false
}

// SetSmb1Enabled gets a reference to the given bool and assigns it to the Smb1Enabled field.
func (o *DfsGatewayGroup) SetSmb1Enabled(v bool) {
	o.Smb1Enabled = &v
}

// GetSmbBrowseable returns the SmbBrowseable field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetSmbBrowseable() bool {
	if o == nil || IsNil(o.SmbBrowseable) {
		var ret bool
		return ret
	}
	return *o.SmbBrowseable
}

// GetSmbBrowseableOk returns a tuple with the SmbBrowseable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetSmbBrowseableOk() (*bool, bool) {
	if o == nil || IsNil(o.SmbBrowseable) {
		return nil, false
	}
	return o.SmbBrowseable, true
}

// HasSmbBrowseable returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasSmbBrowseable() bool {
	if o != nil && !IsNil(o.SmbBrowseable) {
		return true
	}

	return false
}

// SetSmbBrowseable gets a reference to the given bool and assigns it to the SmbBrowseable field.
func (o *DfsGatewayGroup) SetSmbBrowseable(v bool) {
	o.SmbBrowseable = &v
}

// GetSmbHomes returns the SmbHomes field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetSmbHomes() bool {
	if o == nil || IsNil(o.SmbHomes) {
		var ret bool
		return ret
	}
	return *o.SmbHomes
}

// GetSmbHomesOk returns a tuple with the SmbHomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetSmbHomesOk() (*bool, bool) {
	if o == nil || IsNil(o.SmbHomes) {
		return nil, false
	}
	return o.SmbHomes, true
}

// HasSmbHomes returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasSmbHomes() bool {
	if o != nil && !IsNil(o.SmbHomes) {
		return true
	}

	return false
}

// SetSmbHomes gets a reference to the given bool and assigns it to the SmbHomes field.
func (o *DfsGatewayGroup) SetSmbHomes(v bool) {
	o.SmbHomes = &v
}

// GetSmbPorts returns the SmbPorts field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetSmbPorts() []int64 {
	if o == nil || IsNil(o.SmbPorts) {
		var ret []int64
		return ret
	}
	return o.SmbPorts
}

// GetSmbPortsOk returns a tuple with the SmbPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetSmbPortsOk() ([]int64, bool) {
	if o == nil || IsNil(o.SmbPorts) {
		return nil, false
	}
	return o.SmbPorts, true
}

// HasSmbPorts returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasSmbPorts() bool {
	if o != nil && !IsNil(o.SmbPorts) {
		return true
	}

	return false
}

// SetSmbPorts gets a reference to the given []int64 and assigns it to the SmbPorts field.
func (o *DfsGatewayGroup) SetSmbPorts(v []int64) {
	o.SmbPorts = v
}

// GetSmbType returns the SmbType field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetSmbType() string {
	if o == nil || IsNil(o.SmbType) {
		var ret string
		return ret
	}
	return *o.SmbType
}

// GetSmbTypeOk returns a tuple with the SmbType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetSmbTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SmbType) {
		return nil, false
	}
	return o.SmbType, true
}

// HasSmbType returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasSmbType() bool {
	if o != nil && !IsNil(o.SmbType) {
		return true
	}

	return false
}

// SetSmbType gets a reference to the given string and assigns it to the SmbType field.
func (o *DfsGatewayGroup) SetSmbType(v string) {
	o.SmbType = &v
}

// GetSslCertificate returns the SslCertificate field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetSslCertificate() SSLCertificateNestview {
	if o == nil || IsNil(o.SslCertificate) {
		var ret SSLCertificateNestview
		return ret
	}
	return *o.SslCertificate
}

// GetSslCertificateOk returns a tuple with the SslCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetSslCertificateOk() (*SSLCertificateNestview, bool) {
	if o == nil || IsNil(o.SslCertificate) {
		return nil, false
	}
	return o.SslCertificate, true
}

// HasSslCertificate returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasSslCertificate() bool {
	if o != nil && !IsNil(o.SslCertificate) {
		return true
	}

	return false
}

// SetSslCertificate gets a reference to the given SSLCertificateNestview and assigns it to the SslCertificate field.
func (o *DfsGatewayGroup) SetSslCertificate(v SSLCertificateNestview) {
	o.SslCertificate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DfsGatewayGroup) SetStatus(v string) {
	o.Status = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetTypes() []string {
	if o == nil || IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *DfsGatewayGroup) SetTypes(v []string) {
	o.Types = v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DfsGatewayGroup) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetZoneNum returns the ZoneNum field value if set, zero value otherwise.
func (o *DfsGatewayGroup) GetZoneNum() int64 {
	if o == nil || IsNil(o.ZoneNum) {
		var ret int64
		return ret
	}
	return *o.ZoneNum
}

// GetZoneNumOk returns a tuple with the ZoneNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsGatewayGroup) GetZoneNumOk() (*int64, bool) {
	if o == nil || IsNil(o.ZoneNum) {
		return nil, false
	}
	return o.ZoneNum, true
}

// HasZoneNum returns a boolean if a field has been set.
func (o *DfsGatewayGroup) HasZoneNum() bool {
	if o != nil && !IsNil(o.ZoneNum) {
		return true
	}

	return false
}

// SetZoneNum gets a reference to the given int64 and assigns it to the ZoneNum field.
func (o *DfsGatewayGroup) SetZoneNum(v int64) {
	o.ZoneNum = &v
}

func (o DfsGatewayGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsGatewayGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionStatus) {
		toSerialize["action_status"] = o.ActionStatus
	}
	if !IsNil(o.Ad) {
		toSerialize["ad"] = o.Ad
	}
	if !IsNil(o.BucketNum) {
		toSerialize["bucket_num"] = o.BucketNum
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Cpus) {
		toSerialize["cpus"] = o.Cpus
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Encoding) {
		toSerialize["encoding"] = o.Encoding
	}
	if !IsNil(o.GatewayNum) {
		toSerialize["gateway_num"] = o.GatewayNum
	}
	if !IsNil(o.HdfsNum) {
		toSerialize["hdfs_num"] = o.HdfsNum
	}
	if !IsNil(o.HdfsSecurities) {
		toSerialize["hdfs_securities"] = o.HdfsSecurities
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ldap) {
		toSerialize["ldap"] = o.Ldap
	}
	if !IsNil(o.MemoryKbyte) {
		toSerialize["memory_kbyte"] = o.MemoryKbyte
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NfsVersions) {
		toSerialize["nfs_versions"] = o.NfsVersions
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.ProtocolAuthType) {
		toSerialize["protocol_auth_type"] = o.ProtocolAuthType
	}
	if !IsNil(o.RootfsNum) {
		toSerialize["rootfs_num"] = o.RootfsNum
	}
	if !IsNil(o.S3Port) {
		toSerialize["s3_port"] = o.S3Port
	}
	if !IsNil(o.S3Scheme) {
		toSerialize["s3_scheme"] = o.S3Scheme
	}
	if !IsNil(o.S3UserNum) {
		toSerialize["s3_user_num"] = o.S3UserNum
	}
	if !IsNil(o.Securities) {
		toSerialize["securities"] = o.Securities
	}
	if !IsNil(o.ShareNums) {
		toSerialize["share_nums"] = o.ShareNums
	}
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	if !IsNil(o.Smb1Enabled) {
		toSerialize["smb1_enabled"] = o.Smb1Enabled
	}
	if !IsNil(o.SmbBrowseable) {
		toSerialize["smb_browseable"] = o.SmbBrowseable
	}
	if !IsNil(o.SmbHomes) {
		toSerialize["smb_homes"] = o.SmbHomes
	}
	if !IsNil(o.SmbPorts) {
		toSerialize["smb_ports"] = o.SmbPorts
	}
	if !IsNil(o.SmbType) {
		toSerialize["smb_type"] = o.SmbType
	}
	if !IsNil(o.SslCertificate) {
		toSerialize["ssl_certificate"] = o.SslCertificate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.ZoneNum) {
		toSerialize["zone_num"] = o.ZoneNum
	}
	return toSerialize, nil
}

type NullableDfsGatewayGroup struct {
	value *DfsGatewayGroup
	isSet bool
}

func (v NullableDfsGatewayGroup) Get() *DfsGatewayGroup {
	return v.value
}

func (v *NullableDfsGatewayGroup) Set(val *DfsGatewayGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsGatewayGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsGatewayGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsGatewayGroup(val *DfsGatewayGroup) *NullableDfsGatewayGroup {
	return &NullableDfsGatewayGroup{value: val, isSet: true}
}

func (v NullableDfsGatewayGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsGatewayGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


