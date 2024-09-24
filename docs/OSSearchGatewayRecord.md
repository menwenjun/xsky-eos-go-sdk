# OSSearchGatewayRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Pool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**SearchEngine** | Pointer to [**OSSearchEngineNestview**](OSSearchEngineNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Samples** | Pointer to [**[]OSSearchGatewayStat**](OSSearchGatewayStat.md) |  | [optional] 

## Methods

### NewOSSearchGatewayRecord

`func NewOSSearchGatewayRecord() *OSSearchGatewayRecord`

NewOSSearchGatewayRecord instantiates a new OSSearchGatewayRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSSearchGatewayRecordWithDefaults

`func NewOSSearchGatewayRecordWithDefaults() *OSSearchGatewayRecord`

NewOSSearchGatewayRecordWithDefaults instantiates a new OSSearchGatewayRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *OSSearchGatewayRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *OSSearchGatewayRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *OSSearchGatewayRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *OSSearchGatewayRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *OSSearchGatewayRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSSearchGatewayRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSSearchGatewayRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSSearchGatewayRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSSearchGatewayRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSSearchGatewayRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSSearchGatewayRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSSearchGatewayRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetHost

`func (o *OSSearchGatewayRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OSSearchGatewayRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OSSearchGatewayRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *OSSearchGatewayRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *OSSearchGatewayRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSSearchGatewayRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSSearchGatewayRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSSearchGatewayRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPool

`func (o *OSSearchGatewayRecord) GetPool() PoolNestview`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *OSSearchGatewayRecord) GetPoolOk() (*PoolNestview, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *OSSearchGatewayRecord) SetPool(v PoolNestview)`

SetPool sets Pool field to given value.

### HasPool

`func (o *OSSearchGatewayRecord) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetSearchEngine

`func (o *OSSearchGatewayRecord) GetSearchEngine() OSSearchEngineNestview`

GetSearchEngine returns the SearchEngine field if non-nil, zero value otherwise.

### GetSearchEngineOk

`func (o *OSSearchGatewayRecord) GetSearchEngineOk() (*OSSearchEngineNestview, bool)`

GetSearchEngineOk returns a tuple with the SearchEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEngine

`func (o *OSSearchGatewayRecord) SetSearchEngine(v OSSearchEngineNestview)`

SetSearchEngine sets SearchEngine field to given value.

### HasSearchEngine

`func (o *OSSearchGatewayRecord) HasSearchEngine() bool`

HasSearchEngine returns a boolean if a field has been set.

### GetStatus

`func (o *OSSearchGatewayRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSSearchGatewayRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSSearchGatewayRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSSearchGatewayRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *OSSearchGatewayRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSSearchGatewayRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSSearchGatewayRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSSearchGatewayRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *OSSearchGatewayRecord) GetSamples() []OSSearchGatewayStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *OSSearchGatewayRecord) GetSamplesOk() (*[]OSSearchGatewayStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *OSSearchGatewayRecord) SetSamples(v []OSSearchGatewayStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *OSSearchGatewayRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


