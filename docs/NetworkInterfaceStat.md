# NetworkInterfaceStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**DropPps** | Pointer to **float64** |  | [optional] 
**ErrorPps** | Pointer to **float64** |  | [optional] 
**RxBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**RxDropPps** | Pointer to **float64** |  | [optional] 
**RxErrorPps** | Pointer to **float64** |  | [optional] 
**RxPps** | Pointer to **float64** |  | [optional] 
**TxBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**TxDropPps** | Pointer to **float64** |  | [optional] 
**TxErrorPps** | Pointer to **float64** |  | [optional] 
**TxPps** | Pointer to **float64** |  | [optional] 

## Methods

### NewNetworkInterfaceStat

`func NewNetworkInterfaceStat() *NetworkInterfaceStat`

NewNetworkInterfaceStat instantiates a new NetworkInterfaceStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceStatWithDefaults

`func NewNetworkInterfaceStatWithDefaults() *NetworkInterfaceStat`

NewNetworkInterfaceStatWithDefaults instantiates a new NetworkInterfaceStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *NetworkInterfaceStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *NetworkInterfaceStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *NetworkInterfaceStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *NetworkInterfaceStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDropPps

`func (o *NetworkInterfaceStat) GetDropPps() float64`

GetDropPps returns the DropPps field if non-nil, zero value otherwise.

### GetDropPpsOk

`func (o *NetworkInterfaceStat) GetDropPpsOk() (*float64, bool)`

GetDropPpsOk returns a tuple with the DropPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropPps

`func (o *NetworkInterfaceStat) SetDropPps(v float64)`

SetDropPps sets DropPps field to given value.

### HasDropPps

`func (o *NetworkInterfaceStat) HasDropPps() bool`

HasDropPps returns a boolean if a field has been set.

### GetErrorPps

`func (o *NetworkInterfaceStat) GetErrorPps() float64`

GetErrorPps returns the ErrorPps field if non-nil, zero value otherwise.

### GetErrorPpsOk

`func (o *NetworkInterfaceStat) GetErrorPpsOk() (*float64, bool)`

GetErrorPpsOk returns a tuple with the ErrorPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPps

`func (o *NetworkInterfaceStat) SetErrorPps(v float64)`

SetErrorPps sets ErrorPps field to given value.

### HasErrorPps

`func (o *NetworkInterfaceStat) HasErrorPps() bool`

HasErrorPps returns a boolean if a field has been set.

### GetRxBandwidthKbyte

`func (o *NetworkInterfaceStat) GetRxBandwidthKbyte() float64`

GetRxBandwidthKbyte returns the RxBandwidthKbyte field if non-nil, zero value otherwise.

### GetRxBandwidthKbyteOk

`func (o *NetworkInterfaceStat) GetRxBandwidthKbyteOk() (*float64, bool)`

GetRxBandwidthKbyteOk returns a tuple with the RxBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBandwidthKbyte

`func (o *NetworkInterfaceStat) SetRxBandwidthKbyte(v float64)`

SetRxBandwidthKbyte sets RxBandwidthKbyte field to given value.

### HasRxBandwidthKbyte

`func (o *NetworkInterfaceStat) HasRxBandwidthKbyte() bool`

HasRxBandwidthKbyte returns a boolean if a field has been set.

### GetRxDropPps

`func (o *NetworkInterfaceStat) GetRxDropPps() float64`

GetRxDropPps returns the RxDropPps field if non-nil, zero value otherwise.

### GetRxDropPpsOk

`func (o *NetworkInterfaceStat) GetRxDropPpsOk() (*float64, bool)`

GetRxDropPpsOk returns a tuple with the RxDropPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDropPps

`func (o *NetworkInterfaceStat) SetRxDropPps(v float64)`

SetRxDropPps sets RxDropPps field to given value.

### HasRxDropPps

`func (o *NetworkInterfaceStat) HasRxDropPps() bool`

HasRxDropPps returns a boolean if a field has been set.

### GetRxErrorPps

`func (o *NetworkInterfaceStat) GetRxErrorPps() float64`

