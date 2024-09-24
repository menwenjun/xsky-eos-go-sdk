# DfsQuotaRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsPath** | Pointer to [**DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**DomainUser** | Pointer to [**DomainUserNestview**](DomainUserNestview.md) |  | [optional] 
**DomainUserGroup** | Pointer to [**DomainUserGroupNestview**](DomainUserGroupNestview.md) |  | [optional] 
**FilesExceedTime** | Pointer to **time.Time** |  | [optional] 
**FilesGraceTime** | Pointer to **int64** |  | [optional] 
**FilesHardQuota** | Pointer to **int64** |  | [optional] 
**FilesRatio** | Pointer to **float32** |  | [optional] 
**FilesSoftQuota** | Pointer to **int64** |  | [optional] 
**FilesSuggestQuota** | Pointer to **int64** |  | [optional] 
**FsUser** | Pointer to [**FSUserNestview**](FSUserNestview.md) |  | [optional] 
**FsUserGroup** | Pointer to [**FSUserGroupNestview**](FSUserGroupNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Inode** | Pointer to **int64** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**SizeExceedTime** | Pointer to **time.Time** |  | [optional] 
**SizeGraceTime** | Pointer to **int64** |  | [optional] 
**SizeHardQuota** | Pointer to **int64** |  | [optional] 
**SizeRatio** | Pointer to **float32** |  | [optional] 
**SizeSoftQuota** | Pointer to **int64** |  | [optional] 
**SizeSuggestQuota** | Pointer to **int64** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**SourceUuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Samples** | Pointer to [**[]DfsQuotaStat**](DfsQuotaStat.md) |  | [optional] 

## Methods

### NewDfsQuotaRecord

`func NewDfsQuotaRecord() *DfsQuotaRecord`

NewDfsQuotaRecord instantiates a new DfsQuotaRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQuotaRecordWithDefaults

`func NewDfsQuotaRecordWithDefaults() *DfsQuotaRecord`

NewDfsQuotaRecordWithDefaults instantiates a new DfsQuotaRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsQuotaRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsQuotaRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsQuotaRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsQuotaRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DfsQuotaRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsQuotaRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsQuotaRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsQuotaRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsQuotaRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsQuotaRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsQuotaRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsQuotaRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsPath

`func (o *DfsQuotaRecord) GetDfsPath() DfsPathNestview`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DfsQuotaRecord) GetDfsPathOk() (*DfsPathNestview, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DfsQuotaRecord) SetDfsPath(v DfsPathNestview)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DfsQuotaRecord) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsQuotaRecord) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsQuotaRecord) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsQuotaRecord) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsQuotaRecord) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetDomainUser

`func (o *DfsQuotaRecord) GetDomainUser() DomainUserNestview`

GetDomainUser returns the DomainUser field if non-nil, zero value otherwise.

### GetDomainUserOk

`func (o *DfsQuotaRecord) GetDomainUserOk() (*DomainUserNestview, bool)`

GetDomainUserOk returns a tuple with the DomainUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUser

`func (o *DfsQuotaRecord) SetDomainUser(v DomainUserNestview)`

SetDomainUser sets DomainUser field to given value.

### HasDomainUser

`func (o *DfsQuotaRecord) HasDomainUser() bool`

HasDomainUser returns a boolean if a field has been set.

### GetDomainUserGroup

`func (o *DfsQuotaRecord) GetDomainUserGroup() DomainUserGroupNestview`

GetDomainUserGroup returns the DomainUserGroup field if non-nil, zero value otherwise.

### GetDomainUserGroupOk

`func (o *DfsQuotaRecord) GetDomainUserGroupOk() (*DomainUserGroupNestview, bool)`

GetDomainUserGroupOk returns a tuple with the DomainUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUserGroup

`func (o *DfsQuotaRecord) SetDomainUserGroup(v DomainUserGroupNestview)`

SetDomainUserGroup sets DomainUserGroup field to given value.

### HasDomainUserGroup

`func (o *DfsQuotaRecord) HasDomainUserGroup() bool`

HasDomainUserGroup returns a boolean if a field has been set.

### GetFilesExceedTime

