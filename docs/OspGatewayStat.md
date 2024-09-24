# OspGatewayStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllRequests** | Pointer to **int64** |  | [optional] 
**CompactBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**CompactOps** | Pointer to **float64** |  | [optional] 
**CpuUtil** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DeleteLatency** | Pointer to **float64** |  | [optional] 
**DeleteOps** | Pointer to **float64** |  | [optional] 
**DownBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**DownFirstByteLatency** | Pointer to **float64** |  | [optional] 
**DownLatency** | Pointer to **float64** |  | [optional] 
**DownOps** | Pointer to **float64** |  | [optional] 
**ListLatency** | Pointer to **float64** |  | [optional] 
**ListOps** | Pointer to **float64** |  | [optional] 
**MemUsagePercent** | Pointer to **float64** |  | [optional] 
**MergeSpeed** | Pointer to **int64** |  | [optional] 
**ReplicationBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**ReplicationOps** | Pointer to **float64** |  | [optional] 
**Requests** | Pointer to **int64** |  | [optional] 
**SyncDownBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**SyncDownLastByteLatency** | Pointer to **float64** |  | [optional] 
**SyncDownLatency** | Pointer to **float64** |  | [optional] 
**SyncDownOps** | Pointer to **float64** |  | [optional] 
**SyncUpBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**SyncUpLatency** | Pointer to **float64** |  | [optional] 
**SyncUpOps** | Pointer to **float64** |  | [optional] 
**UpBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**UpLastByteLatency** | Pointer to **float64** |  | [optional] 
**UpLatency** | Pointer to **float64** |  | [optional] 
**UpOps** | Pointer to **float64** |  | [optional] 

## Methods

### NewOspGatewayStat

`func NewOspGatewayStat() *OspGatewayStat`

NewOspGatewayStat instantiates a new OspGatewayStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspGatewayStatWithDefaults

`func NewOspGatewayStatWithDefaults() *OspGatewayStat`

NewOspGatewayStatWithDefaults instantiates a new OspGatewayStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllRequests

`func (o *OspGatewayStat) GetAllRequests() int64`

GetAllRequests returns the AllRequests field if non-nil, zero value otherwise.

### GetAllRequestsOk

`func (o *OspGatewayStat) GetAllRequestsOk() (*int64, bool)`

GetAllRequestsOk returns a tuple with the AllRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRequests

`func (o *OspGatewayStat) SetAllRequests(v int64)`

SetAllRequests sets AllRequests field to given value.

### HasAllRequests

`func (o *OspGatewayStat) HasAllRequests() bool`

HasAllRequests returns a boolean if a field has been set.

### GetCompactBandwidthKbyte

`func (o *OspGatewayStat) GetCompactBandwidthKbyte() float64`

GetCompactBandwidthKbyte returns the CompactBandwidthKbyte field if non-nil, zero value otherwise.

### GetCompactBandwidthKbyteOk

`func (o *OspGatewayStat) GetCompactBandwidthKbyteOk() (*float64, bool)`

GetCompactBandwidthKbyteOk returns a tuple with the CompactBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactBandwidthKbyte

`func (o *OspGatewayStat) SetCompactBandwidthKbyte(v float64)`

SetCompactBandwidthKbyte sets CompactBandwidthKbyte field to given value.

### HasCompactBandwidthKbyte

`func (o *OspGatewayStat) HasCompactBandwidthKbyte() bool`

HasCompactBandwidthKbyte returns a boolean if a field has been set.

### GetCompactOps

`func (o *OspGatewayStat) GetCompactOps() float64`

GetCompactOps returns the CompactOps field if non-nil, zero value otherwise.

### GetCompactOpsOk

`func (o *OspGatewayStat) GetCompactOpsOk() (*float64, bool)`

GetCompactOpsOk returns a tuple with the CompactOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactOps

`func (o *OspGatewayStat) SetCompactOps(v float64)`

SetCompactOps sets CompactOps field to given value.

### HasCompactOps

`func (o *OspGatewayStat) HasCompactOps() bool`

HasCompactOps returns a boolean if a field has been set.

### GetCpuUtil

