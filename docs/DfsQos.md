# DfsQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsPath** | Pointer to [**DfsPath**](DfsPath.md) |  | [optional] 
**DfsQosPolicy** | Pointer to [**DfsQosPolicy**](DfsQosPolicy.md) |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfs**](DfsRootfs.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IoStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsQos

`func NewDfsQos() *DfsQos`

NewDfsQos instantiates a new DfsQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQosWithDefaults

`func NewDfsQosWithDefaults() *DfsQos`

NewDfsQosWithDefaults instantiates a new DfsQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsQos) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsQos) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsQos) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsQos) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsQos) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsQos) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsQos) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsQos) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsPath

`func (o *DfsQos) GetDfsPath() DfsPath`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DfsQos) GetDfsPathOk() (*DfsPath, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DfsQos) SetDfsPath(v DfsPath)`

SetDfsPath sets DfsPath field to given value.

### HasDfsPath

`func (o *DfsQos) HasDfsPath() bool`

HasDfsPath returns a boolean if a field has been set.

### GetDfsQosPolicy

`func (o *DfsQos) GetDfsQosPolicy() DfsQosPolicy`

GetDfsQosPolicy returns the DfsQosPolicy field if non-nil, zero value otherwise.

### GetDfsQosPolicyOk

`func (o *DfsQos) GetDfsQosPolicyOk() (*DfsQosPolicy, bool)`

GetDfsQosPolicyOk returns a tuple with the DfsQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsQosPolicy

`func (o *DfsQos) SetDfsQosPolicy(v DfsQosPolicy)`

SetDfsQosPolicy sets DfsQosPolicy field to given value.

### HasDfsQosPolicy

`func (o *DfsQos) HasDfsQosPolicy() bool`

HasDfsQosPolicy returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsQos) GetDfsRootfs() DfsRootfs`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsQos) GetDfsRootfsOk() (*DfsRootfs, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsQos) SetDfsRootfs(v DfsRootfs)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsQos) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetId

`func (o *DfsQos) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsQos) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsQos) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsQos) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIoStatus

`func (o *DfsQos) GetIoStatus() string`

GetIoStatus returns the IoStatus field if non-nil, zero value otherwise.

### GetIoStatusOk

`func (o *DfsQos) GetIoStatusOk() (*string, bool)`

GetIoStatusOk returns a tuple with the IoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoStatus

`func (o *DfsQos) SetIoStatus(v string)`

SetIoStatus sets IoStatus field to given value.

### HasIoStatus

`func (o *DfsQos) HasIoStatus() bool`

HasIoStatus returns a boolean if a field has been set.

### GetStatus

`func (o *DfsQos) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsQos) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsQos) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsQos) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsQos) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsQos) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsQos) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsQos) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


