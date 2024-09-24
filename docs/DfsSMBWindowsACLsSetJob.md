# DfsSMBWindowsACLsSetJob

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

## Methods

### NewDfsSMBWindowsACLsSetJob

`func NewDfsSMBWindowsACLsSetJob() *DfsSMBWindowsACLsSetJob`

NewDfsSMBWindowsACLsSetJob instantiates a new DfsSMBWindowsACLsSetJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBWindowsACLsSetJobWithDefaults

`func NewDfsSMBWindowsACLsSetJobWithDefaults() *DfsSMBWindowsACLsSetJob`

NewDfsSMBWindowsACLsSetJobWithDefaults instantiates a new DfsSMBWindowsACLsSetJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAceInfo

`func (o *DfsSMBWindowsACLsSetJob) GetAceInfo() map[string]interface{}`

GetAceInfo returns the AceInfo field if non-nil, zero value otherwise.

### GetAceInfoOk

`func (o *DfsSMBWindowsACLsSetJob) GetAceInfoOk() (*map[string]interface{}, bool)`

GetAceInfoOk returns a tuple with the AceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAceInfo

`func (o *DfsSMBWindowsACLsSetJob) SetAceInfo(v map[string]interface{})`

SetAceInfo sets AceInfo field to given value.

### HasAceInfo

`func (o *DfsSMBWindowsACLsSetJob) HasAceInfo() bool`

HasAceInfo returns a boolean if a field has been set.

### GetCluster

`func (o *DfsSMBWindowsACLsSetJob) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsSMBWindowsACLsSetJob) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsSMBWindowsACLsSetJob) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsSMBWindowsACLsSetJob) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsSMBWindowsACLsSetJob) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsSMBWindowsACLsSetJob) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsSMBWindowsACLsSetJob) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsSMBWindowsACLsSetJob) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCurrentAceNum

`func (o *DfsSMBWindowsACLsSetJob) GetCurrentAceNum() int64`

GetCurrentAceNum returns the CurrentAceNum field if non-nil, zero value otherwise.

### GetCurrentAceNumOk

`func (o *DfsSMBWindowsACLsSetJob) GetCurrentAceNumOk() (*int64, bool)`

GetCurrentAceNumOk returns a tuple with the CurrentAceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAceNum

`func (o *DfsSMBWindowsACLsSetJob) SetCurrentAceNum(v int64)`

SetCurrentAceNum sets CurrentAceNum field to given value.

### HasCurrentAceNum

`func (o *DfsSMBWindowsACLsSetJob) HasCurrentAceNum() bool`

HasCurrentAceNum returns a boolean if a field has been set.

### GetCurrentPath

`func (o *DfsSMBWindowsACLsSetJob) GetCurrentPath() string`

GetCurrentPath returns the CurrentPath field if non-nil, zero value otherwise.

### GetCurrentPathOk

`func (o *DfsSMBWindowsACLsSetJob) GetCurrentPathOk() (*string, bool)`

GetCurrentPathOk returns a tuple with the CurrentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPath

`func (o *DfsSMBWindowsACLsSetJob) SetCurrentPath(v string)`

SetCurrentPath sets CurrentPath field to given value.

### HasCurrentPath

`func (o *DfsSMBWindowsACLsSetJob) HasCurrentPath() bool`

HasCurrentPath returns a boolean if a field has been set.

### GetDfsGateway

`func (o *DfsSMBWindowsACLsSetJob) GetDfsGateway() DfsGatewayNestview`

GetDfsGateway returns the DfsGateway field if non-nil, zero value otherwise.

### GetDfsGatewayOk

`func (o *DfsSMBWindowsACLsSetJob) GetDfsGatewayOk() (*DfsGatewayNestview, bool)`

GetDfsGatewayOk returns a tuple with the DfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateway

`func (o *DfsSMBWindowsACLsSetJob) SetDfsGateway(v DfsGatewayNestview)`

SetDfsGateway sets DfsGateway field to given value.

### HasDfsGateway

`func (o *DfsSMBWindowsACLsSetJob) HasDfsGateway() bool`

HasDfsGateway returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsSMBWindowsACLsSetJob) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsSMBWindowsACLsSetJob) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsSMBWindowsACLsSetJob) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsSMBWindowsACLsSetJob) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetId

`func (o *DfsSMBWindowsACLsSetJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsSMBWindowsACLsSetJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsSMBWindowsACLsSetJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsSMBWindowsACLsSetJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceType

`func (o *DfsSMBWindowsACLsSetJob) GetInheritanceType() string`

GetInheritanceType returns the InheritanceType field if non-nil, zero value otherwise.

### GetInheritanceTypeOk

`func (o *DfsSMBWindowsACLsSetJob) GetInheritanceTypeOk() (*string, bool)`

GetInheritanceTypeOk returns a tuple with the InheritanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceType

`func (o *DfsSMBWindowsACLsSetJob) SetInheritanceType(v string)`

SetInheritanceType sets InheritanceType field to given value.

### HasInheritanceType

`func (o *DfsSMBWindowsACLsSetJob) HasInheritanceType() bool`

HasInheritanceType returns a boolean if a field has been set.

### GetInode

`func (o *DfsSMBWindowsACLsSetJob) GetInode() int64`

GetInode returns the Inode field if non-nil, zero value otherwise.

### GetInodeOk

`func (o *DfsSMBWindowsACLsSetJob) GetInodeOk() (*int64, bool)`

GetInodeOk returns a tuple with the Inode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInode

`func (o *DfsSMBWindowsACLsSetJob) SetInode(v int64)`

SetInode sets Inode field to given value.

### HasInode

`func (o *DfsSMBWindowsACLsSetJob) HasInode() bool`

HasInode returns a boolean if a field has been set.

### GetPath

`func (o *DfsSMBWindowsACLsSetJob) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSMBWindowsACLsSetJob) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSMBWindowsACLsSetJob) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsSMBWindowsACLsSetJob) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPid

`func (o *DfsSMBWindowsACLsSetJob) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *DfsSMBWindowsACLsSetJob) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *DfsSMBWindowsACLsSetJob) SetPid(v int64)`

SetPid sets Pid field to given value.

### HasPid

`func (o *DfsSMBWindowsACLsSetJob) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetRemainFiles

