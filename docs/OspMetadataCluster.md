# OspMetadataCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**CommitProxyNum** | Pointer to **int64** |  | [optional] 
**CoordinatorHostIds** | Pointer to **[]int64** |  | [optional] 
**CoordinatorNum** | Pointer to **int64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**CurReplicationNum** | Pointer to **int64** |  | [optional] 
**Datacenters** | Pointer to [**[]PlacementNodeNestview**](PlacementNodeNestview.md) |  | [optional] 
**DeployMode** | Pointer to **string** |  | [optional] 
**GrvProxyNum** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LogNum** | **int64** |  | 
**MasterDatacenter** | Pointer to [**PlacementNode**](PlacementNode.md) |  | [optional] 
**MinAvailableSpaceRatio** | Pointer to **float64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectStorage** | Pointer to [**NgObjectStorageNestview**](NgObjectStorageNestview.md) |  | [optional] 
**PlacementNode** | Pointer to [**OspMdClusterPlacementNodeNestview**](OspMdClusterPlacementNodeNestview.md) |  | [optional] 
**ProtectionType** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**RegionNestview**](RegionNestview.md) |  | [optional] 
**ReplicationNum** | Pointer to **int64** |  | [optional] 
**ResolverNum** | Pointer to **int64** |  | [optional] 
**SpaceUsageFactor** | Pointer to **float64** |  | [optional] 
**StatelessServerNum** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOspMetadataCluster

`func NewOspMetadataCluster(logNum int64, ) *OspMetadataCluster`

NewOspMetadataCluster instantiates a new OspMetadataCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspMetadataClusterWithDefaults

`func NewOspMetadataClusterWithDefaults() *OspMetadataCluster`

NewOspMetadataClusterWithDefaults instantiates a new OspMetadataCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *OspMetadataCluster) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *OspMetadataCluster) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *OspMetadataCluster) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *OspMetadataCluster) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *OspMetadataCluster) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OspMetadataCluster) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OspMetadataCluster) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OspMetadataCluster) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCommitProxyNum

`func (o *OspMetadataCluster) GetCommitProxyNum() int64`

GetCommitProxyNum returns the CommitProxyNum field if non-nil, zero value otherwise.

### GetCommitProxyNumOk

`func (o *OspMetadataCluster) GetCommitProxyNumOk() (*int64, bool)`

GetCommitProxyNumOk returns a tuple with the CommitProxyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitProxyNum

`func (o *OspMetadataCluster) SetCommitProxyNum(v int64)`

SetCommitProxyNum sets CommitProxyNum field to given value.

### HasCommitProxyNum

`func (o *OspMetadataCluster) HasCommitProxyNum() bool`

HasCommitProxyNum returns a boolean if a field has been set.

### GetCoordinatorHostIds

`func (o *OspMetadataCluster) GetCoordinatorHostIds() []int64`

GetCoordinatorHostIds returns the CoordinatorHostIds field if non-nil, zero value otherwise.

### GetCoordinatorHostIdsOk

`func (o *OspMetadataCluster) GetCoordinatorHostIdsOk() (*[]int64, bool)`

GetCoordinatorHostIdsOk returns a tuple with the CoordinatorHostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorHostIds

`func (o *OspMetadataCluster) SetCoordinatorHostIds(v []int64)`

SetCoordinatorHostIds sets CoordinatorHostIds field to given value.

### HasCoordinatorHostIds

`func (o *OspMetadataCluster) HasCoordinatorHostIds() bool`

HasCoordinatorHostIds returns a boolean if a field has been set.

### GetCoordinatorNum

`func (o *OspMetadataCluster) GetCoordinatorNum() int64`

GetCoordinatorNum returns the CoordinatorNum field if non-nil, zero value otherwise.

### GetCoordinatorNumOk

`func (o *OspMetadataCluster) GetCoordinatorNumOk() (*int64, bool)`

GetCoordinatorNumOk returns a tuple with the CoordinatorNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorNum

`func (o *OspMetadataCluster) SetCoordinatorNum(v int64)`

SetCoordinatorNum sets CoordinatorNum field to given value.

### HasCoordinatorNum

`func (o *OspMetadataCluster) HasCoordinatorNum() bool`

HasCoordinatorNum returns a boolean if a field has been set.

### GetCreate

