# DfsTierStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualKbyte** | Pointer to **int64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DataKbyte** | Pointer to **int64** |  | [optional] 
**DegradedPercent** | Pointer to **float64** |  | [optional] 
**HealthyPercent** | Pointer to **float64** |  | [optional] 
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

### NewDfsTierStat

`func NewDfsTierStat() *DfsTierStat`

NewDfsTierStat instantiates a new DfsTierStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsTierStatWithDefaults

`func NewDfsTierStatWithDefaults() *DfsTierStat`

NewDfsTierStatWithDefaults instantiates a new DfsTierStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualKbyte

`func (o *DfsTierStat) GetActualKbyte() int64`

GetActualKbyte returns the ActualKbyte field if non-nil, zero value otherwise.

### GetActualKbyteOk

`func (o *DfsTierStat) GetActualKbyteOk() (*int64, bool)`

GetActualKbyteOk returns a tuple with the ActualKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualKbyte

`func (o *DfsTierStat) SetActualKbyte(v int64)`

SetActualKbyte sets ActualKbyte field to given value.

### HasActualKbyte

`func (o *DfsTierStat) HasActualKbyte() bool`

HasActualKbyte returns a boolean if a field has been set.

### GetCreate

`func (o *DfsTierStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsTierStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsTierStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsTierStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataKbyte

`func (o *DfsTierStat) GetDataKbyte() int64`

GetDataKbyte returns the DataKbyte field if non-nil, zero value otherwise.

### GetDataKbyteOk

`func (o *DfsTierStat) GetDataKbyteOk() (*int64, bool)`

GetDataKbyteOk returns a tuple with the DataKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKbyte

`func (o *DfsTierStat) SetDataKbyte(v int64)`

SetDataKbyte sets DataKbyte field to given value.

### HasDataKbyte

`func (o *DfsTierStat) HasDataKbyte() bool`

HasDataKbyte returns a boolean if a field has been set.

### GetDegradedPercent

`func (o *DfsTierStat) GetDegradedPercent() float64`

GetDegradedPercent returns the DegradedPercent field if non-nil, zero value otherwise.

### GetDegradedPercentOk

`func (o *DfsTierStat) GetDegradedPercentOk() (*float64, bool)`

GetDegradedPercentOk returns a tuple with the DegradedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedPercent

`func (o *DfsTierStat) SetDegradedPercent(v float64)`

SetDegradedPercent sets DegradedPercent field to given value.

### HasDegradedPercent

`func (o *DfsTierStat) HasDegradedPercent() bool`

HasDegradedPercent returns a boolean if a field has been set.

### GetHealthyPercent

`func (o *DfsTierStat) GetHealthyPercent() float64`

GetHealthyPercent returns the HealthyPercent field if non-nil, zero value otherwise.

### GetHealthyPercentOk

`func (o *DfsTierStat) GetHealthyPercentOk() (*float64, bool)`

GetHealthyPercentOk returns a tuple with the HealthyPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyPercent

`func (o *DfsTierStat) SetHealthyPercent(v float64)`

SetHealthyPercent sets HealthyPercent field to given value.

### HasHealthyPercent

`func (o *DfsTierStat) HasHealthyPercent() bool`

HasHealthyPercent returns a boolean if a field has been set.

### GetReadBandwidthKbyte

`func (o *DfsTierStat) GetReadBandwidthKbyte() float64`

GetReadBandwidthKbyte returns the ReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadBandwidthKbyteOk

`func (o *DfsTierStat) GetReadBandwidthKbyteOk() (*float64, bool)`

GetReadBandwidthKbyteOk returns a tuple with the ReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidthKbyte

`func (o *DfsTierStat) SetReadBandwidthKbyte(v float64)`

SetReadBandwidthKbyte sets ReadBandwidthKbyte field to given value.

### HasReadBandwidthKbyte

`func (o *DfsTierStat) HasReadBandwidthKbyte() bool`

HasReadBandwidthKbyte returns a boolean if a field has been set.

### GetReadIops

`func (o *DfsTierStat) GetReadIops() float64`

GetReadIops returns the ReadIops field if non-nil, zero value otherwise.

### GetReadIopsOk

`func (o *DfsTierStat) GetReadIopsOk() (*float64, bool)`

GetReadIopsOk returns a tuple with the ReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIops

`func (o *DfsTierStat) SetReadIops(v float64)`

SetReadIops sets ReadIops field to given value.

### HasReadIops

`func (o *DfsTierStat) HasReadIops() bool`

HasReadIops returns a boolean if a field has been set.

### GetReadLatencyUs

`func (o *DfsTierStat) GetReadLatencyUs() float64`

GetReadLatencyUs returns the ReadLatencyUs field if non-nil, zero value otherwise.

### GetReadLatencyUsOk

`func (o *DfsTierStat) GetReadLatencyUsOk() (*float64, bool)`

GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLatencyUs

`func (o *DfsTierStat) SetReadLatencyUs(v float64)`

SetReadLatencyUs sets ReadLatencyUs field to given value.

### HasReadLatencyUs

`func (o *DfsTierStat) HasReadLatencyUs() bool`

HasReadLatencyUs returns a boolean if a field has been set.

### GetRecoveryBandwidthKbyte

`func (o *DfsTierStat) GetRecoveryBandwidthKbyte() float64`

GetRecoveryBandwidthKbyte returns the RecoveryBandwidthKbyte field if non-nil, zero value otherwise.

### GetRecoveryBandwidthKbyteOk

`func (o *DfsTierStat) GetRecoveryBandwidthKbyteOk() (*float64, bool)`

GetRecoveryBandwidthKbyteOk returns a tuple with the RecoveryBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryBandwidthKbyte

`func (o *DfsTierStat) SetRecoveryBandwidthKbyte(v float64)`

SetRecoveryBandwidthKbyte sets RecoveryBandwidthKbyte field to given value.

### HasRecoveryBandwidthKbyte

`func (o *DfsTierStat) HasRecoveryBandwidthKbyte() bool`

HasRecoveryBandwidthKbyte returns a boolean if a field has been set.

### GetRecoveryIops

`func (o *DfsTierStat) GetRecoveryIops() float64`

GetRecoveryIops returns the RecoveryIops field if non-nil, zero value otherwise.

### GetRecoveryIopsOk

`func (o *DfsTierStat) GetRecoveryIopsOk() (*float64, bool)`

GetRecoveryIopsOk returns a tuple with the RecoveryIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryIops

`func (o *DfsTierStat) SetRecoveryIops(v float64)`

SetRecoveryIops sets RecoveryIops field to given value.

### HasRecoveryIops

`func (o *DfsTierStat) HasRecoveryIops() bool`

HasRecoveryIops returns a boolean if a field has been set.

### GetRecoveryPercent

`func (o *DfsTierStat) GetRecoveryPercent() float64`

GetRecoveryPercent returns the RecoveryPercent field if non-nil, zero value otherwise.

### GetRecoveryPercentOk

`func (o *DfsTierStat) GetRecoveryPercentOk() (*float64, bool)`

GetRecoveryPercentOk returns a tuple with the RecoveryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPercent

`func (o *DfsTierStat) SetRecoveryPercent(v float64)`

SetRecoveryPercent sets RecoveryPercent field to given value.

### HasRecoveryPercent

`func (o *DfsTierStat) HasRecoveryPercent() bool`

HasRecoveryPercent returns a boolean if a field has been set.

### GetTotalKbyte

`func (o *DfsTierStat) GetTotalKbyte() int64`

GetTotalKbyte returns the TotalKbyte field if non-nil, zero value otherwise.

### GetTotalKbyteOk

`func (o *DfsTierStat) GetTotalKbyteOk() (*int64, bool)`

GetTotalKbyteOk returns a tuple with the TotalKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKbyte

`func (o *DfsTierStat) SetTotalKbyte(v int64)`

SetTotalKbyte sets TotalKbyte field to given value.

### HasTotalKbyte

`func (o *DfsTierStat) HasTotalKbyte() bool`

HasTotalKbyte returns a boolean if a field has been set.

### GetUnavailablePercent

`func (o *DfsTierStat) GetUnavailablePercent() float64`

GetUnavailablePercent returns the UnavailablePercent field if non-nil, zero value otherwise.

### GetUnavailablePercentOk

`func (o *DfsTierStat) GetUnavailablePercentOk() (*float64, bool)`

GetUnavailablePercentOk returns a tuple with the UnavailablePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailablePercent

`func (o *DfsTierStat) SetUnavailablePercent(v float64)`

SetUnavailablePercent sets UnavailablePercent field to given value.

### HasUnavailablePercent

`func (o *DfsTierStat) HasUnavailablePercent() bool`

HasUnavailablePercent returns a boolean if a field has been set.

### GetUsedKbyte

`func (o *DfsTierStat) GetUsedKbyte() int64`

GetUsedKbyte returns the UsedKbyte field if non-nil, zero value otherwise.

### GetUsedKbyteOk

`func (o *DfsTierStat) GetUsedKbyteOk() (*int64, bool)`

GetUsedKbyteOk returns a tuple with the UsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte

`func (o *DfsTierStat) SetUsedKbyte(v int64)`

SetUsedKbyte sets UsedKbyte field to given value.

### HasUsedKbyte

`func (o *DfsTierStat) HasUsedKbyte() bool`

HasUsedKbyte returns a boolean if a field has been set.

### GetWriteBandwidthKbyte

`func (o *DfsTierStat) GetWriteBandwidthKbyte() float64`

GetWriteBandwidthKbyte returns the WriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetWriteBandwidthKbyteOk

`func (o *DfsTierStat) GetWriteBandwidthKbyteOk() (*float64, bool)`

GetWriteBandwidthKbyteOk returns a tuple with the WriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidthKbyte

`func (o *DfsTierStat) SetWriteBandwidthKbyte(v float64)`

SetWriteBandwidthKbyte sets WriteBandwidthKbyte field to given value.

### HasWriteBandwidthKbyte

`func (o *DfsTierStat) HasWriteBandwidthKbyte() bool`

HasWriteBandwidthKbyte returns a boolean if a field has been set.

### GetWriteIops

`func (o *DfsTierStat) GetWriteIops() float64`

GetWriteIops returns the WriteIops field if non-nil, zero value otherwise.

### GetWriteIopsOk

`func (o *DfsTierStat) GetWriteIopsOk() (*float64, bool)`

GetWriteIopsOk returns a tuple with the WriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIops

`func (o *DfsTierStat) SetWriteIops(v float64)`

SetWriteIops sets WriteIops field to given value.

### HasWriteIops

`func (o *DfsTierStat) HasWriteIops() bool`

HasWriteIops returns a boolean if a field has been set.

### GetWriteLatencyUs

`func (o *DfsTierStat) GetWriteLatencyUs() float64`

GetWriteLatencyUs returns the WriteLatencyUs field if non-nil, zero value otherwise.

### GetWriteLatencyUsOk

`func (o *DfsTierStat) GetWriteLatencyUsOk() (*float64, bool)`

GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLatencyUs

`func (o *DfsTierStat) SetWriteLatencyUs(v float64)`

SetWriteLatencyUs sets WriteLatencyUs field to given value.

### HasWriteLatencyUs

`func (o *DfsTierStat) HasWriteLatencyUs() bool`

HasWriteLatencyUs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


