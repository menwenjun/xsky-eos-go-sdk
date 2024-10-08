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

// checks if the S3LoadBalancer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3LoadBalancer{}

// S3LoadBalancer S3LoadBalancer is a load balancer for rados loadBalancers +X:model:stats;
type S3LoadBalancer struct {
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	Description *string `json:"description,omitempty"`
	Group *S3LoadBalancerGroupNestview `json:"group,omitempty"`
	Host *HostNestview `json:"host,omitempty"`
	Http2Enabled *bool `json:"http2_enabled,omitempty"`
	HttpsPort *int64 `json:"https_port,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Keepaliveds []S3LoadBalancerKeepalived `json:"keepaliveds,omitempty"`
	Name *string `json:"name,omitempty"`
	OspZoneId *int64 `json:"osp_zone_id,omitempty"`
	OssApiEnabled *bool `json:"oss_api_enabled,omitempty"`
	Port *int64 `json:"port,omitempty"`
	Roles []string `json:"roles,omitempty"`
	SearchHttpsPort *int64 `json:"search_https_port,omitempty"`
	SearchPort *int64 `json:"search_port,omitempty"`
	SslCertificate *SSLCertificateNestview `json:"ssl_certificate,omitempty"`
	Status *string `json:"status,omitempty"`
	SyncHttpsPort *int64 `json:"sync_https_port,omitempty"`
	SyncPort *int64 `json:"sync_port,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Vips *string `json:"vips,omitempty"`
	WebServicePort *int64 `json:"web_service_port,omitempty"`
}

// NewS3LoadBalancer instantiates a new S3LoadBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3LoadBalancer() *S3LoadBalancer {
	this := S3LoadBalancer{}
	return &this
}

// NewS3LoadBalancerWithDefaults instantiates a new S3LoadBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3LoadBalancerWithDefaults() *S3LoadBalancer {
	this := S3LoadBalancer{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *S3LoadBalancer) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *S3LoadBalancer) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *S3LoadBalancer) SetDescription(v string) {
	o.Description = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetGroup() S3LoadBalancerGroupNestview {
	if o == nil || IsNil(o.Group) {
		var ret S3LoadBalancerGroupNestview
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetGroupOk() (*S3LoadBalancerGroupNestview, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given S3LoadBalancerGroupNestview and assigns it to the Group field.
func (o *S3LoadBalancer) SetGroup(v S3LoadBalancerGroupNestview) {
	o.Group = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetHost() HostNestview {
	if o == nil || IsNil(o.Host) {
		var ret HostNestview
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetHostOk() (*HostNestview, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostNestview and assigns it to the Host field.
func (o *S3LoadBalancer) SetHost(v HostNestview) {
	o.Host = &v
}

// GetHttp2Enabled returns the Http2Enabled field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetHttp2Enabled() bool {
	if o == nil || IsNil(o.Http2Enabled) {
		var ret bool
		return ret
	}
	return *o.Http2Enabled
}

// GetHttp2EnabledOk returns a tuple with the Http2Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetHttp2EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Http2Enabled) {
		return nil, false
	}
	return o.Http2Enabled, true
}

// HasHttp2Enabled returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasHttp2Enabled() bool {
	if o != nil && !IsNil(o.Http2Enabled) {
		return true
	}

	return false
}

// SetHttp2Enabled gets a reference to the given bool and assigns it to the Http2Enabled field.
func (o *S3LoadBalancer) SetHttp2Enabled(v bool) {
	o.Http2Enabled = &v
}

// GetHttpsPort returns the HttpsPort field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetHttpsPort() int64 {
	if o == nil || IsNil(o.HttpsPort) {
		var ret int64
		return ret
	}
	return *o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetHttpsPortOk() (*int64, bool) {
	if o == nil || IsNil(o.HttpsPort) {
		return nil, false
	}
	return o.HttpsPort, true
}

// HasHttpsPort returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasHttpsPort() bool {
	if o != nil && !IsNil(o.HttpsPort) {
		return true
	}

	return false
}

// SetHttpsPort gets a reference to the given int64 and assigns it to the HttpsPort field.
func (o *S3LoadBalancer) SetHttpsPort(v int64) {
	o.HttpsPort = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *S3LoadBalancer) SetId(v int64) {
	o.Id = &v
}

// GetKeepaliveds returns the Keepaliveds field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetKeepaliveds() []S3LoadBalancerKeepalived {
	if o == nil || IsNil(o.Keepaliveds) {
		var ret []S3LoadBalancerKeepalived
		return ret
	}
	return o.Keepaliveds
}

// GetKeepalivedsOk returns a tuple with the Keepaliveds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetKeepalivedsOk() ([]S3LoadBalancerKeepalived, bool) {
	if o == nil || IsNil(o.Keepaliveds) {
		return nil, false
	}
	return o.Keepaliveds, true
}

// HasKeepaliveds returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasKeepaliveds() bool {
	if o != nil && !IsNil(o.Keepaliveds) {
		return true
	}

	return false
}

