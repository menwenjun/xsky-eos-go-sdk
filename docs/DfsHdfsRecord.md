# DfsHdfsRecord

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
**Samples** | Pointer to [**[]DfsHdfsStat**](DfsHdfsStat.md) |  | [optional] 

## Methods

### NewDfsHdfsRecord

`func NewDfsHdfsRecord() *DfsHdfsRecord`

NewDfsHdfsRecord instantiates a new DfsHdfsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsHdfsRecordWithDefaults

`func NewDfsHdfsRecordWithDefaults() *DfsHdfsRecord`

NewDfsHdfsRecordWithDefaults instantiates a new DfsHdfsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsHdfsRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsHdfsRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsHdfsRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsHdfsRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAuthToLocals

`func (o *DfsHdfsRecord) GetAuthToLocals() []string`

GetAuthToLocals returns the AuthToLocals field if non-nil, zero value otherwise.

### GetAuthToLocalsOk

`func (o *DfsHdfsRecord) GetAuthToLocalsOk() (*[]string, bool)`

GetAuthToLocalsOk returns a tuple with the AuthToLocals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToLocals

`func (o *DfsHdfsRecord) SetAuthToLocals(v []string)`

SetAuthToLocals sets AuthToLocals field to given value.

### HasAuthToLocals

`func (o *DfsHdfsRecord) HasAuthToLocals() bool`

HasAuthToLocals returns a boolean if a field has been set.

### GetAuthType

`func (o *DfsHdfsRecord) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *DfsHdfsRecord) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *DfsHdfsRecord) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *DfsHdfsRecord) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetBlockSize

`func (o *DfsHdfsRecord) GetBlockSize() int64`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *DfsHdfsRecord) GetBlockSizeOk() (*int64, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *DfsHdfsRecord) SetBlockSize(v int64)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *DfsHdfsRecord) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetChecksumType

`func (o *DfsHdfsRecord) GetChecksumType() string`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *DfsHdfsRecord) GetChecksumTypeOk() (*string, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *DfsHdfsRecord) SetChecksumType(v string)`

SetChecksumType sets ChecksumType field to given value.

### HasChecksumType

`func (o *DfsHdfsRecord) HasChecksumType() bool`

HasChecksumType returns a boolean if a field has been set.

### GetCluster

