# DfsFTPShareACLReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminEnabled** | Pointer to **bool** | enable admin user permission | [optional] 
**CreateEnabled** | Pointer to **bool** | enable creating files | [optional] 
**DeleteEnabled** | Pointer to **bool** | enable deleting files | [optional] 
**DownloadBandwidth** | Pointer to **int64** | max bandwidth of downloading | [optional] 
**DownloadEnabled** | Pointer to **bool** | enable downloading files | [optional] 
**FsUserId** | Pointer to **int64** | id of user | [optional] 
**Id** | Pointer to **int64** | id of user group | [optional] 
**ListEnabled** | Pointer to **bool** | enable listing files | [optional] 
**RenameEnabled** | Pointer to **bool** | enable renaming files | [optional] 
**Type** | Pointer to **string** | type of share acl | [optional] 
**UploadBandwidth** | Pointer to **int64** | max bandwidth of uploading | [optional] 
**UploadEnabled** | Pointer to **bool** | enable uploading files | [optional] 

## Methods

### NewDfsFTPShareACLReq

`func NewDfsFTPShareACLReq() *DfsFTPShareACLReq`

NewDfsFTPShareACLReq instantiates a new DfsFTPShareACLReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFTPShareACLReqWithDefaults

`func NewDfsFTPShareACLReqWithDefaults() *DfsFTPShareACLReq`

NewDfsFTPShareACLReqWithDefaults instantiates a new DfsFTPShareACLReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminEnabled

`func (o *DfsFTPShareACLReq) GetAdminEnabled() bool`

GetAdminEnabled returns the AdminEnabled field if non-nil, zero value otherwise.

### GetAdminEnabledOk

`func (o *DfsFTPShareACLReq) GetAdminEnabledOk() (*bool, bool)`

GetAdminEnabledOk returns a tuple with the AdminEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEnabled

`func (o *DfsFTPShareACLReq) SetAdminEnabled(v bool)`

SetAdminEnabled sets AdminEnabled field to given value.

### HasAdminEnabled

`func (o *DfsFTPShareACLReq) HasAdminEnabled() bool`

HasAdminEnabled returns a boolean if a field has been set.

### GetCreateEnabled

`func (o *DfsFTPShareACLReq) GetCreateEnabled() bool`

GetCreateEnabled returns the CreateEnabled field if non-nil, zero value otherwise.

### GetCreateEnabledOk

`func (o *DfsFTPShareACLReq) GetCreateEnabledOk() (*bool, bool)`

GetCreateEnabledOk returns a tuple with the CreateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnabled

`func (o *DfsFTPShareACLReq) SetCreateEnabled(v bool)`

SetCreateEnabled sets CreateEnabled field to given value.

### HasCreateEnabled

`func (o *DfsFTPShareACLReq) HasCreateEnabled() bool`

HasCreateEnabled returns a boolean if a field has been set.

### GetDeleteEnabled

`func (o *DfsFTPShareACLReq) GetDeleteEnabled() bool`

GetDeleteEnabled returns the DeleteEnabled field if non-nil, zero value otherwise.

### GetDeleteEnabledOk

`func (o *DfsFTPShareACLReq) GetDeleteEnabledOk() (*bool, bool)`

GetDeleteEnabledOk returns a tuple with the DeleteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteEnabled

`func (o *DfsFTPShareACLReq) SetDeleteEnabled(v bool)`

SetDeleteEnabled sets DeleteEnabled field to given value.

### HasDeleteEnabled

`func (o *DfsFTPShareACLReq) HasDeleteEnabled() bool`

HasDeleteEnabled returns a boolean if a field has been set.

### GetDownloadBandwidth

`func (o *DfsFTPShareACLReq) GetDownloadBandwidth() int64`

GetDownloadBandwidth returns the DownloadBandwidth field if non-nil, zero value otherwise.

### GetDownloadBandwidthOk

`func (o *DfsFTPShareACLReq) GetDownloadBandwidthOk() (*int64, bool)`

GetDownloadBandwidthOk returns a tuple with the DownloadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadBandwidth

`func (o *DfsFTPShareACLReq) SetDownloadBandwidth(v int64)`

SetDownloadBandwidth sets DownloadBandwidth field to given value.

### HasDownloadBandwidth

`func (o *DfsFTPShareACLReq) HasDownloadBandwidth() bool`

HasDownloadBandwidth returns a boolean if a field has been set.

### GetDownloadEnabled

