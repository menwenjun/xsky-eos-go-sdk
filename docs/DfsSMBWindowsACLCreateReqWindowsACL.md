# DfsSMBWindowsACLCreateReqWindowsACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AceType** | **string** | type of acl, allowed or denied | 
**ApplyTo** | **string** | inherited information | 
**DfsRootfsId** | **int64** | id of rootfs | 
**FsActiveDirectoryId** | Pointer to **int64** | fs active directory id | [optional] 
**Path** | **string** | subdirectory or sub file of share path | 
**Permissions** | **[]string** | principal access permissions | 
**PrincipalType** | **string** | principal type such as Everyone, user, user_group | 
**SecurityType** | **string** | security type like local, ad | 
**UserGroupId** | Pointer to **int64** | local fs user group id | [optional] 
**UserGroupName** | Pointer to **string** | domain user_group_name | [optional] 
**UserId** | Pointer to **int64** | local fs user id | [optional] 
**Username** | Pointer to **string** | domain username | [optional] 

## Methods

### NewDfsSMBWindowsACLCreateReqWindowsACL

`func NewDfsSMBWindowsACLCreateReqWindowsACL(aceType string, applyTo string, dfsRootfsId int64, path string, permissions []string, principalType string, securityType string, ) *DfsSMBWindowsACLCreateReqWindowsACL`

NewDfsSMBWindowsACLCreateReqWindowsACL instantiates a new DfsSMBWindowsACLCreateReqWindowsACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBWindowsACLCreateReqWindowsACLWithDefaults

`func NewDfsSMBWindowsACLCreateReqWindowsACLWithDefaults() *DfsSMBWindowsACLCreateReqWindowsACL`

NewDfsSMBWindowsACLCreateReqWindowsACLWithDefaults instantiates a new DfsSMBWindowsACLCreateReqWindowsACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAceType

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetAceType() string`

GetAceType returns the AceType field if non-nil, zero value otherwise.

### GetAceTypeOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetAceTypeOk() (*string, bool)`

GetAceTypeOk returns a tuple with the AceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAceType

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetAceType(v string)`

SetAceType sets AceType field to given value.


### GetApplyTo

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetApplyTo() string`

GetApplyTo returns the ApplyTo field if non-nil, zero value otherwise.

### GetApplyToOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetApplyToOk() (*string, bool)`

GetApplyToOk returns a tuple with the ApplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTo

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetApplyTo(v string)`

SetApplyTo sets ApplyTo field to given value.


### GetDfsRootfsId

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetFsActiveDirectoryId

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetFsActiveDirectoryId() int64`

GetFsActiveDirectoryId returns the FsActiveDirectoryId field if non-nil, zero value otherwise.

### GetFsActiveDirectoryIdOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetFsActiveDirectoryIdOk() (*int64, bool)`

GetFsActiveDirectoryIdOk returns a tuple with the FsActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsActiveDirectoryId

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetFsActiveDirectoryId(v int64)`

SetFsActiveDirectoryId sets FsActiveDirectoryId field to given value.

### HasFsActiveDirectoryId

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) HasFsActiveDirectoryId() bool`

HasFsActiveDirectoryId returns a boolean if a field has been set.

### GetPath

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetPath(v string)`

SetPath sets Path field to given value.


### GetPermissions

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetPrincipalType

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.


### GetSecurityType

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetSecurityType() string`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetSecurityTypeOk() (*string, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetSecurityType(v string)`

SetSecurityType sets SecurityType field to given value.


### GetUserGroupId

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetUserGroupId() int64`

GetUserGroupId returns the UserGroupId field if non-nil, zero value otherwise.

### GetUserGroupIdOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetUserGroupIdOk() (*int64, bool)`

GetUserGroupIdOk returns a tuple with the UserGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupId

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetUserGroupId(v int64)`

SetUserGroupId sets UserGroupId field to given value.

### HasUserGroupId

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) HasUserGroupId() bool`

HasUserGroupId returns a boolean if a field has been set.

### GetUserGroupName

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.

### HasUserGroupName

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) HasUserGroupName() bool`

HasUserGroupName returns a boolean if a field has been set.

### GetUserId

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DfsSMBWindowsACLCreateReqWindowsACL) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


