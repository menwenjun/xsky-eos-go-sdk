# DfsFTPShareACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminEnabled** | Pointer to **bool** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**CreateEnabled** | Pointer to **bool** |  | [optional] 
**DeleteEnabled** | Pointer to **bool** |  | [optional] 
**DfsFtpShare** | Pointer to [**DfsFTPShareNestview**](DfsFTPShareNestview.md) |  | [optional] 
**DownloadBandwidth** | Pointer to **int64** |  | [optional] 
**DownloadEnabled** | Pointer to **bool** |  | [optional] 
**FsUser** | Pointer to [**FSUserNestview**](FSUserNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ListEnabled** | Pointer to **bool** |  | [optional] 
**RenameEnabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**UploadBandwidth** | Pointer to **int64** |  | [optional] 
**UploadEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewDfsFTPShareACL

`func NewDfsFTPShareACL() *DfsFTPShareACL`

NewDfsFTPShareACL instantiates a new DfsFTPShareACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFTPShareACLWithDefaults

`func NewDfsFTPShareACLWithDefaults() *DfsFTPShareACL`

NewDfsFTPShareACLWithDefaults instantiates a new DfsFTPShareACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminEnabled

`func (o *DfsFTPShareACL) GetAdminEnabled() bool`

GetAdminEnabled returns the AdminEnabled field if non-nil, zero value otherwise.

### GetAdminEnabledOk

`func (o *DfsFTPShareACL) GetAdminEnabledOk() (*bool, bool)`

GetAdminEnabledOk returns a tuple with the AdminEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEnabled

`func (o *DfsFTPShareACL) SetAdminEnabled(v bool)`

SetAdminEnabled sets AdminEnabled field to given value.

### HasAdminEnabled

`func (o *DfsFTPShareACL) HasAdminEnabled() bool`

HasAdminEnabled returns a boolean if a field has been set.

### GetCluster

`func (o *DfsFTPShareACL) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsFTPShareACL) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsFTPShareACL) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsFTPShareACL) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsFTPShareACL) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsFTPShareACL) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsFTPShareACL) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsFTPShareACL) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetCreateEnabled

`func (o *DfsFTPShareACL) GetCreateEnabled() bool`

GetCreateEnabled returns the CreateEnabled field if non-nil, zero value otherwise.

### GetCreateEnabledOk

`func (o *DfsFTPShareACL) GetCreateEnabledOk() (*bool, bool)`

GetCreateEnabledOk returns a tuple with the CreateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnabled

`func (o *DfsFTPShareACL) SetCreateEnabled(v bool)`

SetCreateEnabled sets CreateEnabled field to given value.

### HasCreateEnabled

`func (o *DfsFTPShareACL) HasCreateEnabled() bool`

HasCreateEnabled returns a boolean if a field has been set.

### GetDeleteEnabled

`func (o *DfsFTPShareACL) GetDeleteEnabled() bool`

GetDeleteEnabled returns the DeleteEnabled field if non-nil, zero value otherwise.

### GetDeleteEnabledOk

`func (o *DfsFTPShareACL) GetDeleteEnabledOk() (*bool, bool)`

GetDeleteEnabledOk returns a tuple with the DeleteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteEnabled

`func (o *DfsFTPShareACL) SetDeleteEnabled(v bool)`

SetDeleteEnabled sets DeleteEnabled field to given value.

### HasDeleteEnabled

`func (o *DfsFTPShareACL) HasDeleteEnabled() bool`

HasDeleteEnabled returns a boolean if a field has been set.

### GetDfsFtpShare

`func (o *DfsFTPShareACL) GetDfsFtpShare() DfsFTPShareNestview`

GetDfsFtpShare returns the DfsFtpShare field if non-nil, zero value otherwise.

### GetDfsFtpShareOk

`func (o *DfsFTPShareACL) GetDfsFtpShareOk() (*DfsFTPShareNestview, bool)`

GetDfsFtpShareOk returns a tuple with the DfsFtpShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsFtpShare

`func (o *DfsFTPShareACL) SetDfsFtpShare(v DfsFTPShareNestview)`

SetDfsFtpShare sets DfsFtpShare field to given value.

### HasDfsFtpShare

`func (o *DfsFTPShareACL) HasDfsFtpShare() bool`

HasDfsFtpShare returns a boolean if a field has been set.

### GetDownloadBandwidth

`func (o *DfsFTPShareACL) GetDownloadBandwidth() int64`

GetDownloadBandwidth returns the DownloadBandwidth field if non-nil, zero value otherwise.

### GetDownloadBandwidthOk

`func (o *DfsFTPShareACL) GetDownloadBandwidthOk() (*int64, bool)`

GetDownloadBandwidthOk returns a tuple with the DownloadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadBandwidth

`func (o *DfsFTPShareACL) SetDownloadBandwidth(v int64)`

SetDownloadBandwidth sets DownloadBandwidth field to given value.

### HasDownloadBandwidth

`func (o *DfsFTPShareACL) HasDownloadBandwidth() bool`

HasDownloadBandwidth returns a boolean if a field has been set.

### GetDownloadEnabled

`func (o *DfsFTPShareACL) GetDownloadEnabled() bool`

GetDownloadEnabled returns the DownloadEnabled field if non-nil, zero value otherwise.

### GetDownloadEnabledOk

`func (o *DfsFTPShareACL) GetDownloadEnabledOk() (*bool, bool)`

GetDownloadEnabledOk returns a tuple with the DownloadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadEnabled

`func (o *DfsFTPShareACL) SetDownloadEnabled(v bool)`

SetDownloadEnabled sets DownloadEnabled field to given value.

### HasDownloadEnabled

`func (o *DfsFTPShareACL) HasDownloadEnabled() bool`

HasDownloadEnabled returns a boolean if a field has been set.

### GetFsUser

`func (o *DfsFTPShareACL) GetFsUser() FSUserNestview`

GetFsUser returns the FsUser field if non-nil, zero value otherwise.

### GetFsUserOk

`func (o *DfsFTPShareACL) GetFsUserOk() (*FSUserNestview, bool)`

GetFsUserOk returns a tuple with the FsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUser

`func (o *DfsFTPShareACL) SetFsUser(v FSUserNestview)`

SetFsUser sets FsUser field to given value.

### HasFsUser

`func (o *DfsFTPShareACL) HasFsUser() bool`

HasFsUser returns a boolean if a field has been set.

### GetId

`func (o *DfsFTPShareACL) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsFTPShareACL) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsFTPShareACL) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsFTPShareACL) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListEnabled

`func (o *DfsFTPShareACL) GetListEnabled() bool`

GetListEnabled returns the ListEnabled field if non-nil, zero value otherwise.

### GetListEnabledOk

`func (o *DfsFTPShareACL) GetListEnabledOk() (*bool, bool)`

GetListEnabledOk returns a tuple with the ListEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListEnabled

`func (o *DfsFTPShareACL) SetListEnabled(v bool)`

SetListEnabled sets ListEnabled field to given value.

### HasListEnabled

`func (o *DfsFTPShareACL) HasListEnabled() bool`

HasListEnabled returns a boolean if a field has been set.

### GetRenameEnabled

`func (o *DfsFTPShareACL) GetRenameEnabled() bool`

GetRenameEnabled returns the RenameEnabled field if non-nil, zero value otherwise.

### GetRenameEnabledOk

`func (o *DfsFTPShareACL) GetRenameEnabledOk() (*bool, bool)`

GetRenameEnabledOk returns a tuple with the RenameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameEnabled

`func (o *DfsFTPShareACL) SetRenameEnabled(v bool)`

SetRenameEnabled sets RenameEnabled field to given value.

### HasRenameEnabled

`func (o *DfsFTPShareACL) HasRenameEnabled() bool`

HasRenameEnabled returns a boolean if a field has been set.

### GetType

`func (o *DfsFTPShareACL) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsFTPShareACL) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsFTPShareACL) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsFTPShareACL) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsFTPShareACL) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsFTPShareACL) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsFTPShareACL) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsFTPShareACL) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUploadBandwidth

`func (o *DfsFTPShareACL) GetUploadBandwidth() int64`

GetUploadBandwidth returns the UploadBandwidth field if non-nil, zero value otherwise.

### GetUploadBandwidthOk

`func (o *DfsFTPShareACL) GetUploadBandwidthOk() (*int64, bool)`

GetUploadBandwidthOk returns a tuple with the UploadBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadBandwidth

`func (o *DfsFTPShareACL) SetUploadBandwidth(v int64)`

SetUploadBandwidth sets UploadBandwidth field to given value.

### HasUploadBandwidth

`func (o *DfsFTPShareACL) HasUploadBandwidth() bool`

HasUploadBandwidth returns a boolean if a field has been set.

### GetUploadEnabled

`func (o *DfsFTPShareACL) GetUploadEnabled() bool`

GetUploadEnabled returns the UploadEnabled field if non-nil, zero value otherwise.

### GetUploadEnabledOk

`func (o *DfsFTPShareACL) GetUploadEnabledOk() (*bool, bool)`

GetUploadEnabledOk returns a tuple with the UploadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadEnabled

`func (o *DfsFTPShareACL) SetUploadEnabled(v bool)`

SetUploadEnabled sets UploadEnabled field to given value.

### HasUploadEnabled

`func (o *DfsFTPShareACL) HasUploadEnabled() bool`

HasUploadEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


