# NetworkDiagnosisStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | Pointer to **int64** |  | [optional] 
**CollPkgs** | Pointer to **int64** |  | [optional] 
**ErrPkgRatio** | Pointer to **float64** |  | [optional] 
**LostPkgRatio** | Pointer to **float64** |  | [optional] 
**RecvBytes** | Pointer to **int64** |  | [optional] 
**RecvPkgs** | Pointer to **int64** |  | [optional] 
**SendBytes** | Pointer to **int64** |  | [optional] 
**SendPkgs** | Pointer to **int64** |  | [optional] 
**TcpConnErrs** | Pointer to **int64** |  | [optional] 
**TcpSegmentRetrans** | Pointer to **int64** |  | [optional] 

## Methods

### NewNetworkDiagnosisStat

`func NewNetworkDiagnosisStat() *NetworkDiagnosisStat`

NewNetworkDiagnosisStat instantiates a new NetworkDiagnosisStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDiagnosisStatWithDefaults

`func NewNetworkDiagnosisStatWithDefaults() *NetworkDiagnosisStat`

NewNetworkDiagnosisStatWithDefaults instantiates a new NetworkDiagnosisStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *NetworkDiagnosisStat) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *NetworkDiagnosisStat) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *NetworkDiagnosisStat) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *NetworkDiagnosisStat) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetCollPkgs

`func (o *NetworkDiagnosisStat) GetCollPkgs() int64`

GetCollPkgs returns the CollPkgs field if non-nil, zero value otherwise.

### GetCollPkgsOk

`func (o *NetworkDiagnosisStat) GetCollPkgsOk() (*int64, bool)`

GetCollPkgsOk returns a tuple with the CollPkgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollPkgs

`func (o *NetworkDiagnosisStat) SetCollPkgs(v int64)`

SetCollPkgs sets CollPkgs field to given value.

### HasCollPkgs

`func (o *NetworkDiagnosisStat) HasCollPkgs() bool`

HasCollPkgs returns a boolean if a field has been set.

### GetErrPkgRatio

`func (o *NetworkDiagnosisStat) GetErrPkgRatio() float64`

GetErrPkgRatio returns the ErrPkgRatio field if non-nil, zero value otherwise.

### GetErrPkgRatioOk

`func (o *NetworkDiagnosisStat) GetErrPkgRatioOk() (*float64, bool)`

GetErrPkgRatioOk returns a tuple with the ErrPkgRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrPkgRatio

`func (o *NetworkDiagnosisStat) SetErrPkgRatio(v float64)`

SetErrPkgRatio sets ErrPkgRatio field to given value.

### HasErrPkgRatio

`func (o *NetworkDiagnosisStat) HasErrPkgRatio() bool`

HasErrPkgRatio returns a boolean if a field has been set.

### GetLostPkgRatio

`func (o *NetworkDiagnosisStat) GetLostPkgRatio() float64`

GetLostPkgRatio returns the LostPkgRatio field if non-nil, zero value otherwise.

### GetLostPkgRatioOk

`func (o *NetworkDiagnosisStat) GetLostPkgRatioOk() (*float64, bool)`

GetLostPkgRatioOk returns a tuple with the LostPkgRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostPkgRatio

`func (o *NetworkDiagnosisStat) SetLostPkgRatio(v float64)`

SetLostPkgRatio sets LostPkgRatio field to given value.

### HasLostPkgRatio

`func (o *NetworkDiagnosisStat) HasLostPkgRatio() bool`

HasLostPkgRatio returns a boolean if a field has been set.

### GetRecvBytes

`func (o *NetworkDiagnosisStat) GetRecvBytes() int64`

GetRecvBytes returns the RecvBytes field if non-nil, zero value otherwise.

### GetRecvBytesOk

`func (o *NetworkDiagnosisStat) GetRecvBytesOk() (*int64, bool)`

GetRecvBytesOk returns a tuple with the RecvBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvBytes

