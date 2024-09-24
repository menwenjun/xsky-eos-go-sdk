# ApplicationUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**ApplicationUpdateReqApplication**](ApplicationUpdateReqApplication.md) |  | [optional] 

## Methods

### NewApplicationUpdateReq

`func NewApplicationUpdateReq() *ApplicationUpdateReq`

NewApplicationUpdateReq instantiates a new ApplicationUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUpdateReqWithDefaults

`func NewApplicationUpdateReqWithDefaults() *ApplicationUpdateReq`

NewApplicationUpdateReqWithDefaults instantiates a new ApplicationUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ApplicationUpdateReq) GetApplication() ApplicationUpdateReqApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApplicationUpdateReq) GetApplicationOk() (*ApplicationUpdateReqApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApplicationUpdateReq) SetApplication(v ApplicationUpdateReqApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApplicationUpdateReq) HasApplication() bool`

HasApplication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


