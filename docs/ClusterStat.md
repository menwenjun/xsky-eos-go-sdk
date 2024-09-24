# ClusterStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualKbyte** | Pointer to **int64** |  | [optional] 
**AllocatedSize** | Pointer to **int64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DataKbyte** | Pointer to **int64** |  | [optional] 
**DegradedPercent** | Pointer to **float64** |  | [optional] 
**ErrorKbyte** | Pointer to **int64** |  | [optional] 
**HealthyPercent** | Pointer to **float64** |  | [optional] 
**MaxAvailableKbyte** | Pointer to **int64** |  | [optional] 
**MinAvailableKbyte** | Pointer to **int64** |  | [optional] 
**OsDownBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**OsDownIops** | Pointer to **float64** |  | [optional] 
**OsMergeSpeed** | Pointer to **int64** |  | [optional] 
**OsUpBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**OsUpIops** | Pointer to **float64** |  | [optional] 
**ReadBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**ReadIops** | Pointer to **float64** |  | [optional] 
**ReadLatencyUs** | Pointer to **float64** |  | [optional] 
**RecoveryBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**RecoveryIops** | Pointer to **float64** |  | [optional] 
**RecoveryPercent** | Pointer to **float64** |  | [optional] 
**TotalKbyte** | Pointer to **int64** |  | [optional] 
**UnavailablePercent** | Pointer to **float64** |  | [optional] 
**UsedKbyte** | Pointer to **int64** |  | [optional] 
**WriteBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**WriteIops** | Pointer to **float64** |  | [optional] 
**WriteLatencyUs** | Pointer to **float64** |  | [optional] 

## Methods

### NewClusterStat

`func NewClusterStat() *ClusterStat`

NewClusterStat instantiates a new ClusterStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatWithDefaults

`func NewClusterStatWithDefaults() *ClusterStat`

NewClusterStatWithDefaults instantiates a new ClusterStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualKbyte

`func (o *ClusterStat) GetActualKbyte() int64`

GetActualKbyte returns the ActualKbyte field if non-nil, zero value otherwise.

### GetActualKbyteOk

`func (o *ClusterStat) GetActualKbyteOk() (*int64, bool)`

GetActualKbyteOk returns a tuple with the ActualKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualKbyte

`func (o *ClusterStat) SetActualKbyte(v int64)`

SetActualKbyte sets ActualKbyte field to given value.

### HasActualKbyte

`func (o *ClusterStat) HasActualKbyte() bool`

HasActualKbyte returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *ClusterStat) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *ClusterStat) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *ClusterStat) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *ClusterStat) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetCreate

