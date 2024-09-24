# DfsGatewayConnectionsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FtpSessions** | Pointer to [**[]DfsFTPSession**](DfsFTPSession.md) | dfs ftp sessions | [optional] 
**HdfsConnections** | Pointer to [**[]DfsHdfsConnection**](DfsHdfsConnection.md) | dfs hdfs connections | [optional] 
**NfsConnections** | Pointer to [**[]DfsNFSConnection**](DfsNFSConnection.md) | dfs nfs connections | [optional] 
**S3Connections** | Pointer to [**[]DfsS3BucketConnection**](DfsS3BucketConnection.md) | dfs s3 bucket connections | [optional] 
**SmbSessions** | Pointer to [**[]DfsSMBSession**](DfsSMBSession.md) | dfs smb sessions | [optional] 

## Methods

### NewDfsGatewayConnectionsResp

`func NewDfsGatewayConnectionsResp() *DfsGatewayConnectionsResp`

NewDfsGatewayConnectionsResp instantiates a new DfsGatewayConnectionsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayConnectionsRespWithDefaults

`func NewDfsGatewayConnectionsRespWithDefaults() *DfsGatewayConnectionsResp`

NewDfsGatewayConnectionsRespWithDefaults instantiates a new DfsGatewayConnectionsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFtpSessions

`func (o *DfsGatewayConnectionsResp) GetFtpSessions() []DfsFTPSession`

GetFtpSessions returns the FtpSessions field if non-nil, zero value otherwise.

### GetFtpSessionsOk

`func (o *DfsGatewayConnectionsResp) GetFtpSessionsOk() (*[]DfsFTPSession, bool)`

GetFtpSessionsOk returns a tuple with the FtpSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpSessions

`func (o *DfsGatewayConnectionsResp) SetFtpSessions(v []DfsFTPSession)`

SetFtpSessions sets FtpSessions field to given value.

### HasFtpSessions

`func (o *DfsGatewayConnectionsResp) HasFtpSessions() bool`

HasFtpSessions returns a boolean if a field has been set.

### GetHdfsConnections

`func (o *DfsGatewayConnectionsResp) GetHdfsConnections() []DfsHdfsConnection`

GetHdfsConnections returns the HdfsConnections field if non-nil, zero value otherwise.

### GetHdfsConnectionsOk

`func (o *DfsGatewayConnectionsResp) GetHdfsConnectionsOk() (*[]DfsHdfsConnection, bool)`

GetHdfsConnectionsOk returns a tuple with the HdfsConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsConnections

`func (o *DfsGatewayConnectionsResp) SetHdfsConnections(v []DfsHdfsConnection)`

SetHdfsConnections sets HdfsConnections field to given value.

### HasHdfsConnections

`func (o *DfsGatewayConnectionsResp) HasHdfsConnections() bool`

HasHdfsConnections returns a boolean if a field has been set.

### GetNfsConnections

`func (o *DfsGatewayConnectionsResp) GetNfsConnections() []DfsNFSConnection`

GetNfsConnections returns the NfsConnections field if non-nil, zero value otherwise.

### GetNfsConnectionsOk

`func (o *DfsGatewayConnectionsResp) GetNfsConnectionsOk() (*[]DfsNFSConnection, bool)`

GetNfsConnectionsOk returns a tuple with the NfsConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsConnections

`func (o *DfsGatewayConnectionsResp) SetNfsConnections(v []DfsNFSConnection)`

SetNfsConnections sets NfsConnections field to given value.

### HasNfsConnections

`func (o *DfsGatewayConnectionsResp) HasNfsConnections() bool`

HasNfsConnections returns a boolean if a field has been set.

### GetS3Connections

`func (o *DfsGatewayConnectionsResp) GetS3Connections() []DfsS3BucketConnection`

GetS3Connections returns the S3Connections field if non-nil, zero value otherwise.

### GetS3ConnectionsOk

`func (o *DfsGatewayConnectionsResp) GetS3ConnectionsOk() (*[]DfsS3BucketConnection, bool)`

GetS3ConnectionsOk returns a tuple with the S3Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Connections

`func (o *DfsGatewayConnectionsResp) SetS3Connections(v []DfsS3BucketConnection)`

SetS3Connections sets S3Connections field to given value.

### HasS3Connections

`func (o *DfsGatewayConnectionsResp) HasS3Connections() bool`

HasS3Connections returns a boolean if a field has been set.

### GetSmbSessions

`func (o *DfsGatewayConnectionsResp) GetSmbSessions() []DfsSMBSession`

GetSmbSessions returns the SmbSessions field if non-nil, zero value otherwise.

### GetSmbSessionsOk

`func (o *DfsGatewayConnectionsResp) GetSmbSessionsOk() (*[]DfsSMBSession, bool)`

GetSmbSessionsOk returns a tuple with the SmbSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbSessions

`func (o *DfsGatewayConnectionsResp) SetSmbSessions(v []DfsSMBSession)`

SetSmbSessions sets SmbSessions field to given value.

### HasSmbSessions

`func (o *DfsGatewayConnectionsResp) HasSmbSessions() bool`

HasSmbSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


