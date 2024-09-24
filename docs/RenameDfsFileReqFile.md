# RenameDfsFileReqFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsRootfsId** | **int64** | dfs rootfs id | 
**DstPath** | **string** | dst path | 
**SrcPath** | **string** | src path | 

## Methods

### NewRenameDfsFileReqFile

`func NewRenameDfsFileReqFile(dfsRootfsId int64, dstPath string, srcPath string, ) *RenameDfsFileReqFile`

NewRenameDfsFileReqFile instantiates a new RenameDfsFileReqFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameDfsFileReqFileWithDefaults

`func NewRenameDfsFileReqFileWithDefaults() *RenameDfsFileReqFile`

NewRenameDfsFileReqFileWithDefaults instantiates a new RenameDfsFileReqFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsRootfsId

`func (o *RenameDfsFileReqFile) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *RenameDfsFileReqFile) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *RenameDfsFileReqFile) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetDstPath

`func (o *RenameDfsFileReqFile) GetDstPath() string`

GetDstPath returns the DstPath field if non-nil, zero value otherwise.

### GetDstPathOk

`func (o *RenameDfsFileReqFile) GetDstPathOk() (*string, bool)`

GetDstPathOk returns a tuple with the DstPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPath

`func (o *RenameDfsFileReqFile) SetDstPath(v string)`

SetDstPath sets DstPath field to given value.


### GetSrcPath

`func (o *RenameDfsFileReqFile) GetSrcPath() string`

GetSrcPath returns the SrcPath field if non-nil, zero value otherwise.

### GetSrcPathOk

`func (o *RenameDfsFileReqFile) GetSrcPathOk() (*string, bool)`

GetSrcPathOk returns a tuple with the SrcPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPath

`func (o *RenameDfsFileReqFile) SetSrcPath(v string)`

SetSrcPath sets SrcPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


