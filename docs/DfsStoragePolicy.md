# DfsStoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DfsPaths** | Pointer to [**[]DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfs**](DfsRootfs.md) |  | [optional] 
**DfsStorageClass** | Pointer to [**DfsTier**](DfsTier.md) |  | [optional] 
**FileCreate** | Pointer to **time.Time** |  | [optional] 
**FileName** | Pointer to **string** | useless fields, for compatability | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**DfsPolicyFilter**](DfsPolicyFilter.md) |  | [optional] 
**FsUser** | Pointer to [**FSUser**](FSUser.md) |  | [optional] 
**FsUserGroup** | Pointer to [**FSUserGroup**](FSUserGroup.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsDefault** | Pointer to **bool** | common properties | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PathNum** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Stretched** | Pointer to **bool** | indicate that all pools under this policy is stretched pool | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsStoragePolicy

`func NewDfsStoragePolicy() *DfsStoragePolicy`

NewDfsStoragePolicy instantiates a new DfsStoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsStoragePolicyWithDefaults

`func NewDfsStoragePolicyWithDefaults() *DfsStoragePolicy`

NewDfsStoragePolicyWithDefaults instantiates a new DfsStoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsStoragePolicy) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsStoragePolicy) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsStoragePolicy) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsStoragePolicy) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *DfsStoragePolicy) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsStoragePolicy) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsStoragePolicy) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsStoragePolicy) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsStoragePolicy) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsStoragePolicy) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsStoragePolicy) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsStoragePolicy) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DfsStoragePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsStoragePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsStoragePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsStoragePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsPaths

`func (o *DfsStoragePolicy) GetDfsPaths() []DfsPathNestview`

GetDfsPaths returns the DfsPaths field if non-nil, zero value otherwise.

### GetDfsPathsOk

`func (o *DfsStoragePolicy) GetDfsPathsOk() (*[]DfsPathNestview, bool)`

GetDfsPathsOk returns a tuple with the DfsPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPaths

`func (o *DfsStoragePolicy) SetDfsPaths(v []DfsPathNestview)`

SetDfsPaths sets DfsPaths field to given value.

### HasDfsPaths

`func (o *DfsStoragePolicy) HasDfsPaths() bool`

HasDfsPaths returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsStoragePolicy) GetDfsRootfs() DfsRootfs`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsStoragePolicy) GetDfsRootfsOk() (*DfsRootfs, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsStoragePolicy) SetDfsRootfs(v DfsRootfs)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsStoragePolicy) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetDfsStorageClass

`func (o *DfsStoragePolicy) GetDfsStorageClass() DfsTier`

GetDfsStorageClass returns the DfsStorageClass field if non-nil, zero value otherwise.

### GetDfsStorageClassOk

`func (o *DfsStoragePolicy) GetDfsStorageClassOk() (*DfsTier, bool)`

GetDfsStorageClassOk returns a tuple with the DfsStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsStorageClass

`func (o *DfsStoragePolicy) SetDfsStorageClass(v DfsTier)`

SetDfsStorageClass sets DfsStorageClass field to given value.

### HasDfsStorageClass

`func (o *DfsStoragePolicy) HasDfsStorageClass() bool`

HasDfsStorageClass returns a boolean if a field has been set.

### GetFileCreate

`func (o *DfsStoragePolicy) GetFileCreate() time.Time`

GetFileCreate returns the FileCreate field if non-nil, zero value otherwise.

### GetFileCreateOk

`func (o *DfsStoragePolicy) GetFileCreateOk() (*time.Time, bool)`

GetFileCreateOk returns a tuple with the FileCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCreate

`func (o *DfsStoragePolicy) SetFileCreate(v time.Time)`

SetFileCreate sets FileCreate field to given value.

### HasFileCreate

`func (o *DfsStoragePolicy) HasFileCreate() bool`

HasFileCreate returns a boolean if a field has been set.

### GetFileName

`func (o *DfsStoragePolicy) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DfsStoragePolicy) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DfsStoragePolicy) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DfsStoragePolicy) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFilePath

`func (o *DfsStoragePolicy) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *DfsStoragePolicy) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *DfsStoragePolicy) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *DfsStoragePolicy) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFilter

`func (o *DfsStoragePolicy) GetFilter() DfsPolicyFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DfsStoragePolicy) GetFilterOk() (*DfsPolicyFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DfsStoragePolicy) SetFilter(v DfsPolicyFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *DfsStoragePolicy) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFsUser

`func (o *DfsStoragePolicy) GetFsUser() FSUser`

GetFsUser returns the FsUser field if non-nil, zero value otherwise.

### GetFsUserOk

`func (o *DfsStoragePolicy) GetFsUserOk() (*FSUser, bool)`

GetFsUserOk returns a tuple with the FsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUser

`func (o *DfsStoragePolicy) SetFsUser(v FSUser)`

SetFsUser sets FsUser field to given value.

### HasFsUser

`func (o *DfsStoragePolicy) HasFsUser() bool`

HasFsUser returns a boolean if a field has been set.

### GetFsUserGroup

`func (o *DfsStoragePolicy) GetFsUserGroup() FSUserGroup`

GetFsUserGroup returns the FsUserGroup field if non-nil, zero value otherwise.

### GetFsUserGroupOk

`func (o *DfsStoragePolicy) GetFsUserGroupOk() (*FSUserGroup, bool)`

GetFsUserGroupOk returns a tuple with the FsUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroup

`func (o *DfsStoragePolicy) SetFsUserGroup(v FSUserGroup)`

SetFsUserGroup sets FsUserGroup field to given value.

### HasFsUserGroup

`func (o *DfsStoragePolicy) HasFsUserGroup() bool`

HasFsUserGroup returns a boolean if a field has been set.

### GetId

`func (o *DfsStoragePolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsStoragePolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsStoragePolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsStoragePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *DfsStoragePolicy) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *DfsStoragePolicy) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *DfsStoragePolicy) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *DfsStoragePolicy) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *DfsStoragePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsStoragePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsStoragePolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsStoragePolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPathNum

`func (o *DfsStoragePolicy) GetPathNum() int64`

GetPathNum returns the PathNum field if non-nil, zero value otherwise.

### GetPathNumOk

`func (o *DfsStoragePolicy) GetPathNumOk() (*int64, bool)`

GetPathNumOk returns a tuple with the PathNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathNum

`func (o *DfsStoragePolicy) SetPathNum(v int64)`

SetPathNum sets PathNum field to given value.

### HasPathNum

`func (o *DfsStoragePolicy) HasPathNum() bool`

HasPathNum returns a boolean if a field has been set.

### GetStatus

`func (o *DfsStoragePolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsStoragePolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsStoragePolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsStoragePolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStretched

`func (o *DfsStoragePolicy) GetStretched() bool`

GetStretched returns the Stretched field if non-nil, zero value otherwise.

### GetStretchedOk

`func (o *DfsStoragePolicy) GetStretchedOk() (*bool, bool)`

GetStretchedOk returns a tuple with the Stretched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStretched

`func (o *DfsStoragePolicy) SetStretched(v bool)`

SetStretched sets Stretched field to given value.

### HasStretched

`func (o *DfsStoragePolicy) HasStretched() bool`

HasStretched returns a boolean if a field has been set.

### GetType

`func (o *DfsStoragePolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsStoragePolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsStoragePolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsStoragePolicy) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsStoragePolicy) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsStoragePolicy) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsStoragePolicy) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsStoragePolicy) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