`func (o *DfsHdfsRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsHdfsRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsHdfsRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsHdfsRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsHdfsRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsHdfsRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsHdfsRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsHdfsRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DfsHdfsRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsHdfsRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsHdfsRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsHdfsRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsHdfsRecord) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsHdfsRecord) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsHdfsRecord) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsHdfsRecord) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetDfsGatewayZone

`func (o *DfsHdfsRecord) GetDfsGatewayZone() DfsGatewayZoneNestview`

GetDfsGatewayZone returns the DfsGatewayZone field if non-nil, zero value otherwise.

### GetDfsGatewayZoneOk

`func (o *DfsHdfsRecord) GetDfsGatewayZoneOk() (*DfsGatewayZoneNestview, bool)`

GetDfsGatewayZoneOk returns a tuple with the DfsGatewayZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayZone

`func (o *DfsHdfsRecord) SetDfsGatewayZone(v DfsGatewayZoneNestview)`

SetDfsGatewayZone sets DfsGatewayZone field to given value.

### HasDfsGatewayZone

`func (o *DfsHdfsRecord) HasDfsGatewayZone() bool`

HasDfsGatewayZone returns a boolean if a field has been set.

### GetDfsHdfsAclNum

`func (o *DfsHdfsRecord) GetDfsHdfsAclNum() int64`

GetDfsHdfsAclNum returns the DfsHdfsAclNum field if non-nil, zero value otherwise.

### GetDfsHdfsAclNumOk

`func (o *DfsHdfsRecord) GetDfsHdfsAclNumOk() (*int64, bool)`

GetDfsHdfsAclNumOk returns a tuple with the DfsHdfsAclNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsAclNum

`func (o *DfsHdfsRecord) SetDfsHdfsAclNum(v int64)`

SetDfsHdfsAclNum sets DfsHdfsAclNum field to given value.

### HasDfsHdfsAclNum

`func (o *DfsHdfsRecord) HasDfsHdfsAclNum() bool`

HasDfsHdfsAclNum returns a boolean if a field has been set.

### GetDfsHdfsIpWhiteLists

`func (o *DfsHdfsRecord) GetDfsHdfsIpWhiteLists() []DfsHdfsIPWhiteList`

GetDfsHdfsIpWhiteLists returns the DfsHdfsIpWhiteLists field if non-nil, zero value otherwise.

### GetDfsHdfsIpWhiteListsOk

`func (o *DfsHdfsRecord) GetDfsHdfsIpWhiteListsOk() (*[]DfsHdfsIPWhiteList, bool)`

GetDfsHdfsIpWhiteListsOk returns a tuple with the DfsHdfsIpWhiteLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsIpWhiteLists

`func (o *DfsHdfsRecord) SetDfsHdfsIpWhiteLists(v []DfsHdfsIPWhiteList)`

SetDfsHdfsIpWhiteLists sets DfsHdfsIpWhiteLists field to given value.

### HasDfsHdfsIpWhiteLists

`func (o *DfsHdfsRecord) HasDfsHdfsIpWhiteLists() bool`

HasDfsHdfsIpWhiteLists returns a boolean if a field has been set.

### GetDfsHdfsIpWhiteListsNum

`func (o *DfsHdfsRecord) GetDfsHdfsIpWhiteListsNum() int64`

GetDfsHdfsIpWhiteListsNum returns the DfsHdfsIpWhiteListsNum field if non-nil, zero value otherwise.

### GetDfsHdfsIpWhiteListsNumOk

`func (o *DfsHdfsRecord) GetDfsHdfsIpWhiteListsNumOk() (*int64, bool)`

GetDfsHdfsIpWhiteListsNumOk returns a tuple with the DfsHdfsIpWhiteListsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsIpWhiteListsNum

`func (o *DfsHdfsRecord) SetDfsHdfsIpWhiteListsNum(v int64)`

SetDfsHdfsIpWhiteListsNum sets DfsHdfsIpWhiteListsNum field to given value.

### HasDfsHdfsIpWhiteListsNum

`func (o *DfsHdfsRecord) HasDfsHdfsIpWhiteListsNum() bool`

HasDfsHdfsIpWhiteListsNum returns a boolean if a field has been set.

### GetDfsHdfsProxyUserNum

`func (o *DfsHdfsRecord) GetDfsHdfsProxyUserNum() int64`

GetDfsHdfsProxyUserNum returns the DfsHdfsProxyUserNum field if non-nil, zero value otherwise.

### GetDfsHdfsProxyUserNumOk

`func (o *DfsHdfsRecord) GetDfsHdfsProxyUserNumOk() (*int64, bool)`

GetDfsHdfsProxyUserNumOk returns a tuple with the DfsHdfsProxyUserNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsProxyUserNum

`func (o *DfsHdfsRecord) SetDfsHdfsProxyUserNum(v int64)`

SetDfsHdfsProxyUserNum sets DfsHdfsProxyUserNum field to given value.

### HasDfsHdfsProxyUserNum

`func (o *DfsHdfsRecord) HasDfsHdfsProxyUserNum() bool`

HasDfsHdfsProxyUserNum returns a boolean if a field has been set.

### GetDfsPath

`func (o *DfsHdfsRecord) GetDfsPath() DfsPathNestview`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DfsHdfsRecord) GetDfsPathOk() (*DfsPathNestview, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DfsHdfsRecord) SetDfsPath(v DfsPathNestview)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DfsHdfsRecord) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetId

`func (o *DfsHdfsRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsHdfsRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsHdfsRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsHdfsRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKerberos

`func (o *DfsHdfsRecord) GetKerberos() FSKerberosNestview`

GetKerberos returns the Kerberos field if non-nil, zero value otherwise.

### GetKerberosOk

`func (o *DfsHdfsRecord) GetKerberosOk() (*FSKerberosNestview, bool)`

GetKerberosOk returns a tuple with the Kerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberos

`func (o *DfsHdfsRecord) SetKerberos(v FSKerberosNestview)`

SetKerberos sets Kerberos field to given value.

### HasKerberos

`func (o *DfsHdfsRecord) HasKerberos() bool`

HasKerberos returns a boolean if a field has been set.

### GetName

`func (o *DfsHdfsRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsHdfsRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsHdfsRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsHdfsRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *DfsHdfsRecord) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DfsHdfsRecord) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DfsHdfsRecord) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DfsHdfsRecord) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrincipalName

`func (o *DfsHdfsRecord) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *DfsHdfsRecord) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *DfsHdfsRecord) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *DfsHdfsRecord) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### GetRanger

