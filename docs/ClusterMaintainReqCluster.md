# ClusterMaintainReqCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **int64** |  | [optional] 
**Maintained** | **bool** |  | 

## Methods

### NewClusterMaintainReqCluster

`func NewClusterMaintainReqCluster(maintained bool, ) *ClusterMaintainReqCluster`

NewClusterMaintainReqCluster instantiates a new ClusterMaintainReqCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMaintainReqClusterWithDefaults

`func NewClusterMaintainReqClusterWithDefaults() *ClusterMaintainReqCluster`

NewClusterMaintainReqClusterWithDefaults instantiates a new ClusterMaintainReqCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterMaintainReqCluster) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterMaintainReqCluster) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterMaintainReqCluster) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClusterMaintainReqCluster) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetMaintained

`func (o *ClusterMaintainReqCluster) GetMaintained() bool`

GetMaintained returns the Maintained field if non-nil, zero value otherwise.

### GetMaintainedOk

`func (o *ClusterMaintainReqCluster) GetMaintainedOk() (*bool, bool)`

GetMaintainedOk returns a tuple with the Maintained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintained

`func (o *ClusterMaintainReqCluster) SetMaintained(v bool)`

SetMaintained sets Maintained field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


