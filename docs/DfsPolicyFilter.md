# DfsPolicyFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainUserGroups** | Pointer to [**[]DomainUserGroup**](DomainUserGroup.md) |  | [optional] 
**AdDomainUsers** | Pointer to [**[]DomainUser**](DomainUser.md) |  | [optional] 
**AdUserGroupIds** | Pointer to **[]int64** |  | [optional] 
**AdUserIds** | Pointer to **[]int64** |  | [optional] 
**DfsStoragePolicy** | Pointer to [**DfsStoragePolicyNestview**](DfsStoragePolicyNestview.md) |  | [optional] 
**ExcludeFileName** | Pointer to **bool** |  | [optional] 
**ExcludeUserName** | Pointer to **bool** |  | [optional] 
**FileNames** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LdapDomainUserGroups** | Pointer to [**[]DomainUserGroup**](DomainUserGroup.md) |  | [optional] 
**LdapDomainUsers** | Pointer to [**[]DomainUser**](DomainUser.md) |  | [optional] 
**LdapUserGroupIds** | Pointer to **[]int64** |  | [optional] 
**LdapUserIds** | Pointer to **[]int64** |  | [optional] 
**UserGroupNames** | Pointer to **[]string** |  | [optional] 
**UserNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDfsPolicyFilter

`func NewDfsPolicyFilter() *DfsPolicyFilter`

NewDfsPolicyFilter instantiates a new DfsPolicyFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsPolicyFilterWithDefaults

`func NewDfsPolicyFilterWithDefaults() *DfsPolicyFilter`

NewDfsPolicyFilterWithDefaults instantiates a new DfsPolicyFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainUserGroups

`func (o *DfsPolicyFilter) GetAdDomainUserGroups() []DomainUserGroup`

GetAdDomainUserGroups returns the AdDomainUserGroups field if non-nil, zero value otherwise.

### GetAdDomainUserGroupsOk

`func (o *DfsPolicyFilter) GetAdDomainUserGroupsOk() (*[]DomainUserGroup, bool)`

GetAdDomainUserGroupsOk returns a tuple with the AdDomainUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainUserGroups

`func (o *DfsPolicyFilter) SetAdDomainUserGroups(v []DomainUserGroup)`

SetAdDomainUserGroups sets AdDomainUserGroups field to given value.

### HasAdDomainUserGroups

`func (o *DfsPolicyFilter) HasAdDomainUserGroups() bool`

HasAdDomainUserGroups returns a boolean if a field has been set.

### GetAdDomainUsers

`func (o *DfsPolicyFilter) GetAdDomainUsers() []DomainUser`

GetAdDomainUsers returns the AdDomainUsers field if non-nil, zero value otherwise.

### GetAdDomainUsersOk

`func (o *DfsPolicyFilter) GetAdDomainUsersOk() (*[]DomainUser, bool)`

GetAdDomainUsersOk returns a tuple with the AdDomainUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainUsers

`func (o *DfsPolicyFilter) SetAdDomainUsers(v []DomainUser)`

SetAdDomainUsers sets AdDomainUsers field to given value.

### HasAdDomainUsers

`func (o *DfsPolicyFilter) HasAdDomainUsers() bool`

HasAdDomainUsers returns a boolean if a field has been set.

### GetAdUserGroupIds

`func (o *DfsPolicyFilter) GetAdUserGroupIds() []int64`

GetAdUserGroupIds returns the AdUserGroupIds field if non-nil, zero value otherwise.

### GetAdUserGroupIdsOk

`func (o *DfsPolicyFilter) GetAdUserGroupIdsOk() (*[]int64, bool)`

GetAdUserGroupIdsOk returns a tuple with the AdUserGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserGroupIds

`func (o *DfsPolicyFilter) SetAdUserGroupIds(v []int64)`

SetAdUserGroupIds sets AdUserGroupIds field to given value.

### HasAdUserGroupIds

`func (o *DfsPolicyFilter) HasAdUserGroupIds() bool`

HasAdUserGroupIds returns a boolean if a field has been set.

### GetAdUserIds

