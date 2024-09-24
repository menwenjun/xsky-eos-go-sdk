# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPath** | Pointer to [**AccessPathNestview**](AccessPathNestview.md) |  | [optional] 
**ActionStatus** | Pointer to **string** |  | [optional] 
**AllocatedSize** | Pointer to **int64** |  | [optional] 
**AluaEnabled** | Pointer to **bool** |  | [optional] 
**BlockSnapshotNum** | Pointer to **int64** |  | [optional] 
**BlockVolumeGroup** | Pointer to [**VolumeGroupNestview**](VolumeGroupNestview.md) |  | [optional] 
**ClientGroupNum** | Pointer to **int64** |  | [optional] 
**CloudPlatform** | Pointer to [**CloudPlatformNestview**](CloudPlatformNestview.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**CrcCheck** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DpBlockBackupPolicyNum** | Pointer to **int64** |  | [optional] 
**DpBlockReplicationPolicy** | Pointer to [**DpBlockReplicationPolicyNestview**](DpBlockReplicationPolicyNestview.md) |  | [optional] 
**DpBlockSnapshotPolicy** | Pointer to [**DpBlockSnapshotPolicyNestview**](DpBlockSnapshotPolicyNestview.md) |  | [optional] 
**Flattened** | Pointer to **bool** |  | [optional] 
**Format** | Pointer to **int64** | 128 or 129 (advanced) | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**LatestSnapshotTime** | Pointer to **time.Time** |  | [optional] 
**MigrationStripes** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OriginalName** | Pointer to **string** |  | [optional] 
**OriginalVolumeName** | Pointer to **string** |  | [optional] 
**Passive** | Pointer to **bool** |  | [optional] 
**PerformancePriority** | Pointer to **int64** | 0 or 1 | [optional] 
**Pool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Qos** | Pointer to [**VolumeQosSpec**](VolumeQosSpec.md) |  | [optional] 
**QosEnabled** | Pointer to **bool** |  | [optional] 
**RbdClient** | Pointer to [**RBDClient**](RBDClient.md) |  | [optional] 
**Recycled** | Pointer to **bool** | trash | [optional] 
**RemoteCluster** | Pointer to [**RemoteClusterNestview**](RemoteClusterNestview.md) |  | [optional] 
**ReplicationDest** | Pointer to **string** |  | [optional] 
**ReplicationPool** | Pointer to **string** |  | [optional] 
**ReplicationPoolId** | Pointer to **int64** |  | [optional] 
**ReplicationPoolName** | Pointer to **string** |  | [optional] 
**ReplicationRole** | Pointer to **string** |  | [optional] 
**ReplicationStatus** | Pointer to **string** |  | [optional] 
**ReplicationVersion** | Pointer to **int64** |  | [optional] 
**ReplicationVolume** | Pointer to **string** |  | [optional] 
**ReplicationVolumeId** | Pointer to **int64** |  | [optional] 
**ReplicationVolumeName** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Sn** | Pointer to **string** |  | [optional] 
**Snapshot** | Pointer to [**SnapshotNestview**](SnapshotNestview.md) |  | [optional] 
**SnapshotReplicationRole** | Pointer to **string** |  | [optional] 
**SnapshotReplicationStatus** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**SourceUuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TrashResource** | Pointer to [**TrashResourceNestview**](TrashResourceNestview.md) |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**VolumeName** | Pointer to **string** |  | [optional] 

## Methods

### NewVolume

`func NewVolume() *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPath

`func (o *Volume) GetAccessPath() AccessPathNestview`

GetAccessPath returns the AccessPath field if non-nil, zero value otherwise.

### GetAccessPathOk

`func (o *Volume) GetAccessPathOk() (*AccessPathNestview, bool)`

GetAccessPathOk returns a tuple with the AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPath

`func (o *Volume) SetAccessPath(v AccessPathNestview)`

SetAccessPath sets AccessPath field to given value.

### HasAccessPath

`func (o *Volume) HasAccessPath() bool`

HasAccessPath returns a boolean if a field has been set.

### GetActionStatus

`func (o *Volume) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *Volume) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *Volume) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *Volume) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *Volume) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *Volume) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *Volume) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *Volume) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetAluaEnabled

