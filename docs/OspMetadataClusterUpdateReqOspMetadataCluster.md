# OspMetadataClusterUpdateReqOspMetadataCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoordinatorHostIds** | Pointer to **[]int64** |  | [optional] 
**CoordinatorNum** | Pointer to **int64** |  | [optional] 
**DeployMode** | Pointer to **string** |  | [optional] 
**LogNum** | Pointer to **int64** |  | [optional] 
**NewOsdIds** | Pointer to [**[]OspMetadataClusterService**](OspMetadataClusterService.md) |  | [optional] 
**NewPartitionIds** | Pointer to [**[]OspMetadataClusterService**](OspMetadataClusterService.md) |  | [optional] 
**ProtectionType** | Pointer to **string** |  | [optional] 
**RemoveMetadataServerIds** | Pointer to **[]int64** |  | [optional] 
**RemoveOsdIds** | Pointer to **[]int64** |  | [optional] 
**RemovePartitionIds** | Pointer to **[]int64** |  | [optional] 
**ReplicationNum** | Pointer to **int64** |  | [optional] 

## Methods

### NewOspMetadataClusterUpdateReqOspMetadataCluster

`func NewOspMetadataClusterUpdateReqOspMetadataCluster() *OspMetadataClusterUpdateReqOspMetadataCluster`

NewOspMetadataClusterUpdateReqOspMetadataCluster instantiates a new OspMetadataClusterUpdateReqOspMetadataCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspMetadataClusterUpdateReqOspMetadataClusterWithDefaults

`func NewOspMetadataClusterUpdateReqOspMetadataClusterWithDefaults() *OspMetadataClusterUpdateReqOspMetadataCluster`

NewOspMetadataClusterUpdateReqOspMetadataClusterWithDefaults instantiates a new OspMetadataClusterUpdateReqOspMetadataCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinatorHostIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetCoordinatorHostIds() []int64`

GetCoordinatorHostIds returns the CoordinatorHostIds field if non-nil, zero value otherwise.

### GetCoordinatorHostIdsOk

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetCoordinatorHostIdsOk() (*[]int64, bool)`

GetCoordinatorHostIdsOk returns a tuple with the CoordinatorHostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorHostIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) SetCoordinatorHostIds(v []int64)`

SetCoordinatorHostIds sets CoordinatorHostIds field to given value.

### HasCoordinatorHostIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) HasCoordinatorHostIds() bool`

HasCoordinatorHostIds returns a boolean if a field has been set.

### GetCoordinatorNum

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetCoordinatorNum() int64`

GetCoordinatorNum returns the CoordinatorNum field if non-nil, zero value otherwise.

### GetCoordinatorNumOk

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetCoordinatorNumOk() (*int64, bool)`

GetCoordinatorNumOk returns a tuple with the CoordinatorNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorNum

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) SetCoordinatorNum(v int64)`

SetCoordinatorNum sets CoordinatorNum field to given value.

### HasCoordinatorNum

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) HasCoordinatorNum() bool`

HasCoordinatorNum returns a boolean if a field has been set.

### GetDeployMode

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetDeployMode() string`

GetDeployMode returns the DeployMode field if non-nil, zero value otherwise.

### GetDeployModeOk

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetDeployModeOk() (*string, bool)`

GetDeployModeOk returns a tuple with the DeployMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployMode

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) SetDeployMode(v string)`

SetDeployMode sets DeployMode field to given value.

### HasDeployMode

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) HasDeployMode() bool`

HasDeployMode returns a boolean if a field has been set.

### GetLogNum

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetLogNum() int64`

GetLogNum returns the LogNum field if non-nil, zero value otherwise.

### GetLogNumOk

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetLogNumOk() (*int64, bool)`

GetLogNumOk returns a tuple with the LogNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNum

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) SetLogNum(v int64)`

SetLogNum sets LogNum field to given value.

### HasLogNum

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) HasLogNum() bool`

