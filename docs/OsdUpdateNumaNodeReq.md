# OsdUpdateNumaNodeReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumaApplyPolicy** | Pointer to **string** | apply policy in binding numa | [optional] 
**NumaNode** | Pointer to **int64** | numa node id | [optional] 

## Methods

### NewOsdUpdateNumaNodeReq

`func NewOsdUpdateNumaNodeReq() *OsdUpdateNumaNodeReq`

NewOsdUpdateNumaNodeReq instantiates a new OsdUpdateNumaNodeReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsdUpdateNumaNodeReqWithDefaults

`func NewOsdUpdateNumaNodeReqWithDefaults() *OsdUpdateNumaNodeReq`

NewOsdUpdateNumaNodeReqWithDefaults instantiates a new OsdUpdateNumaNodeReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumaApplyPolicy

`func (o *OsdUpdateNumaNodeReq) GetNumaApplyPolicy() string`

GetNumaApplyPolicy returns the NumaApplyPolicy field if non-nil, zero value otherwise.

### GetNumaApplyPolicyOk

`func (o *OsdUpdateNumaNodeReq) GetNumaApplyPolicyOk() (*string, bool)`

GetNumaApplyPolicyOk returns a tuple with the NumaApplyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaApplyPolicy

`func (o *OsdUpdateNumaNodeReq) SetNumaApplyPolicy(v string)`

SetNumaApplyPolicy sets NumaApplyPolicy field to given value.

### HasNumaApplyPolicy

`func (o *OsdUpdateNumaNodeReq) HasNumaApplyPolicy() bool`

HasNumaApplyPolicy returns a boolean if a field has been set.

### GetNumaNode

`func (o *OsdUpdateNumaNodeReq) GetNumaNode() int64`

GetNumaNode returns the NumaNode field if non-nil, zero value otherwise.

### GetNumaNodeOk

`func (o *OsdUpdateNumaNodeReq) GetNumaNodeOk() (*int64, bool)`

GetNumaNodeOk returns a tuple with the NumaNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaNode

`func (o *OsdUpdateNumaNodeReq) SetNumaNode(v int64)`

SetNumaNode sets NumaNode field to given value.

### HasNumaNode

`func (o *OsdUpdateNumaNodeReq) HasNumaNode() bool`

HasNumaNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