`func (o *ClusterStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ClusterStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ClusterStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ClusterStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataKbyte

`func (o *ClusterStat) GetDataKbyte() int64`

GetDataKbyte returns the DataKbyte field if non-nil, zero value otherwise.

### GetDataKbyteOk

`func (o *ClusterStat) GetDataKbyteOk() (*int64, bool)`

GetDataKbyteOk returns a tuple with the DataKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKbyte

`func (o *ClusterStat) SetDataKbyte(v int64)`

SetDataKbyte sets DataKbyte field to given value.

### HasDataKbyte

`func (o *ClusterStat) HasDataKbyte() bool`

HasDataKbyte returns a boolean if a field has been set.

### GetDegradedPercent

`func (o *ClusterStat) GetDegradedPercent() float64`

GetDegradedPercent returns the DegradedPercent field if non-nil, zero value otherwise.

### GetDegradedPercentOk

`func (o *ClusterStat) GetDegradedPercentOk() (*float64, bool)`

GetDegradedPercentOk returns a tuple with the DegradedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedPercent

`func (o *ClusterStat) SetDegradedPercent(v float64)`

SetDegradedPercent sets DegradedPercent field to given value.

### HasDegradedPercent

`func (o *ClusterStat) HasDegradedPercent() bool`

HasDegradedPercent returns a boolean if a field has been set.

### GetErrorKbyte

`func (o *ClusterStat) GetErrorKbyte() int64`

GetErrorKbyte returns the ErrorKbyte field if non-nil, zero value otherwise.

### GetErrorKbyteOk

`func (o *ClusterStat) GetErrorKbyteOk() (*int64, bool)`

GetErrorKbyteOk returns a tuple with the ErrorKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorKbyte

`func (o *ClusterStat) SetErrorKbyte(v int64)`

SetErrorKbyte sets ErrorKbyte field to given value.

### HasErrorKbyte

`func (o *ClusterStat) HasErrorKbyte() bool`

HasErrorKbyte returns a boolean if a field has been set.

### GetHealthyPercent

`func (o *ClusterStat) GetHealthyPercent() float64`

GetHealthyPercent returns the HealthyPercent field if non-nil, zero value otherwise.

### GetHealthyPercentOk

`func (o *ClusterStat) GetHealthyPercentOk() (*float64, bool)`

GetHealthyPercentOk returns a tuple with the HealthyPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyPercent

`func (o *ClusterStat) SetHealthyPercent(v float64)`

SetHealthyPercent sets HealthyPercent field to given value.

### HasHealthyPercent

`func (o *ClusterStat) HasHealthyPercent() bool`

HasHealthyPercent returns a boolean if a field has been set.

### GetMaxAvailableKbyte

`func (o *ClusterStat) GetMaxAvailableKbyte() int64`

GetMaxAvailableKbyte returns the MaxAvailableKbyte field if non-nil, zero value otherwise.

### GetMaxAvailableKbyteOk

`func (o *ClusterStat) GetMaxAvailableKbyteOk() (*int64, bool)`

GetMaxAvailableKbyteOk returns a tuple with the MaxAvailableKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvailableKbyte

`func (o *ClusterStat) SetMaxAvailableKbyte(v int64)`

SetMaxAvailableKbyte sets MaxAvailableKbyte field to given value.

### HasMaxAvailableKbyte

`func (o *ClusterStat) HasMaxAvailableKbyte() bool`

HasMaxAvailableKbyte returns a boolean if a field has been set.

### GetMinAvailableKbyte

`func (o *ClusterStat) GetMinAvailableKbyte() int64`

GetMinAvailableKbyte returns the MinAvailableKbyte field if non-nil, zero value otherwise.

### GetMinAvailableKbyteOk

`func (o *ClusterStat) GetMinAvailableKbyteOk() (*int64, bool)`

GetMinAvailableKbyteOk returns a tuple with the MinAvailableKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAvailableKbyte

`func (o *ClusterStat) SetMinAvailableKbyte(v int64)`

SetMinAvailableKbyte sets MinAvailableKbyte field to given value.

### HasMinAvailableKbyte

`func (o *ClusterStat) HasMinAvailableKbyte() bool`

HasMinAvailableKbyte returns a boolean if a field has been set.

### GetOsDownBandwidthKbyte

`func (o *ClusterStat) GetOsDownBandwidthKbyte() float64`

GetOsDownBandwidthKbyte returns the OsDownBandwidthKbyte field if non-nil, zero value otherwise.

### GetOsDownBandwidthKbyteOk

`func (o *ClusterStat) GetOsDownBandwidthKbyteOk() (*float64, bool)`

GetOsDownBandwidthKbyteOk returns a tuple with the OsDownBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDownBandwidthKbyte

`func (o *ClusterStat) SetOsDownBandwidthKbyte(v float64)`

SetOsDownBandwidthKbyte sets OsDownBandwidthKbyte field to given value.

### HasOsDownBandwidthKbyte

`func (o *ClusterStat) HasOsDownBandwidthKbyte() bool`

HasOsDownBandwidthKbyte returns a boolean if a field has been set.

### GetOsDownIops

`func (o *ClusterStat) GetOsDownIops() float64`

GetOsDownIops returns the OsDownIops field if non-nil, zero value otherwise.

### GetOsDownIopsOk

`func (o *ClusterStat) GetOsDownIopsOk() (*float64, bool)`

GetOsDownIopsOk returns a tuple with the OsDownIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDownIops

`func (o *ClusterStat) SetOsDownIops(v float64)`

SetOsDownIops sets OsDownIops field to given value.

### HasOsDownIops

`func (o *ClusterStat) HasOsDownIops() bool`

HasOsDownIops returns a boolean if a field has been set.

### GetOsMergeSpeed

`func (o *ClusterStat) GetOsMergeSpeed() int64`

GetOsMergeSpeed returns the OsMergeSpeed field if non-nil, zero value otherwise.

### GetOsMergeSpeedOk

`func (o *ClusterStat) GetOsMergeSpeedOk() (*int64, bool)`

GetOsMergeSpeedOk returns a tuple with the OsMergeSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsMergeSpeed

`func (o *ClusterStat) SetOsMergeSpeed(v int64)`

SetOsMergeSpeed sets OsMergeSpeed field to given value.

### HasOsMergeSpeed

`func (o *ClusterStat) HasOsMergeSpeed() bool`

HasOsMergeSpeed returns a boolean if a field has been set.

### GetOsUpBandwidthKbyte

`func (o *ClusterStat) GetOsUpBandwidthKbyte() float64`

GetOsUpBandwidthKbyte returns the OsUpBandwidthKbyte field if non-nil, zero value otherwise.

### GetOsUpBandwidthKbyteOk

`func (o *ClusterStat) GetOsUpBandwidthKbyteOk() (*float64, bool)`

GetOsUpBandwidthKbyteOk returns a tuple with the OsUpBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUpBandwidthKbyte

`func (o *ClusterStat) SetOsUpBandwidthKbyte(v float64)`

SetOsUpBandwidthKbyte sets OsUpBandwidthKbyte field to given value.

### HasOsUpBandwidthKbyte

`func (o *ClusterStat) HasOsUpBandwidthKbyte() bool`

HasOsUpBandwidthKbyte returns a boolean if a field has been set.

### GetOsUpIops

`func (o *ClusterStat) GetOsUpIops() float64`

GetOsUpIops returns the OsUpIops field if non-nil, zero value otherwise.

### GetOsUpIopsOk

`func (o *ClusterStat) GetOsUpIopsOk() (*float64, bool)`

GetOsUpIopsOk returns a tuple with the OsUpIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUpIops

`func (o *ClusterStat) SetOsUpIops(v float64)`

SetOsUpIops sets OsUpIops field to given value.

### HasOsUpIops

`func (o *ClusterStat) HasOsUpIops() bool`

HasOsUpIops returns a boolean if a field has been set.

### GetReadBandwidthKbyte

`func (o *ClusterStat) GetReadBandwidthKbyte() float64`

GetReadBandwidthKbyte returns the ReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadBandwidthKbyteOk

`func (o *ClusterStat) GetReadBandwidthKbyteOk() (*float64, bool)`

GetReadBandwidthKbyteOk returns a tuple with the ReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidthKbyte

`func (o *ClusterStat) SetReadBandwidthKbyte(v float64)`

SetReadBandwidthKbyte sets ReadBandwidthKbyte field to given value.

### HasReadBandwidthKbyte

`func (o *ClusterStat) HasReadBandwidthKbyte() bool`

HasReadBandwidthKbyte returns a boolean if a field has been set.

### GetReadIops

`func (o *ClusterStat) GetReadIops() float64`

GetReadIops returns the ReadIops field if non-nil, zero value otherwise.

### GetReadIopsOk

`func (o *ClusterStat) GetReadIopsOk() (*float64, bool)`

GetReadIopsOk returns a tuple with the ReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIops

`func (o *ClusterStat) SetReadIops(v float64)`

SetReadIops sets ReadIops field to given value.

### HasReadIops

`func (o *ClusterStat) HasReadIops() bool`

HasReadIops returns a boolean if a field has been set.

### GetReadLatencyUs

`func (o *ClusterStat) GetReadLatencyUs() float64`

GetReadLatencyUs returns the ReadLatencyUs field if non-nil, zero value otherwise.

### GetReadLatencyUsOk

`func (o *ClusterStat) GetReadLatencyUsOk() (*float64, bool)`

GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLatencyUs

`func (o *ClusterStat) SetReadLatencyUs(v float64)`

SetReadLatencyUs sets ReadLatencyUs field to given value.

### HasReadLatencyUs

`func (o *ClusterStat) HasReadLatencyUs() bool`

HasReadLatencyUs returns a boolean if a field has been set.

### GetRecoveryBandwidthKbyte

`func (o *ClusterStat) GetRecoveryBandwidthKbyte() float64`

GetRecoveryBandwidthKbyte returns the RecoveryBandwidthKbyte field if non-nil, zero value otherwise.

### GetRecoveryBandwidthKbyteOk

`func (o *ClusterStat) GetRecoveryBandwidthKbyteOk() (*float64, bool)`

GetRecoveryBandwidthKbyteOk returns a tuple with the RecoveryBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryBandwidthKbyte

`func (o *ClusterStat) SetRecoveryBandwidthKbyte(v float64)`

SetRecoveryBandwidthKbyte sets RecoveryBandwidthKbyte field to given value.

### HasRecoveryBandwidthKbyte

`func (o *ClusterStat) HasRecoveryBandwidthKbyte() bool`

HasRecoveryBandwidthKbyte returns a boolean if a field has been set.

### GetRecoveryIops

`func (o *ClusterStat) GetRecoveryIops() float64`

GetRecoveryIops returns the RecoveryIops field if non-nil, zero value otherwise.

### GetRecoveryIopsOk

`func (o *ClusterStat) GetRecoveryIopsOk() (*float64, bool)`

GetRecoveryIopsOk returns a tuple with the RecoveryIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryIops

`func (o *ClusterStat) SetRecoveryIops(v float64)`

SetRecoveryIops sets RecoveryIops field to given value.

### HasRecoveryIops

`func (o *ClusterStat) HasRecoveryIops() bool`

HasRecoveryIops returns a boolean if a field has been set.

### GetRecoveryPercent

`func (o *ClusterStat) GetRecoveryPercent() float64`

GetRecoveryPercent returns the RecoveryPercent field if non-nil, zero value otherwise.

### GetRecoveryPercentOk

`func (o *ClusterStat) GetRecoveryPercentOk() (*float64, bool)`

GetRecoveryPercentOk returns a tuple with the RecoveryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPercent

`func (o *ClusterStat) SetRecoveryPercent(v float64)`

SetRecoveryPercent sets RecoveryPercent field to given value.

### HasRecoveryPercent

`func (o *ClusterStat) HasRecoveryPercent() bool`

HasRecoveryPercent returns a boolean if a field has been set.

### GetTotalKbyte

`func (o *ClusterStat) GetTotalKbyte() int64`

GetTotalKbyte returns the TotalKbyte field if non-nil, zero value otherwise.

### GetTotalKbyteOk

`func (o *ClusterStat) GetTotalKbyteOk() (*int64, bool)`

GetTotalKbyteOk returns a tuple with the TotalKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKbyte

`func (o *ClusterStat) SetTotalKbyte(v int64)`

SetTotalKbyte sets TotalKbyte field to given value.

### HasTotalKbyte

`func (o *ClusterStat) HasTotalKbyte() bool`

HasTotalKbyte returns a boolean if a field has been set.

### GetUnavailablePercent

`func (o *ClusterStat) GetUnavailablePercent() float64`

GetUnavailablePercent returns the UnavailablePercent field if non-nil, zero value otherwise.

### GetUnavailablePercentOk

`func (o *ClusterStat) GetUnavailablePercentOk() (*float64, bool)`

GetUnavailablePercentOk returns a tuple with the UnavailablePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailablePercent

`func (o *ClusterStat) SetUnavailablePercent(v float64)`

SetUnavailablePercent sets UnavailablePercent field to given value.

### HasUnavailablePercent

`func (o *ClusterStat) HasUnavailablePercent() bool`

HasUnavailablePercent returns a boolean if a field has been set.

### GetUsedKbyte

`func (o *ClusterStat) GetUsedKbyte() int64`

GetUsedKbyte returns the UsedKbyte field if non-nil, zero value otherwise.

### GetUsedKbyteOk

`func (o *ClusterStat) GetUsedKbyteOk() (*int64, bool)`

GetUsedKbyteOk returns a tuple with the UsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte

`func (o *ClusterStat) SetUsedKbyte(v int64)`

SetUsedKbyte sets UsedKbyte field to given value.

### HasUsedKbyte

`func (o *ClusterStat) HasUsedKbyte() bool`

HasUsedKbyte returns a boolean if a field has been set.

### GetWriteBandwidthKbyte

`func (o *ClusterStat) GetWriteBandwidthKbyte() float64`

GetWriteBandwidthKbyte returns the WriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetWriteBandwidthKbyteOk

`func (o *ClusterStat) GetWriteBandwidthKbyteOk() (*float64, bool)`

GetWriteBandwidthKbyteOk returns a tuple with the WriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidthKbyte

`func (o *ClusterStat) SetWriteBandwidthKbyte(v float64)`

SetWriteBandwidthKbyte sets WriteBandwidthKbyte field to given value.

### HasWriteBandwidthKbyte

`func (o *ClusterStat) HasWriteBandwidthKbyte() bool`

HasWriteBandwidthKbyte returns a boolean if a field has been set.

### GetWriteIops

`func (o *ClusterStat) GetWriteIops() float64`

GetWriteIops returns the WriteIops field if non-nil, zero value otherwise.

### GetWriteIopsOk

`func (o *ClusterStat) GetWriteIopsOk() (*float64, bool)`

GetWriteIopsOk returns a tuple with the WriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIops

`func (o *ClusterStat) SetWriteIops(v float64)`

SetWriteIops sets WriteIops field to given value.

### HasWriteIops

`func (o *ClusterStat) HasWriteIops() bool`

HasWriteIops returns a boolean if a field has been set.

### GetWriteLatencyUs

`func (o *ClusterStat) GetWriteLatencyUs() float64`

GetWriteLatencyUs returns the WriteLatencyUs field if non-nil, zero value otherwise.

### GetWriteLatencyUsOk

`func (o *ClusterStat) GetWriteLatencyUsOk() (*float64, bool)`

GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLatencyUs

`func (o *ClusterStat) SetWriteLatencyUs(v float64)`

SetWriteLatencyUs sets WriteLatencyUs field to given value.

### HasWriteLatencyUs

`func (o *ClusterStat) HasWriteLatencyUs() bool`

HasWriteLatencyUs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


