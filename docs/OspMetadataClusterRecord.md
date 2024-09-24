# OspMetadataClusterRecord

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
**Samples** | Pointer to [**[]OspMetadataClusterStat**](OspMetadataClusterStat.md) |  | [optional] 

## Methods

### NewOspMetadataClusterRecord

`func NewOspMetadataClusterRecord(logNum int64, ) *OspMetadataClusterRecord`

NewOspMetadataClusterRecord instantiates a new OspMetadataClusterRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspMetadataClusterRecordWithDefaults

`func NewOspMetadataClusterRecordWithDefaults() *OspMetadataClusterRecord`

NewOspMetadataClusterRecordWithDefaults instantiates a new OspMetadataClusterRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *OspMetadataClusterRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *OspMetadataClusterRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *OspMetadataClusterRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *OspMetadataClusterRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *OspMetadataClusterRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OspMetadataClusterRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OspMetadataClusterRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OspMetadataClusterRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCommitProxyNum

`func (o *OspMetadataClusterRecord) GetCommitProxyNum() int64`

GetCommitProxyNum returns the CommitProxyNum field if non-nil, zero value otherwise.

### GetCommitProxyNumOk

`func (o *OspMetadataClusterRecord) GetCommitProxyNumOk() (*int64, bool)`

GetCommitProxyNumOk returns a tuple with the CommitProxyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitProxyNum

`func (o *OspMetadataClusterRecord) SetCommitProxyNum(v int64)`

SetCommitProxyNum sets CommitProxyNum field to given value.

### HasCommitProxyNum

`func (o *OspMetadataClusterRecord) HasCommitProxyNum() bool`

HasCommitProxyNum returns a boolean if a field has been set.

### GetCoordinatorHostIds

`func (o *OspMetadataClusterRecord) GetCoordinatorHostIds() []int64`

GetCoordinatorHostIds returns the CoordinatorHostIds field if non-nil, zero value otherwise.

### GetCoordinatorHostIdsOk

`func (o *OspMetadataClusterRecord) GetCoordinatorHostIdsOk() (*[]int64, bool)`

GetCoordinatorHostIdsOk returns a tuple with the CoordinatorHostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorHostIds

`func (o *OspMetadataClusterRecord) SetCoordinatorHostIds(v []int64)`

SetCoordinatorHostIds sets CoordinatorHostIds field to given value.

### HasCoordinatorHostIds

`func (o *OspMetadataClusterRecord) HasCoordinatorHostIds() bool`

HasCoordinatorHostIds returns a boolean if a field has been set.

### GetCoordinatorNum

`func (o *OspMetadataClusterRecord) GetCoordinatorNum() int64`

GetCoordinatorNum returns the CoordinatorNum field if non-nil, zero value otherwise.

### GetCoordinatorNumOk

`func (o *OspMetadataClusterRecord) GetCoordinatorNumOk() (*int64, bool)`

GetCoordinatorNumOk returns a tuple with the CoordinatorNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatorNum

`func (o *OspMetadataClusterRecord) SetCoordinatorNum(v int64)`

SetCoordinatorNum sets CoordinatorNum field to given value.

### HasCoordinatorNum

`func (o *OspMetadataClusterRecord) HasCoordinatorNum() bool`

HasCoordinatorNum returns a boolean if a field has been set.

### GetCreate

