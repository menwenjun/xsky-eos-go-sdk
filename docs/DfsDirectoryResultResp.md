# DfsDirectoryResultResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directory** | [**DfsFile**](DfsFile.md) |  | 
**Result** | **string** | operation result | 

## Methods

### NewDfsDirectoryResultResp

`func NewDfsDirectoryResultResp(directory DfsFile, result string, ) *DfsDirectoryResultResp`

NewDfsDirectoryResultResp instantiates a new DfsDirectoryResultResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsDirectoryResultRespWithDefaults

`func NewDfsDirectoryResultRespWithDefaults() *DfsDirectoryResultResp`

NewDfsDirectoryResultRespWithDefaults instantiates a new DfsDirectoryResultResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectory

`func (o *DfsDirectoryResultResp) GetDirectory() DfsFile`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *DfsDirectoryResultResp) GetDirectoryOk() (*DfsFile, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *DfsDirectoryResultResp) SetDirectory(v DfsFile)`

SetDirectory sets Directory field to given value.


### GetResult

`func (o *DfsDirectoryResultResp) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DfsDirectoryResultResp) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DfsDirectoryResultResp) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


