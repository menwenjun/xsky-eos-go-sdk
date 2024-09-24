# RBDClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminIp** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**CpuArch** | Pointer to **string** |  | [optional] 
**CpuCoreNum** | Pointer to **int64** |  | [optional] 
**CpuModel** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MemoryKbyte** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**PublicNetwork** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**VolumeNum** | Pointer to **int64** |  | [optional] 
**XdcStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewRBDClient

`func NewRBDClient() *RBDClient`

NewRBDClient instantiates a new RBDClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRBDClientWithDefaults

`func NewRBDClientWithDefaults() *RBDClient`

NewRBDClientWithDefaults instantiates a new RBDClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminIp

`func (o *RBDClient) GetAdminIp() string`

GetAdminIp returns the AdminIp field if non-nil, zero value otherwise.

### GetAdminIpOk

`func (o *RBDClient) GetAdminIpOk() (*string, bool)`

GetAdminIpOk returns a tuple with the AdminIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIp

`func (o *RBDClient) SetAdminIp(v string)`

SetAdminIp sets AdminIp field to given value.

### HasAdminIp

`func (o *RBDClient) HasAdminIp() bool`

HasAdminIp returns a boolean if a field has been set.

### GetCluster

`func (o *RBDClient) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *RBDClient) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *RBDClient) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *RBDClient) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCpuArch

`func (o *RBDClient) GetCpuArch() string`

GetCpuArch returns the CpuArch field if non-nil, zero value otherwise.

### GetCpuArchOk

`func (o *RBDClient) GetCpuArchOk() (*string, bool)`

GetCpuArchOk returns a tuple with the CpuArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArch

`func (o *RBDClient) SetCpuArch(v string)`

SetCpuArch sets CpuArch field to given value.

### HasCpuArch

`func (o *RBDClient) HasCpuArch() bool`

HasCpuArch returns a boolean if a field has been set.

### GetCpuCoreNum

`func (o *RBDClient) GetCpuCoreNum() int64`

GetCpuCoreNum returns the CpuCoreNum field if non-nil, zero value otherwise.

### GetCpuCoreNumOk

`func (o *RBDClient) GetCpuCoreNumOk() (*int64, bool)`

GetCpuCoreNumOk returns a tuple with the CpuCoreNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCoreNum

`func (o *RBDClient) SetCpuCoreNum(v int64)`

SetCpuCoreNum sets CpuCoreNum field to given value.

### HasCpuCoreNum

`func (o *RBDClient) HasCpuCoreNum() bool`

HasCpuCoreNum returns a boolean if a field has been set.

### GetCpuModel

`func (o *RBDClient) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *RBDClient) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *RBDClient) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *RBDClient) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetCreate

`func (o *RBDClient) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *RBDClient) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *RBDClient) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *RBDClient) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *RBDClient) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RBDClient) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RBDClient) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RBDClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemoryKbyte

`func (o *RBDClient) GetMemoryKbyte() int64`

GetMemoryKbyte returns the MemoryKbyte field if non-nil, zero value otherwise.

### GetMemoryKbyteOk

`func (o *RBDClient) GetMemoryKbyteOk() (*int64, bool)`

GetMemoryKbyteOk returns a tuple with the MemoryKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryKbyte

`func (o *RBDClient) SetMemoryKbyte(v int64)`

SetMemoryKbyte sets MemoryKbyte field to given value.

### HasMemoryKbyte

`func (o *RBDClient) HasMemoryKbyte() bool`

HasMemoryKbyte returns a boolean if a field has been set.

### GetName

`func (o *RBDClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RBDClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RBDClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RBDClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *RBDClient) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *RBDClient) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *RBDClient) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *RBDClient) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPort

`func (o *RBDClient) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RBDClient) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RBDClient) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RBDClient) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPublicIp

`func (o *RBDClient) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *RBDClient) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *RBDClient) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *RBDClient) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPublicNetwork

`func (o *RBDClient) GetPublicNetwork() string`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *RBDClient) GetPublicNetworkOk() (*string, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *RBDClient) SetPublicNetwork(v string)`

SetPublicNetwork sets PublicNetwork field to given value.

### HasPublicNetwork

`func (o *RBDClient) HasPublicNetwork() bool`

HasPublicNetwork returns a boolean if a field has been set.

### GetStatus

`func (o *RBDClient) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RBDClient) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RBDClient) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RBDClient) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *RBDClient) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RBDClient) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RBDClient) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RBDClient) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *RBDClient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RBDClient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RBDClient) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RBDClient) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *RBDClient) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *RBDClient) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *RBDClient) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *RBDClient) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVendor

`func (o *RBDClient) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *RBDClient) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *RBDClient) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *RBDClient) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *RBDClient) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RBDClient) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RBDClient) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RBDClient) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVolumeNum

`func (o *RBDClient) GetVolumeNum() int64`

GetVolumeNum returns the VolumeNum field if non-nil, zero value otherwise.

### GetVolumeNumOk

`func (o *RBDClient) GetVolumeNumOk() (*int64, bool)`

GetVolumeNumOk returns a tuple with the VolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeNum

`func (o *RBDClient) SetVolumeNum(v int64)`

SetVolumeNum sets VolumeNum field to given value.

### HasVolumeNum

`func (o *RBDClient) HasVolumeNum() bool`

HasVolumeNum returns a boolean if a field has been set.

### GetXdcStatus

`func (o *RBDClient) GetXdcStatus() string`

GetXdcStatus returns the XdcStatus field if non-nil, zero value otherwise.

### GetXdcStatusOk

`func (o *RBDClient) GetXdcStatusOk() (*string, bool)`

GetXdcStatusOk returns a tuple with the XdcStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXdcStatus

`func (o *RBDClient) SetXdcStatus(v string)`

SetXdcStatus sets XdcStatus field to given value.

### HasXdcStatus

`func (o *RBDClient) HasXdcStatus() bool`

HasXdcStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


