# DfsSMBWindowsACLsSetJobRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AceInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**CurrentAceNum** | Pointer to **int64** |  | [optional] 
**CurrentPath** | Pointer to **string** |  | [optional] 
**DfsGateway** | Pointer to [**DfsGatewayNestview**](DfsGatewayNestview.md) |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InheritanceType** | Pointer to **string** |  | [optional] 
**Inode** | Pointer to **int64** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **int64** |  | [optional] 
**RemainFiles** | Pointer to **int64** |  | [optional] 
**ReplacePermission** | Pointer to **bool** |  | [optional] 
**Resume** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalFiles** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Aces** | Pointer to [**[]DfsSMBWindowsACE**](DfsSMBWindowsACE.md) |  | [optional] 

## Methods

### NewDfsSMBWindowsACLsSetJobRecord

`func NewDfsSMBWindowsACLsSetJobRecord() *DfsSMBWindowsACLsSetJobRecord`

NewDfsSMBWindowsACLsSetJobRecord instantiates a new DfsSMBWindowsACLsSetJobRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBWindowsACLsSetJobRecordWithDefaults

`func NewDfsSMBWindowsACLsSetJobRecordWithDefaults() *DfsSMBWindowsACLsSetJobRecord`

NewDfsSMBWindowsACLsSetJobRecordWithDefaults instantiates a new DfsSMBWindowsACLsSetJobRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAceInfo

`func (o *DfsSMBWindowsACLsSetJobRecord) GetAceInfo() map[string]interface{}`

GetAceInfo returns the AceInfo field if non-nil, zero value otherwise.

### GetAceInfoOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetAceInfoOk() (*map[string]interface{}, bool)`

GetAceInfoOk returns a tuple with the AceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAceInfo

`func (o *DfsSMBWindowsACLsSetJobRecord) SetAceInfo(v map[string]interface{})`

SetAceInfo sets AceInfo field to given value.

### HasAceInfo

`func (o *DfsSMBWindowsACLsSetJobRecord) HasAceInfo() bool`

HasAceInfo returns a boolean if a field has been set.

### GetCluster

`func (o *DfsSMBWindowsACLsSetJobRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsSMBWindowsACLsSetJobRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsSMBWindowsACLsSetJobRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsSMBWindowsACLsSetJobRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsSMBWindowsACLsSetJobRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsSMBWindowsACLsSetJobRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCurrentAceNum

`func (o *DfsSMBWindowsACLsSetJobRecord) GetCurrentAceNum() int64`

GetCurrentAceNum returns the CurrentAceNum field if non-nil, zero value otherwise.

### GetCurrentAceNumOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetCurrentAceNumOk() (*int64, bool)`

GetCurrentAceNumOk returns a tuple with the CurrentAceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAceNum

`func (o *DfsSMBWindowsACLsSetJobRecord) SetCurrentAceNum(v int64)`

SetCurrentAceNum sets CurrentAceNum field to given value.

### HasCurrentAceNum

`func (o *DfsSMBWindowsACLsSetJobRecord) HasCurrentAceNum() bool`

HasCurrentAceNum returns a boolean if a field has been set.

### GetCurrentPath

`func (o *DfsSMBWindowsACLsSetJobRecord) GetCurrentPath() string`

GetCurrentPath returns the CurrentPath field if non-nil, zero value otherwise.

### GetCurrentPathOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetCurrentPathOk() (*string, bool)`

GetCurrentPathOk returns a tuple with the CurrentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPath

`func (o *DfsSMBWindowsACLsSetJobRecord) SetCurrentPath(v string)`

SetCurrentPath sets CurrentPath field to given value.

### HasCurrentPath

`func (o *DfsSMBWindowsACLsSetJobRecord) HasCurrentPath() bool`

HasCurrentPath returns a boolean if a field has been set.

### GetDfsGateway

`func (o *DfsSMBWindowsACLsSetJobRecord) GetDfsGateway() DfsGatewayNestview`

GetDfsGateway returns the DfsGateway field if non-nil, zero value otherwise.

### GetDfsGatewayOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetDfsGatewayOk() (*DfsGatewayNestview, bool)`

GetDfsGatewayOk returns a tuple with the DfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateway

`func (o *DfsSMBWindowsACLsSetJobRecord) SetDfsGateway(v DfsGatewayNestview)`

SetDfsGateway sets DfsGateway field to given value.

### HasDfsGateway

`func (o *DfsSMBWindowsACLsSetJobRecord) HasDfsGateway() bool`

HasDfsGateway returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsSMBWindowsACLsSetJobRecord) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsSMBWindowsACLsSetJobRecord) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsSMBWindowsACLsSetJobRecord) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetId

`func (o *DfsSMBWindowsACLsSetJobRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsSMBWindowsACLsSetJobRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsSMBWindowsACLsSetJobRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceType

`func (o *DfsSMBWindowsACLsSetJobRecord) GetInheritanceType() string`

GetInheritanceType returns the InheritanceType field if non-nil, zero value otherwise.

### GetInheritanceTypeOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetInheritanceTypeOk() (*string, bool)`

