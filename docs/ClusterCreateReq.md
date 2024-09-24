# ClusterCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterCreateReqCluster**](ClusterCreateReqCluster.md) |  | [optional] 

## Methods

### NewClusterCreateReq

`func NewClusterCreateReq() *ClusterCreateReq`

NewClusterCreateReq instantiates a new ClusterCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateReqWithDefaults

`func NewClusterCreateReqWithDefaults() *ClusterCreateReq`

NewClusterCreateReqWithDefaults instantiates a new ClusterCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ClusterCreateReq) GetCluster() ClusterCreateReqCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClusterCreateReq) GetClusterOk() (*ClusterCreateReqCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClusterCreateReq) SetCluster(v ClusterCreateReqCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ClusterCreateReq) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