HasLogNum returns a boolean if a field has been set.

### GetNewOsdIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetNewOsdIds() []OspMetadataClusterService`

GetNewOsdIds returns the NewOsdIds field if non-nil, zero value otherwise.

### GetNewOsdIdsOk

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetNewOsdIdsOk() (*[]OspMetadataClusterService, bool)`

GetNewOsdIdsOk returns a tuple with the NewOsdIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOsdIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) SetNewOsdIds(v []OspMetadataClusterService)`

SetNewOsdIds sets NewOsdIds field to given value.

### HasNewOsdIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) HasNewOsdIds() bool`

HasNewOsdIds returns a boolean if a field has been set.

### GetNewPartitionIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetNewPartitionIds() []OspMetadataClusterService`

GetNewPartitionIds returns the NewPartitionIds field if non-nil, zero value otherwise.

### GetNewPartitionIdsOk

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetNewPartitionIdsOk() (*[]OspMetadataClusterService, bool)`

GetNewPartitionIdsOk returns a tuple with the NewPartitionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPartitionIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) SetNewPartitionIds(v []OspMetadataClusterService)`

SetNewPartitionIds sets NewPartitionIds field to given value.

### HasNewPartitionIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) HasNewPartitionIds() bool`

HasNewPartitionIds returns a boolean if a field has been set.

### GetProtectionType

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### GetRemoveMetadataServerIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetRemoveMetadataServerIds() []int64`

GetRemoveMetadataServerIds returns the RemoveMetadataServerIds field if non-nil, zero value otherwise.

### GetRemoveMetadataServerIdsOk

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetRemoveMetadataServerIdsOk() (*[]int64, bool)`

GetRemoveMetadataServerIdsOk returns a tuple with the RemoveMetadataServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveMetadataServerIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) SetRemoveMetadataServerIds(v []int64)`

SetRemoveMetadataServerIds sets RemoveMetadataServerIds field to given value.

### HasRemoveMetadataServerIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) HasRemoveMetadataServerIds() bool`

HasRemoveMetadataServerIds returns a boolean if a field has been set.

### GetRemoveOsdIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetRemoveOsdIds() []int64`

GetRemoveOsdIds returns the RemoveOsdIds field if non-nil, zero value otherwise.

### GetRemoveOsdIdsOk

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetRemoveOsdIdsOk() (*[]int64, bool)`

GetRemoveOsdIdsOk returns a tuple with the RemoveOsdIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveOsdIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) SetRemoveOsdIds(v []int64)`

SetRemoveOsdIds sets RemoveOsdIds field to given value.

### HasRemoveOsdIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) HasRemoveOsdIds() bool`

HasRemoveOsdIds returns a boolean if a field has been set.

### GetRemovePartitionIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetRemovePartitionIds() []int64`

GetRemovePartitionIds returns the RemovePartitionIds field if non-nil, zero value otherwise.

### GetRemovePartitionIdsOk

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetRemovePartitionIdsOk() (*[]int64, bool)`

GetRemovePartitionIdsOk returns a tuple with the RemovePartitionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePartitionIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) SetRemovePartitionIds(v []int64)`

SetRemovePartitionIds sets RemovePartitionIds field to given value.

### HasRemovePartitionIds

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) HasRemovePartitionIds() bool`

HasRemovePartitionIds returns a boolean if a field has been set.

### GetReplicationNum

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetReplicationNum() int64`

GetReplicationNum returns the ReplicationNum field if non-nil, zero value otherwise.

### GetReplicationNumOk

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) GetReplicationNumOk() (*int64, bool)`

GetReplicationNumOk returns a tuple with the ReplicationNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationNum

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) SetReplicationNum(v int64)`

SetReplicationNum sets ReplicationNum field to given value.

### HasReplicationNum

`func (o *OspMetadataClusterUpdateReqOspMetadataCluster) HasReplicationNum() bool`

HasReplicationNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


