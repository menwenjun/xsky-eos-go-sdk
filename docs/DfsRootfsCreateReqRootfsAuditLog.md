# DfsRootfsCreateReqRootfsAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | actions need to audit | [optional] 
**LogPath** | Pointer to **string** | audit log path | [optional] 
**Size** | Pointer to **int64** | max size of audit log | [optional] 

## Methods

### NewDfsRootfsCreateReqRootfsAuditLog

`func NewDfsRootfsCreateReqRootfsAuditLog() *DfsRootfsCreateReqRootfsAuditLog`

NewDfsRootfsCreateReqRootfsAuditLog instantiates a new DfsRootfsCreateReqRootfsAuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsRootfsCreateReqRootfsAuditLogWithDefaults

`func NewDfsRootfsCreateReqRootfsAuditLogWithDefaults() *DfsRootfsCreateReqRootfsAuditLog`

NewDfsRootfsCreateReqRootfsAuditLogWithDefaults instantiates a new DfsRootfsCreateReqRootfsAuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *DfsRootfsCreateReqRootfsAuditLog) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DfsRootfsCreateReqRootfsAuditLog) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DfsRootfsCreateReqRootfsAuditLog) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *DfsRootfsCreateReqRootfsAuditLog) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetLogPath

`func (o *DfsRootfsCreateReqRootfsAuditLog) GetLogPath() string`

GetLogPath returns the LogPath field if non-nil, zero value otherwise.

### GetLogPathOk

`func (o *DfsRootfsCreateReqRootfsAuditLog) GetLogPathOk() (*string, bool)`

GetLogPathOk returns a tuple with the LogPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogPath

`func (o *DfsRootfsCreateReqRootfsAuditLog) SetLogPath(v string)`

SetLogPath sets LogPath field to given value.

### HasLogPath

`func (o *DfsRootfsCreateReqRootfsAuditLog) HasLogPath() bool`

HasLogPath returns a boolean if a field has been set.

### GetSize

`func (o *DfsRootfsCreateReqRootfsAuditLog) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DfsRootfsCreateReqRootfsAuditLog) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DfsRootfsCreateReqRootfsAuditLog) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DfsRootfsCreateReqRootfsAuditLog) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


