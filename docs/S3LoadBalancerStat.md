# S3LoadBalancerStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveConnections** | Pointer to **float64** |  | [optional] 
**BackendDownBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**BackendDownloadIops** | Pointer to **float64** |  | [optional] 
**BackendUpBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**BackendUploadIops** | Pointer to **float64** |  | [optional] 
**CpuUtil** | Pointer to **float64** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DownBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**DownloadIops** | Pointer to **float64** |  | [optional] 
**FailureRequests** | Pointer to **float64** |  | [optional] 
**MemUsagePercent** | Pointer to **float64** |  | [optional] 
**SuccessRequests** | Pointer to **float64** |  | [optional] 
**UpBandwidthKbyte** | Pointer to **float64** |  | [optional] 
**UploadIops** | Pointer to **float64** |  | [optional] 
**VipsConnStat** | Pointer to [**[]VipConnectionStat**](VipConnectionStat.md) |  | [optional] 

## Methods

### NewS3LoadBalancerStat

`func NewS3LoadBalancerStat() *S3LoadBalancerStat`

NewS3LoadBalancerStat instantiates a new S3LoadBalancerStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancerStatWithDefaults

`func NewS3LoadBalancerStatWithDefaults() *S3LoadBalancerStat`

NewS3LoadBalancerStatWithDefaults instantiates a new S3LoadBalancerStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveConnections

`func (o *S3LoadBalancerStat) GetActiveConnections() float64`

GetActiveConnections returns the ActiveConnections field if non-nil, zero value otherwise.

### GetActiveConnectionsOk

`func (o *S3LoadBalancerStat) GetActiveConnectionsOk() (*float64, bool)`

GetActiveConnectionsOk returns a tuple with the ActiveConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveConnections

`func (o *S3LoadBalancerStat) SetActiveConnections(v float64)`

SetActiveConnections sets ActiveConnections field to given value.

### HasActiveConnections

`func (o *S3LoadBalancerStat) HasActiveConnections() bool`

HasActiveConnections returns a boolean if a field has been set.

### GetBackendDownBandwidthKbyte

`func (o *S3LoadBalancerStat) GetBackendDownBandwidthKbyte() float64`

GetBackendDownBandwidthKbyte returns the BackendDownBandwidthKbyte field if non-nil, zero value otherwise.

### GetBackendDownBandwidthKbyteOk

`func (o *S3LoadBalancerStat) GetBackendDownBandwidthKbyteOk() (*float64, bool)`

GetBackendDownBandwidthKbyteOk returns a tuple with the BackendDownBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendDownBandwidthKbyte

`func (o *S3LoadBalancerStat) SetBackendDownBandwidthKbyte(v float64)`

SetBackendDownBandwidthKbyte sets BackendDownBandwidthKbyte field to given value.

### HasBackendDownBandwidthKbyte

`func (o *S3LoadBalancerStat) HasBackendDownBandwidthKbyte() bool`

HasBackendDownBandwidthKbyte returns a boolean if a field has been set.

### GetBackendDownloadIops

`func (o *S3LoadBalancerStat) GetBackendDownloadIops() float64`

GetBackendDownloadIops returns the BackendDownloadIops field if non-nil, zero value otherwise.

### GetBackendDownloadIopsOk

`func (o *S3LoadBalancerStat) GetBackendDownloadIopsOk() (*float64, bool)`

GetBackendDownloadIopsOk returns a tuple with the BackendDownloadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendDownloadIops

`func (o *S3LoadBalancerStat) SetBackendDownloadIops(v float64)`

SetBackendDownloadIops sets BackendDownloadIops field to given value.

### HasBackendDownloadIops

`func (o *S3LoadBalancerStat) HasBackendDownloadIops() bool`

HasBackendDownloadIops returns a boolean if a field has been set.

### GetBackendUpBandwidthKbyte

`func (o *S3LoadBalancerStat) GetBackendUpBandwidthKbyte() float64`

