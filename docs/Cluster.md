# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to [**AccessToken**](AccessToken.md) |  | [optional] 
**AccessUrl** | Pointer to **string** |  | [optional] 
**ActionStatus** | Pointer to **string** |  | [optional] 
**AdminAccessTokenNum** | Pointer to **int64** |  | [optional] 
**AdminNetwork** | Pointer to **string** |  | [optional] 
**BootNode** | Pointer to [**BootNode**](BootNode.md) |  | [optional] 
**ContainOsOsd** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**DiskLightingMode** | Pointer to **string** |  | [optional] 
**DownOutInterval** | Pointer to **int64** |  | [optional] 
**ElasticsearchEnabled** | Pointer to **bool** |  | [optional] 
**FsId** | Pointer to **string** |  | [optional] 
**GatewayNetwork** | Pointer to **string** |  | [optional] 
**HostNum** | Pointer to **int64** |  | [optional] 
**HostNumInDc** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsSetPartition** | Pointer to **bool** |  | [optional] 
**IsSetTopology** | Pointer to **bool** |  | [optional] 
**IsSkipPartition** | Pointer to **bool** |  | [optional] 
**JournalPartitionNum** | Pointer to **int64** |  | [optional] 
**Maintained** | Pointer to **bool** |  | [optional] 
**MetadataPartitionNum** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OsGatewayOplogSwitch** | Pointer to **bool** |  | [optional] 
**OspClusterNetwork** | Pointer to **string** |  | [optional] 
**OspGatewayIsConnected** | Pointer to **string** |  | [optional] 
**OspGatewayNetwork** | Pointer to **string** |  | [optional] 
**PrivateNetwork** | Pointer to **string** |  | [optional] 
**PublicNetwork** | Pointer to **string** |  | [optional] 
**SnmpEnabled** | Pointer to **bool** |  | [optional] 
**StatsReservedDays** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *Cluster) GetAccessToken() AccessToken`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Cluster) GetAccessTokenOk() (*AccessToken, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Cluster) SetAccessToken(v AccessToken)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *Cluster) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAccessUrl

`func (o *Cluster) GetAccessUrl() string`

GetAccessUrl returns the AccessUrl field if non-nil, zero value otherwise.

### GetAccessUrlOk

`func (o *Cluster) GetAccessUrlOk() (*string, bool)`

GetAccessUrlOk returns a tuple with the AccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessUrl

`func (o *Cluster) SetAccessUrl(v string)`

SetAccessUrl sets AccessUrl field to given value.

### HasAccessUrl

`func (o *Cluster) HasAccessUrl() bool`

HasAccessUrl returns a boolean if a field has been set.

### GetActionStatus

`func (o *Cluster) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *Cluster) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *Cluster) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *Cluster) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAdminAccessTokenNum

`func (o *Cluster) GetAdminAccessTokenNum() int64`

GetAdminAccessTokenNum returns the AdminAccessTokenNum field if non-nil, zero value otherwise.

### GetAdminAccessTokenNumOk

`func (o *Cluster) GetAdminAccessTokenNumOk() (*int64, bool)`

GetAdminAccessTokenNumOk returns a tuple with the AdminAccessTokenNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAccessTokenNum

`func (o *Cluster) SetAdminAccessTokenNum(v int64)`

SetAdminAccessTokenNum sets AdminAccessTokenNum field to given value.

### HasAdminAccessTokenNum

`func (o *Cluster) HasAdminAccessTokenNum() bool`

HasAdminAccessTokenNum returns a boolean if a field has been set.

### GetAdminNetwork

`func (o *Cluster) GetAdminNetwork() string`

GetAdminNetwork returns the AdminNetwork field if non-nil, zero value otherwise.

### GetAdminNetworkOk

`func (o *Cluster) GetAdminNetworkOk() (*string, bool)`

GetAdminNetworkOk returns a tuple with the AdminNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminNetwork

`func (o *Cluster) SetAdminNetwork(v string)`

SetAdminNetwork sets AdminNetwork field to given value.

### HasAdminNetwork

`func (o *Cluster) HasAdminNetwork() bool`

HasAdminNetwork returns a boolean if a field has been set.

### GetBootNode

`func (o *Cluster) GetBootNode() BootNode`

GetBootNode returns the BootNode field if non-nil, zero value otherwise.

### GetBootNodeOk

`func (o *Cluster) GetBootNodeOk() (*BootNode, bool)`

GetBootNodeOk returns a tuple with the BootNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootNode

`func (o *Cluster) SetBootNode(v BootNode)`

SetBootNode sets BootNode field to given value.

### HasBootNode

`func (o *Cluster) HasBootNode() bool`

HasBootNode returns a boolean if a field has been set.

