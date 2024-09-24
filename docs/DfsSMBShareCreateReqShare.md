# DfsSMBShareCreateReqShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessBasedShareEnum** | Pointer to **bool** | smb access-based-share-enum enable | [optional] 
**AclInherited** | Pointer to **bool** | default acl status | [optional] 
**CaseSensitive** | Pointer to **bool** | case sensitive | [optional] 
**Description** | Pointer to **string** | description of share | [optional] 
**DfsGatewayGroupId** | Pointer to **int64** | gateway group id | [optional] 
**DfsRootfsId** | **int64** | rootfs id | 
**DfsSmbShareAcls** | Pointer to [**[]DfsSMBShareACLReq**](DfsSMBShareACLReq.md) | access control array | [optional] 
**HostsAllow** | Pointer to **string** | allow ips in share | [optional] 
**Name** | **string** | name of share | 
**Oplocks** | Pointer to **bool** | smb oplocks enable | [optional] 
**Path** | **string** | directory path | 
**Recycled** | Pointer to **bool** | recycle status | [optional] 
**WindowsAcl** | Pointer to **bool** | windows acl status | [optional] 

## Methods

### NewDfsSMBShareCreateReqShare

`func NewDfsSMBShareCreateReqShare(dfsRootfsId int64, name string, path string, ) *DfsSMBShareCreateReqShare`

NewDfsSMBShareCreateReqShare instantiates a new DfsSMBShareCreateReqShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBShareCreateReqShareWithDefaults

`func NewDfsSMBShareCreateReqShareWithDefaults() *DfsSMBShareCreateReqShare`

NewDfsSMBShareCreateReqShareWithDefaults instantiates a new DfsSMBShareCreateReqShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessBasedShareEnum

`func (o *DfsSMBShareCreateReqShare) GetAccessBasedShareEnum() bool`

GetAccessBasedShareEnum returns the AccessBasedShareEnum field if non-nil, zero value otherwise.

### GetAccessBasedShareEnumOk

`func (o *DfsSMBShareCreateReqShare) GetAccessBasedShareEnumOk() (*bool, bool)`

GetAccessBasedShareEnumOk returns a tuple with the AccessBasedShareEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessBasedShareEnum

`func (o *DfsSMBShareCreateReqShare) SetAccessBasedShareEnum(v bool)`

SetAccessBasedShareEnum sets AccessBasedShareEnum field to given value.

### HasAccessBasedShareEnum

`func (o *DfsSMBShareCreateReqShare) HasAccessBasedShareEnum() bool`

HasAccessBasedShareEnum returns a boolean if a field has been set.

### GetAclInherited

`func (o *DfsSMBShareCreateReqShare) GetAclInherited() bool`

GetAclInherited returns the AclInherited field if non-nil, zero value otherwise.

### GetAclInheritedOk

`func (o *DfsSMBShareCreateReqShare) GetAclInheritedOk() (*bool, bool)`

GetAclInheritedOk returns a tuple with the AclInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclInherited

`func (o *DfsSMBShareCreateReqShare) SetAclInherited(v bool)`

SetAclInherited sets AclInherited field to given value.

### HasAclInherited

`func (o *DfsSMBShareCreateReqShare) HasAclInherited() bool`

HasAclInherited returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *DfsSMBShareCreateReqShare) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *DfsSMBShareCreateReqShare) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *DfsSMBShareCreateReqShare) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *DfsSMBShareCreateReqShare) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetDescription

`func (o *DfsSMBShareCreateReqShare) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsSMBShareCreateReqShare) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsSMBShareCreateReqShare) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsSMBShareCreateReqShare) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroupId

`func (o *DfsSMBShareCreateReqShare) GetDfsGatewayGroupId() int64`

GetDfsGatewayGroupId returns the DfsGatewayGroupId field if non-nil, zero value otherwise.

### GetDfsGatewayGroupIdOk

`func (o *DfsSMBShareCreateReqShare) GetDfsGatewayGroupIdOk() (*int64, bool)`

GetDfsGatewayGroupIdOk returns a tuple with the DfsGatewayGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroupId

`func (o *DfsSMBShareCreateReqShare) SetDfsGatewayGroupId(v int64)`

SetDfsGatewayGroupId sets DfsGatewayGroupId field to given value.

### HasDfsGatewayGroupId

`func (o *DfsSMBShareCreateReqShare) HasDfsGatewayGroupId() bool`

HasDfsGatewayGroupId returns a boolean if a field has been set.

### GetDfsRootfsId

`func (o *DfsSMBShareCreateReqShare) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsSMBShareCreateReqShare) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsSMBShareCreateReqShare) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.