`func (o *DfsPolicyFilter) GetAdUserIds() []int64`

GetAdUserIds returns the AdUserIds field if non-nil, zero value otherwise.

### GetAdUserIdsOk

`func (o *DfsPolicyFilter) GetAdUserIdsOk() (*[]int64, bool)`

GetAdUserIdsOk returns a tuple with the AdUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserIds

`func (o *DfsPolicyFilter) SetAdUserIds(v []int64)`

SetAdUserIds sets AdUserIds field to given value.

### HasAdUserIds

`func (o *DfsPolicyFilter) HasAdUserIds() bool`

HasAdUserIds returns a boolean if a field has been set.

### GetDfsStoragePolicy

`func (o *DfsPolicyFilter) GetDfsStoragePolicy() DfsStoragePolicyNestview`

GetDfsStoragePolicy returns the DfsStoragePolicy field if non-nil, zero value otherwise.

### GetDfsStoragePolicyOk

`func (o *DfsPolicyFilter) GetDfsStoragePolicyOk() (*DfsStoragePolicyNestview, bool)`

GetDfsStoragePolicyOk returns a tuple with the DfsStoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsStoragePolicy

`func (o *DfsPolicyFilter) SetDfsStoragePolicy(v DfsStoragePolicyNestview)`

SetDfsStoragePolicy sets DfsStoragePolicy field to given value.

### HasDfsStoragePolicy

`func (o *DfsPolicyFilter) HasDfsStoragePolicy() bool`

HasDfsStoragePolicy returns a boolean if a field has been set.

### GetExcludeFileName

`func (o *DfsPolicyFilter) GetExcludeFileName() bool`

GetExcludeFileName returns the ExcludeFileName field if non-nil, zero value otherwise.

### GetExcludeFileNameOk

`func (o *DfsPolicyFilter) GetExcludeFileNameOk() (*bool, bool)`

GetExcludeFileNameOk returns a tuple with the ExcludeFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFileName

`func (o *DfsPolicyFilter) SetExcludeFileName(v bool)`

SetExcludeFileName sets ExcludeFileName field to given value.

### HasExcludeFileName

`func (o *DfsPolicyFilter) HasExcludeFileName() bool`

HasExcludeFileName returns a boolean if a field has been set.

### GetExcludeUserName

`func (o *DfsPolicyFilter) GetExcludeUserName() bool`

GetExcludeUserName returns the ExcludeUserName field if non-nil, zero value otherwise.

### GetExcludeUserNameOk

`func (o *DfsPolicyFilter) GetExcludeUserNameOk() (*bool, bool)`

GetExcludeUserNameOk returns a tuple with the ExcludeUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUserName

`func (o *DfsPolicyFilter) SetExcludeUserName(v bool)`

SetExcludeUserName sets ExcludeUserName field to given value.

### HasExcludeUserName

`func (o *DfsPolicyFilter) HasExcludeUserName() bool`

HasExcludeUserName returns a boolean if a field has been set.

### GetFileNames

`func (o *DfsPolicyFilter) GetFileNames() string`

GetFileNames returns the FileNames field if non-nil, zero value otherwise.

### GetFileNamesOk

`func (o *DfsPolicyFilter) GetFileNamesOk() (*string, bool)`

GetFileNamesOk returns a tuple with the FileNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNames

`func (o *DfsPolicyFilter) SetFileNames(v string)`

SetFileNames sets FileNames field to given value.

### HasFileNames

`func (o *DfsPolicyFilter) HasFileNames() bool`

HasFileNames returns a boolean if a field has been set.

### GetId