`func (o *DfsQuotaRecord) GetFilesExceedTime() time.Time`

GetFilesExceedTime returns the FilesExceedTime field if non-nil, zero value otherwise.

### GetFilesExceedTimeOk

`func (o *DfsQuotaRecord) GetFilesExceedTimeOk() (*time.Time, bool)`

GetFilesExceedTimeOk returns a tuple with the FilesExceedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesExceedTime

`func (o *DfsQuotaRecord) SetFilesExceedTime(v time.Time)`

SetFilesExceedTime sets FilesExceedTime field to given value.

### HasFilesExceedTime

`func (o *DfsQuotaRecord) HasFilesExceedTime() bool`

HasFilesExceedTime returns a boolean if a field has been set.

### GetFilesGraceTime

`func (o *DfsQuotaRecord) GetFilesGraceTime() int64`

GetFilesGraceTime returns the FilesGraceTime field if non-nil, zero value otherwise.

### GetFilesGraceTimeOk

`func (o *DfsQuotaRecord) GetFilesGraceTimeOk() (*int64, bool)`

GetFilesGraceTimeOk returns a tuple with the FilesGraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesGraceTime

`func (o *DfsQuotaRecord) SetFilesGraceTime(v int64)`

SetFilesGraceTime sets FilesGraceTime field to given value.

### HasFilesGraceTime

`func (o *DfsQuotaRecord) HasFilesGraceTime() bool`

HasFilesGraceTime returns a boolean if a field has been set.

### GetFilesHardQuota

`func (o *DfsQuotaRecord) GetFilesHardQuota() int64`

GetFilesHardQuota returns the FilesHardQuota field if non-nil, zero value otherwise.

### GetFilesHardQuotaOk

`func (o *DfsQuotaRecord) GetFilesHardQuotaOk() (*int64, bool)`

GetFilesHardQuotaOk returns a tuple with the FilesHardQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesHardQuota

`func (o *DfsQuotaRecord) SetFilesHardQuota(v int64)`

SetFilesHardQuota sets FilesHardQuota field to given value.

### HasFilesHardQuota

`func (o *DfsQuotaRecord) HasFilesHardQuota() bool`

HasFilesHardQuota returns a boolean if a field has been set.

### GetFilesRatio

`func (o *DfsQuotaRecord) GetFilesRatio() float32`

GetFilesRatio returns the FilesRatio field if non-nil, zero value otherwise.

### GetFilesRatioOk

`func (o *DfsQuotaRecord) GetFilesRatioOk() (*float32, bool)`

GetFilesRatioOk returns a tuple with the FilesRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesRatio

`func (o *DfsQuotaRecord) SetFilesRatio(v float32)`

SetFilesRatio sets FilesRatio field to given value.

### HasFilesRatio

`func (o *DfsQuotaRecord) HasFilesRatio() bool`

HasFilesRatio returns a boolean if a field has been set.

### GetFilesSoftQuota

`func (o *DfsQuotaRecord) GetFilesSoftQuota() int64`

GetFilesSoftQuota returns the FilesSoftQuota field if non-nil, zero value otherwise.

### GetFilesSoftQuotaOk

`func (o *DfsQuotaRecord) GetFilesSoftQuotaOk() (*int64, bool)`

GetFilesSoftQuotaOk returns a tuple with the FilesSoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSoftQuota

`func (o *DfsQuotaRecord) SetFilesSoftQuota(v int64)`

SetFilesSoftQuota sets FilesSoftQuota field to given value.

### HasFilesSoftQuota

`func (o *DfsQuotaRecord) HasFilesSoftQuota() bool`

HasFilesSoftQuota returns a boolean if a field has been set.

### GetFilesSuggestQuota

`func (o *DfsQuotaRecord) GetFilesSuggestQuota() int64`

GetFilesSuggestQuota returns the FilesSuggestQuota field if non-nil, zero value otherwise.

### GetFilesSuggestQuotaOk

`func (o *DfsQuotaRecord) GetFilesSuggestQuotaOk() (*int64, bool)`

GetFilesSuggestQuotaOk returns a tuple with the FilesSuggestQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSuggestQuota

