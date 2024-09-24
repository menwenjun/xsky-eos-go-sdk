# OspMetadataClusterCreateReqOspMetadataCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **int64** |  | [optional] 
**CommitProxyNum** | **int64** |  | 
**CoordinatorHostIds** | **[]int64** |  | 
**CoordinatorNum** | **int64** |  | 
**DeployMode** | Pointer to **string** |  | [optional] 
**GrvProxyNum** | **int64** |  | 
**LogNum** | **int64** |  | 
**Name** | **string** |  | 
**OsdIds** | Pointer to [**[]OspMetadataClusterService**](OspMetadataClusterService.md) |  | [optional] 
**PartitionIds** | Pointer to [**[]OspMetadataClusterService**](OspMetadataClusterService.md) |  | [optional] 
**ProtectionType** | **string** |  | 
**ReplicationNum** | **int64** |  | 
**ResolverNum** | **int64** |  | 
**StatelessServerNum** | **int64** |  | 
**Type** | **string** |  | 

## Methods

### NewOspMetadataClusterCreateReqOspMetadataCluster

`func NewOspMetadataClusterCreateReqOspMetadataCluster(commitProxyNum int64, coordinatorHostIds []int64, coordinatorNum int64, grvProxyNum int64, logNum int64, name string, protectionType string, replicationNum int64, resolverNum int64, statelessServerNum int64, type_ string, ) *OspMetadataClusterCreateReqOspMetadataCluster`

NewOspMetadataClusterCreateReqOspMetadataCluster instantiates a new OspMetadataClusterCreateReqOspMetadataCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspMetadataClusterCreateReqOspMetadataClusterWithDefaults

`func NewOspMetadataClusterCreateReqOspMetadataClusterWithDefaults() *OspMetadataClusterCreateReqOspMetadataCluster`

NewOspMetadataClusterCreateReqOspMetadataClusterWithDefaults instantiates a new OspMetadataClusterCreateReqOspMetadataCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCommitProxyNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCommitProxyNum() int64`

GetCommitProxyNum returns the CommitProxyNum field if non-nil, zero value otherwise.

### GetCommitProxyNumOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCommitProxyNumOk() (*int64, bool)`

GetCommitProxyNumOk returns a tuple with the CommitProxyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitProxyNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetCommitProxyNum(v int64)`

SetCommitProxyNum sets CommitProxyNum field to given value.


### GetCoordinatorHostIds

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCoordinatorHostIds() []int64`

GetCoordinatorHostIds returns the CoordinatorHostIds field if non-nil, zero value otherwise.

### GetCoordinatorHostIdsOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCoordinatorHostIdsOk() (*[]int64, bool)`

GetCoordinatorHostIdsOk returns a tuple with the CoordinatorHostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorHostIds

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetCoordinatorHostIds(v []int64)`

SetCoordinatorHostIds sets CoordinatorHostIds field to given value.


### GetCoordinatorNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCoordinatorNum() int64`

GetCoordinatorNum returns the CoordinatorNum field if non-nil, zero value otherwise.

### GetCoordinatorNumOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetCoordinatorNumOk() (*int64, bool)`

GetCoordinatorNumOk returns a tuple with the CoordinatorNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetCoordinatorNum(v int64)`

SetCoordinatorNum sets CoordinatorNum field to given value.


### GetDeployMode

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetDeployMode() string`

GetDeployMode returns the DeployMode field if non-nil, zero value otherwise.

### GetDeployModeOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetDeployModeOk() (*string, bool)`

GetDeployModeOk returns a tuple with the DeployMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployMode

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetDeployMode(v string)`

SetDeployMode sets DeployMode field to given value.

### HasDeployMode

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) HasDeployMode() bool`

HasDeployMode returns a boolean if a field has been set.

### GetGrvProxyNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetGrvProxyNum() int64`

GetGrvProxyNum returns the GrvProxyNum field if non-nil, zero value otherwise.

### GetGrvProxyNumOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetGrvProxyNumOk() (*int64, bool)`

GetGrvProxyNumOk returns a tuple with the GrvProxyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrvProxyNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetGrvProxyNum(v int64)`

SetGrvProxyNum sets GrvProxyNum field to given value.


### GetLogNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetLogNum() int64`

GetLogNum returns the LogNum field if non-nil, zero value otherwise.

### GetLogNumOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetLogNumOk() (*int64, bool)`

GetLogNumOk returns a tuple with the LogNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetLogNum(v int64)`

SetLogNum sets LogNum field to given value.


### GetName

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetName(v string)`

SetName sets Name field to given value.


### GetOsdIds

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetOsdIds() []OspMetadataClusterService`

GetOsdIds returns the OsdIds field if non-nil, zero value otherwise.

### GetOsdIdsOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetOsdIdsOk() (*[]OspMetadataClusterService, bool)`

GetOsdIdsOk returns a tuple with the OsdIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdIds

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetOsdIds(v []OspMetadataClusterService)`

SetOsdIds sets OsdIds field to given value.

### HasOsdIds

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) HasOsdIds() bool`

HasOsdIds returns a boolean if a field has been set.

### GetPartitionIds

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetPartitionIds() []OspMetadataClusterService`

GetPartitionIds returns the PartitionIds field if non-nil, zero value otherwise.

### GetPartitionIdsOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetPartitionIdsOk() (*[]OspMetadataClusterService, bool)`

GetPartitionIdsOk returns a tuple with the PartitionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionIds

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetPartitionIds(v []OspMetadataClusterService)`

SetPartitionIds sets PartitionIds field to given value.

### HasPartitionIds

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) HasPartitionIds() bool`

HasPartitionIds returns a boolean if a field has been set.

### GetProtectionType

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.


### GetReplicationNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetReplicationNum() int64`

GetReplicationNum returns the ReplicationNum field if non-nil, zero value otherwise.

### GetReplicationNumOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetReplicationNumOk() (*int64, bool)`

GetReplicationNumOk returns a tuple with the ReplicationNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetReplicationNum(v int64)`

SetReplicationNum sets ReplicationNum field to given value.


### GetResolverNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetResolverNum() int64`

GetResolverNum returns the ResolverNum field if non-nil, zero value otherwise.

### GetResolverNumOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetResolverNumOk() (*int64, bool)`

GetResolverNumOk returns a tuple with the ResolverNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetResolverNum(v int64)`

SetResolverNum sets ResolverNum field to given value.


### GetStatelessServerNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetStatelessServerNum() int64`

GetStatelessServerNum returns the StatelessServerNum field if non-nil, zero value otherwise.

### GetStatelessServerNumOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetStatelessServerNumOk() (*int64, bool)`

GetStatelessServerNumOk returns a tuple with the StatelessServerNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatelessServerNum

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetStatelessServerNum(v int64)`

SetStatelessServerNum sets StatelessServerNum field to given value.


### GetType

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OspMetadataClusterCreateReqOspMetadataCluster) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


