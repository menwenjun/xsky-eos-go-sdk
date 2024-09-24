# DfsRootfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DfsGatewayGroup** | Pointer to [**DfsGatewayGroupNestview**](DfsGatewayGroupNestview.md) |  | [optional] 
**GatewayGroupIds** | Pointer to **[]int64** |  | [optional] 
**GcSpeed** | Pointer to **string** |  | [optional] 
**HdfsNum** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MetadataCluster** | Pointer to [**MetadataClusterNestview**](MetadataClusterNestview.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PriorityPolicyIds** | Pointer to **[]int64** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**ProgressInfo** | Pointer to [**ProgressInfo**](ProgressInfo.md) |  | [optional] 
**QosNum** | Pointer to **int64** |  | [optional] 
**S3BucketNum** | Pointer to **int64** |  | [optional] 
**ShareNums** | Pointer to **map[string]int64** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StoragePolicyNum** | Pointer to **int64** |  | [optional] 
**TierNum** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**WormLogPath** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsRootfs

`func NewDfsRootfs() *DfsRootfs`

NewDfsRootfs instantiates a new DfsRootfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsRootfsWithDefaults

`func NewDfsRootfsWithDefaults() *DfsRootfs`

NewDfsRootfsWithDefaults instantiates a new DfsRootfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsRootfs) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsRootfs) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsRootfs) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsRootfs) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DfsRootfs) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsRootfs) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsRootfs) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsRootfs) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsRootfs) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsRootfs) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsRootfs) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsRootfs) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DfsRootfs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsRootfs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsRootfs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsRootfs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsRootfs) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsRootfs) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsRootfs) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsRootfs) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetGatewayGroupIds

`func (o *DfsRootfs) GetGatewayGroupIds() []int64`

GetGatewayGroupIds returns the GatewayGroupIds field if non-nil, zero value otherwise.

### GetGatewayGroupIdsOk

`func (o *DfsRootfs) GetGatewayGroupIdsOk() (*[]int64, bool)`

GetGatewayGroupIdsOk returns a tuple with the GatewayGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayGroupIds

`func (o *DfsRootfs) SetGatewayGroupIds(v []int64)`

SetGatewayGroupIds sets GatewayGroupIds field to given value.

### HasGatewayGroupIds

`func (o *DfsRootfs) HasGatewayGroupIds() bool`

HasGatewayGroupIds returns a boolean if a field has been set.

### GetGcSpeed

`func (o *DfsRootfs) GetGcSpeed() string`

GetGcSpeed returns the GcSpeed field if non-nil, zero value otherwise.

### GetGcSpeedOk

`func (o *DfsRootfs) GetGcSpeedOk() (*string, bool)`

GetGcSpeedOk returns a tuple with the GcSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcSpeed

`func (o *DfsRootfs) SetGcSpeed(v string)`

SetGcSpeed sets GcSpeed field to given value.

### HasGcSpeed

`func (o *DfsRootfs) HasGcSpeed() bool`

HasGcSpeed returns a boolean if a field has been set.

### GetHdfsNum

`func (o *DfsRootfs) GetHdfsNum() int64`

GetHdfsNum returns the HdfsNum field if non-nil, zero value otherwise.

### GetHdfsNumOk

`func (o *DfsRootfs) GetHdfsNumOk() (*int64, bool)`

GetHdfsNumOk returns a tuple with the HdfsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsNum

`func (o *DfsRootfs) SetHdfsNum(v int64)`

SetHdfsNum sets HdfsNum field to given value.

### HasHdfsNum

`func (o *DfsRootfs) HasHdfsNum() bool`

HasHdfsNum returns a boolean if a field has been set.

### GetId