`func (o *Volume) GetAluaEnabled() bool`

GetAluaEnabled returns the AluaEnabled field if non-nil, zero value otherwise.

### GetAluaEnabledOk

`func (o *Volume) GetAluaEnabledOk() (*bool, bool)`

GetAluaEnabledOk returns a tuple with the AluaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAluaEnabled

`func (o *Volume) SetAluaEnabled(v bool)`

SetAluaEnabled sets AluaEnabled field to given value.

### HasAluaEnabled

`func (o *Volume) HasAluaEnabled() bool`

HasAluaEnabled returns a boolean if a field has been set.

### GetBlockSnapshotNum

`func (o *Volume) GetBlockSnapshotNum() int64`

GetBlockSnapshotNum returns the BlockSnapshotNum field if non-nil, zero value otherwise.

### GetBlockSnapshotNumOk

`func (o *Volume) GetBlockSnapshotNumOk() (*int64, bool)`

GetBlockSnapshotNumOk returns a tuple with the BlockSnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSnapshotNum

`func (o *Volume) SetBlockSnapshotNum(v int64)`

SetBlockSnapshotNum sets BlockSnapshotNum field to given value.

### HasBlockSnapshotNum

`func (o *Volume) HasBlockSnapshotNum() bool`

HasBlockSnapshotNum returns a boolean if a field has been set.

### GetBlockVolumeGroup

`func (o *Volume) GetBlockVolumeGroup() VolumeGroupNestview`

GetBlockVolumeGroup returns the BlockVolumeGroup field if non-nil, zero value otherwise.

### GetBlockVolumeGroupOk

`func (o *Volume) GetBlockVolumeGroupOk() (*VolumeGroupNestview, bool)`

GetBlockVolumeGroupOk returns a tuple with the BlockVolumeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeGroup

`func (o *Volume) SetBlockVolumeGroup(v VolumeGroupNestview)`

SetBlockVolumeGroup sets BlockVolumeGroup field to given value.

### HasBlockVolumeGroup

`func (o *Volume) HasBlockVolumeGroup() bool`

HasBlockVolumeGroup returns a boolean if a field has been set.

### GetClientGroupNum

`func (o *Volume) GetClientGroupNum() int64`

GetClientGroupNum returns the ClientGroupNum field if non-nil, zero value otherwise.

### GetClientGroupNumOk

`func (o *Volume) GetClientGroupNumOk() (*int64, bool)`

GetClientGroupNumOk returns a tuple with the ClientGroupNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGroupNum

`func (o *Volume) SetClientGroupNum(v int64)`

SetClientGroupNum sets ClientGroupNum field to given value.

### HasClientGroupNum

`func (o *Volume) HasClientGroupNum() bool`

HasClientGroupNum returns a boolean if a field has been set.

### GetCloudPlatform

`func (o *Volume) GetCloudPlatform() CloudPlatformNestview`

GetCloudPlatform returns the CloudPlatform field if non-nil, zero value otherwise.

### GetCloudPlatformOk

`func (o *Volume) GetCloudPlatformOk() (*CloudPlatformNestview, bool)`

GetCloudPlatformOk returns a tuple with the CloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatform

`func (o *Volume) SetCloudPlatform(v CloudPlatformNestview)`

SetCloudPlatform sets CloudPlatform field to given value.

### HasCloudPlatform

`func (o *Volume) HasCloudPlatform() bool`

HasCloudPlatform returns a boolean if a field has been set.

### GetCluster

