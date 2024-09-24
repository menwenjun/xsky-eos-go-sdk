# SSLCertificateUpdateReqCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | certificate description | [optional] 
**Enabled** | Pointer to **bool** | enabled or not | [optional] 
**ForceHttps** | Pointer to **bool** | redirect http request to https | [optional] 
**GlobalApply** | Pointer to **bool** | global_apply or not | [optional] 
**Name** | Pointer to **string** | certificate name | [optional] 
**S3LoadBalancerGroupIds** | Pointer to **[]int64** | s3 load balencer ids | [optional] 
**Type** | Pointer to **string** | applied type: admin, s3, dfs_s3 | [optional] 

## Methods

### NewSSLCertificateUpdateReqCertificate

`func NewSSLCertificateUpdateReqCertificate() *SSLCertificateUpdateReqCertificate`

NewSSLCertificateUpdateReqCertificate instantiates a new SSLCertificateUpdateReqCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLCertificateUpdateReqCertificateWithDefaults

`func NewSSLCertificateUpdateReqCertificateWithDefaults() *SSLCertificateUpdateReqCertificate`

NewSSLCertificateUpdateReqCertificateWithDefaults instantiates a new SSLCertificateUpdateReqCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SSLCertificateUpdateReqCertificate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SSLCertificateUpdateReqCertificate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SSLCertificateUpdateReqCertificate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SSLCertificateUpdateReqCertificate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SSLCertificateUpdateReqCertificate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SSLCertificateUpdateReqCertificate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SSLCertificateUpdateReqCertificate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SSLCertificateUpdateReqCertificate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForceHttps

`func (o *SSLCertificateUpdateReqCertificate) GetForceHttps() bool`

GetForceHttps returns the ForceHttps field if non-nil, zero value otherwise.

### GetForceHttpsOk

`func (o *SSLCertificateUpdateReqCertificate) GetForceHttpsOk() (*bool, bool)`

GetForceHttpsOk returns a tuple with the ForceHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHttps

`func (o *SSLCertificateUpdateReqCertificate) SetForceHttps(v bool)`

SetForceHttps sets ForceHttps field to given value.

### HasForceHttps

`func (o *SSLCertificateUpdateReqCertificate) HasForceHttps() bool`

HasForceHttps returns a boolean if a field has been set.

### GetGlobalApply

`func (o *SSLCertificateUpdateReqCertificate) GetGlobalApply() bool`

GetGlobalApply returns the GlobalApply field if non-nil, zero value otherwise.

### GetGlobalApplyOk

`func (o *SSLCertificateUpdateReqCertificate) GetGlobalApplyOk() (*bool, bool)`

GetGlobalApplyOk returns a tuple with the GlobalApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalApply

`func (o *SSLCertificateUpdateReqCertificate) SetGlobalApply(v bool)`

SetGlobalApply sets GlobalApply field to given value.

### HasGlobalApply

`func (o *SSLCertificateUpdateReqCertificate) HasGlobalApply() bool`

HasGlobalApply returns a boolean if a field has been set.

### GetName

`func (o *SSLCertificateUpdateReqCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SSLCertificateUpdateReqCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SSLCertificateUpdateReqCertificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SSLCertificateUpdateReqCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetS3LoadBalancerGroupIds

`func (o *SSLCertificateUpdateReqCertificate) GetS3LoadBalancerGroupIds() []int64`

GetS3LoadBalancerGroupIds returns the S3LoadBalancerGroupIds field if non-nil, zero value otherwise.

### GetS3LoadBalancerGroupIdsOk

`func (o *SSLCertificateUpdateReqCertificate) GetS3LoadBalancerGroupIdsOk() (*[]int64, bool)`

GetS3LoadBalancerGroupIdsOk returns a tuple with the S3LoadBalancerGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3LoadBalancerGroupIds

`func (o *SSLCertificateUpdateReqCertificate) SetS3LoadBalancerGroupIds(v []int64)`

SetS3LoadBalancerGroupIds sets S3LoadBalancerGroupIds field to given value.

### HasS3LoadBalancerGroupIds

`func (o *SSLCertificateUpdateReqCertificate) HasS3LoadBalancerGroupIds() bool`

HasS3LoadBalancerGroupIds returns a boolean if a field has been set.

### GetType

`func (o *SSLCertificateUpdateReqCertificate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SSLCertificateUpdateReqCertificate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SSLCertificateUpdateReqCertificate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SSLCertificateUpdateReqCertificate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


