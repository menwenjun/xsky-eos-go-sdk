# NFSGatewayStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllRequests** | Pointer to **int64** |  | [optional] 
**CompactBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**CompactOps** | Pointer to **float64** |  | [optional] 
**CpuUtil** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DeleteLatencyUs** | Pointer to **float64** |  | [optional] 
**DeleteOps** | Pointer to **float64** |  | [optional] 
**DownBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**DownFirstByteLatencyUs** | Pointer to **float64** |  | [optional] 
**DownLatencyUs** | Pointer to **float64** |  | [optional] 
**DownOps** | Pointer to **float64** |  | [optional] 
**ListLatencyUs** | Pointer to **float64** |  | [optional] 
**ListOps** | Pointer to **float64** |  | [optional] 
**MemUsagePercent** | Pointer to **float64** |  | [optional] 
**MergeSpeed** | Pointer to **int64** |  | [optional] 
**ReplicationBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**ReplicationOps** | Pointer to **float64** |  | [optional] 
**Requests** | Pointer to **int64** |  | [optional] 
**SyncDownBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**SyncDownLastByteLatencyUs** | Pointer to **float64** |  | [optional] 
**SyncDownLatencyUs** | Pointer to **float64** |  | [optional] 
**SyncDownOps** | Pointer to **float64** |  | [optional] 
**SyncUpBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**SyncUpLatencyUs** | Pointer to **float64** |  | [optional] 
**SyncUpOps** | Pointer to **float64** |  | [optional] 
**UpBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**UpLastByteLatencyUs** | Pointer to **float64** |  | [optional] 
**UpLatencyUs** | Pointer to **float64** |  | [optional] 
**UpOps** | Pointer to **float64** |  | [optional] 

## Methods

### NewNFSGatewayStat

`func NewNFSGatewayStat() *NFSGatewayStat`

NewNFSGatewayStat instantiates a new NFSGatewayStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFSGatewayStatWithDefaults

`func NewNFSGatewayStatWithDefaults() *NFSGatewayStat`

NewNFSGatewayStatWithDefaults instantiates a new NFSGatewayStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllRequests

`func (o *NFSGatewayStat) GetAllRequests() int64`

GetAllRequests returns the AllRequests field if non-nil, zero value otherwise.

### GetAllRequestsOk

`func (o *NFSGatewayStat) GetAllRequestsOk() (*int64, bool)`

GetAllRequestsOk returns a tuple with the AllRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRequests

`func (o *NFSGatewayStat) SetAllRequests(v int64)`

SetAllRequests sets AllRequests field to given value.

### HasAllRequests

`func (o *NFSGatewayStat) HasAllRequests() bool`

HasAllRequests returns a boolean if a field has been set.

### GetCompactBandwidthKbyte

`func (o *NFSGatewayStat) GetCompactBandwidthKbyte() float64`

GetCompactBandwidthKbyte returns the CompactBandwidthKbyte field if non-nil, zero value otherwise.

### GetCompactBandwidthKbyteOk

`func (o *NFSGatewayStat) GetCompactBandwidthKbyteOk() (*float64, bool)`

GetCompactBandwidthKbyteOk returns a tuple with the CompactBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactBandwidthKbyte

`func (o *NFSGatewayStat) SetCompactBandwidthKbyte(v float64)`

SetCompactBandwidthKbyte sets CompactBandwidthKbyte field to given value.

### HasCompactBandwidthKbyte

`func (o *NFSGatewayStat) HasCompactBandwidthKbyte() bool`

HasCompactBandwidthKbyte returns a boolean if a field has been set.

### GetCompactOps

`func (o *NFSGatewayStat) GetCompactOps() float64`

GetCompactOps returns the CompactOps field if non-nil, zero value otherwise.

### GetCompactOpsOk

`func (o *NFSGatewayStat) GetCompactOpsOk() (*float64, bool)`

GetCompactOpsOk returns a tuple with the CompactOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactOps

`func (o *NFSGatewayStat) SetCompactOps(v float64)`

SetCompactOps sets CompactOps field to given value.

### HasCompactOps

`func (o *NFSGatewayStat) HasCompactOps() bool`

HasCompactOps returns a boolean if a field has been set.

### GetCpuUtil

