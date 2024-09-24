# DfsFileDeleteReqFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsRootfsId** | **int64** |  | 
**Path** | **string** |  | 
**PrivilegedToken** | **string** |  | 

## Methods

### NewDfsFileDeleteReqFile

`func NewDfsFileDeleteReqFile(dfsRootfsId int64, path string, privilegedToken string, ) *DfsFileDeleteReqFile`

NewDfsFileDeleteReqFile instantiates a new DfsFileDeleteReqFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFileDeleteReqFileWithDefaults

`func NewDfsFileDeleteReqFileWithDefaults() *DfsFileDeleteReqFile`

NewDfsFileDeleteReqFileWithDefaults instantiates a new DfsFileDeleteReqFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsRootfsId

`func (o *DfsFileDeleteReqFile) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsFileDeleteReqFile) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsFileDeleteReqFile) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetPath

`func (o *DfsFileDeleteReqFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsFileDeleteReqFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsFileDeleteReqFile) SetPath(v string)`

SetPath sets Path field to given value.


### GetPrivilegedToken

`func (o *DfsFileDeleteReqFile) GetPrivilegedToken() string`

GetPrivilegedToken returns the PrivilegedToken field if non-nil, zero value otherwise.

### GetPrivilegedTokenOk

`func (o *DfsFileDeleteReqFile) GetPrivilegedTokenOk() (*string, bool)`

GetPrivilegedTokenOk returns a tuple with the PrivilegedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedToken

`func (o *DfsFileDeleteReqFile) SetPrivilegedToken(v string)`

SetPrivilegedToken sets PrivilegedToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


