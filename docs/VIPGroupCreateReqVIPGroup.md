# VIPGroupCreateReqVIPGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** |  | 
**Preempt** | Pointer to **bool** |  | [optional] 
**ResourceId** | **int64** |  | 
**ResourceType** | **string** |  | 
**Vips** | [**[]VIPArgs**](VIPArgs.md) |  | 

## Methods

### NewVIPGroupCreateReqVIPGroup

`func NewVIPGroupCreateReqVIPGroup(network string, resourceId int64, resourceType string, vips []VIPArgs, ) *VIPGroupCreateReqVIPGroup`

NewVIPGroupCreateReqVIPGroup instantiates a new VIPGroupCreateReqVIPGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVIPGroupCreateReqVIPGroupWithDefaults

`func NewVIPGroupCreateReqVIPGroupWithDefaults() *VIPGroupCreateReqVIPGroup`

NewVIPGroupCreateReqVIPGroupWithDefaults instantiates a new VIPGroupCreateReqVIPGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *VIPGroupCreateReqVIPGroup) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VIPGroupCreateReqVIPGroup) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VIPGroupCreateReqVIPGroup) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetPreempt

`func (o *VIPGroupCreateReqVIPGroup) GetPreempt() bool`

GetPreempt returns the Preempt field if non-nil, zero value otherwise.

### GetPreemptOk

`func (o *VIPGroupCreateReqVIPGroup) GetPreemptOk() (*bool, bool)`

GetPreemptOk returns a tuple with the Preempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreempt

`func (o *VIPGroupCreateReqVIPGroup) SetPreempt(v bool)`

SetPreempt sets Preempt field to given value.

### HasPreempt

`func (o *VIPGroupCreateReqVIPGroup) HasPreempt() bool`

HasPreempt returns a boolean if a field has been set.

### GetResourceId

`func (o *VIPGroupCreateReqVIPGroup) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *VIPGroupCreateReqVIPGroup) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *VIPGroupCreateReqVIPGroup) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *VIPGroupCreateReqVIPGroup) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *VIPGroupCreateReqVIPGroup) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *VIPGroupCreateReqVIPGroup) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetVips

`func (o *VIPGroupCreateReqVIPGroup) GetVips() []VIPArgs`

GetVips returns the Vips field if non-nil, zero value otherwise.

### GetVipsOk

`func (o *VIPGroupCreateReqVIPGroup) GetVipsOk() (*[]VIPArgs, bool)`

GetVipsOk returns a tuple with the Vips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVips

`func (o *VIPGroupCreateReqVIPGroup) SetVips(v []VIPArgs)`

SetVips sets Vips field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


