# DfsFTPShareUpdateACLReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminEnabled** | Pointer to **bool** | enable admin user permission | [optional] 
**CreateEnabled** | Pointer to **bool** | enable creating files | [optional] 
**DeleteEnabled** | Pointer to **bool** | enable deleting files | [optional] 
**DownloadBandwidth** | Pointer to **int64** | max bandwidth of downloading | [optional] 
**DownloadEnabled** | Pointer to **bool** | enable downloading files | [optional] 
**Id** | Pointer to **int64** | id of user group | [optional] 
**ListEnabled** | Pointer to **bool** | enable listing files | [optional] 
**RenameEnabled** | Pointer to **bool** | enable renaming files | [optional] 
**UploadBandwidth** | Pointer to **int64** | max bandwidth of uploading | [optional] 
**UploadEnabled** | Pointer to **bool** | enable uploading files | [optional] 

## Methods

### NewDfsFTPShareUpdateACLReq

`func NewDfsFTPShareUpdateACLReq() *DfsFTPShareUpdateACLReq`

NewDfsFTPShareUpdateACLReq instantiates a new DfsFTPShareUpdateACLReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFTPShareUpdateACLReqWithDefaults

`func NewDfsFTPShareUpdateACLReqWithDefaults() *DfsFTPShareUpdateACLReq`

NewDfsFTPShareUpdateACLReqWithDefaults instantiates a new DfsFTPShareUpdateACLReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminEnabled

`func (o *DfsFTPShareUpdateACLReq) GetAdminEnabled() bool`

GetAdminEnabled returns the AdminEnabled field if non-nil, zero value otherwise.

### GetAdminEnabledOk

`func (o *DfsFTPShareUpdateACLReq) GetAdminEnabledOk() (*bool, bool)`

GetAdminEnabledOk returns a tuple with the AdminEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEnabled

`func (o *DfsFTPShareUpdateACLReq) SetAdminEnabled(v bool)`

SetAdminEnabled sets AdminEnabled field to given value.

### HasAdminEnabled

`func (o *DfsFTPShareUpdateACLReq) HasAdminEnabled() bool`

HasAdminEnabled returns a boolean if a field has been set.

### GetCreateEnabled

`func (o *DfsFTPShareUpdateACLReq) GetCreateEnabled() bool`

GetCreateEnabled returns the CreateEnabled field if non-nil, zero value otherwise.

### GetCreateEnabledOk

`func (o *DfsFTPShareUpdateACLReq) GetCreateEnabledOk() (*bool, bool)`

GetCreateEnabledOk returns a tuple with the CreateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnabled

`func (o *DfsFTPShareUpdateACLReq) SetCreateEnabled(v bool)`

SetCreateEnabled sets CreateEnabled field to given value.

### HasCreateEnabled

`func (o *DfsFTPShareUpdateACLReq) HasCreateEnabled() bool`

HasCreateEnabled returns a boolean if a field has been set.

### GetDeleteEnabled

`func (o *DfsFTPShareUpdateACLReq) GetDeleteEnabled() bool`

GetDeleteEnabled returns the DeleteEnabled field if non-nil, zero value otherwise.

### GetDeleteEnabledOk

`func (o *DfsFTPShareUpdateACLReq) GetDeleteEnabledOk() (*bool, bool)`

GetDeleteEnabledOk returns a tuple with the DeleteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteEnabled

`func (o *DfsFTPShareUpdateACLReq) SetDeleteEnabled(v bool)`

SetDeleteEnabled sets DeleteEnabled field to given value.

### HasDeleteEnabled

`func (o *DfsFTPShareUpdateACLReq) HasDeleteEnabled() bool`

HasDeleteEnabled returns a boolean if a field has been set.

### GetDownloadBandwidth

`func (o *DfsFTPShareUpdateACLReq) GetDownloadBandwidth() int64`

GetDownloadBandwidth returns the DownloadBandwidth field if non-nil, zero value otherwise.

### GetDownloadBandwidthOk

`func (o *DfsFTPShareUpdateACLReq) GetDownloadBandwidthOk() (*int64, bool)`

GetDownloadBandwidthOk returns a tuple with the DownloadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadBandwidth

`func (o *DfsFTPShareUpdateACLReq) SetDownloadBandwidth(v int64)`

