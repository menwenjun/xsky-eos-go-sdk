# TopologyFromOsdReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsdIds** | **[]int64** |  | 
**ProtectDomainId** | Pointer to **int64** |  | [optional] 

## Methods

### NewTopologyFromOsdReq

`func NewTopologyFromOsdReq(osdIds []int64, ) *TopologyFromOsdReq`

NewTopologyFromOsdReq instantiates a new TopologyFromOsdReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyFromOsdReqWithDefaults

`func NewTopologyFromOsdReqWithDefaults() *TopologyFromOsdReq`

NewTopologyFromOsdReqWithDefaults instantiates a new TopologyFromOsdReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsdIds

`func (o *TopologyFromOsdReq) GetOsdIds() []int64`

GetOsdIds returns the OsdIds field if non-nil, zero value otherwise.

### GetOsdIdsOk

`func (o *TopologyFromOsdReq) GetOsdIdsOk() (*[]int64, bool)`

GetOsdIdsOk returns a tuple with the OsdIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdIds

`func (o *TopologyFromOsdReq) SetOsdIds(v []int64)`

SetOsdIds sets OsdIds field to given value.


### GetProtectDomainId

`func (o *TopologyFromOsdReq) GetProtectDomainId() int64`

GetProtectDomainId returns the ProtectDomainId field if non-nil, zero value otherwise.

### GetProtectDomainIdOk

`func (o *TopologyFromOsdReq) GetProtectDomainIdOk() (*int64, bool)`

GetProtectDomainIdOk returns a tuple with the ProtectDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectDomainId

`func (o *TopologyFromOsdReq) SetProtectDomainId(v int64)`

SetProtectDomainId sets ProtectDomainId field to given value.

### HasProtectDomainId

`func (o *TopologyFromOsdReq) HasProtectDomainId() bool`

HasProtectDomainId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