GetInheritanceTypeOk returns a tuple with the InheritanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceType

`func (o *DfsSMBWindowsACLsSetJobRecord) SetInheritanceType(v string)`

SetInheritanceType sets InheritanceType field to given value.

### HasInheritanceType

`func (o *DfsSMBWindowsACLsSetJobRecord) HasInheritanceType() bool`

HasInheritanceType returns a boolean if a field has been set.

### GetInode

`func (o *DfsSMBWindowsACLsSetJobRecord) GetInode() int64`

GetInode returns the Inode field if non-nil, zero value otherwise.

### GetInodeOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetInodeOk() (*int64, bool)`

GetInodeOk returns a tuple with the Inode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInode

`func (o *DfsSMBWindowsACLsSetJobRecord) SetInode(v int64)`

SetInode sets Inode field to given value.

### HasInode

`func (o *DfsSMBWindowsACLsSetJobRecord) HasInode() bool`

HasInode returns a boolean if a field has been set.

### GetPath

`func (o *DfsSMBWindowsACLsSetJobRecord) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSMBWindowsACLsSetJobRecord) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsSMBWindowsACLsSetJobRecord) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPid

`func (o *DfsSMBWindowsACLsSetJobRecord) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *DfsSMBWindowsACLsSetJobRecord) SetPid(v int64)`

SetPid sets Pid field to given value.

### HasPid

`func (o *DfsSMBWindowsACLsSetJobRecord) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetRemainFiles

`func (o *DfsSMBWindowsACLsSetJobRecord) GetRemainFiles() int64`

GetRemainFiles returns the RemainFiles field if non-nil, zero value otherwise.

### GetRemainFilesOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetRemainFilesOk() (*int64, bool)`

GetRemainFilesOk returns a tuple with the RemainFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainFiles

`func (o *DfsSMBWindowsACLsSetJobRecord) SetRemainFiles(v int64)`

SetRemainFiles sets RemainFiles field to given value.

### HasRemainFiles

`func (o *DfsSMBWindowsACLsSetJobRecord) HasRemainFiles() bool`

HasRemainFiles returns a boolean if a field has been set.

### GetReplacePermission

`func (o *DfsSMBWindowsACLsSetJobRecord) GetReplacePermission() bool`

GetReplacePermission returns the ReplacePermission field if non-nil, zero value otherwise.

### GetReplacePermissionOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetReplacePermissionOk() (*bool, bool)`

GetReplacePermissionOk returns a tuple with the ReplacePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePermission

`func (o *DfsSMBWindowsACLsSetJobRecord) SetReplacePermission(v bool)`

SetReplacePermission sets ReplacePermission field to given value.

### HasReplacePermission

`func (o *DfsSMBWindowsACLsSetJobRecord) HasReplacePermission() bool`

HasReplacePermission returns a boolean if a field has been set.

### GetResume

`func (o *DfsSMBWindowsACLsSetJobRecord) GetResume() bool`

GetResume returns the Resume field if non-nil, zero value otherwise.

### GetResumeOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetResumeOk() (*bool, bool)`

GetResumeOk returns a tuple with the Resume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResume

`func (o *DfsSMBWindowsACLsSetJobRecord) SetResume(v bool)`

SetResume sets Resume field to given value.

### HasResume

`func (o *DfsSMBWindowsACLsSetJobRecord) HasResume() bool`

HasResume returns a boolean if a field has been set.

### GetStatus

`func (o *DfsSMBWindowsACLsSetJobRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsSMBWindowsACLsSetJobRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsSMBWindowsACLsSetJobRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalFiles

`func (o *DfsSMBWindowsACLsSetJobRecord) GetTotalFiles() int64`

GetTotalFiles returns the TotalFiles field if non-nil, zero value otherwise.

### GetTotalFilesOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetTotalFilesOk() (*int64, bool)`

GetTotalFilesOk returns a tuple with the TotalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFiles

`func (o *DfsSMBWindowsACLsSetJobRecord) SetTotalFiles(v int64)`

SetTotalFiles sets TotalFiles field to given value.

### HasTotalFiles

`func (o *DfsSMBWindowsACLsSetJobRecord) HasTotalFiles() bool`

HasTotalFiles returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsSMBWindowsACLsSetJobRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsSMBWindowsACLsSetJobRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsSMBWindowsACLsSetJobRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetAces

`func (o *DfsSMBWindowsACLsSetJobRecord) GetAces() []DfsSMBWindowsACE`

GetAces returns the Aces field if non-nil, zero value otherwise.

### GetAcesOk

`func (o *DfsSMBWindowsACLsSetJobRecord) GetAcesOk() (*[]DfsSMBWindowsACE, bool)`

GetAcesOk returns a tuple with the Aces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAces

`func (o *DfsSMBWindowsACLsSetJobRecord) SetAces(v []DfsSMBWindowsACE)`

SetAces sets Aces field to given value.

### HasAces

`func (o *DfsSMBWindowsACLsSetJobRecord) HasAces() bool`

HasAces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


