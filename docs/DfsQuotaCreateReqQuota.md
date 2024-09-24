# DfsQuotaCreateReqQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdId** | Pointer to **int64** | ad id | [optional] 
**DfsRootfsId** | **int64** | id of dfs rootfs | 
**FilesGraceTime** | Pointer to **int64** | grace time for files soft quota | [optional] 
**FilesHardQuota** | Pointer to **int64** | hard quota of files | [optional] 
**FilesRatio** | Pointer to **float32** | ratio of files | [optional] 
**FilesSoftQuota** | Pointer to **int64** | soft quota of files | [optional] 
**FilesSuggestQuota** | Pointer to **int64** | suggest quota of files | [optional] 
**FsUserGroupId** | Pointer to **int64** | fs user group id | [optional] 
**FsUserId** | Pointer to **int64** | fs user id | [optional] 
**LdapId** | Pointer to **int64** | ldap id | [optional] 
**Path** | **string** | path of quota | 
**SizeGraceTime** | Pointer to **int64** | grace time for size soft quota | [optional] 
**SizeHardQuota** | Pointer to **int64** | hard quota of size | [optional] 
**SizeRatio** | Pointer to **float32** | ratio of size | [optional] 
**SizeSoftQuota** | Pointer to **int64** | soft quota of size | [optional] 
**SizeSuggestQuota** | Pointer to **int64** | suggest quota of size | [optional] 
**Type** | Pointer to **string** | type of dfs quota | [optional] 
**UserGroupName** | Pointer to **string** | user group name | [optional] 
**UserName** | Pointer to **string** | user name | [optional] 

## Methods

### NewDfsQuotaCreateReqQuota

`func NewDfsQuotaCreateReqQuota(dfsRootfsId int64, path string, ) *DfsQuotaCreateReqQuota`

NewDfsQuotaCreateReqQuota instantiates a new DfsQuotaCreateReqQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQuotaCreateReqQuotaWithDefaults

`func NewDfsQuotaCreateReqQuotaWithDefaults() *DfsQuotaCreateReqQuota`

NewDfsQuotaCreateReqQuotaWithDefaults instantiates a new DfsQuotaCreateReqQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdId

`func (o *DfsQuotaCreateReqQuota) GetAdId() int64`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *DfsQuotaCreateReqQuota) GetAdIdOk() (*int64, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *DfsQuotaCreateReqQuota) SetAdId(v int64)`

SetAdId sets AdId field to given value.

### HasAdId

`func (o *DfsQuotaCreateReqQuota) HasAdId() bool`

HasAdId returns a boolean if a field has been set.

### GetDfsRootfsId

`func (o *DfsQuotaCreateReqQuota) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsQuotaCreateReqQuota) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsQuotaCreateReqQuota) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetFilesGraceTime

`func (o *DfsQuotaCreateReqQuota) GetFilesGraceTime() int64`

GetFilesGraceTime returns the FilesGraceTime field if non-nil, zero value otherwise.

### GetFilesGraceTimeOk

`func (o *DfsQuotaCreateReqQuota) GetFilesGraceTimeOk() (*int64, bool)`

GetFilesGraceTimeOk returns a tuple with the FilesGraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesGraceTime

`func (o *DfsQuotaCreateReqQuota) SetFilesGraceTime(v int64)`

SetFilesGraceTime sets FilesGraceTime field to given value.

### HasFilesGraceTime

`func (o *DfsQuotaCreateReqQuota) HasFilesGraceTime() bool`

HasFilesGraceTime returns a boolean if a field has been set.

### GetFilesHardQuota

`func (o *DfsQuotaCreateReqQuota) GetFilesHardQuota() int64`

GetFilesHardQuota returns the FilesHardQuota field if non-nil, zero value otherwise.

### GetFilesHardQuotaOk

`func (o *DfsQuotaCreateReqQuota) GetFilesHardQuotaOk() (*int64, bool)`

GetFilesHardQuotaOk returns a tuple with the FilesHardQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesHardQuota

`func (o *DfsQuotaCreateReqQuota) SetFilesHardQuota(v int64)`

SetFilesHardQuota sets FilesHardQuota field to given value.

### HasFilesHardQuota

`func (o *DfsQuotaCreateReqQuota) HasFilesHardQuota() bool`

HasFilesHardQuota returns a boolean if a field has been set.

### GetFilesRatio

`func (o *DfsQuotaCreateReqQuota) GetFilesRatio() float32`

GetFilesRatio returns the FilesRatio field if non-nil, zero value otherwise.

### GetFilesRatioOk

