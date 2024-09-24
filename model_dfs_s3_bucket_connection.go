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

// checks if the DfsS3BucketConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfsS3BucketConnection{}

// DfsS3BucketConnection DfsS3BucketConnection defines model of dfs s3 bucket connection
type DfsS3BucketConnection struct {
	Api *string `json:"api,omitempty"`
	Bucket *DfsS3BucketNestview `json:"bucket,omitempty"`
	ClientIp *string `json:"client_ip,omitempty"`
	ClientPort *int64 `json:"client_port,omitempty"`
	Cluster *ClusterNestview `json:"cluster,omitempty"`
	ConnectedAt *time.Time `json:"connected_at,omitempty"`
	Create *time.Time `json:"create,omitempty"`
	DfsGateway *DfsGatewayNestview `json:"dfs_gateway,omitempty"`
	Id *int64 `json:"id,omitempty"`
	ProtocolVersion *string `json:"protocol_version,omitempty"`
	RequestId *string `json:"request_id,omitempty"`
	Update *time.Time `json:"update,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewDfsS3BucketConnection instantiates a new DfsS3BucketConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfsS3BucketConnection() *DfsS3BucketConnection {
	this := DfsS3BucketConnection{}
	return &this
}

// NewDfsS3BucketConnectionWithDefaults instantiates a new DfsS3BucketConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfsS3BucketConnectionWithDefaults() *DfsS3BucketConnection {
	this := DfsS3BucketConnection{}
	return &this
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetApi() string {
	if o == nil || IsNil(o.Api) {
		var ret string
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetApiOk() (*string, bool) {
	if o == nil || IsNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasApi() bool {
	if o != nil && !IsNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given string and assigns it to the Api field.
func (o *DfsS3BucketConnection) SetApi(v string) {
	o.Api = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetBucket() DfsS3BucketNestview {
	if o == nil || IsNil(o.Bucket) {
		var ret DfsS3BucketNestview
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetBucketOk() (*DfsS3BucketNestview, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given DfsS3BucketNestview and assigns it to the Bucket field.
func (o *DfsS3BucketConnection) SetBucket(v DfsS3BucketNestview) {
	o.Bucket = &v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetClientIp() string {
	if o == nil || IsNil(o.ClientIp) {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetClientIpOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIp) {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasClientIp() bool {
	if o != nil && !IsNil(o.ClientIp) {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *DfsS3BucketConnection) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetClientPort returns the ClientPort field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetClientPort() int64 {
	if o == nil || IsNil(o.ClientPort) {
		var ret int64
		return ret
	}
	return *o.ClientPort
}

// GetClientPortOk returns a tuple with the ClientPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetClientPortOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientPort) {
		return nil, false
	}
	return o.ClientPort, true
}

// HasClientPort returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasClientPort() bool {
	if o != nil && !IsNil(o.ClientPort) {
		return true
	}

	return false
}

// SetClientPort gets a reference to the given int64 and assigns it to the ClientPort field.
func (o *DfsS3BucketConnection) SetClientPort(v int64) {
	o.ClientPort = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetCluster() ClusterNestview {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterNestview
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetClusterOk() (*ClusterNestview, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterNestview and assigns it to the Cluster field.
func (o *DfsS3BucketConnection) SetCluster(v ClusterNestview) {
	o.Cluster = &v
}

// GetConnectedAt returns the ConnectedAt field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetConnectedAt() time.Time {
	if o == nil || IsNil(o.ConnectedAt) {
		var ret time.Time
		return ret
	}
	return *o.ConnectedAt
}

// GetConnectedAtOk returns a tuple with the ConnectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetConnectedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ConnectedAt) {
		return nil, false
	}
	return o.ConnectedAt, true
}

// HasConnectedAt returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasConnectedAt() bool {
	if o != nil && !IsNil(o.ConnectedAt) {
		return true
	}

	return false
}

// SetConnectedAt gets a reference to the given time.Time and assigns it to the ConnectedAt field.
func (o *DfsS3BucketConnection) SetConnectedAt(v time.Time) {
	o.ConnectedAt = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *DfsS3BucketConnection) SetCreate(v time.Time) {
	o.Create = &v
}

// GetDfsGateway returns the DfsGateway field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetDfsGateway() DfsGatewayNestview {
	if o == nil || IsNil(o.DfsGateway) {
		var ret DfsGatewayNestview
		return ret
	}
	return *o.DfsGateway
}

// GetDfsGatewayOk returns a tuple with the DfsGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetDfsGatewayOk() (*DfsGatewayNestview, bool) {
	if o == nil || IsNil(o.DfsGateway) {
		return nil, false
	}
	return o.DfsGateway, true
}

// HasDfsGateway returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasDfsGateway() bool {
	if o != nil && !IsNil(o.DfsGateway) {
		return true
	}

	return false
}

// SetDfsGateway gets a reference to the given DfsGatewayNestview and assigns it to the DfsGateway field.
func (o *DfsS3BucketConnection) SetDfsGateway(v DfsGatewayNestview) {
	o.DfsGateway = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DfsS3BucketConnection) SetId(v int64) {
	o.Id = &v
}

// GetProtocolVersion returns the ProtocolVersion field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetProtocolVersion() string {
	if o == nil || IsNil(o.ProtocolVersion) {
		var ret string
		return ret
	}
	return *o.ProtocolVersion
}

// GetProtocolVersionOk returns a tuple with the ProtocolVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetProtocolVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolVersion) {
		return nil, false
	}
	return o.ProtocolVersion, true
}

// HasProtocolVersion returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasProtocolVersion() bool {
	if o != nil && !IsNil(o.ProtocolVersion) {
		return true
	}

	return false
}

// SetProtocolVersion gets a reference to the given string and assigns it to the ProtocolVersion field.
func (o *DfsS3BucketConnection) SetProtocolVersion(v string) {
	o.ProtocolVersion = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DfsS3BucketConnection) SetRequestId(v string) {
	o.RequestId = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *DfsS3BucketConnection) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DfsS3BucketConnection) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfsS3BucketConnection) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DfsS3BucketConnection) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DfsS3BucketConnection) SetUsername(v string) {
	o.Username = &v
}

func (o DfsS3BucketConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfsS3BucketConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
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
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProtocolVersion) {
		toSerialize["protocol_version"] = o.ProtocolVersion
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableDfsS3BucketConnection struct {
	value *DfsS3BucketConnection
	isSet bool
}

func (v NullableDfsS3BucketConnection) Get() *DfsS3BucketConnection {
	return v.value
}

func (v *NullableDfsS3BucketConnection) Set(val *DfsS3BucketConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableDfsS3BucketConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableDfsS3BucketConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfsS3BucketConnection(val *DfsS3BucketConnection) *NullableDfsS3BucketConnection {
	return &NullableDfsS3BucketConnection{value: val, isSet: true}
}

func (v NullableDfsS3BucketConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfsS3BucketConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