### GetContainOsOsd

`func (o *Cluster) GetContainOsOsd() bool`

GetContainOsOsd returns the ContainOsOsd field if non-nil, zero value otherwise.

### GetContainOsOsdOk

`func (o *Cluster) GetContainOsOsdOk() (*bool, bool)`

GetContainOsOsdOk returns a tuple with the ContainOsOsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainOsOsd

`func (o *Cluster) SetContainOsOsd(v bool)`

SetContainOsOsd sets ContainOsOsd field to given value.

### HasContainOsOsd

`func (o *Cluster) HasContainOsOsd() bool`

HasContainOsOsd returns a boolean if a field has been set.

### GetCreate

`func (o *Cluster) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Cluster) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Cluster) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Cluster) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDefault

`func (o *Cluster) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Cluster) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Cluster) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Cluster) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDiskLightingMode

`func (o *Cluster) GetDiskLightingMode() string`

GetDiskLightingMode returns the DiskLightingMode field if non-nil, zero value otherwise.

### GetDiskLightingModeOk

`func (o *Cluster) GetDiskLightingModeOk() (*string, bool)`

GetDiskLightingModeOk returns a tuple with the DiskLightingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLightingMode

`func (o *Cluster) SetDiskLightingMode(v string)`

SetDiskLightingMode sets DiskLightingMode field to given value.

### HasDiskLightingMode

`func (o *Cluster) HasDiskLightingMode() bool`

HasDiskLightingMode returns a boolean if a field has been set.

### GetDownOutInterval

`func (o *Cluster) GetDownOutInterval() int64`

GetDownOutInterval returns the DownOutInterval field if non-nil, zero value otherwise.

### GetDownOutIntervalOk

`func (o *Cluster) GetDownOutIntervalOk() (*int64, bool)`

GetDownOutIntervalOk returns a tuple with the DownOutInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownOutInterval

`func (o *Cluster) SetDownOutInterval(v int64)`

SetDownOutInterval sets DownOutInterval field to given value.

### HasDownOutInterval

`func (o *Cluster) HasDownOutInterval() bool`

HasDownOutInterval returns a boolean if a field has been set.

### GetElasticsearchEnabled

`func (o *Cluster) GetElasticsearchEnabled() bool`

GetElasticsearchEnabled returns the ElasticsearchEnabled field if non-nil, zero value otherwise.

### GetElasticsearchEnabledOk

`func (o *Cluster) GetElasticsearchEnabledOk() (*bool, bool)`

GetElasticsearchEnabledOk returns a tuple with the ElasticsearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticsearchEnabled

`func (o *Cluster) SetElasticsearchEnabled(v bool)`

SetElasticsearchEnabled sets ElasticsearchEnabled field to given value.

### HasElasticsearchEnabled

`func (o *Cluster) HasElasticsearchEnabled() bool`

HasElasticsearchEnabled returns a boolean if a field has been set.

### GetFsId

`func (o *Cluster) GetFsId() string`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *Cluster) GetFsIdOk() (*string, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *Cluster) SetFsId(v string)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *Cluster) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetGatewayNetwork

`func (o *Cluster) GetGatewayNetwork() string`

GetGatewayNetwork returns the GatewayNetwork field if non-nil, zero value otherwise.

### GetGatewayNetworkOk

`func (o *Cluster) GetGatewayNetworkOk() (*string, bool)`

GetGatewayNetworkOk returns a tuple with the GatewayNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNetwork

`func (o *Cluster) SetGatewayNetwork(v string)`

SetGatewayNetwork sets GatewayNetwork field to given value.

### HasGatewayNetwork

`func (o *Cluster) HasGatewayNetwork() bool`

HasGatewayNetwork returns a boolean if a field has been set.

### GetHostNum

`func (o *Cluster) GetHostNum() int64`

GetHostNum returns the HostNum field if non-nil, zero value otherwise.

### GetHostNumOk

`func (o *Cluster) GetHostNumOk() (*int64, bool)`

GetHostNumOk returns a tuple with the HostNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNum

`func (o *Cluster) SetHostNum(v int64)`

SetHostNum sets HostNum field to given value.

### HasHostNum

`func (o *Cluster) HasHostNum() bool`

HasHostNum returns a boolean if a field has been set.

### GetHostNumInDc

`func (o *Cluster) GetHostNumInDc() int64`

GetHostNumInDc returns the HostNumInDc field if non-nil, zero value otherwise.

### GetHostNumInDcOk

`func (o *Cluster) GetHostNumInDcOk() (*int64, bool)`

GetHostNumInDcOk returns a tuple with the HostNumInDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNumInDc

`func (o *Cluster) SetHostNumInDc(v int64)`

