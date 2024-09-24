# S3LoadBalancerGroupCreateReqGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | group description | [optional] 
**Http2Enabled** | Pointer to **bool** | group access http2 enabled | [optional] 
**HttpsPort** | Pointer to **int64** | group access https port | [optional] 
**Name** | **string** | group name | 
**OspZoneId** | Pointer to **int64** | osp zone id | [optional] 
**OssApiEnabled** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **int64** | group access http port | [optional] 
**Roles** | Pointer to **[]string** | group roles | [optional] 
**S3LoadBalancers** | [**[]S3LoadBalancerGroupCreateReqGroupLoadBalancersElt**](S3LoadBalancerGroupCreateReqGroupLoadBalancersElt.md) | s3 load balancers | 
**SearchHttpsPort** | Pointer to **int64** | group search https port | [optional] 
**SearchPort** | Pointer to **int64** | group search http port | [optional] 
**SyncHttpsPort** | Pointer to **int64** | group sync https port | [optional] 
**SyncPort** | Pointer to **int64** | group sync http port | [optional] 
**WebServiceConfig** | Pointer to [**S3LbGroupWebServiceConfig**](S3LbGroupWebServiceConfig.md) |  | [optional] 
**WebServicePort** | Pointer to **int64** | group web service http port | [optional] 

## Methods

### NewS3LoadBalancerGroupCreateReqGroup

`func NewS3LoadBalancerGroupCreateReqGroup(name string, s3LoadBalancers []S3LoadBalancerGroupCreateReqGroupLoadBalancersElt, ) *S3LoadBalancerGroupCreateReqGroup`

NewS3LoadBalancerGroupCreateReqGroup instantiates a new S3LoadBalancerGroupCreateReqGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancerGroupCreateReqGroupWithDefaults

`func NewS3LoadBalancerGroupCreateReqGroupWithDefaults() *S3LoadBalancerGroupCreateReqGroup`

NewS3LoadBalancerGroupCreateReqGroupWithDefaults instantiates a new S3LoadBalancerGroupCreateReqGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *S3LoadBalancerGroupCreateReqGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *S3LoadBalancerGroupCreateReqGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *S3LoadBalancerGroupCreateReqGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHttp2Enabled

`func (o *S3LoadBalancerGroupCreateReqGroup) GetHttp2Enabled() bool`

GetHttp2Enabled returns the Http2Enabled field if non-nil, zero value otherwise.

### GetHttp2EnabledOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetHttp2EnabledOk() (*bool, bool)`

GetHttp2EnabledOk returns a tuple with the Http2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Enabled

`func (o *S3LoadBalancerGroupCreateReqGroup) SetHttp2Enabled(v bool)`

SetHttp2Enabled sets Http2Enabled field to given value.

### HasHttp2Enabled

`func (o *S3LoadBalancerGroupCreateReqGroup) HasHttp2Enabled() bool`

HasHttp2Enabled returns a boolean if a field has been set.

### GetHttpsPort

`func (o *S3LoadBalancerGroupCreateReqGroup) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *S3LoadBalancerGroupCreateReqGroup) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *S3LoadBalancerGroupCreateReqGroup) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetName

`func (o *S3LoadBalancerGroupCreateReqGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3LoadBalancerGroupCreateReqGroup) SetName(v string)`

SetName sets Name field to given value.


### GetOspZoneId

`func (o *S3LoadBalancerGroupCreateReqGroup) GetOspZoneId() int64`

GetOspZoneId returns the OspZoneId field if non-nil, zero value otherwise.

### GetOspZoneIdOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetOspZoneIdOk() (*int64, bool)`

GetOspZoneIdOk returns a tuple with the OspZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspZoneId

`func (o *S3LoadBalancerGroupCreateReqGroup) SetOspZoneId(v int64)`

SetOspZoneId sets OspZoneId field to given value.

### HasOspZoneId

`func (o *S3LoadBalancerGroupCreateReqGroup) HasOspZoneId() bool`

HasOspZoneId returns a boolean if a field has been set.

### GetOssApiEnabled

`func (o *S3LoadBalancerGroupCreateReqGroup) GetOssApiEnabled() bool`

GetOssApiEnabled returns the OssApiEnabled field if non-nil, zero value otherwise.

### GetOssApiEnabledOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetOssApiEnabledOk() (*bool, bool)`

GetOssApiEnabledOk returns a tuple with the OssApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOssApiEnabled

`func (o *S3LoadBalancerGroupCreateReqGroup) SetOssApiEnabled(v bool)`

SetOssApiEnabled sets OssApiEnabled field to given value.

### HasOssApiEnabled

`func (o *S3LoadBalancerGroupCreateReqGroup) HasOssApiEnabled() bool`

HasOssApiEnabled returns a boolean if a field has been set.

### GetPort

`func (o *S3LoadBalancerGroupCreateReqGroup) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *S3LoadBalancerGroupCreateReqGroup) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *S3LoadBalancerGroupCreateReqGroup) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRoles

`func (o *S3LoadBalancerGroupCreateReqGroup) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *S3LoadBalancerGroupCreateReqGroup) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *S3LoadBalancerGroupCreateReqGroup) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetS3LoadBalancers

