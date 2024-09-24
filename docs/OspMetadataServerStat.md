# OspMetadataServerStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DiskUtil** | Pointer to **float64** |  | [optional] 
**MemUsagePercent** | Pointer to **float64** |  | [optional] 
**RecvBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**SendBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**TotalSizeByte** | Pointer to **int64** |  | [optional] 
**UsedSizeByte** | Pointer to **int64** |  | [optional] 

## Methods

### NewOspMetadataServerStat

`func NewOspMetadataServerStat() *OspMetadataServerStat`

NewOspMetadataServerStat instantiates a new OspMetadataServerStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspMetadataServerStatWithDefaults

`func NewOspMetadataServerStatWithDefaults() *OspMetadataServerStat`

NewOspMetadataServerStatWithDefaults instantiates a new OspMetadataServerStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *OspMetadataServerStat) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *OspMetadataServerStat) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *OspMetadataServerStat) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *OspMetadataServerStat) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCreate

`func (o *OspMetadataServerStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OspMetadataServerStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OspMetadataServerStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OspMetadataServerStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDiskUtil

`func (o *OspMetadataServerStat) GetDiskUtil() float64`

GetDiskUtil returns the DiskUtil field if non-nil, zero value otherwise.

### GetDiskUtilOk

`func (o *OspMetadataServerStat) GetDiskUtilOk() (*float64, bool)`

GetDiskUtilOk returns a tuple with the DiskUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUtil

`func (o *OspMetadataServerStat) SetDiskUtil(v float64)`

SetDiskUtil sets DiskUtil field to given value.

### HasDiskUtil

`func (o *OspMetadataServerStat) HasDiskUtil() bool`

HasDiskUtil returns a boolean if a field has been set.

### GetMemUsagePercent

`func (o *OspMetadataServerStat) GetMemUsagePercent() float64`

GetMemUsagePercent returns the MemUsagePercent field if non-nil, zero value otherwise.

### GetMemUsagePercentOk

`func (o *OspMetadataServerStat) GetMemUsagePercentOk() (*float64, bool)`

GetMemUsagePercentOk returns a tuple with the MemUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsagePercent

`func (o *OspMetadataServerStat) SetMemUsagePercent(v float64)`

SetMemUsagePercent sets MemUsagePercent field to given value.

### HasMemUsagePercent

`func (o *OspMetadataServerStat) HasMemUsagePercent() bool`

HasMemUsagePercent returns a boolean if a field has been set.

### GetRecvBandwidthKbyte

`func (o *OspMetadataServerStat) GetRecvBandwidthKbyte() float64`

GetRecvBandwidthKbyte returns the RecvBandwidthKbyte field if non-nil, zero value otherwise.

### GetRecvBandwidthKbyteOk

`func (o *OspMetadataServerStat) GetRecvBandwidthKbyteOk() (*float64, bool)`

GetRecvBandwidthKbyteOk returns a tuple with the RecvBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvBandwidthKbyte

`func (o *OspMetadataServerStat) SetRecvBandwidthKbyte(v float64)`

SetRecvBandwidthKbyte sets RecvBandwidthKbyte field to given value.

### HasRecvBandwidthKbyte

`func (o *OspMetadataServerStat) HasRecvBandwidthKbyte() bool`

HasRecvBandwidthKbyte returns a boolean if a field has been set.

### GetSendBandwidthKbyte

`func (o *OspMetadataServerStat) GetSendBandwidthKbyte() float64`

GetSendBandwidthKbyte returns the SendBandwidthKbyte field if non-nil, zero value otherwise.

### GetSendBandwidthKbyteOk

`func (o *OspMetadataServerStat) GetSendBandwidthKbyteOk() (*float64, bool)`

GetSendBandwidthKbyteOk returns a tuple with the SendBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBandwidthKbyte

`func (o *OspMetadataServerStat) SetSendBandwidthKbyte(v float64)`

SetSendBandwidthKbyte sets SendBandwidthKbyte field to given value.

### HasSendBandwidthKbyte

`func (o *OspMetadataServerStat) HasSendBandwidthKbyte() bool`

HasSendBandwidthKbyte returns a boolean if a field has been set.

### GetTotalSizeByte

`func (o *OspMetadataServerStat) GetTotalSizeByte() int64`

GetTotalSizeByte returns the TotalSizeByte field if non-nil, zero value otherwise.

### GetTotalSizeByteOk

`func (o *OspMetadataServerStat) GetTotalSizeByteOk() (*int64, bool)`

GetTotalSizeByteOk returns a tuple with the TotalSizeByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeByte

`func (o *OspMetadataServerStat) SetTotalSizeByte(v int64)`

SetTotalSizeByte sets TotalSizeByte field to given value.

### HasTotalSizeByte

`func (o *OspMetadataServerStat) HasTotalSizeByte() bool`

HasTotalSizeByte returns a boolean if a field has been set.

### GetUsedSizeByte

`func (o *OspMetadataServerStat) GetUsedSizeByte() int64`

GetUsedSizeByte returns the UsedSizeByte field if non-nil, zero value otherwise.

### GetUsedSizeByteOk

`func (o *OspMetadataServerStat) GetUsedSizeByteOk() (*int64, bool)`

GetUsedSizeByteOk returns a tuple with the UsedSizeByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSizeByte

`func (o *OspMetadataServerStat) SetUsedSizeByte(v int64)`

SetUsedSizeByte sets UsedSizeByte field to given value.

### HasUsedSizeByte

`func (o *OspMetadataServerStat) HasUsedSizeByte() bool`

HasUsedSizeByte returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


