# DfsQosCreateReqQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DfsPath** | **string** | path of qos | 
**DfsQosPolicyId** | **int64** | qos policy id of qos | 
**DfsRootfsId** | **int64** | id of dfs rootfs | 

## Methods

### NewDfsQosCreateReqQos

`func NewDfsQosCreateReqQos(dfsPath string, dfsQosPolicyId int64, dfsRootfsId int64, ) *DfsQosCreateReqQos`

NewDfsQosCreateReqQos instantiates a new DfsQosCreateReqQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsQosCreateReqQosWithDefaults

`func NewDfsQosCreateReqQosWithDefaults() *DfsQosCreateReqQos`

NewDfsQosCreateReqQosWithDefaults instantiates a new DfsQosCreateReqQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDfsPath

`func (o *DfsQosCreateReqQos) GetDfsPath() string`

GetDfsPath returns the DfsPath field if non-nil, zero value otherwise.

### GetDfsPathOk

`func (o *DfsQosCreateReqQos) GetDfsPathOk() (*string, bool)`

GetDfsPathOk returns a tuple with the DfsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsPath

`func (o *DfsQosCreateReqQos) SetDfsPath(v string)`

SetDfsPath sets DfsPath field to given value.


### GetDfsQosPolicyId

`func (o *DfsQosCreateReqQos) GetDfsQosPolicyId() int64`

GetDfsQosPolicyId returns the DfsQosPolicyId field if non-nil, zero value otherwise.

### GetDfsQosPolicyIdOk

`func (o *DfsQosCreateReqQos) GetDfsQosPolicyIdOk() (*int64, bool)`

GetDfsQosPolicyIdOk returns a tuple with the DfsQosPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsQosPolicyId

`func (o *DfsQosCreateReqQos) SetDfsQosPolicyId(v int64)`

SetDfsQosPolicyId sets DfsQosPolicyId field to given value.


### GetDfsRootfsId

`func (o *DfsQosCreateReqQos) GetDfsRootfsId() int64`

GetDfsRootfsId returns the DfsRootfsId field if non-nil, zero value otherwise.

### GetDfsRootfsIdOk

`func (o *DfsQosCreateReqQos) GetDfsRootfsIdOk() (*int64, bool)`

GetDfsRootfsIdOk returns a tuple with the DfsRootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfsId

`func (o *DfsQosCreateReqQos) SetDfsRootfsId(v int64)`

SetDfsRootfsId sets DfsRootfsId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


