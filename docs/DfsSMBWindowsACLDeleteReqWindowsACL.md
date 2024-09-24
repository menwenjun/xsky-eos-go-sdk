# DfsSMBWindowsACLDeleteReqWindowsACL

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

### NewDfsSMBWindowsACLDeleteReqWindowsACL

`func NewDfsSMBWindowsACLDeleteReqWindowsACL(aceType string, applyTo string, dfsRootfsId int64, path string, permissions []string, principal string, ) *DfsSMBWindowsACLDeleteReqWindowsACL`

NewDfsSMBWindowsACLDeleteReqWindowsACL instantiates a new DfsSMBWindowsACLDeleteReqWindowsACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBWindowsACLDeleteReqWindowsACLWithDefaults

`func NewDfsSMBWindowsACLDeleteReqWindowsACLWithDefaults() *DfsSMBWindowsACLDeleteReqWindowsACL`

NewDfsSMBWindowsACLDeleteReqWindowsACLWithDefaults instantiates a new DfsSMBWindowsACLDeleteReqWindowsACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAceType

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetAceType() string`

GetAceType returns the AceType field if non-nil, zero value otherwise.

### GetAceTypeOk

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetAceTypeOk() (*string, bool)`

GetAceTypeOk returns a tuple with the AceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAceType

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) SetAceType(v string)`

SetAceType sets AceType field to given value.


### GetApplyTo

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetApplyTo() string`

GetApplyTo returns the ApplyTo field if non-nil, zero value otherwise.

### GetApplyToOk

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetApplyToOk() (*string, bool)`

GetApplyToOk returns a tuple with the ApplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTo

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) SetApplyTo(v string)`

SetApplyTo sets ApplyTo field to given value.


### GetDfsRootfsId

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetPath

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) SetPath(v string)`

SetPath sets Path field to given value.


### GetPermissions

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetPrincipal

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *DfsSMBWindowsACLDeleteReqWindowsACL) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