GetBackendUpBandwidthKbyte returns the BackendUpBandwidthKbyte field if non-nil, zero value otherwise.

### GetBackendUpBandwidthKbyteOk

`func (o *S3LoadBalancerStat) GetBackendUpBandwidthKbyteOk() (*float64, bool)`

GetBackendUpBandwidthKbyteOk returns a tuple with the BackendUpBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendUpBandwidthKbyte

`func (o *S3LoadBalancerStat) SetBackendUpBandwidthKbyte(v float64)`

SetBackendUpBandwidthKbyte sets BackendUpBandwidthKbyte field to given value.

### HasBackendUpBandwidthKbyte

`func (o *S3LoadBalancerStat) HasBackendUpBandwidthKbyte() bool`

HasBackendUpBandwidthKbyte returns a boolean if a field has been set.

### GetBackendUploadIops

`func (o *S3LoadBalancerStat) GetBackendUploadIops() float64`

GetBackendUploadIops returns the BackendUploadIops field if non-nil, zero value otherwise.

### GetBackendUploadIopsOk

`func (o *S3LoadBalancerStat) GetBackendUploadIopsOk() (*float64, bool)`

GetBackendUploadIopsOk returns a tuple with the BackendUploadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendUploadIops

`func (o *S3LoadBalancerStat) SetBackendUploadIops(v float64)`

SetBackendUploadIops sets BackendUploadIops field to given value.

### HasBackendUploadIops

`func (o *S3LoadBalancerStat) HasBackendUploadIops() bool`

HasBackendUploadIops returns a boolean if a field has been set.

### GetCpuUtil

`func (o *S3LoadBalancerStat) GetCpuUtil() float64`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *S3LoadBalancerStat) GetCpuUtilOk() (*float64, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *S3LoadBalancerStat) SetCpuUtil(v float64)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *S3LoadBalancerStat) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCreate

`func (o *S3LoadBalancerStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *S3LoadBalancerStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *S3LoadBalancerStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *S3LoadBalancerStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDownBandwidthKbyte

`func (o *S3LoadBalancerStat) GetDownBandwidthKbyte() float64`

GetDownBandwidthKbyte returns the DownBandwidthKbyte field if non-nil, zero value otherwise.

### GetDownBandwidthKbyteOk

`func (o *S3LoadBalancerStat) GetDownBandwidthKbyteOk() (*float64, bool)`

GetDownBandwidthKbyteOk returns a tuple with the DownBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownBandwidthKbyte

`func (o *S3LoadBalancerStat) SetDownBandwidthKbyte(v float64)`

SetDownBandwidthKbyte sets DownBandwidthKbyte field to given value.

### HasDownBandwidthKbyte

`func (o *S3LoadBalancerStat) HasDownBandwidthKbyte() bool`

HasDownBandwidthKbyte returns a boolean if a field has been set.

### GetDownloadIops

`func (o *S3LoadBalancerStat) GetDownloadIops() float64`

GetDownloadIops returns the DownloadIops field if non-nil, zero value otherwise.

### GetDownloadIopsOk

`func (o *S3LoadBalancerStat) GetDownloadIopsOk() (*float64, bool)`

GetDownloadIopsOk returns a tuple with the DownloadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadIops

`func (o *S3LoadBalancerStat) SetDownloadIops(v float64)`

SetDownloadIops sets DownloadIops field to given value.

### HasDownloadIops

`func (o *S3LoadBalancerStat) HasDownloadIops() bool`

HasDownloadIops returns a boolean if a field has been set.

### GetFailureRequests

`func (o *S3LoadBalancerStat) GetFailureRequests() float64`

GetFailureRequests returns the FailureRequests field if non-nil, zero value otherwise.

### GetFailureRequestsOk

`func (o *S3LoadBalancerStat) GetFailureRequestsOk() (*float64, bool)`

GetFailureRequestsOk returns a tuple with the FailureRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureRequests

`func (o *S3LoadBalancerStat) SetFailureRequests(v float64)`

SetFailureRequests sets FailureRequests field to given value.

### HasFailureRequests

`func (o *S3LoadBalancerStat) HasFailureRequests() bool`

HasFailureRequests returns a boolean if a field has been set.

### GetMemUsagePercent

`func (o *S3LoadBalancerStat) GetMemUsagePercent() float64`

GetMemUsagePercent returns the MemUsagePercent field if non-nil, zero value otherwise.

### GetMemUsagePercentOk

`func (o *S3LoadBalancerStat) GetMemUsagePercentOk() (*float64, bool)`

GetMemUsagePercentOk returns a tuple with the MemUsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsagePercent

`func (o *S3LoadBalancerStat) SetMemUsagePercent(v float64)`

SetMemUsagePercent sets MemUsagePercent field to given value.

### HasMemUsagePercent

`func (o *S3LoadBalancerStat) HasMemUsagePercent() bool`

HasMemUsagePercent returns a boolean if a field has been set.

### GetSuccessRequests

`func (o *S3LoadBalancerStat) GetSuccessRequests() float64`

GetSuccessRequests returns the SuccessRequests field if non-nil, zero value otherwise.

### GetSuccessRequestsOk

`func (o *S3LoadBalancerStat) GetSuccessRequestsOk() (*float64, bool)`

GetSuccessRequestsOk returns a tuple with the SuccessRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRequests

`func (o *S3LoadBalancerStat) SetSuccessRequests(v float64)`

SetSuccessRequests sets SuccessRequests field to given value.

### HasSuccessRequests

`func (o *S3LoadBalancerStat) HasSuccessRequests() bool`

HasSuccessRequests returns a boolean if a field has been set.

### GetUpBandwidthKbyte

`func (o *S3LoadBalancerStat) GetUpBandwidthKbyte() float64`

GetUpBandwidthKbyte returns the UpBandwidthKbyte field if non-nil, zero value otherwise.

### GetUpBandwidthKbyteOk

`func (o *S3LoadBalancerStat) GetUpBandwidthKbyteOk() (*float64, bool)`

GetUpBandwidthKbyteOk returns a tuple with the UpBandwidthKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpBandwidthKbyte

`func (o *S3LoadBalancerStat) SetUpBandwidthKbyte(v float64)`

SetUpBandwidthKbyte sets UpBandwidthKbyte field to given value.

### HasUpBandwidthKbyte

`func (o *S3LoadBalancerStat) HasUpBandwidthKbyte() bool`

HasUpBandwidthKbyte returns a boolean if a field has been set.

### GetUploadIops

`func (o *S3LoadBalancerStat) GetUploadIops() float64`

GetUploadIops returns the UploadIops field if non-nil, zero value otherwise.

### GetUploadIopsOk

`func (o *S3LoadBalancerStat) GetUploadIopsOk() (*float64, bool)`

GetUploadIopsOk returns a tuple with the UploadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadIops

`func (o *S3LoadBalancerStat) SetUploadIops(v float64)`

SetUploadIops sets UploadIops field to given value.

### HasUploadIops

`func (o *S3LoadBalancerStat) HasUploadIops() bool`

HasUploadIops returns a boolean if a field has been set.

### GetVipsConnStat

`func (o *S3LoadBalancerStat) GetVipsConnStat() []VipConnectionStat`

GetVipsConnStat returns the VipsConnStat field if non-nil, zero value otherwise.

### GetVipsConnStatOk

`func (o *S3LoadBalancerStat) GetVipsConnStatOk() (*[]VipConnectionStat, bool)`

GetVipsConnStatOk returns a tuple with the VipsConnStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipsConnStat

`func (o *S3LoadBalancerStat) SetVipsConnStat(v []VipConnectionStat)`

SetVipsConnStat sets VipsConnStat field to given value.

### HasVipsConnStat

`func (o *S3LoadBalancerStat) HasVipsConnStat() bool`

HasVipsConnStat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


