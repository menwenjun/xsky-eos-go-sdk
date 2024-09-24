# DfsHdfsUpdateReqHdfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToLocals** | Pointer to **[]string** | auth to local of hdfs | [optional] 
**AuthType** | Pointer to **string** | auth type of hdfs | [optional] 
**BlockSize** | Pointer to **int64** | block size of hdfs | [optional] 
**ChecksumType** | Pointer to **string** | checksum type of hdfs | [optional] 
**Description** | Pointer to **string** | description of hdfs | [optional] 
**KerberosId** | Pointer to **int64** | fs kerberos id | [optional] 
**Path** | Pointer to **string** | path of hdfs | [optional] 
**Port** | Pointer to **int64** | port of dfs hdfs | [optional] 
**Ranger** | Pointer to **bool** | enabled of hdfs ranger | [optional] 
**RangerIp** | Pointer to **string** | ranger ip of hdfs | [optional] 
**RangerServiceName** | Pointer to **string** | ranger service name of hdfs | [optional] 
**RangerUrl** | Pointer to **string** | ranger url of hdfs | [optional] 

## Methods

### NewDfsHdfsUpdateReqHdfs

`func NewDfsHdfsUpdateReqHdfs() *DfsHdfsUpdateReqHdfs`

NewDfsHdfsUpdateReqHdfs instantiates a new DfsHdfsUpdateReqHdfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsHdfsUpdateReqHdfsWithDefaults

`func NewDfsHdfsUpdateReqHdfsWithDefaults() *DfsHdfsUpdateReqHdfs`

NewDfsHdfsUpdateReqHdfsWithDefaults instantiates a new DfsHdfsUpdateReqHdfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToLocals

`func (o *DfsHdfsUpdateReqHdfs) GetAuthToLocals() []string`

GetAuthToLocals returns the AuthToLocals field if non-nil, zero value otherwise.

### GetAuthToLocalsOk

`func (o *DfsHdfsUpdateReqHdfs) GetAuthToLocalsOk() (*[]string, bool)`

GetAuthToLocalsOk returns a tuple with the AuthToLocals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToLocals

`func (o *DfsHdfsUpdateReqHdfs) SetAuthToLocals(v []string)`

SetAuthToLocals sets AuthToLocals field to given value.

### HasAuthToLocals

`func (o *DfsHdfsUpdateReqHdfs) HasAuthToLocals() bool`

HasAuthToLocals returns a boolean if a field has been set.

### GetAuthType

`func (o *DfsHdfsUpdateReqHdfs) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *DfsHdfsUpdateReqHdfs) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *DfsHdfsUpdateReqHdfs) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *DfsHdfsUpdateReqHdfs) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetBlockSize

`func (o *DfsHdfsUpdateReqHdfs) GetBlockSize() int64`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *DfsHdfsUpdateReqHdfs) GetBlockSizeOk() (*int64, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *DfsHdfsUpdateReqHdfs) SetBlockSize(v int64)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *DfsHdfsUpdateReqHdfs) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetChecksumType

`func (o *DfsHdfsUpdateReqHdfs) GetChecksumType() string`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *DfsHdfsUpdateReqHdfs) GetChecksumTypeOk() (*string, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *DfsHdfsUpdateReqHdfs) SetChecksumType(v string)`

SetChecksumType sets ChecksumType field to given value.

### HasChecksumType

`func (o *DfsHdfsUpdateReqHdfs) HasChecksumType() bool`

HasChecksumType returns a boolean if a field has been set.

### GetDescription

`func (o *DfsHdfsUpdateReqHdfs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsHdfsUpdateReqHdfs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsHdfsUpdateReqHdfs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsHdfsUpdateReqHdfs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKerberosId

`func (o *DfsHdfsUpdateReqHdfs) GetKerberosId() int64`

GetKerberosId returns the KerberosId field if non-nil, zero value otherwise.

### GetKerberosIdOk

`func (o *DfsHdfsUpdateReqHdfs) GetKerberosIdOk() (*int64, bool)`

GetKerberosIdOk returns a tuple with the KerberosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosId

`func (o *DfsHdfsUpdateReqHdfs) SetKerberosId(v int64)`

SetKerberosId sets KerberosId field to given value.

### HasKerberosId

`func (o *DfsHdfsUpdateReqHdfs) HasKerberosId() bool`

HasKerberosId returns a boolean if a field has been set.

### GetPath

`func (o *DfsHdfsUpdateReqHdfs) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsHdfsUpdateReqHdfs) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsHdfsUpdateReqHdfs) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsHdfsUpdateReqHdfs) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *DfsHdfsUpdateReqHdfs) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DfsHdfsUpdateReqHdfs) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DfsHdfsUpdateReqHdfs) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DfsHdfsUpdateReqHdfs) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRanger

`func (o *DfsHdfsUpdateReqHdfs) GetRanger() bool`

GetRanger returns the Ranger field if non-nil, zero value otherwise.

### GetRangerOk

`func (o *DfsHdfsUpdateReqHdfs) GetRangerOk() (*bool, bool)`

GetRangerOk returns a tuple with the Ranger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanger

`func (o *DfsHdfsUpdateReqHdfs) SetRanger(v bool)`

SetRanger sets Ranger field to given value.

### HasRanger

`func (o *DfsHdfsUpdateReqHdfs) HasRanger() bool`

HasRanger returns a boolean if a field has been set.

### GetRangerIp

`func (o *DfsHdfsUpdateReqHdfs) GetRangerIp() string`

GetRangerIp returns the RangerIp field if non-nil, zero value otherwise.

### GetRangerIpOk

`func (o *DfsHdfsUpdateReqHdfs) GetRangerIpOk() (*string, bool)`

GetRangerIpOk returns a tuple with the RangerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerIp

`func (o *DfsHdfsUpdateReqHdfs) SetRangerIp(v string)`

SetRangerIp sets RangerIp field to given value.

### HasRangerIp

`func (o *DfsHdfsUpdateReqHdfs) HasRangerIp() bool`

HasRangerIp returns a boolean if a field has been set.

### GetRangerServiceName

`func (o *DfsHdfsUpdateReqHdfs) GetRangerServiceName() string`

GetRangerServiceName returns the RangerServiceName field if non-nil, zero value otherwise.

### GetRangerServiceNameOk

`func (o *DfsHdfsUpdateReqHdfs) GetRangerServiceNameOk() (*string, bool)`

GetRangerServiceNameOk returns a tuple with the RangerServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerServiceName

`func (o *DfsHdfsUpdateReqHdfs) SetRangerServiceName(v string)`

SetRangerServiceName sets RangerServiceName field to given value.

### HasRangerServiceName

`func (o *DfsHdfsUpdateReqHdfs) HasRangerServiceName() bool`

HasRangerServiceName returns a boolean if a field has been set.

### GetRangerUrl

`func (o *DfsHdfsUpdateReqHdfs) GetRangerUrl() string`

GetRangerUrl returns the RangerUrl field if non-nil, zero value otherwise.

### GetRangerUrlOk

`func (o *DfsHdfsUpdateReqHdfs) GetRangerUrlOk() (*string, bool)`

GetRangerUrlOk returns a tuple with the RangerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerUrl

`func (o *DfsHdfsUpdateReqHdfs) SetRangerUrl(v string)`

SetRangerUrl sets RangerUrl field to given value.

### HasRangerUrl

`func (o *DfsHdfsUpdateReqHdfs) HasRangerUrl() bool`

HasRangerUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


