# UsersResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]UserRecord**](UserRecord.md) | users | 

## Methods

### NewUsersResp

`func NewUsersResp(users []UserRecord, ) *UsersResp`

NewUsersResp instantiates a new UsersResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersRespWithDefaults

`func NewUsersRespWithDefaults() *UsersResp`

NewUsersRespWithDefaults instantiates a new UsersResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UsersResp) GetUsers() []UserRecord`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersResp) GetUsersOk() (*[]UserRecord, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersResp) SetUsers(v []UserRecord)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


