# DfsHdfsCreateReqHdfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToLocals** | Pointer to **[]string** | auth to local of hdfs | [optional] 
**AuthType** | **string** | auth type of hdfs | 
**BlockSize** | **int64** | block size | 
**ChecksumType** | **string** | checksum type | 
**Description** | Pointer to **string** | description of hdfs | [optional] 
**DfsGatewayGroupId** | Pointer to **int64** | id of dfs gateway group | [optional] 
**DfsGatewayZoneId** | **int64** | id of dfs gateway zone | 
**DfsHdfsAcls** | Pointer to [**[]DfsHdfsACLReq**](DfsHdfsACLReq.md) | dfs hdfs acl list | [optional] 
**DfsHdfsIpWhiteLists** | Pointer to [**[]DfsHdfsIPWhiteList**](DfsHdfsIPWhiteList.md) | dfs hdfs access ip list | [optional] 
**DfsHdfsProxyUsers** | Pointer to [**[]DfsHdfsProxyUserReq**](DfsHdfsProxyUserReq.md) | dfs hdfs proxy user list | [optional] 
**KerberosId** | Pointer to **int64** | fs kerberos id | [optional] 
**Name** | **string** | dfs hdfs name | 
**Path** | **string** | path | 
**Port** | Pointer to **int64** | port of dfs hdfs | [optional] 
**Ranger** | Pointer to **bool** | enabled of hdfs ranger | [optional] 
**RangerIp** | Pointer to **string** | ranger ip of hdfs | [optional] 
**RangerServiceName** | Pointer to **string** | ranger service name of hdfs | [optional] 
**RangerUrl** | Pointer to **string** | ranger url of hdfs | [optional] 
**RootfsId** | **int64** | id of dfs rootfs | 

## Methods

### NewDfsHdfsCreateReqHdfs

`func NewDfsHdfsCreateReqHdfs(authType string, blockSize int64, checksumType string, dfsGatewayZoneId int64, name string, path string, rootfsId int64, ) *DfsHdfsCreateReqHdfs`

NewDfsHdfsCreateReqHdfs instantiates a new DfsHdfsCreateReqHdfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsHdfsCreateReqHdfsWithDefaults

`func NewDfsHdfsCreateReqHdfsWithDefaults() *DfsHdfsCreateReqHdfs`

NewDfsHdfsCreateReqHdfsWithDefaults instantiates a new DfsHdfsCreateReqHdfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToLocals

`func (o *DfsHdfsCreateReqHdfs) GetAuthToLocals() []string`

GetAuthToLocals returns the AuthToLocals field if non-nil, zero value otherwise.

### GetAuthToLocalsOk

`func (o *DfsHdfsCreateReqHdfs) GetAuthToLocalsOk() (*[]string, bool)`

GetAuthToLocalsOk returns a tuple with the AuthToLocals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToLocals

`func (o *DfsHdfsCreateReqHdfs) SetAuthToLocals(v []string)`

SetAuthToLocals sets AuthToLocals field to given value.

### HasAuthToLocals

`func (o *DfsHdfsCreateReqHdfs) HasAuthToLocals() bool`

HasAuthToLocals returns a boolean if a field has been set.

### GetAuthType

`func (o *DfsHdfsCreateReqHdfs) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *DfsHdfsCreateReqHdfs) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *DfsHdfsCreateReqHdfs) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetBlockSize

`func (o *DfsHdfsCreateReqHdfs) GetBlockSize() int64`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *DfsHdfsCreateReqHdfs) GetBlockSizeOk() (*int64, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *DfsHdfsCreateReqHdfs) SetBlockSize(v int64)`

SetBlockSize sets BlockSize field to given value.


### GetChecksumType

