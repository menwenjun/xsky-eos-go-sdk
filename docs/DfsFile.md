# DfsFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **time.Time** | access time | [optional] 
**Bucket** | Pointer to [**DfsS3Bucket**](DfsS3Bucket.md) |  | [optional] 
**Change** | Pointer to **time.Time** | change time | [optional] 
**DfsPathId** | Pointer to **int64** | dfs path id | [optional] 
**DfsPathPerformanceId** | Pointer to **int64** | dfs path performance id | [optional] 
**DfsRootfs** | Pointer to [**NestedRootfs**](NestedRootfs.md) |  | [optional] 
**DfsStoragePolicies** | Pointer to [**[]DfsStoragePolicy**](DfsStoragePolicy.md) | dfs storage policies | [optional] 
**DirQuotaNum** | Pointer to **int64** | count of dir quota | [optional] 
**DpSnapshotNum** | Pointer to **int64** | count of data protection snapshot | [optional] 
**Files** | Pointer to **int64** | sub file count when it is a directory | [optional] 
**FtpShare** | Pointer to [**DfsFTPShare**](DfsFTPShare.md) |  | [optional] 
**FullPathSnapNum** | Pointer to **int64** | count of full path snapshot | [optional] 
**Group** | Pointer to **int64** | owner group | [optional] 
**HdfsNum** | Pointer to **int64** | count of hdfs | [optional] 
**Hdfses** | Pointer to [**[]DfsHdfs**](DfsHdfs.md) | hdfses | [optional] 
**Inode** | Pointer to **int64** | inode count | [optional] 
**IsBucket** | Pointer to **bool** | is bucket path | [optional] 
**IsBucketParent** | Pointer to **bool** | is bucked parent path | [optional] 
**LocalGroupInfo** | Pointer to [**LocalGroupInfo**](LocalGroupInfo.md) |  | [optional] 
**LocalUserInfo** | Pointer to [**LocalUserInfo**](LocalUserInfo.md) |  | [optional] 
**Mode** | Pointer to **string** | file mode | [optional] 
**Modify** | Pointer to **time.Time** | modify time | [optional] 
**Name** | Pointer to **string** | file name | [optional] 
**NfsShare** | Pointer to [**DfsNFSShare**](DfsNFSShare.md) |  | [optional] 
**OriginalName** | Pointer to **string** | original name before file moved to trash | [optional] 
**Owner** | Pointer to **int64** | owner user | [optional] 
**Parent** | Pointer to **string** | parent path | [optional] 
**ParentPathStoragePolicies** | Pointer to [**[]DfsStoragePolicy**](DfsStoragePolicy.md) | par path link dfs storage policies | [optional] 
**Path** | Pointer to **string** | full path | [optional] 
**Qos** | Pointer to [**DfsFileQos**](DfsFileQos.md) |  | [optional] 
**QuotaNum** | Pointer to **int64** | count of quota | [optional] 
**Shared** | Pointer to **bool** | shared | [optional] 
**Shares** | Pointer to **[]string** | share types | [optional] 
**Size** | Pointer to **int64** | file size | [optional] 
**SmbShares** | Pointer to [**[]DfsSMBShare**](DfsSMBShare.md) | smb shares | [optional] 
**SnapshotNum** | Pointer to **int64** | count of snapshot | [optional] 
**Stretched** | Pointer to **bool** | is stretched directory | [optional] 
**TotalSnapshotNum** | Pointer to **int64** | count of total snapshot | [optional] 
**Trash** | Pointer to [**DfsFileTrash**](DfsFileTrash.md) |  | [optional] 
**Type** | Pointer to **string** | file type | [optional] 
**Worm** | Pointer to [**DfsFileWorm**](DfsFileWorm.md) |  | [optional] 

## Methods

### NewDfsFile

`func NewDfsFile() *DfsFile`

NewDfsFile instantiates a new DfsFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFileWithDefaults

`func NewDfsFileWithDefaults() *DfsFile`

NewDfsFileWithDefaults instantiates a new DfsFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *DfsFile) GetAccess() time.Time`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DfsFile) GetAccessOk() (*time.Time, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DfsFile) SetAccess(v time.Time)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *DfsFile) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetBucket

`func (o *DfsFile) GetBucket() DfsS3Bucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *DfsFile) GetBucketOk() (*DfsS3Bucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *DfsFile) SetBucket(v DfsS3Bucket)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *DfsFile) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetChange

`func (o *DfsFile) GetChange() time.Time`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *DfsFile) GetChangeOk() (*time.Time, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *DfsFile) SetChange(v time.Time)`

SetChange sets Change field to given value.

### HasChange

`func (o *DfsFile) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetDfsPathId

`func (o *DfsFile) GetDfsPathId() int64`

GetDfsPathId returns the DfsPathId field if non-nil, zero value otherwise.

### GetDfsPathIdOk

`func (o *DfsFile) GetDfsPathIdOk() (*int64, bool)`

GetDfsPathIdOk returns a tuple with the DfsPathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPathId

`func (o *DfsFile) SetDfsPathId(v int64)`

SetDfsPathId sets DfsPathId field to given value.

### HasDfsPathId

`func (o *DfsFile) HasDfsPathId() bool`

HasDfsPathId returns a boolean if a field has been set.

### GetDfsPathPerformanceId

`func (o *DfsFile) GetDfsPathPerformanceId() int64`

GetDfsPathPerformanceId returns the DfsPathPerformanceId field if non-nil, zero value otherwise.

### GetDfsPathPerformanceIdOk

`func (o *DfsFile) GetDfsPathPerformanceIdOk() (*int64, bool)`

GetDfsPathPerformanceIdOk returns a tuple with the DfsPathPerformanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPathPerformanceId

`func (o *DfsFile) SetDfsPathPerformanceId(v int64)`

SetDfsPathPerformanceId sets DfsPathPerformanceId field to given value.

### HasDfsPathPerformanceId

`func (o *DfsFile) HasDfsPathPerformanceId() bool`

