# DfsTrashFilesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsTrashFiles** | Pointer to [**[]DfsTrashFile**](DfsTrashFile.md) | files in dfs trash | [optional] 
**Eof** | Pointer to **bool** | for ls directory, set to true when reach end | [optional] 

## Methods

### NewDfsTrashFilesResp

`func NewDfsTrashFilesResp() *DfsTrashFilesResp`

NewDfsTrashFilesResp instantiates a new DfsTrashFilesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsTrashFilesRespWithDefaults

`func NewDfsTrashFilesRespWithDefaults() *DfsTrashFilesResp`

NewDfsTrashFilesRespWithDefaults instantiates a new DfsTrashFilesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsTrashFiles

`func (o *DfsTrashFilesResp) GetDfsTrashFiles() []DfsTrashFile`

GetDfsTrashFiles returns the DfsTrashFiles field if non-nil, zero value otherwise.

### GetDfsTrashFilesOk

`func (o *DfsTrashFilesResp) GetDfsTrashFilesOk() (*[]DfsTrashFile, bool)`

GetDfsTrashFilesOk returns a tuple with the DfsTrashFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsTrashFiles

`func (o *DfsTrashFilesResp) SetDfsTrashFiles(v []DfsTrashFile)`

SetDfsTrashFiles sets DfsTrashFiles field to given value.

### HasDfsTrashFiles

`func (o *DfsTrashFilesResp) HasDfsTrashFiles() bool`

HasDfsTrashFiles returns a boolean if a field has been set.

### GetEof

`func (o *DfsTrashFilesResp) GetEof() bool`

GetEof returns the Eof field if non-nil, zero value otherwise.

### GetEofOk

`func (o *DfsTrashFilesResp) GetEofOk() (*bool, bool)`

GetEofOk returns a tuple with the Eof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEof

`func (o *DfsTrashFilesResp) SetEof(v bool)`

SetEof sets Eof field to given value.

### HasEof

`func (o *DfsTrashFilesResp) HasEof() bool`

HasEof returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


