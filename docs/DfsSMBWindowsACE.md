# DfsSMBWindowsACE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AceType** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**FsActiveDirectory** | Pointer to [**FSActiveDirectory**](FSActiveDirectory.md) |  | [optional] 
**FsUserGroup** | Pointer to [**FSUserGroup**](FSUserGroup.md) |  | [optional] 
**Inheritance** | Pointer to [**WindowsACLInheritance**](WindowsACLInheritance.md) |  | [optional] 
**OriginalPermission** | Pointer to **string** | original permission is convenient in the xms-cli output | [optional] 
**Permissions** | Pointer to [**[]WindowsACLPermission**](WindowsACLPermission.md) |  | [optional] 
**Principal** | Pointer to **string** |  | [optional] 
**PrincipalType** | Pointer to **string** |  | [optional] 
**SecurityType** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**FSUser**](FSUser.md) |  | [optional] 
**UserGroupName** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsSMBWindowsACE

`func NewDfsSMBWindowsACE() *DfsSMBWindowsACE`

NewDfsSMBWindowsACE instantiates a new DfsSMBWindowsACE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBWindowsACEWithDefaults

`func NewDfsSMBWindowsACEWithDefaults() *DfsSMBWindowsACE`

NewDfsSMBWindowsACEWithDefaults instantiates a new DfsSMBWindowsACE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAceType

`func (o *DfsSMBWindowsACE) GetAceType() string`

GetAceType returns the AceType field if non-nil, zero value otherwise.

### GetAceTypeOk

`func (o *DfsSMBWindowsACE) GetAceTypeOk() (*string, bool)`

GetAceTypeOk returns a tuple with the AceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAceType

`func (o *DfsSMBWindowsACE) SetAceType(v string)`

SetAceType sets AceType field to given value.

### HasAceType

`func (o *DfsSMBWindowsACE) HasAceType() bool`

HasAceType returns a boolean if a field has been set.

### GetAction

`func (o *DfsSMBWindowsACE) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DfsSMBWindowsACE) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DfsSMBWindowsACE) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DfsSMBWindowsACE) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFsActiveDirectory

`func (o *DfsSMBWindowsACE) GetFsActiveDirectory() FSActiveDirectory`

GetFsActiveDirectory returns the FsActiveDirectory field if non-nil, zero value otherwise.

### GetFsActiveDirectoryOk

`func (o *DfsSMBWindowsACE) GetFsActiveDirectoryOk() (*FSActiveDirectory, bool)`

GetFsActiveDirectoryOk returns a tuple with the FsActiveDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsActiveDirectory

`func (o *DfsSMBWindowsACE) SetFsActiveDirectory(v FSActiveDirectory)`

SetFsActiveDirectory sets FsActiveDirectory field to given value.

### HasFsActiveDirectory

`func (o *DfsSMBWindowsACE) HasFsActiveDirectory() bool`

HasFsActiveDirectory returns a boolean if a field has been set.

### GetFsUserGroup

`func (o *DfsSMBWindowsACE) GetFsUserGroup() FSUserGroup`

GetFsUserGroup returns the FsUserGroup field if non-nil, zero value otherwise.

### GetFsUserGroupOk

`func (o *DfsSMBWindowsACE) GetFsUserGroupOk() (*FSUserGroup, bool)`

GetFsUserGroupOk returns a tuple with the FsUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroup

`func (o *DfsSMBWindowsACE) SetFsUserGroup(v FSUserGroup)`

SetFsUserGroup sets FsUserGroup field to given value.

### HasFsUserGroup

`func (o *DfsSMBWindowsACE) HasFsUserGroup() bool`

HasFsUserGroup returns a boolean if a field has been set.

### GetInheritance

`func (o *DfsSMBWindowsACE) GetInheritance() WindowsACLInheritance`

GetInheritance returns the Inheritance field if non-nil, zero value otherwise.

### GetInheritanceOk

`func (o *DfsSMBWindowsACE) GetInheritanceOk() (*WindowsACLInheritance, bool)`

GetInheritanceOk returns a tuple with the Inheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritance

`func (o *DfsSMBWindowsACE) SetInheritance(v WindowsACLInheritance)`

SetInheritance sets Inheritance field to given value.

### HasInheritance

`func (o *DfsSMBWindowsACE) HasInheritance() bool`

HasInheritance returns a boolean if a field has been set.

### GetOriginalPermission

`func (o *DfsSMBWindowsACE) GetOriginalPermission() string`

GetOriginalPermission returns the OriginalPermission field if non-nil, zero value otherwise.

### GetOriginalPermissionOk

`func (o *DfsSMBWindowsACE) GetOriginalPermissionOk() (*string, bool)`

GetOriginalPermissionOk returns a tuple with the OriginalPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPermission

`func (o *DfsSMBWindowsACE) SetOriginalPermission(v string)`

SetOriginalPermission sets OriginalPermission field to given value.

### HasOriginalPermission

`func (o *DfsSMBWindowsACE) HasOriginalPermission() bool`

HasOriginalPermission returns a boolean if a field has been set.

### GetPermissions

`func (o *DfsSMBWindowsACE) GetPermissions() []WindowsACLPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DfsSMBWindowsACE) GetPermissionsOk() (*[]WindowsACLPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DfsSMBWindowsACE) SetPermissions(v []WindowsACLPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DfsSMBWindowsACE) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPrincipal

`func (o *DfsSMBWindowsACE) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *DfsSMBWindowsACE) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *DfsSMBWindowsACE) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *DfsSMBWindowsACE) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetPrincipalType

`func (o *DfsSMBWindowsACE) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *DfsSMBWindowsACE) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *DfsSMBWindowsACE) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.

### HasPrincipalType

`func (o *DfsSMBWindowsACE) HasPrincipalType() bool`

HasPrincipalType returns a boolean if a field has been set.

### GetSecurityType

`func (o *DfsSMBWindowsACE) GetSecurityType() string`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *DfsSMBWindowsACE) GetSecurityTypeOk() (*string, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *DfsSMBWindowsACE) SetSecurityType(v string)`

SetSecurityType sets SecurityType field to given value.

### HasSecurityType

`func (o *DfsSMBWindowsACE) HasSecurityType() bool`

HasSecurityType returns a boolean if a field has been set.

### GetUser

`func (o *DfsSMBWindowsACE) GetUser() FSUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DfsSMBWindowsACE) GetUserOk() (*FSUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DfsSMBWindowsACE) SetUser(v FSUser)`

SetUser sets User field to given value.

### HasUser

`func (o *DfsSMBWindowsACE) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserGroupName

`func (o *DfsSMBWindowsACE) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *DfsSMBWindowsACE) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *DfsSMBWindowsACE) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.

### HasUserGroupName

`func (o *DfsSMBWindowsACE) HasUserGroupName() bool`

HasUserGroupName returns a boolean if a field has been set.

### GetUserName

`func (o *DfsSMBWindowsACE) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DfsSMBWindowsACE) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DfsSMBWindowsACE) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DfsSMBWindowsACE) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


