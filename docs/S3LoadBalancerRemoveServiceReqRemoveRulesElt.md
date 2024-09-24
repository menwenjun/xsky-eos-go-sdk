# S3LoadBalancerRemoveServiceReqRemoveRulesElt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OspZoneId** | **int64** | zone id of load balancers | 
**Rules** | **[]string** | rules of service | 

## Methods

### NewS3LoadBalancerRemoveServiceReqRemoveRulesElt

`func NewS3LoadBalancerRemoveServiceReqRemoveRulesElt(ospZoneId int64, rules []string, ) *S3LoadBalancerRemoveServiceReqRemoveRulesElt`

NewS3LoadBalancerRemoveServiceReqRemoveRulesElt instantiates a new S3LoadBalancerRemoveServiceReqRemoveRulesElt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancerRemoveServiceReqRemoveRulesEltWithDefaults

`func NewS3LoadBalancerRemoveServiceReqRemoveRulesEltWithDefaults() *S3LoadBalancerRemoveServiceReqRemoveRulesElt`

NewS3LoadBalancerRemoveServiceReqRemoveRulesEltWithDefaults instantiates a new S3LoadBalancerRemoveServiceReqRemoveRulesElt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOspZoneId

`func (o *S3LoadBalancerRemoveServiceReqRemoveRulesElt) GetOspZoneId() int64`

GetOspZoneId returns the OspZoneId field if non-nil, zero value otherwise.

### GetOspZoneIdOk

`func (o *S3LoadBalancerRemoveServiceReqRemoveRulesElt) GetOspZoneIdOk() (*int64, bool)`

GetOspZoneIdOk returns a tuple with the OspZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspZoneId

`func (o *S3LoadBalancerRemoveServiceReqRemoveRulesElt) SetOspZoneId(v int64)`

SetOspZoneId sets OspZoneId field to given value.


### GetRules

`func (o *S3LoadBalancerRemoveServiceReqRemoveRulesElt) GetRules() []string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *S3LoadBalancerRemoveServiceReqRemoveRulesElt) GetRulesOk() (*[]string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *S3LoadBalancerRemoveServiceReqRemoveRulesElt) SetRules(v []string)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


