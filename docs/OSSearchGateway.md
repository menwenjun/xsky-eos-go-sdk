# OSSearchGateway

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

## Methods

### NewOSSearchGateway

`func NewOSSearchGateway() *OSSearchGateway`

NewOSSearchGateway instantiates a new OSSearchGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSSearchGatewayWithDefaults

`func NewOSSearchGatewayWithDefaults() *OSSearchGateway`

NewOSSearchGatewayWithDefaults instantiates a new OSSearchGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *OSSearchGateway) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *OSSearchGateway) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *OSSearchGateway) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *OSSearchGateway) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *OSSearchGateway) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSSearchGateway) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSSearchGateway) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSSearchGateway) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSSearchGateway) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSSearchGateway) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSSearchGateway) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSSearchGateway) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetHost

`func (o *OSSearchGateway) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OSSearchGateway) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OSSearchGateway) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *OSSearchGateway) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *OSSearchGateway) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSSearchGateway) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSSearchGateway) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSSearchGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPool

`func (o *OSSearchGateway) GetPool() PoolNestview`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *OSSearchGateway) GetPoolOk() (*PoolNestview, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *OSSearchGateway) SetPool(v PoolNestview)`

SetPool sets Pool field to given value.

### HasPool

`func (o *OSSearchGateway) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetSearchEngine

`func (o *OSSearchGateway) GetSearchEngine() OSSearchEngineNestview`

GetSearchEngine returns the SearchEngine field if non-nil, zero value otherwise.

### GetSearchEngineOk

`func (o *OSSearchGateway) GetSearchEngineOk() (*OSSearchEngineNestview, bool)`

GetSearchEngineOk returns a tuple with the SearchEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEngine

`func (o *OSSearchGateway) SetSearchEngine(v OSSearchEngineNestview)`

SetSearchEngine sets SearchEngine field to given value.

### HasSearchEngine

`func (o *OSSearchGateway) HasSearchEngine() bool`

HasSearchEngine returns a boolean if a field has been set.

### GetStatus

`func (o *OSSearchGateway) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSSearchGateway) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSSearchGateway) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSSearchGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *OSSearchGateway) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSSearchGateway) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSSearchGateway) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSSearchGateway) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