`func (o *Volume) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Volume) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Volume) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Volume) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCrcCheck

`func (o *Volume) GetCrcCheck() bool`

GetCrcCheck returns the CrcCheck field if non-nil, zero value otherwise.

### GetCrcCheckOk

`func (o *Volume) GetCrcCheckOk() (*bool, bool)`

GetCrcCheckOk returns a tuple with the CrcCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrcCheck

`func (o *Volume) SetCrcCheck(v bool)`

SetCrcCheck sets CrcCheck field to given value.

### HasCrcCheck

`func (o *Volume) HasCrcCheck() bool`

HasCrcCheck returns a boolean if a field has been set.

### GetCreate

`func (o *Volume) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Volume) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Volume) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Volume) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *Volume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Volume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Volume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Volume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDpBlockBackupPolicyNum

`func (o *Volume) GetDpBlockBackupPolicyNum() int64`

GetDpBlockBackupPolicyNum returns the DpBlockBackupPolicyNum field if non-nil, zero value otherwise.

### GetDpBlockBackupPolicyNumOk

`func (o *Volume) GetDpBlockBackupPolicyNumOk() (*int64, bool)`

GetDpBlockBackupPolicyNumOk returns a tuple with the DpBlockBackupPolicyNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpBlockBackupPolicyNum

`func (o *Volume) SetDpBlockBackupPolicyNum(v int64)`

SetDpBlockBackupPolicyNum sets DpBlockBackupPolicyNum field to given value.

### HasDpBlockBackupPolicyNum

`func (o *Volume) HasDpBlockBackupPolicyNum() bool`

HasDpBlockBackupPolicyNum returns a boolean if a field has been set.

### GetDpBlockReplicationPolicy

`func (o *Volume) GetDpBlockReplicationPolicy() DpBlockReplicationPolicyNestview`

GetDpBlockReplicationPolicy returns the DpBlockReplicationPolicy field if non-nil, zero value otherwise.

### GetDpBlockReplicationPolicyOk

`func (o *Volume) GetDpBlockReplicationPolicyOk() (*DpBlockReplicationPolicyNestview, bool)`

GetDpBlockReplicationPolicyOk returns a tuple with the DpBlockReplicationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpBlockReplicationPolicy

`func (o *Volume) SetDpBlockReplicationPolicy(v DpBlockReplicationPolicyNestview)`

SetDpBlockReplicationPolicy sets DpBlockReplicationPolicy field to given value.

### HasDpBlockReplicationPolicy

`func (o *Volume) HasDpBlockReplicationPolicy() bool`

HasDpBlockReplicationPolicy returns a boolean if a field has been set.

### GetDpBlockSnapshotPolicy

`func (o *Volume) GetDpBlockSnapshotPolicy() DpBlockSnapshotPolicyNestview`

GetDpBlockSnapshotPolicy returns the DpBlockSnapshotPolicy field if non-nil, zero value otherwise.

### GetDpBlockSnapshotPolicyOk

`func (o *Volume) GetDpBlockSnapshotPolicyOk() (*DpBlockSnapshotPolicyNestview, bool)`

GetDpBlockSnapshotPolicyOk returns a tuple with the DpBlockSnapshotPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpBlockSnapshotPolicy

`func (o *Volume) SetDpBlockSnapshotPolicy(v DpBlockSnapshotPolicyNestview)`

SetDpBlockSnapshotPolicy sets DpBlockSnapshotPolicy field to given value.

### HasDpBlockSnapshotPolicy

`func (o *Volume) HasDpBlockSnapshotPolicy() bool`

HasDpBlockSnapshotPolicy returns a boolean if a field has been set.

### GetFlattened

`func (o *Volume) GetFlattened() bool`

GetFlattened returns the Flattened field if non-nil, zero value otherwise.

### GetFlattenedOk

`func (o *Volume) GetFlattenedOk() (*bool, bool)`

GetFlattenedOk returns a tuple with the Flattened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattened

`func (o *Volume) SetFlattened(v bool)`

SetFlattened sets Flattened field to given value.

### HasFlattened

`func (o *Volume) HasFlattened() bool`

HasFlattened returns a boolean if a field has been set.

### GetFormat

`func (o *Volume) GetFormat() int64`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Volume) GetFormatOk() (*int64, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Volume) SetFormat(v int64)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Volume) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetHidden

`func (o *Volume) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Volume) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Volume) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Volume) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *Volume) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Volume) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Volume) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Volume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageId

`func (o *Volume) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *Volume) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *Volume) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *Volume) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetLatestSnapshotTime

`func (o *Volume) GetLatestSnapshotTime() time.Time`

GetLatestSnapshotTime returns the LatestSnapshotTime field if non-nil, zero value otherwise.

### GetLatestSnapshotTimeOk

`func (o *Volume) GetLatestSnapshotTimeOk() (*time.Time, bool)`

GetLatestSnapshotTimeOk returns a tuple with the LatestSnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSnapshotTime

`func (o *Volume) SetLatestSnapshotTime(v time.Time)`

SetLatestSnapshotTime sets LatestSnapshotTime field to given value.

### HasLatestSnapshotTime

`func (o *Volume) HasLatestSnapshotTime() bool`

HasLatestSnapshotTime returns a boolean if a field has been set.

### GetMigrationStripes

`func (o *Volume) GetMigrationStripes() int64`

GetMigrationStripes returns the MigrationStripes field if non-nil, zero value otherwise.

### GetMigrationStripesOk

`func (o *Volume) GetMigrationStripesOk() (*int64, bool)`

GetMigrationStripesOk returns a tuple with the MigrationStripes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationStripes

`func (o *Volume) SetMigrationStripes(v int64)`

SetMigrationStripes sets MigrationStripes field to given value.

### HasMigrationStripes

`func (o *Volume) HasMigrationStripes() bool`

HasMigrationStripes returns a boolean if a field has been set.

### GetName

`func (o *Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Volume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Volume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalName

`func (o *Volume) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *Volume) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *Volume) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *Volume) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetOriginalVolumeName

`func (o *Volume) GetOriginalVolumeName() string`

GetOriginalVolumeName returns the OriginalVolumeName field if non-nil, zero value otherwise.

### GetOriginalVolumeNameOk

`func (o *Volume) GetOriginalVolumeNameOk() (*string, bool)`

GetOriginalVolumeNameOk returns a tuple with the OriginalVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalVolumeName

`func (o *Volume) SetOriginalVolumeName(v string)`

SetOriginalVolumeName sets OriginalVolumeName field to given value.

### HasOriginalVolumeName

`func (o *Volume) HasOriginalVolumeName() bool`

HasOriginalVolumeName returns a boolean if a field has been set.

### GetPassive

`func (o *Volume) GetPassive() bool`

GetPassive returns the Passive field if non-nil, zero value otherwise.

### GetPassiveOk

`func (o *Volume) GetPassiveOk() (*bool, bool)`

GetPassiveOk returns a tuple with the Passive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassive

`func (o *Volume) SetPassive(v bool)`

SetPassive sets Passive field to given value.

### HasPassive

`func (o *Volume) HasPassive() bool`

HasPassive returns a boolean if a field has been set.

### GetPerformancePriority

`func (o *Volume) GetPerformancePriority() int64`

GetPerformancePriority returns the PerformancePriority field if non-nil, zero value otherwise.

### GetPerformancePriorityOk

`func (o *Volume) GetPerformancePriorityOk() (*int64, bool)`

GetPerformancePriorityOk returns a tuple with the PerformancePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformancePriority

`func (o *Volume) SetPerformancePriority(v int64)`

SetPerformancePriority sets PerformancePriority field to given value.

### HasPerformancePriority

`func (o *Volume) HasPerformancePriority() bool`

HasPerformancePriority returns a boolean if a field has been set.

### GetPool

`func (o *Volume) GetPool() PoolNestview`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *Volume) GetPoolOk() (*PoolNestview, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *Volume) SetPool(v PoolNestview)`

SetPool sets Pool field to given value.

### HasPool

`func (o *Volume) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetProgress

`func (o *Volume) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Volume) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Volume) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Volume) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetQos

`func (o *Volume) GetQos() VolumeQosSpec`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Volume) GetQosOk() (*VolumeQosSpec, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Volume) SetQos(v VolumeQosSpec)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Volume) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetQosEnabled

`func (o *Volume) GetQosEnabled() bool`

GetQosEnabled returns the QosEnabled field if non-nil, zero value otherwise.

### GetQosEnabledOk

`func (o *Volume) GetQosEnabledOk() (*bool, bool)`

GetQosEnabledOk returns a tuple with the QosEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosEnabled

`func (o *Volume) SetQosEnabled(v bool)`

SetQosEnabled sets QosEnabled field to given value.

### HasQosEnabled

`func (o *Volume) HasQosEnabled() bool`

HasQosEnabled returns a boolean if a field has been set.

### GetRbdClient

`func (o *Volume) GetRbdClient() RBDClient`

GetRbdClient returns the RbdClient field if non-nil, zero value otherwise.

### GetRbdClientOk

`func (o *Volume) GetRbdClientOk() (*RBDClient, bool)`

GetRbdClientOk returns a tuple with the RbdClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbdClient

`func (o *Volume) SetRbdClient(v RBDClient)`

SetRbdClient sets RbdClient field to given value.

### HasRbdClient

`func (o *Volume) HasRbdClient() bool`

HasRbdClient returns a boolean if a field has been set.

### GetRecycled

`func (o *Volume) GetRecycled() bool`

GetRecycled returns the Recycled field if non-nil, zero value otherwise.

### GetRecycledOk

`func (o *Volume) GetRecycledOk() (*bool, bool)`

GetRecycledOk returns a tuple with the Recycled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycled

`func (o *Volume) SetRecycled(v bool)`

SetRecycled sets Recycled field to given value.

### HasRecycled

`func (o *Volume) HasRecycled() bool`

HasRecycled returns a boolean if a field has been set.

### GetRemoteCluster

`func (o *Volume) GetRemoteCluster() RemoteClusterNestview`

GetRemoteCluster returns the RemoteCluster field if non-nil, zero value otherwise.

### GetRemoteClusterOk

`func (o *Volume) GetRemoteClusterOk() (*RemoteClusterNestview, bool)`

GetRemoteClusterOk returns a tuple with the RemoteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCluster

`func (o *Volume) SetRemoteCluster(v RemoteClusterNestview)`

SetRemoteCluster sets RemoteCluster field to given value.

### HasRemoteCluster

`func (o *Volume) HasRemoteCluster() bool`

HasRemoteCluster returns a boolean if a field has been set.

### GetReplicationDest

`func (o *Volume) GetReplicationDest() string`

GetReplicationDest returns the ReplicationDest field if non-nil, zero value otherwise.

### GetReplicationDestOk

`func (o *Volume) GetReplicationDestOk() (*string, bool)`

GetReplicationDestOk returns a tuple with the ReplicationDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationDest

`func (o *Volume) SetReplicationDest(v string)`

SetReplicationDest sets ReplicationDest field to given value.

### HasReplicationDest

`func (o *Volume) HasReplicationDest() bool`

HasReplicationDest returns a boolean if a field has been set.

### GetReplicationPool

`func (o *Volume) GetReplicationPool() string`

GetReplicationPool returns the ReplicationPool field if non-nil, zero value otherwise.

### GetReplicationPoolOk

`func (o *Volume) GetReplicationPoolOk() (*string, bool)`

GetReplicationPoolOk returns a tuple with the ReplicationPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPool

`func (o *Volume) SetReplicationPool(v string)`

SetReplicationPool sets ReplicationPool field to given value.

### HasReplicationPool

`func (o *Volume) HasReplicationPool() bool`

HasReplicationPool returns a boolean if a field has been set.

### GetReplicationPoolId

`func (o *Volume) GetReplicationPoolId() int64`

GetReplicationPoolId returns the ReplicationPoolId field if non-nil, zero value otherwise.

### GetReplicationPoolIdOk

`func (o *Volume) GetReplicationPoolIdOk() (*int64, bool)`

GetReplicationPoolIdOk returns a tuple with the ReplicationPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPoolId

`func (o *Volume) SetReplicationPoolId(v int64)`

SetReplicationPoolId sets ReplicationPoolId field to given value.

### HasReplicationPoolId

`func (o *Volume) HasReplicationPoolId() bool`

HasReplicationPoolId returns a boolean if a field has been set.

### GetReplicationPoolName

`func (o *Volume) GetReplicationPoolName() string`

GetReplicationPoolName returns the ReplicationPoolName field if non-nil, zero value otherwise.

### GetReplicationPoolNameOk

`func (o *Volume) GetReplicationPoolNameOk() (*string, bool)`

GetReplicationPoolNameOk returns a tuple with the ReplicationPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPoolName

`func (o *Volume) SetReplicationPoolName(v string)`

SetReplicationPoolName sets ReplicationPoolName field to given value.

### HasReplicationPoolName

`func (o *Volume) HasReplicationPoolName() bool`

HasReplicationPoolName returns a boolean if a field has been set.

### GetReplicationRole

`func (o *Volume) GetReplicationRole() string`

GetReplicationRole returns the ReplicationRole field if non-nil, zero value otherwise.

### GetReplicationRoleOk

`func (o *Volume) GetReplicationRoleOk() (*string, bool)`

GetReplicationRoleOk returns a tuple with the ReplicationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRole

`func (o *Volume) SetReplicationRole(v string)`

SetReplicationRole sets ReplicationRole field to given value.

### HasReplicationRole

`func (o *Volume) HasReplicationRole() bool`

HasReplicationRole returns a boolean if a field has been set.

### GetReplicationStatus

`func (o *Volume) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *Volume) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *Volume) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *Volume) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetReplicationVersion

`func (o *Volume) GetReplicationVersion() int64`

GetReplicationVersion returns the ReplicationVersion field if non-nil, zero value otherwise.

### GetReplicationVersionOk

`func (o *Volume) GetReplicationVersionOk() (*int64, bool)`

GetReplicationVersionOk returns a tuple with the ReplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVersion

`func (o *Volume) SetReplicationVersion(v int64)`

SetReplicationVersion sets ReplicationVersion field to given value.

### HasReplicationVersion

`func (o *Volume) HasReplicationVersion() bool`

HasReplicationVersion returns a boolean if a field has been set.

### GetReplicationVolume

`func (o *Volume) GetReplicationVolume() string`

GetReplicationVolume returns the ReplicationVolume field if non-nil, zero value otherwise.

### GetReplicationVolumeOk

`func (o *Volume) GetReplicationVolumeOk() (*string, bool)`

GetReplicationVolumeOk returns a tuple with the ReplicationVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolume

`func (o *Volume) SetReplicationVolume(v string)`

SetReplicationVolume sets ReplicationVolume field to given value.

### HasReplicationVolume

`func (o *Volume) HasReplicationVolume() bool`

HasReplicationVolume returns a boolean if a field has been set.

### GetReplicationVolumeId

`func (o *Volume) GetReplicationVolumeId() int64`

GetReplicationVolumeId returns the ReplicationVolumeId field if non-nil, zero value otherwise.

### GetReplicationVolumeIdOk

`func (o *Volume) GetReplicationVolumeIdOk() (*int64, bool)`

GetReplicationVolumeIdOk returns a tuple with the ReplicationVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolumeId

`func (o *Volume) SetReplicationVolumeId(v int64)`

SetReplicationVolumeId sets ReplicationVolumeId field to given value.

### HasReplicationVolumeId

`func (o *Volume) HasReplicationVolumeId() bool`

HasReplicationVolumeId returns a boolean if a field has been set.

### GetReplicationVolumeName

`func (o *Volume) GetReplicationVolumeName() string`

GetReplicationVolumeName returns the ReplicationVolumeName field if non-nil, zero value otherwise.

### GetReplicationVolumeNameOk

