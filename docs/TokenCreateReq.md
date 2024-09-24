# TokenCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | [**TokenCreateReqAuth**](TokenCreateReqAuth.md) |  | 

## Methods

### NewTokenCreateReq

`func NewTokenCreateReq(auth TokenCreateReqAuth, ) *TokenCreateReq`

NewTokenCreateReq instantiates a new TokenCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreateReqWithDefaults

`func NewTokenCreateReqWithDefaults() *TokenCreateReq`

NewTokenCreateReqWithDefaults instantiates a new TokenCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *TokenCreateReq) GetAuth() TokenCreateReqAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *TokenCreateReq) GetAuthOk() (*TokenCreateReqAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *TokenCreateReq) SetAuth(v TokenCreateReqAuth)`

SetAuth sets Auth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