GetRxErrorPps returns the RxErrorPps field if non-nil, zero value otherwise.

### GetRxErrorPpsOk

`func (o *NetworkInterfaceStat) GetRxErrorPpsOk() (*float64, bool)`

GetRxErrorPpsOk returns a tuple with the RxErrorPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrorPps

`func (o *NetworkInterfaceStat) SetRxErrorPps(v float64)`

SetRxErrorPps sets RxErrorPps field to given value.

### HasRxErrorPps

`func (o *NetworkInterfaceStat) HasRxErrorPps() bool`

HasRxErrorPps returns a boolean if a field has been set.

### GetRxPps

`func (o *NetworkInterfaceStat) GetRxPps() float64`

GetRxPps returns the RxPps field if non-nil, zero value otherwise.

### GetRxPpsOk

`func (o *NetworkInterfaceStat) GetRxPpsOk() (*float64, bool)`

GetRxPpsOk returns a tuple with the RxPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPps

`func (o *NetworkInterfaceStat) SetRxPps(v float64)`

SetRxPps sets RxPps field to given value.

### HasRxPps

`func (o *NetworkInterfaceStat) HasRxPps() bool`

HasRxPps returns a boolean if a field has been set.

### GetTxBandwidthKbyte

`func (o *NetworkInterfaceStat) GetTxBandwidthKbyte() float64`

GetTxBandwidthKbyte returns the TxBandwidthKbyte field if non-nil, zero value otherwise.

### GetTxBandwidthKbyteOk

`func (o *NetworkInterfaceStat) GetTxBandwidthKbyteOk() (*float64, bool)`

GetTxBandwidthKbyteOk returns a tuple with the TxBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBandwidthKbyte

`func (o *NetworkInterfaceStat) SetTxBandwidthKbyte(v float64)`

SetTxBandwidthKbyte sets TxBandwidthKbyte field to given value.

### HasTxBandwidthKbyte

`func (o *NetworkInterfaceStat) HasTxBandwidthKbyte() bool`

HasTxBandwidthKbyte returns a boolean if a field has been set.

### GetTxDropPps

`func (o *NetworkInterfaceStat) GetTxDropPps() float64`

GetTxDropPps returns the TxDropPps field if non-nil, zero value otherwise.

### GetTxDropPpsOk

`func (o *NetworkInterfaceStat) GetTxDropPpsOk() (*float64, bool)`

GetTxDropPpsOk returns a tuple with the TxDropPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDropPps

`func (o *NetworkInterfaceStat) SetTxDropPps(v float64)`

SetTxDropPps sets TxDropPps field to given value.

### HasTxDropPps

`func (o *NetworkInterfaceStat) HasTxDropPps() bool`

HasTxDropPps returns a boolean if a field has been set.

### GetTxErrorPps

`func (o *NetworkInterfaceStat) GetTxErrorPps() float64`

GetTxErrorPps returns the TxErrorPps field if non-nil, zero value otherwise.

### GetTxErrorPpsOk

`func (o *NetworkInterfaceStat) GetTxErrorPpsOk() (*float64, bool)`

GetTxErrorPpsOk returns a tuple with the TxErrorPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrorPps

`func (o *NetworkInterfaceStat) SetTxErrorPps(v float64)`

SetTxErrorPps sets TxErrorPps field to given value.

### HasTxErrorPps

`func (o *NetworkInterfaceStat) HasTxErrorPps() bool`

HasTxErrorPps returns a boolean if a field has been set.

### GetTxPps

`func (o *NetworkInterfaceStat) GetTxPps() float64`

GetTxPps returns the TxPps field if non-nil, zero value otherwise.

### GetTxPpsOk

`func (o *NetworkInterfaceStat) GetTxPpsOk() (*float64, bool)`

GetTxPpsOk returns a tuple with the TxPps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPps

`func (o *NetworkInterfaceStat) SetTxPps(v float64)`

SetTxPps sets TxPps field to given value.

### HasTxPps

`func (o *NetworkInterfaceStat) HasTxPps() bool`

HasTxPps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


