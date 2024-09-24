# S3LoadBalancerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Http2Enabled** | Pointer to **bool** |  | [optional] 
**HttpsPort** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OspZoneId** | Pointer to **int64** |  | [optional] 
**OssApiEnabled** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**SearchHttpsPort** | Pointer to **int64** |  | [optional] 
**SearchPort** | Pointer to **int64** |  | [optional] 
**SslCertificate** | Pointer to [**SSLCertificateNestview**](SSLCertificateNestview.md) |  | [optional] 
**SslCertificates** | Pointer to [**[]SSLCertificateNestview**](SSLCertificateNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SyncHttpsPort** | Pointer to **int64** |  | [optional] 
**SyncPort** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**WebServiceConfig** | Pointer to [**S3LbGroupWebServiceConfig**](S3LbGroupWebServiceConfig.md) |  | [optional] 
**WebServicePort** | Pointer to **int64** |  | [optional] 

## Methods

### NewS3LoadBalancerGroup

`func NewS3LoadBalancerGroup() *S3LoadBalancerGroup`

NewS3LoadBalancerGroup instantiates a new S3LoadBalancerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancerGroupWithDefaults

`func NewS3LoadBalancerGroupWithDefaults() *S3LoadBalancerGroup`

NewS3LoadBalancerGroupWithDefaults instantiates a new S3LoadBalancerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *S3LoadBalancerGroup) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *S3LoadBalancerGroup) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *S3LoadBalancerGroup) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *S3LoadBalancerGroup) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *S3LoadBalancerGroup) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *S3LoadBalancerGroup) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *S3LoadBalancerGroup) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *S3LoadBalancerGroup) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *S3LoadBalancerGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *S3LoadBalancerGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *S3LoadBalancerGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *S3LoadBalancerGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *S3LoadBalancerGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *S3LoadBalancerGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *S3LoadBalancerGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *S3LoadBalancerGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHttp2Enabled

`func (o *S3LoadBalancerGroup) GetHttp2Enabled() bool`

GetHttp2Enabled returns the Http2Enabled field if non-nil, zero value otherwise.

### GetHttp2EnabledOk

`func (o *S3LoadBalancerGroup) GetHttp2EnabledOk() (*bool, bool)`

GetHttp2EnabledOk returns a tuple with the Http2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Enabled

`func (o *S3LoadBalancerGroup) SetHttp2Enabled(v bool)`

SetHttp2Enabled sets Http2Enabled field to given value.

### HasHttp2Enabled

`func (o *S3LoadBalancerGroup) HasHttp2Enabled() bool`

HasHttp2Enabled returns a boolean if a field has been set.

### GetHttpsPort

`func (o *S3LoadBalancerGroup) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *S3LoadBalancerGroup) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *S3LoadBalancerGroup) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *S3LoadBalancerGroup) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetId

`func (o *S3LoadBalancerGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3LoadBalancerGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3LoadBalancerGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *S3LoadBalancerGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *S3LoadBalancerGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3LoadBalancerGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3LoadBalancerGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *S3LoadBalancerGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOspZoneId

`func (o *S3LoadBalancerGroup) GetOspZoneId() int64`

GetOspZoneId returns the OspZoneId field if non-nil, zero value otherwise.

### GetOspZoneIdOk

`func (o *S3LoadBalancerGroup) GetOspZoneIdOk() (*int64, bool)`

GetOspZoneIdOk returns a tuple with the OspZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspZoneId

`func (o *S3LoadBalancerGroup) SetOspZoneId(v int64)`

SetOspZoneId sets OspZoneId field to given value.

### HasOspZoneId

`func (o *S3LoadBalancerGroup) HasOspZoneId() bool`

HasOspZoneId returns a boolean if a field has been set.

### GetOssApiEnabled

`func (o *S3LoadBalancerGroup) GetOssApiEnabled() bool`

GetOssApiEnabled returns the OssApiEnabled field if non-nil, zero value otherwise.

### GetOssApiEnabledOk

`func (o *S3LoadBalancerGroup) GetOssApiEnabledOk() (*bool, bool)`

GetOssApiEnabledOk returns a tuple with the OssApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOssApiEnabled

`func (o *S3LoadBalancerGroup) SetOssApiEnabled(v bool)`

SetOssApiEnabled sets OssApiEnabled field to given value.

### HasOssApiEnabled

`func (o *S3LoadBalancerGroup) HasOssApiEnabled() bool`

HasOssApiEnabled returns a boolean if a field has been set.

### GetPort

`func (o *S3LoadBalancerGroup) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *S3LoadBalancerGroup) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *S3LoadBalancerGroup) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *S3LoadBalancerGroup) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRoles

`func (o *S3LoadBalancerGroup) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *S3LoadBalancerGroup) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *S3LoadBalancerGroup) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *S3LoadBalancerGroup) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSearchHttpsPort

`func (o *S3LoadBalancerGroup) GetSearchHttpsPort() int64`

GetSearchHttpsPort returns the SearchHttpsPort field if non-nil, zero value otherwise.

### GetSearchHttpsPortOk

`func (o *S3LoadBalancerGroup) GetSearchHttpsPortOk() (*int64, bool)`

GetSearchHttpsPortOk returns a tuple with the SearchHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchHttpsPort

`func (o *S3LoadBalancerGroup) SetSearchHttpsPort(v int64)`

SetSearchHttpsPort sets SearchHttpsPort field to given value.

### HasSearchHttpsPort

`func (o *S3LoadBalancerGroup) HasSearchHttpsPort() bool`

HasSearchHttpsPort returns a boolean if a field has been set.

### GetSearchPort

`func (o *S3LoadBalancerGroup) GetSearchPort() int64`

GetSearchPort returns the SearchPort field if non-nil, zero value otherwise.

### GetSearchPortOk

`func (o *S3LoadBalancerGroup) GetSearchPortOk() (*int64, bool)`

GetSearchPortOk returns a tuple with the SearchPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPort

`func (o *S3LoadBalancerGroup) SetSearchPort(v int64)`

SetSearchPort sets SearchPort field to given value.

### HasSearchPort

`func (o *S3LoadBalancerGroup) HasSearchPort() bool`

HasSearchPort returns a boolean if a field has been set.

### GetSslCertificate

`func (o *S3LoadBalancerGroup) GetSslCertificate() SSLCertificateNestview`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *S3LoadBalancerGroup) GetSslCertificateOk() (*SSLCertificateNestview, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *S3LoadBalancerGroup) SetSslCertificate(v SSLCertificateNestview)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *S3LoadBalancerGroup) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetSslCertificates

`func (o *S3LoadBalancerGroup) GetSslCertificates() []SSLCertificateNestview`

GetSslCertificates returns the SslCertificates field if non-nil, zero value otherwise.

### GetSslCertificatesOk

`func (o *S3LoadBalancerGroup) GetSslCertificatesOk() (*[]SSLCertificateNestview, bool)`

GetSslCertificatesOk returns a tuple with the SslCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificates

`func (o *S3LoadBalancerGroup) SetSslCertificates(v []SSLCertificateNestview)`

SetSslCertificates sets SslCertificates field to given value.

### HasSslCertificates

`func (o *S3LoadBalancerGroup) HasSslCertificates() bool`

HasSslCertificates returns a boolean if a field has been set.

### GetStatus

`func (o *S3LoadBalancerGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *S3LoadBalancerGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *S3LoadBalancerGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *S3LoadBalancerGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyncHttpsPort

`func (o *S3LoadBalancerGroup) GetSyncHttpsPort() int64`

GetSyncHttpsPort returns the SyncHttpsPort field if non-nil, zero value otherwise.

### GetSyncHttpsPortOk

`func (o *S3LoadBalancerGroup) GetSyncHttpsPortOk() (*int64, bool)`

GetSyncHttpsPortOk returns a tuple with the SyncHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncHttpsPort

`func (o *S3LoadBalancerGroup) SetSyncHttpsPort(v int64)`

SetSyncHttpsPort sets SyncHttpsPort field to given value.

### HasSyncHttpsPort

`func (o *S3LoadBalancerGroup) HasSyncHttpsPort() bool`

HasSyncHttpsPort returns a boolean if a field has been set.

### GetSyncPort

`func (o *S3LoadBalancerGroup) GetSyncPort() int64`

GetSyncPort returns the SyncPort field if non-nil, zero value otherwise.

### GetSyncPortOk

`func (o *S3LoadBalancerGroup) GetSyncPortOk() (*int64, bool)`

GetSyncPortOk returns a tuple with the SyncPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPort

`func (o *S3LoadBalancerGroup) SetSyncPort(v int64)`

SetSyncPort sets SyncPort field to given value.

### HasSyncPort

`func (o *S3LoadBalancerGroup) HasSyncPort() bool`

HasSyncPort returns a boolean if a field has been set.

### GetUpdate

`func (o *S3LoadBalancerGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *S3LoadBalancerGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *S3LoadBalancerGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *S3LoadBalancerGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetWebServiceConfig

`func (o *S3LoadBalancerGroup) GetWebServiceConfig() S3LbGroupWebServiceConfig`

GetWebServiceConfig returns the WebServiceConfig field if non-nil, zero value otherwise.

### GetWebServiceConfigOk

`func (o *S3LoadBalancerGroup) GetWebServiceConfigOk() (*S3LbGroupWebServiceConfig, bool)`

GetWebServiceConfigOk returns a tuple with the WebServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceConfig

`func (o *S3LoadBalancerGroup) SetWebServiceConfig(v S3LbGroupWebServiceConfig)`

SetWebServiceConfig sets WebServiceConfig field to given value.

### HasWebServiceConfig

`func (o *S3LoadBalancerGroup) HasWebServiceConfig() bool`

HasWebServiceConfig returns a boolean if a field has been set.

### GetWebServicePort

`func (o *S3LoadBalancerGroup) GetWebServicePort() int64`

GetWebServicePort returns the WebServicePort field if non-nil, zero value otherwise.

### GetWebServicePortOk

`func (o *S3LoadBalancerGroup) GetWebServicePortOk() (*int64, bool)`

GetWebServicePortOk returns a tuple with the WebServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServicePort

`func (o *S3LoadBalancerGroup) SetWebServicePort(v int64)`

SetWebServicePort sets WebServicePort field to given value.

### HasWebServicePort

`func (o *S3LoadBalancerGroup) HasWebServicePort() bool`

HasWebServicePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