`func (o *DfsHdfsRecord) GetRanger() bool`

GetRanger returns the Ranger field if non-nil, zero value otherwise.

### GetRangerOk

`func (o *DfsHdfsRecord) GetRangerOk() (*bool, bool)`

GetRangerOk returns a tuple with the Ranger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanger

`func (o *DfsHdfsRecord) SetRanger(v bool)`

SetRanger sets Ranger field to given value.

### HasRanger

`func (o *DfsHdfsRecord) HasRanger() bool`

HasRanger returns a boolean if a field has been set.

### GetRangerIp

`func (o *DfsHdfsRecord) GetRangerIp() string`

GetRangerIp returns the RangerIp field if non-nil, zero value otherwise.

### GetRangerIpOk

`func (o *DfsHdfsRecord) GetRangerIpOk() (*string, bool)`

GetRangerIpOk returns a tuple with the RangerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerIp

`func (o *DfsHdfsRecord) SetRangerIp(v string)`

SetRangerIp sets RangerIp field to given value.

### HasRangerIp

`func (o *DfsHdfsRecord) HasRangerIp() bool`

HasRangerIp returns a boolean if a field has been set.

### GetRangerServiceName

`func (o *DfsHdfsRecord) GetRangerServiceName() string`

GetRangerServiceName returns the RangerServiceName field if non-nil, zero value otherwise.

### GetRangerServiceNameOk

`func (o *DfsHdfsRecord) GetRangerServiceNameOk() (*string, bool)`

GetRangerServiceNameOk returns a tuple with the RangerServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerServiceName

`func (o *DfsHdfsRecord) SetRangerServiceName(v string)`

SetRangerServiceName sets RangerServiceName field to given value.

### HasRangerServiceName

`func (o *DfsHdfsRecord) HasRangerServiceName() bool`

HasRangerServiceName returns a boolean if a field has been set.

### GetRangerUrl

`func (o *DfsHdfsRecord) GetRangerUrl() string`

GetRangerUrl returns the RangerUrl field if non-nil, zero value otherwise.

### GetRangerUrlOk

`func (o *DfsHdfsRecord) GetRangerUrlOk() (*string, bool)`

GetRangerUrlOk returns a tuple with the RangerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerUrl

`func (o *DfsHdfsRecord) SetRangerUrl(v string)`

SetRangerUrl sets RangerUrl field to given value.

### HasRangerUrl

`func (o *DfsHdfsRecord) HasRangerUrl() bool`

HasRangerUrl returns a boolean if a field has been set.

### GetSecurities

`func (o *DfsHdfsRecord) GetSecurities() []string`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *DfsHdfsRecord) GetSecuritiesOk() (*[]string, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *DfsHdfsRecord) SetSecurities(v []string)`

SetSecurities sets Securities field to given value.

### HasSecurities

`func (o *DfsHdfsRecord) HasSecurities() bool`

HasSecurities returns a boolean if a field has been set.

### GetStatus

`func (o *DfsHdfsRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsHdfsRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsHdfsRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsHdfsRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsHdfsRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsHdfsRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsHdfsRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsHdfsRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVersion

`func (o *DfsHdfsRecord) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DfsHdfsRecord) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DfsHdfsRecord) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DfsHdfsRecord) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSamples

`func (o *DfsHdfsRecord) GetSamples() []DfsHdfsStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DfsHdfsRecord) GetSamplesOk() (*[]DfsHdfsStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DfsHdfsRecord) SetSamples(v []DfsHdfsStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *DfsHdfsRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