`func (o *DfsQuotaRecord) SetFilesSuggestQuota(v int64)`

SetFilesSuggestQuota sets FilesSuggestQuota field to given value.

### HasFilesSuggestQuota

`func (o *DfsQuotaRecord) HasFilesSuggestQuota() bool`

HasFilesSuggestQuota returns a boolean if a field has been set.

### GetFsUser

`func (o *DfsQuotaRecord) GetFsUser() FSUserNestview`

GetFsUser returns the FsUser field if non-nil, zero value otherwise.

### GetFsUserOk

`func (o *DfsQuotaRecord) GetFsUserOk() (*FSUserNestview, bool)`

GetFsUserOk returns a tuple with the FsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUser

`func (o *DfsQuotaRecord) SetFsUser(v FSUserNestview)`

SetFsUser sets FsUser field to given value.

### HasFsUser

`func (o *DfsQuotaRecord) HasFsUser() bool`

HasFsUser returns a boolean if a field has been set.

### GetFsUserGroup

`func (o *DfsQuotaRecord) GetFsUserGroup() FSUserGroupNestview`

GetFsUserGroup returns the FsUserGroup field if non-nil, zero value otherwise.

### GetFsUserGroupOk

`func (o *DfsQuotaRecord) GetFsUserGroupOk() (*FSUserGroupNestview, bool)`

GetFsUserGroupOk returns a tuple with the FsUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroup

`func (o *DfsQuotaRecord) SetFsUserGroup(v FSUserGroupNestview)`

SetFsUserGroup sets FsUserGroup field to given value.

### HasFsUserGroup

`func (o *DfsQuotaRecord) HasFsUserGroup() bool`

HasFsUserGroup returns a boolean if a field has been set.

### GetId

`func (o *DfsQuotaRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsQuotaRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsQuotaRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsQuotaRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInode

`func (o *DfsQuotaRecord) GetInode() int64`

GetInode returns the Inode field if non-nil, zero value otherwise.

### GetInodeOk

`func (o *DfsQuotaRecord) GetInodeOk() (*int64, bool)`

GetInodeOk returns a tuple with the Inode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInode

`func (o *DfsQuotaRecord) SetInode(v int64)`

SetInode sets Inode field to given value.

### HasInode

`func (o *DfsQuotaRecord) HasInode() bool`

HasInode returns a boolean if a field has been set.

### GetShared

`func (o *DfsQuotaRecord) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *DfsQuotaRecord) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *DfsQuotaRecord) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *DfsQuotaRecord) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSizeExceedTime

`func (o *DfsQuotaRecord) GetSizeExceedTime() time.Time`

GetSizeExceedTime returns the SizeExceedTime field if non-nil, zero value otherwise.

### GetSizeExceedTimeOk

`func (o *DfsQuotaRecord) GetSizeExceedTimeOk() (*time.Time, bool)`

GetSizeExceedTimeOk returns a tuple with the SizeExceedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeExceedTime

`func (o *DfsQuotaRecord) SetSizeExceedTime(v time.Time)`

SetSizeExceedTime sets SizeExceedTime field to given value.

### HasSizeExceedTime

`func (o *DfsQuotaRecord) HasSizeExceedTime() bool`

HasSizeExceedTime returns a boolean if a field has been set.

### GetSizeGraceTime

`func (o *DfsQuotaRecord) GetSizeGraceTime() int64`

GetSizeGraceTime returns the SizeGraceTime field if non-nil, zero value otherwise.

### GetSizeGraceTimeOk

`func (o *DfsQuotaRecord) GetSizeGraceTimeOk() (*int64, bool)`

GetSizeGraceTimeOk returns a tuple with the SizeGraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGraceTime

`func (o *DfsQuotaRecord) SetSizeGraceTime(v int64)`

SetSizeGraceTime sets SizeGraceTime field to given value.

### HasSizeGraceTime

`func (o *DfsQuotaRecord) HasSizeGraceTime() bool`

HasSizeGraceTime returns a boolean if a field has been set.

### GetSizeHardQuota

`func (o *DfsQuotaRecord) GetSizeHardQuota() int64`

GetSizeHardQuota returns the SizeHardQuota field if non-nil, zero value otherwise.

