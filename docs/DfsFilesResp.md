# DfsFilesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsFiles** | Pointer to [**[]DfsFile**](DfsFile.md) | dfs file list | [optional] 
**Eof** | Pointer to **bool** | for ls directory, set to true when reach end | [optional] 

## Methods

### NewDfsFilesResp

`func NewDfsFilesResp() *DfsFilesResp`

NewDfsFilesResp instantiates a new DfsFilesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFilesRespWithDefaults

`func NewDfsFilesRespWithDefaults() *DfsFilesResp`

NewDfsFilesRespWithDefaults instantiates a new DfsFilesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsFiles

`func (o *DfsFilesResp) GetDfsFiles() []DfsFile`

GetDfsFiles returns the DfsFiles field if non-nil, zero value otherwise.

### GetDfsFilesOk

`func (o *DfsFilesResp) GetDfsFilesOk() (*[]DfsFile, bool)`

GetDfsFilesOk returns a tuple with the DfsFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsFiles

`func (o *DfsFilesResp) SetDfsFiles(v []DfsFile)`

SetDfsFiles sets DfsFiles field to given value.

### HasDfsFiles

`func (o *DfsFilesResp) HasDfsFiles() bool`

HasDfsFiles returns a boolean if a field has been set.

### GetEof

`func (o *DfsFilesResp) GetEof() bool`

GetEof returns the Eof field if non-nil, zero value otherwise.

### GetEofOk

`func (o *DfsFilesResp) GetEofOk() (*bool, bool)`

GetEofOk returns a tuple with the Eof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEof

`func (o *DfsFilesResp) SetEof(v bool)`

SetEof sets Eof field to given value.

### HasEof

`func (o *DfsFilesResp) HasEof() bool`

HasEof returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


