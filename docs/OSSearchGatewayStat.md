# OSSearchGatewayStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**MemUsagePercent** | Pointer to **float64** |  | [optional] 
**ReadBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**ReadIops** | Pointer to **float64** |  | [optional] 
**ReadLatencyUs** | Pointer to **float64** |  | [optional] 
**WriteBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**WriteIops** | Pointer to **float64** |  | [optional] 
**WriteLatencyUs** | Pointer to **float64** |  | [optional] 

## Methods

### NewOSSearchGatewayStat

`func NewOSSearchGatewayStat() *OSSearchGatewayStat`

NewOSSearchGatewayStat instantiates a new OSSearchGatewayStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSSearchGatewayStatWithDefaults

`func NewOSSearchGatewayStatWithDefaults() *OSSearchGatewayStat`

NewOSSearchGatewayStatWithDefaults instantiates a new OSSearchGatewayStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *OSSearchGatewayStat) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *OSSearchGatewayStat) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *OSSearchGatewayStat) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *OSSearchGatewayStat) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCreate

`func (o *OSSearchGatewayStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSSearchGatewayStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSSearchGatewayStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSSearchGatewayStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetMemUsagePercent

`func (o *OSSearchGatewayStat) GetMemUsagePercent() float64`

GetMemUsagePercent returns the MemUsagePercent field if non-nil, zero value otherwise.

### GetMemUsagePercentOk

`func (o *OSSearchGatewayStat) GetMemUsagePercentOk() (*float64, bool)`

GetMemUsagePercentOk returns a tuple with the MemUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsagePercent

`func (o *OSSearchGatewayStat) SetMemUsagePercent(v float64)`

SetMemUsagePercent sets MemUsagePercent field to given value.

### HasMemUsagePercent

`func (o *OSSearchGatewayStat) HasMemUsagePercent() bool`

HasMemUsagePercent returns a boolean if a field has been set.

### GetReadBandwidthKbyte

`func (o *OSSearchGatewayStat) GetReadBandwidthKbyte() float64`

GetReadBandwidthKbyte returns the ReadBandwidthKbyte field if non-nil, zero value otherwise.

### GetReadBandwidthKbyteOk

`func (o *OSSearchGatewayStat) GetReadBandwidthKbyteOk() (*float64, bool)`

GetReadBandwidthKbyteOk returns a tuple with the ReadBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadBandwidthKbyte

`func (o *OSSearchGatewayStat) SetReadBandwidthKbyte(v float64)`

SetReadBandwidthKbyte sets ReadBandwidthKbyte field to given value.

### HasReadBandwidthKbyte

`func (o *OSSearchGatewayStat) HasReadBandwidthKbyte() bool`

HasReadBandwidthKbyte returns a boolean if a field has been set.

### GetReadIops

`func (o *OSSearchGatewayStat) GetReadIops() float64`

GetReadIops returns the ReadIops field if non-nil, zero value otherwise.

### GetReadIopsOk

`func (o *OSSearchGatewayStat) GetReadIopsOk() (*float64, bool)`

GetReadIopsOk returns a tuple with the ReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIops

`func (o *OSSearchGatewayStat) SetReadIops(v float64)`

SetReadIops sets ReadIops field to given value.

### HasReadIops

`func (o *OSSearchGatewayStat) HasReadIops() bool`

HasReadIops returns a boolean if a field has been set.

### GetReadLatencyUs

`func (o *OSSearchGatewayStat) GetReadLatencyUs() float64`

GetReadLatencyUs returns the ReadLatencyUs field if non-nil, zero value otherwise.

### GetReadLatencyUsOk

`func (o *OSSearchGatewayStat) GetReadLatencyUsOk() (*float64, bool)`

GetReadLatencyUsOk returns a tuple with the ReadLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLatencyUs

`func (o *OSSearchGatewayStat) SetReadLatencyUs(v float64)`

SetReadLatencyUs sets ReadLatencyUs field to given value.

### HasReadLatencyUs

`func (o *OSSearchGatewayStat) HasReadLatencyUs() bool`

HasReadLatencyUs returns a boolean if a field has been set.

### GetWriteBandwidthKbyte

`func (o *OSSearchGatewayStat) GetWriteBandwidthKbyte() float64`

GetWriteBandwidthKbyte returns the WriteBandwidthKbyte field if non-nil, zero value otherwise.

### GetWriteBandwidthKbyteOk

`func (o *OSSearchGatewayStat) GetWriteBandwidthKbyteOk() (*float64, bool)`

GetWriteBandwidthKbyteOk returns a tuple with the WriteBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBandwidthKbyte

`func (o *OSSearchGatewayStat) SetWriteBandwidthKbyte(v float64)`

SetWriteBandwidthKbyte sets WriteBandwidthKbyte field to given value.

### HasWriteBandwidthKbyte

`func (o *OSSearchGatewayStat) HasWriteBandwidthKbyte() bool`

HasWriteBandwidthKbyte returns a boolean if a field has been set.

### GetWriteIops

`func (o *OSSearchGatewayStat) GetWriteIops() float64`

GetWriteIops returns the WriteIops field if non-nil, zero value otherwise.

### GetWriteIopsOk

`func (o *OSSearchGatewayStat) GetWriteIopsOk() (*float64, bool)`

GetWriteIopsOk returns a tuple with the WriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIops

`func (o *OSSearchGatewayStat) SetWriteIops(v float64)`

SetWriteIops sets WriteIops field to given value.

### HasWriteIops

`func (o *OSSearchGatewayStat) HasWriteIops() bool`

HasWriteIops returns a boolean if a field has been set.

### GetWriteLatencyUs

`func (o *OSSearchGatewayStat) GetWriteLatencyUs() float64`

GetWriteLatencyUs returns the WriteLatencyUs field if non-nil, zero value otherwise.

### GetWriteLatencyUsOk

`func (o *OSSearchGatewayStat) GetWriteLatencyUsOk() (*float64, bool)`

GetWriteLatencyUsOk returns a tuple with the WriteLatencyUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLatencyUs

`func (o *OSSearchGatewayStat) SetWriteLatencyUs(v float64)`

SetWriteLatencyUs sets WriteLatencyUs field to given value.

### HasWriteLatencyUs

`func (o *OSSearchGatewayStat) HasWriteLatencyUs() bool`

HasWriteLatencyUs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