### GetSizeHardQuotaOk

`func (o *DfsQuotaRecord) GetSizeHardQuotaOk() (*int64, bool)`

GetSizeHardQuotaOk returns a tuple with the SizeHardQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeHardQuota

`func (o *DfsQuotaRecord) SetSizeHardQuota(v int64)`

SetSizeHardQuota sets SizeHardQuota field to given value.

### HasSizeHardQuota

`func (o *DfsQuotaRecord) HasSizeHardQuota() bool`

HasSizeHardQuota returns a boolean if a field has been set.

### GetSizeRatio

`func (o *DfsQuotaRecord) GetSizeRatio() float32`

GetSizeRatio returns the SizeRatio field if non-nil, zero value otherwise.

### GetSizeRatioOk

`func (o *DfsQuotaRecord) GetSizeRatioOk() (*float32, bool)`

GetSizeRatioOk returns a tuple with the SizeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeRatio

`func (o *DfsQuotaRecord) SetSizeRatio(v float32)`

SetSizeRatio sets SizeRatio field to given value.

### HasSizeRatio

`func (o *DfsQuotaRecord) HasSizeRatio() bool`

HasSizeRatio returns a boolean if a field has been set.

### GetSizeSoftQuota

`func (o *DfsQuotaRecord) GetSizeSoftQuota() int64`

GetSizeSoftQuota returns the SizeSoftQuota field if non-nil, zero value otherwise.

### GetSizeSoftQuotaOk

`func (o *DfsQuotaRecord) GetSizeSoftQuotaOk() (*int64, bool)`

GetSizeSoftQuotaOk returns a tuple with the SizeSoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSoftQuota

`func (o *DfsQuotaRecord) SetSizeSoftQuota(v int64)`

SetSizeSoftQuota sets SizeSoftQuota field to given value.

### HasSizeSoftQuota

`func (o *DfsQuotaRecord) HasSizeSoftQuota() bool`

HasSizeSoftQuota returns a boolean if a field has been set.

### GetSizeSuggestQuota

`func (o *DfsQuotaRecord) GetSizeSuggestQuota() int64`

GetSizeSuggestQuota returns the SizeSuggestQuota field if non-nil, zero value otherwise.

### GetSizeSuggestQuotaOk

`func (o *DfsQuotaRecord) GetSizeSuggestQuotaOk() (*int64, bool)`

GetSizeSuggestQuotaOk returns a tuple with the SizeSuggestQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSuggestQuota

`func (o *DfsQuotaRecord) SetSizeSuggestQuota(v int64)`

SetSizeSuggestQuota sets SizeSuggestQuota field to given value.

### HasSizeSuggestQuota

`func (o *DfsQuotaRecord) HasSizeSuggestQuota() bool`

HasSizeSuggestQuota returns a boolean if a field has been set.

### GetSourceType

`func (o *DfsQuotaRecord) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *DfsQuotaRecord) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *DfsQuotaRecord) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *DfsQuotaRecord) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSourceUuid

`func (o *DfsQuotaRecord) GetSourceUuid() string`

GetSourceUuid returns the SourceUuid field if non-nil, zero value otherwise.

### GetSourceUuidOk

`func (o *DfsQuotaRecord) GetSourceUuidOk() (*string, bool)`

GetSourceUuidOk returns a tuple with the SourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUuid

`func (o *DfsQuotaRecord) SetSourceUuid(v string)`

SetSourceUuid sets SourceUuid field to given value.

### HasSourceUuid

`func (o *DfsQuotaRecord) HasSourceUuid() bool`

HasSourceUuid returns a boolean if a field has been set.

### GetStatus

`func (o *DfsQuotaRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsQuotaRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsQuotaRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsQuotaRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *DfsQuotaRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsQuotaRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsQuotaRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsQuotaRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsQuotaRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsQuotaRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsQuotaRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsQuotaRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *DfsQuotaRecord) GetSamples() []DfsQuotaStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DfsQuotaRecord) GetSamplesOk() (*[]DfsQuotaStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DfsQuotaRecord) SetSamples(v []DfsQuotaStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *DfsQuotaRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


