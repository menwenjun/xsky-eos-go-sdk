# OSSearchEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AllocatedSize** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**GatewayDataSize** | Pointer to **int64** |  | [optional] 
**GatewayFlavor** | Pointer to [**VMFlavorNestview**](VMFlavorNestview.md) |  | [optional] 
**GatewayHttpPort** | Pointer to **int64** |  | [optional] 
**GatewayTransportPort** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**SearchS3LoadBalancerGroup** | Pointer to [**S3LoadBalancerGroupNestview**](S3LoadBalancerGroupNestview.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SyncS3LoadBalancerGroup** | Pointer to [**S3LoadBalancerGroupNestview**](S3LoadBalancerGroupNestview.md) |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOSSearchEngine

`func NewOSSearchEngine() *OSSearchEngine`

NewOSSearchEngine instantiates a new OSSearchEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSSearchEngineWithDefaults

`func NewOSSearchEngineWithDefaults() *OSSearchEngine`

NewOSSearchEngineWithDefaults instantiates a new OSSearchEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *OSSearchEngine) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *OSSearchEngine) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *OSSearchEngine) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *OSSearchEngine) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *OSSearchEngine) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *OSSearchEngine) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *OSSearchEngine) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *OSSearchEngine) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetCluster

`func (o *OSSearchEngine) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSSearchEngine) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSSearchEngine) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSSearchEngine) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSSearchEngine) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSSearchEngine) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSSearchEngine) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSSearchEngine) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetGatewayDataSize

`func (o *OSSearchEngine) GetGatewayDataSize() int64`

GetGatewayDataSize returns the GatewayDataSize field if non-nil, zero value otherwise.

### GetGatewayDataSizeOk

`func (o *OSSearchEngine) GetGatewayDataSizeOk() (*int64, bool)`

GetGatewayDataSizeOk returns a tuple with the GatewayDataSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayDataSize

`func (o *OSSearchEngine) SetGatewayDataSize(v int64)`

SetGatewayDataSize sets GatewayDataSize field to given value.

### HasGatewayDataSize

`func (o *OSSearchEngine) HasGatewayDataSize() bool`

HasGatewayDataSize returns a boolean if a field has been set.

### GetGatewayFlavor

`func (o *OSSearchEngine) GetGatewayFlavor() VMFlavorNestview`

GetGatewayFlavor returns the GatewayFlavor field if non-nil, zero value otherwise.

### GetGatewayFlavorOk

`func (o *OSSearchEngine) GetGatewayFlavorOk() (*VMFlavorNestview, bool)`

GetGatewayFlavorOk returns a tuple with the GatewayFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayFlavor

`func (o *OSSearchEngine) SetGatewayFlavor(v VMFlavorNestview)`

SetGatewayFlavor sets GatewayFlavor field to given value.

### HasGatewayFlavor

`func (o *OSSearchEngine) HasGatewayFlavor() bool`

HasGatewayFlavor returns a boolean if a field has been set.

### GetGatewayHttpPort

`func (o *OSSearchEngine) GetGatewayHttpPort() int64`

GetGatewayHttpPort returns the GatewayHttpPort field if non-nil, zero value otherwise.

### GetGatewayHttpPortOk

`func (o *OSSearchEngine) GetGatewayHttpPortOk() (*int64, bool)`

GetGatewayHttpPortOk returns a tuple with the GatewayHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayHttpPort

`func (o *OSSearchEngine) SetGatewayHttpPort(v int64)`

SetGatewayHttpPort sets GatewayHttpPort field to given value.

### HasGatewayHttpPort

`func (o *OSSearchEngine) HasGatewayHttpPort() bool`

HasGatewayHttpPort returns a boolean if a field has been set.

### GetGatewayTransportPort

`func (o *OSSearchEngine) GetGatewayTransportPort() int64`

GetGatewayTransportPort returns the GatewayTransportPort field if non-nil, zero value otherwise.

### GetGatewayTransportPortOk

`func (o *OSSearchEngine) GetGatewayTransportPortOk() (*int64, bool)`

GetGatewayTransportPortOk returns a tuple with the GatewayTransportPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransportPort

`func (o *OSSearchEngine) SetGatewayTransportPort(v int64)`

SetGatewayTransportPort sets GatewayTransportPort field to given value.

### HasGatewayTransportPort

`func (o *OSSearchEngine) HasGatewayTransportPort() bool`

HasGatewayTransportPort returns a boolean if a field has been set.

### GetId

`func (o *OSSearchEngine) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSSearchEngine) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSSearchEngine) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSSearchEngine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSearchS3LoadBalancerGroup

`func (o *OSSearchEngine) GetSearchS3LoadBalancerGroup() S3LoadBalancerGroupNestview`

GetSearchS3LoadBalancerGroup returns the SearchS3LoadBalancerGroup field if non-nil, zero value otherwise.

### GetSearchS3LoadBalancerGroupOk

`func (o *OSSearchEngine) GetSearchS3LoadBalancerGroupOk() (*S3LoadBalancerGroupNestview, bool)`

GetSearchS3LoadBalancerGroupOk returns a tuple with the SearchS3LoadBalancerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchS3LoadBalancerGroup

`func (o *OSSearchEngine) SetSearchS3LoadBalancerGroup(v S3LoadBalancerGroupNestview)`

SetSearchS3LoadBalancerGroup sets SearchS3LoadBalancerGroup field to given value.

### HasSearchS3LoadBalancerGroup

`func (o *OSSearchEngine) HasSearchS3LoadBalancerGroup() bool`

HasSearchS3LoadBalancerGroup returns a boolean if a field has been set.

### GetSize

`func (o *OSSearchEngine) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OSSearchEngine) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OSSearchEngine) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *OSSearchEngine) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *OSSearchEngine) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSSearchEngine) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSSearchEngine) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSSearchEngine) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyncS3LoadBalancerGroup

`func (o *OSSearchEngine) GetSyncS3LoadBalancerGroup() S3LoadBalancerGroupNestview`

GetSyncS3LoadBalancerGroup returns the SyncS3LoadBalancerGroup field if non-nil, zero value otherwise.

### GetSyncS3LoadBalancerGroupOk

`func (o *OSSearchEngine) GetSyncS3LoadBalancerGroupOk() (*S3LoadBalancerGroupNestview, bool)`

GetSyncS3LoadBalancerGroupOk returns a tuple with the SyncS3LoadBalancerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncS3LoadBalancerGroup

`func (o *OSSearchEngine) SetSyncS3LoadBalancerGroup(v S3LoadBalancerGroupNestview)`

SetSyncS3LoadBalancerGroup sets SyncS3LoadBalancerGroup field to given value.

### HasSyncS3LoadBalancerGroup

`func (o *OSSearchEngine) HasSyncS3LoadBalancerGroup() bool`

HasSyncS3LoadBalancerGroup returns a boolean if a field has been set.

### GetUpdate

`func (o *OSSearchEngine) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSSearchEngine) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSSearchEngine) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSSearchEngine) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


