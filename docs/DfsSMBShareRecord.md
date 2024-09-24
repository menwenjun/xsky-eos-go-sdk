# DfsSMBShareRecord

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
**Samples** | Pointer to [**[]DfsSMBShareStat**](DfsSMBShareStat.md) |  | [optional] 

## Methods

### NewDfsSMBShareRecord

`func NewDfsSMBShareRecord() *DfsSMBShareRecord`

NewDfsSMBShareRecord instantiates a new DfsSMBShareRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBShareRecordWithDefaults

`func NewDfsSMBShareRecordWithDefaults() *DfsSMBShareRecord`

NewDfsSMBShareRecordWithDefaults instantiates a new DfsSMBShareRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessBasedShareEnum

`func (o *DfsSMBShareRecord) GetAccessBasedShareEnum() bool`

GetAccessBasedShareEnum returns the AccessBasedShareEnum field if non-nil, zero value otherwise.

### GetAccessBasedShareEnumOk

`func (o *DfsSMBShareRecord) GetAccessBasedShareEnumOk() (*bool, bool)`

GetAccessBasedShareEnumOk returns a tuple with the AccessBasedShareEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessBasedShareEnum

`func (o *DfsSMBShareRecord) SetAccessBasedShareEnum(v bool)`

SetAccessBasedShareEnum sets AccessBasedShareEnum field to given value.

### HasAccessBasedShareEnum

`func (o *DfsSMBShareRecord) HasAccessBasedShareEnum() bool`

HasAccessBasedShareEnum returns a boolean if a field has been set.

### GetAclInherited

`func (o *DfsSMBShareRecord) GetAclInherited() bool`

GetAclInherited returns the AclInherited field if non-nil, zero value otherwise.

### GetAclInheritedOk

`func (o *DfsSMBShareRecord) GetAclInheritedOk() (*bool, bool)`

GetAclInheritedOk returns a tuple with the AclInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclInherited

`func (o *DfsSMBShareRecord) SetAclInherited(v bool)`

SetAclInherited sets AclInherited field to given value.

### HasAclInherited

`func (o *DfsSMBShareRecord) HasAclInherited() bool`

HasAclInherited returns a boolean if a field has been set.

### GetActionStatus

`func (o *DfsSMBShareRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsSMBShareRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsSMBShareRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsSMBShareRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *DfsSMBShareRecord) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *DfsSMBShareRecord) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *DfsSMBShareRecord) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *DfsSMBShareRecord) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetCluster

`func (o *DfsSMBShareRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsSMBShareRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsSMBShareRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsSMBShareRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsSMBShareRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsSMBShareRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsSMBShareRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsSMBShareRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *DfsSMBShareRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsSMBShareRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsSMBShareRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsSMBShareRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsSMBShareRecord) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsSMBShareRecord) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsSMBShareRecord) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsSMBShareRecord) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetDfsPath

`func (o *DfsSMBShareRecord) GetDfsPath() DfsPathNestview`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DfsSMBShareRecord) GetDfsPathOk() (*DfsPathNestview, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DfsSMBShareRecord) SetDfsPath(v DfsPathNestview)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DfsSMBShareRecord) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsSMBShareRecord) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsSMBShareRecord) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsSMBShareRecord) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsSMBShareRecord) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetHostsAllow

`func (o *DfsSMBShareRecord) GetHostsAllow() string`

GetHostsAllow returns the HostsAllow field if non-nil, zero value otherwise.

### GetHostsAllowOk

`func (o *DfsSMBShareRecord) GetHostsAllowOk() (*string, bool)`

GetHostsAllowOk returns a tuple with the HostsAllow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsAllow

`func (o *DfsSMBShareRecord) SetHostsAllow(v string)`

SetHostsAllow sets HostsAllow field to given value.

### HasHostsAllow

`func (o *DfsSMBShareRecord) HasHostsAllow() bool`

HasHostsAllow returns a boolean if a field has been set.

### GetId

`func (o *DfsSMBShareRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsSMBShareRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsSMBShareRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsSMBShareRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DfsSMBShareRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsSMBShareRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsSMBShareRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsSMBShareRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOplocks

`func (o *DfsSMBShareRecord) GetOplocks() bool`

GetOplocks returns the Oplocks field if non-nil, zero value otherwise.

### GetOplocksOk

`func (o *DfsSMBShareRecord) GetOplocksOk() (*bool, bool)`

GetOplocksOk returns a tuple with the Oplocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplocks

`func (o *DfsSMBShareRecord) SetOplocks(v bool)`

SetOplocks sets Oplocks field to given value.

### HasOplocks

`func (o *DfsSMBShareRecord) HasOplocks() bool`

HasOplocks returns a boolean if a field has been set.

### GetRecycled

`func (o *DfsSMBShareRecord) GetRecycled() bool`

GetRecycled returns the Recycled field if non-nil, zero value otherwise.

### GetRecycledOk

`func (o *DfsSMBShareRecord) GetRecycledOk() (*bool, bool)`

GetRecycledOk returns a tuple with the Recycled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycled

`func (o *DfsSMBShareRecord) SetRecycled(v bool)`

SetRecycled sets Recycled field to given value.

### HasRecycled

`func (o *DfsSMBShareRecord) HasRecycled() bool`

HasRecycled returns a boolean if a field has been set.

### GetSecurities

`func (o *DfsSMBShareRecord) GetSecurities() []string`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *DfsSMBShareRecord) GetSecuritiesOk() (*[]string, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *DfsSMBShareRecord) SetSecurities(v []string)`

SetSecurities sets Securities field to given value.

### HasSecurities

`func (o *DfsSMBShareRecord) HasSecurities() bool`

HasSecurities returns a boolean if a field has been set.

### GetStatus

`func (o *DfsSMBShareRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsSMBShareRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsSMBShareRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsSMBShareRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalSnapshotNum

`func (o *DfsSMBShareRecord) GetTotalSnapshotNum() int64`

GetTotalSnapshotNum returns the TotalSnapshotNum field if non-nil, zero value otherwise.

### GetTotalSnapshotNumOk

`func (o *DfsSMBShareRecord) GetTotalSnapshotNumOk() (*int64, bool)`

GetTotalSnapshotNumOk returns a tuple with the TotalSnapshotNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSnapshotNum

`func (o *DfsSMBShareRecord) SetTotalSnapshotNum(v int64)`

SetTotalSnapshotNum sets TotalSnapshotNum field to given value.

### HasTotalSnapshotNum

`func (o *DfsSMBShareRecord) HasTotalSnapshotNum() bool`

HasTotalSnapshotNum returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsSMBShareRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsSMBShareRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsSMBShareRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsSMBShareRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetWindowsAcl

`func (o *DfsSMBShareRecord) GetWindowsAcl() bool`

GetWindowsAcl returns the WindowsAcl field if non-nil, zero value otherwise.

### GetWindowsAclOk

`func (o *DfsSMBShareRecord) GetWindowsAclOk() (*bool, bool)`

GetWindowsAclOk returns a tuple with the WindowsAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsAcl

`func (o *DfsSMBShareRecord) SetWindowsAcl(v bool)`

SetWindowsAcl sets WindowsAcl field to given value.

### HasWindowsAcl

`func (o *DfsSMBShareRecord) HasWindowsAcl() bool`

HasWindowsAcl returns a boolean if a field has been set.

### GetSamples

`func (o *DfsSMBShareRecord) GetSamples() []DfsSMBShareStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DfsSMBShareRecord) GetSamplesOk() (*[]DfsSMBShareStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DfsSMBShareRecord) SetSamples(v []DfsSMBShareStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *DfsSMBShareRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


