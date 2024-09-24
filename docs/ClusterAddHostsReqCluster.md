# ClusterAddHostsReqCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSelectOspMonitorRole** | Pointer to **bool** | auto select osp monitor role if true | [optional] 
**CephHosts** | Pointer to [**[]CephHostReq**](CephHostReq.md) | ceph host list | [optional] 
**OspHosts** | Pointer to [**[]OspHostReq**](OspHostReq.md) | osp host list | [optional] 

## Methods

### NewClusterAddHostsReqCluster

`func NewClusterAddHostsReqCluster() *ClusterAddHostsReqCluster`

NewClusterAddHostsReqCluster instantiates a new ClusterAddHostsReqCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAddHostsReqClusterWithDefaults

`func NewClusterAddHostsReqClusterWithDefaults() *ClusterAddHostsReqCluster`

NewClusterAddHostsReqClusterWithDefaults instantiates a new ClusterAddHostsReqCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSelectOspMonitorRole

`func (o *ClusterAddHostsReqCluster) GetAutoSelectOspMonitorRole() bool`

GetAutoSelectOspMonitorRole returns the AutoSelectOspMonitorRole field if non-nil, zero value otherwise.

### GetAutoSelectOspMonitorRoleOk

`func (o *ClusterAddHostsReqCluster) GetAutoSelectOspMonitorRoleOk() (*bool, bool)`

GetAutoSelectOspMonitorRoleOk returns a tuple with the AutoSelectOspMonitorRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelectOspMonitorRole

`func (o *ClusterAddHostsReqCluster) SetAutoSelectOspMonitorRole(v bool)`

SetAutoSelectOspMonitorRole sets AutoSelectOspMonitorRole field to given value.

### HasAutoSelectOspMonitorRole

`func (o *ClusterAddHostsReqCluster) HasAutoSelectOspMonitorRole() bool`

HasAutoSelectOspMonitorRole returns a boolean if a field has been set.

### GetCephHosts

`func (o *ClusterAddHostsReqCluster) GetCephHosts() []CephHostReq`

GetCephHosts returns the CephHosts field if non-nil, zero value otherwise.

### GetCephHostsOk

`func (o *ClusterAddHostsReqCluster) GetCephHostsOk() (*[]CephHostReq, bool)`

GetCephHostsOk returns a tuple with the CephHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephHosts

`func (o *ClusterAddHostsReqCluster) SetCephHosts(v []CephHostReq)`

SetCephHosts sets CephHosts field to given value.

### HasCephHosts

`func (o *ClusterAddHostsReqCluster) HasCephHosts() bool`

HasCephHosts returns a boolean if a field has been set.

### GetOspHosts

`func (o *ClusterAddHostsReqCluster) GetOspHosts() []OspHostReq`

GetOspHosts returns the OspHosts field if non-nil, zero value otherwise.

### GetOspHostsOk

`func (o *ClusterAddHostsReqCluster) GetOspHostsOk() (*[]OspHostReq, bool)`

GetOspHostsOk returns a tuple with the OspHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspHosts

`func (o *ClusterAddHostsReqCluster) SetOspHosts(v []OspHostReq)`

SetOspHosts sets OspHosts field to given value.

### HasOspHosts

`func (o *ClusterAddHostsReqCluster) HasOspHosts() bool`

HasOspHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