`func (o *DfsRootfs) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsRootfs) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsRootfs) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsRootfs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadataCluster

`func (o *DfsRootfs) GetMetadataCluster() MetadataClusterNestview`

GetMetadataCluster returns the MetadataCluster field if non-nil, zero value otherwise.

### GetMetadataClusterOk

`func (o *DfsRootfs) GetMetadataClusterOk() (*MetadataClusterNestview, bool)`

GetMetadataClusterOk returns a tuple with the MetadataCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCluster

`func (o *DfsRootfs) SetMetadataCluster(v MetadataClusterNestview)`

SetMetadataCluster sets MetadataCluster field to given value.

### HasMetadataCluster

`func (o *DfsRootfs) HasMetadataCluster() bool`

HasMetadataCluster returns a boolean if a field has been set.

### GetName

`func (o *DfsRootfs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsRootfs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsRootfs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsRootfs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriorityPolicyIds

`func (o *DfsRootfs) GetPriorityPolicyIds() []int64`

GetPriorityPolicyIds returns the PriorityPolicyIds field if non-nil, zero value otherwise.

### GetPriorityPolicyIdsOk

`func (o *DfsRootfs) GetPriorityPolicyIdsOk() (*[]int64, bool)`

GetPriorityPolicyIdsOk returns a tuple with the PriorityPolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityPolicyIds

`func (o *DfsRootfs) SetPriorityPolicyIds(v []int64)`

SetPriorityPolicyIds sets PriorityPolicyIds field to given value.

### HasPriorityPolicyIds

`func (o *DfsRootfs) HasPriorityPolicyIds() bool`

HasPriorityPolicyIds returns a boolean if a field has been set.

### GetProgress

`func (o *DfsRootfs) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DfsRootfs) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DfsRootfs) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DfsRootfs) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProgressInfo

`func (o *DfsRootfs) GetProgressInfo() ProgressInfo`

GetProgressInfo returns the ProgressInfo field if non-nil, zero value otherwise.

### GetProgressInfoOk

`func (o *DfsRootfs) GetProgressInfoOk() (*ProgressInfo, bool)`

GetProgressInfoOk returns a tuple with the ProgressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressInfo

`func (o *DfsRootfs) SetProgressInfo(v ProgressInfo)`

SetProgressInfo sets ProgressInfo field to given value.

### HasProgressInfo

`func (o *DfsRootfs) HasProgressInfo() bool`

HasProgressInfo returns a boolean if a field has been set.

### GetQosNum

`func (o *DfsRootfs) GetQosNum() int64`

GetQosNum returns the QosNum field if non-nil, zero value otherwise.

### GetQosNumOk

`func (o *DfsRootfs) GetQosNumOk() (*int64, bool)`

GetQosNumOk returns a tuple with the QosNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosNum

`func (o *DfsRootfs) SetQosNum(v int64)`

SetQosNum sets QosNum field to given value.

### HasQosNum

`func (o *DfsRootfs) HasQosNum() bool`

HasQosNum returns a boolean if a field has been set.

### GetS3BucketNum

`func (o *DfsRootfs) GetS3BucketNum() int64`

GetS3BucketNum returns the S3BucketNum field if non-nil, zero value otherwise.

### GetS3BucketNumOk

`func (o *DfsRootfs) GetS3BucketNumOk() (*int64, bool)`

GetS3BucketNumOk returns a tuple with the S3BucketNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketNum

`func (o *DfsRootfs) SetS3BucketNum(v int64)`

SetS3BucketNum sets S3BucketNum field to given value.

### HasS3BucketNum

`func (o *DfsRootfs) HasS3BucketNum() bool`

HasS3BucketNum returns a boolean if a field has been set.

### GetShareNums

`func (o *DfsRootfs) GetShareNums() map[string]int64`

GetShareNums returns the ShareNums field if non-nil, zero value otherwise.

### GetShareNumsOk

`func (o *DfsRootfs) GetShareNumsOk() (*map[string]int64, bool)`

GetShareNumsOk returns a tuple with the ShareNums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareNums

`func (o *DfsRootfs) SetShareNums(v map[string]int64)`

SetShareNums sets ShareNums field to given value.

### HasShareNums

`func (o *DfsRootfs) HasShareNums() bool`

HasShareNums returns a boolean if a field has been set.

### GetShared

`func (o *DfsRootfs) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *DfsRootfs) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *DfsRootfs) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *DfsRootfs) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSize

`func (o *DfsRootfs) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DfsRootfs) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DfsRootfs) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DfsRootfs) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *DfsRootfs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsRootfs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsRootfs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsRootfs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStoragePolicyNum

`func (o *DfsRootfs) GetStoragePolicyNum() int64`

GetStoragePolicyNum returns the StoragePolicyNum field if non-nil, zero value otherwise.

### GetStoragePolicyNumOk

`func (o *DfsRootfs) GetStoragePolicyNumOk() (*int64, bool)`

GetStoragePolicyNumOk returns a tuple with the StoragePolicyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyNum

`func (o *DfsRootfs) SetStoragePolicyNum(v int64)`

SetStoragePolicyNum sets StoragePolicyNum field to given value.

### HasStoragePolicyNum

`func (o *DfsRootfs) HasStoragePolicyNum() bool`

HasStoragePolicyNum returns a boolean if a field has been set.

### GetTierNum

`func (o *DfsRootfs) GetTierNum() int64`

GetTierNum returns the TierNum field if non-nil, zero value otherwise.

### GetTierNumOk

`func (o *DfsRootfs) GetTierNumOk() (*int64, bool)`

GetTierNumOk returns a tuple with the TierNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierNum

`func (o *DfsRootfs) SetTierNum(v int64)`

SetTierNum sets TierNum field to given value.

### HasTierNum

`func (o *DfsRootfs) HasTierNum() bool`

HasTierNum returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsRootfs) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsRootfs) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsRootfs) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsRootfs) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUuid

`func (o *DfsRootfs) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DfsRootfs) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DfsRootfs) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DfsRootfs) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWormLogPath

`func (o *DfsRootfs) GetWormLogPath() string`

GetWormLogPath returns the WormLogPath field if non-nil, zero value otherwise.

### GetWormLogPathOk

`func (o *DfsRootfs) GetWormLogPathOk() (*string, bool)`

GetWormLogPathOk returns a tuple with the WormLogPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWormLogPath

`func (o *DfsRootfs) SetWormLogPath(v string)`

SetWormLogPath sets WormLogPath field to given value.

### HasWormLogPath

`func (o *DfsRootfs) HasWormLogPath() bool`

HasWormLogPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


