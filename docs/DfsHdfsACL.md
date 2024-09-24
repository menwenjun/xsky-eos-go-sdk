# DfsHdfsACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsHdfs** | Pointer to [**DfsHdfsNestview**](DfsHdfsNestview.md) |  | [optional] 
**DomainUser** | Pointer to [**DomainUser**](DomainUser.md) |  | [optional] 
**DomainUserGroup** | Pointer to [**DomainUserGroup**](DomainUserGroup.md) |  | [optional] 
**FsUser** | Pointer to [**FSUserNestview**](FSUserNestview.md) |  | [optional] 
**FsUserGroup** | Pointer to [**FSUserGroupNestview**](FSUserGroupNestview.md) |  | [optional] 
**FsUserGroupName** | Pointer to **string** |  | [optional] 
**FsUserName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsProxyUser** | Pointer to **bool** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**Security** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsHdfsACL

`func NewDfsHdfsACL() *DfsHdfsACL`

NewDfsHdfsACL instantiates a new DfsHdfsACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsHdfsACLWithDefaults

`func NewDfsHdfsACLWithDefaults() *DfsHdfsACL`

NewDfsHdfsACLWithDefaults instantiates a new DfsHdfsACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsHdfsACL) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsHdfsACL) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsHdfsACL) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsHdfsACL) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsHdfsACL) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsHdfsACL) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsHdfsACL) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsHdfsACL) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsHdfs

`func (o *DfsHdfsACL) GetDfsHdfs() DfsHdfsNestview`

GetDfsHdfs returns the DfsHdfs field if non-nil, zero value otherwise.

### GetDfsHdfsOk

`func (o *DfsHdfsACL) GetDfsHdfsOk() (*DfsHdfsNestview, bool)`

GetDfsHdfsOk returns a tuple with the DfsHdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfs

`func (o *DfsHdfsACL) SetDfsHdfs(v DfsHdfsNestview)`

SetDfsHdfs sets DfsHdfs field to given value.

### HasDfsHdfs

`func (o *DfsHdfsACL) HasDfsHdfs() bool`

HasDfsHdfs returns a boolean if a field has been set.

### GetDomainUser

`func (o *DfsHdfsACL) GetDomainUser() DomainUser`

GetDomainUser returns the DomainUser field if non-nil, zero value otherwise.

### GetDomainUserOk

`func (o *DfsHdfsACL) GetDomainUserOk() (*DomainUser, bool)`

GetDomainUserOk returns a tuple with the DomainUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUser

`func (o *DfsHdfsACL) SetDomainUser(v DomainUser)`

SetDomainUser sets DomainUser field to given value.

### HasDomainUser

`func (o *DfsHdfsACL) HasDomainUser() bool`

HasDomainUser returns a boolean if a field has been set.

### GetDomainUserGroup

`func (o *DfsHdfsACL) GetDomainUserGroup() DomainUserGroup`

GetDomainUserGroup returns the DomainUserGroup field if non-nil, zero value otherwise.

### GetDomainUserGroupOk

`func (o *DfsHdfsACL) GetDomainUserGroupOk() (*DomainUserGroup, bool)`

GetDomainUserGroupOk returns a tuple with the DomainUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUserGroup

`func (o *DfsHdfsACL) SetDomainUserGroup(v DomainUserGroup)`

SetDomainUserGroup sets DomainUserGroup field to given value.

### HasDomainUserGroup

`func (o *DfsHdfsACL) HasDomainUserGroup() bool`

HasDomainUserGroup returns a boolean if a field has been set.

### GetFsUser

`func (o *DfsHdfsACL) GetFsUser() FSUserNestview`

GetFsUser returns the FsUser field if non-nil, zero value otherwise.

### GetFsUserOk

`func (o *DfsHdfsACL) GetFsUserOk() (*FSUserNestview, bool)`

GetFsUserOk returns a tuple with the FsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUser

`func (o *DfsHdfsACL) SetFsUser(v FSUserNestview)`

SetFsUser sets FsUser field to given value.

### HasFsUser

`func (o *DfsHdfsACL) HasFsUser() bool`

HasFsUser returns a boolean if a field has been set.

### GetFsUserGroup

`func (o *DfsHdfsACL) GetFsUserGroup() FSUserGroupNestview`

GetFsUserGroup returns the FsUserGroup field if non-nil, zero value otherwise.

### GetFsUserGroupOk

`func (o *DfsHdfsACL) GetFsUserGroupOk() (*FSUserGroupNestview, bool)`

GetFsUserGroupOk returns a tuple with the FsUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroup

`func (o *DfsHdfsACL) SetFsUserGroup(v FSUserGroupNestview)`

SetFsUserGroup sets FsUserGroup field to given value.

### HasFsUserGroup

`func (o *DfsHdfsACL) HasFsUserGroup() bool`

HasFsUserGroup returns a boolean if a field has been set.

### GetFsUserGroupName

`func (o *DfsHdfsACL) GetFsUserGroupName() string`

GetFsUserGroupName returns the FsUserGroupName field if non-nil, zero value otherwise.

### GetFsUserGroupNameOk

`func (o *DfsHdfsACL) GetFsUserGroupNameOk() (*string, bool)`

GetFsUserGroupNameOk returns a tuple with the FsUserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroupName

`func (o *DfsHdfsACL) SetFsUserGroupName(v string)`

SetFsUserGroupName sets FsUserGroupName field to given value.

### HasFsUserGroupName

`func (o *DfsHdfsACL) HasFsUserGroupName() bool`

HasFsUserGroupName returns a boolean if a field has been set.

### GetFsUserName

`func (o *DfsHdfsACL) GetFsUserName() string`

GetFsUserName returns the FsUserName field if non-nil, zero value otherwise.

### GetFsUserNameOk

`func (o *DfsHdfsACL) GetFsUserNameOk() (*string, bool)`

GetFsUserNameOk returns a tuple with the FsUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserName

`func (o *DfsHdfsACL) SetFsUserName(v string)`

SetFsUserName sets FsUserName field to given value.

### HasFsUserName

`func (o *DfsHdfsACL) HasFsUserName() bool`

HasFsUserName returns a boolean if a field has been set.

### GetId

`func (o *DfsHdfsACL) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsHdfsACL) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsHdfsACL) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsHdfsACL) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsProxyUser

`func (o *DfsHdfsACL) GetIsProxyUser() bool`

GetIsProxyUser returns the IsProxyUser field if non-nil, zero value otherwise.

### GetIsProxyUserOk

`func (o *DfsHdfsACL) GetIsProxyUserOk() (*bool, bool)`

GetIsProxyUserOk returns a tuple with the IsProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxyUser

`func (o *DfsHdfsACL) SetIsProxyUser(v bool)`

SetIsProxyUser sets IsProxyUser field to given value.

### HasIsProxyUser

`func (o *DfsHdfsACL) HasIsProxyUser() bool`

HasIsProxyUser returns a boolean if a field has been set.

### GetPermission

`func (o *DfsHdfsACL) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *DfsHdfsACL) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *DfsHdfsACL) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *DfsHdfsACL) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetSecurity

`func (o *DfsHdfsACL) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DfsHdfsACL) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DfsHdfsACL) SetSecurity(v string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DfsHdfsACL) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetStatus

`func (o *DfsHdfsACL) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsHdfsACL) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsHdfsACL) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsHdfsACL) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *DfsHdfsACL) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsHdfsACL) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsHdfsACL) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsHdfsACL) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsHdfsACL) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsHdfsACL) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsHdfsACL) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsHdfsACL) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