`func (o *DfsQuotaCreateReqQuota) GetFilesRatioOk() (*float32, bool)`

GetFilesRatioOk returns a tuple with the FilesRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesRatio

`func (o *DfsQuotaCreateReqQuota) SetFilesRatio(v float32)`

SetFilesRatio sets FilesRatio field to given value.

### HasFilesRatio

`func (o *DfsQuotaCreateReqQuota) HasFilesRatio() bool`

HasFilesRatio returns a boolean if a field has been set.

### GetFilesSoftQuota

`func (o *DfsQuotaCreateReqQuota) GetFilesSoftQuota() int64`

GetFilesSoftQuota returns the FilesSoftQuota field if non-nil, zero value otherwise.

### GetFilesSoftQuotaOk

`func (o *DfsQuotaCreateReqQuota) GetFilesSoftQuotaOk() (*int64, bool)`

GetFilesSoftQuotaOk returns a tuple with the FilesSoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSoftQuota

`func (o *DfsQuotaCreateReqQuota) SetFilesSoftQuota(v int64)`

SetFilesSoftQuota sets FilesSoftQuota field to given value.

### HasFilesSoftQuota

`func (o *DfsQuotaCreateReqQuota) HasFilesSoftQuota() bool`

HasFilesSoftQuota returns a boolean if a field has been set.

### GetFilesSuggestQuota

`func (o *DfsQuotaCreateReqQuota) GetFilesSuggestQuota() int64`

GetFilesSuggestQuota returns the FilesSuggestQuota field if non-nil, zero value otherwise.

### GetFilesSuggestQuotaOk

`func (o *DfsQuotaCreateReqQuota) GetFilesSuggestQuotaOk() (*int64, bool)`

GetFilesSuggestQuotaOk returns a tuple with the FilesSuggestQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSuggestQuota

`func (o *DfsQuotaCreateReqQuota) SetFilesSuggestQuota(v int64)`

SetFilesSuggestQuota sets FilesSuggestQuota field to given value.

### HasFilesSuggestQuota

`func (o *DfsQuotaCreateReqQuota) HasFilesSuggestQuota() bool`

HasFilesSuggestQuota returns a boolean if a field has been set.

### GetFsUserGroupId

`func (o *DfsQuotaCreateReqQuota) GetFsUserGroupId() int64`

GetFsUserGroupId returns the FsUserGroupId field if non-nil, zero value otherwise.

### GetFsUserGroupIdOk

`func (o *DfsQuotaCreateReqQuota) GetFsUserGroupIdOk() (*int64, bool)`

GetFsUserGroupIdOk returns a tuple with the FsUserGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroupId

`func (o *DfsQuotaCreateReqQuota) SetFsUserGroupId(v int64)`

SetFsUserGroupId sets FsUserGroupId field to given value.

### HasFsUserGroupId

`func (o *DfsQuotaCreateReqQuota) HasFsUserGroupId() bool`

HasFsUserGroupId returns a boolean if a field has been set.

### GetFsUserId

`func (o *DfsQuotaCreateReqQuota) GetFsUserId() int64`

GetFsUserId returns the FsUserId field if non-nil, zero value otherwise.

### GetFsUserIdOk

`func (o *DfsQuotaCreateReqQuota) GetFsUserIdOk() (*int64, bool)`

GetFsUserIdOk returns a tuple with the FsUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserId

`func (o *DfsQuotaCreateReqQuota) SetFsUserId(v int64)`

SetFsUserId sets FsUserId field to given value.

### HasFsUserId

`func (o *DfsQuotaCreateReqQuota) HasFsUserId() bool`

HasFsUserId returns a boolean if a field has been set.

### GetLdapId

`func (o *DfsQuotaCreateReqQuota) GetLdapId() int64`

GetLdapId returns the LdapId field if non-nil, zero value otherwise.

### GetLdapIdOk

`func (o *DfsQuotaCreateReqQuota) GetLdapIdOk() (*int64, bool)`

GetLdapIdOk returns a tuple with the LdapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapId

`func (o *DfsQuotaCreateReqQuota) SetLdapId(v int64)`

SetLdapId sets LdapId field to given value.

### HasLdapId

`func (o *DfsQuotaCreateReqQuota) HasLdapId() bool`

HasLdapId returns a boolean if a field has been set.

### GetPath

`func (o *DfsQuotaCreateReqQuota) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsQuotaCreateReqQuota) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsQuotaCreateReqQuota) SetPath(v string)`

SetPath sets Path field to given value.


### GetSizeGraceTime

`func (o *DfsQuotaCreateReqQuota) GetSizeGraceTime() int64`