`func (o *NFSGatewayStat) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *NFSGatewayStat) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *NFSGatewayStat) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *NFSGatewayStat) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCreate

`func (o *NFSGatewayStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *NFSGatewayStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *NFSGatewayStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *NFSGatewayStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDeleteLatencyUs

`func (o *NFSGatewayStat) GetDeleteLatencyUs() float64`

GetDeleteLatencyUs returns the DeleteLatencyUs field if non-nil, zero value otherwise.

### GetDeleteLatencyUsOk

`func (o *NFSGatewayStat) GetDeleteLatencyUsOk() (*float64, bool)`

GetDeleteLatencyUsOk returns a tuple with the DeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteLatencyUs

`func (o *NFSGatewayStat) SetDeleteLatencyUs(v float64)`

SetDeleteLatencyUs sets DeleteLatencyUs field to given value.

### HasDeleteLatencyUs

`func (o *NFSGatewayStat) HasDeleteLatencyUs() bool`

HasDeleteLatencyUs returns a boolean if a field has been set.

### GetDeleteOps

`func (o *NFSGatewayStat) GetDeleteOps() float64`

GetDeleteOps returns the DeleteOps field if non-nil, zero value otherwise.

### GetDeleteOpsOk

`func (o *NFSGatewayStat) GetDeleteOpsOk() (*float64, bool)`

GetDeleteOpsOk returns a tuple with the DeleteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOps

`func (o *NFSGatewayStat) SetDeleteOps(v float64)`

SetDeleteOps sets DeleteOps field to given value.

### HasDeleteOps

`func (o *NFSGatewayStat) HasDeleteOps() bool`

HasDeleteOps returns a boolean if a field has been set.

### GetDownBandwidthKbyte

`func (o *NFSGatewayStat) GetDownBandwidthKbyte() float64`

GetDownBandwidthKbyte returns the DownBandwidthKbyte field if non-nil, zero value otherwise.

### GetDownBandwidthKbyteOk

`func (o *NFSGatewayStat) GetDownBandwidthKbyteOk() (*float64, bool)`

GetDownBandwidthKbyteOk returns a tuple with the DownBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownBandwidthKbyte

`func (o *NFSGatewayStat) SetDownBandwidthKbyte(v float64)`

SetDownBandwidthKbyte sets DownBandwidthKbyte field to given value.

### HasDownBandwidthKbyte

`func (o *NFSGatewayStat) HasDownBandwidthKbyte() bool`

HasDownBandwidthKbyte returns a boolean if a field has been set.

### GetDownFirstByteLatencyUs

`func (o *NFSGatewayStat) GetDownFirstByteLatencyUs() float64`

GetDownFirstByteLatencyUs returns the DownFirstByteLatencyUs field if non-nil, zero value otherwise.

### GetDownFirstByteLatencyUsOk

`func (o *NFSGatewayStat) GetDownFirstByteLatencyUsOk() (*float64, bool)`

GetDownFirstByteLatencyUsOk returns a tuple with the DownFirstByteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownFirstByteLatencyUs

`func (o *NFSGatewayStat) SetDownFirstByteLatencyUs(v float64)`

SetDownFirstByteLatencyUs sets DownFirstByteLatencyUs field to given value.

### HasDownFirstByteLatencyUs

`func (o *NFSGatewayStat) HasDownFirstByteLatencyUs() bool`

HasDownFirstByteLatencyUs returns a boolean if a field has been set.

### GetDownLatencyUs

`func (o *NFSGatewayStat) GetDownLatencyUs() float64`

GetDownLatencyUs returns the DownLatencyUs field if non-nil, zero value otherwise.

### GetDownLatencyUsOk

`func (o *NFSGatewayStat) GetDownLatencyUsOk() (*float64, bool)`

GetDownLatencyUsOk returns a tuple with the DownLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLatencyUs

`func (o *NFSGatewayStat) SetDownLatencyUs(v float64)`

SetDownLatencyUs sets DownLatencyUs field to given value.

### HasDownLatencyUs

`func (o *NFSGatewayStat) HasDownLatencyUs() bool`

HasDownLatencyUs returns a boolean if a field has been set.

### GetDownOps

`func (o *NFSGatewayStat) GetDownOps() float64`

GetDownOps returns the DownOps field if non-nil, zero value otherwise.

### GetDownOpsOk

`func (o *NFSGatewayStat) GetDownOpsOk() (*float64, bool)`

GetDownOpsOk returns a tuple with the DownOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownOps

`func (o *NFSGatewayStat) SetDownOps(v float64)`

SetDownOps sets DownOps field to given value.

### HasDownOps

`func (o *NFSGatewayStat) HasDownOps() bool`

HasDownOps returns a boolean if a field has been set.

### GetListLatencyUs

`func (o *NFSGatewayStat) GetListLatencyUs() float64`

GetListLatencyUs returns the ListLatencyUs field if non-nil, zero value otherwise.

### GetListLatencyUsOk

`func (o *NFSGatewayStat) GetListLatencyUsOk() (*float64, bool)`

GetListLatencyUsOk returns a tuple with the ListLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListLatencyUs

`func (o *NFSGatewayStat) SetListLatencyUs(v float64)`

SetListLatencyUs sets ListLatencyUs field to given value.

### HasListLatencyUs

`func (o *NFSGatewayStat) HasListLatencyUs() bool`

HasListLatencyUs returns a boolean if a field has been set.

### GetListOps

`func (o *NFSGatewayStat) GetListOps() float64`

GetListOps returns the ListOps field if non-nil, zero value otherwise.

### GetListOpsOk

`func (o *NFSGatewayStat) GetListOpsOk() (*float64, bool)`

GetListOpsOk returns a tuple with the ListOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOps

`func (o *NFSGatewayStat) SetListOps(v float64)`

SetListOps sets ListOps field to given value.

### HasListOps

`func (o *NFSGatewayStat) HasListOps() bool`

HasListOps returns a boolean if a field has been set.

### GetMemUsagePercent

`func (o *NFSGatewayStat) GetMemUsagePercent() float64`

GetMemUsagePercent returns the MemUsagePercent field if non-nil, zero value otherwise.

### GetMemUsagePercentOk

`func (o *NFSGatewayStat) GetMemUsagePercentOk() (*float64, bool)`

GetMemUsagePercentOk returns a tuple with the MemUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsagePercent

`func (o *NFSGatewayStat) SetMemUsagePercent(v float64)`

SetMemUsagePercent sets MemUsagePercent field to given value.

### HasMemUsagePercent

`func (o *NFSGatewayStat) HasMemUsagePercent() bool`

HasMemUsagePercent returns a boolean if a field has been set.

### GetMergeSpeed

`func (o *NFSGatewayStat) GetMergeSpeed() int64`

GetMergeSpeed returns the MergeSpeed field if non-nil, zero value otherwise.

### GetMergeSpeedOk

`func (o *NFSGatewayStat) GetMergeSpeedOk() (*int64, bool)`

GetMergeSpeedOk returns a tuple with the MergeSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeSpeed

`func (o *NFSGatewayStat) SetMergeSpeed(v int64)`

SetMergeSpeed sets MergeSpeed field to given value.

### HasMergeSpeed

`func (o *NFSGatewayStat) HasMergeSpeed() bool`

HasMergeSpeed returns a boolean if a field has been set.

### GetReplicationBandwidthKbyte

`func (o *NFSGatewayStat) GetReplicationBandwidthKbyte() float64`

GetReplicationBandwidthKbyte returns the ReplicationBandwidthKbyte field if non-nil, zero value otherwise.

### GetReplicationBandwidthKbyteOk

`func (o *NFSGatewayStat) GetReplicationBandwidthKbyteOk() (*float64, bool)`

GetReplicationBandwidthKbyteOk returns a tuple with the ReplicationBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBandwidthKbyte

`func (o *NFSGatewayStat) SetReplicationBandwidthKbyte(v float64)`

SetReplicationBandwidthKbyte sets ReplicationBandwidthKbyte field to given value.

### HasReplicationBandwidthKbyte

`func (o *NFSGatewayStat) HasReplicationBandwidthKbyte() bool`

HasReplicationBandwidthKbyte returns a boolean if a field has been set.

### GetReplicationOps

`func (o *NFSGatewayStat) GetReplicationOps() float64`

GetReplicationOps returns the ReplicationOps field if non-nil, zero value otherwise.

### GetReplicationOpsOk

`func (o *NFSGatewayStat) GetReplicationOpsOk() (*float64, bool)`

GetReplicationOpsOk returns a tuple with the ReplicationOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationOps

`func (o *NFSGatewayStat) SetReplicationOps(v float64)`

SetReplicationOps sets ReplicationOps field to given value.

### HasReplicationOps

`func (o *NFSGatewayStat) HasReplicationOps() bool`

HasReplicationOps returns a boolean if a field has been set.

### GetRequests

`func (o *NFSGatewayStat) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *NFSGatewayStat) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *NFSGatewayStat) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *NFSGatewayStat) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSyncDownBandwidthKbyte

`func (o *NFSGatewayStat) GetSyncDownBandwidthKbyte() float64`

GetSyncDownBandwidthKbyte returns the SyncDownBandwidthKbyte field if non-nil, zero value otherwise.

### GetSyncDownBandwidthKbyteOk

`func (o *NFSGatewayStat) GetSyncDownBandwidthKbyteOk() (*float64, bool)`

GetSyncDownBandwidthKbyteOk returns a tuple with the SyncDownBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDownBandwidthKbyte

`func (o *NFSGatewayStat) SetSyncDownBandwidthKbyte(v float64)`

SetSyncDownBandwidthKbyte sets SyncDownBandwidthKbyte field to given value.

### HasSyncDownBandwidthKbyte

`func (o *NFSGatewayStat) HasSyncDownBandwidthKbyte() bool`

HasSyncDownBandwidthKbyte returns a boolean if a field has been set.

### GetSyncDownLastByteLatencyUs

`func (o *NFSGatewayStat) GetSyncDownLastByteLatencyUs() float64`

GetSyncDownLastByteLatencyUs returns the SyncDownLastByteLatencyUs field if non-nil, zero value otherwise.

### GetSyncDownLastByteLatencyUsOk

`func (o *NFSGatewayStat) GetSyncDownLastByteLatencyUsOk() (*float64, bool)`

GetSyncDownLastByteLatencyUsOk returns a tuple with the SyncDownLastByteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDownLastByteLatencyUs

`func (o *NFSGatewayStat) SetSyncDownLastByteLatencyUs(v float64)`

SetSyncDownLastByteLatencyUs sets SyncDownLastByteLatencyUs field to given value.

### HasSyncDownLastByteLatencyUs

`func (o *NFSGatewayStat) HasSyncDownLastByteLatencyUs() bool`

HasSyncDownLastByteLatencyUs returns a boolean if a field has been set.

### GetSyncDownLatencyUs

`func (o *NFSGatewayStat) GetSyncDownLatencyUs() float64`

GetSyncDownLatencyUs returns the SyncDownLatencyUs field if non-nil, zero value otherwise.

### GetSyncDownLatencyUsOk

`func (o *NFSGatewayStat) GetSyncDownLatencyUsOk() (*float64, bool)`

GetSyncDownLatencyUsOk returns a tuple with the SyncDownLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDownLatencyUs

`func (o *NFSGatewayStat) SetSyncDownLatencyUs(v float64)`

SetSyncDownLatencyUs sets SyncDownLatencyUs field to given value.

### HasSyncDownLatencyUs

`func (o *NFSGatewayStat) HasSyncDownLatencyUs() bool`

HasSyncDownLatencyUs returns a boolean if a field has been set.

### GetSyncDownOps

`func (o *NFSGatewayStat) GetSyncDownOps() float64`

GetSyncDownOps returns the SyncDownOps field if non-nil, zero value otherwise.

### GetSyncDownOpsOk

`func (o *NFSGatewayStat) GetSyncDownOpsOk() (*float64, bool)`

GetSyncDownOpsOk returns a tuple with the SyncDownOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDownOps

`func (o *NFSGatewayStat) SetSyncDownOps(v float64)`

SetSyncDownOps sets SyncDownOps field to given value.

### HasSyncDownOps

`func (o *NFSGatewayStat) HasSyncDownOps() bool`

HasSyncDownOps returns a boolean if a field has been set.

### GetSyncUpBandwidthKbyte

`func (o *NFSGatewayStat) GetSyncUpBandwidthKbyte() float64`

GetSyncUpBandwidthKbyte returns the SyncUpBandwidthKbyte field if non-nil, zero value otherwise.

### GetSyncUpBandwidthKbyteOk

`func (o *NFSGatewayStat) GetSyncUpBandwidthKbyteOk() (*float64, bool)`

GetSyncUpBandwidthKbyteOk returns a tuple with the SyncUpBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUpBandwidthKbyte

`func (o *NFSGatewayStat) SetSyncUpBandwidthKbyte(v float64)`

SetSyncUpBandwidthKbyte sets SyncUpBandwidthKbyte field to given value.

### HasSyncUpBandwidthKbyte

`func (o *NFSGatewayStat) HasSyncUpBandwidthKbyte() bool`

HasSyncUpBandwidthKbyte returns a boolean if a field has been set.

### GetSyncUpLatencyUs

`func (o *NFSGatewayStat) GetSyncUpLatencyUs() float64`

GetSyncUpLatencyUs returns the SyncUpLatencyUs field if non-nil, zero value otherwise.

### GetSyncUpLatencyUsOk

`func (o *NFSGatewayStat) GetSyncUpLatencyUsOk() (*float64, bool)`

GetSyncUpLatencyUsOk returns a tuple with the SyncUpLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUpLatencyUs

`func (o *NFSGatewayStat) SetSyncUpLatencyUs(v float64)`

SetSyncUpLatencyUs sets SyncUpLatencyUs field to given value.

### HasSyncUpLatencyUs

`func (o *NFSGatewayStat) HasSyncUpLatencyUs() bool`

HasSyncUpLatencyUs returns a boolean if a field has been set.

### GetSyncUpOps

`func (o *NFSGatewayStat) GetSyncUpOps() float64`

GetSyncUpOps returns the SyncUpOps field if non-nil, zero value otherwise.

### GetSyncUpOpsOk

`func (o *NFSGatewayStat) GetSyncUpOpsOk() (*float64, bool)`

GetSyncUpOpsOk returns a tuple with the SyncUpOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUpOps

`func (o *NFSGatewayStat) SetSyncUpOps(v float64)`

SetSyncUpOps sets SyncUpOps field to given value.

### HasSyncUpOps

`func (o *NFSGatewayStat) HasSyncUpOps() bool`

HasSyncUpOps returns a boolean if a field has been set.

### GetUpBandwidthKbyte

`func (o *NFSGatewayStat) GetUpBandwidthKbyte() float64`

GetUpBandwidthKbyte returns the UpBandwidthKbyte field if non-nil, zero value otherwise.

### GetUpBandwidthKbyteOk

`func (o *NFSGatewayStat) GetUpBandwidthKbyteOk() (*float64, bool)`

GetUpBandwidthKbyteOk returns a tuple with the UpBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpBandwidthKbyte

`func (o *NFSGatewayStat) SetUpBandwidthKbyte(v float64)`

SetUpBandwidthKbyte sets UpBandwidthKbyte field to given value.

### HasUpBandwidthKbyte

`func (o *NFSGatewayStat) HasUpBandwidthKbyte() bool`

HasUpBandwidthKbyte returns a boolean if a field has been set.

### GetUpLastByteLatencyUs

`func (o *NFSGatewayStat) GetUpLastByteLatencyUs() float64`

GetUpLastByteLatencyUs returns the UpLastByteLatencyUs field if non-nil, zero value otherwise.

### GetUpLastByteLatencyUsOk

`func (o *NFSGatewayStat) GetUpLastByteLatencyUsOk() (*float64, bool)`

GetUpLastByteLatencyUsOk returns a tuple with the UpLastByteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLastByteLatencyUs

`func (o *NFSGatewayStat) SetUpLastByteLatencyUs(v float64)`

SetUpLastByteLatencyUs sets UpLastByteLatencyUs field to given value.

### HasUpLastByteLatencyUs

`func (o *NFSGatewayStat) HasUpLastByteLatencyUs() bool`

HasUpLastByteLatencyUs returns a boolean if a field has been set.

### GetUpLatencyUs

`func (o *NFSGatewayStat) GetUpLatencyUs() float64`

GetUpLatencyUs returns the UpLatencyUs field if non-nil, zero value otherwise.

### GetUpLatencyUsOk

`func (o *NFSGatewayStat) GetUpLatencyUsOk() (*float64, bool)`

GetUpLatencyUsOk returns a tuple with the UpLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLatencyUs

`func (o *NFSGatewayStat) SetUpLatencyUs(v float64)`

SetUpLatencyUs sets UpLatencyUs field to given value.

### HasUpLatencyUs

`func (o *NFSGatewayStat) HasUpLatencyUs() bool`

HasUpLatencyUs returns a boolean if a field has been set.

### GetUpOps

`func (o *NFSGatewayStat) GetUpOps() float64`

GetUpOps returns the UpOps field if non-nil, zero value otherwise.

### GetUpOpsOk

`func (o *NFSGatewayStat) GetUpOpsOk() (*float64, bool)`

GetUpOpsOk returns a tuple with the UpOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpOps

`func (o *NFSGatewayStat) SetUpOps(v float64)`

SetUpOps sets UpOps field to given value.

### HasUpOps

`func (o *NFSGatewayStat) HasUpOps() bool`

HasUpOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


