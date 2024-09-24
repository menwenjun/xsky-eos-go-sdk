# SSLCertificateCreateReqCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | certificate in pem | 
**Description** | Pointer to **string** | certificate description | [optional] 
**Name** | **string** | certificate name | 
**PrivateKey** | **string** | private key in pem | 

## Methods

### NewSSLCertificateCreateReqCertificate

`func NewSSLCertificateCreateReqCertificate(certificate string, name string, privateKey string, ) *SSLCertificateCreateReqCertificate`

NewSSLCertificateCreateReqCertificate instantiates a new SSLCertificateCreateReqCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLCertificateCreateReqCertificateWithDefaults

`func NewSSLCertificateCreateReqCertificateWithDefaults() *SSLCertificateCreateReqCertificate`

NewSSLCertificateCreateReqCertificateWithDefaults instantiates a new SSLCertificateCreateReqCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *SSLCertificateCreateReqCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SSLCertificateCreateReqCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SSLCertificateCreateReqCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetDescription

`func (o *SSLCertificateCreateReqCertificate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SSLCertificateCreateReqCertificate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SSLCertificateCreateReqCertificate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SSLCertificateCreateReqCertificate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SSLCertificateCreateReqCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SSLCertificateCreateReqCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SSLCertificateCreateReqCertificate) SetName(v string)`

SetName sets Name field to given value.


### GetPrivateKey

`func (o *SSLCertificateCreateReqCertificate) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SSLCertificateCreateReqCertificate) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SSLCertificateCreateReqCertificate) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


