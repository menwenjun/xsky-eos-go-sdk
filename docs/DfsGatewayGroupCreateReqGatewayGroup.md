# DfsGatewayGroupCreateReqGatewayGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdId** | Pointer to **int64** | active directory id | [optional] 
**Cpus** | Pointer to **int64** | cpus of gateway container | [optional] 
**Description** | Pointer to **string** | description of gateway group | [optional] 
**DfsGateways** | [**[]DfsGatewayReq**](DfsGatewayReq.md) | dfs gateways list | 
**DfsVipGateways** | Pointer to **[]string** | VIPGateways contains all the gateways of VIP network | [optional] 
**DfsVips** | **[]string** | VIPs of gateway group | 
**Encoding** | Pointer to **string** | ftp encoding format, default is utf8 | [optional] 
**HdfsSecurities** | Pointer to **[]string** | security type for hdfs (local, ldap) | [optional] 
**LdapId** | Pointer to **int64** | ldap id | [optional] 
**MemoryKbyte** | Pointer to **int64** | memory limit of gateway container, unit KB | [optional] 
**Name** | **string** | name of gateway group | 
**NfsVersions** | Pointer to **[]string** | nfs versions of nfs supported | [optional] 
**Securities** | Pointer to **[]string** | security type for smb/quota (local, ad, ldap) | [optional] 
**Smb1Enabled** | Pointer to **bool** | smb version 1.0 enabled | [optional] 
**SmbBrowseable** | Pointer to **bool** | smb browseable enable | [optional] 
**SmbHomes** | Pointer to **bool** | smb Homes share enable | [optional] 
**SmbPorts** | Pointer to **[]int64** | smb ports | [optional] 
**SmbType** | Pointer to **string** | smb service type (smb, xsmb) | [optional] 
**Types** | **[]string** | types of supported (smb, nfs, ftp) | 
**ZoneName** | Pointer to **string** | name of gateway zone | [optional] 

## Methods

### NewDfsGatewayGroupCreateReqGatewayGroup

`func NewDfsGatewayGroupCreateReqGatewayGroup(dfsGateways []DfsGatewayReq, dfsVips []string, name string, types []string, ) *DfsGatewayGroupCreateReqGatewayGroup`

NewDfsGatewayGroupCreateReqGatewayGroup instantiates a new DfsGatewayGroupCreateReqGatewayGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayGroupCreateReqGatewayGroupWithDefaults

`func NewDfsGatewayGroupCreateReqGatewayGroupWithDefaults() *DfsGatewayGroupCreateReqGatewayGroup`

NewDfsGatewayGroupCreateReqGatewayGroupWithDefaults instantiates a new DfsGatewayGroupCreateReqGatewayGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdId

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetAdId() int64`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetAdIdOk() (*int64, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetAdId(v int64)`

SetAdId sets AdId field to given value.

