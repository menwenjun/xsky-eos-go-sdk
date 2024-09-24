# DfsTrashFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **time.Time** | access time | [optional] 
**BaseFileId** | Pointer to **string** | base file id in trash | [optional] 
**BaseName** | Pointer to **string** | file base name in trash | [optional] 
**Change** | Pointer to **time.Time** | change time | [optional] 
**DfsRootfs** | Pointer to [**NestedRootfs**](NestedRootfs.md) |  | [optional] 
**DfsTrash** | Pointer to [**NestedTrash**](NestedTrash.md) |  | [optional] 
**Files** | Pointer to **int64** | sub file count when it is a directory | [optional] 
**Group** | Pointer to **int64** | file owner user group | [optional] 
**Inode** | Pointer to **int64** | file inode | [optional] 
**Modify** | Pointer to **time.Time** | modify time | [optional] 
**Name** | Pointer to **string** | file name | [optional] 
**OriginalName** | Pointer to **string** | original name before file moved to trash | [optional] 
**Owner** | Pointer to **int64** | file owner user | [optional] 
**Parent** | Pointer to **string** | parent path | [optional] 
**Size** | Pointer to **int64** | file size | [optional] 
**TrashPath** | Pointer to **string** | trash path | [optional] 
**Type** | Pointer to **string** | file type | [optional] 

## Methods

### NewDfsTrashFile

`func NewDfsTrashFile() *DfsTrashFile`

NewDfsTrashFile instantiates a new DfsTrashFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsTrashFileWithDefaults

`func NewDfsTrashFileWithDefaults() *DfsTrashFile`

NewDfsTrashFileWithDefaults instantiates a new DfsTrashFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *DfsTrashFile) GetAccess() time.Time`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DfsTrashFile) GetAccessOk() (*time.Time, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DfsTrashFile) SetAccess(v time.Time)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *DfsTrashFile) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetBaseFileId

`func (o *DfsTrashFile) GetBaseFileId() string`

GetBaseFileId returns the BaseFileId field if non-nil, zero value otherwise.

### GetBaseFileIdOk

`func (o *DfsTrashFile) GetBaseFileIdOk() (*string, bool)`

GetBaseFileIdOk returns a tuple with the BaseFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFileId

`func (o *DfsTrashFile) SetBaseFileId(v string)`

SetBaseFileId sets BaseFileId field to given value.

### HasBaseFileId

`func (o *DfsTrashFile) HasBaseFileId() bool`

HasBaseFileId returns a boolean if a field has been set.

### GetBaseName

`func (o *DfsTrashFile) GetBaseName() string`

GetBaseName returns the BaseName field if non-nil, zero value otherwise.

### GetBaseNameOk

`func (o *DfsTrashFile) GetBaseNameOk() (*string, bool)`

GetBaseNameOk returns a tuple with the BaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseName

`func (o *DfsTrashFile) SetBaseName(v string)`

SetBaseName sets BaseName field to given value.

### HasBaseName

`func (o *DfsTrashFile) HasBaseName() bool`

HasBaseName returns a boolean if a field has been set.

### GetChange

`func (o *DfsTrashFile) GetChange() time.Time`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *DfsTrashFile) GetChangeOk() (*time.Time, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *DfsTrashFile) SetChange(v time.Time)`

SetChange sets Change field to given value.

### HasChange

`func (o *DfsTrashFile) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsTrashFile) GetDfsRootfs() NestedRootfs`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsTrashFile) GetDfsRootfsOk() (*NestedRootfs, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsTrashFile) SetDfsRootfs(v NestedRootfs)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsTrashFile) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetDfsTrash

`func (o *DfsTrashFile) GetDfsTrash() NestedTrash`

GetDfsTrash returns the DfsTrash field if non-nil, zero value otherwise.

### GetDfsTrashOk

`func (o *DfsTrashFile) GetDfsTrashOk() (*NestedTrash, bool)`

GetDfsTrashOk returns a tuple with the DfsTrash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsTrash

`func (o *DfsTrashFile) SetDfsTrash(v NestedTrash)`

SetDfsTrash sets DfsTrash field to given value.

### HasDfsTrash

`func (o *DfsTrashFile) HasDfsTrash() bool`

HasDfsTrash returns a boolean if a field has been set.

### GetFiles

`func (o *DfsTrashFile) GetFiles() int64`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *DfsTrashFile) GetFilesOk() (*int64, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *DfsTrashFile) SetFiles(v int64)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *DfsTrashFile) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetGroup

`func (o *DfsTrashFile) GetGroup() int64`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DfsTrashFile) GetGroupOk() (*int64, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DfsTrashFile) SetGroup(v int64)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DfsTrashFile) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInode

`func (o *DfsTrashFile) GetInode() int64`

GetInode returns the Inode field if non-nil, zero value otherwise.

### GetInodeOk

`func (o *DfsTrashFile) GetInodeOk() (*int64, bool)`

GetInodeOk returns a tuple with the Inode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInode

`func (o *DfsTrashFile) SetInode(v int64)`

SetInode sets Inode field to given value.

### HasInode

`func (o *DfsTrashFile) HasInode() bool`

HasInode returns a boolean if a field has been set.

### GetModify

`func (o *DfsTrashFile) GetModify() time.Time`

GetModify returns the Modify field if non-nil, zero value otherwise.

### GetModifyOk

`func (o *DfsTrashFile) GetModifyOk() (*time.Time, bool)`

GetModifyOk returns a tuple with the Modify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModify

`func (o *DfsTrashFile) SetModify(v time.Time)`

SetModify sets Modify field to given value.

### HasModify

`func (o *DfsTrashFile) HasModify() bool`

HasModify returns a boolean if a field has been set.

### GetName

`func (o *DfsTrashFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsTrashFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsTrashFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsTrashFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalName

`func (o *DfsTrashFile) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *DfsTrashFile) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *DfsTrashFile) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *DfsTrashFile) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetOwner

`func (o *DfsTrashFile) GetOwner() int64`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DfsTrashFile) GetOwnerOk() (*int64, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DfsTrashFile) SetOwner(v int64)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DfsTrashFile) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParent

`func (o *DfsTrashFile) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *DfsTrashFile) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *DfsTrashFile) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *DfsTrashFile) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSize

`func (o *DfsTrashFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DfsTrashFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DfsTrashFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DfsTrashFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTrashPath

`func (o *DfsTrashFile) GetTrashPath() string`

GetTrashPath returns the TrashPath field if non-nil, zero value otherwise.

### GetTrashPathOk

`func (o *DfsTrashFile) GetTrashPathOk() (*string, bool)`

GetTrashPathOk returns a tuple with the TrashPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashPath

`func (o *DfsTrashFile) SetTrashPath(v string)`

SetTrashPath sets TrashPath field to given value.

### HasTrashPath

`func (o *DfsTrashFile) HasTrashPath() bool`

HasTrashPath returns a boolean if a field has been set.

### GetType

`func (o *DfsTrashFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsTrashFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsTrashFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsTrashFile) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


