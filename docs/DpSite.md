# DpSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Config** | Pointer to [**DpSiteConfig**](DpSiteConfig.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**CryptoKey** | Pointer to [**CryptoKeyNestview**](CryptoKeyNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**ProtectionType** | Pointer to **string** |  | [optional] 
**RemoteCluster** | Pointer to [**RemoteClusterNestview**](RemoteClusterNestview.md) |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDpSite

`func NewDpSite() *DpSite`

NewDpSite instantiates a new DpSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpSiteWithDefaults

`func NewDpSiteWithDefaults() *DpSite`

NewDpSiteWithDefaults instantiates a new DpSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DpSite) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DpSite) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DpSite) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DpSite) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCluster

`func (o *DpSite) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DpSite) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DpSite) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DpSite) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConfig

`func (o *DpSite) GetConfig() DpSiteConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DpSite) GetConfigOk() (*DpSiteConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DpSite) SetConfig(v DpSiteConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DpSite) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreate

`func (o *DpSite) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DpSite) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DpSite) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DpSite) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCryptoKey

`func (o *DpSite) GetCryptoKey() CryptoKeyNestview`

GetCryptoKey returns the CryptoKey field if non-nil, zero value otherwise.

### GetCryptoKeyOk

`func (o *DpSite) GetCryptoKeyOk() (*CryptoKeyNestview, bool)`

GetCryptoKeyOk returns a tuple with the CryptoKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoKey

`func (o *DpSite) SetCryptoKey(v CryptoKeyNestview)`

SetCryptoKey sets CryptoKey field to given value.

### HasCryptoKey

`func (o *DpSite) HasCryptoKey() bool`

HasCryptoKey returns a boolean if a field has been set.

### GetId

`func (o *DpSite) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpSite) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpSite) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DpSite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DpSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpSite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DpSite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *DpSite) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DpSite) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DpSite) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DpSite) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProtectionType

`func (o *DpSite) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *DpSite) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *DpSite) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *DpSite) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### GetRemoteCluster

`func (o *DpSite) GetRemoteCluster() RemoteClusterNestview`

GetRemoteCluster returns the RemoteCluster field if non-nil, zero value otherwise.

### GetRemoteClusterOk

`func (o *DpSite) GetRemoteClusterOk() (*RemoteClusterNestview, bool)`

GetRemoteClusterOk returns a tuple with the RemoteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCluster

`func (o *DpSite) SetRemoteCluster(v RemoteClusterNestview)`

SetRemoteCluster sets RemoteCluster field to given value.

### HasRemoteCluster

`func (o *DpSite) HasRemoteCluster() bool`

HasRemoteCluster returns a boolean if a field has been set.

### GetService

`func (o *DpSite) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DpSite) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DpSite) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DpSite) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *DpSite) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpSite) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpSite) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DpSite) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DpSite) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DpSite) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DpSite) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DpSite) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