`func (o *DfsSMBWindowsACLsSetJob) GetRemainFiles() int64`

GetRemainFiles returns the RemainFiles field if non-nil, zero value otherwise.

### GetRemainFilesOk

`func (o *DfsSMBWindowsACLsSetJob) GetRemainFilesOk() (*int64, bool)`

GetRemainFilesOk returns a tuple with the RemainFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainFiles

`func (o *DfsSMBWindowsACLsSetJob) SetRemainFiles(v int64)`

SetRemainFiles sets RemainFiles field to given value.

### HasRemainFiles

`func (o *DfsSMBWindowsACLsSetJob) HasRemainFiles() bool`

HasRemainFiles returns a boolean if a field has been set.

### GetReplacePermission

`func (o *DfsSMBWindowsACLsSetJob) GetReplacePermission() bool`

GetReplacePermission returns the ReplacePermission field if non-nil, zero value otherwise.

### GetReplacePermissionOk

`func (o *DfsSMBWindowsACLsSetJob) GetReplacePermissionOk() (*bool, bool)`

GetReplacePermissionOk returns a tuple with the ReplacePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePermission

`func (o *DfsSMBWindowsACLsSetJob) SetReplacePermission(v bool)`

SetReplacePermission sets ReplacePermission field to given value.

### HasReplacePermission

`func (o *DfsSMBWindowsACLsSetJob) HasReplacePermission() bool`

HasReplacePermission returns a boolean if a field has been set.

### GetResume

`func (o *DfsSMBWindowsACLsSetJob) GetResume() bool`

GetResume returns the Resume field if non-nil, zero value otherwise.

### GetResumeOk

`func (o *DfsSMBWindowsACLsSetJob) GetResumeOk() (*bool, bool)`

GetResumeOk returns a tuple with the Resume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResume

`func (o *DfsSMBWindowsACLsSetJob) SetResume(v bool)`

SetResume sets Resume field to given value.

### HasResume

`func (o *DfsSMBWindowsACLsSetJob) HasResume() bool`

HasResume returns a boolean if a field has been set.

### GetStatus

`func (o *DfsSMBWindowsACLsSetJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsSMBWindowsACLsSetJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsSMBWindowsACLsSetJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsSMBWindowsACLsSetJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalFiles

`func (o *DfsSMBWindowsACLsSetJob) GetTotalFiles() int64`

GetTotalFiles returns the TotalFiles field if non-nil, zero value otherwise.

### GetTotalFilesOk

`func (o *DfsSMBWindowsACLsSetJob) GetTotalFilesOk() (*int64, bool)`

GetTotalFilesOk returns a tuple with the TotalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFiles

`func (o *DfsSMBWindowsACLsSetJob) SetTotalFiles(v int64)`

SetTotalFiles sets TotalFiles field to given value.

### HasTotalFiles

`func (o *DfsSMBWindowsACLsSetJob) HasTotalFiles() bool`

HasTotalFiles returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsSMBWindowsACLsSetJob) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsSMBWindowsACLsSetJob) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsSMBWindowsACLsSetJob) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsSMBWindowsACLsSetJob) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


