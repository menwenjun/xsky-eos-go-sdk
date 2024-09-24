# MetadataServiceStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DeleteLatencyUs** | Pointer to **float64** |  | [optional] 
**DeleteOps** | Pointer to **float64** |  | [optional] 
**Disk** | Pointer to [**DiskStat**](DiskStat.md) |  | [optional] 
**GetattrsLatencyUs** | Pointer to **float64** |  | [optional] 
**ListLatencyUs** | Pointer to **float64** |  | [optional] 
**ListOps** | Pointer to **float64** |  | [optional] 
**MemUsagePercent** | Pointer to **float64** |  | [optional] 
**OpLatencyUs** | Pointer to **float64** |  | [optional] 
**OpenLatencyUs** | Pointer to **float64** |  | [optional] 
**ReadLatencyUs** | Pointer to **float64** |  | [optional] 
**ReadOps** | Pointer to **float64** |  | [optional] 
**RecoveryDone** | Pointer to **int64** |  | [optional] 
**RecoveryLeftSecond** | Pointer to **float64** |  | [optional] 
**RecoveryOps** | Pointer to **float64** |  | [optional] 
**RecoveryTotal** | Pointer to **int64** | those fields are only for primary xmds | [optional] 
**WriteLatencyUs** | Pointer to **float64** |  | [optional] 
**WriteOps** | Pointer to **float64** |  | [optional] 

## Methods

### NewMetadataServiceStat

`func NewMetadataServiceStat() *MetadataServiceStat`

NewMetadataServiceStat instantiates a new MetadataServiceStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataServiceStatWithDefaults

`func NewMetadataServiceStatWithDefaults() *MetadataServiceStat`

NewMetadataServiceStatWithDefaults instantiates a new MetadataServiceStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *MetadataServiceStat) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *MetadataServiceStat) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *MetadataServiceStat) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *MetadataServiceStat) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCreate

`func (o *MetadataServiceStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *MetadataServiceStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *MetadataServiceStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *MetadataServiceStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDeleteLatencyUs

`func (o *MetadataServiceStat) GetDeleteLatencyUs() float64`

GetDeleteLatencyUs returns the DeleteLatencyUs field if non-nil, zero value otherwise.

### GetDeleteLatencyUsOk

`func (o *MetadataServiceStat) GetDeleteLatencyUsOk() (*float64, bool)`

GetDeleteLatencyUsOk returns a tuple with the DeleteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteLatencyUs

`func (o *MetadataServiceStat) SetDeleteLatencyUs(v float64)`

SetDeleteLatencyUs sets DeleteLatencyUs field to given value.

### HasDeleteLatencyUs

`func (o *MetadataServiceStat) HasDeleteLatencyUs() bool`

HasDeleteLatencyUs returns a boolean if a field has been set.

### GetDeleteOps

`func (o *MetadataServiceStat) GetDeleteOps() float64`

GetDeleteOps returns the DeleteOps field if non-nil, zero value otherwise.

### GetDeleteOpsOk

`func (o *MetadataServiceStat) GetDeleteOpsOk() (*float64, bool)`

GetDeleteOpsOk returns a tuple with the DeleteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOps

`func (o *MetadataServiceStat) SetDeleteOps(v float64)`

SetDeleteOps sets DeleteOps field to given value.

### HasDeleteOps

`func (o *MetadataServiceStat) HasDeleteOps() bool`

HasDeleteOps returns a boolean if a field has been set.

### GetDisk

`func (o *MetadataServiceStat) GetDisk() DiskStat`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *MetadataServiceStat) GetDiskOk() (*DiskStat, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *MetadataServiceStat) SetDisk(v DiskStat)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *MetadataServiceStat) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetGetattrsLatencyUs

`func (o *MetadataServiceStat) GetGetattrsLatencyUs() float64`

GetGetattrsLatencyUs returns the GetattrsLatencyUs field if non-nil, zero value otherwise.

### GetGetattrsLatencyUsOk

`func (o *MetadataServiceStat) GetGetattrsLatencyUsOk() (*float64, bool)`

GetGetattrsLatencyUsOk returns a tuple with the GetattrsLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetattrsLatencyUs

`func (o *MetadataServiceStat) SetGetattrsLatencyUs(v float64)`

SetGetattrsLatencyUs sets GetattrsLatencyUs field to given value.

### HasGetattrsLatencyUs

`func (o *MetadataServiceStat) HasGetattrsLatencyUs() bool`

HasGetattrsLatencyUs returns a boolean if a field has been set.

### GetListLatencyUs

`func (o *MetadataServiceStat) GetListLatencyUs() float64`

GetListLatencyUs returns the ListLatencyUs field if non-nil, zero value otherwise.

### GetListLatencyUsOk

`func (o *MetadataServiceStat) GetListLatencyUsOk() (*float64, bool)`

GetListLatencyUsOk returns a tuple with the ListLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListLatencyUs

`func (o *MetadataServiceStat) SetListLatencyUs(v float64)`

SetListLatencyUs sets ListLatencyUs field to given value.

### HasListLatencyUs

`func (o *MetadataServiceStat) HasListLatencyUs() bool`

HasListLatencyUs returns a boolean if a field has been set.

### GetListOps

`func (o *MetadataServiceStat) GetListOps() float64`

GetListOps returns the ListOps field if non-nil, zero value otherwise.

### GetListOpsOk

`func (o *MetadataServiceStat) GetListOpsOk() (*float64, bool)`

GetListOpsOk returns a tuple with the ListOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOps

`func (o *MetadataServiceStat) SetListOps(v float64)`

SetListOps sets ListOps field to given value.

### HasListOps

`func (o *MetadataServiceStat) HasListOps() bool`

HasListOps returns a boolean if a field has been set.

### GetMemUsagePercent

`func (o *MetadataServiceStat) GetMemUsagePercent() float64`

GetMemUsagePercent returns the MemUsagePercent field if non-nil, zero value otherwise.

### GetMemUsagePercentOk

`func (o *MetadataServiceStat) GetMemUsagePercentOk() (*float64, bool)`

GetMemUsagePercentOk returns a tuple with the MemUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsagePercent

`func (o *MetadataServiceStat) SetMemUsagePercent(v float64)`

SetMemUsagePercent sets MemUsagePercent field to given value.

### HasMemUsagePercent

`func (o *MetadataServiceStat) HasMemUsagePercent() bool`

HasMemUsagePercent returns a boolean if a field has been set.

### GetOpLatencyUs

`func (o *MetadataServiceStat) GetOpLatencyUs() float64`

GetOpLatencyUs returns the OpLatencyUs field if non-nil, zero value otherwise.

### GetOpLatencyUsOk

`func (o *MetadataServiceStat) GetOpLatencyUsOk() (*float64, bool)`

GetOpLatencyUsOk returns a tuple with the OpLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpLatencyUs

`func (o *MetadataServiceStat) SetOpLatencyUs(v float64)`

SetOpLatencyUs sets OpLatencyUs field to given value.

### HasOpLatencyUs

`func (o *MetadataServiceStat) HasOpLatencyUs() bool`

HasOpLatencyUs returns a boolean if a field has been set.

### GetOpenLatencyUs

`func (o *MetadataServiceStat) GetOpenLatencyUs() float64`

GetOpenLatencyUs returns the OpenLatencyUs field if non-nil, zero value otherwise.

### GetOpenLatencyUsOk

`func (o *MetadataServiceStat) GetOpenLatencyUsOk() (*float64, bool)`

GetOpenLatencyUsOk returns a tuple with the OpenLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenLatencyUs

`func (o *MetadataServiceStat) SetOpenLatencyUs(v float64)`

SetOpenLatencyUs sets OpenLatencyUs field to given value.

### HasOpenLatencyUs

`func (o *MetadataServiceStat) HasOpenLatencyUs() bool`

HasOpenLatencyUs returns a boolean if a field has been set.

### GetReadLatencyUs

`func (o *MetadataServiceStat) GetReadLatencyUs() float64`

GetReadLatencyUs returns the ReadLatencyUs field if non-nil, zero value otherwise.

### GetReadLatencyUsOk

`func (o *MetadataServiceStat) GetReadLatencyUsOk() (*float64, bool)`

GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLatencyUs

`func (o *MetadataServiceStat) SetReadLatencyUs(v float64)`

SetReadLatencyUs sets ReadLatencyUs field to given value.

### HasReadLatencyUs

`func (o *MetadataServiceStat) HasReadLatencyUs() bool`

HasReadLatencyUs returns a boolean if a field has been set.

### GetReadOps

`func (o *MetadataServiceStat) GetReadOps() float64`

GetReadOps returns the ReadOps field if non-nil, zero value otherwise.

### GetReadOpsOk

`func (o *MetadataServiceStat) GetReadOpsOk() (*float64, bool)`

GetReadOpsOk returns a tuple with the ReadOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOps

`func (o *MetadataServiceStat) SetReadOps(v float64)`

SetReadOps sets ReadOps field to given value.

### HasReadOps

`func (o *MetadataServiceStat) HasReadOps() bool`

HasReadOps returns a boolean if a field has been set.

### GetRecoveryDone

`func (o *MetadataServiceStat) GetRecoveryDone() int64`

GetRecoveryDone returns the RecoveryDone field if non-nil, zero value otherwise.

### GetRecoveryDoneOk

`func (o *MetadataServiceStat) GetRecoveryDoneOk() (*int64, bool)`

GetRecoveryDoneOk returns a tuple with the RecoveryDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryDone

`func (o *MetadataServiceStat) SetRecoveryDone(v int64)`

SetRecoveryDone sets RecoveryDone field to given value.

### HasRecoveryDone

`func (o *MetadataServiceStat) HasRecoveryDone() bool`

HasRecoveryDone returns a boolean if a field has been set.

### GetRecoveryLeftSecond

`func (o *MetadataServiceStat) GetRecoveryLeftSecond() float64`

GetRecoveryLeftSecond returns the RecoveryLeftSecond field if non-nil, zero value otherwise.

### GetRecoveryLeftSecondOk

`func (o *MetadataServiceStat) GetRecoveryLeftSecondOk() (*float64, bool)`

GetRecoveryLeftSecondOk returns a tuple with the RecoveryLeftSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryLeftSecond

`func (o *MetadataServiceStat) SetRecoveryLeftSecond(v float64)`

SetRecoveryLeftSecond sets RecoveryLeftSecond field to given value.

### HasRecoveryLeftSecond

`func (o *MetadataServiceStat) HasRecoveryLeftSecond() bool`

HasRecoveryLeftSecond returns a boolean if a field has been set.

### GetRecoveryOps

`func (o *MetadataServiceStat) GetRecoveryOps() float64`

GetRecoveryOps returns the RecoveryOps field if non-nil, zero value otherwise.

### GetRecoveryOpsOk

`func (o *MetadataServiceStat) GetRecoveryOpsOk() (*float64, bool)`

GetRecoveryOpsOk returns a tuple with the RecoveryOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryOps

`func (o *MetadataServiceStat) SetRecoveryOps(v float64)`

SetRecoveryOps sets RecoveryOps field to given value.

### HasRecoveryOps

`func (o *MetadataServiceStat) HasRecoveryOps() bool`

HasRecoveryOps returns a boolean if a field has been set.

### GetRecoveryTotal

`func (o *MetadataServiceStat) GetRecoveryTotal() int64`

GetRecoveryTotal returns the RecoveryTotal field if non-nil, zero value otherwise.

### GetRecoveryTotalOk

`func (o *MetadataServiceStat) GetRecoveryTotalOk() (*int64, bool)`

GetRecoveryTotalOk returns a tuple with the RecoveryTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTotal

`func (o *MetadataServiceStat) SetRecoveryTotal(v int64)`

SetRecoveryTotal sets RecoveryTotal field to given value.

### HasRecoveryTotal

`func (o *MetadataServiceStat) HasRecoveryTotal() bool`

HasRecoveryTotal returns a boolean if a field has been set.

### GetWriteLatencyUs

`func (o *MetadataServiceStat) GetWriteLatencyUs() float64`

GetWriteLatencyUs returns the WriteLatencyUs field if non-nil, zero value otherwise.

### GetWriteLatencyUsOk

`func (o *MetadataServiceStat) GetWriteLatencyUsOk() (*float64, bool)`

GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLatencyUs

`func (o *MetadataServiceStat) SetWriteLatencyUs(v float64)`

SetWriteLatencyUs sets WriteLatencyUs field to given value.

### HasWriteLatencyUs

`func (o *MetadataServiceStat) HasWriteLatencyUs() bool`

HasWriteLatencyUs returns a boolean if a field has been set.

### GetWriteOps

`func (o *MetadataServiceStat) GetWriteOps() float64`

GetWriteOps returns the WriteOps field if non-nil, zero value otherwise.

### GetWriteOpsOk

`func (o *MetadataServiceStat) GetWriteOpsOk() (*float64, bool)`

GetWriteOpsOk returns a tuple with the WriteOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteOps

`func (o *MetadataServiceStat) SetWriteOps(v float64)`

SetWriteOps sets WriteOps field to given value.

### HasWriteOps

`func (o *MetadataServiceStat) HasWriteOps() bool`

HasWriteOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


