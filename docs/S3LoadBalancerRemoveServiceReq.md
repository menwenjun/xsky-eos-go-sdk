# S3LoadBalancerRemoveServiceReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoveRules** | [**[]S3LoadBalancerRemoveServiceReqRemoveRulesElt**](S3LoadBalancerRemoveServiceReqRemoveRulesElt.md) |  | 

## Methods

### NewS3LoadBalancerRemoveServiceReq

`func NewS3LoadBalancerRemoveServiceReq(removeRules []S3LoadBalancerRemoveServiceReqRemoveRulesElt, ) *S3LoadBalancerRemoveServiceReq`

NewS3LoadBalancerRemoveServiceReq instantiates a new S3LoadBalancerRemoveServiceReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancerRemoveServiceReqWithDefaults

`func NewS3LoadBalancerRemoveServiceReqWithDefaults() *S3LoadBalancerRemoveServiceReq`

NewS3LoadBalancerRemoveServiceReqWithDefaults instantiates a new S3LoadBalancerRemoveServiceReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoveRules

`func (o *S3LoadBalancerRemoveServiceReq) GetRemoveRules() []S3LoadBalancerRemoveServiceReqRemoveRulesElt`

GetRemoveRules returns the RemoveRules field if non-nil, zero value otherwise.

### GetRemoveRulesOk

`func (o *S3LoadBalancerRemoveServiceReq) GetRemoveRulesOk() (*[]S3LoadBalancerRemoveServiceReqRemoveRulesElt, bool)`

GetRemoveRulesOk returns a tuple with the RemoveRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveRules

`func (o *S3LoadBalancerRemoveServiceReq) SetRemoveRules(v []S3LoadBalancerRemoveServiceReqRemoveRulesElt)`

SetRemoveRules sets RemoveRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