SetHostNumInDc sets HostNumInDc field to given value.

### HasHostNumInDc

`func (o *Cluster) HasHostNumInDc() bool`

HasHostNumInDc returns a boolean if a field has been set.

### GetId

`func (o *Cluster) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Cluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSetPartition

`func (o *Cluster) GetIsSetPartition() bool`

GetIsSetPartition returns the IsSetPartition field if non-nil, zero value otherwise.

### GetIsSetPartitionOk

`func (o *Cluster) GetIsSetPartitionOk() (*bool, bool)`

GetIsSetPartitionOk returns a tuple with the IsSetPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetPartition

`func (o *Cluster) SetIsSetPartition(v bool)`

SetIsSetPartition sets IsSetPartition field to given value.

### HasIsSetPartition

`func (o *Cluster) HasIsSetPartition() bool`

HasIsSetPartition returns a boolean if a field has been set.

### GetIsSetTopology

`func (o *Cluster) GetIsSetTopology() bool`

GetIsSetTopology returns the IsSetTopology field if non-nil, zero value otherwise.

### GetIsSetTopologyOk

`func (o *Cluster) GetIsSetTopologyOk() (*bool, bool)`

GetIsSetTopologyOk returns a tuple with the IsSetTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetTopology

`func (o *Cluster) SetIsSetTopology(v bool)`

SetIsSetTopology sets IsSetTopology field to given value.

### HasIsSetTopology

`func (o *Cluster) HasIsSetTopology() bool`

HasIsSetTopology returns a boolean if a field has been set.

### GetIsSkipPartition

`func (o *Cluster) GetIsSkipPartition() bool`

GetIsSkipPartition returns the IsSkipPartition field if non-nil, zero value otherwise.

### GetIsSkipPartitionOk

`func (o *Cluster) GetIsSkipPartitionOk() (*bool, bool)`

GetIsSkipPartitionOk returns a tuple with the IsSkipPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSkipPartition

`func (o *Cluster) SetIsSkipPartition(v bool)`

SetIsSkipPartition sets IsSkipPartition field to given value.

### HasIsSkipPartition

`func (o *Cluster) HasIsSkipPartition() bool`

HasIsSkipPartition returns a boolean if a field has been set.

### GetJournalPartitionNum

`func (o *Cluster) GetJournalPartitionNum() int64`

GetJournalPartitionNum returns the JournalPartitionNum field if non-nil, zero value otherwise.

### GetJournalPartitionNumOk

`func (o *Cluster) GetJournalPartitionNumOk() (*int64, bool)`

GetJournalPartitionNumOk returns a tuple with the JournalPartitionNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalPartitionNum

`func (o *Cluster) SetJournalPartitionNum(v int64)`

SetJournalPartitionNum sets JournalPartitionNum field to given value.

### HasJournalPartitionNum

`func (o *Cluster) HasJournalPartitionNum() bool`

HasJournalPartitionNum returns a boolean if a field has been set.

### GetMaintained

`func (o *Cluster) GetMaintained() bool`

GetMaintained returns the Maintained field if non-nil, zero value otherwise.

### GetMaintainedOk

`func (o *Cluster) GetMaintainedOk() (*bool, bool)`

GetMaintainedOk returns a tuple with the Maintained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintained

`func (o *Cluster) SetMaintained(v bool)`

SetMaintained sets Maintained field to given value.

### HasMaintained

`func (o *Cluster) HasMaintained() bool`

HasMaintained returns a boolean if a field has been set.

### GetMetadataPartitionNum

`func (o *Cluster) GetMetadataPartitionNum() int64`

GetMetadataPartitionNum returns the MetadataPartitionNum field if non-nil, zero value otherwise.

### GetMetadataPartitionNumOk

`func (o *Cluster) GetMetadataPartitionNumOk() (*int64, bool)`

GetMetadataPartitionNumOk returns a tuple with the MetadataPartitionNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataPartitionNum

`func (o *Cluster) SetMetadataPartitionNum(v int64)`

SetMetadataPartitionNum sets MetadataPartitionNum field to given value.

### HasMetadataPartitionNum

`func (o *Cluster) HasMetadataPartitionNum() bool`

HasMetadataPartitionNum returns a boolean if a field has been set.

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsGatewayOplogSwitch

`func (o *Cluster) GetOsGatewayOplogSwitch() bool`

GetOsGatewayOplogSwitch returns the OsGatewayOplogSwitch field if non-nil, zero value otherwise.

### GetOsGatewayOplogSwitchOk

`func (o *Cluster) GetOsGatewayOplogSwitchOk() (*bool, bool)`

GetOsGatewayOplogSwitchOk returns a tuple with the OsGatewayOplogSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGatewayOplogSwitch

`func (o *Cluster) SetOsGatewayOplogSwitch(v bool)`

SetOsGatewayOplogSwitch sets OsGatewayOplogSwitch field to given value.

### HasOsGatewayOplogSwitch

`func (o *Cluster) HasOsGatewayOplogSwitch() bool`

HasOsGatewayOplogSwitch returns a boolean if a field has been set.

### GetOspClusterNetwork

`func (o *Cluster) GetOspClusterNetwork() string`

GetOspClusterNetwork returns the OspClusterNetwork field if non-nil, zero value otherwise.

### GetOspClusterNetworkOk

`func (o *Cluster) GetOspClusterNetworkOk() (*string, bool)`

GetOspClusterNetworkOk returns a tuple with the OspClusterNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterNetwork

`func (o *Cluster) SetOspClusterNetwork(v string)`

SetOspClusterNetwork sets OspClusterNetwork field to given value.

### HasOspClusterNetwork

`func (o *Cluster) HasOspClusterNetwork() bool`

HasOspClusterNetwork returns a boolean if a field has been set.

### GetOspGatewayIsConnected

`func (o *Cluster) GetOspGatewayIsConnected() string`

GetOspGatewayIsConnected returns the OspGatewayIsConnected field if non-nil, zero value otherwise.

### GetOspGatewayIsConnectedOk

`func (o *Cluster) GetOspGatewayIsConnectedOk() (*string, bool)`

GetOspGatewayIsConnectedOk returns a tuple with the OspGatewayIsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayIsConnected

`func (o *Cluster) SetOspGatewayIsConnected(v string)`

SetOspGatewayIsConnected sets OspGatewayIsConnected field to given value.

### HasOspGatewayIsConnected

`func (o *Cluster) HasOspGatewayIsConnected() bool`

HasOspGatewayIsConnected returns a boolean if a field has been set.

### GetOspGatewayNetwork

`func (o *Cluster) GetOspGatewayNetwork() string`

GetOspGatewayNetwork returns the OspGatewayNetwork field if non-nil, zero value otherwise.

### GetOspGatewayNetworkOk

`func (o *Cluster) GetOspGatewayNetworkOk() (*string, bool)`

GetOspGatewayNetworkOk returns a tuple with the OspGatewayNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayNetwork

`func (o *Cluster) SetOspGatewayNetwork(v string)`

SetOspGatewayNetwork sets OspGatewayNetwork field to given value.

### HasOspGatewayNetwork

`func (o *Cluster) HasOspGatewayNetwork() bool`

HasOspGatewayNetwork returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *Cluster) GetPrivateNetwork() string`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *Cluster) GetPrivateNetworkOk() (*string, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *Cluster) SetPrivateNetwork(v string)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *Cluster) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.