`func (o *Volume) GetReplicationVolumeNameOk() (*string, bool)`

GetReplicationVolumeNameOk returns a tuple with the ReplicationVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolumeName

`func (o *Volume) SetReplicationVolumeName(v string)`

SetReplicationVolumeName sets ReplicationVolumeName field to given value.

### HasReplicationVolumeName

`func (o *Volume) HasReplicationVolumeName() bool`

HasReplicationVolumeName returns a boolean if a field has been set.

### GetSize

`func (o *Volume) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Volume) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Volume) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Volume) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSn

`func (o *Volume) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *Volume) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *Volume) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *Volume) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetSnapshot

`func (o *Volume) GetSnapshot() SnapshotNestview`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *Volume) GetSnapshotOk() (*SnapshotNestview, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *Volume) SetSnapshot(v SnapshotNestview)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *Volume) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetSnapshotReplicationRole

`func (o *Volume) GetSnapshotReplicationRole() string`

GetSnapshotReplicationRole returns the SnapshotReplicationRole field if non-nil, zero value otherwise.

### GetSnapshotReplicationRoleOk

`func (o *Volume) GetSnapshotReplicationRoleOk() (*string, bool)`

GetSnapshotReplicationRoleOk returns a tuple with the SnapshotReplicationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotReplicationRole

`func (o *Volume) SetSnapshotReplicationRole(v string)`

SetSnapshotReplicationRole sets SnapshotReplicationRole field to given value.

### HasSnapshotReplicationRole

`func (o *Volume) HasSnapshotReplicationRole() bool`

HasSnapshotReplicationRole returns a boolean if a field has been set.

### GetSnapshotReplicationStatus

`func (o *Volume) GetSnapshotReplicationStatus() string`

GetSnapshotReplicationStatus returns the SnapshotReplicationStatus field if non-nil, zero value otherwise.

### GetSnapshotReplicationStatusOk

`func (o *Volume) GetSnapshotReplicationStatusOk() (*string, bool)`

GetSnapshotReplicationStatusOk returns a tuple with the SnapshotReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotReplicationStatus

`func (o *Volume) SetSnapshotReplicationStatus(v string)`

SetSnapshotReplicationStatus sets SnapshotReplicationStatus field to given value.

### HasSnapshotReplicationStatus

`func (o *Volume) HasSnapshotReplicationStatus() bool`

HasSnapshotReplicationStatus returns a boolean if a field has been set.

### GetSourceType

`func (o *Volume) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Volume) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Volume) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *Volume) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSourceUuid

`func (o *Volume) GetSourceUuid() string`

GetSourceUuid returns the SourceUuid field if non-nil, zero value otherwise.

### GetSourceUuidOk

`func (o *Volume) GetSourceUuidOk() (*string, bool)`

GetSourceUuidOk returns a tuple with the SourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUuid

`func (o *Volume) SetSourceUuid(v string)`

SetSourceUuid sets SourceUuid field to given value.

### HasSourceUuid

`func (o *Volume) HasSourceUuid() bool`

HasSourceUuid returns a boolean if a field has been set.

### GetStatus

`func (o *Volume) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Volume) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Volume) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Volume) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrashResource

`func (o *Volume) GetTrashResource() TrashResourceNestview`

GetTrashResource returns the TrashResource field if non-nil, zero value otherwise.

### GetTrashResourceOk

`func (o *Volume) GetTrashResourceOk() (*TrashResourceNestview, bool)`

GetTrashResourceOk returns a tuple with the TrashResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashResource

`func (o *Volume) SetTrashResource(v TrashResourceNestview)`

SetTrashResource sets TrashResource field to given value.

### HasTrashResource

`func (o *Volume) HasTrashResource() bool`

HasTrashResource returns a boolean if a field has been set.

### GetUid

`func (o *Volume) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Volume) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Volume) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Volume) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUpdate

`func (o *Volume) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Volume) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Volume) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Volume) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVolumeName

`func (o *Volume) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *Volume) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *Volume) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *Volume) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