`func (o *OspMetadataClusterRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OspMetadataClusterRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OspMetadataClusterRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OspMetadataClusterRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCurReplicationNum

`func (o *OspMetadataClusterRecord) GetCurReplicationNum() int64`

GetCurReplicationNum returns the CurReplicationNum field if non-nil, zero value otherwise.

### GetCurReplicationNumOk

`func (o *OspMetadataClusterRecord) GetCurReplicationNumOk() (*int64, bool)`

GetCurReplicationNumOk returns a tuple with the CurReplicationNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurReplicationNum

`func (o *OspMetadataClusterRecord) SetCurReplicationNum(v int64)`

SetCurReplicationNum sets CurReplicationNum field to given value.

### HasCurReplicationNum

`func (o *OspMetadataClusterRecord) HasCurReplicationNum() bool`

HasCurReplicationNum returns a boolean if a field has been set.

### GetDatacenters

`func (o *OspMetadataClusterRecord) GetDatacenters() []PlacementNodeNestview`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *OspMetadataClusterRecord) GetDatacentersOk() (*[]PlacementNodeNestview, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *OspMetadataClusterRecord) SetDatacenters(v []PlacementNodeNestview)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *OspMetadataClusterRecord) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetDeployMode

`func (o *OspMetadataClusterRecord) GetDeployMode() string`

GetDeployMode returns the DeployMode field if non-nil, zero value otherwise.

### GetDeployModeOk

`func (o *OspMetadataClusterRecord) GetDeployModeOk() (*string, bool)`

GetDeployModeOk returns a tuple with the DeployMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployMode

`func (o *OspMetadataClusterRecord) SetDeployMode(v string)`

SetDeployMode sets DeployMode field to given value.

### HasDeployMode

`func (o *OspMetadataClusterRecord) HasDeployMode() bool`

HasDeployMode returns a boolean if a field has been set.

### GetGrvProxyNum

`func (o *OspMetadataClusterRecord) GetGrvProxyNum() int64`

GetGrvProxyNum returns the GrvProxyNum field if non-nil, zero value otherwise.

### GetGrvProxyNumOk

`func (o *OspMetadataClusterRecord) GetGrvProxyNumOk() (*int64, bool)`

GetGrvProxyNumOk returns a tuple with the GrvProxyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrvProxyNum

`func (o *OspMetadataClusterRecord) SetGrvProxyNum(v int64)`

SetGrvProxyNum sets GrvProxyNum field to given value.

### HasGrvProxyNum

`func (o *OspMetadataClusterRecord) HasGrvProxyNum() bool`

HasGrvProxyNum returns a boolean if a field has been set.

### GetId

`func (o *OspMetadataClusterRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OspMetadataClusterRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OspMetadataClusterRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OspMetadataClusterRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogNum

`func (o *OspMetadataClusterRecord) GetLogNum() int64`

GetLogNum returns the LogNum field if non-nil, zero value otherwise.

### GetLogNumOk

`func (o *OspMetadataClusterRecord) GetLogNumOk() (*int64, bool)`

GetLogNumOk returns a tuple with the LogNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNum

`func (o *OspMetadataClusterRecord) SetLogNum(v int64)`

SetLogNum sets LogNum field to given value.


### GetMasterDatacenter

`func (o *OspMetadataClusterRecord) GetMasterDatacenter() PlacementNode`

GetMasterDatacenter returns the MasterDatacenter field if non-nil, zero value otherwise.

### GetMasterDatacenterOk

`func (o *OspMetadataClusterRecord) GetMasterDatacenterOk() (*PlacementNode, bool)`

GetMasterDatacenterOk returns a tuple with the MasterDatacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDatacenter

`func (o *OspMetadataClusterRecord) SetMasterDatacenter(v PlacementNode)`

SetMasterDatacenter sets MasterDatacenter field to given value.

### HasMasterDatacenter

`func (o *OspMetadataClusterRecord) HasMasterDatacenter() bool`

HasMasterDatacenter returns a boolean if a field has been set.

### GetMinAvailableSpaceRatio

`func (o *OspMetadataClusterRecord) GetMinAvailableSpaceRatio() float64`

GetMinAvailableSpaceRatio returns the MinAvailableSpaceRatio field if non-nil, zero value otherwise.

### GetMinAvailableSpaceRatioOk

`func (o *OspMetadataClusterRecord) GetMinAvailableSpaceRatioOk() (*float64, bool)`

GetMinAvailableSpaceRatioOk returns a tuple with the MinAvailableSpaceRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAvailableSpaceRatio

`func (o *OspMetadataClusterRecord) SetMinAvailableSpaceRatio(v float64)`

SetMinAvailableSpaceRatio sets MinAvailableSpaceRatio field to given value.

### HasMinAvailableSpaceRatio

`func (o *OspMetadataClusterRecord) HasMinAvailableSpaceRatio() bool`

HasMinAvailableSpaceRatio returns a boolean if a field has been set.

### GetName

`func (o *OspMetadataClusterRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OspMetadataClusterRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OspMetadataClusterRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OspMetadataClusterRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectStorage

`func (o *OspMetadataClusterRecord) GetObjectStorage() NgObjectStorageNestview`

GetObjectStorage returns the ObjectStorage field if non-nil, zero value otherwise.

### GetObjectStorageOk

`func (o *OspMetadataClusterRecord) GetObjectStorageOk() (*NgObjectStorageNestview, bool)`

GetObjectStorageOk returns a tuple with the ObjectStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorage

`func (o *OspMetadataClusterRecord) SetObjectStorage(v NgObjectStorageNestview)`

SetObjectStorage sets ObjectStorage field to given value.

### HasObjectStorage

`func (o *OspMetadataClusterRecord) HasObjectStorage() bool`

HasObjectStorage returns a boolean if a field has been set.

### GetPlacementNode

`func (o *OspMetadataClusterRecord) GetPlacementNode() OspMdClusterPlacementNodeNestview`

GetPlacementNode returns the PlacementNode field if non-nil, zero value otherwise.

### GetPlacementNodeOk

`func (o *OspMetadataClusterRecord) GetPlacementNodeOk() (*OspMdClusterPlacementNodeNestview, bool)`

GetPlacementNodeOk returns a tuple with the PlacementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementNode

`func (o *OspMetadataClusterRecord) SetPlacementNode(v OspMdClusterPlacementNodeNestview)`

SetPlacementNode sets PlacementNode field to given value.

### HasPlacementNode

`func (o *OspMetadataClusterRecord) HasPlacementNode() bool`

HasPlacementNode returns a boolean if a field has been set.

### GetProtectionType

`func (o *OspMetadataClusterRecord) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *OspMetadataClusterRecord) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *OspMetadataClusterRecord) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *OspMetadataClusterRecord) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### GetRegion

