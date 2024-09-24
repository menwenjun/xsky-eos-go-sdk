# SendVerificationReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verification** | Pointer to [**SendVerificationReqVerification**](SendVerificationReqVerification.md) |  | [optional] 

## Methods

### NewSendVerificationReq

`func NewSendVerificationReq() *SendVerificationReq`

NewSendVerificationReq instantiates a new SendVerificationReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendVerificationReqWithDefaults

`func NewSendVerificationReqWithDefaults() *SendVerificationReq`

NewSendVerificationReqWithDefaults instantiates a new SendVerificationReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerification

`func (o *SendVerificationReq) GetVerification() SendVerificationReqVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *SendVerificationReq) GetVerificationOk() (*SendVerificationReqVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *SendVerificationReq) SetVerification(v SendVerificationReqVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *SendVerificationReq) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