`func (o *NetworkDiagnosisStat) SetRecvBytes(v int64)`

SetRecvBytes sets RecvBytes field to given value.

### HasRecvBytes

`func (o *NetworkDiagnosisStat) HasRecvBytes() bool`

HasRecvBytes returns a boolean if a field has been set.

### GetRecvPkgs

`func (o *NetworkDiagnosisStat) GetRecvPkgs() int64`

GetRecvPkgs returns the RecvPkgs field if non-nil, zero value otherwise.

### GetRecvPkgsOk

`func (o *NetworkDiagnosisStat) GetRecvPkgsOk() (*int64, bool)`

GetRecvPkgsOk returns a tuple with the RecvPkgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvPkgs

`func (o *NetworkDiagnosisStat) SetRecvPkgs(v int64)`

SetRecvPkgs sets RecvPkgs field to given value.

### HasRecvPkgs

`func (o *NetworkDiagnosisStat) HasRecvPkgs() bool`

HasRecvPkgs returns a boolean if a field has been set.

### GetSendBytes

`func (o *NetworkDiagnosisStat) GetSendBytes() int64`

GetSendBytes returns the SendBytes field if non-nil, zero value otherwise.

### GetSendBytesOk

`func (o *NetworkDiagnosisStat) GetSendBytesOk() (*int64, bool)`

GetSendBytesOk returns a tuple with the SendBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBytes

`func (o *NetworkDiagnosisStat) SetSendBytes(v int64)`

SetSendBytes sets SendBytes field to given value.

### HasSendBytes

`func (o *NetworkDiagnosisStat) HasSendBytes() bool`

HasSendBytes returns a boolean if a field has been set.

### GetSendPkgs

`func (o *NetworkDiagnosisStat) GetSendPkgs() int64`

GetSendPkgs returns the SendPkgs field if non-nil, zero value otherwise.

### GetSendPkgsOk

`func (o *NetworkDiagnosisStat) GetSendPkgsOk() (*int64, bool)`

GetSendPkgsOk returns a tuple with the SendPkgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPkgs

`func (o *NetworkDiagnosisStat) SetSendPkgs(v int64)`

SetSendPkgs sets SendPkgs field to given value.

### HasSendPkgs

`func (o *NetworkDiagnosisStat) HasSendPkgs() bool`

HasSendPkgs returns a boolean if a field has been set.

### GetTcpConnErrs

`func (o *NetworkDiagnosisStat) GetTcpConnErrs() int64`

GetTcpConnErrs returns the TcpConnErrs field if non-nil, zero value otherwise.

### GetTcpConnErrsOk

`func (o *NetworkDiagnosisStat) GetTcpConnErrsOk() (*int64, bool)`

GetTcpConnErrsOk returns a tuple with the TcpConnErrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpConnErrs

`func (o *NetworkDiagnosisStat) SetTcpConnErrs(v int64)`

SetTcpConnErrs sets TcpConnErrs field to given value.

### HasTcpConnErrs

`func (o *NetworkDiagnosisStat) HasTcpConnErrs() bool`

HasTcpConnErrs returns a boolean if a field has been set.

### GetTcpSegmentRetrans

`func (o *NetworkDiagnosisStat) GetTcpSegmentRetrans() int64`

GetTcpSegmentRetrans returns the TcpSegmentRetrans field if non-nil, zero value otherwise.

### GetTcpSegmentRetransOk

`func (o *NetworkDiagnosisStat) GetTcpSegmentRetransOk() (*int64, bool)`

GetTcpSegmentRetransOk returns a tuple with the TcpSegmentRetrans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSegmentRetrans

`func (o *NetworkDiagnosisStat) SetTcpSegmentRetrans(v int64)`

SetTcpSegmentRetrans sets TcpSegmentRetrans field to given value.

### HasTcpSegmentRetrans

`func (o *NetworkDiagnosisStat) HasTcpSegmentRetrans() bool`

HasTcpSegmentRetrans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