`func (o *OspMetadataClusterRecord) GetRegion() RegionNestview`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OspMetadataClusterRecord) GetRegionOk() (*RegionNestview, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OspMetadataClusterRecord) SetRegion(v RegionNestview)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *OspMetadataClusterRecord) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReplicationNum

`func (o *OspMetadataClusterRecord) GetReplicationNum() int64`

GetReplicationNum returns the ReplicationNum field if non-nil, zero value otherwise.

### GetReplicationNumOk

`func (o *OspMetadataClusterRecord) GetReplicationNumOk() (*int64, bool)`

GetReplicationNumOk returns a tuple with the ReplicationNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationNum

`func (o *OspMetadataClusterRecord) SetReplicationNum(v int64)`

SetReplicationNum sets ReplicationNum field to given value.

### HasReplicationNum

`func (o *OspMetadataClusterRecord) HasReplicationNum() bool`

HasReplicationNum returns a boolean if a field has been set.

### GetResolverNum

`func (o *OspMetadataClusterRecord) GetResolverNum() int64`

GetResolverNum returns the ResolverNum field if non-nil, zero value otherwise.

### GetResolverNumOk

`func (o *OspMetadataClusterRecord) GetResolverNumOk() (*int64, bool)`

GetResolverNumOk returns a tuple with the ResolverNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverNum

`func (o *OspMetadataClusterRecord) SetResolverNum(v int64)`

SetResolverNum sets ResolverNum field to given value.

### HasResolverNum

`func (o *OspMetadataClusterRecord) HasResolverNum() bool`

HasResolverNum returns a boolean if a field has been set.

### GetSpaceUsageFactor

`func (o *OspMetadataClusterRecord) GetSpaceUsageFactor() float64`

GetSpaceUsageFactor returns the SpaceUsageFactor field if non-nil, zero value otherwise.

### GetSpaceUsageFactorOk

`func (o *OspMetadataClusterRecord) GetSpaceUsageFactorOk() (*float64, bool)`

GetSpaceUsageFactorOk returns a tuple with the SpaceUsageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceUsageFactor

`func (o *OspMetadataClusterRecord) SetSpaceUsageFactor(v float64)`

SetSpaceUsageFactor sets SpaceUsageFactor field to given value.

### HasSpaceUsageFactor

`func (o *OspMetadataClusterRecord) HasSpaceUsageFactor() bool`

HasSpaceUsageFactor returns a boolean if a field has been set.

### GetStatelessServerNum

`func (o *OspMetadataClusterRecord) GetStatelessServerNum() int64`

GetStatelessServerNum returns the StatelessServerNum field if non-nil, zero value otherwise.

### GetStatelessServerNumOk

`func (o *OspMetadataClusterRecord) GetStatelessServerNumOk() (*int64, bool)`

GetStatelessServerNumOk returns a tuple with the StatelessServerNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatelessServerNum

`func (o *OspMetadataClusterRecord) SetStatelessServerNum(v int64)`

SetStatelessServerNum sets StatelessServerNum field to given value.

### HasStatelessServerNum

`func (o *OspMetadataClusterRecord) HasStatelessServerNum() bool`

HasStatelessServerNum returns a boolean if a field has been set.

### GetStatus

`func (o *OspMetadataClusterRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OspMetadataClusterRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OspMetadataClusterRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OspMetadataClusterRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *OspMetadataClusterRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OspMetadataClusterRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OspMetadataClusterRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OspMetadataClusterRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *OspMetadataClusterRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OspMetadataClusterRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OspMetadataClusterRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OspMetadataClusterRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *OspMetadataClusterRecord) GetSamples() []OspMetadataClusterStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *OspMetadataClusterRecord) GetSamplesOk() (*[]OspMetadataClusterStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *OspMetadataClusterRecord) SetSamples(v []OspMetadataClusterStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *OspMetadataClusterRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