`func (o *S3LoadBalancerGroupCreateReqGroup) GetS3LoadBalancers() []S3LoadBalancerGroupCreateReqGroupLoadBalancersElt`

GetS3LoadBalancers returns the S3LoadBalancers field if non-nil, zero value otherwise.

### GetS3LoadBalancersOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetS3LoadBalancersOk() (*[]S3LoadBalancerGroupCreateReqGroupLoadBalancersElt, bool)`

GetS3LoadBalancersOk returns a tuple with the S3LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3LoadBalancers

`func (o *S3LoadBalancerGroupCreateReqGroup) SetS3LoadBalancers(v []S3LoadBalancerGroupCreateReqGroupLoadBalancersElt)`

SetS3LoadBalancers sets S3LoadBalancers field to given value.


### GetSearchHttpsPort

`func (o *S3LoadBalancerGroupCreateReqGroup) GetSearchHttpsPort() int64`

GetSearchHttpsPort returns the SearchHttpsPort field if non-nil, zero value otherwise.

### GetSearchHttpsPortOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetSearchHttpsPortOk() (*int64, bool)`

GetSearchHttpsPortOk returns a tuple with the SearchHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchHttpsPort

`func (o *S3LoadBalancerGroupCreateReqGroup) SetSearchHttpsPort(v int64)`

SetSearchHttpsPort sets SearchHttpsPort field to given value.

### HasSearchHttpsPort

`func (o *S3LoadBalancerGroupCreateReqGroup) HasSearchHttpsPort() bool`

HasSearchHttpsPort returns a boolean if a field has been set.

### GetSearchPort

`func (o *S3LoadBalancerGroupCreateReqGroup) GetSearchPort() int64`

GetSearchPort returns the SearchPort field if non-nil, zero value otherwise.

### GetSearchPortOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetSearchPortOk() (*int64, bool)`

GetSearchPortOk returns a tuple with the SearchPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPort

`func (o *S3LoadBalancerGroupCreateReqGroup) SetSearchPort(v int64)`

SetSearchPort sets SearchPort field to given value.

### HasSearchPort

`func (o *S3LoadBalancerGroupCreateReqGroup) HasSearchPort() bool`

HasSearchPort returns a boolean if a field has been set.

### GetSyncHttpsPort

`func (o *S3LoadBalancerGroupCreateReqGroup) GetSyncHttpsPort() int64`

GetSyncHttpsPort returns the SyncHttpsPort field if non-nil, zero value otherwise.

### GetSyncHttpsPortOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetSyncHttpsPortOk() (*int64, bool)`

GetSyncHttpsPortOk returns a tuple with the SyncHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncHttpsPort

`func (o *S3LoadBalancerGroupCreateReqGroup) SetSyncHttpsPort(v int64)`

SetSyncHttpsPort sets SyncHttpsPort field to given value.

### HasSyncHttpsPort

`func (o *S3LoadBalancerGroupCreateReqGroup) HasSyncHttpsPort() bool`

HasSyncHttpsPort returns a boolean if a field has been set.

### GetSyncPort

`func (o *S3LoadBalancerGroupCreateReqGroup) GetSyncPort() int64`

GetSyncPort returns the SyncPort field if non-nil, zero value otherwise.

### GetSyncPortOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetSyncPortOk() (*int64, bool)`

GetSyncPortOk returns a tuple with the SyncPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPort

`func (o *S3LoadBalancerGroupCreateReqGroup) SetSyncPort(v int64)`

SetSyncPort sets SyncPort field to given value.

### HasSyncPort

`func (o *S3LoadBalancerGroupCreateReqGroup) HasSyncPort() bool`

HasSyncPort returns a boolean if a field has been set.

### GetWebServiceConfig

`func (o *S3LoadBalancerGroupCreateReqGroup) GetWebServiceConfig() S3LbGroupWebServiceConfig`

GetWebServiceConfig returns the WebServiceConfig field if non-nil, zero value otherwise.

### GetWebServiceConfigOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetWebServiceConfigOk() (*S3LbGroupWebServiceConfig, bool)`

GetWebServiceConfigOk returns a tuple with the WebServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceConfig

`func (o *S3LoadBalancerGroupCreateReqGroup) SetWebServiceConfig(v S3LbGroupWebServiceConfig)`

SetWebServiceConfig sets WebServiceConfig field to given value.

### HasWebServiceConfig

`func (o *S3LoadBalancerGroupCreateReqGroup) HasWebServiceConfig() bool`

HasWebServiceConfig returns a boolean if a field has been set.

### GetWebServicePort

`func (o *S3LoadBalancerGroupCreateReqGroup) GetWebServicePort() int64`

GetWebServicePort returns the WebServicePort field if non-nil, zero value otherwise.

### GetWebServicePortOk

`func (o *S3LoadBalancerGroupCreateReqGroup) GetWebServicePortOk() (*int64, bool)`

GetWebServicePortOk returns a tuple with the WebServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServicePort

`func (o *S3LoadBalancerGroupCreateReqGroup) SetWebServicePort(v int64)`

SetWebServicePort sets WebServicePort field to given value.

### HasWebServicePort

`func (o *S3LoadBalancerGroupCreateReqGroup) HasWebServicePort() bool`

HasWebServicePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