// SetKeepaliveds gets a reference to the given []S3LoadBalancerKeepalived and assigns it to the Keepaliveds field.
func (o *S3LoadBalancer) SetKeepaliveds(v []S3LoadBalancerKeepalived) {
	o.Keepaliveds = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *S3LoadBalancer) SetName(v string) {
	o.Name = &v
}

// GetOspZoneId returns the OspZoneId field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetOspZoneId() int64 {
	if o == nil || IsNil(o.OspZoneId) {
		var ret int64
		return ret
	}
	return *o.OspZoneId
}

// GetOspZoneIdOk returns a tuple with the OspZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetOspZoneIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OspZoneId) {
		return nil, false
	}
	return o.OspZoneId, true
}

// HasOspZoneId returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasOspZoneId() bool {
	if o != nil && !IsNil(o.OspZoneId) {
		return true
	}

	return false
}

// SetOspZoneId gets a reference to the given int64 and assigns it to the OspZoneId field.
func (o *S3LoadBalancer) SetOspZoneId(v int64) {
	o.OspZoneId = &v
}

// GetOssApiEnabled returns the OssApiEnabled field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetOssApiEnabled() bool {
	if o == nil || IsNil(o.OssApiEnabled) {
		var ret bool
		return ret
	}
	return *o.OssApiEnabled
}

// GetOssApiEnabledOk returns a tuple with the OssApiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetOssApiEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OssApiEnabled) {
		return nil, false
	}
	return o.OssApiEnabled, true
}

// HasOssApiEnabled returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasOssApiEnabled() bool {
	if o != nil && !IsNil(o.OssApiEnabled) {
		return true
	}

	return false
}

// SetOssApiEnabled gets a reference to the given bool and assigns it to the OssApiEnabled field.
func (o *S3LoadBalancer) SetOssApiEnabled(v bool) {
	o.OssApiEnabled = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *S3LoadBalancer) SetPort(v int64) {
	o.Port = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *S3LoadBalancer) SetRoles(v []string) {
	o.Roles = v
}

// GetSearchHttpsPort returns the SearchHttpsPort field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetSearchHttpsPort() int64 {
	if o == nil || IsNil(o.SearchHttpsPort) {
		var ret int64
		return ret
	}
	return *o.SearchHttpsPort
}

// GetSearchHttpsPortOk returns a tuple with the SearchHttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetSearchHttpsPortOk() (*int64, bool) {
	if o == nil || IsNil(o.SearchHttpsPort) {
		return nil, false
	}
	return o.SearchHttpsPort, true
}

// HasSearchHttpsPort returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasSearchHttpsPort() bool {
	if o != nil && !IsNil(o.SearchHttpsPort) {
		return true
	}

	return false
}

// SetSearchHttpsPort gets a reference to the given int64 and assigns it to the SearchHttpsPort field.
func (o *S3LoadBalancer) SetSearchHttpsPort(v int64) {
	o.SearchHttpsPort = &v
}

// GetSearchPort returns the SearchPort field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetSearchPort() int64 {
	if o == nil || IsNil(o.SearchPort) {
		var ret int64
		return ret
	}
	return *o.SearchPort
}

// GetSearchPortOk returns a tuple with the SearchPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetSearchPortOk() (*int64, bool) {
	if o == nil || IsNil(o.SearchPort) {
		return nil, false
	}
	return o.SearchPort, true
}

// HasSearchPort returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasSearchPort() bool {
	if o != nil && !IsNil(o.SearchPort) {
		return true
	}

	return false
}

// SetSearchPort gets a reference to the given int64 and assigns it to the SearchPort field.
func (o *S3LoadBalancer) SetSearchPort(v int64) {
	o.SearchPort = &v
}

// GetSslCertificate returns the SslCertificate field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetSslCertificate() SSLCertificateNestview {
	if o == nil || IsNil(o.SslCertificate) {
		var ret SSLCertificateNestview
		return ret
	}
	return *o.SslCertificate
}

// GetSslCertificateOk returns a tuple with the SslCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetSslCertificateOk() (*SSLCertificateNestview, bool) {
	if o == nil || IsNil(o.SslCertificate) {
		return nil, false
	}
	return o.SslCertificate, true
}

// HasSslCertificate returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasSslCertificate() bool {
	if o != nil && !IsNil(o.SslCertificate) {
		return true
	}

	return false
}

// SetSslCertificate gets a reference to the given SSLCertificateNestview and assigns it to the SslCertificate field.
func (o *S3LoadBalancer) SetSslCertificate(v SSLCertificateNestview) {
	o.SslCertificate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *S3LoadBalancer) SetStatus(v string) {
	o.Status = &v
}

// GetSyncHttpsPort returns the SyncHttpsPort field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetSyncHttpsPort() int64 {
	if o == nil || IsNil(o.SyncHttpsPort) {
		var ret int64
		return ret
	}
	return *o.SyncHttpsPort
}

