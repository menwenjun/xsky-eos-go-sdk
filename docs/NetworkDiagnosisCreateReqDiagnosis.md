# NetworkDiagnosisCreateReqDiagnosis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostIds** | **[]int64** | ids of host | 
**HostScope** | **string** | host scope | 
**Networks** | **[]string** | network types | 

## Methods

### NewNetworkDiagnosisCreateReqDiagnosis

`func NewNetworkDiagnosisCreateReqDiagnosis(hostIds []int64, hostScope string, networks []string, ) *NetworkDiagnosisCreateReqDiagnosis`

NewNetworkDiagnosisCreateReqDiagnosis instantiates a new NetworkDiagnosisCreateReqDiagnosis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDiagnosisCreateReqDiagnosisWithDefaults

`func NewNetworkDiagnosisCreateReqDiagnosisWithDefaults() *NetworkDiagnosisCreateReqDiagnosis`

NewNetworkDiagnosisCreateReqDiagnosisWithDefaults instantiates a new NetworkDiagnosisCreateReqDiagnosis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostIds

`func (o *NetworkDiagnosisCreateReqDiagnosis) GetHostIds() []int64`

GetHostIds returns the HostIds field if non-nil, zero value otherwise.

### GetHostIdsOk

`func (o *NetworkDiagnosisCreateReqDiagnosis) GetHostIdsOk() (*[]int64, bool)`

GetHostIdsOk returns a tuple with the HostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIds

`func (o *NetworkDiagnosisCreateReqDiagnosis) SetHostIds(v []int64)`

SetHostIds sets HostIds field to given value.


### GetHostScope

`func (o *NetworkDiagnosisCreateReqDiagnosis) GetHostScope() string`

GetHostScope returns the HostScope field if non-nil, zero value otherwise.

### GetHostScopeOk

`func (o *NetworkDiagnosisCreateReqDiagnosis) GetHostScopeOk() (*string, bool)`

GetHostScopeOk returns a tuple with the HostScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostScope

`func (o *NetworkDiagnosisCreateReqDiagnosis) SetHostScope(v string)`

SetHostScope sets HostScope field to given value.


### GetNetworks

`func (o *NetworkDiagnosisCreateReqDiagnosis) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *NetworkDiagnosisCreateReqDiagnosis) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *NetworkDiagnosisCreateReqDiagnosis) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


