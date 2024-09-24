# AuthLoginReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | [**AuthLoginReqAuth**](AuthLoginReqAuth.md) |  | 

## Methods

### NewAuthLoginReq

`func NewAuthLoginReq(auth AuthLoginReqAuth, ) *AuthLoginReq`

NewAuthLoginReq instantiates a new AuthLoginReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLoginReqWithDefaults

`func NewAuthLoginReqWithDefaults() *AuthLoginReq`

NewAuthLoginReqWithDefaults instantiates a new AuthLoginReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *AuthLoginReq) GetAuth() AuthLoginReqAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *AuthLoginReq) GetAuthOk() (*AuthLoginReqAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *AuthLoginReq) SetAuth(v AuthLoginReqAuth)`

SetAuth sets Auth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


