# DfsHdfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AuthToLocals** | Pointer to **[]string** |  | [optional] 
**AuthType** | Pointer to **string** |  | [optional] 
**BlockSize** | Pointer to **int64** |  | [optional] 
**ChecksumType** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DfsGatewayGroup** | Pointer to [**DfsGatewayGroupNestview**](DfsGatewayGroupNestview.md) |  | [optional] 
**DfsGatewayZone** | Pointer to [**DfsGatewayZoneNestview**](DfsGatewayZoneNestview.md) |  | [optional] 
**DfsHdfsAclNum** | Pointer to **int64** |  | [optional] 
**DfsHdfsIpWhiteLists** | Pointer to [**[]DfsHdfsIPWhiteList**](DfsHdfsIPWhiteList.md) |  | [optional] 
**DfsHdfsIpWhiteListsNum** | Pointer to **int64** |  | [optional] 
**DfsHdfsProxyUserNum** | Pointer to **int64** |  | [optional] 
**DfsPath** | Pointer to [**DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Kerberos** | Pointer to [**FSKerberosNestview**](FSKerberosNestview.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**PrincipalName** | Pointer to **string** |  | [optional] 
**Ranger** | Pointer to **bool** |  | [optional] 
**RangerIp** | Pointer to **string** |  | [optional] 
**RangerServiceName** | Pointer to **string** |  | [optional] 
**RangerUrl** | Pointer to **string** |  | [optional] 
**Securities** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsHdfs

`func NewDfsHdfs() *DfsHdfs`

NewDfsHdfs instantiates a new DfsHdfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsHdfsWithDefaults

`func NewDfsHdfsWithDefaults() *DfsHdfs`

NewDfsHdfsWithDefaults instantiates a new DfsHdfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsHdfs) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsHdfs) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsHdfs) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsHdfs) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAuthToLocals

`func (o *DfsHdfs) GetAuthToLocals() []string`

GetAuthToLocals returns the AuthToLocals field if non-nil, zero value otherwise.

### GetAuthToLocalsOk

`func (o *DfsHdfs) GetAuthToLocalsOk() (*[]string, bool)`

GetAuthToLocalsOk returns a tuple with the AuthToLocals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToLocals

`func (o *DfsHdfs) SetAuthToLocals(v []string)`

SetAuthToLocals sets AuthToLocals field to given value.

### HasAuthToLocals

`func (o *DfsHdfs) HasAuthToLocals() bool`

HasAuthToLocals returns a boolean if a field has been set.

### GetAuthType

`func (o *DfsHdfs) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *DfsHdfs) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *DfsHdfs) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *DfsHdfs) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetBlockSize

`func (o *DfsHdfs) GetBlockSize() int64`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *DfsHdfs) GetBlockSizeOk() (*int64, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *DfsHdfs) SetBlockSize(v int64)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *DfsHdfs) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetChecksumType

`func (o *DfsHdfs) GetChecksumType() string`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *DfsHdfs) GetChecksumTypeOk() (*string, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *DfsHdfs) SetChecksumType(v string)`

SetChecksumType sets ChecksumType field to given value.

### HasChecksumType

`func (o *DfsHdfs) HasChecksumType() bool`

HasChecksumType returns a boolean if a field has been set.

### GetCluster