### GetPublicNetwork

`func (o *Cluster) GetPublicNetwork() string`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *Cluster) GetPublicNetworkOk() (*string, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *Cluster) SetPublicNetwork(v string)`

SetPublicNetwork sets PublicNetwork field to given value.

### HasPublicNetwork

`func (o *Cluster) HasPublicNetwork() bool`

HasPublicNetwork returns a boolean if a field has been set.

### GetSnmpEnabled

`func (o *Cluster) GetSnmpEnabled() bool`

GetSnmpEnabled returns the SnmpEnabled field if non-nil, zero value otherwise.

### GetSnmpEnabledOk

`func (o *Cluster) GetSnmpEnabledOk() (*bool, bool)`

GetSnmpEnabledOk returns a tuple with the SnmpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpEnabled

`func (o *Cluster) SetSnmpEnabled(v bool)`

SetSnmpEnabled sets SnmpEnabled field to given value.

### HasSnmpEnabled

`func (o *Cluster) HasSnmpEnabled() bool`

HasSnmpEnabled returns a boolean if a field has been set.

### GetStatsReservedDays

`func (o *Cluster) GetStatsReservedDays() int64`

GetStatsReservedDays returns the StatsReservedDays field if non-nil, zero value otherwise.

### GetStatsReservedDaysOk

`func (o *Cluster) GetStatsReservedDaysOk() (*int64, bool)`

GetStatsReservedDaysOk returns a tuple with the StatsReservedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsReservedDays

`func (o *Cluster) SetStatsReservedDays(v int64)`

SetStatsReservedDays sets StatsReservedDays field to given value.

### HasStatsReservedDays

`func (o *Cluster) HasStatsReservedDays() bool`

HasStatsReservedDays returns a boolean if a field has been set.

### GetStatus

`func (o *Cluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Cluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Cluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cluster) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Cluster) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *Cluster) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Cluster) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Cluster) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Cluster) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVersion

`func (o *Cluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cluster) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Cluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


