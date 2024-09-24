# DfsDirectoriesRespDirectories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsRootfs** | [**DfsRootfs**](DfsRootfs.md) |  | 
**DirectoryResults** | [**[]DfsDirectoryResultRespOld**](DfsDirectoryResultRespOld.md) | dfs directories operation results | 

## Methods

### NewDfsDirectoriesRespDirectories

`func NewDfsDirectoriesRespDirectories(dfsRootfs DfsRootfs, directoryResults []DfsDirectoryResultRespOld, ) *DfsDirectoriesRespDirectories`

NewDfsDirectoriesRespDirectories instantiates a new DfsDirectoriesRespDirectories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsDirectoriesRespDirectoriesWithDefaults

`func NewDfsDirectoriesRespDirectoriesWithDefaults() *DfsDirectoriesRespDirectories`

NewDfsDirectoriesRespDirectoriesWithDefaults instantiates a new DfsDirectoriesRespDirectories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsRootfs

`func (o *DfsDirectoriesRespDirectories) GetDfsRootfs() DfsRootfs`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsDirectoriesRespDirectories) GetDfsRootfsOk() (*DfsRootfs, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsDirectoriesRespDirectories) SetDfsRootfs(v DfsRootfs)`

SetDfsRootfs sets DfsRootfs field to given value.


### GetDirectoryResults

`func (o *DfsDirectoriesRespDirectories) GetDirectoryResults() []DfsDirectoryResultRespOld`

GetDirectoryResults returns the DirectoryResults field if non-nil, zero value otherwise.

### GetDirectoryResultsOk

`func (o *DfsDirectoriesRespDirectories) GetDirectoryResultsOk() (*[]DfsDirectoryResultRespOld, bool)`

GetDirectoryResultsOk returns a tuple with the DirectoryResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResults

`func (o *DfsDirectoriesRespDirectories) SetDirectoryResults(v []DfsDirectoryResultRespOld)`

SetDirectoryResults sets DirectoryResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