### HasAdId

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasAdId() bool`

HasAdId returns a boolean if a field has been set.

### GetCpus

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetCpus() int64`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetCpusOk() (*int64, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetCpus(v int64)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetDescription

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGateways

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetDfsGateways() []DfsGatewayReq`

GetDfsGateways returns the DfsGateways field if non-nil, zero value otherwise.

### GetDfsGatewaysOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetDfsGatewaysOk() (*[]DfsGatewayReq, bool)`

GetDfsGatewaysOk returns a tuple with the DfsGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateways

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetDfsGateways(v []DfsGatewayReq)`

SetDfsGateways sets DfsGateways field to given value.


### GetDfsVipGateways

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetDfsVipGateways() []string`

GetDfsVipGateways returns the DfsVipGateways field if non-nil, zero value otherwise.

### GetDfsVipGatewaysOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetDfsVipGatewaysOk() (*[]string, bool)`

GetDfsVipGatewaysOk returns a tuple with the DfsVipGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsVipGateways

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetDfsVipGateways(v []string)`

SetDfsVipGateways sets DfsVipGateways field to given value.

### HasDfsVipGateways

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasDfsVipGateways() bool`

HasDfsVipGateways returns a boolean if a field has been set.

### GetDfsVips

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetDfsVips() []string`

GetDfsVips returns the DfsVips field if non-nil, zero value otherwise.

### GetDfsVipsOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetDfsVipsOk() (*[]string, bool)`

GetDfsVipsOk returns a tuple with the DfsVips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsVips

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetDfsVips(v []string)`

SetDfsVips sets DfsVips field to given value.


### GetEncoding

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetHdfsSecurities

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetHdfsSecurities() []string`

GetHdfsSecurities returns the HdfsSecurities field if non-nil, zero value otherwise.

### GetHdfsSecuritiesOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetHdfsSecuritiesOk() (*[]string, bool)`

GetHdfsSecuritiesOk returns a tuple with the HdfsSecurities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsSecurities

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetHdfsSecurities(v []string)`

SetHdfsSecurities sets HdfsSecurities field to given value.

### HasHdfsSecurities

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasHdfsSecurities() bool`

HasHdfsSecurities returns a boolean if a field has been set.

### GetLdapId

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetLdapId() int64`

GetLdapId returns the LdapId field if non-nil, zero value otherwise.

### GetLdapIdOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetLdapIdOk() (*int64, bool)`

GetLdapIdOk returns a tuple with the LdapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapId

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetLdapId(v int64)`

SetLdapId sets LdapId field to given value.

### HasLdapId

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasLdapId() bool`

HasLdapId returns a boolean if a field has been set.

### GetMemoryKbyte

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetMemoryKbyte() int64`

GetMemoryKbyte returns the MemoryKbyte field if non-nil, zero value otherwise.

### GetMemoryKbyteOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetMemoryKbyteOk() (*int64, bool)`

GetMemoryKbyteOk returns a tuple with the MemoryKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryKbyte

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetMemoryKbyte(v int64)`

SetMemoryKbyte sets MemoryKbyte field to given value.

### HasMemoryKbyte

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasMemoryKbyte() bool`

HasMemoryKbyte returns a boolean if a field has been set.

### GetName

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetName(v string)`

SetName sets Name field to given value.


### GetNfsVersions

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetNfsVersions() []string`

GetNfsVersions returns the NfsVersions field if non-nil, zero value otherwise.

### GetNfsVersionsOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetNfsVersionsOk() (*[]string, bool)`

GetNfsVersionsOk returns a tuple with the NfsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsVersions

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetNfsVersions(v []string)`

SetNfsVersions sets NfsVersions field to given value.

### HasNfsVersions

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasNfsVersions() bool`

HasNfsVersions returns a boolean if a field has been set.

### GetSecurities

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSecurities() []string`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSecuritiesOk() (*[]string, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetSecurities(v []string)`

SetSecurities sets Securities field to given value.

### HasSecurities

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasSecurities() bool`

HasSecurities returns a boolean if a field has been set.

### GetSmb1Enabled

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSmb1Enabled() bool`

GetSmb1Enabled returns the Smb1Enabled field if non-nil, zero value otherwise.

### GetSmb1EnabledOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSmb1EnabledOk() (*bool, bool)`

GetSmb1EnabledOk returns a tuple with the Smb1Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmb1Enabled

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetSmb1Enabled(v bool)`

SetSmb1Enabled sets Smb1Enabled field to given value.

### HasSmb1Enabled

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasSmb1Enabled() bool`

HasSmb1Enabled returns a boolean if a field has been set.

### GetSmbBrowseable

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSmbBrowseable() bool`

GetSmbBrowseable returns the SmbBrowseable field if non-nil, zero value otherwise.

### GetSmbBrowseableOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSmbBrowseableOk() (*bool, bool)`

GetSmbBrowseableOk returns a tuple with the SmbBrowseable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbBrowseable

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetSmbBrowseable(v bool)`

SetSmbBrowseable sets SmbBrowseable field to given value.

### HasSmbBrowseable

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasSmbBrowseable() bool`

HasSmbBrowseable returns a boolean if a field has been set.

### GetSmbHomes

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSmbHomes() bool`

GetSmbHomes returns the SmbHomes field if non-nil, zero value otherwise.

### GetSmbHomesOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSmbHomesOk() (*bool, bool)`

GetSmbHomesOk returns a tuple with the SmbHomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbHomes

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetSmbHomes(v bool)`

SetSmbHomes sets SmbHomes field to given value.

### HasSmbHomes

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasSmbHomes() bool`

HasSmbHomes returns a boolean if a field has been set.

### GetSmbPorts

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSmbPorts() []int64`

GetSmbPorts returns the SmbPorts field if non-nil, zero value otherwise.

### GetSmbPortsOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSmbPortsOk() (*[]int64, bool)`

GetSmbPortsOk returns a tuple with the SmbPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbPorts

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetSmbPorts(v []int64)`

SetSmbPorts sets SmbPorts field to given value.

### HasSmbPorts

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasSmbPorts() bool`

HasSmbPorts returns a boolean if a field has been set.

### GetSmbType

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSmbType() string`

GetSmbType returns the SmbType field if non-nil, zero value otherwise.

### GetSmbTypeOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetSmbTypeOk() (*string, bool)`

GetSmbTypeOk returns a tuple with the SmbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbType

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetSmbType(v string)`

SetSmbType sets SmbType field to given value.

### HasSmbType

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasSmbType() bool`

HasSmbType returns a boolean if a field has been set.

### GetTypes

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetTypes(v []string)`

SetTypes sets Types field to given value.


### GetZoneName

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *DfsGatewayGroupCreateReqGatewayGroup) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *DfsGatewayGroupCreateReqGatewayGroup) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *DfsGatewayGroupCreateReqGatewayGroup) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