`func (o *OspGatewayStat) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *OspGatewayStat) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *OspGatewayStat) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *OspGatewayStat) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCreate

`func (o *OspGatewayStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OspGatewayStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OspGatewayStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OspGatewayStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDeleteLatency

`func (o *OspGatewayStat) GetDeleteLatency() float64`

GetDeleteLatency returns the DeleteLatency field if non-nil, zero value otherwise.

### GetDeleteLatencyOk

`func (o *OspGatewayStat) GetDeleteLatencyOk() (*float64, bool)`

GetDeleteLatencyOk returns a tuple with the DeleteLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteLatency

`func (o *OspGatewayStat) SetDeleteLatency(v float64)`

SetDeleteLatency sets DeleteLatency field to given value.

### HasDeleteLatency

`func (o *OspGatewayStat) HasDeleteLatency() bool`

HasDeleteLatency returns a boolean if a field has been set.

### GetDeleteOps

`func (o *OspGatewayStat) GetDeleteOps() float64`

GetDeleteOps returns the DeleteOps field if non-nil, zero value otherwise.

### GetDeleteOpsOk

`func (o *OspGatewayStat) GetDeleteOpsOk() (*float64, bool)`

GetDeleteOpsOk returns a tuple with the DeleteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOps

`func (o *OspGatewayStat) SetDeleteOps(v float64)`

SetDeleteOps sets DeleteOps field to given value.

### HasDeleteOps

`func (o *OspGatewayStat) HasDeleteOps() bool`

HasDeleteOps returns a boolean if a field has been set.

### GetDownBandwidthKbyte

`func (o *OspGatewayStat) GetDownBandwidthKbyte() float64`

GetDownBandwidthKbyte returns the DownBandwidthKbyte field if non-nil, zero value otherwise.

### GetDownBandwidthKbyteOk

`func (o *OspGatewayStat) GetDownBandwidthKbyteOk() (*float64, bool)`

GetDownBandwidthKbyteOk returns a tuple with the DownBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownBandwidthKbyte

`func (o *OspGatewayStat) SetDownBandwidthKbyte(v float64)`

SetDownBandwidthKbyte sets DownBandwidthKbyte field to given value.

### HasDownBandwidthKbyte

`func (o *OspGatewayStat) HasDownBandwidthKbyte() bool`

HasDownBandwidthKbyte returns a boolean if a field has been set.

### GetDownFirstByteLatency

`func (o *OspGatewayStat) GetDownFirstByteLatency() float64`

GetDownFirstByteLatency returns the DownFirstByteLatency field if non-nil, zero value otherwise.

### GetDownFirstByteLatencyOk

`func (o *OspGatewayStat) GetDownFirstByteLatencyOk() (*float64, bool)`

GetDownFirstByteLatencyOk returns a tuple with the DownFirstByteLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownFirstByteLatency

`func (o *OspGatewayStat) SetDownFirstByteLatency(v float64)`

SetDownFirstByteLatency sets DownFirstByteLatency field to given value.

### HasDownFirstByteLatency

`func (o *OspGatewayStat) HasDownFirstByteLatency() bool`

HasDownFirstByteLatency returns a boolean if a field has been set.

### GetDownLatency

`func (o *OspGatewayStat) GetDownLatency() float64`

GetDownLatency returns the DownLatency field if non-nil, zero value otherwise.

### GetDownLatencyOk

`func (o *OspGatewayStat) GetDownLatencyOk() (*float64, bool)`

GetDownLatencyOk returns a tuple with the DownLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLatency

`func (o *OspGatewayStat) SetDownLatency(v float64)`

SetDownLatency sets DownLatency field to given value.

### HasDownLatency

`func (o *OspGatewayStat) HasDownLatency() bool`

HasDownLatency returns a boolean if a field has been set.

### GetDownOps

`func (o *OspGatewayStat) GetDownOps() float64`

GetDownOps returns the DownOps field if non-nil, zero value otherwise.

### GetDownOpsOk

`func (o *OspGatewayStat) GetDownOpsOk() (*float64, bool)`

GetDownOpsOk returns a tuple with the DownOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownOps

`func (o *OspGatewayStat) SetDownOps(v float64)`

SetDownOps sets DownOps field to given value.

### HasDownOps

`func (o *OspGatewayStat) HasDownOps() bool`

HasDownOps returns a boolean if a field has been set.

### GetListLatency

`func (o *OspGatewayStat) GetListLatency() float64`

GetListLatency returns the ListLatency field if non-nil, zero value otherwise.

### GetListLatencyOk

`func (o *OspGatewayStat) GetListLatencyOk() (*float64, bool)`

GetListLatencyOk returns a tuple with the ListLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListLatency

`func (o *OspGatewayStat) SetListLatency(v float64)`

SetListLatency sets ListLatency field to given value.

### HasListLatency

`func (o *OspGatewayStat) HasListLatency() bool`

HasListLatency returns a boolean if a field has been set.

### GetListOps

`func (o *OspGatewayStat) GetListOps() float64`

GetListOps returns the ListOps field if non-nil, zero value otherwise.

### GetListOpsOk

`func (o *OspGatewayStat) GetListOpsOk() (*float64, bool)`

GetListOpsOk returns a tuple with the ListOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOps

`func (o *OspGatewayStat) SetListOps(v float64)`

SetListOps sets ListOps field to given value.

### HasListOps

`func (o *OspGatewayStat) HasListOps() bool`

HasListOps returns a boolean if a field has been set.

### GetMemUsagePercent

`func (o *OspGatewayStat) GetMemUsagePercent() float64`

GetMemUsagePercent returns the MemUsagePercent field if non-nil, zero value otherwise.

### GetMemUsagePercentOk

`func (o *OspGatewayStat) GetMemUsagePercentOk() (*float64, bool)`

GetMemUsagePercentOk returns a tuple with the MemUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsagePercent

`func (o *OspGatewayStat) SetMemUsagePercent(v float64)`

SetMemUsagePercent sets MemUsagePercent field to given value.

### HasMemUsagePercent

`func (o *OspGatewayStat) HasMemUsagePercent() bool`

HasMemUsagePercent returns a boolean if a field has been set.

### GetMergeSpeed

`func (o *OspGatewayStat) GetMergeSpeed() int64`

GetMergeSpeed returns the MergeSpeed field if non-nil, zero value otherwise.

### GetMergeSpeedOk

`func (o *OspGatewayStat) GetMergeSpeedOk() (*int64, bool)`

GetMergeSpeedOk returns a tuple with the MergeSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeSpeed

`func (o *OspGatewayStat) SetMergeSpeed(v int64)`

SetMergeSpeed sets MergeSpeed field to given value.

### HasMergeSpeed

`func (o *OspGatewayStat) HasMergeSpeed() bool`

HasMergeSpeed returns a boolean if a field has been set.

### GetReplicationBandwidthKbyte

`func (o *OspGatewayStat) GetReplicationBandwidthKbyte() float64`

GetReplicationBandwidthKbyte returns the ReplicationBandwidthKbyte field if non-nil, zero value otherwise.

### GetReplicationBandwidthKbyteOk

`func (o *OspGatewayStat) GetReplicationBandwidthKbyteOk() (*float64, bool)`

GetReplicationBandwidthKbyteOk returns a tuple with the ReplicationBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBandwidthKbyte

`func (o *OspGatewayStat) SetReplicationBandwidthKbyte(v float64)`

SetReplicationBandwidthKbyte sets ReplicationBandwidthKbyte field to given value.

### HasReplicationBandwidthKbyte

`func (o *OspGatewayStat) HasReplicationBandwidthKbyte() bool`

HasReplicationBandwidthKbyte returns a boolean if a field has been set.

### GetReplicationOps

`func (o *OspGatewayStat) GetReplicationOps() float64`

GetReplicationOps returns the ReplicationOps field if non-nil, zero value otherwise.

### GetReplicationOpsOk

`func (o *OspGatewayStat) GetReplicationOpsOk() (*float64, bool)`

GetReplicationOpsOk returns a tuple with the ReplicationOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationOps

`func (o *OspGatewayStat) SetReplicationOps(v float64)`

SetReplicationOps sets ReplicationOps field to given value.

### HasReplicationOps

`func (o *OspGatewayStat) HasReplicationOps() bool`

HasReplicationOps returns a boolean if a field has been set.

### GetRequests

`func (o *OspGatewayStat) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *OspGatewayStat) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *OspGatewayStat) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *OspGatewayStat) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSyncDownBandwidthKbyte

`func (o *OspGatewayStat) GetSyncDownBandwidthKbyte() float64`

GetSyncDownBandwidthKbyte returns the SyncDownBandwidthKbyte field if non-nil, zero value otherwise.

### GetSyncDownBandwidthKbyteOk

`func (o *OspGatewayStat) GetSyncDownBandwidthKbyteOk() (*float64, bool)`

GetSyncDownBandwidthKbyteOk returns a tuple with the SyncDownBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDownBandwidthKbyte

`func (o *OspGatewayStat) SetSyncDownBandwidthKbyte(v float64)`

SetSyncDownBandwidthKbyte sets SyncDownBandwidthKbyte field to given value.

### HasSyncDownBandwidthKbyte

`func (o *OspGatewayStat) HasSyncDownBandwidthKbyte() bool`

HasSyncDownBandwidthKbyte returns a boolean if a field has been set.

### GetSyncDownLastByteLatency

`func (o *OspGatewayStat) GetSyncDownLastByteLatency() float64`

GetSyncDownLastByteLatency returns the SyncDownLastByteLatency field if non-nil, zero value otherwise.

### GetSyncDownLastByteLatencyOk

`func (o *OspGatewayStat) GetSyncDownLastByteLatencyOk() (*float64, bool)`

GetSyncDownLastByteLatencyOk returns a tuple with the SyncDownLastByteLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDownLastByteLatency

`func (o *OspGatewayStat) SetSyncDownLastByteLatency(v float64)`

SetSyncDownLastByteLatency sets SyncDownLastByteLatency field to given value.

### HasSyncDownLastByteLatency

`func (o *OspGatewayStat) HasSyncDownLastByteLatency() bool`

HasSyncDownLastByteLatency returns a boolean if a field has been set.

### GetSyncDownLatency

`func (o *OspGatewayStat) GetSyncDownLatency() float64`

GetSyncDownLatency returns the SyncDownLatency field if non-nil, zero value otherwise.

### GetSyncDownLatencyOk

`func (o *OspGatewayStat) GetSyncDownLatencyOk() (*float64, bool)`

GetSyncDownLatencyOk returns a tuple with the SyncDownLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDownLatency

`func (o *OspGatewayStat) SetSyncDownLatency(v float64)`

SetSyncDownLatency sets SyncDownLatency field to given value.

### HasSyncDownLatency

`func (o *OspGatewayStat) HasSyncDownLatency() bool`

HasSyncDownLatency returns a boolean if a field has been set.

### GetSyncDownOps

`func (o *OspGatewayStat) GetSyncDownOps() float64`

GetSyncDownOps returns the SyncDownOps field if non-nil, zero value otherwise.

### GetSyncDownOpsOk

`func (o *OspGatewayStat) GetSyncDownOpsOk() (*float64, bool)`

GetSyncDownOpsOk returns a tuple with the SyncDownOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDownOps

`func (o *OspGatewayStat) SetSyncDownOps(v float64)`

SetSyncDownOps sets SyncDownOps field to given value.

### HasSyncDownOps

`func (o *OspGatewayStat) HasSyncDownOps() bool`

HasSyncDownOps returns a boolean if a field has been set.

### GetSyncUpBandwidthKbyte

`func (o *OspGatewayStat) GetSyncUpBandwidthKbyte() float64`

GetSyncUpBandwidthKbyte returns the SyncUpBandwidthKbyte field if non-nil, zero value otherwise.

### GetSyncUpBandwidthKbyteOk

`func (o *OspGatewayStat) GetSyncUpBandwidthKbyteOk() (*float64, bool)`

GetSyncUpBandwidthKbyteOk returns a tuple with the SyncUpBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUpBandwidthKbyte

`func (o *OspGatewayStat) SetSyncUpBandwidthKbyte(v float64)`

SetSyncUpBandwidthKbyte sets SyncUpBandwidthKbyte field to given value.

### HasSyncUpBandwidthKbyte

`func (o *OspGatewayStat) HasSyncUpBandwidthKbyte() bool`

HasSyncUpBandwidthKbyte returns a boolean if a field has been set.

### GetSyncUpLatency

`func (o *OspGatewayStat) GetSyncUpLatency() float64`

GetSyncUpLatency returns the SyncUpLatency field if non-nil, zero value otherwise.

### GetSyncUpLatencyOk

`func (o *OspGatewayStat) GetSyncUpLatencyOk() (*float64, bool)`

GetSyncUpLatencyOk returns a tuple with the SyncUpLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUpLatency

`func (o *OspGatewayStat) SetSyncUpLatency(v float64)`

SetSyncUpLatency sets SyncUpLatency field to given value.

### HasSyncUpLatency

`func (o *OspGatewayStat) HasSyncUpLatency() bool`

HasSyncUpLatency returns a boolean if a field has been set.

### GetSyncUpOps

`func (o *OspGatewayStat) GetSyncUpOps() float64`

GetSyncUpOps returns the SyncUpOps field if non-nil, zero value otherwise.

### GetSyncUpOpsOk

`func (o *OspGatewayStat) GetSyncUpOpsOk() (*float64, bool)`

GetSyncUpOpsOk returns a tuple with the SyncUpOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUpOps

`func (o *OspGatewayStat) SetSyncUpOps(v float64)`

SetSyncUpOps sets SyncUpOps field to given value.

### HasSyncUpOps

`func (o *OspGatewayStat) HasSyncUpOps() bool`

HasSyncUpOps returns a boolean if a field has been set.

### GetUpBandwidthKbyte

`func (o *OspGatewayStat) GetUpBandwidthKbyte() float64`

GetUpBandwidthKbyte returns the UpBandwidthKbyte field if non-nil, zero value otherwise.

### GetUpBandwidthKbyteOk

`func (o *OspGatewayStat) GetUpBandwidthKbyteOk() (*float64, bool)`

GetUpBandwidthKbyteOk returns a tuple with the UpBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpBandwidthKbyte

`func (o *OspGatewayStat) SetUpBandwidthKbyte(v float64)`

SetUpBandwidthKbyte sets UpBandwidthKbyte field to given value.

### HasUpBandwidthKbyte

`func (o *OspGatewayStat) HasUpBandwidthKbyte() bool`

HasUpBandwidthKbyte returns a boolean if a field has been set.

### GetUpLastByteLatency

`func (o *OspGatewayStat) GetUpLastByteLatency() float64`

GetUpLastByteLatency returns the UpLastByteLatency field if non-nil, zero value otherwise.

### GetUpLastByteLatencyOk

`func (o *OspGatewayStat) GetUpLastByteLatencyOk() (*float64, bool)`

GetUpLastByteLatencyOk returns a tuple with the UpLastByteLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLastByteLatency

`func (o *OspGatewayStat) SetUpLastByteLatency(v float64)`

SetUpLastByteLatency sets UpLastByteLatency field to given value.

### HasUpLastByteLatency

`func (o *OspGatewayStat) HasUpLastByteLatency() bool`

HasUpLastByteLatency returns a boolean if a field has been set.

### GetUpLatency

`func (o *OspGatewayStat) GetUpLatency() float64`

GetUpLatency returns the UpLatency field if non-nil, zero value otherwise.

### GetUpLatencyOk

`func (o *OspGatewayStat) GetUpLatencyOk() (*float64, bool)`

GetUpLatencyOk returns a tuple with the UpLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLatency

`func (o *OspGatewayStat) SetUpLatency(v float64)`

SetUpLatency sets UpLatency field to given value.

### HasUpLatency

`func (o *OspGatewayStat) HasUpLatency() bool`

HasUpLatency returns a boolean if a field has been set.

### GetUpOps

`func (o *OspGatewayStat) GetUpOps() float64`

GetUpOps returns the UpOps field if non-nil, zero value otherwise.

### GetUpOpsOk

`func (o *OspGatewayStat) GetUpOpsOk() (*float64, bool)`

GetUpOpsOk returns a tuple with the UpOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpOps

`func (o *OspGatewayStat) SetUpOps(v float64)`

SetUpOps sets UpOps field to given value.

### HasUpOps

`func (o *OspGatewayStat) HasUpOps() bool`

HasUpOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


