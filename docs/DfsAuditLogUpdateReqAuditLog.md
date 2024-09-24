# DfsAuditLogUpdateReqAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | actions need to audit | [optional] 
**DfsRootfsId** | **int64** | rootfs id | 
**LogPath** | Pointer to **string** | audit log path | [optional] 
**Size** | Pointer to **int64** | max size of audit log | [optional] 

## Methods

### NewDfsAuditLogUpdateReqAuditLog

`func NewDfsAuditLogUpdateReqAuditLog(dfsRootfsId int64, ) *DfsAuditLogUpdateReqAuditLog`

NewDfsAuditLogUpdateReqAuditLog instantiates a new DfsAuditLogUpdateReqAuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsAuditLogUpdateReqAuditLogWithDefaults

`func NewDfsAuditLogUpdateReqAuditLogWithDefaults() *DfsAuditLogUpdateReqAuditLog`

NewDfsAuditLogUpdateReqAuditLogWithDefaults instantiates a new DfsAuditLogUpdateReqAuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *DfsAuditLogUpdateReqAuditLog) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DfsAuditLogUpdateReqAuditLog) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DfsAuditLogUpdateReqAuditLog) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *DfsAuditLogUpdateReqAuditLog) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDfsRootfsId

`func (o *DfsAuditLogUpdateReqAuditLog) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsAuditLogUpdateReqAuditLog) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsAuditLogUpdateReqAuditLog) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetLogPath

`func (o *DfsAuditLogUpdateReqAuditLog) GetLogPath() string`

GetLogPath returns the LogPath field if non-nil, zero value otherwise.

### GetLogPathOk

`func (o *DfsAuditLogUpdateReqAuditLog) GetLogPathOk() (*string, bool)`

GetLogPathOk returns a tuple with the LogPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogPath

`func (o *DfsAuditLogUpdateReqAuditLog) SetLogPath(v string)`

SetLogPath sets LogPath field to given value.

### HasLogPath

`func (o *DfsAuditLogUpdateReqAuditLog) HasLogPath() bool`

HasLogPath returns a boolean if a field has been set.

### GetSize

`func (o *DfsAuditLogUpdateReqAuditLog) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DfsAuditLogUpdateReqAuditLog) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DfsAuditLogUpdateReqAuditLog) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DfsAuditLogUpdateReqAuditLog) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


