# UserResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**UserRecord**](UserRecord.md) |  | 

## Methods

### NewUserResp

`func NewUserResp(user UserRecord, ) *UserResp`

NewUserResp instantiates a new UserResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRespWithDefaults

`func NewUserRespWithDefaults() *UserResp`

NewUserRespWithDefaults instantiates a new UserResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserResp) GetUser() UserRecord`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserResp) GetUserOk() (*UserRecord, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserResp) SetUser(v UserRecord)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


