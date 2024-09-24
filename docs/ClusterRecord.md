# ClusterRecord

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
**Samples** | Pointer to [**[]ClusterStat**](ClusterStat.md) |  | [optional] 

## Methods

### NewClusterRecord

`func NewClusterRecord() *ClusterRecord`

NewClusterRecord instantiates a new ClusterRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRecordWithDefaults

`func NewClusterRecordWithDefaults() *ClusterRecord`

NewClusterRecordWithDefaults instantiates a new ClusterRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *ClusterRecord) GetAccessToken() AccessToken`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ClusterRecord) GetAccessTokenOk() (*AccessToken, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ClusterRecord) SetAccessToken(v AccessToken)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *ClusterRecord) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAccessUrl

`func (o *ClusterRecord) GetAccessUrl() string`

GetAccessUrl returns the AccessUrl field if non-nil, zero value otherwise.

### GetAccessUrlOk

`func (o *ClusterRecord) GetAccessUrlOk() (*string, bool)`

GetAccessUrlOk returns a tuple with the AccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessUrl

`func (o *ClusterRecord) SetAccessUrl(v string)`

SetAccessUrl sets AccessUrl field to given value.

### HasAccessUrl

`func (o *ClusterRecord) HasAccessUrl() bool`

HasAccessUrl returns a boolean if a field has been set.

### GetActionStatus

`func (o *ClusterRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *ClusterRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *ClusterRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *ClusterRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAdminAccessTokenNum

`func (o *ClusterRecord) GetAdminAccessTokenNum() int64`

GetAdminAccessTokenNum returns the AdminAccessTokenNum field if non-nil, zero value otherwise.

### GetAdminAccessTokenNumOk

`func (o *ClusterRecord) GetAdminAccessTokenNumOk() (*int64, bool)`

GetAdminAccessTokenNumOk returns a tuple with the AdminAccessTokenNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAccessTokenNum

`func (o *ClusterRecord) SetAdminAccessTokenNum(v int64)`

SetAdminAccessTokenNum sets AdminAccessTokenNum field to given value.

### HasAdminAccessTokenNum

`func (o *ClusterRecord) HasAdminAccessTokenNum() bool`

HasAdminAccessTokenNum returns a boolean if a field has been set.

### GetAdminNetwork

`func (o *ClusterRecord) GetAdminNetwork() string`

GetAdminNetwork returns the AdminNetwork field if non-nil, zero value otherwise.

### GetAdminNetworkOk

`func (o *ClusterRecord) GetAdminNetworkOk() (*string, bool)`

GetAdminNetworkOk returns a tuple with the AdminNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminNetwork

`func (o *ClusterRecord) SetAdminNetwork(v string)`

SetAdminNetwork sets AdminNetwork field to given value.

### HasAdminNetwork

`func (o *ClusterRecord) HasAdminNetwork() bool`

HasAdminNetwork returns a boolean if a field has been set.

### GetBootNode

`func (o *ClusterRecord) GetBootNode() BootNode`

GetBootNode returns the BootNode field if non-nil, zero value otherwise.

### GetBootNodeOk

`func (o *ClusterRecord) GetBootNodeOk() (*BootNode, bool)`

GetBootNodeOk returns a tuple with the BootNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootNode

`func (o *ClusterRecord) SetBootNode(v BootNode)`

SetBootNode sets BootNode field to given value.

### HasBootNode

`func (o *ClusterRecord) HasBootNode() bool`

HasBootNode returns a boolean if a field has been set.

### GetContainOsOsd

`func (o *ClusterRecord) GetContainOsOsd() bool`

GetContainOsOsd returns the ContainOsOsd field if non-nil, zero value otherwise.

### GetContainOsOsdOk

`func (o *ClusterRecord) GetContainOsOsdOk() (*bool, bool)`

GetContainOsOsdOk returns a tuple with the ContainOsOsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainOsOsd

`func (o *ClusterRecord) SetContainOsOsd(v bool)`

SetContainOsOsd sets ContainOsOsd field to given value.

### HasContainOsOsd

`func (o *ClusterRecord) HasContainOsOsd() bool`

HasContainOsOsd returns a boolean if a field has been set.

### GetCreate

`func (o *ClusterRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ClusterRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ClusterRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ClusterRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDefault

`func (o *ClusterRecord) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ClusterRecord) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ClusterRecord) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ClusterRecord) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDiskLightingMode

`func (o *ClusterRecord) GetDiskLightingMode() string`

GetDiskLightingMode returns the DiskLightingMode field if non-nil, zero value otherwise.

### GetDiskLightingModeOk