`func (o *DfsHdfs) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsHdfs) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsHdfs) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsHdfs) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsHdfs) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsHdfs) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsHdfs) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsHdfs) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DfsHdfs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsHdfs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsHdfs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsHdfs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsHdfs) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsHdfs) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsHdfs) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsHdfs) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetDfsGatewayZone

`func (o *DfsHdfs) GetDfsGatewayZone() DfsGatewayZoneNestview`

GetDfsGatewayZone returns the DfsGatewayZone field if non-nil, zero value otherwise.

### GetDfsGatewayZoneOk

`func (o *DfsHdfs) GetDfsGatewayZoneOk() (*DfsGatewayZoneNestview, bool)`

GetDfsGatewayZoneOk returns a tuple with the DfsGatewayZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayZone

`func (o *DfsHdfs) SetDfsGatewayZone(v DfsGatewayZoneNestview)`

SetDfsGatewayZone sets DfsGatewayZone field to given value.

### HasDfsGatewayZone

`func (o *DfsHdfs) HasDfsGatewayZone() bool`

HasDfsGatewayZone returns a boolean if a field has been set.

### GetDfsHdfsAclNum

`func (o *DfsHdfs) GetDfsHdfsAclNum() int64`

GetDfsHdfsAclNum returns the DfsHdfsAclNum field if non-nil, zero value otherwise.

### GetDfsHdfsAclNumOk

`func (o *DfsHdfs) GetDfsHdfsAclNumOk() (*int64, bool)`

GetDfsHdfsAclNumOk returns a tuple with the DfsHdfsAclNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsAclNum

`func (o *DfsHdfs) SetDfsHdfsAclNum(v int64)`

SetDfsHdfsAclNum sets DfsHdfsAclNum field to given value.

### HasDfsHdfsAclNum

`func (o *DfsHdfs) HasDfsHdfsAclNum() bool`

HasDfsHdfsAclNum returns a boolean if a field has been set.

### GetDfsHdfsIpWhiteLists

`func (o *DfsHdfs) GetDfsHdfsIpWhiteLists() []DfsHdfsIPWhiteList`

GetDfsHdfsIpWhiteLists returns the DfsHdfsIpWhiteLists field if non-nil, zero value otherwise.

### GetDfsHdfsIpWhiteListsOk

`func (o *DfsHdfs) GetDfsHdfsIpWhiteListsOk() (*[]DfsHdfsIPWhiteList, bool)`

GetDfsHdfsIpWhiteListsOk returns a tuple with the DfsHdfsIpWhiteLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsIpWhiteLists

`func (o *DfsHdfs) SetDfsHdfsIpWhiteLists(v []DfsHdfsIPWhiteList)`

SetDfsHdfsIpWhiteLists sets DfsHdfsIpWhiteLists field to given value.

### HasDfsHdfsIpWhiteLists

`func (o *DfsHdfs) HasDfsHdfsIpWhiteLists() bool`

HasDfsHdfsIpWhiteLists returns a boolean if a field has been set.

### GetDfsHdfsIpWhiteListsNum

`func (o *DfsHdfs) GetDfsHdfsIpWhiteListsNum() int64`

GetDfsHdfsIpWhiteListsNum returns the DfsHdfsIpWhiteListsNum field if non-nil, zero value otherwise.

### GetDfsHdfsIpWhiteListsNumOk

`func (o *DfsHdfs) GetDfsHdfsIpWhiteListsNumOk() (*int64, bool)`

GetDfsHdfsIpWhiteListsNumOk returns a tuple with the DfsHdfsIpWhiteListsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsIpWhiteListsNum

`func (o *DfsHdfs) SetDfsHdfsIpWhiteListsNum(v int64)`

SetDfsHdfsIpWhiteListsNum sets DfsHdfsIpWhiteListsNum field to given value.

### HasDfsHdfsIpWhiteListsNum

`func (o *DfsHdfs) HasDfsHdfsIpWhiteListsNum() bool`

HasDfsHdfsIpWhiteListsNum returns a boolean if a field has been set.

### GetDfsHdfsProxyUserNum

`func (o *DfsHdfs) GetDfsHdfsProxyUserNum() int64`

GetDfsHdfsProxyUserNum returns the DfsHdfsProxyUserNum field if non-nil, zero value otherwise.

### GetDfsHdfsProxyUserNumOk

`func (o *DfsHdfs) GetDfsHdfsProxyUserNumOk() (*int64, bool)`

GetDfsHdfsProxyUserNumOk returns a tuple with the DfsHdfsProxyUserNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsProxyUserNum

`func (o *DfsHdfs) SetDfsHdfsProxyUserNum(v int64)`

SetDfsHdfsProxyUserNum sets DfsHdfsProxyUserNum field to given value.

### HasDfsHdfsProxyUserNum

`func (o *DfsHdfs) HasDfsHdfsProxyUserNum() bool`

HasDfsHdfsProxyUserNum returns a boolean if a field has been set.

### GetDfsPath

`func (o *DfsHdfs) GetDfsPath() DfsPathNestview`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DfsHdfs) GetDfsPathOk() (*DfsPathNestview, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DfsHdfs) SetDfsPath(v DfsPathNestview)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DfsHdfs) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetId

`func (o *DfsHdfs) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsHdfs) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsHdfs) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsHdfs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKerberos

`func (o *DfsHdfs) GetKerberos() FSKerberosNestview`

GetKerberos returns the Kerberos field if non-nil, zero value otherwise.

### GetKerberosOk

`func (o *DfsHdfs) GetKerberosOk() (*FSKerberosNestview, bool)`

GetKerberosOk returns a tuple with the Kerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberos

`func (o *DfsHdfs) SetKerberos(v FSKerberosNestview)`

SetKerberos sets Kerberos field to given value.

### HasKerberos

`func (o *DfsHdfs) HasKerberos() bool`

HasKerberos returns a boolean if a field has been set.

### GetName

`func (o *DfsHdfs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsHdfs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsHdfs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsHdfs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *DfsHdfs) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DfsHdfs) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DfsHdfs) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DfsHdfs) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrincipalName

`func (o *DfsHdfs) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *DfsHdfs) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *DfsHdfs) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *DfsHdfs) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### GetRanger

`func (o *DfsHdfs) GetRanger() bool`

GetRanger returns the Ranger field if non-nil, zero value otherwise.

### GetRangerOk

`func (o *DfsHdfs) GetRangerOk() (*bool, bool)`

GetRangerOk returns a tuple with the Ranger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanger

`func (o *DfsHdfs) SetRanger(v bool)`

SetRanger sets Ranger field to given value.

### HasRanger

`func (o *DfsHdfs) HasRanger() bool`

HasRanger returns a boolean if a field has been set.

### GetRangerIp

`func (o *DfsHdfs) GetRangerIp() string`

GetRangerIp returns the RangerIp field if non-nil, zero value otherwise.

### GetRangerIpOk

`func (o *DfsHdfs) GetRangerIpOk() (*string, bool)`

GetRangerIpOk returns a tuple with the RangerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerIp

`func (o *DfsHdfs) SetRangerIp(v string)`

SetRangerIp sets RangerIp field to given value.

### HasRangerIp

`func (o *DfsHdfs) HasRangerIp() bool`

HasRangerIp returns a boolean if a field has been set.

### GetRangerServiceName

`func (o *DfsHdfs) GetRangerServiceName() string`

GetRangerServiceName returns the RangerServiceName field if non-nil, zero value otherwise.

### GetRangerServiceNameOk

`func (o *DfsHdfs) GetRangerServiceNameOk() (*string, bool)`

GetRangerServiceNameOk returns a tuple with the RangerServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerServiceName

`func (o *DfsHdfs) SetRangerServiceName(v string)`

SetRangerServiceName sets RangerServiceName field to given value.

### HasRangerServiceName

`func (o *DfsHdfs) HasRangerServiceName() bool`

HasRangerServiceName returns a boolean if a field has been set.

### GetRangerUrl

`func (o *DfsHdfs) GetRangerUrl() string`

GetRangerUrl returns the RangerUrl field if non-nil, zero value otherwise.

### GetRangerUrlOk

`func (o *DfsHdfs) GetRangerUrlOk() (*string, bool)`

GetRangerUrlOk returns a tuple with the RangerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerUrl

`func (o *DfsHdfs) SetRangerUrl(v string)`

SetRangerUrl sets RangerUrl field to given value.

### HasRangerUrl

`func (o *DfsHdfs) HasRangerUrl() bool`

HasRangerUrl returns a boolean if a field has been set.

### GetSecurities

`func (o *DfsHdfs) GetSecurities() []string`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *DfsHdfs) GetSecuritiesOk() (*[]string, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *DfsHdfs) SetSecurities(v []string)`

SetSecurities sets Securities field to given value.

### HasSecurities

`func (o *DfsHdfs) HasSecurities() bool`

HasSecurities returns a boolean if a field has been set.

### GetStatus

`func (o *DfsHdfs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsHdfs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsHdfs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsHdfs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsHdfs) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsHdfs) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsHdfs) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsHdfs) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVersion

`func (o *DfsHdfs) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DfsHdfs) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DfsHdfs) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DfsHdfs) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


