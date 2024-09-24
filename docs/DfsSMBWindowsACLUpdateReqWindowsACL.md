# DfsSMBWindowsACLUpdateReqWindowsACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AceType** | **string** | type of acl, allowed or denied | 
**ApplyTo** | **string** | inherited information | 
**DfsRootfsId** | **int64** | id of rootfs | 
**Path** | **string** | directory path | 
**Permissions** | **[]string** | principal access permissions | 
**Principal** | **string** | the entity of performing access control | 

## Methods

### NewDfsSMBWindowsACLUpdateReqWindowsACL

`func NewDfsSMBWindowsACLUpdateReqWindowsACL(aceType string, applyTo string, dfsRootfsId int64, path string, permissions []string, principal string, ) *DfsSMBWindowsACLUpdateReqWindowsACL`

NewDfsSMBWindowsACLUpdateReqWindowsACL instantiates a new DfsSMBWindowsACLUpdateReqWindowsACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBWindowsACLUpdateReqWindowsACLWithDefaults

`func NewDfsSMBWindowsACLUpdateReqWindowsACLWithDefaults() *DfsSMBWindowsACLUpdateReqWindowsACL`

NewDfsSMBWindowsACLUpdateReqWindowsACLWithDefaults instantiates a new DfsSMBWindowsACLUpdateReqWindowsACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAceType

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetAceType() string`

GetAceType returns the AceType field if non-nil, zero value otherwise.

### GetAceTypeOk

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetAceTypeOk() (*string, bool)`

GetAceTypeOk returns a tuple with the AceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAceType

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) SetAceType(v string)`

SetAceType sets AceType field to given value.


### GetApplyTo

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetApplyTo() string`

GetApplyTo returns the ApplyTo field if non-nil, zero value otherwise.

### GetApplyToOk

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetApplyToOk() (*string, bool)`

GetApplyToOk returns a tuple with the ApplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTo

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) SetApplyTo(v string)`

SetApplyTo sets ApplyTo field to given value.


### GetDfsRootfsId

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetPath

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) SetPath(v string)`

SetPath sets Path field to given value.


### GetPermissions

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetPrincipal

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *DfsSMBWindowsACLUpdateReqWindowsACL) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


