# DfsFileBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to **int64** | sub file count when it is a directory | [optional] 
**Inode** | Pointer to **int64** | inode count | [optional] 
**Name** | Pointer to **string** | file name | [optional] 
**Path** | Pointer to **string** | full path | [optional] 
**Size** | Pointer to **int64** | file size | [optional] 
**Type** | Pointer to **string** | file type | [optional] 

## Methods

### NewDfsFileBase

`func NewDfsFileBase() *DfsFileBase`

NewDfsFileBase instantiates a new DfsFileBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFileBaseWithDefaults

`func NewDfsFileBaseWithDefaults() *DfsFileBase`

NewDfsFileBaseWithDefaults instantiates a new DfsFileBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *DfsFileBase) GetFiles() int64`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *DfsFileBase) GetFilesOk() (*int64, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *DfsFileBase) SetFiles(v int64)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *DfsFileBase) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetInode

`func (o *DfsFileBase) GetInode() int64`

GetInode returns the Inode field if non-nil, zero value otherwise.

### GetInodeOk

`func (o *DfsFileBase) GetInodeOk() (*int64, bool)`

GetInodeOk returns a tuple with the Inode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInode

`func (o *DfsFileBase) SetInode(v int64)`

SetInode sets Inode field to given value.

### HasInode

`func (o *DfsFileBase) HasInode() bool`

HasInode returns a boolean if a field has been set.

### GetName

`func (o *DfsFileBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsFileBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsFileBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsFileBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *DfsFileBase) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsFileBase) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsFileBase) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsFileBase) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *DfsFileBase) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DfsFileBase) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DfsFileBase) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DfsFileBase) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *DfsFileBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsFileBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsFileBase) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsFileBase) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


