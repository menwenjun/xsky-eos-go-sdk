# ObjectStorageGatewayStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DownBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**DownOps** | Pointer to **float64** |  | [optional] 
**MemUsagePercent** | Pointer to **float64** |  | [optional] 
**MergeSpeed** | Pointer to **int64** |  | [optional] 
**Requests** | Pointer to **int64** |  | [optional] 
**SyncDownBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**SyncDownOps** | Pointer to **float64** |  | [optional] 
**UpBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**UpOps** | Pointer to **float64** |  | [optional] 

## Methods

### NewObjectStorageGatewayStat

`func NewObjectStorageGatewayStat() *ObjectStorageGatewayStat`

NewObjectStorageGatewayStat instantiates a new ObjectStorageGatewayStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageGatewayStatWithDefaults

`func NewObjectStorageGatewayStatWithDefaults() *ObjectStorageGatewayStat`

NewObjectStorageGatewayStatWithDefaults instantiates a new ObjectStorageGatewayStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *ObjectStorageGatewayStat) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *ObjectStorageGatewayStat) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *ObjectStorageGatewayStat) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *ObjectStorageGatewayStat) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageGatewayStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageGatewayStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageGatewayStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageGatewayStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDownBandwidthKbyte

`func (o *ObjectStorageGatewayStat) GetDownBandwidthKbyte() float64`

GetDownBandwidthKbyte returns the DownBandwidthKbyte field if non-nil, zero value otherwise.

### GetDownBandwidthKbyteOk

`func (o *ObjectStorageGatewayStat) GetDownBandwidthKbyteOk() (*float64, bool)`

GetDownBandwidthKbyteOk returns a tuple with the DownBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownBandwidthKbyte

`func (o *ObjectStorageGatewayStat) SetDownBandwidthKbyte(v float64)`

SetDownBandwidthKbyte sets DownBandwidthKbyte field to given value.

### HasDownBandwidthKbyte

`func (o *ObjectStorageGatewayStat) HasDownBandwidthKbyte() bool`

HasDownBandwidthKbyte returns a boolean if a field has been set.

### GetDownOps

`func (o *ObjectStorageGatewayStat) GetDownOps() float64`

GetDownOps returns the DownOps field if non-nil, zero value otherwise.

### GetDownOpsOk

`func (o *ObjectStorageGatewayStat) GetDownOpsOk() (*float64, bool)`

GetDownOpsOk returns a tuple with the DownOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownOps

`func (o *ObjectStorageGatewayStat) SetDownOps(v float64)`

SetDownOps sets DownOps field to given value.

### HasDownOps

`func (o *ObjectStorageGatewayStat) HasDownOps() bool`

HasDownOps returns a boolean if a field has been set.

### GetMemUsagePercent

`func (o *ObjectStorageGatewayStat) GetMemUsagePercent() float64`

GetMemUsagePercent returns the MemUsagePercent field if non-nil, zero value otherwise.

### GetMemUsagePercentOk

`func (o *ObjectStorageGatewayStat) GetMemUsagePercentOk() (*float64, bool)`

GetMemUsagePercentOk returns a tuple with the MemUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsagePercent

`func (o *ObjectStorageGatewayStat) SetMemUsagePercent(v float64)`

SetMemUsagePercent sets MemUsagePercent field to given value.

### HasMemUsagePercent

`func (o *ObjectStorageGatewayStat) HasMemUsagePercent() bool`

HasMemUsagePercent returns a boolean if a field has been set.

### GetMergeSpeed

`func (o *ObjectStorageGatewayStat) GetMergeSpeed() int64`

GetMergeSpeed returns the MergeSpeed field if non-nil, zero value otherwise.

### GetMergeSpeedOk

`func (o *ObjectStorageGatewayStat) GetMergeSpeedOk() (*int64, bool)`

GetMergeSpeedOk returns a tuple with the MergeSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeSpeed

`func (o *ObjectStorageGatewayStat) SetMergeSpeed(v int64)`

SetMergeSpeed sets MergeSpeed field to given value.

### HasMergeSpeed

`func (o *ObjectStorageGatewayStat) HasMergeSpeed() bool`

HasMergeSpeed returns a boolean if a field has been set.

### GetRequests

`func (o *ObjectStorageGatewayStat) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ObjectStorageGatewayStat) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ObjectStorageGatewayStat) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ObjectStorageGatewayStat) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSyncDownBandwidthKbyte

`func (o *ObjectStorageGatewayStat) GetSyncDownBandwidthKbyte() float64`

GetSyncDownBandwidthKbyte returns the SyncDownBandwidthKbyte field if non-nil, zero value otherwise.

### GetSyncDownBandwidthKbyteOk

`func (o *ObjectStorageGatewayStat) GetSyncDownBandwidthKbyteOk() (*float64, bool)`

GetSyncDownBandwidthKbyteOk returns a tuple with the SyncDownBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDownBandwidthKbyte

`func (o *ObjectStorageGatewayStat) SetSyncDownBandwidthKbyte(v float64)`

SetSyncDownBandwidthKbyte sets SyncDownBandwidthKbyte field to given value.

### HasSyncDownBandwidthKbyte

`func (o *ObjectStorageGatewayStat) HasSyncDownBandwidthKbyte() bool`

HasSyncDownBandwidthKbyte returns a boolean if a field has been set.

### GetSyncDownOps

`func (o *ObjectStorageGatewayStat) GetSyncDownOps() float64`

GetSyncDownOps returns the SyncDownOps field if non-nil, zero value otherwise.

### GetSyncDownOpsOk

`func (o *ObjectStorageGatewayStat) GetSyncDownOpsOk() (*float64, bool)`

GetSyncDownOpsOk returns a tuple with the SyncDownOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDownOps

`func (o *ObjectStorageGatewayStat) SetSyncDownOps(v float64)`

SetSyncDownOps sets SyncDownOps field to given value.

### HasSyncDownOps

`func (o *ObjectStorageGatewayStat) HasSyncDownOps() bool`

HasSyncDownOps returns a boolean if a field has been set.

### GetUpBandwidthKbyte

`func (o *ObjectStorageGatewayStat) GetUpBandwidthKbyte() float64`

GetUpBandwidthKbyte returns the UpBandwidthKbyte field if non-nil, zero value otherwise.

### GetUpBandwidthKbyteOk

`func (o *ObjectStorageGatewayStat) GetUpBandwidthKbyteOk() (*float64, bool)`

GetUpBandwidthKbyteOk returns a tuple with the UpBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpBandwidthKbyte

`func (o *ObjectStorageGatewayStat) SetUpBandwidthKbyte(v float64)`

SetUpBandwidthKbyte sets UpBandwidthKbyte field to given value.

### HasUpBandwidthKbyte

`func (o *ObjectStorageGatewayStat) HasUpBandwidthKbyte() bool`

HasUpBandwidthKbyte returns a boolean if a field has been set.

### GetUpOps

`func (o *ObjectStorageGatewayStat) GetUpOps() float64`

GetUpOps returns the UpOps field if non-nil, zero value otherwise.

### GetUpOpsOk

`func (o *ObjectStorageGatewayStat) GetUpOpsOk() (*float64, bool)`

GetUpOpsOk returns a tuple with the UpOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpOps

`func (o *ObjectStorageGatewayStat) SetUpOps(v float64)`

SetUpOps sets UpOps field to given value.

### HasUpOps

`func (o *ObjectStorageGatewayStat) HasUpOps() bool`

HasUpOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