`func (o *DfsFTPShareACLReq) GetDownloadEnabled() bool`

GetDownloadEnabled returns the DownloadEnabled field if non-nil, zero value otherwise.

### GetDownloadEnabledOk

`func (o *DfsFTPShareACLReq) GetDownloadEnabledOk() (*bool, bool)`

GetDownloadEnabledOk returns a tuple with the DownloadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadEnabled

`func (o *DfsFTPShareACLReq) SetDownloadEnabled(v bool)`

SetDownloadEnabled sets DownloadEnabled field to given value.

### HasDownloadEnabled

`func (o *DfsFTPShareACLReq) HasDownloadEnabled() bool`

HasDownloadEnabled returns a boolean if a field has been set.

### GetFsUserId

`func (o *DfsFTPShareACLReq) GetFsUserId() int64`

GetFsUserId returns the FsUserId field if non-nil, zero value otherwise.

### GetFsUserIdOk

`func (o *DfsFTPShareACLReq) GetFsUserIdOk() (*int64, bool)`

GetFsUserIdOk returns a tuple with the FsUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserId

`func (o *DfsFTPShareACLReq) SetFsUserId(v int64)`

SetFsUserId sets FsUserId field to given value.

### HasFsUserId

`func (o *DfsFTPShareACLReq) HasFsUserId() bool`

HasFsUserId returns a boolean if a field has been set.

### GetId

`func (o *DfsFTPShareACLReq) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsFTPShareACLReq) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsFTPShareACLReq) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsFTPShareACLReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListEnabled

`func (o *DfsFTPShareACLReq) GetListEnabled() bool`

GetListEnabled returns the ListEnabled field if non-nil, zero value otherwise.

### GetListEnabledOk

`func (o *DfsFTPShareACLReq) GetListEnabledOk() (*bool, bool)`

GetListEnabledOk returns a tuple with the ListEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListEnabled

`func (o *DfsFTPShareACLReq) SetListEnabled(v bool)`

SetListEnabled sets ListEnabled field to given value.

### HasListEnabled

`func (o *DfsFTPShareACLReq) HasListEnabled() bool`

HasListEnabled returns a boolean if a field has been set.

### GetRenameEnabled

`func (o *DfsFTPShareACLReq) GetRenameEnabled() bool`

GetRenameEnabled returns the RenameEnabled field if non-nil, zero value otherwise.

### GetRenameEnabledOk

`func (o *DfsFTPShareACLReq) GetRenameEnabledOk() (*bool, bool)`

GetRenameEnabledOk returns a tuple with the RenameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameEnabled

`func (o *DfsFTPShareACLReq) SetRenameEnabled(v bool)`

SetRenameEnabled sets RenameEnabled field to given value.

### HasRenameEnabled

`func (o *DfsFTPShareACLReq) HasRenameEnabled() bool`

HasRenameEnabled returns a boolean if a field has been set.

### GetType

`func (o *DfsFTPShareACLReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsFTPShareACLReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsFTPShareACLReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsFTPShareACLReq) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUploadBandwidth

`func (o *DfsFTPShareACLReq) GetUploadBandwidth() int64`

GetUploadBandwidth returns the UploadBandwidth field if non-nil, zero value otherwise.

### GetUploadBandwidthOk

`func (o *DfsFTPShareACLReq) GetUploadBandwidthOk() (*int64, bool)`

GetUploadBandwidthOk returns a tuple with the UploadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadBandwidth

`func (o *DfsFTPShareACLReq) SetUploadBandwidth(v int64)`

SetUploadBandwidth sets UploadBandwidth field to given value.

### HasUploadBandwidth

`func (o *DfsFTPShareACLReq) HasUploadBandwidth() bool`

HasUploadBandwidth returns a boolean if a field has been set.

### GetUploadEnabled

`func (o *DfsFTPShareACLReq) GetUploadEnabled() bool`

GetUploadEnabled returns the UploadEnabled field if non-nil, zero value otherwise.

### GetUploadEnabledOk

`func (o *DfsFTPShareACLReq) GetUploadEnabledOk() (*bool, bool)`

GetUploadEnabledOk returns a tuple with the UploadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadEnabled

`func (o *DfsFTPShareACLReq) SetUploadEnabled(v bool)`

SetUploadEnabled sets UploadEnabled field to given value.

### HasUploadEnabled

`func (o *DfsFTPShareACLReq) HasUploadEnabled() bool`

HasUploadEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


