# S3LoadBalancerRegisterServiceReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegisterRules** | [**[]S3LoadBalancerRegisterServiceReqRegisterRulesElt**](S3LoadBalancerRegisterServiceReqRegisterRulesElt.md) |  | 

## Methods

### NewS3LoadBalancerRegisterServiceReq

`func NewS3LoadBalancerRegisterServiceReq(registerRules []S3LoadBalancerRegisterServiceReqRegisterRulesElt, ) *S3LoadBalancerRegisterServiceReq`

NewS3LoadBalancerRegisterServiceReq instantiates a new S3LoadBalancerRegisterServiceReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancerRegisterServiceReqWithDefaults

`func NewS3LoadBalancerRegisterServiceReqWithDefaults() *S3LoadBalancerRegisterServiceReq`

NewS3LoadBalancerRegisterServiceReqWithDefaults instantiates a new S3LoadBalancerRegisterServiceReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegisterRules

`func (o *S3LoadBalancerRegisterServiceReq) GetRegisterRules() []S3LoadBalancerRegisterServiceReqRegisterRulesElt`

GetRegisterRules returns the RegisterRules field if non-nil, zero value otherwise.

### GetRegisterRulesOk

`func (o *S3LoadBalancerRegisterServiceReq) GetRegisterRulesOk() (*[]S3LoadBalancerRegisterServiceReqRegisterRulesElt, bool)`

GetRegisterRulesOk returns a tuple with the RegisterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterRules

`func (o *S3LoadBalancerRegisterServiceReq) SetRegisterRules(v []S3LoadBalancerRegisterServiceReqRegisterRulesElt)`

SetRegisterRules sets RegisterRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