`func (o *OspMetadataCluster) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OspMetadataCluster) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OspMetadataCluster) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OspMetadataCluster) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCurReplicationNum

`func (o *OspMetadataCluster) GetCurReplicationNum() int64`

GetCurReplicationNum returns the CurReplicationNum field if non-nil, zero value otherwise.

### GetCurReplicationNumOk

`func (o *OspMetadataCluster) GetCurReplicationNumOk() (*int64, bool)`

GetCurReplicationNumOk returns a tuple with the CurReplicationNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurReplicationNum

`func (o *OspMetadataCluster) SetCurReplicationNum(v int64)`

SetCurReplicationNum sets CurReplicationNum field to given value.

### HasCurReplicationNum

`func (o *OspMetadataCluster) HasCurReplicationNum() bool`

HasCurReplicationNum returns a boolean if a field has been set.

### GetDatacenters

`func (o *OspMetadataCluster) GetDatacenters() []PlacementNodeNestview`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *OspMetadataCluster) GetDatacentersOk() (*[]PlacementNodeNestview, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *OspMetadataCluster) SetDatacenters(v []PlacementNodeNestview)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *OspMetadataCluster) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetDeployMode

`func (o *OspMetadataCluster) GetDeployMode() string`

GetDeployMode returns the DeployMode field if non-nil, zero value otherwise.

### GetDeployModeOk

`func (o *OspMetadataCluster) GetDeployModeOk() (*string, bool)`

GetDeployModeOk returns a tuple with the DeployMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployMode

`func (o *OspMetadataCluster) SetDeployMode(v string)`

SetDeployMode sets DeployMode field to given value.

### HasDeployMode

`func (o *OspMetadataCluster) HasDeployMode() bool`

HasDeployMode returns a boolean if a field has been set.

### GetGrvProxyNum

`func (o *OspMetadataCluster) GetGrvProxyNum() int64`

GetGrvProxyNum returns the GrvProxyNum field if non-nil, zero value otherwise.

### GetGrvProxyNumOk

`func (o *OspMetadataCluster) GetGrvProxyNumOk() (*int64, bool)`

GetGrvProxyNumOk returns a tuple with the GrvProxyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrvProxyNum

`func (o *OspMetadataCluster) SetGrvProxyNum(v int64)`

SetGrvProxyNum sets GrvProxyNum field to given value.

### HasGrvProxyNum

`func (o *OspMetadataCluster) HasGrvProxyNum() bool`

HasGrvProxyNum returns a boolean if a field has been set.

### GetId

`func (o *OspMetadataCluster) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OspMetadataCluster) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OspMetadataCluster) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OspMetadataCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogNum

`func (o *OspMetadataCluster) GetLogNum() int64`

GetLogNum returns the LogNum field if non-nil, zero value otherwise.

### GetLogNumOk

`func (o *OspMetadataCluster) GetLogNumOk() (*int64, bool)`

GetLogNumOk returns a tuple with the LogNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNum

`func (o *OspMetadataCluster) SetLogNum(v int64)`

SetLogNum sets LogNum field to given value.


### GetMasterDatacenter

`func (o *OspMetadataCluster) GetMasterDatacenter() PlacementNode`

GetMasterDatacenter returns the MasterDatacenter field if non-nil, zero value otherwise.

### GetMasterDatacenterOk

`func (o *OspMetadataCluster) GetMasterDatacenterOk() (*PlacementNode, bool)`

GetMasterDatacenterOk returns a tuple with the MasterDatacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDatacenter

`func (o *OspMetadataCluster) SetMasterDatacenter(v PlacementNode)`

SetMasterDatacenter sets MasterDatacenter field to given value.

### HasMasterDatacenter

`func (o *OspMetadataCluster) HasMasterDatacenter() bool`

HasMasterDatacenter returns a boolean if a field has been set.

### GetMinAvailableSpaceRatio

`func (o *OspMetadataCluster) GetMinAvailableSpaceRatio() float64`

GetMinAvailableSpaceRatio returns the MinAvailableSpaceRatio field if non-nil, zero value otherwise.

### GetMinAvailableSpaceRatioOk

`func (o *OspMetadataCluster) GetMinAvailableSpaceRatioOk() (*float64, bool)`

GetMinAvailableSpaceRatioOk returns a tuple with the MinAvailableSpaceRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAvailableSpaceRatio

`func (o *OspMetadataCluster) SetMinAvailableSpaceRatio(v float64)`

SetMinAvailableSpaceRatio sets MinAvailableSpaceRatio field to given value.

### HasMinAvailableSpaceRatio

`func (o *OspMetadataCluster) HasMinAvailableSpaceRatio() bool`

HasMinAvailableSpaceRatio returns a boolean if a field has been set.

### GetName

