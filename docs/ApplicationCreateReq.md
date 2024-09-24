# ApplicationCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**ApplicationCreateReqApplication**](ApplicationCreateReqApplication.md) |  | [optional] 

## Methods

### NewApplicationCreateReq

`func NewApplicationCreateReq() *ApplicationCreateReq`

NewApplicationCreateReq instantiates a new ApplicationCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCreateReqWithDefaults

`func NewApplicationCreateReqWithDefaults() *ApplicationCreateReq`

NewApplicationCreateReqWithDefaults instantiates a new ApplicationCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ApplicationCreateReq) GetApplication() ApplicationCreateReqApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApplicationCreateReq) GetApplicationOk() (*ApplicationCreateReqApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApplicationCreateReq) SetApplication(v ApplicationCreateReqApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApplicationCreateReq) HasApplication() bool`

HasApplication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


