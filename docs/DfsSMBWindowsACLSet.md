# DfsSMBWindowsACLSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AceType** | **string** | type of acl like allowed, denied | 
**ActionType** | **string** | api action type like add, modify, delete | 
**ApplyTo** | **string** | inherited information | 
**FsActiveDirectoryId** | Pointer to **int64** | fs active directory id | [optional] 
**Permissions** | **[]string** | principal access permissions | 
**Principal** | Pointer to **string** | the entity of performing access control | [optional] 
**PrincipalType** | **string** | principal type like Everyone, user, user_group | 
**SecurityType** | Pointer to **string** | security type like local, ad | [optional] 
**UserGroupId** | Pointer to **int64** | local fs user group id | [optional] 
**UserGroupName** | Pointer to **string** | domain user group name | [optional] 
**UserId** | Pointer to **int64** | local fs user id | [optional] 
**Username** | Pointer to **string** | domain username | [optional] 

## Methods

### NewDfsSMBWindowsACLSet

`func NewDfsSMBWindowsACLSet(aceType string, actionType string, applyTo string, permissions []string, principalType string, ) *DfsSMBWindowsACLSet`

NewDfsSMBWindowsACLSet instantiates a new DfsSMBWindowsACLSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBWindowsACLSetWithDefaults

`func NewDfsSMBWindowsACLSetWithDefaults() *DfsSMBWindowsACLSet`

NewDfsSMBWindowsACLSetWithDefaults instantiates a new DfsSMBWindowsACLSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAceType

`func (o *DfsSMBWindowsACLSet) GetAceType() string`

GetAceType returns the AceType field if non-nil, zero value otherwise.

### GetAceTypeOk

`func (o *DfsSMBWindowsACLSet) GetAceTypeOk() (*string, bool)`

GetAceTypeOk returns a tuple with the AceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAceType

`func (o *DfsSMBWindowsACLSet) SetAceType(v string)`

SetAceType sets AceType field to given value.


### GetActionType

`func (o *DfsSMBWindowsACLSet) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *DfsSMBWindowsACLSet) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *DfsSMBWindowsACLSet) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetApplyTo

`func (o *DfsSMBWindowsACLSet) GetApplyTo() string`

GetApplyTo returns the ApplyTo field if non-nil, zero value otherwise.

### GetApplyToOk

`func (o *DfsSMBWindowsACLSet) GetApplyToOk() (*string, bool)`

GetApplyToOk returns a tuple with the ApplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTo

`func (o *DfsSMBWindowsACLSet) SetApplyTo(v string)`

SetApplyTo sets ApplyTo field to given value.


### GetFsActiveDirectoryId

`func (o *DfsSMBWindowsACLSet) GetFsActiveDirectoryId() int64`

GetFsActiveDirectoryId returns the FsActiveDirectoryId field if non-nil, zero value otherwise.

### GetFsActiveDirectoryIdOk

`func (o *DfsSMBWindowsACLSet) GetFsActiveDirectoryIdOk() (*int64, bool)`

GetFsActiveDirectoryIdOk returns a tuple with the FsActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsActiveDirectoryId

`func (o *DfsSMBWindowsACLSet) SetFsActiveDirectoryId(v int64)`

SetFsActiveDirectoryId sets FsActiveDirectoryId field to given value.

### HasFsActiveDirectoryId

`func (o *DfsSMBWindowsACLSet) HasFsActiveDirectoryId() bool`

HasFsActiveDirectoryId returns a boolean if a field has been set.

### GetPermissions

`func (o *DfsSMBWindowsACLSet) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DfsSMBWindowsACLSet) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DfsSMBWindowsACLSet) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetPrincipal

`func (o *DfsSMBWindowsACLSet) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *DfsSMBWindowsACLSet) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *DfsSMBWindowsACLSet) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *DfsSMBWindowsACLSet) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetPrincipalType

`func (o *DfsSMBWindowsACLSet) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *DfsSMBWindowsACLSet) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *DfsSMBWindowsACLSet) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.


### GetSecurityType

`func (o *DfsSMBWindowsACLSet) GetSecurityType() string`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *DfsSMBWindowsACLSet) GetSecurityTypeOk() (*string, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *DfsSMBWindowsACLSet) SetSecurityType(v string)`

SetSecurityType sets SecurityType field to given value.

### HasSecurityType

`func (o *DfsSMBWindowsACLSet) HasSecurityType() bool`

HasSecurityType returns a boolean if a field has been set.

### GetUserGroupId

`func (o *DfsSMBWindowsACLSet) GetUserGroupId() int64`

GetUserGroupId returns the UserGroupId field if non-nil, zero value otherwise.

### GetUserGroupIdOk

`func (o *DfsSMBWindowsACLSet) GetUserGroupIdOk() (*int64, bool)`

GetUserGroupIdOk returns a tuple with the UserGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupId

`func (o *DfsSMBWindowsACLSet) SetUserGroupId(v int64)`

SetUserGroupId sets UserGroupId field to given value.

### HasUserGroupId

`func (o *DfsSMBWindowsACLSet) HasUserGroupId() bool`

HasUserGroupId returns a boolean if a field has been set.

### GetUserGroupName

`func (o *DfsSMBWindowsACLSet) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *DfsSMBWindowsACLSet) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *DfsSMBWindowsACLSet) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.

### HasUserGroupName

`func (o *DfsSMBWindowsACLSet) HasUserGroupName() bool`

HasUserGroupName returns a boolean if a field has been set.

### GetUserId

`func (o *DfsSMBWindowsACLSet) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DfsSMBWindowsACLSet) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DfsSMBWindowsACLSet) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DfsSMBWindowsACLSet) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *DfsSMBWindowsACLSet) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DfsSMBWindowsACLSet) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DfsSMBWindowsACLSet) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DfsSMBWindowsACLSet) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