`func (o *OspMetadataCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OspMetadataCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OspMetadataCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OspMetadataCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectStorage

`func (o *OspMetadataCluster) GetObjectStorage() NgObjectStorageNestview`

GetObjectStorage returns the ObjectStorage field if non-nil, zero value otherwise.

### GetObjectStorageOk

`func (o *OspMetadataCluster) GetObjectStorageOk() (*NgObjectStorageNestview, bool)`

GetObjectStorageOk returns a tuple with the ObjectStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorage

`func (o *OspMetadataCluster) SetObjectStorage(v NgObjectStorageNestview)`

SetObjectStorage sets ObjectStorage field to given value.

### HasObjectStorage

`func (o *OspMetadataCluster) HasObjectStorage() bool`

HasObjectStorage returns a boolean if a field has been set.

### GetPlacementNode

`func (o *OspMetadataCluster) GetPlacementNode() OspMdClusterPlacementNodeNestview`

GetPlacementNode returns the PlacementNode field if non-nil, zero value otherwise.

### GetPlacementNodeOk

`func (o *OspMetadataCluster) GetPlacementNodeOk() (*OspMdClusterPlacementNodeNestview, bool)`

GetPlacementNodeOk returns a tuple with the PlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementNode

`func (o *OspMetadataCluster) SetPlacementNode(v OspMdClusterPlacementNodeNestview)`

SetPlacementNode sets PlacementNode field to given value.

### HasPlacementNode

`func (o *OspMetadataCluster) HasPlacementNode() bool`

HasPlacementNode returns a boolean if a field has been set.

### GetProtectionType

`func (o *OspMetadataCluster) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *OspMetadataCluster) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *OspMetadataCluster) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *OspMetadataCluster) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### GetRegion

`func (o *OspMetadataCluster) GetRegion() RegionNestview`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OspMetadataCluster) GetRegionOk() (*RegionNestview, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OspMetadataCluster) SetRegion(v RegionNestview)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *OspMetadataCluster) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReplicationNum

`func (o *OspMetadataCluster) GetReplicationNum() int64`

GetReplicationNum returns the ReplicationNum field if non-nil, zero value otherwise.

### GetReplicationNumOk

`func (o *OspMetadataCluster) GetReplicationNumOk() (*int64, bool)`

GetReplicationNumOk returns a tuple with the ReplicationNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationNum

`func (o *OspMetadataCluster) SetReplicationNum(v int64)`

SetReplicationNum sets ReplicationNum field to given value.

### HasReplicationNum

`func (o *OspMetadataCluster) HasReplicationNum() bool`

HasReplicationNum returns a boolean if a field has been set.

### GetResolverNum

`func (o *OspMetadataCluster) GetResolverNum() int64`

GetResolverNum returns the ResolverNum field if non-nil, zero value otherwise.

### GetResolverNumOk

`func (o *OspMetadataCluster) GetResolverNumOk() (*int64, bool)`

GetResolverNumOk returns a tuple with the ResolverNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverNum

`func (o *OspMetadataCluster) SetResolverNum(v int64)`

SetResolverNum sets ResolverNum field to given value.

### HasResolverNum

`func (o *OspMetadataCluster) HasResolverNum() bool`

HasResolverNum returns a boolean if a field has been set.

### GetSpaceUsageFactor

`func (o *OspMetadataCluster) GetSpaceUsageFactor() float64`

GetSpaceUsageFactor returns the SpaceUsageFactor field if non-nil, zero value otherwise.

### GetSpaceUsageFactorOk

`func (o *OspMetadataCluster) GetSpaceUsageFactorOk() (*float64, bool)`

GetSpaceUsageFactorOk returns a tuple with the SpaceUsageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceUsageFactor

`func (o *OspMetadataCluster) SetSpaceUsageFactor(v float64)`

SetSpaceUsageFactor sets SpaceUsageFactor field to given value.

### HasSpaceUsageFactor

`func (o *OspMetadataCluster) HasSpaceUsageFactor() bool`

HasSpaceUsageFactor returns a boolean if a field has been set.

### GetStatelessServerNum

`func (o *OspMetadataCluster) GetStatelessServerNum() int64`

GetStatelessServerNum returns the StatelessServerNum field if non-nil, zero value otherwise.

### GetStatelessServerNumOk

`func (o *OspMetadataCluster) GetStatelessServerNumOk() (*int64, bool)`

GetStatelessServerNumOk returns a tuple with the StatelessServerNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatelessServerNum

`func (o *OspMetadataCluster) SetStatelessServerNum(v int64)`

SetStatelessServerNum sets StatelessServerNum field to given value.

### HasStatelessServerNum

`func (o *OspMetadataCluster) HasStatelessServerNum() bool`

HasStatelessServerNum returns a boolean if a field has been set.

### GetStatus

`func (o *OspMetadataCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OspMetadataCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OspMetadataCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OspMetadataCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *OspMetadataCluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OspMetadataCluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OspMetadataCluster) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OspMetadataCluster) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *OspMetadataCluster) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OspMetadataCluster) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OspMetadataCluster) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OspMetadataCluster) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