HasDfsPathPerformanceId returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsFile) GetDfsRootfs() NestedRootfs`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsFile) GetDfsRootfsOk() (*NestedRootfs, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsFile) SetDfsRootfs(v NestedRootfs)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsFile) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetDfsStoragePolicies

`func (o *DfsFile) GetDfsStoragePolicies() []DfsStoragePolicy`

GetDfsStoragePolicies returns the DfsStoragePolicies field if non-nil, zero value otherwise.

### GetDfsStoragePoliciesOk

`func (o *DfsFile) GetDfsStoragePoliciesOk() (*[]DfsStoragePolicy, bool)`

GetDfsStoragePoliciesOk returns a tuple with the DfsStoragePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsStoragePolicies

`func (o *DfsFile) SetDfsStoragePolicies(v []DfsStoragePolicy)`

SetDfsStoragePolicies sets DfsStoragePolicies field to given value.

### HasDfsStoragePolicies

`func (o *DfsFile) HasDfsStoragePolicies() bool`

HasDfsStoragePolicies returns a boolean if a field has been set.

### GetDirQuotaNum

`func (o *DfsFile) GetDirQuotaNum() int64`

GetDirQuotaNum returns the DirQuotaNum field if non-nil, zero value otherwise.

### GetDirQuotaNumOk

`func (o *DfsFile) GetDirQuotaNumOk() (*int64, bool)`

GetDirQuotaNumOk returns a tuple with the DirQuotaNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirQuotaNum

`func (o *DfsFile) SetDirQuotaNum(v int64)`

SetDirQuotaNum sets DirQuotaNum field to given value.

### HasDirQuotaNum

`func (o *DfsFile) HasDirQuotaNum() bool`

HasDirQuotaNum returns a boolean if a field has been set.

### GetDpSnapshotNum

`func (o *DfsFile) GetDpSnapshotNum() int64`

GetDpSnapshotNum returns the DpSnapshotNum field if non-nil, zero value otherwise.

### GetDpSnapshotNumOk

`func (o *DfsFile) GetDpSnapshotNumOk() (*int64, bool)`

GetDpSnapshotNumOk returns a tuple with the DpSnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpSnapshotNum

`func (o *DfsFile) SetDpSnapshotNum(v int64)`

SetDpSnapshotNum sets DpSnapshotNum field to given value.

### HasDpSnapshotNum

`func (o *DfsFile) HasDpSnapshotNum() bool`

HasDpSnapshotNum returns a boolean if a field has been set.

### GetFiles

`func (o *DfsFile) GetFiles() int64`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *DfsFile) GetFilesOk() (*int64, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *DfsFile) SetFiles(v int64)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *DfsFile) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFtpShare

`func (o *DfsFile) GetFtpShare() DfsFTPShare`

GetFtpShare returns the FtpShare field if non-nil, zero value otherwise.

### GetFtpShareOk

`func (o *DfsFile) GetFtpShareOk() (*DfsFTPShare, bool)`

GetFtpShareOk returns a tuple with the FtpShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpShare

`func (o *DfsFile) SetFtpShare(v DfsFTPShare)`

SetFtpShare sets FtpShare field to given value.

### HasFtpShare

`func (o *DfsFile) HasFtpShare() bool`

HasFtpShare returns a boolean if a field has been set.

### GetFullPathSnapNum

`func (o *DfsFile) GetFullPathSnapNum() int64`

GetFullPathSnapNum returns the FullPathSnapNum field if non-nil, zero value otherwise.

### GetFullPathSnapNumOk

`func (o *DfsFile) GetFullPathSnapNumOk() (*int64, bool)`

GetFullPathSnapNumOk returns a tuple with the FullPathSnapNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPathSnapNum

`func (o *DfsFile) SetFullPathSnapNum(v int64)`

SetFullPathSnapNum sets FullPathSnapNum field to given value.

### HasFullPathSnapNum

`func (o *DfsFile) HasFullPathSnapNum() bool`

HasFullPathSnapNum returns a boolean if a field has been set.

### GetGroup

`func (o *DfsFile) GetGroup() int64`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DfsFile) GetGroupOk() (*int64, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DfsFile) SetGroup(v int64)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DfsFile) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHdfsNum

`func (o *DfsFile) GetHdfsNum() int64`

GetHdfsNum returns the HdfsNum field if non-nil, zero value otherwise.

### GetHdfsNumOk

`func (o *DfsFile) GetHdfsNumOk() (*int64, bool)`

GetHdfsNumOk returns a tuple with the HdfsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsNum

`func (o *DfsFile) SetHdfsNum(v int64)`

SetHdfsNum sets HdfsNum field to given value.

### HasHdfsNum

`func (o *DfsFile) HasHdfsNum() bool`

HasHdfsNum returns a boolean if a field has been set.

### GetHdfses

`func (o *DfsFile) GetHdfses() []DfsHdfs`

GetHdfses returns the Hdfses field if non-nil, zero value otherwise.

### GetHdfsesOk

`func (o *DfsFile) GetHdfsesOk() (*[]DfsHdfs, bool)`

GetHdfsesOk returns a tuple with the Hdfses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfses

`func (o *DfsFile) SetHdfses(v []DfsHdfs)`

SetHdfses sets Hdfses field to given value.

### HasHdfses

`func (o *DfsFile) HasHdfses() bool`

HasHdfses returns a boolean if a field has been set.

### GetInode

`func (o *DfsFile) GetInode() int64`

GetInode returns the Inode field if non-nil, zero value otherwise.

### GetInodeOk

`func (o *DfsFile) GetInodeOk() (*int64, bool)`

GetInodeOk returns a tuple with the Inode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInode

`func (o *DfsFile) SetInode(v int64)`

SetInode sets Inode field to given value.

### HasInode

`func (o *DfsFile) HasInode() bool`

HasInode returns a boolean if a field has been set.

### GetIsBucket

`func (o *DfsFile) GetIsBucket() bool`

GetIsBucket returns the IsBucket field if non-nil, zero value otherwise.

### GetIsBucketOk

`func (o *DfsFile) GetIsBucketOk() (*bool, bool)`

GetIsBucketOk returns a tuple with the IsBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBucket

`func (o *DfsFile) SetIsBucket(v bool)`

SetIsBucket sets IsBucket field to given value.

### HasIsBucket

`func (o *DfsFile) HasIsBucket() bool`

HasIsBucket returns a boolean if a field has been set.

### GetIsBucketParent

`func (o *DfsFile) GetIsBucketParent() bool`

GetIsBucketParent returns the IsBucketParent field if non-nil, zero value otherwise.

### GetIsBucketParentOk

`func (o *DfsFile) GetIsBucketParentOk() (*bool, bool)`

GetIsBucketParentOk returns a tuple with the IsBucketParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBucketParent

`func (o *DfsFile) SetIsBucketParent(v bool)`

SetIsBucketParent sets IsBucketParent field to given value.

### HasIsBucketParent

`func (o *DfsFile) HasIsBucketParent() bool`

HasIsBucketParent returns a boolean if a field has been set.

### GetLocalGroupInfo

`func (o *DfsFile) GetLocalGroupInfo() LocalGroupInfo`

GetLocalGroupInfo returns the LocalGroupInfo field if non-nil, zero value otherwise.

### GetLocalGroupInfoOk

`func (o *DfsFile) GetLocalGroupInfoOk() (*LocalGroupInfo, bool)`

GetLocalGroupInfoOk returns a tuple with the LocalGroupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGroupInfo

`func (o *DfsFile) SetLocalGroupInfo(v LocalGroupInfo)`

SetLocalGroupInfo sets LocalGroupInfo field to given value.

### HasLocalGroupInfo

`func (o *DfsFile) HasLocalGroupInfo() bool`

HasLocalGroupInfo returns a boolean if a field has been set.

### GetLocalUserInfo

`func (o *DfsFile) GetLocalUserInfo() LocalUserInfo`

GetLocalUserInfo returns the LocalUserInfo field if non-nil, zero value otherwise.

### GetLocalUserInfoOk

`func (o *DfsFile) GetLocalUserInfoOk() (*LocalUserInfo, bool)`

GetLocalUserInfoOk returns a tuple with the LocalUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserInfo

`func (o *DfsFile) SetLocalUserInfo(v LocalUserInfo)`

SetLocalUserInfo sets LocalUserInfo field to given value.

### HasLocalUserInfo

`func (o *DfsFile) HasLocalUserInfo() bool`

HasLocalUserInfo returns a boolean if a field has been set.

### GetMode

`func (o *DfsFile) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DfsFile) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DfsFile) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DfsFile) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetModify

`func (o *DfsFile) GetModify() time.Time`

GetModify returns the Modify field if non-nil, zero value otherwise.

### GetModifyOk

`func (o *DfsFile) GetModifyOk() (*time.Time, bool)`

GetModifyOk returns a tuple with the Modify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModify

`func (o *DfsFile) SetModify(v time.Time)`

SetModify sets Modify field to given value.

### HasModify

`func (o *DfsFile) HasModify() bool`

HasModify returns a boolean if a field has been set.

### GetName

`func (o *DfsFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNfsShare

`func (o *DfsFile) GetNfsShare() DfsNFSShare`

GetNfsShare returns the NfsShare field if non-nil, zero value otherwise.

### GetNfsShareOk

`func (o *DfsFile) GetNfsShareOk() (*DfsNFSShare, bool)`

GetNfsShareOk returns a tuple with the NfsShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsShare

`func (o *DfsFile) SetNfsShare(v DfsNFSShare)`

SetNfsShare sets NfsShare field to given value.

### HasNfsShare

`func (o *DfsFile) HasNfsShare() bool`

HasNfsShare returns a boolean if a field has been set.

### GetOriginalName

`func (o *DfsFile) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *DfsFile) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *DfsFile) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *DfsFile) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetOwner

`func (o *DfsFile) GetOwner() int64`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DfsFile) GetOwnerOk() (*int64, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DfsFile) SetOwner(v int64)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DfsFile) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParent

`func (o *DfsFile) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *DfsFile) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *DfsFile) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *DfsFile) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetParentPathStoragePolicies

`func (o *DfsFile) GetParentPathStoragePolicies() []DfsStoragePolicy`

GetParentPathStoragePolicies returns the ParentPathStoragePolicies field if non-nil, zero value otherwise.

### GetParentPathStoragePoliciesOk

`func (o *DfsFile) GetParentPathStoragePoliciesOk() (*[]DfsStoragePolicy, bool)`

GetParentPathStoragePoliciesOk returns a tuple with the ParentPathStoragePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPathStoragePolicies

`func (o *DfsFile) SetParentPathStoragePolicies(v []DfsStoragePolicy)`

SetParentPathStoragePolicies sets ParentPathStoragePolicies field to given value.

### HasParentPathStoragePolicies

`func (o *DfsFile) HasParentPathStoragePolicies() bool`

HasParentPathStoragePolicies returns a boolean if a field has been set.

### GetPath

`func (o *DfsFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetQos

`func (o *DfsFile) GetQos() DfsFileQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *DfsFile) GetQosOk() (*DfsFileQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *DfsFile) SetQos(v DfsFileQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *DfsFile) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetQuotaNum

`func (o *DfsFile) GetQuotaNum() int64`

GetQuotaNum returns the QuotaNum field if non-nil, zero value otherwise.

### GetQuotaNumOk

`func (o *DfsFile) GetQuotaNumOk() (*int64, bool)`

GetQuotaNumOk returns a tuple with the QuotaNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaNum

`func (o *DfsFile) SetQuotaNum(v int64)`

SetQuotaNum sets QuotaNum field to given value.

### HasQuotaNum

`func (o *DfsFile) HasQuotaNum() bool`

HasQuotaNum returns a boolean if a field has been set.

### GetShared

`func (o *DfsFile) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *DfsFile) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *DfsFile) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *DfsFile) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetShares

`func (o *DfsFile) GetShares() []string`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *DfsFile) GetSharesOk() (*[]string, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *DfsFile) SetShares(v []string)`

SetShares sets Shares field to given value.

### HasShares

`func (o *DfsFile) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetSize

`func (o *DfsFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DfsFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DfsFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DfsFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSmbShares

`func (o *DfsFile) GetSmbShares() []DfsSMBShare`

GetSmbShares returns the SmbShares field if non-nil, zero value otherwise.

### GetSmbSharesOk

`func (o *DfsFile) GetSmbSharesOk() (*[]DfsSMBShare, bool)`

GetSmbSharesOk returns a tuple with the SmbShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbShares

`func (o *DfsFile) SetSmbShares(v []DfsSMBShare)`

SetSmbShares sets SmbShares field to given value.

### HasSmbShares

`func (o *DfsFile) HasSmbShares() bool`

HasSmbShares returns a boolean if a field has been set.

### GetSnapshotNum

`func (o *DfsFile) GetSnapshotNum() int64`

GetSnapshotNum returns the SnapshotNum field if non-nil, zero value otherwise.

### GetSnapshotNumOk

`func (o *DfsFile) GetSnapshotNumOk() (*int64, bool)`

GetSnapshotNumOk returns a tuple with the SnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotNum

`func (o *DfsFile) SetSnapshotNum(v int64)`

SetSnapshotNum sets SnapshotNum field to given value.

### HasSnapshotNum

`func (o *DfsFile) HasSnapshotNum() bool`

HasSnapshotNum returns a boolean if a field has been set.

### GetStretched

`func (o *DfsFile) GetStretched() bool`

GetStretched returns the Stretched field if non-nil, zero value otherwise.

### GetStretchedOk

`func (o *DfsFile) GetStretchedOk() (*bool, bool)`

GetStretchedOk returns a tuple with the Stretched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStretched

`func (o *DfsFile) SetStretched(v bool)`

SetStretched sets Stretched field to given value.

### HasStretched

`func (o *DfsFile) HasStretched() bool`

HasStretched returns a boolean if a field has been set.

### GetTotalSnapshotNum

`func (o *DfsFile) GetTotalSnapshotNum() int64`

GetTotalSnapshotNum returns the TotalSnapshotNum field if non-nil, zero value otherwise.

### GetTotalSnapshotNumOk

`func (o *DfsFile) GetTotalSnapshotNumOk() (*int64, bool)`

GetTotalSnapshotNumOk returns a tuple with the TotalSnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSnapshotNum

`func (o *DfsFile) SetTotalSnapshotNum(v int64)`

SetTotalSnapshotNum sets TotalSnapshotNum field to given value.

### HasTotalSnapshotNum

`func (o *DfsFile) HasTotalSnapshotNum() bool`

HasTotalSnapshotNum returns a boolean if a field has been set.

### GetTrash

`func (o *DfsFile) GetTrash() DfsFileTrash`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *DfsFile) GetTrashOk() (*DfsFileTrash, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *DfsFile) SetTrash(v DfsFileTrash)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *DfsFile) HasTrash() bool`

HasTrash returns a boolean if a field has been set.

### GetType

`func (o *DfsFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsFile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorm

`func (o *DfsFile) GetWorm() DfsFileWorm`

GetWorm returns the Worm field if non-nil, zero value otherwise.

### GetWormOk

`func (o *DfsFile) GetWormOk() (*DfsFileWorm, bool)`

GetWormOk returns a tuple with the Worm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorm

`func (o *DfsFile) SetWorm(v DfsFileWorm)`

SetWorm sets Worm field to given value.

### HasWorm

`func (o *DfsFile) HasWorm() bool`

HasWorm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


