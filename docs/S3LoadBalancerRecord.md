# S3LoadBalancerRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**S3LoadBalancerGroupNestview**](S3LoadBalancerGroupNestview.md) |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Http2Enabled** | Pointer to **bool** |  | [optional] 
**HttpsPort** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Keepaliveds** | Pointer to [**[]S3LoadBalancerKeepalived**](S3LoadBalancerKeepalived.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OspZoneId** | Pointer to **int64** |  | [optional] 
**OssApiEnabled** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**SearchHttpsPort** | Pointer to **int64** |  | [optional] 
**SearchPort** | Pointer to **int64** |  | [optional] 
**SslCertificate** | Pointer to [**SSLCertificateNestview**](SSLCertificateNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SyncHttpsPort** | Pointer to **int64** |  | [optional] 
**SyncPort** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Vips** | Pointer to **string** |  | [optional] 
**WebServicePort** | Pointer to **int64** |  | [optional] 
**Samples** | Pointer to [**[]S3LoadBalancerStat**](S3LoadBalancerStat.md) |  | [optional] 

## Methods

### NewS3LoadBalancerRecord

`func NewS3LoadBalancerRecord() *S3LoadBalancerRecord`

NewS3LoadBalancerRecord instantiates a new S3LoadBalancerRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancerRecordWithDefaults

`func NewS3LoadBalancerRecordWithDefaults() *S3LoadBalancerRecord`

NewS3LoadBalancerRecordWithDefaults instantiates a new S3LoadBalancerRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *S3LoadBalancerRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *S3LoadBalancerRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *S3LoadBalancerRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *S3LoadBalancerRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *S3LoadBalancerRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *S3LoadBalancerRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *S3LoadBalancerRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *S3LoadBalancerRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *S3LoadBalancerRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *S3LoadBalancerRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *S3LoadBalancerRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *S3LoadBalancerRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *S3LoadBalancerRecord) GetGroup() S3LoadBalancerGroupNestview`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *S3LoadBalancerRecord) GetGroupOk() (*S3LoadBalancerGroupNestview, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *S3LoadBalancerRecord) SetGroup(v S3LoadBalancerGroupNestview)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *S3LoadBalancerRecord) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHost

`func (o *S3LoadBalancerRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *S3LoadBalancerRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *S3LoadBalancerRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *S3LoadBalancerRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHttp2Enabled

`func (o *S3LoadBalancerRecord) GetHttp2Enabled() bool`

GetHttp2Enabled returns the Http2Enabled field if non-nil, zero value otherwise.

### GetHttp2EnabledOk

`func (o *S3LoadBalancerRecord) GetHttp2EnabledOk() (*bool, bool)`

GetHttp2EnabledOk returns a tuple with the Http2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Enabled

`func (o *S3LoadBalancerRecord) SetHttp2Enabled(v bool)`

SetHttp2Enabled sets Http2Enabled field to given value.

### HasHttp2Enabled

`func (o *S3LoadBalancerRecord) HasHttp2Enabled() bool`

HasHttp2Enabled returns a boolean if a field has been set.

### GetHttpsPort

`func (o *S3LoadBalancerRecord) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *S3LoadBalancerRecord) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *S3LoadBalancerRecord) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *S3LoadBalancerRecord) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetId

`func (o *S3LoadBalancerRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3LoadBalancerRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3LoadBalancerRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *S3LoadBalancerRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeepaliveds

`func (o *S3LoadBalancerRecord) GetKeepaliveds() []S3LoadBalancerKeepalived`

GetKeepaliveds returns the Keepaliveds field if non-nil, zero value otherwise.

### GetKeepalivedsOk

`func (o *S3LoadBalancerRecord) GetKeepalivedsOk() (*[]S3LoadBalancerKeepalived, bool)`

GetKeepalivedsOk returns a tuple with the Keepaliveds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveds

`func (o *S3LoadBalancerRecord) SetKeepaliveds(v []S3LoadBalancerKeepalived)`

SetKeepaliveds sets Keepaliveds field to given value.

### HasKeepaliveds

`func (o *S3LoadBalancerRecord) HasKeepaliveds() bool`

HasKeepaliveds returns a boolean if a field has been set.

### GetName

`func (o *S3LoadBalancerRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3LoadBalancerRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3LoadBalancerRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *S3LoadBalancerRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOspZoneId

`func (o *S3LoadBalancerRecord) GetOspZoneId() int64`

GetOspZoneId returns the OspZoneId field if non-nil, zero value otherwise.

### GetOspZoneIdOk

`func (o *S3LoadBalancerRecord) GetOspZoneIdOk() (*int64, bool)`

GetOspZoneIdOk returns a tuple with the OspZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspZoneId

`func (o *S3LoadBalancerRecord) SetOspZoneId(v int64)`

SetOspZoneId sets OspZoneId field to given value.

### HasOspZoneId

`func (o *S3LoadBalancerRecord) HasOspZoneId() bool`

HasOspZoneId returns a boolean if a field has been set.

### GetOssApiEnabled

`func (o *S3LoadBalancerRecord) GetOssApiEnabled() bool`

GetOssApiEnabled returns the OssApiEnabled field if non-nil, zero value otherwise.

### GetOssApiEnabledOk

`func (o *S3LoadBalancerRecord) GetOssApiEnabledOk() (*bool, bool)`

GetOssApiEnabledOk returns a tuple with the OssApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOssApiEnabled

`func (o *S3LoadBalancerRecord) SetOssApiEnabled(v bool)`

SetOssApiEnabled sets OssApiEnabled field to given value.

### HasOssApiEnabled

`func (o *S3LoadBalancerRecord) HasOssApiEnabled() bool`

HasOssApiEnabled returns a boolean if a field has been set.

### GetPort

`func (o *S3LoadBalancerRecord) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *S3LoadBalancerRecord) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *S3LoadBalancerRecord) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *S3LoadBalancerRecord) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRoles

`func (o *S3LoadBalancerRecord) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *S3LoadBalancerRecord) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *S3LoadBalancerRecord) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *S3LoadBalancerRecord) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSearchHttpsPort

`func (o *S3LoadBalancerRecord) GetSearchHttpsPort() int64`

GetSearchHttpsPort returns the SearchHttpsPort field if non-nil, zero value otherwise.

### GetSearchHttpsPortOk

`func (o *S3LoadBalancerRecord) GetSearchHttpsPortOk() (*int64, bool)`

GetSearchHttpsPortOk returns a tuple with the SearchHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchHttpsPort

`func (o *S3LoadBalancerRecord) SetSearchHttpsPort(v int64)`

SetSearchHttpsPort sets SearchHttpsPort field to given value.

### HasSearchHttpsPort

`func (o *S3LoadBalancerRecord) HasSearchHttpsPort() bool`

HasSearchHttpsPort returns a boolean if a field has been set.

### GetSearchPort

`func (o *S3LoadBalancerRecord) GetSearchPort() int64`

GetSearchPort returns the SearchPort field if non-nil, zero value otherwise.

### GetSearchPortOk

`func (o *S3LoadBalancerRecord) GetSearchPortOk() (*int64, bool)`

GetSearchPortOk returns a tuple with the SearchPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPort

`func (o *S3LoadBalancerRecord) SetSearchPort(v int64)`

SetSearchPort sets SearchPort field to given value.

### HasSearchPort

`func (o *S3LoadBalancerRecord) HasSearchPort() bool`

HasSearchPort returns a boolean if a field has been set.

### GetSslCertificate

`func (o *S3LoadBalancerRecord) GetSslCertificate() SSLCertificateNestview`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *S3LoadBalancerRecord) GetSslCertificateOk() (*SSLCertificateNestview, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *S3LoadBalancerRecord) SetSslCertificate(v SSLCertificateNestview)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *S3LoadBalancerRecord) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetStatus

`func (o *S3LoadBalancerRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *S3LoadBalancerRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *S3LoadBalancerRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *S3LoadBalancerRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyncHttpsPort

`func (o *S3LoadBalancerRecord) GetSyncHttpsPort() int64`

GetSyncHttpsPort returns the SyncHttpsPort field if non-nil, zero value otherwise.

### GetSyncHttpsPortOk

`func (o *S3LoadBalancerRecord) GetSyncHttpsPortOk() (*int64, bool)`

GetSyncHttpsPortOk returns a tuple with the SyncHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncHttpsPort

`func (o *S3LoadBalancerRecord) SetSyncHttpsPort(v int64)`

SetSyncHttpsPort sets SyncHttpsPort field to given value.

### HasSyncHttpsPort

`func (o *S3LoadBalancerRecord) HasSyncHttpsPort() bool`

HasSyncHttpsPort returns a boolean if a field has been set.

### GetSyncPort

`func (o *S3LoadBalancerRecord) GetSyncPort() int64`

GetSyncPort returns the SyncPort field if non-nil, zero value otherwise.

### GetSyncPortOk

`func (o *S3LoadBalancerRecord) GetSyncPortOk() (*int64, bool)`

GetSyncPortOk returns a tuple with the SyncPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPort

`func (o *S3LoadBalancerRecord) SetSyncPort(v int64)`

SetSyncPort sets SyncPort field to given value.

### HasSyncPort

`func (o *S3LoadBalancerRecord) HasSyncPort() bool`

HasSyncPort returns a boolean if a field has been set.

### GetUpdate

`func (o *S3LoadBalancerRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *S3LoadBalancerRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *S3LoadBalancerRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *S3LoadBalancerRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVips

`func (o *S3LoadBalancerRecord) GetVips() string`

GetVips returns the Vips field if non-nil, zero value otherwise.

### GetVipsOk

`func (o *S3LoadBalancerRecord) GetVipsOk() (*string, bool)`

GetVipsOk returns a tuple with the Vips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVips

`func (o *S3LoadBalancerRecord) SetVips(v string)`

SetVips sets Vips field to given value.

### HasVips

`func (o *S3LoadBalancerRecord) HasVips() bool`

HasVips returns a boolean if a field has been set.

### GetWebServicePort

`func (o *S3LoadBalancerRecord) GetWebServicePort() int64`

GetWebServicePort returns the WebServicePort field if non-nil, zero value otherwise.

### GetWebServicePortOk

`func (o *S3LoadBalancerRecord) GetWebServicePortOk() (*int64, bool)`

GetWebServicePortOk returns a tuple with the WebServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServicePort

`func (o *S3LoadBalancerRecord) SetWebServicePort(v int64)`

SetWebServicePort sets WebServicePort field to given value.

### HasWebServicePort

`func (o *S3LoadBalancerRecord) HasWebServicePort() bool`

HasWebServicePort returns a boolean if a field has been set.

### GetSamples

`func (o *S3LoadBalancerRecord) GetSamples() []S3LoadBalancerStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *S3LoadBalancerRecord) GetSamplesOk() (*[]S3LoadBalancerStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *S3LoadBalancerRecord) SetSamples(v []S3LoadBalancerStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *S3LoadBalancerRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


