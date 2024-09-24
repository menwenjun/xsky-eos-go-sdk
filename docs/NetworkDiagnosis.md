# NetworkDiagnosis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DiagnoseActiveNum** | Pointer to **int64** |  | [optional] 
**DiagnoseErrorNum** | Pointer to **int64** |  | [optional] 
**DiagnoseNum** | Pointer to **int64** |  | [optional] 
**DiagnoseWarningNum** | Pointer to **int64** |  | [optional] 
**DiagnosingNum** | Pointer to **int64** |  | [optional] 
**HostNum** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Networks** | Pointer to **[]string** |  | [optional] 
**NotDiagnoseNum** | Pointer to **int64** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNetworkDiagnosis

`func NewNetworkDiagnosis() *NetworkDiagnosis`

NewNetworkDiagnosis instantiates a new NetworkDiagnosis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDiagnosisWithDefaults

`func NewNetworkDiagnosisWithDefaults() *NetworkDiagnosis`

NewNetworkDiagnosisWithDefaults instantiates a new NetworkDiagnosis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *NetworkDiagnosis) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *NetworkDiagnosis) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *NetworkDiagnosis) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *NetworkDiagnosis) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *NetworkDiagnosis) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *NetworkDiagnosis) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *NetworkDiagnosis) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *NetworkDiagnosis) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDiagnoseActiveNum

`func (o *NetworkDiagnosis) GetDiagnoseActiveNum() int64`

GetDiagnoseActiveNum returns the DiagnoseActiveNum field if non-nil, zero value otherwise.

### GetDiagnoseActiveNumOk

`func (o *NetworkDiagnosis) GetDiagnoseActiveNumOk() (*int64, bool)`

GetDiagnoseActiveNumOk returns a tuple with the DiagnoseActiveNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnoseActiveNum

`func (o *NetworkDiagnosis) SetDiagnoseActiveNum(v int64)`

SetDiagnoseActiveNum sets DiagnoseActiveNum field to given value.

### HasDiagnoseActiveNum

`func (o *NetworkDiagnosis) HasDiagnoseActiveNum() bool`

HasDiagnoseActiveNum returns a boolean if a field has been set.

### GetDiagnoseErrorNum

`func (o *NetworkDiagnosis) GetDiagnoseErrorNum() int64`

GetDiagnoseErrorNum returns the DiagnoseErrorNum field if non-nil, zero value otherwise.

### GetDiagnoseErrorNumOk

`func (o *NetworkDiagnosis) GetDiagnoseErrorNumOk() (*int64, bool)`

GetDiagnoseErrorNumOk returns a tuple with the DiagnoseErrorNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnoseErrorNum

`func (o *NetworkDiagnosis) SetDiagnoseErrorNum(v int64)`

SetDiagnoseErrorNum sets DiagnoseErrorNum field to given value.

### HasDiagnoseErrorNum

`func (o *NetworkDiagnosis) HasDiagnoseErrorNum() bool`

HasDiagnoseErrorNum returns a boolean if a field has been set.

### GetDiagnoseNum

`func (o *NetworkDiagnosis) GetDiagnoseNum() int64`

GetDiagnoseNum returns the DiagnoseNum field if non-nil, zero value otherwise.

### GetDiagnoseNumOk

`func (o *NetworkDiagnosis) GetDiagnoseNumOk() (*int64, bool)`

GetDiagnoseNumOk returns a tuple with the DiagnoseNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnoseNum

`func (o *NetworkDiagnosis) SetDiagnoseNum(v int64)`

SetDiagnoseNum sets DiagnoseNum field to given value.

### HasDiagnoseNum

`func (o *NetworkDiagnosis) HasDiagnoseNum() bool`

HasDiagnoseNum returns a boolean if a field has been set.

### GetDiagnoseWarningNum

`func (o *NetworkDiagnosis) GetDiagnoseWarningNum() int64`

GetDiagnoseWarningNum returns the DiagnoseWarningNum field if non-nil, zero value otherwise.

### GetDiagnoseWarningNumOk

`func (o *NetworkDiagnosis) GetDiagnoseWarningNumOk() (*int64, bool)`

GetDiagnoseWarningNumOk returns a tuple with the DiagnoseWarningNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnoseWarningNum

`func (o *NetworkDiagnosis) SetDiagnoseWarningNum(v int64)`

SetDiagnoseWarningNum sets DiagnoseWarningNum field to given value.

### HasDiagnoseWarningNum

`func (o *NetworkDiagnosis) HasDiagnoseWarningNum() bool`

HasDiagnoseWarningNum returns a boolean if a field has been set.

### GetDiagnosingNum

`func (o *NetworkDiagnosis) GetDiagnosingNum() int64`

GetDiagnosingNum returns the DiagnosingNum field if non-nil, zero value otherwise.

### GetDiagnosingNumOk

`func (o *NetworkDiagnosis) GetDiagnosingNumOk() (*int64, bool)`

GetDiagnosingNumOk returns a tuple with the DiagnosingNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosingNum

`func (o *NetworkDiagnosis) SetDiagnosingNum(v int64)`

SetDiagnosingNum sets DiagnosingNum field to given value.

### HasDiagnosingNum

`func (o *NetworkDiagnosis) HasDiagnosingNum() bool`

HasDiagnosingNum returns a boolean if a field has been set.

### GetHostNum

`func (o *NetworkDiagnosis) GetHostNum() int64`

GetHostNum returns the HostNum field if non-nil, zero value otherwise.

### GetHostNumOk

`func (o *NetworkDiagnosis) GetHostNumOk() (*int64, bool)`

GetHostNumOk returns a tuple with the HostNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNum

`func (o *NetworkDiagnosis) SetHostNum(v int64)`

SetHostNum sets HostNum field to given value.

### HasHostNum

`func (o *NetworkDiagnosis) HasHostNum() bool`

HasHostNum returns a boolean if a field has been set.

### GetId

`func (o *NetworkDiagnosis) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDiagnosis) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDiagnosis) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkDiagnosis) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworks

`func (o *NetworkDiagnosis) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *NetworkDiagnosis) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *NetworkDiagnosis) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *NetworkDiagnosis) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNotDiagnoseNum

`func (o *NetworkDiagnosis) GetNotDiagnoseNum() int64`

GetNotDiagnoseNum returns the NotDiagnoseNum field if non-nil, zero value otherwise.

### GetNotDiagnoseNumOk

`func (o *NetworkDiagnosis) GetNotDiagnoseNumOk() (*int64, bool)`

GetNotDiagnoseNumOk returns a tuple with the NotDiagnoseNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotDiagnoseNum

`func (o *NetworkDiagnosis) SetNotDiagnoseNum(v int64)`

SetNotDiagnoseNum sets NotDiagnoseNum field to given value.

### HasNotDiagnoseNum

`func (o *NetworkDiagnosis) HasNotDiagnoseNum() bool`

HasNotDiagnoseNum returns a boolean if a field has been set.

### GetProgress

`func (o *NetworkDiagnosis) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *NetworkDiagnosis) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *NetworkDiagnosis) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *NetworkDiagnosis) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkDiagnosis) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkDiagnosis) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkDiagnosis) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkDiagnosis) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *NetworkDiagnosis) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *NetworkDiagnosis) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *NetworkDiagnosis) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *NetworkDiagnosis) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


