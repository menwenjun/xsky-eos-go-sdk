# HostStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextSwitchPs** | Pointer to **float64** |  | [optional] 
**CpuUtil** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**InterruptsPs** | Pointer to **float64** |  | [optional] 
**Load15min** | Pointer to **float64** |  | [optional] 
**Load1min** | Pointer to **float64** |  | [optional] 
**Load5min** | Pointer to **float64** |  | [optional] 
**MemUsagePercent** | Pointer to **float64** |  | [optional] 
**MetaDockerUsedKbyte** | Pointer to **float64** |  | [optional] 
**MetaElasticsearchUsedKbyte** | Pointer to **float64** |  | [optional] 
**MetaEtcdUsedKbyte** | Pointer to **float64** |  | [optional] 
**MetaLogUsedKbyte** | Pointer to **float64** |  | [optional] 
**MetaMonUsedKbyte** | Pointer to **float64** |  | [optional] 
**MetaPostgresUsedKbyte** | Pointer to **float64** |  | [optional] 
**MetaPrometheusUsedKbyte** | Pointer to **float64** |  | [optional] 
**NetworkDropPps** | Pointer to **float64** |  | [optional] 
**NetworkDropRatio** | Pointer to **float64** |  | [optional] 
**NetworkErrorPps** | Pointer to **float64** |  | [optional] 
**NetworkErrorRatio** | Pointer to **float64** |  | [optional] 
**NetworkRxBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**NetworkRxDropPps** | Pointer to **float64** |  | [optional] 
**NetworkRxDropRatio** | Pointer to **float64** |  | [optional] 
**NetworkRxErrorPps** | Pointer to **float64** |  | [optional] 
**NetworkRxErrorRatio** | Pointer to **float64** |  | [optional] 
**NetworkRxPps** | Pointer to **float64** |  | [optional] 
**NetworkTxBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**NetworkTxDropPps** | Pointer to **float64** |  | [optional] 
**NetworkTxDropRatio** | Pointer to **float64** |  | [optional] 
**NetworkTxErrorPps** | Pointer to **float64** |  | [optional] 
**NetworkTxErrorRatio** | Pointer to **float64** |  | [optional] 
**NetworkTxPps** | Pointer to **float64** |  | [optional] 
**PagePagingPs** | Pointer to **float64** |  | [optional] 

## Methods

### NewHostStat

`func NewHostStat() *HostStat`

NewHostStat instantiates a new HostStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostStatWithDefaults

`func NewHostStatWithDefaults() *HostStat`

NewHostStatWithDefaults instantiates a new HostStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextSwitchPs

`func (o *HostStat) GetContextSwitchPs() float64`

GetContextSwitchPs returns the ContextSwitchPs field if non-nil, zero value otherwise.

### GetContextSwitchPsOk

`func (o *HostStat) GetContextSwitchPsOk() (*float64, bool)`

GetContextSwitchPsOk returns a tuple with the ContextSwitchPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextSwitchPs

`func (o *HostStat) SetContextSwitchPs(v float64)`

SetContextSwitchPs sets ContextSwitchPs field to given value.

### HasContextSwitchPs

`func (o *HostStat) HasContextSwitchPs() bool`

HasContextSwitchPs returns a boolean if a field has been set.

### GetCpuUtil

