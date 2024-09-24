# ChmodReqFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsRootfsId** | **int64** | dfs rootfs id | 
**Mode** | **string** | file mode | 
**Path** | **string** | full path | 

## Methods

### NewChmodReqFile

`func NewChmodReqFile(dfsRootfsId int64, mode string, path string, ) *ChmodReqFile`

NewChmodReqFile instantiates a new ChmodReqFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChmodReqFileWithDefaults

`func NewChmodReqFileWithDefaults() *ChmodReqFile`

NewChmodReqFileWithDefaults instantiates a new ChmodReqFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsRootfsId

`func (o *ChmodReqFile) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *ChmodReqFile) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *ChmodReqFile) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetMode

`func (o *ChmodReqFile) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ChmodReqFile) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ChmodReqFile) SetMode(v string)`

SetMode sets Mode field to given value.


### GetPath

`func (o *ChmodReqFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ChmodReqFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ChmodReqFile) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