### GetDfsSmbShareAcls

`func (o *DfsSMBShareCreateReqShare) GetDfsSmbShareAcls() []DfsSMBShareACLReq`

GetDfsSmbShareAcls returns the DfsSmbShareAcls field if non-nil, zero value otherwise.

### GetDfsSmbShareAclsOk

`func (o *DfsSMBShareCreateReqShare) GetDfsSmbShareAclsOk() (*[]DfsSMBShareACLReq, bool)`

GetDfsSmbShareAclsOk returns a tuple with the DfsSmbShareAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsSmbShareAcls

`func (o *DfsSMBShareCreateReqShare) SetDfsSmbShareAcls(v []DfsSMBShareACLReq)`

SetDfsSmbShareAcls sets DfsSmbShareAcls field to given value.

### HasDfsSmbShareAcls

`func (o *DfsSMBShareCreateReqShare) HasDfsSmbShareAcls() bool`

HasDfsSmbShareAcls returns a boolean if a field has been set.

### GetHostsAllow

`func (o *DfsSMBShareCreateReqShare) GetHostsAllow() string`

GetHostsAllow returns the HostsAllow field if non-nil, zero value otherwise.

### GetHostsAllowOk

`func (o *DfsSMBShareCreateReqShare) GetHostsAllowOk() (*string, bool)`

GetHostsAllowOk returns a tuple with the HostsAllow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsAllow

`func (o *DfsSMBShareCreateReqShare) SetHostsAllow(v string)`

SetHostsAllow sets HostsAllow field to given value.

### HasHostsAllow

`func (o *DfsSMBShareCreateReqShare) HasHostsAllow() bool`

HasHostsAllow returns a boolean if a field has been set.

### GetName

`func (o *DfsSMBShareCreateReqShare) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsSMBShareCreateReqShare) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsSMBShareCreateReqShare) SetName(v string)`

SetName sets Name field to given value.


### GetOplocks

`func (o *DfsSMBShareCreateReqShare) GetOplocks() bool`

GetOplocks returns the Oplocks field if non-nil, zero value otherwise.

### GetOplocksOk

`func (o *DfsSMBShareCreateReqShare) GetOplocksOk() (*bool, bool)`

GetOplocksOk returns a tuple with the Oplocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplocks

`func (o *DfsSMBShareCreateReqShare) SetOplocks(v bool)`

SetOplocks sets Oplocks field to given value.

### HasOplocks

`func (o *DfsSMBShareCreateReqShare) HasOplocks() bool`

HasOplocks returns a boolean if a field has been set.

### GetPath

`func (o *DfsSMBShareCreateReqShare) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsSMBShareCreateReqShare) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsSMBShareCreateReqShare) SetPath(v string)`

SetPath sets Path field to given value.


### GetRecycled

`func (o *DfsSMBShareCreateReqShare) GetRecycled() bool`

GetRecycled returns the Recycled field if non-nil, zero value otherwise.

### GetRecycledOk

`func (o *DfsSMBShareCreateReqShare) GetRecycledOk() (*bool, bool)`

GetRecycledOk returns a tuple with the Recycled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycled

`func (o *DfsSMBShareCreateReqShare) SetRecycled(v bool)`

SetRecycled sets Recycled field to given value.

### HasRecycled

`func (o *DfsSMBShareCreateReqShare) HasRecycled() bool`

HasRecycled returns a boolean if a field has been set.

### GetWindowsAcl

`func (o *DfsSMBShareCreateReqShare) GetWindowsAcl() bool`

GetWindowsAcl returns the WindowsAcl field if non-nil, zero value otherwise.

### GetWindowsAclOk

`func (o *DfsSMBShareCreateReqShare) GetWindowsAclOk() (*bool, bool)`

GetWindowsAclOk returns a tuple with the WindowsAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsAcl

`func (o *DfsSMBShareCreateReqShare) SetWindowsAcl(v bool)`

SetWindowsAcl sets WindowsAcl field to given value.

### HasWindowsAcl

`func (o *DfsSMBShareCreateReqShare) HasWindowsAcl() bool`

HasWindowsAcl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


