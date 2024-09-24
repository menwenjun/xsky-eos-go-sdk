# SSLCertificateNestview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | domain | [optional] 
**ForceHttps** | Pointer to **bool** | redirect http request to https | [optional] 
**GlobalApply** | Pointer to **bool** | apply to all s3 load balancer groups | [optional] 
**Id** | Pointer to **int64** | certificate id | [optional] 
**Name** | Pointer to **string** | certificate name | [optional] 

## Methods

### NewSSLCertificateNestview

`func NewSSLCertificateNestview() *SSLCertificateNestview`

NewSSLCertificateNestview instantiates a new SSLCertificateNestview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLCertificateNestviewWithDefaults

`func NewSSLCertificateNestviewWithDefaults() *SSLCertificateNestview`

NewSSLCertificateNestviewWithDefaults instantiates a new SSLCertificateNestview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SSLCertificateNestview) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SSLCertificateNestview) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SSLCertificateNestview) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SSLCertificateNestview) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetForceHttps

`func (o *SSLCertificateNestview) GetForceHttps() bool`

GetForceHttps returns the ForceHttps field if non-nil, zero value otherwise.

### GetForceHttpsOk

`func (o *SSLCertificateNestview) GetForceHttpsOk() (*bool, bool)`

GetForceHttpsOk returns a tuple with the ForceHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHttps

`func (o *SSLCertificateNestview) SetForceHttps(v bool)`

SetForceHttps sets ForceHttps field to given value.

### HasForceHttps

`func (o *SSLCertificateNestview) HasForceHttps() bool`

HasForceHttps returns a boolean if a field has been set.

### GetGlobalApply

`func (o *SSLCertificateNestview) GetGlobalApply() bool`

GetGlobalApply returns the GlobalApply field if non-nil, zero value otherwise.

### GetGlobalApplyOk

`func (o *SSLCertificateNestview) GetGlobalApplyOk() (*bool, bool)`

GetGlobalApplyOk returns a tuple with the GlobalApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalApply

`func (o *SSLCertificateNestview) SetGlobalApply(v bool)`

SetGlobalApply sets GlobalApply field to given value.

### HasGlobalApply

`func (o *SSLCertificateNestview) HasGlobalApply() bool`

HasGlobalApply returns a boolean if a field has been set.

### GetId

`func (o *SSLCertificateNestview) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SSLCertificateNestview) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SSLCertificateNestview) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SSLCertificateNestview) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SSLCertificateNestview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SSLCertificateNestview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SSLCertificateNestview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SSLCertificateNestview) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