`func (o *DfsPolicyFilter) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsPolicyFilter) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsPolicyFilter) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsPolicyFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLdapDomainUserGroups

`func (o *DfsPolicyFilter) GetLdapDomainUserGroups() []DomainUserGroup`

GetLdapDomainUserGroups returns the LdapDomainUserGroups field if non-nil, zero value otherwise.

### GetLdapDomainUserGroupsOk

`func (o *DfsPolicyFilter) GetLdapDomainUserGroupsOk() (*[]DomainUserGroup, bool)`

GetLdapDomainUserGroupsOk returns a tuple with the LdapDomainUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDomainUserGroups

`func (o *DfsPolicyFilter) SetLdapDomainUserGroups(v []DomainUserGroup)`

SetLdapDomainUserGroups sets LdapDomainUserGroups field to given value.

### HasLdapDomainUserGroups

`func (o *DfsPolicyFilter) HasLdapDomainUserGroups() bool`

HasLdapDomainUserGroups returns a boolean if a field has been set.

### GetLdapDomainUsers

`func (o *DfsPolicyFilter) GetLdapDomainUsers() []DomainUser`

GetLdapDomainUsers returns the LdapDomainUsers field if non-nil, zero value otherwise.

### GetLdapDomainUsersOk

`func (o *DfsPolicyFilter) GetLdapDomainUsersOk() (*[]DomainUser, bool)`

GetLdapDomainUsersOk returns a tuple with the LdapDomainUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDomainUsers

`func (o *DfsPolicyFilter) SetLdapDomainUsers(v []DomainUser)`

SetLdapDomainUsers sets LdapDomainUsers field to given value.

### HasLdapDomainUsers

`func (o *DfsPolicyFilter) HasLdapDomainUsers() bool`

HasLdapDomainUsers returns a boolean if a field has been set.

### GetLdapUserGroupIds

`func (o *DfsPolicyFilter) GetLdapUserGroupIds() []int64`

GetLdapUserGroupIds returns the LdapUserGroupIds field if non-nil, zero value otherwise.

### GetLdapUserGroupIdsOk

`func (o *DfsPolicyFilter) GetLdapUserGroupIdsOk() (*[]int64, bool)`

GetLdapUserGroupIdsOk returns a tuple with the LdapUserGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserGroupIds

`func (o *DfsPolicyFilter) SetLdapUserGroupIds(v []int64)`

SetLdapUserGroupIds sets LdapUserGroupIds field to given value.

### HasLdapUserGroupIds

`func (o *DfsPolicyFilter) HasLdapUserGroupIds() bool`

HasLdapUserGroupIds returns a boolean if a field has been set.

### GetLdapUserIds

`func (o *DfsPolicyFilter) GetLdapUserIds() []int64`

GetLdapUserIds returns the LdapUserIds field if non-nil, zero value otherwise.

### GetLdapUserIdsOk

`func (o *DfsPolicyFilter) GetLdapUserIdsOk() (*[]int64, bool)`

GetLdapUserIdsOk returns a tuple with the LdapUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserIds

`func (o *DfsPolicyFilter) SetLdapUserIds(v []int64)`

SetLdapUserIds sets LdapUserIds field to given value.

### HasLdapUserIds

`func (o *DfsPolicyFilter) HasLdapUserIds() bool`

HasLdapUserIds returns a boolean if a field has been set.

### GetUserGroupNames

`func (o *DfsPolicyFilter) GetUserGroupNames() []string`

GetUserGroupNames returns the UserGroupNames field if non-nil, zero value otherwise.

### GetUserGroupNamesOk

`func (o *DfsPolicyFilter) GetUserGroupNamesOk() (*[]string, bool)`

GetUserGroupNamesOk returns a tuple with the UserGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupNames

`func (o *DfsPolicyFilter) SetUserGroupNames(v []string)`

SetUserGroupNames sets UserGroupNames field to given value.

### HasUserGroupNames

`func (o *DfsPolicyFilter) HasUserGroupNames() bool`

HasUserGroupNames returns a boolean if a field has been set.

### GetUserNames

`func (o *DfsPolicyFilter) GetUserNames() []string`

GetUserNames returns the UserNames field if non-nil, zero value otherwise.

### GetUserNamesOk

`func (o *DfsPolicyFilter) GetUserNamesOk() (*[]string, bool)`

GetUserNamesOk returns a tuple with the UserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNames

`func (o *DfsPolicyFilter) SetUserNames(v []string)`

SetUserNames sets UserNames field to given value.

### HasUserNames

`func (o *DfsPolicyFilter) HasUserNames() bool`

HasUserNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


