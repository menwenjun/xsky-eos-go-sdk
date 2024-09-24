# DpSiteCreateReqSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | site address | [optional] 
**Config** | [**DpSiteConfig**](DpSiteConfig.md) |  | 
**CryptoKeyId** | Pointer to **int64** | crypto key | [optional] 
**Name** | **string** | site name | 
**Platform** | **string** | platform of site | 
**ProtectionType** | **string** | protection type | 
**RemoteClusterId** | Pointer to **int64** | remote cluster | [optional] 
**Service** | **string** | service of site | 

## Methods

### NewDpSiteCreateReqSite

`func NewDpSiteCreateReqSite(config DpSiteConfig, name string, platform string, protectionType string, service string, ) *DpSiteCreateReqSite`

NewDpSiteCreateReqSite instantiates a new DpSiteCreateReqSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpSiteCreateReqSiteWithDefaults

`func NewDpSiteCreateReqSiteWithDefaults() *DpSiteCreateReqSite`

NewDpSiteCreateReqSiteWithDefaults instantiates a new DpSiteCreateReqSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DpSiteCreateReqSite) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DpSiteCreateReqSite) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DpSiteCreateReqSite) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DpSiteCreateReqSite) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetConfig

`func (o *DpSiteCreateReqSite) GetConfig() DpSiteConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DpSiteCreateReqSite) GetConfigOk() (*DpSiteConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DpSiteCreateReqSite) SetConfig(v DpSiteConfig)`

SetConfig sets Config field to given value.


### GetCryptoKeyId

`func (o *DpSiteCreateReqSite) GetCryptoKeyId() int64`

GetCryptoKeyId returns the CryptoKeyId field if non-nil, zero value otherwise.

### GetCryptoKeyIdOk

`func (o *DpSiteCreateReqSite) GetCryptoKeyIdOk() (*int64, bool)`

GetCryptoKeyIdOk returns a tuple with the CryptoKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoKeyId

`func (o *DpSiteCreateReqSite) SetCryptoKeyId(v int64)`

SetCryptoKeyId sets CryptoKeyId field to given value.

### HasCryptoKeyId

`func (o *DpSiteCreateReqSite) HasCryptoKeyId() bool`

HasCryptoKeyId returns a boolean if a field has been set.

### GetName

`func (o *DpSiteCreateReqSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpSiteCreateReqSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpSiteCreateReqSite) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *DpSiteCreateReqSite) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DpSiteCreateReqSite) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DpSiteCreateReqSite) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetProtectionType

`func (o *DpSiteCreateReqSite) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *DpSiteCreateReqSite) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *DpSiteCreateReqSite) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.


### GetRemoteClusterId

`func (o *DpSiteCreateReqSite) GetRemoteClusterId() int64`

GetRemoteClusterId returns the RemoteClusterId field if non-nil, zero value otherwise.

### GetRemoteClusterIdOk

`func (o *DpSiteCreateReqSite) GetRemoteClusterIdOk() (*int64, bool)`

GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterId

`func (o *DpSiteCreateReqSite) SetRemoteClusterId(v int64)`

SetRemoteClusterId sets RemoteClusterId field to given value.

### HasRemoteClusterId

`func (o *DpSiteCreateReqSite) HasRemoteClusterId() bool`

HasRemoteClusterId returns a boolean if a field has been set.

### GetService

`func (o *DpSiteCreateReqSite) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DpSiteCreateReqSite) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DpSiteCreateReqSite) SetService(v string)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


