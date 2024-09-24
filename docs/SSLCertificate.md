# SSLCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** | created time of certificate | [optional] 
**Description** | Pointer to **string** | certificate description | [optional] 
**Domain** | Pointer to **string** | domain | [optional] 
**Enabled** | Pointer to **bool** | enabled or not | [optional] 
**ForceHttps** | Pointer to **bool** | redirect http request to https | [optional] 
**GlobalApply** | Pointer to **bool** | apply to all s3 load balancer groups | [optional] 
**Id** | Pointer to **int64** | certificate id | [optional] 
**Issuer** | Pointer to **map[string]interface{}** | issuer info | [optional] 
**Name** | Pointer to **string** | certificate name | [optional] 
**NotAfter** | Pointer to **time.Time** | validity is not after the time | [optional] 
**NotBefore** | Pointer to **time.Time** | validity is not before the time | [optional] 
**PermittedDnsDomains** | Pointer to **[]map[string]interface{}** | permitted dns domains | [optional] 
**PublicKeyAlgorithm** | Pointer to **string** | public key algorithm | [optional] 
**RawCertificate** | Pointer to **string** | public certificate | [optional] 
**S3LoadBalancerGroups** | Pointer to [**[]S3LoadBalancerGroupNestview**](S3LoadBalancerGroupNestview.md) | s3 load balancer groups | [optional] 
**SignatureAlgorithm** | Pointer to **string** | signature algorithm | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **map[string]interface{}** | subject info | [optional] 
**Type** | Pointer to **string** | applied type: admin, s3, dfs_s3 | [optional] 
**Update** | Pointer to **time.Time** | updated time of certificate | [optional] 
**Version** | Pointer to **int64** | certificate version | [optional] 

## Methods

### NewSSLCertificate

`func NewSSLCertificate() *SSLCertificate`

NewSSLCertificate instantiates a new SSLCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLCertificateWithDefaults

`func NewSSLCertificateWithDefaults() *SSLCertificate`

NewSSLCertificateWithDefaults instantiates a new SSLCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *SSLCertificate) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *SSLCertificate) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *SSLCertificate) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *SSLCertificate) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *SSLCertificate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SSLCertificate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SSLCertificate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SSLCertificate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *SSLCertificate) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SSLCertificate) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SSLCertificate) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SSLCertificate) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEnabled

