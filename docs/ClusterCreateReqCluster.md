# ClusterCreateReqCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootnode** | Pointer to [**ClusterCreateBootNodeReq**](ClusterCreateBootNodeReq.md) |  | [optional] 
**FsId** | **string** | cluster fs id | 
**Name** | **string** | cluster name | 
**Network** | Pointer to [**ClusterCreateNetworkReq**](ClusterCreateNetworkReq.md) |  | [optional] 
**Type** | Pointer to **string** | cluster type: ceph or osp | [optional] 

## Methods

### NewClusterCreateReqCluster

`func NewClusterCreateReqCluster(fsId string, name string, ) *ClusterCreateReqCluster`

NewClusterCreateReqCluster instantiates a new ClusterCreateReqCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateReqClusterWithDefaults

`func NewClusterCreateReqClusterWithDefaults() *ClusterCreateReqCluster`

NewClusterCreateReqClusterWithDefaults instantiates a new ClusterCreateReqCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootnode

`func (o *ClusterCreateReqCluster) GetBootnode() ClusterCreateBootNodeReq`

GetBootnode returns the Bootnode field if non-nil, zero value otherwise.

### GetBootnodeOk

`func (o *ClusterCreateReqCluster) GetBootnodeOk() (*ClusterCreateBootNodeReq, bool)`

GetBootnodeOk returns a tuple with the Bootnode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootnode

`func (o *ClusterCreateReqCluster) SetBootnode(v ClusterCreateBootNodeReq)`

SetBootnode sets Bootnode field to given value.

### HasBootnode

`func (o *ClusterCreateReqCluster) HasBootnode() bool`

HasBootnode returns a boolean if a field has been set.

### GetFsId

`func (o *ClusterCreateReqCluster) GetFsId() string`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *ClusterCreateReqCluster) GetFsIdOk() (*string, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *ClusterCreateReqCluster) SetFsId(v string)`

SetFsId sets FsId field to given value.


### GetName

`func (o *ClusterCreateReqCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterCreateReqCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterCreateReqCluster) SetName(v string)`

SetName sets Name field to given value.


### GetNetwork

`func (o *ClusterCreateReqCluster) GetNetwork() ClusterCreateNetworkReq`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClusterCreateReqCluster) GetNetworkOk() (*ClusterCreateNetworkReq, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClusterCreateReqCluster) SetNetwork(v ClusterCreateNetworkReq)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ClusterCreateReqCluster) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetType

`func (o *ClusterCreateReqCluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterCreateReqCluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterCreateReqCluster) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterCreateReqCluster) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


