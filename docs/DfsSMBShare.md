# DfsSMBShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessBasedShareEnum** | Pointer to **bool** |  | [optional] 
**AclInherited** | Pointer to **bool** |  | [optional] 
**ActionStatus** | Pointer to **string** |  | [optional] 
**CaseSensitive** | Pointer to **bool** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DfsGatewayGroup** | Pointer to [**DfsGatewayGroupNestview**](DfsGatewayGroupNestview.md) |  | [optional] 
**DfsPath** | Pointer to [**DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**HostsAllow** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Oplocks** | Pointer to **bool** |  | [optional] 
**Recycled** | Pointer to **bool** |  | [optional] 
**Securities** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalSnapshotNum** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**WindowsAcl** | Pointer to **bool** |  | [optional] 

## Methods

### NewDfsSMBShare

`func NewDfsSMBShare() *DfsSMBShare`

NewDfsSMBShare instantiates a new DfsSMBShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBShareWithDefaults

`func NewDfsSMBShareWithDefaults() *DfsSMBShare`

NewDfsSMBShareWithDefaults instantiates a new DfsSMBShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessBasedShareEnum

`func (o *DfsSMBShare) GetAccessBasedShareEnum() bool`

GetAccessBasedShareEnum returns the AccessBasedShareEnum field if non-nil, zero value otherwise.

### GetAccessBasedShareEnumOk

`func (o *DfsSMBShare) GetAccessBasedShareEnumOk() (*bool, bool)`

GetAccessBasedShareEnumOk returns a tuple with the AccessBasedShareEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessBasedShareEnum

`func (o *DfsSMBShare) SetAccessBasedShareEnum(v bool)`

SetAccessBasedShareEnum sets AccessBasedShareEnum field to given value.

### HasAccessBasedShareEnum

`func (o *DfsSMBShare) HasAccessBasedShareEnum() bool`

HasAccessBasedShareEnum returns a boolean if a field has been set.

### GetAclInherited

`func (o *DfsSMBShare) GetAclInherited() bool`

GetAclInherited returns the AclInherited field if non-nil, zero value otherwise.

### GetAclInheritedOk

`func (o *DfsSMBShare) GetAclInheritedOk() (*bool, bool)`

GetAclInheritedOk returns a tuple with the AclInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclInherited

`func (o *DfsSMBShare) SetAclInherited(v bool)`

SetAclInherited sets AclInherited field to given value.

### HasAclInherited

`func (o *DfsSMBShare) HasAclInherited() bool`

HasAclInherited returns a boolean if a field has been set.

### GetActionStatus

`func (o *DfsSMBShare) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsSMBShare) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsSMBShare) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsSMBShare) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *DfsSMBShare) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *DfsSMBShare) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *DfsSMBShare) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *DfsSMBShare) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetCluster

`func (o *DfsSMBShare) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsSMBShare) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsSMBShare) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsSMBShare) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsSMBShare) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsSMBShare) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsSMBShare) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsSMBShare) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DfsSMBShare) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsSMBShare) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsSMBShare) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsSMBShare) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsSMBShare) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsSMBShare) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsSMBShare) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsSMBShare) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetDfsPath

`func (o *DfsSMBShare) GetDfsPath() DfsPathNestview`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DfsSMBShare) GetDfsPathOk() (*DfsPathNestview, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DfsSMBShare) SetDfsPath(v DfsPathNestview)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DfsSMBShare) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsSMBShare) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsSMBShare) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsSMBShare) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsSMBShare) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetHostsAllow

`func (o *DfsSMBShare) GetHostsAllow() string`

GetHostsAllow returns the HostsAllow field if non-nil, zero value otherwise.

### GetHostsAllowOk

`func (o *DfsSMBShare) GetHostsAllowOk() (*string, bool)`

GetHostsAllowOk returns a tuple with the HostsAllow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsAllow

`func (o *DfsSMBShare) SetHostsAllow(v string)`

SetHostsAllow sets HostsAllow field to given value.

### HasHostsAllow

`func (o *DfsSMBShare) HasHostsAllow() bool`

HasHostsAllow returns a boolean if a field has been set.

### GetId

`func (o *DfsSMBShare) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsSMBShare) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsSMBShare) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsSMBShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DfsSMBShare) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsSMBShare) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsSMBShare) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsSMBShare) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOplocks

`func (o *DfsSMBShare) GetOplocks() bool`

GetOplocks returns the Oplocks field if non-nil, zero value otherwise.

### GetOplocksOk

`func (o *DfsSMBShare) GetOplocksOk() (*bool, bool)`

GetOplocksOk returns a tuple with the Oplocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplocks

`func (o *DfsSMBShare) SetOplocks(v bool)`

SetOplocks sets Oplocks field to given value.

### HasOplocks

`func (o *DfsSMBShare) HasOplocks() bool`

HasOplocks returns a boolean if a field has been set.

### GetRecycled

`func (o *DfsSMBShare) GetRecycled() bool`

GetRecycled returns the Recycled field if non-nil, zero value otherwise.

### GetRecycledOk

`func (o *DfsSMBShare) GetRecycledOk() (*bool, bool)`

GetRecycledOk returns a tuple with the Recycled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycled

`func (o *DfsSMBShare) SetRecycled(v bool)`

SetRecycled sets Recycled field to given value.

### HasRecycled

`func (o *DfsSMBShare) HasRecycled() bool`

HasRecycled returns a boolean if a field has been set.

### GetSecurities

`func (o *DfsSMBShare) GetSecurities() []string`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *DfsSMBShare) GetSecuritiesOk() (*[]string, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *DfsSMBShare) SetSecurities(v []string)`

SetSecurities sets Securities field to given value.

### HasSecurities

`func (o *DfsSMBShare) HasSecurities() bool`

HasSecurities returns a boolean if a field has been set.

### GetStatus

`func (o *DfsSMBShare) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsSMBShare) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsSMBShare) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsSMBShare) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalSnapshotNum

`func (o *DfsSMBShare) GetTotalSnapshotNum() int64`

GetTotalSnapshotNum returns the TotalSnapshotNum field if non-nil, zero value otherwise.

### GetTotalSnapshotNumOk

`func (o *DfsSMBShare) GetTotalSnapshotNumOk() (*int64, bool)`

GetTotalSnapshotNumOk returns a tuple with the TotalSnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSnapshotNum

`func (o *DfsSMBShare) SetTotalSnapshotNum(v int64)`

SetTotalSnapshotNum sets TotalSnapshotNum field to given value.

### HasTotalSnapshotNum

`func (o *DfsSMBShare) HasTotalSnapshotNum() bool`

HasTotalSnapshotNum returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsSMBShare) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsSMBShare) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsSMBShare) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsSMBShare) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetWindowsAcl

`func (o *DfsSMBShare) GetWindowsAcl() bool`

GetWindowsAcl returns the WindowsAcl field if non-nil, zero value otherwise.

### GetWindowsAclOk

`func (o *DfsSMBShare) GetWindowsAclOk() (*bool, bool)`

GetWindowsAclOk returns a tuple with the WindowsAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsAcl

`func (o *DfsSMBShare) SetWindowsAcl(v bool)`

SetWindowsAcl sets WindowsAcl field to given value.

### HasWindowsAcl

`func (o *DfsSMBShare) HasWindowsAcl() bool`

HasWindowsAcl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


