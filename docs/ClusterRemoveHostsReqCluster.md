# ClusterRemoveHostsReqCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CephHostIds** | Pointer to **[]int64** | ceph host ids | [optional] 
**OspHostIds** | Pointer to **[]int64** | osp host ids | [optional] 

## Methods

### NewClusterRemoveHostsReqCluster

`func NewClusterRemoveHostsReqCluster() *ClusterRemoveHostsReqCluster`

NewClusterRemoveHostsReqCluster instantiates a new ClusterRemoveHostsReqCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRemoveHostsReqClusterWithDefaults

`func NewClusterRemoveHostsReqClusterWithDefaults() *ClusterRemoveHostsReqCluster`

NewClusterRemoveHostsReqClusterWithDefaults instantiates a new ClusterRemoveHostsReqCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCephHostIds

`func (o *ClusterRemoveHostsReqCluster) GetCephHostIds() []int64`

GetCephHostIds returns the CephHostIds field if non-nil, zero value otherwise.

### GetCephHostIdsOk

`func (o *ClusterRemoveHostsReqCluster) GetCephHostIdsOk() (*[]int64, bool)`

GetCephHostIdsOk returns a tuple with the CephHostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephHostIds

`func (o *ClusterRemoveHostsReqCluster) SetCephHostIds(v []int64)`

SetCephHostIds sets CephHostIds field to given value.

### HasCephHostIds

`func (o *ClusterRemoveHostsReqCluster) HasCephHostIds() bool`

HasCephHostIds returns a boolean if a field has been set.

### GetOspHostIds

`func (o *ClusterRemoveHostsReqCluster) GetOspHostIds() []int64`

GetOspHostIds returns the OspHostIds field if non-nil, zero value otherwise.

### GetOspHostIdsOk

`func (o *ClusterRemoveHostsReqCluster) GetOspHostIdsOk() (*[]int64, bool)`

GetOspHostIdsOk returns a tuple with the OspHostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspHostIds

`func (o *ClusterRemoveHostsReqCluster) SetOspHostIds(v []int64)`

SetOspHostIds sets OspHostIds field to given value.

### HasOspHostIds

`func (o *ClusterRemoveHostsReqCluster) HasOspHostIds() bool`

HasOspHostIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


