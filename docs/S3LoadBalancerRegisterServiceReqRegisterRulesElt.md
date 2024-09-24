# S3LoadBalancerRegisterServiceReqRegisterRulesElt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecognitionRules** | **[]string** | recognition rules of service | 
**ServiceType** | Pointer to **string** | type of service | [optional] 
**TargetEndpoints** | [**[]S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt**](S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt.md) | target endpoints of service | 

## Methods

### NewS3LoadBalancerRegisterServiceReqRegisterRulesElt

`func NewS3LoadBalancerRegisterServiceReqRegisterRulesElt(recognitionRules []string, targetEndpoints []S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt, ) *S3LoadBalancerRegisterServiceReqRegisterRulesElt`

NewS3LoadBalancerRegisterServiceReqRegisterRulesElt instantiates a new S3LoadBalancerRegisterServiceReqRegisterRulesElt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancerRegisterServiceReqRegisterRulesEltWithDefaults

`func NewS3LoadBalancerRegisterServiceReqRegisterRulesEltWithDefaults() *S3LoadBalancerRegisterServiceReqRegisterRulesElt`

NewS3LoadBalancerRegisterServiceReqRegisterRulesEltWithDefaults instantiates a new S3LoadBalancerRegisterServiceReqRegisterRulesElt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecognitionRules

`func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetRecognitionRules() []string`

GetRecognitionRules returns the RecognitionRules field if non-nil, zero value otherwise.

### GetRecognitionRulesOk

`func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetRecognitionRulesOk() (*[]string, bool)`

GetRecognitionRulesOk returns a tuple with the RecognitionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecognitionRules

`func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) SetRecognitionRules(v []string)`

SetRecognitionRules sets RecognitionRules field to given value.


### GetServiceType

`func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetTargetEndpoints

`func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetTargetEndpoints() []S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt`

GetTargetEndpoints returns the TargetEndpoints field if non-nil, zero value otherwise.

### GetTargetEndpointsOk

`func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) GetTargetEndpointsOk() (*[]S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt, bool)`

GetTargetEndpointsOk returns a tuple with the TargetEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEndpoints

`func (o *S3LoadBalancerRegisterServiceReqRegisterRulesElt) SetTargetEndpoints(v []S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt)`

SetTargetEndpoints sets TargetEndpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


