# DfsSMBShareUpdateReqShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessBasedShareEnum** | Pointer to **bool** | smb access-based-share-enum enable | [optional] 
**AclInherited** | Pointer to **bool** | default acl status | [optional] 
**CaseSensitive** | Pointer to **bool** | case sensitive | [optional] 
**Description** | Pointer to **string** | description of share | [optional] 
**DfsGatewayGroupId** | Pointer to **int64** | dfs gateway group id of share | [optional] 
**HostsAllow** | Pointer to **string** | allow ips in share | [optional] 
**Name** | Pointer to **string** | name of share | [optional] 
**Oplocks** | Pointer to **bool** | smb oplocks enable | [optional] 
**Recycled** | Pointer to **bool** | recycle status | [optional] 
**WindowsAcl** | Pointer to **bool** | windows acl status | [optional] 

## Methods

### NewDfsSMBShareUpdateReqShare

`func NewDfsSMBShareUpdateReqShare() *DfsSMBShareUpdateReqShare`

NewDfsSMBShareUpdateReqShare instantiates a new DfsSMBShareUpdateReqShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBShareUpdateReqShareWithDefaults

`func NewDfsSMBShareUpdateReqShareWithDefaults() *DfsSMBShareUpdateReqShare`

NewDfsSMBShareUpdateReqShareWithDefaults instantiates a new DfsSMBShareUpdateReqShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessBasedShareEnum

`func (o *DfsSMBShareUpdateReqShare) GetAccessBasedShareEnum() bool`

GetAccessBasedShareEnum returns the AccessBasedShareEnum field if non-nil, zero value otherwise.

### GetAccessBasedShareEnumOk

`func (o *DfsSMBShareUpdateReqShare) GetAccessBasedShareEnumOk() (*bool, bool)`

GetAccessBasedShareEnumOk returns a tuple with the AccessBasedShareEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessBasedShareEnum

`func (o *DfsSMBShareUpdateReqShare) SetAccessBasedShareEnum(v bool)`

SetAccessBasedShareEnum sets AccessBasedShareEnum field to given value.

### HasAccessBasedShareEnum

`func (o *DfsSMBShareUpdateReqShare) HasAccessBasedShareEnum() bool`

HasAccessBasedShareEnum returns a boolean if a field has been set.

### GetAclInherited

`func (o *DfsSMBShareUpdateReqShare) GetAclInherited() bool`

GetAclInherited returns the AclInherited field if non-nil, zero value otherwise.

### GetAclInheritedOk

`func (o *DfsSMBShareUpdateReqShare) GetAclInheritedOk() (*bool, bool)`

GetAclInheritedOk returns a tuple with the AclInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclInherited

`func (o *DfsSMBShareUpdateReqShare) SetAclInherited(v bool)`

SetAclInherited sets AclInherited field to given value.

### HasAclInherited

`func (o *DfsSMBShareUpdateReqShare) HasAclInherited() bool`

HasAclInherited returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *DfsSMBShareUpdateReqShare) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *DfsSMBShareUpdateReqShare) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *DfsSMBShareUpdateReqShare) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *DfsSMBShareUpdateReqShare) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetDescription

`func (o *DfsSMBShareUpdateReqShare) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsSMBShareUpdateReqShare) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsSMBShareUpdateReqShare) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsSMBShareUpdateReqShare) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroupId

`func (o *DfsSMBShareUpdateReqShare) GetDfsGatewayGroupId() int64`

GetDfsGatewayGroupId returns the DfsGatewayGroupId field if non-nil, zero value otherwise.

### GetDfsGatewayGroupIdOk

`func (o *DfsSMBShareUpdateReqShare) GetDfsGatewayGroupIdOk() (*int64, bool)`

GetDfsGatewayGroupIdOk returns a tuple with the DfsGatewayGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroupId

`func (o *DfsSMBShareUpdateReqShare) SetDfsGatewayGroupId(v int64)`

SetDfsGatewayGroupId sets DfsGatewayGroupId field to given value.

### HasDfsGatewayGroupId

`func (o *DfsSMBShareUpdateReqShare) HasDfsGatewayGroupId() bool`

HasDfsGatewayGroupId returns a boolean if a field has been set.

### GetHostsAllow

`func (o *DfsSMBShareUpdateReqShare) GetHostsAllow() string`

GetHostsAllow returns the HostsAllow field if non-nil, zero value otherwise.

### GetHostsAllowOk

`func (o *DfsSMBShareUpdateReqShare) GetHostsAllowOk() (*string, bool)`

GetHostsAllowOk returns a tuple with the HostsAllow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsAllow

`func (o *DfsSMBShareUpdateReqShare) SetHostsAllow(v string)`

SetHostsAllow sets HostsAllow field to given value.

### HasHostsAllow

`func (o *DfsSMBShareUpdateReqShare) HasHostsAllow() bool`

HasHostsAllow returns a boolean if a field has been set.

### GetName

`func (o *DfsSMBShareUpdateReqShare) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsSMBShareUpdateReqShare) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsSMBShareUpdateReqShare) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsSMBShareUpdateReqShare) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOplocks

`func (o *DfsSMBShareUpdateReqShare) GetOplocks() bool`

GetOplocks returns the Oplocks field if non-nil, zero value otherwise.

### GetOplocksOk

`func (o *DfsSMBShareUpdateReqShare) GetOplocksOk() (*bool, bool)`

GetOplocksOk returns a tuple with the Oplocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplocks

`func (o *DfsSMBShareUpdateReqShare) SetOplocks(v bool)`

SetOplocks sets Oplocks field to given value.

### HasOplocks

`func (o *DfsSMBShareUpdateReqShare) HasOplocks() bool`

HasOplocks returns a boolean if a field has been set.

### GetRecycled

`func (o *DfsSMBShareUpdateReqShare) GetRecycled() bool`

GetRecycled returns the Recycled field if non-nil, zero value otherwise.

### GetRecycledOk

`func (o *DfsSMBShareUpdateReqShare) GetRecycledOk() (*bool, bool)`

GetRecycledOk returns a tuple with the Recycled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycled

`func (o *DfsSMBShareUpdateReqShare) SetRecycled(v bool)`

SetRecycled sets Recycled field to given value.

### HasRecycled

`func (o *DfsSMBShareUpdateReqShare) HasRecycled() bool`

HasRecycled returns a boolean if a field has been set.

### GetWindowsAcl

`func (o *DfsSMBShareUpdateReqShare) GetWindowsAcl() bool`

GetWindowsAcl returns the WindowsAcl field if non-nil, zero value otherwise.

### GetWindowsAclOk

`func (o *DfsSMBShareUpdateReqShare) GetWindowsAclOk() (*bool, bool)`

GetWindowsAclOk returns a tuple with the WindowsAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsAcl

`func (o *DfsSMBShareUpdateReqShare) SetWindowsAcl(v bool)`

SetWindowsAcl sets WindowsAcl field to given value.

### HasWindowsAcl

`func (o *DfsSMBShareUpdateReqShare) HasWindowsAcl() bool`

HasWindowsAcl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


