# DfsGatewayServiceStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**MemUsage** | Pointer to **float64** |  | [optional] 
**MemUsedKbyte** | Pointer to **int64** |  | [optional] 

## Methods

### NewDfsGatewayServiceStat

`func NewDfsGatewayServiceStat() *DfsGatewayServiceStat`

NewDfsGatewayServiceStat instantiates a new DfsGatewayServiceStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayServiceStatWithDefaults

`func NewDfsGatewayServiceStatWithDefaults() *DfsGatewayServiceStat`

NewDfsGatewayServiceStatWithDefaults instantiates a new DfsGatewayServiceStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *DfsGatewayServiceStat) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *DfsGatewayServiceStat) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *DfsGatewayServiceStat) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *DfsGatewayServiceStat) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCreate

`func (o *DfsGatewayServiceStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsGatewayServiceStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsGatewayServiceStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsGatewayServiceStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetMemUsage

`func (o *DfsGatewayServiceStat) GetMemUsage() float64`

GetMemUsage returns the MemUsage field if non-nil, zero value otherwise.

### GetMemUsageOk

`func (o *DfsGatewayServiceStat) GetMemUsageOk() (*float64, bool)`

GetMemUsageOk returns a tuple with the MemUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsage

`func (o *DfsGatewayServiceStat) SetMemUsage(v float64)`

SetMemUsage sets MemUsage field to given value.

### HasMemUsage

`func (o *DfsGatewayServiceStat) HasMemUsage() bool`

HasMemUsage returns a boolean if a field has been set.

### GetMemUsedKbyte

`func (o *DfsGatewayServiceStat) GetMemUsedKbyte() int64`

GetMemUsedKbyte returns the MemUsedKbyte field if non-nil, zero value otherwise.

### GetMemUsedKbyteOk

`func (o *DfsGatewayServiceStat) GetMemUsedKbyteOk() (*int64, bool)`

GetMemUsedKbyteOk returns a tuple with the MemUsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsedKbyte

`func (o *DfsGatewayServiceStat) SetMemUsedKbyte(v int64)`

SetMemUsedKbyte sets MemUsedKbyte field to given value.

### HasMemUsedKbyte

`func (o *DfsGatewayServiceStat) HasMemUsedKbyte() bool`

HasMemUsedKbyte returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