`func (o *HostStat) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *HostStat) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *HostStat) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *HostStat) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCreate

`func (o *HostStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *HostStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *HostStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *HostStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetInterruptsPs

`func (o *HostStat) GetInterruptsPs() float64`

GetInterruptsPs returns the InterruptsPs field if non-nil, zero value otherwise.

### GetInterruptsPsOk

`func (o *HostStat) GetInterruptsPsOk() (*float64, bool)`

GetInterruptsPsOk returns a tuple with the InterruptsPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptsPs

`func (o *HostStat) SetInterruptsPs(v float64)`

SetInterruptsPs sets InterruptsPs field to given value.

### HasInterruptsPs

`func (o *HostStat) HasInterruptsPs() bool`

HasInterruptsPs returns a boolean if a field has been set.

### GetLoad15min

`func (o *HostStat) GetLoad15min() float64`

GetLoad15min returns the Load15min field if non-nil, zero value otherwise.

### GetLoad15minOk

`func (o *HostStat) GetLoad15minOk() (*float64, bool)`

GetLoad15minOk returns a tuple with the Load15min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad15min

`func (o *HostStat) SetLoad15min(v float64)`

SetLoad15min sets Load15min field to given value.

### HasLoad15min

`func (o *HostStat) HasLoad15min() bool`

HasLoad15min returns a boolean if a field has been set.

### GetLoad1min

`func (o *HostStat) GetLoad1min() float64`

GetLoad1min returns the Load1min field if non-nil, zero value otherwise.

### GetLoad1minOk

`func (o *HostStat) GetLoad1minOk() (*float64, bool)`

GetLoad1minOk returns a tuple with the Load1min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad1min

`func (o *HostStat) SetLoad1min(v float64)`

SetLoad1min sets Load1min field to given value.

### HasLoad1min

`func (o *HostStat) HasLoad1min() bool`

HasLoad1min returns a boolean if a field has been set.

### GetLoad5min

`func (o *HostStat) GetLoad5min() float64`

GetLoad5min returns the Load5min field if non-nil, zero value otherwise.

### GetLoad5minOk

`func (o *HostStat) GetLoad5minOk() (*float64, bool)`

GetLoad5minOk returns a tuple with the Load5min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad5min

`func (o *HostStat) SetLoad5min(v float64)`

SetLoad5min sets Load5min field to given value.

### HasLoad5min

`func (o *HostStat) HasLoad5min() bool`

HasLoad5min returns a boolean if a field has been set.

### GetMemUsagePercent

`func (o *HostStat) GetMemUsagePercent() float64`

GetMemUsagePercent returns the MemUsagePercent field if non-nil, zero value otherwise.

### GetMemUsagePercentOk

`func (o *HostStat) GetMemUsagePercentOk() (*float64, bool)`

GetMemUsagePercentOk returns a tuple with the MemUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsagePercent

`func (o *HostStat) SetMemUsagePercent(v float64)`

SetMemUsagePercent sets MemUsagePercent field to given value.

### HasMemUsagePercent

`func (o *HostStat) HasMemUsagePercent() bool`

HasMemUsagePercent returns a boolean if a field has been set.

### GetMetaDockerUsedKbyte

`func (o *HostStat) GetMetaDockerUsedKbyte() float64`

GetMetaDockerUsedKbyte returns the MetaDockerUsedKbyte field if non-nil, zero value otherwise.

### GetMetaDockerUsedKbyteOk

`func (o *HostStat) GetMetaDockerUsedKbyteOk() (*float64, bool)`

GetMetaDockerUsedKbyteOk returns a tuple with the MetaDockerUsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDockerUsedKbyte

`func (o *HostStat) SetMetaDockerUsedKbyte(v float64)`

SetMetaDockerUsedKbyte sets MetaDockerUsedKbyte field to given value.

### HasMetaDockerUsedKbyte

`func (o *HostStat) HasMetaDockerUsedKbyte() bool`

HasMetaDockerUsedKbyte returns a boolean if a field has been set.

### GetMetaElasticsearchUsedKbyte

`func (o *HostStat) GetMetaElasticsearchUsedKbyte() float64`

GetMetaElasticsearchUsedKbyte returns the MetaElasticsearchUsedKbyte field if non-nil, zero value otherwise.

### GetMetaElasticsearchUsedKbyteOk

`func (o *HostStat) GetMetaElasticsearchUsedKbyteOk() (*float64, bool)`

GetMetaElasticsearchUsedKbyteOk returns a tuple with the MetaElasticsearchUsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaElasticsearchUsedKbyte

`func (o *HostStat) SetMetaElasticsearchUsedKbyte(v float64)`

SetMetaElasticsearchUsedKbyte sets MetaElasticsearchUsedKbyte field to given value.

### HasMetaElasticsearchUsedKbyte

`func (o *HostStat) HasMetaElasticsearchUsedKbyte() bool`

HasMetaElasticsearchUsedKbyte returns a boolean if a field has been set.

### GetMetaEtcdUsedKbyte

`func (o *HostStat) GetMetaEtcdUsedKbyte() float64`

GetMetaEtcdUsedKbyte returns the MetaEtcdUsedKbyte field if non-nil, zero value otherwise.

### GetMetaEtcdUsedKbyteOk

`func (o *HostStat) GetMetaEtcdUsedKbyteOk() (*float64, bool)`

GetMetaEtcdUsedKbyteOk returns a tuple with the MetaEtcdUsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaEtcdUsedKbyte

`func (o *HostStat) SetMetaEtcdUsedKbyte(v float64)`

SetMetaEtcdUsedKbyte sets MetaEtcdUsedKbyte field to given value.

### HasMetaEtcdUsedKbyte

`func (o *HostStat) HasMetaEtcdUsedKbyte() bool`

HasMetaEtcdUsedKbyte returns a boolean if a field has been set.

### GetMetaLogUsedKbyte

`func (o *HostStat) GetMetaLogUsedKbyte() float64`

GetMetaLogUsedKbyte returns the MetaLogUsedKbyte field if non-nil, zero value otherwise.

### GetMetaLogUsedKbyteOk

`func (o *HostStat) GetMetaLogUsedKbyteOk() (*float64, bool)`

GetMetaLogUsedKbyteOk returns a tuple with the MetaLogUsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaLogUsedKbyte

`func (o *HostStat) SetMetaLogUsedKbyte(v float64)`

SetMetaLogUsedKbyte sets MetaLogUsedKbyte field to given value.

### HasMetaLogUsedKbyte

`func (o *HostStat) HasMetaLogUsedKbyte() bool`

HasMetaLogUsedKbyte returns a boolean if a field has been set.

### GetMetaMonUsedKbyte

`func (o *HostStat) GetMetaMonUsedKbyte() float64`

GetMetaMonUsedKbyte returns the MetaMonUsedKbyte field if non-nil, zero value otherwise.

### GetMetaMonUsedKbyteOk

`func (o *HostStat) GetMetaMonUsedKbyteOk() (*float64, bool)`

GetMetaMonUsedKbyteOk returns a tuple with the MetaMonUsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaMonUsedKbyte

`func (o *HostStat) SetMetaMonUsedKbyte(v float64)`

SetMetaMonUsedKbyte sets MetaMonUsedKbyte field to given value.

### HasMetaMonUsedKbyte

`func (o *HostStat) HasMetaMonUsedKbyte() bool`

HasMetaMonUsedKbyte returns a boolean if a field has been set.

### GetMetaPostgresUsedKbyte

`func (o *HostStat) GetMetaPostgresUsedKbyte() float64`

GetMetaPostgresUsedKbyte returns the MetaPostgresUsedKbyte field if non-nil, zero value otherwise.

### GetMetaPostgresUsedKbyteOk

`func (o *HostStat) GetMetaPostgresUsedKbyteOk() (*float64, bool)`

GetMetaPostgresUsedKbyteOk returns a tuple with the MetaPostgresUsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaPostgresUsedKbyte

`func (o *HostStat) SetMetaPostgresUsedKbyte(v float64)`

SetMetaPostgresUsedKbyte sets MetaPostgresUsedKbyte field to given value.

### HasMetaPostgresUsedKbyte

`func (o *HostStat) HasMetaPostgresUsedKbyte() bool`

HasMetaPostgresUsedKbyte returns a boolean if a field has been set.

### GetMetaPrometheusUsedKbyte

`func (o *HostStat) GetMetaPrometheusUsedKbyte() float64`

GetMetaPrometheusUsedKbyte returns the MetaPrometheusUsedKbyte field if non-nil, zero value otherwise.

### GetMetaPrometheusUsedKbyteOk

`func (o *HostStat) GetMetaPrometheusUsedKbyteOk() (*float64, bool)`

GetMetaPrometheusUsedKbyteOk returns a tuple with the MetaPrometheusUsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaPrometheusUsedKbyte

`func (o *HostStat) SetMetaPrometheusUsedKbyte(v float64)`

SetMetaPrometheusUsedKbyte sets MetaPrometheusUsedKbyte field to given value.

### HasMetaPrometheusUsedKbyte

`func (o *HostStat) HasMetaPrometheusUsedKbyte() bool`

HasMetaPrometheusUsedKbyte returns a boolean if a field has been set.

### GetNetworkDropPps

`func (o *HostStat) GetNetworkDropPps() float64`

GetNetworkDropPps returns the NetworkDropPps field if non-nil, zero value otherwise.

### GetNetworkDropPpsOk

`func (o *HostStat) GetNetworkDropPpsOk() (*float64, bool)`

GetNetworkDropPpsOk returns a tuple with the NetworkDropPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDropPps

`func (o *HostStat) SetNetworkDropPps(v float64)`

SetNetworkDropPps sets NetworkDropPps field to given value.

### HasNetworkDropPps

`func (o *HostStat) HasNetworkDropPps() bool`

HasNetworkDropPps returns a boolean if a field has been set.

### GetNetworkDropRatio

`func (o *HostStat) GetNetworkDropRatio() float64`

GetNetworkDropRatio returns the NetworkDropRatio field if non-nil, zero value otherwise.

### GetNetworkDropRatioOk

`func (o *HostStat) GetNetworkDropRatioOk() (*float64, bool)`

GetNetworkDropRatioOk returns a tuple with the NetworkDropRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDropRatio

`func (o *HostStat) SetNetworkDropRatio(v float64)`

SetNetworkDropRatio sets NetworkDropRatio field to given value.

### HasNetworkDropRatio

`func (o *HostStat) HasNetworkDropRatio() bool`

HasNetworkDropRatio returns a boolean if a field has been set.

### GetNetworkErrorPps

`func (o *HostStat) GetNetworkErrorPps() float64`

GetNetworkErrorPps returns the NetworkErrorPps field if non-nil, zero value otherwise.

### GetNetworkErrorPpsOk

`func (o *HostStat) GetNetworkErrorPpsOk() (*float64, bool)`

GetNetworkErrorPpsOk returns a tuple with the NetworkErrorPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkErrorPps

`func (o *HostStat) SetNetworkErrorPps(v float64)`

SetNetworkErrorPps sets NetworkErrorPps field to given value.

### HasNetworkErrorPps

`func (o *HostStat) HasNetworkErrorPps() bool`

HasNetworkErrorPps returns a boolean if a field has been set.

### GetNetworkErrorRatio

`func (o *HostStat) GetNetworkErrorRatio() float64`

GetNetworkErrorRatio returns the NetworkErrorRatio field if non-nil, zero value otherwise.

### GetNetworkErrorRatioOk

`func (o *HostStat) GetNetworkErrorRatioOk() (*float64, bool)`

GetNetworkErrorRatioOk returns a tuple with the NetworkErrorRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkErrorRatio

`func (o *HostStat) SetNetworkErrorRatio(v float64)`

SetNetworkErrorRatio sets NetworkErrorRatio field to given value.

### HasNetworkErrorRatio

`func (o *HostStat) HasNetworkErrorRatio() bool`

HasNetworkErrorRatio returns a boolean if a field has been set.

### GetNetworkRxBandwidthKbyte

`func (o *HostStat) GetNetworkRxBandwidthKbyte() float64`

GetNetworkRxBandwidthKbyte returns the NetworkRxBandwidthKbyte field if non-nil, zero value otherwise.

### GetNetworkRxBandwidthKbyteOk

`func (o *HostStat) GetNetworkRxBandwidthKbyteOk() (*float64, bool)`

GetNetworkRxBandwidthKbyteOk returns a tuple with the NetworkRxBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRxBandwidthKbyte

`func (o *HostStat) SetNetworkRxBandwidthKbyte(v float64)`

SetNetworkRxBandwidthKbyte sets NetworkRxBandwidthKbyte field to given value.

### HasNetworkRxBandwidthKbyte

`func (o *HostStat) HasNetworkRxBandwidthKbyte() bool`

HasNetworkRxBandwidthKbyte returns a boolean if a field has been set.

### GetNetworkRxDropPps

`func (o *HostStat) GetNetworkRxDropPps() float64`

GetNetworkRxDropPps returns the NetworkRxDropPps field if non-nil, zero value otherwise.

### GetNetworkRxDropPpsOk

`func (o *HostStat) GetNetworkRxDropPpsOk() (*float64, bool)`

GetNetworkRxDropPpsOk returns a tuple with the NetworkRxDropPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRxDropPps

`func (o *HostStat) SetNetworkRxDropPps(v float64)`

SetNetworkRxDropPps sets NetworkRxDropPps field to given value.

### HasNetworkRxDropPps

`func (o *HostStat) HasNetworkRxDropPps() bool`

HasNetworkRxDropPps returns a boolean if a field has been set.

### GetNetworkRxDropRatio

`func (o *HostStat) GetNetworkRxDropRatio() float64`

GetNetworkRxDropRatio returns the NetworkRxDropRatio field if non-nil, zero value otherwise.

### GetNetworkRxDropRatioOk

`func (o *HostStat) GetNetworkRxDropRatioOk() (*float64, bool)`

GetNetworkRxDropRatioOk returns a tuple with the NetworkRxDropRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRxDropRatio

`func (o *HostStat) SetNetworkRxDropRatio(v float64)`

SetNetworkRxDropRatio sets NetworkRxDropRatio field to given value.

### HasNetworkRxDropRatio

`func (o *HostStat) HasNetworkRxDropRatio() bool`

HasNetworkRxDropRatio returns a boolean if a field has been set.

### GetNetworkRxErrorPps

`func (o *HostStat) GetNetworkRxErrorPps() float64`

GetNetworkRxErrorPps returns the NetworkRxErrorPps field if non-nil, zero value otherwise.

### GetNetworkRxErrorPpsOk

`func (o *HostStat) GetNetworkRxErrorPpsOk() (*float64, bool)`

GetNetworkRxErrorPpsOk returns a tuple with the NetworkRxErrorPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRxErrorPps

`func (o *HostStat) SetNetworkRxErrorPps(v float64)`

SetNetworkRxErrorPps sets NetworkRxErrorPps field to given value.

### HasNetworkRxErrorPps

`func (o *HostStat) HasNetworkRxErrorPps() bool`

HasNetworkRxErrorPps returns a boolean if a field has been set.

### GetNetworkRxErrorRatio

`func (o *HostStat) GetNetworkRxErrorRatio() float64`

GetNetworkRxErrorRatio returns the NetworkRxErrorRatio field if non-nil, zero value otherwise.

### GetNetworkRxErrorRatioOk

`func (o *HostStat) GetNetworkRxErrorRatioOk() (*float64, bool)`

GetNetworkRxErrorRatioOk returns a tuple with the NetworkRxErrorRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRxErrorRatio

`func (o *HostStat) SetNetworkRxErrorRatio(v float64)`

SetNetworkRxErrorRatio sets NetworkRxErrorRatio field to given value.

### HasNetworkRxErrorRatio

`func (o *HostStat) HasNetworkRxErrorRatio() bool`

HasNetworkRxErrorRatio returns a boolean if a field has been set.

### GetNetworkRxPps

`func (o *HostStat) GetNetworkRxPps() float64`

GetNetworkRxPps returns the NetworkRxPps field if non-nil, zero value otherwise.

### GetNetworkRxPpsOk

`func (o *HostStat) GetNetworkRxPpsOk() (*float64, bool)`

GetNetworkRxPpsOk returns a tuple with the NetworkRxPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRxPps

`func (o *HostStat) SetNetworkRxPps(v float64)`

SetNetworkRxPps sets NetworkRxPps field to given value.

### HasNetworkRxPps

`func (o *HostStat) HasNetworkRxPps() bool`

HasNetworkRxPps returns a boolean if a field has been set.

### GetNetworkTxBandwidthKbyte

`func (o *HostStat) GetNetworkTxBandwidthKbyte() float64`

GetNetworkTxBandwidthKbyte returns the NetworkTxBandwidthKbyte field if non-nil, zero value otherwise.

### GetNetworkTxBandwidthKbyteOk

`func (o *HostStat) GetNetworkTxBandwidthKbyteOk() (*float64, bool)`

GetNetworkTxBandwidthKbyteOk returns a tuple with the NetworkTxBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTxBandwidthKbyte

`func (o *HostStat) SetNetworkTxBandwidthKbyte(v float64)`

SetNetworkTxBandwidthKbyte sets NetworkTxBandwidthKbyte field to given value.

### HasNetworkTxBandwidthKbyte

`func (o *HostStat) HasNetworkTxBandwidthKbyte() bool`

HasNetworkTxBandwidthKbyte returns a boolean if a field has been set.

### GetNetworkTxDropPps

`func (o *HostStat) GetNetworkTxDropPps() float64`

GetNetworkTxDropPps returns the NetworkTxDropPps field if non-nil, zero value otherwise.

### GetNetworkTxDropPpsOk

`func (o *HostStat) GetNetworkTxDropPpsOk() (*float64, bool)`

GetNetworkTxDropPpsOk returns a tuple with the NetworkTxDropPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTxDropPps

`func (o *HostStat) SetNetworkTxDropPps(v float64)`

SetNetworkTxDropPps sets NetworkTxDropPps field to given value.

### HasNetworkTxDropPps

`func (o *HostStat) HasNetworkTxDropPps() bool`

HasNetworkTxDropPps returns a boolean if a field has been set.

### GetNetworkTxDropRatio

`func (o *HostStat) GetNetworkTxDropRatio() float64`

GetNetworkTxDropRatio returns the NetworkTxDropRatio field if non-nil, zero value otherwise.

### GetNetworkTxDropRatioOk

`func (o *HostStat) GetNetworkTxDropRatioOk() (*float64, bool)`

GetNetworkTxDropRatioOk returns a tuple with the NetworkTxDropRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTxDropRatio

`func (o *HostStat) SetNetworkTxDropRatio(v float64)`

SetNetworkTxDropRatio sets NetworkTxDropRatio field to given value.

### HasNetworkTxDropRatio

`func (o *HostStat) HasNetworkTxDropRatio() bool`

HasNetworkTxDropRatio returns a boolean if a field has been set.

### GetNetworkTxErrorPps

`func (o *HostStat) GetNetworkTxErrorPps() float64`

GetNetworkTxErrorPps returns the NetworkTxErrorPps field if non-nil, zero value otherwise.

### GetNetworkTxErrorPpsOk

`func (o *HostStat) GetNetworkTxErrorPpsOk() (*float64, bool)`

GetNetworkTxErrorPpsOk returns a tuple with the NetworkTxErrorPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTxErrorPps

`func (o *HostStat) SetNetworkTxErrorPps(v float64)`

SetNetworkTxErrorPps sets NetworkTxErrorPps field to given value.

### HasNetworkTxErrorPps

`func (o *HostStat) HasNetworkTxErrorPps() bool`

HasNetworkTxErrorPps returns a boolean if a field has been set.

### GetNetworkTxErrorRatio

`func (o *HostStat) GetNetworkTxErrorRatio() float64`

GetNetworkTxErrorRatio returns the NetworkTxErrorRatio field if non-nil, zero value otherwise.

### GetNetworkTxErrorRatioOk

`func (o *HostStat) GetNetworkTxErrorRatioOk() (*float64, bool)`

GetNetworkTxErrorRatioOk returns a tuple with the NetworkTxErrorRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTxErrorRatio

`func (o *HostStat) SetNetworkTxErrorRatio(v float64)`

SetNetworkTxErrorRatio sets NetworkTxErrorRatio field to given value.

### HasNetworkTxErrorRatio

`func (o *HostStat) HasNetworkTxErrorRatio() bool`

HasNetworkTxErrorRatio returns a boolean if a field has been set.

### GetNetworkTxPps

`func (o *HostStat) GetNetworkTxPps() float64`

GetNetworkTxPps returns the NetworkTxPps field if non-nil, zero value otherwise.

### GetNetworkTxPpsOk

`func (o *HostStat) GetNetworkTxPpsOk() (*float64, bool)`

GetNetworkTxPpsOk returns a tuple with the NetworkTxPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTxPps

`func (o *HostStat) SetNetworkTxPps(v float64)`

SetNetworkTxPps sets NetworkTxPps field to given value.

### HasNetworkTxPps

`func (o *HostStat) HasNetworkTxPps() bool`

HasNetworkTxPps returns a boolean if a field has been set.

### GetPagePagingPs

`func (o *HostStat) GetPagePagingPs() float64`

GetPagePagingPs returns the PagePagingPs field if non-nil, zero value otherwise.

### GetPagePagingPsOk

`func (o *HostStat) GetPagePagingPsOk() (*float64, bool)`

GetPagePagingPsOk returns a tuple with the PagePagingPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagePagingPs

`func (o *HostStat) SetPagePagingPs(v float64)`

SetPagePagingPs sets PagePagingPs field to given value.

### HasPagePagingPs

`func (o *HostStat) HasPagePagingPs() bool`

HasPagePagingPs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