// GetSyncHttpsPortOk returns a tuple with the SyncHttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetSyncHttpsPortOk() (*int64, bool) {
	if o == nil || IsNil(o.SyncHttpsPort) {
		return nil, false
	}
	return o.SyncHttpsPort, true
}

// HasSyncHttpsPort returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasSyncHttpsPort() bool {
	if o != nil && !IsNil(o.SyncHttpsPort) {
		return true
	}

	return false
}

// SetSyncHttpsPort gets a reference to the given int64 and assigns it to the SyncHttpsPort field.
func (o *S3LoadBalancer) SetSyncHttpsPort(v int64) {
	o.SyncHttpsPort = &v
}

// GetSyncPort returns the SyncPort field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetSyncPort() int64 {
	if o == nil || IsNil(o.SyncPort) {
		var ret int64
		return ret
	}
	return *o.SyncPort
}

// GetSyncPortOk returns a tuple with the SyncPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetSyncPortOk() (*int64, bool) {
	if o == nil || IsNil(o.SyncPort) {
		return nil, false
	}
	return o.SyncPort, true
}

// HasSyncPort returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasSyncPort() bool {
	if o != nil && !IsNil(o.SyncPort) {
		return true
	}

	return false
}

// SetSyncPort gets a reference to the given int64 and assigns it to the SyncPort field.
func (o *S3LoadBalancer) SetSyncPort(v int64) {
	o.SyncPort = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *S3LoadBalancer) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetVips returns the Vips field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetVips() string {
	if o == nil || IsNil(o.Vips) {
		var ret string
		return ret
	}
	return *o.Vips
}

// GetVipsOk returns a tuple with the Vips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetVipsOk() (*string, bool) {
	if o == nil || IsNil(o.Vips) {
		return nil, false
	}
	return o.Vips, true
}

// HasVips returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasVips() bool {
	if o != nil && !IsNil(o.Vips) {
		return true
	}

	return false
}

// SetVips gets a reference to the given string and assigns it to the Vips field.
func (o *S3LoadBalancer) SetVips(v string) {
	o.Vips = &v
}

// GetWebServicePort returns the WebServicePort field value if set, zero value otherwise.
func (o *S3LoadBalancer) GetWebServicePort() int64 {
	if o == nil || IsNil(o.WebServicePort) {
		var ret int64
		return ret
	}
	return *o.WebServicePort
}

// GetWebServicePortOk returns a tuple with the WebServicePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3LoadBalancer) GetWebServicePortOk() (*int64, bool) {
	if o == nil || IsNil(o.WebServicePort) {
		return nil, false
	}
	return o.WebServicePort, true
}

// HasWebServicePort returns a boolean if a field has been set.
func (o *S3LoadBalancer) HasWebServicePort() bool {
	if o != nil && !IsNil(o.WebServicePort) {
		return true
	}

	return false
}

// SetWebServicePort gets a reference to the given int64 and assigns it to the WebServicePort field.
func (o *S3LoadBalancer) SetWebServicePort(v int64) {
	o.WebServicePort = &v
}

func (o S3LoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3LoadBalancer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Http2Enabled) {
		toSerialize["http2_enabled"] = o.Http2Enabled
	}
	if !IsNil(o.HttpsPort) {
		toSerialize["https_port"] = o.HttpsPort
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Keepaliveds) {
		toSerialize["keepaliveds"] = o.Keepaliveds
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OspZoneId) {
		toSerialize["osp_zone_id"] = o.OspZoneId
	}
	if !IsNil(o.OssApiEnabled) {
		toSerialize["oss_api_enabled"] = o.OssApiEnabled
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.SearchHttpsPort) {
		toSerialize["search_https_port"] = o.SearchHttpsPort
	}
	if !IsNil(o.SearchPort) {
		toSerialize["search_port"] = o.SearchPort
	}
	if !IsNil(o.SslCertificate) {
		toSerialize["ssl_certificate"] = o.SslCertificate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SyncHttpsPort) {
		toSerialize["sync_https_port"] = o.SyncHttpsPort
	}
	if !IsNil(o.SyncPort) {
		toSerialize["sync_port"] = o.SyncPort
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Vips) {
		toSerialize["vips"] = o.Vips
	}
	if !IsNil(o.WebServicePort) {
		toSerialize["web_service_port"] = o.WebServicePort
	}
	return toSerialize, nil
}

type NullableS3LoadBalancer struct {
	value *S3LoadBalancer
	isSet bool
}

func (v NullableS3LoadBalancer) Get() *S3LoadBalancer {
	return v.value
}

func (v *NullableS3LoadBalancer) Set(val *S3LoadBalancer) {
	v.value = val
	v.isSet = true
}

func (v NullableS3LoadBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableS3LoadBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3LoadBalancer(val *S3LoadBalancer) *NullableS3LoadBalancer {
	return &NullableS3LoadBalancer{value: val, isSet: true}
}

func (v NullableS3LoadBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3LoadBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