`func (o *DfsHdfsCreateReqHdfs) GetChecksumType() string`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *DfsHdfsCreateReqHdfs) GetChecksumTypeOk() (*string, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *DfsHdfsCreateReqHdfs) SetChecksumType(v string)`

SetChecksumType sets ChecksumType field to given value.


### GetDescription

`func (o *DfsHdfsCreateReqHdfs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsHdfsCreateReqHdfs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsHdfsCreateReqHdfs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsHdfsCreateReqHdfs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroupId

`func (o *DfsHdfsCreateReqHdfs) GetDfsGatewayGroupId() int64`

GetDfsGatewayGroupId returns the DfsGatewayGroupId field if non-nil, zero value otherwise.

### GetDfsGatewayGroupIdOk

`func (o *DfsHdfsCreateReqHdfs) GetDfsGatewayGroupIdOk() (*int64, bool)`

GetDfsGatewayGroupIdOk returns a tuple with the DfsGatewayGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroupId

`func (o *DfsHdfsCreateReqHdfs) SetDfsGatewayGroupId(v int64)`

SetDfsGatewayGroupId sets DfsGatewayGroupId field to given value.

### HasDfsGatewayGroupId

`func (o *DfsHdfsCreateReqHdfs) HasDfsGatewayGroupId() bool`

HasDfsGatewayGroupId returns a boolean if a field has been set.

### GetDfsGatewayZoneId

`func (o *DfsHdfsCreateReqHdfs) GetDfsGatewayZoneId() int64`

GetDfsGatewayZoneId returns the DfsGatewayZoneId field if non-nil, zero value otherwise.

### GetDfsGatewayZoneIdOk

`func (o *DfsHdfsCreateReqHdfs) GetDfsGatewayZoneIdOk() (*int64, bool)`

GetDfsGatewayZoneIdOk returns a tuple with the DfsGatewayZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayZoneId

`func (o *DfsHdfsCreateReqHdfs) SetDfsGatewayZoneId(v int64)`

SetDfsGatewayZoneId sets DfsGatewayZoneId field to given value.


### GetDfsHdfsAcls

`func (o *DfsHdfsCreateReqHdfs) GetDfsHdfsAcls() []DfsHdfsACLReq`

GetDfsHdfsAcls returns the DfsHdfsAcls field if non-nil, zero value otherwise.

### GetDfsHdfsAclsOk

`func (o *DfsHdfsCreateReqHdfs) GetDfsHdfsAclsOk() (*[]DfsHdfsACLReq, bool)`

GetDfsHdfsAclsOk returns a tuple with the DfsHdfsAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsAcls

`func (o *DfsHdfsCreateReqHdfs) SetDfsHdfsAcls(v []DfsHdfsACLReq)`

SetDfsHdfsAcls sets DfsHdfsAcls field to given value.

### HasDfsHdfsAcls

`func (o *DfsHdfsCreateReqHdfs) HasDfsHdfsAcls() bool`

HasDfsHdfsAcls returns a boolean if a field has been set.

### GetDfsHdfsIpWhiteLists

`func (o *DfsHdfsCreateReqHdfs) GetDfsHdfsIpWhiteLists() []DfsHdfsIPWhiteList`

GetDfsHdfsIpWhiteLists returns the DfsHdfsIpWhiteLists field if non-nil, zero value otherwise.

### GetDfsHdfsIpWhiteListsOk

`func (o *DfsHdfsCreateReqHdfs) GetDfsHdfsIpWhiteListsOk() (*[]DfsHdfsIPWhiteList, bool)`

GetDfsHdfsIpWhiteListsOk returns a tuple with the DfsHdfsIpWhiteLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsIpWhiteLists

`func (o *DfsHdfsCreateReqHdfs) SetDfsHdfsIpWhiteLists(v []DfsHdfsIPWhiteList)`

SetDfsHdfsIpWhiteLists sets DfsHdfsIpWhiteLists field to given value.

### HasDfsHdfsIpWhiteLists

`func (o *DfsHdfsCreateReqHdfs) HasDfsHdfsIpWhiteLists() bool`

HasDfsHdfsIpWhiteLists returns a boolean if a field has been set.

### GetDfsHdfsProxyUsers

`func (o *DfsHdfsCreateReqHdfs) GetDfsHdfsProxyUsers() []DfsHdfsProxyUserReq`

GetDfsHdfsProxyUsers returns the DfsHdfsProxyUsers field if non-nil, zero value otherwise.

### GetDfsHdfsProxyUsersOk

`func (o *DfsHdfsCreateReqHdfs) GetDfsHdfsProxyUsersOk() (*[]DfsHdfsProxyUserReq, bool)`

GetDfsHdfsProxyUsersOk returns a tuple with the DfsHdfsProxyUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfsProxyUsers

`func (o *DfsHdfsCreateReqHdfs) SetDfsHdfsProxyUsers(v []DfsHdfsProxyUserReq)`

SetDfsHdfsProxyUsers sets DfsHdfsProxyUsers field to given value.

### HasDfsHdfsProxyUsers

`func (o *DfsHdfsCreateReqHdfs) HasDfsHdfsProxyUsers() bool`

HasDfsHdfsProxyUsers returns a boolean if a field has been set.

### GetKerberosId

`func (o *DfsHdfsCreateReqHdfs) GetKerberosId() int64`

GetKerberosId returns the KerberosId field if non-nil, zero value otherwise.

### GetKerberosIdOk

`func (o *DfsHdfsCreateReqHdfs) GetKerberosIdOk() (*int64, bool)`

GetKerberosIdOk returns a tuple with the KerberosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosId

`func (o *DfsHdfsCreateReqHdfs) SetKerberosId(v int64)`

SetKerberosId sets KerberosId field to given value.

### HasKerberosId

`func (o *DfsHdfsCreateReqHdfs) HasKerberosId() bool`

HasKerberosId returns a boolean if a field has been set.

### GetName

`func (o *DfsHdfsCreateReqHdfs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsHdfsCreateReqHdfs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsHdfsCreateReqHdfs) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *DfsHdfsCreateReqHdfs) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsHdfsCreateReqHdfs) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsHdfsCreateReqHdfs) SetPath(v string)`

SetPath sets Path field to given value.


### GetPort

`func (o *DfsHdfsCreateReqHdfs) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DfsHdfsCreateReqHdfs) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DfsHdfsCreateReqHdfs) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DfsHdfsCreateReqHdfs) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRanger

`func (o *DfsHdfsCreateReqHdfs) GetRanger() bool`

GetRanger returns the Ranger field if non-nil, zero value otherwise.

### GetRangerOk

`func (o *DfsHdfsCreateReqHdfs) GetRangerOk() (*bool, bool)`

GetRangerOk returns a tuple with the Ranger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanger

`func (o *DfsHdfsCreateReqHdfs) SetRanger(v bool)`

SetRanger sets Ranger field to given value.

### HasRanger

`func (o *DfsHdfsCreateReqHdfs) HasRanger() bool`

HasRanger returns a boolean if a field has been set.

### GetRangerIp

`func (o *DfsHdfsCreateReqHdfs) GetRangerIp() string`

GetRangerIp returns the RangerIp field if non-nil, zero value otherwise.

### GetRangerIpOk

`func (o *DfsHdfsCreateReqHdfs) GetRangerIpOk() (*string, bool)`

GetRangerIpOk returns a tuple with the RangerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerIp

`func (o *DfsHdfsCreateReqHdfs) SetRangerIp(v string)`

SetRangerIp sets RangerIp field to given value.

### HasRangerIp

`func (o *DfsHdfsCreateReqHdfs) HasRangerIp() bool`

HasRangerIp returns a boolean if a field has been set.

### GetRangerServiceName

`func (o *DfsHdfsCreateReqHdfs) GetRangerServiceName() string`

GetRangerServiceName returns the RangerServiceName field if non-nil, zero value otherwise.

### GetRangerServiceNameOk

`func (o *DfsHdfsCreateReqHdfs) GetRangerServiceNameOk() (*string, bool)`

GetRangerServiceNameOk returns a tuple with the RangerServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerServiceName

`func (o *DfsHdfsCreateReqHdfs) SetRangerServiceName(v string)`

SetRangerServiceName sets RangerServiceName field to given value.

### HasRangerServiceName

`func (o *DfsHdfsCreateReqHdfs) HasRangerServiceName() bool`

HasRangerServiceName returns a boolean if a field has been set.

### GetRangerUrl

`func (o *DfsHdfsCreateReqHdfs) GetRangerUrl() string`

GetRangerUrl returns the RangerUrl field if non-nil, zero value otherwise.

### GetRangerUrlOk

`func (o *DfsHdfsCreateReqHdfs) GetRangerUrlOk() (*string, bool)`

GetRangerUrlOk returns a tuple with the RangerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerUrl

`func (o *DfsHdfsCreateReqHdfs) SetRangerUrl(v string)`

SetRangerUrl sets RangerUrl field to given value.

### HasRangerUrl

`func (o *DfsHdfsCreateReqHdfs) HasRangerUrl() bool`

HasRangerUrl returns a boolean if a field has been set.

### GetRootfsId

`func (o *DfsHdfsCreateReqHdfs) GetRootfsId() int64`

GetRootfsId returns the RootfsId field if non-nil, zero value otherwise.

### GetRootfsIdOk

`func (o *DfsHdfsCreateReqHdfs) GetRootfsIdOk() (*int64, bool)`

GetRootfsIdOk returns a tuple with the RootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfsId

`func (o *DfsHdfsCreateReqHdfs) SetRootfsId(v int64)`

SetRootfsId sets RootfsId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


