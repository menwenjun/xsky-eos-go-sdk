# DfsSMBShareACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsSmbShare** | Pointer to [**DfsSMBShareNestview**](DfsSMBShareNestview.md) |  | [optional] 
**DomainUser** | Pointer to [**DomainUser**](DomainUser.md) |  | [optional] 
**DomainUserGroup** | Pointer to [**DomainUserGroup**](DomainUserGroup.md) |  | [optional] 
**FsUser** | Pointer to [**FSUserNestview**](FSUserNestview.md) |  | [optional] 
**FsUserGroup** | Pointer to [**FSUserGroupNestview**](FSUserGroupNestview.md) |  | [optional] 
**FsUserGroupName** | Pointer to **string** |  | [optional] 
**FsUserName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**Security** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsSMBShareACL

`func NewDfsSMBShareACL() *DfsSMBShareACL`

NewDfsSMBShareACL instantiates a new DfsSMBShareACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBShareACLWithDefaults

`func NewDfsSMBShareACLWithDefaults() *DfsSMBShareACL`

NewDfsSMBShareACLWithDefaults instantiates a new DfsSMBShareACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsSMBShareACL) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsSMBShareACL) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsSMBShareACL) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsSMBShareACL) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsSMBShareACL) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsSMBShareACL) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsSMBShareACL) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsSMBShareACL) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsSmbShare

`func (o *DfsSMBShareACL) GetDfsSmbShare() DfsSMBShareNestview`

GetDfsSmbShare returns the DfsSmbShare field if non-nil, zero value otherwise.

### GetDfsSmbShareOk

`func (o *DfsSMBShareACL) GetDfsSmbShareOk() (*DfsSMBShareNestview, bool)`

GetDfsSmbShareOk returns a tuple with the DfsSmbShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsSmbShare

`func (o *DfsSMBShareACL) SetDfsSmbShare(v DfsSMBShareNestview)`

SetDfsSmbShare sets DfsSmbShare field to given value.

### HasDfsSmbShare

`func (o *DfsSMBShareACL) HasDfsSmbShare() bool`

HasDfsSmbShare returns a boolean if a field has been set.

### GetDomainUser

`func (o *DfsSMBShareACL) GetDomainUser() DomainUser`

GetDomainUser returns the DomainUser field if non-nil, zero value otherwise.

### GetDomainUserOk

`func (o *DfsSMBShareACL) GetDomainUserOk() (*DomainUser, bool)`

GetDomainUserOk returns a tuple with the DomainUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUser

`func (o *DfsSMBShareACL) SetDomainUser(v DomainUser)`

SetDomainUser sets DomainUser field to given value.

### HasDomainUser

`func (o *DfsSMBShareACL) HasDomainUser() bool`

HasDomainUser returns a boolean if a field has been set.

### GetDomainUserGroup

`func (o *DfsSMBShareACL) GetDomainUserGroup() DomainUserGroup`

GetDomainUserGroup returns the DomainUserGroup field if non-nil, zero value otherwise.

### GetDomainUserGroupOk

`func (o *DfsSMBShareACL) GetDomainUserGroupOk() (*DomainUserGroup, bool)`

GetDomainUserGroupOk returns a tuple with the DomainUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUserGroup

`func (o *DfsSMBShareACL) SetDomainUserGroup(v DomainUserGroup)`

SetDomainUserGroup sets DomainUserGroup field to given value.

### HasDomainUserGroup

`func (o *DfsSMBShareACL) HasDomainUserGroup() bool`

HasDomainUserGroup returns a boolean if a field has been set.

### GetFsUser

`func (o *DfsSMBShareACL) GetFsUser() FSUserNestview`

GetFsUser returns the FsUser field if non-nil, zero value otherwise.

### GetFsUserOk

`func (o *DfsSMBShareACL) GetFsUserOk() (*FSUserNestview, bool)`

GetFsUserOk returns a tuple with the FsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUser

`func (o *DfsSMBShareACL) SetFsUser(v FSUserNestview)`

SetFsUser sets FsUser field to given value.

### HasFsUser

`func (o *DfsSMBShareACL) HasFsUser() bool`

HasFsUser returns a boolean if a field has been set.

### GetFsUserGroup

`func (o *DfsSMBShareACL) GetFsUserGroup() FSUserGroupNestview`

GetFsUserGroup returns the FsUserGroup field if non-nil, zero value otherwise.

### GetFsUserGroupOk

`func (o *DfsSMBShareACL) GetFsUserGroupOk() (*FSUserGroupNestview, bool)`

GetFsUserGroupOk returns a tuple with the FsUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroup

`func (o *DfsSMBShareACL) SetFsUserGroup(v FSUserGroupNestview)`

SetFsUserGroup sets FsUserGroup field to given value.

### HasFsUserGroup

`func (o *DfsSMBShareACL) HasFsUserGroup() bool`

HasFsUserGroup returns a boolean if a field has been set.

### GetFsUserGroupName

`func (o *DfsSMBShareACL) GetFsUserGroupName() string`

GetFsUserGroupName returns the FsUserGroupName field if non-nil, zero value otherwise.

### GetFsUserGroupNameOk

`func (o *DfsSMBShareACL) GetFsUserGroupNameOk() (*string, bool)`

GetFsUserGroupNameOk returns a tuple with the FsUserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroupName

`func (o *DfsSMBShareACL) SetFsUserGroupName(v string)`

SetFsUserGroupName sets FsUserGroupName field to given value.

### HasFsUserGroupName

`func (o *DfsSMBShareACL) HasFsUserGroupName() bool`

HasFsUserGroupName returns a boolean if a field has been set.

### GetFsUserName

`func (o *DfsSMBShareACL) GetFsUserName() string`

GetFsUserName returns the FsUserName field if non-nil, zero value otherwise.

### GetFsUserNameOk

`func (o *DfsSMBShareACL) GetFsUserNameOk() (*string, bool)`

GetFsUserNameOk returns a tuple with the FsUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserName

`func (o *DfsSMBShareACL) SetFsUserName(v string)`

SetFsUserName sets FsUserName field to given value.

### HasFsUserName

`func (o *DfsSMBShareACL) HasFsUserName() bool`

HasFsUserName returns a boolean if a field has been set.

### GetId

`func (o *DfsSMBShareACL) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsSMBShareACL) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsSMBShareACL) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsSMBShareACL) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermission

`func (o *DfsSMBShareACL) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *DfsSMBShareACL) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *DfsSMBShareACL) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *DfsSMBShareACL) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetSecurity

`func (o *DfsSMBShareACL) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DfsSMBShareACL) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DfsSMBShareACL) SetSecurity(v string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DfsSMBShareACL) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetType

`func (o *DfsSMBShareACL) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DfsSMBShareACL) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DfsSMBShareACL) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DfsSMBShareACL) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsSMBShareACL) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsSMBShareACL) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsSMBShareACL) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsSMBShareACL) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