GetSizeGraceTime returns the SizeGraceTime field if non-nil, zero value otherwise.

### GetSizeGraceTimeOk

`func (o *DfsQuotaCreateReqQuota) GetSizeGraceTimeOk() (*int64, bool)`

GetSizeGraceTimeOk returns a tuple with the SizeGraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGraceTime

`func (o *DfsQuotaCreateReqQuota) SetSizeGraceTime(v int64)`

SetSizeGraceTime sets SizeGraceTime field to given value.

### HasSizeGraceTime

`func (o *DfsQuotaCreateReqQuota) HasSizeGraceTime() bool`

HasSizeGraceTime returns a boolean if a field has been set.

### GetSizeHardQuota

`func (o *DfsQuotaCreateReqQuota) GetSizeHardQuota() int64`

GetSizeHardQuota returns the SizeHardQuota field if non-nil, zero value otherwise.

### GetSizeHardQuotaOk

`func (o *DfsQuotaCreateReqQuota) GetSizeHardQuotaOk() (*int64, bool)`

GetSizeHardQuotaOk returns a tuple with the SizeHardQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeHardQuota

`func (o *DfsQuotaCreateReqQuota) SetSizeHardQuota(v int64)`

SetSizeHardQuota sets SizeHardQuota field to given value.

### HasSizeHardQuota

`func (o *DfsQuotaCreateReqQuota) HasSizeHardQuota() bool`

HasSizeHardQuota returns a boolean if a field has been set.

### GetSizeRatio

`func (o *DfsQuotaCreateReqQuota) GetSizeRatio() float32`

GetSizeRatio returns the SizeRatio field if non-nil, zero value otherwise.

### GetSizeRatioOk

`func (o *DfsQuotaCreateReqQuota) GetSizeRatioOk() (*float32, bool)`

GetSizeRatioOk returns a tuple with the SizeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeRatio

`func (o *DfsQuotaCreateReqQuota) SetSizeRatio(v float32)`

SetSizeRatio sets SizeRatio field to given value.

### HasSizeRatio

`func (o *DfsQuotaCreateReqQuota) HasSizeRatio() bool`

HasSizeRatio returns a boolean if a field has been set.

### GetSizeSoftQuota

`func (o *DfsQuotaCreateReqQuota) GetSizeSoftQuota() int64`

GetSizeSoftQuota returns the SizeSoftQuota field if non-nil, zero value otherwise.

### GetSizeSoftQuotaOk

`func (o *DfsQuotaCreateReqQuota) GetSizeSoftQuotaOk() (*int64, bool)`

GetSizeSoftQuotaOk returns a tuple with the SizeSoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSoftQuota

`func (o *DfsQuotaCreateReqQuota) SetSizeSoftQuota(v int64)`

SetSizeSoftQuota sets SizeSoftQuota field to given value.

### HasSizeSoftQuota

`func (o *DfsQuotaCreateReqQuota) HasSizeSoftQuota() bool`

HasSizeSoftQuota returns a boolean if a field has been set.

### GetSizeSuggestQuota

`func (o *DfsQuotaCreateReqQuota) GetSizeSuggestQuota() int64`

GetSizeSuggestQuota returns the SizeSuggestQuota field if non-nil, zero value otherwise.

### GetSizeSuggestQuotaOk

`func (o *DfsQuotaCreateReqQuota) GetSizeSuggestQuotaOk() (*int64, bool)`

GetSizeSuggestQuotaOk returns a tuple with the SizeSuggestQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSuggestQuota

`func (o *DfsQuotaCreateReqQuota) SetSizeSuggestQuota(v int64)`

SetSizeSuggestQuota sets SizeSuggestQuota field to given value.

### HasSizeSuggestQuota

`func (o *DfsQuotaCreateReqQuota) HasSizeSuggestQuota() bool`

HasSizeSuggestQuota returns a boolean if a field has been set.

### GetType

`func (o *DfsQuotaCreateReqQuota) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsQuotaCreateReqQuota) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsQuotaCreateReqQuota) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsQuotaCreateReqQuota) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserGroupName

`func (o *DfsQuotaCreateReqQuota) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *DfsQuotaCreateReqQuota) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *DfsQuotaCreateReqQuota) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.

### HasUserGroupName

`func (o *DfsQuotaCreateReqQuota) HasUserGroupName() bool`

HasUserGroupName returns a boolean if a field has been set.

### GetUserName

`func (o *DfsQuotaCreateReqQuota) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DfsQuotaCreateReqQuota) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DfsQuotaCreateReqQuota) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DfsQuotaCreateReqQuota) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