`func (o *ClusterRecord) GetDiskLightingModeOk() (*string, bool)`

GetDiskLightingModeOk returns a tuple with the DiskLightingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLightingMode

`func (o *ClusterRecord) SetDiskLightingMode(v string)`

SetDiskLightingMode sets DiskLightingMode field to given value.

### HasDiskLightingMode

`func (o *ClusterRecord) HasDiskLightingMode() bool`

HasDiskLightingMode returns a boolean if a field has been set.

### GetDownOutInterval

`func (o *ClusterRecord) GetDownOutInterval() int64`

GetDownOutInterval returns the DownOutInterval field if non-nil, zero value otherwise.

### GetDownOutIntervalOk

`func (o *ClusterRecord) GetDownOutIntervalOk() (*int64, bool)`

GetDownOutIntervalOk returns a tuple with the DownOutInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownOutInterval

`func (o *ClusterRecord) SetDownOutInterval(v int64)`

SetDownOutInterval sets DownOutInterval field to given value.

### HasDownOutInterval

`func (o *ClusterRecord) HasDownOutInterval() bool`

HasDownOutInterval returns a boolean if a field has been set.

### GetElasticsearchEnabled

`func (o *ClusterRecord) GetElasticsearchEnabled() bool`

GetElasticsearchEnabled returns the ElasticsearchEnabled field if non-nil, zero value otherwise.

### GetElasticsearchEnabledOk

`func (o *ClusterRecord) GetElasticsearchEnabledOk() (*bool, bool)`

GetElasticsearchEnabledOk returns a tuple with the ElasticsearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticsearchEnabled

`func (o *ClusterRecord) SetElasticsearchEnabled(v bool)`

SetElasticsearchEnabled sets ElasticsearchEnabled field to given value.

### HasElasticsearchEnabled

`func (o *ClusterRecord) HasElasticsearchEnabled() bool`

HasElasticsearchEnabled returns a boolean if a field has been set.

### GetFsId

`func (o *ClusterRecord) GetFsId() string`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *ClusterRecord) GetFsIdOk() (*string, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *ClusterRecord) SetFsId(v string)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *ClusterRecord) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetGatewayNetwork

`func (o *ClusterRecord) GetGatewayNetwork() string`

GetGatewayNetwork returns the GatewayNetwork field if non-nil, zero value otherwise.

### GetGatewayNetworkOk

`func (o *ClusterRecord) GetGatewayNetworkOk() (*string, bool)`

GetGatewayNetworkOk returns a tuple with the GatewayNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNetwork

`func (o *ClusterRecord) SetGatewayNetwork(v string)`

SetGatewayNetwork sets GatewayNetwork field to given value.

### HasGatewayNetwork

`func (o *ClusterRecord) HasGatewayNetwork() bool`

HasGatewayNetwork returns a boolean if a field has been set.

### GetHostNum

`func (o *ClusterRecord) GetHostNum() int64`

GetHostNum returns the HostNum field if non-nil, zero value otherwise.

### GetHostNumOk

`func (o *ClusterRecord) GetHostNumOk() (*int64, bool)`

GetHostNumOk returns a tuple with the HostNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNum

`func (o *ClusterRecord) SetHostNum(v int64)`

SetHostNum sets HostNum field to given value.

### HasHostNum

`func (o *ClusterRecord) HasHostNum() bool`

HasHostNum returns a boolean if a field has been set.

### GetHostNumInDc

`func (o *ClusterRecord) GetHostNumInDc() int64`

GetHostNumInDc returns the HostNumInDc field if non-nil, zero value otherwise.

### GetHostNumInDcOk

`func (o *ClusterRecord) GetHostNumInDcOk() (*int64, bool)`

GetHostNumInDcOk returns a tuple with the HostNumInDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNumInDc

`func (o *ClusterRecord) SetHostNumInDc(v int64)`

SetHostNumInDc sets HostNumInDc field to given value.

### HasHostNumInDc

`func (o *ClusterRecord) HasHostNumInDc() bool`

HasHostNumInDc returns a boolean if a field has been set.

### GetId

