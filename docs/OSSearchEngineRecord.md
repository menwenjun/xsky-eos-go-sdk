# OSSearchEngineRecord

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
**Samples** | Pointer to [**[]OSSearchEngineStat**](OSSearchEngineStat.md) |  | [optional] 

## Methods

### NewOSSearchEngineRecord

`func NewOSSearchEngineRecord() *OSSearchEngineRecord`

NewOSSearchEngineRecord instantiates a new OSSearchEngineRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSSearchEngineRecordWithDefaults

`func NewOSSearchEngineRecordWithDefaults() *OSSearchEngineRecord`

NewOSSearchEngineRecordWithDefaults instantiates a new OSSearchEngineRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *OSSearchEngineRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *OSSearchEngineRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *OSSearchEngineRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *OSSearchEngineRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *OSSearchEngineRecord) GetAllocatedSize() int64`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *OSSearchEngineRecord) GetAllocatedSizeOk() (*int64, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *OSSearchEngineRecord) SetAllocatedSize(v int64)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *OSSearchEngineRecord) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetCluster

`func (o *OSSearchEngineRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSSearchEngineRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSSearchEngineRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSSearchEngineRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSSearchEngineRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSSearchEngineRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSSearchEngineRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSSearchEngineRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetGatewayDataSize

`func (o *OSSearchEngineRecord) GetGatewayDataSize() int64`

GetGatewayDataSize returns the GatewayDataSize field if non-nil, zero value otherwise.

### GetGatewayDataSizeOk

`func (o *OSSearchEngineRecord) GetGatewayDataSizeOk() (*int64, bool)`

GetGatewayDataSizeOk returns a tuple with the GatewayDataSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayDataSize

`func (o *OSSearchEngineRecord) SetGatewayDataSize(v int64)`

SetGatewayDataSize sets GatewayDataSize field to given value.

### HasGatewayDataSize

`func (o *OSSearchEngineRecord) HasGatewayDataSize() bool`

HasGatewayDataSize returns a boolean if a field has been set.

### GetGatewayFlavor

`func (o *OSSearchEngineRecord) GetGatewayFlavor() VMFlavorNestview`

GetGatewayFlavor returns the GatewayFlavor field if non-nil, zero value otherwise.

### GetGatewayFlavorOk

`func (o *OSSearchEngineRecord) GetGatewayFlavorOk() (*VMFlavorNestview, bool)`

GetGatewayFlavorOk returns a tuple with the GatewayFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayFlavor

`func (o *OSSearchEngineRecord) SetGatewayFlavor(v VMFlavorNestview)`

SetGatewayFlavor sets GatewayFlavor field to given value.

### HasGatewayFlavor

`func (o *OSSearchEngineRecord) HasGatewayFlavor() bool`

HasGatewayFlavor returns a boolean if a field has been set.

### GetGatewayHttpPort

`func (o *OSSearchEngineRecord) GetGatewayHttpPort() int64`

GetGatewayHttpPort returns the GatewayHttpPort field if non-nil, zero value otherwise.

### GetGatewayHttpPortOk

`func (o *OSSearchEngineRecord) GetGatewayHttpPortOk() (*int64, bool)`

GetGatewayHttpPortOk returns a tuple with the GatewayHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayHttpPort

`func (o *OSSearchEngineRecord) SetGatewayHttpPort(v int64)`

SetGatewayHttpPort sets GatewayHttpPort field to given value.

### HasGatewayHttpPort

`func (o *OSSearchEngineRecord) HasGatewayHttpPort() bool`

HasGatewayHttpPort returns a boolean if a field has been set.

### GetGatewayTransportPort

`func (o *OSSearchEngineRecord) GetGatewayTransportPort() int64`

GetGatewayTransportPort returns the GatewayTransportPort field if non-nil, zero value otherwise.

### GetGatewayTransportPortOk

`func (o *OSSearchEngineRecord) GetGatewayTransportPortOk() (*int64, bool)`

GetGatewayTransportPortOk returns a tuple with the GatewayTransportPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransportPort

`func (o *OSSearchEngineRecord) SetGatewayTransportPort(v int64)`

SetGatewayTransportPort sets GatewayTransportPort field to given value.

### HasGatewayTransportPort

`func (o *OSSearchEngineRecord) HasGatewayTransportPort() bool`

HasGatewayTransportPort returns a boolean if a field has been set.

### GetId

`func (o *OSSearchEngineRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSSearchEngineRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSSearchEngineRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSSearchEngineRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSearchS3LoadBalancerGroup

`func (o *OSSearchEngineRecord) GetSearchS3LoadBalancerGroup() S3LoadBalancerGroupNestview`

GetSearchS3LoadBalancerGroup returns the SearchS3LoadBalancerGroup field if non-nil, zero value otherwise.

### GetSearchS3LoadBalancerGroupOk

`func (o *OSSearchEngineRecord) GetSearchS3LoadBalancerGroupOk() (*S3LoadBalancerGroupNestview, bool)`

GetSearchS3LoadBalancerGroupOk returns a tuple with the SearchS3LoadBalancerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchS3LoadBalancerGroup

`func (o *OSSearchEngineRecord) SetSearchS3LoadBalancerGroup(v S3LoadBalancerGroupNestview)`

SetSearchS3LoadBalancerGroup sets SearchS3LoadBalancerGroup field to given value.

### HasSearchS3LoadBalancerGroup

`func (o *OSSearchEngineRecord) HasSearchS3LoadBalancerGroup() bool`

HasSearchS3LoadBalancerGroup returns a boolean if a field has been set.

### GetSize

`func (o *OSSearchEngineRecord) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OSSearchEngineRecord) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OSSearchEngineRecord) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *OSSearchEngineRecord) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *OSSearchEngineRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSSearchEngineRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSSearchEngineRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSSearchEngineRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyncS3LoadBalancerGroup

`func (o *OSSearchEngineRecord) GetSyncS3LoadBalancerGroup() S3LoadBalancerGroupNestview`

GetSyncS3LoadBalancerGroup returns the SyncS3LoadBalancerGroup field if non-nil, zero value otherwise.

### GetSyncS3LoadBalancerGroupOk

`func (o *OSSearchEngineRecord) GetSyncS3LoadBalancerGroupOk() (*S3LoadBalancerGroupNestview, bool)`

GetSyncS3LoadBalancerGroupOk returns a tuple with the SyncS3LoadBalancerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncS3LoadBalancerGroup

`func (o *OSSearchEngineRecord) SetSyncS3LoadBalancerGroup(v S3LoadBalancerGroupNestview)`

SetSyncS3LoadBalancerGroup sets SyncS3LoadBalancerGroup field to given value.

### HasSyncS3LoadBalancerGroup

`func (o *OSSearchEngineRecord) HasSyncS3LoadBalancerGroup() bool`

HasSyncS3LoadBalancerGroup returns a boolean if a field has been set.

### GetUpdate

`func (o *OSSearchEngineRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSSearchEngineRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSSearchEngineRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSSearchEngineRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *OSSearchEngineRecord) GetSamples() []OSSearchEngineStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *OSSearchEngineRecord) GetSamplesOk() (*[]OSSearchEngineStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *OSSearchEngineRecord) SetSamples(v []OSSearchEngineStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *OSSearchEngineRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