`func (o *SSLCertificate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SSLCertificate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SSLCertificate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SSLCertificate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForceHttps

`func (o *SSLCertificate) GetForceHttps() bool`

GetForceHttps returns the ForceHttps field if non-nil, zero value otherwise.

### GetForceHttpsOk

`func (o *SSLCertificate) GetForceHttpsOk() (*bool, bool)`

GetForceHttpsOk returns a tuple with the ForceHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHttps

`func (o *SSLCertificate) SetForceHttps(v bool)`

SetForceHttps sets ForceHttps field to given value.

### HasForceHttps

`func (o *SSLCertificate) HasForceHttps() bool`

HasForceHttps returns a boolean if a field has been set.

### GetGlobalApply

`func (o *SSLCertificate) GetGlobalApply() bool`

GetGlobalApply returns the GlobalApply field if non-nil, zero value otherwise.

### GetGlobalApplyOk

`func (o *SSLCertificate) GetGlobalApplyOk() (*bool, bool)`

GetGlobalApplyOk returns a tuple with the GlobalApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalApply

`func (o *SSLCertificate) SetGlobalApply(v bool)`

SetGlobalApply sets GlobalApply field to given value.

### HasGlobalApply

`func (o *SSLCertificate) HasGlobalApply() bool`

HasGlobalApply returns a boolean if a field has been set.

### GetId

`func (o *SSLCertificate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SSLCertificate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SSLCertificate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SSLCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuer

`func (o *SSLCertificate) GetIssuer() map[string]interface{}`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SSLCertificate) GetIssuerOk() (*map[string]interface{}, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SSLCertificate) SetIssuer(v map[string]interface{})`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SSLCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetName

`func (o *SSLCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SSLCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SSLCertificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SSLCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotAfter

`func (o *SSLCertificate) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *SSLCertificate) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *SSLCertificate) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *SSLCertificate) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *SSLCertificate) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *SSLCertificate) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *SSLCertificate) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *SSLCertificate) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetPermittedDnsDomains

`func (o *SSLCertificate) GetPermittedDnsDomains() []map[string]interface{}`

GetPermittedDnsDomains returns the PermittedDnsDomains field if non-nil, zero value otherwise.

### GetPermittedDnsDomainsOk

`func (o *SSLCertificate) GetPermittedDnsDomainsOk() (*[]map[string]interface{}, bool)`

GetPermittedDnsDomainsOk returns a tuple with the PermittedDnsDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedDnsDomains

`func (o *SSLCertificate) SetPermittedDnsDomains(v []map[string]interface{})`

SetPermittedDnsDomains sets PermittedDnsDomains field to given value.

### HasPermittedDnsDomains

`func (o *SSLCertificate) HasPermittedDnsDomains() bool`

HasPermittedDnsDomains returns a boolean if a field has been set.

### GetPublicKeyAlgorithm

`func (o *SSLCertificate) GetPublicKeyAlgorithm() string`

GetPublicKeyAlgorithm returns the PublicKeyAlgorithm field if non-nil, zero value otherwise.

### GetPublicKeyAlgorithmOk

`func (o *SSLCertificate) GetPublicKeyAlgorithmOk() (*string, bool)`

GetPublicKeyAlgorithmOk returns a tuple with the PublicKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyAlgorithm

`func (o *SSLCertificate) SetPublicKeyAlgorithm(v string)`

SetPublicKeyAlgorithm sets PublicKeyAlgorithm field to given value.

### HasPublicKeyAlgorithm

`func (o *SSLCertificate) HasPublicKeyAlgorithm() bool`

HasPublicKeyAlgorithm returns a boolean if a field has been set.

### GetRawCertificate

`func (o *SSLCertificate) GetRawCertificate() string`

GetRawCertificate returns the RawCertificate field if non-nil, zero value otherwise.

### GetRawCertificateOk

`func (o *SSLCertificate) GetRawCertificateOk() (*string, bool)`

GetRawCertificateOk returns a tuple with the RawCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawCertificate

`func (o *SSLCertificate) SetRawCertificate(v string)`

SetRawCertificate sets RawCertificate field to given value.

### HasRawCertificate

`func (o *SSLCertificate) HasRawCertificate() bool`

HasRawCertificate returns a boolean if a field has been set.

### GetS3LoadBalancerGroups

`func (o *SSLCertificate) GetS3LoadBalancerGroups() []S3LoadBalancerGroupNestview`

GetS3LoadBalancerGroups returns the S3LoadBalancerGroups field if non-nil, zero value otherwise.

### GetS3LoadBalancerGroupsOk

`func (o *SSLCertificate) GetS3LoadBalancerGroupsOk() (*[]S3LoadBalancerGroupNestview, bool)`

GetS3LoadBalancerGroupsOk returns a tuple with the S3LoadBalancerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3LoadBalancerGroups

`func (o *SSLCertificate) SetS3LoadBalancerGroups(v []S3LoadBalancerGroupNestview)`

SetS3LoadBalancerGroups sets S3LoadBalancerGroups field to given value.

### HasS3LoadBalancerGroups

`func (o *SSLCertificate) HasS3LoadBalancerGroups() bool`

HasS3LoadBalancerGroups returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *SSLCertificate) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *SSLCertificate) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *SSLCertificate) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *SSLCertificate) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetStatus

`func (o *SSLCertificate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SSLCertificate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SSLCertificate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SSLCertificate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *SSLCertificate) GetSubject() map[string]interface{}`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SSLCertificate) GetSubjectOk() (*map[string]interface{}, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SSLCertificate) SetSubject(v map[string]interface{})`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SSLCertificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetType

`func (o *SSLCertificate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SSLCertificate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SSLCertificate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SSLCertificate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *SSLCertificate) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *SSLCertificate) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *SSLCertificate) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *SSLCertificate) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVersion

`func (o *SSLCertificate) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SSLCertificate) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SSLCertificate) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SSLCertificate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