`func (o *ClusterRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSetPartition

`func (o *ClusterRecord) GetIsSetPartition() bool`

GetIsSetPartition returns the IsSetPartition field if non-nil, zero value otherwise.

### GetIsSetPartitionOk

`func (o *ClusterRecord) GetIsSetPartitionOk() (*bool, bool)`

GetIsSetPartitionOk returns a tuple with the IsSetPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetPartition

`func (o *ClusterRecord) SetIsSetPartition(v bool)`

SetIsSetPartition sets IsSetPartition field to given value.

### HasIsSetPartition

`func (o *ClusterRecord) HasIsSetPartition() bool`

HasIsSetPartition returns a boolean if a field has been set.

### GetIsSetTopology

`func (o *ClusterRecord) GetIsSetTopology() bool`

GetIsSetTopology returns the IsSetTopology field if non-nil, zero value otherwise.

### GetIsSetTopologyOk

`func (o *ClusterRecord) GetIsSetTopologyOk() (*bool, bool)`

GetIsSetTopologyOk returns a tuple with the IsSetTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetTopology

`func (o *ClusterRecord) SetIsSetTopology(v bool)`

SetIsSetTopology sets IsSetTopology field to given value.

### HasIsSetTopology

`func (o *ClusterRecord) HasIsSetTopology() bool`

HasIsSetTopology returns a boolean if a field has been set.

### GetIsSkipPartition

`func (o *ClusterRecord) GetIsSkipPartition() bool`

GetIsSkipPartition returns the IsSkipPartition field if non-nil, zero value otherwise.

### GetIsSkipPartitionOk

`func (o *ClusterRecord) GetIsSkipPartitionOk() (*bool, bool)`

GetIsSkipPartitionOk returns a tuple with the IsSkipPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSkipPartition

`func (o *ClusterRecord) SetIsSkipPartition(v bool)`

SetIsSkipPartition sets IsSkipPartition field to given value.

### HasIsSkipPartition

`func (o *ClusterRecord) HasIsSkipPartition() bool`

HasIsSkipPartition returns a boolean if a field has been set.

### GetJournalPartitionNum

`func (o *ClusterRecord) GetJournalPartitionNum() int64`

GetJournalPartitionNum returns the JournalPartitionNum field if non-nil, zero value otherwise.

### GetJournalPartitionNumOk

`func (o *ClusterRecord) GetJournalPartitionNumOk() (*int64, bool)`

GetJournalPartitionNumOk returns a tuple with the JournalPartitionNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalPartitionNum

`func (o *ClusterRecord) SetJournalPartitionNum(v int64)`

SetJournalPartitionNum sets JournalPartitionNum field to given value.

### HasJournalPartitionNum

`func (o *ClusterRecord) HasJournalPartitionNum() bool`

HasJournalPartitionNum returns a boolean if a field has been set.

### GetMaintained

`func (o *ClusterRecord) GetMaintained() bool`

GetMaintained returns the Maintained field if non-nil, zero value otherwise.

### GetMaintainedOk

`func (o *ClusterRecord) GetMaintainedOk() (*bool, bool)`

GetMaintainedOk returns a tuple with the Maintained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintained

`func (o *ClusterRecord) SetMaintained(v bool)`

SetMaintained sets Maintained field to given value.

### HasMaintained

`func (o *ClusterRecord) HasMaintained() bool`

HasMaintained returns a boolean if a field has been set.

### GetMetadataPartitionNum

`func (o *ClusterRecord) GetMetadataPartitionNum() int64`

GetMetadataPartitionNum returns the MetadataPartitionNum field if non-nil, zero value otherwise.

### GetMetadataPartitionNumOk

`func (o *ClusterRecord) GetMetadataPartitionNumOk() (*int64, bool)`

GetMetadataPartitionNumOk returns a tuple with the MetadataPartitionNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataPartitionNum

`func (o *ClusterRecord) SetMetadataPartitionNum(v int64)`

SetMetadataPartitionNum sets MetadataPartitionNum field to given value.

### HasMetadataPartitionNum

`func (o *ClusterRecord) HasMetadataPartitionNum() bool`

HasMetadataPartitionNum returns a boolean if a field has been set.

### GetName

`func (o *ClusterRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsGatewayOplogSwitch

`func (o *ClusterRecord) GetOsGatewayOplogSwitch() bool`

GetOsGatewayOplogSwitch returns the OsGatewayOplogSwitch field if non-nil, zero value otherwise.

### GetOsGatewayOplogSwitchOk

`func (o *ClusterRecord) GetOsGatewayOplogSwitchOk() (*bool, bool)`

GetOsGatewayOplogSwitchOk returns a tuple with the OsGatewayOplogSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGatewayOplogSwitch

`func (o *ClusterRecord) SetOsGatewayOplogSwitch(v bool)`

SetOsGatewayOplogSwitch sets OsGatewayOplogSwitch field to given value.

### HasOsGatewayOplogSwitch

`func (o *ClusterRecord) HasOsGatewayOplogSwitch() bool`

HasOsGatewayOplogSwitch returns a boolean if a field has been set.

### GetOspClusterNetwork

`func (o *ClusterRecord) GetOspClusterNetwork() string`

GetOspClusterNetwork returns the OspClusterNetwork field if non-nil, zero value otherwise.

### GetOspClusterNetworkOk

`func (o *ClusterRecord) GetOspClusterNetworkOk() (*string, bool)`

GetOspClusterNetworkOk returns a tuple with the OspClusterNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspClusterNetwork

`func (o *ClusterRecord) SetOspClusterNetwork(v string)`

SetOspClusterNetwork sets OspClusterNetwork field to given value.

### HasOspClusterNetwork

`func (o *ClusterRecord) HasOspClusterNetwork() bool`

HasOspClusterNetwork returns a boolean if a field has been set.

### GetOspGatewayIsConnected

`func (o *ClusterRecord) GetOspGatewayIsConnected() string`

GetOspGatewayIsConnected returns the OspGatewayIsConnected field if non-nil, zero value otherwise.

### GetOspGatewayIsConnectedOk

`func (o *ClusterRecord) GetOspGatewayIsConnectedOk() (*string, bool)`

GetOspGatewayIsConnectedOk returns a tuple with the OspGatewayIsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayIsConnected

`func (o *ClusterRecord) SetOspGatewayIsConnected(v string)`

SetOspGatewayIsConnected sets OspGatewayIsConnected field to given value.

### HasOspGatewayIsConnected

`func (o *ClusterRecord) HasOspGatewayIsConnected() bool`

HasOspGatewayIsConnected returns a boolean if a field has been set.

### GetOspGatewayNetwork

`func (o *ClusterRecord) GetOspGatewayNetwork() string`

GetOspGatewayNetwork returns the OspGatewayNetwork field if non-nil, zero value otherwise.

### GetOspGatewayNetworkOk

`func (o *ClusterRecord) GetOspGatewayNetworkOk() (*string, bool)`

GetOspGatewayNetworkOk returns a tuple with the OspGatewayNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspGatewayNetwork

`func (o *ClusterRecord) SetOspGatewayNetwork(v string)`

SetOspGatewayNetwork sets OspGatewayNetwork field to given value.

### HasOspGatewayNetwork

`func (o *ClusterRecord) HasOspGatewayNetwork() bool`

HasOspGatewayNetwork returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *ClusterRecord) GetPrivateNetwork() string`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *ClusterRecord) GetPrivateNetworkOk() (*string, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *ClusterRecord) SetPrivateNetwork(v string)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *ClusterRecord) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.

### GetPublicNetwork

`func (o *ClusterRecord) GetPublicNetwork() string`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *ClusterRecord) GetPublicNetworkOk() (*string, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *ClusterRecord) SetPublicNetwork(v string)`