SetDownloadBandwidth sets DownloadBandwidth field to given value.

### HasDownloadBandwidth

`func (o *DfsFTPShareUpdateACLReq) HasDownloadBandwidth() bool`

HasDownloadBandwidth returns a boolean if a field has been set.

### GetDownloadEnabled

`func (o *DfsFTPShareUpdateACLReq) GetDownloadEnabled() bool`

GetDownloadEnabled returns the DownloadEnabled field if non-nil, zero value otherwise.

### GetDownloadEnabledOk

`func (o *DfsFTPShareUpdateACLReq) GetDownloadEnabledOk() (*bool, bool)`

GetDownloadEnabledOk returns a tuple with the DownloadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadEnabled

`func (o *DfsFTPShareUpdateACLReq) SetDownloadEnabled(v bool)`

SetDownloadEnabled sets DownloadEnabled field to given value.

### HasDownloadEnabled

`func (o *DfsFTPShareUpdateACLReq) HasDownloadEnabled() bool`

HasDownloadEnabled returns a boolean if a field has been set.

### GetId

`func (o *DfsFTPShareUpdateACLReq) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsFTPShareUpdateACLReq) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsFTPShareUpdateACLReq) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsFTPShareUpdateACLReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListEnabled

`func (o *DfsFTPShareUpdateACLReq) GetListEnabled() bool`

GetListEnabled returns the ListEnabled field if non-nil, zero value otherwise.

### GetListEnabledOk

`func (o *DfsFTPShareUpdateACLReq) GetListEnabledOk() (*bool, bool)`

GetListEnabledOk returns a tuple with the ListEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListEnabled

`func (o *DfsFTPShareUpdateACLReq) SetListEnabled(v bool)`

SetListEnabled sets ListEnabled field to given value.

### HasListEnabled

`func (o *DfsFTPShareUpdateACLReq) HasListEnabled() bool`

HasListEnabled returns a boolean if a field has been set.

### GetRenameEnabled

`func (o *DfsFTPShareUpdateACLReq) GetRenameEnabled() bool`

GetRenameEnabled returns the RenameEnabled field if non-nil, zero value otherwise.

### GetRenameEnabledOk

`func (o *DfsFTPShareUpdateACLReq) GetRenameEnabledOk() (*bool, bool)`

GetRenameEnabledOk returns a tuple with the RenameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameEnabled

`func (o *DfsFTPShareUpdateACLReq) SetRenameEnabled(v bool)`

SetRenameEnabled sets RenameEnabled field to given value.

### HasRenameEnabled

`func (o *DfsFTPShareUpdateACLReq) HasRenameEnabled() bool`

HasRenameEnabled returns a boolean if a field has been set.

### GetUploadBandwidth

`func (o *DfsFTPShareUpdateACLReq) GetUploadBandwidth() int64`

GetUploadBandwidth returns the UploadBandwidth field if non-nil, zero value otherwise.

### GetUploadBandwidthOk

`func (o *DfsFTPShareUpdateACLReq) GetUploadBandwidthOk() (*int64, bool)`

GetUploadBandwidthOk returns a tuple with the UploadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadBandwidth

`func (o *DfsFTPShareUpdateACLReq) SetUploadBandwidth(v int64)`

SetUploadBandwidth sets UploadBandwidth field to given value.

### HasUploadBandwidth

`func (o *DfsFTPShareUpdateACLReq) HasUploadBandwidth() bool`

HasUploadBandwidth returns a boolean if a field has been set.

### GetUploadEnabled

`func (o *DfsFTPShareUpdateACLReq) GetUploadEnabled() bool`

GetUploadEnabled returns the UploadEnabled field if non-nil, zero value otherwise.

### GetUploadEnabledOk

`func (o *DfsFTPShareUpdateACLReq) GetUploadEnabledOk() (*bool, bool)`

GetUploadEnabledOk returns a tuple with the UploadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadEnabled

`func (o *DfsFTPShareUpdateACLReq) SetUploadEnabled(v bool)`

SetUploadEnabled sets UploadEnabled field to given value.

### HasUploadEnabled

`func (o *DfsFTPShareUpdateACLReq) HasUploadEnabled() bool`

HasUploadEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