SetPublicNetwork sets PublicNetwork field to given value.

### HasPublicNetwork

`func (o *ClusterRecord) HasPublicNetwork() bool`

HasPublicNetwork returns a boolean if a field has been set.

### GetSnmpEnabled

`func (o *ClusterRecord) GetSnmpEnabled() bool`

GetSnmpEnabled returns the SnmpEnabled field if non-nil, zero value otherwise.

### GetSnmpEnabledOk

`func (o *ClusterRecord) GetSnmpEnabledOk() (*bool, bool)`

GetSnmpEnabledOk returns a tuple with the SnmpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpEnabled

`func (o *ClusterRecord) SetSnmpEnabled(v bool)`

SetSnmpEnabled sets SnmpEnabled field to given value.

### HasSnmpEnabled

`func (o *ClusterRecord) HasSnmpEnabled() bool`

HasSnmpEnabled returns a boolean if a field has been set.

### GetStatsReservedDays

`func (o *ClusterRecord) GetStatsReservedDays() int64`

GetStatsReservedDays returns the StatsReservedDays field if non-nil, zero value otherwise.

### GetStatsReservedDaysOk

`func (o *ClusterRecord) GetStatsReservedDaysOk() (*int64, bool)`

GetStatsReservedDaysOk returns a tuple with the StatsReservedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsReservedDays

`func (o *ClusterRecord) SetStatsReservedDays(v int64)`

SetStatsReservedDays sets StatsReservedDays field to given value.

### HasStatsReservedDays

`func (o *ClusterRecord) HasStatsReservedDays() bool`

HasStatsReservedDays returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ClusterRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *ClusterRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ClusterRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ClusterRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ClusterRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterRecord) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterRecord) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterRecord) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterRecord) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSamples

`func (o *ClusterRecord) GetSamples() []ClusterStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ClusterRecord) GetSamplesOk() (*[]ClusterStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ClusterRecord) SetSamples(v []ClusterStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ClusterRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


