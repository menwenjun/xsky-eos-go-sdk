# ObjectStorageUsersResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsUsers** | [**[]ObjectStorageUserRecord**](ObjectStorageUserRecord.md) | object storage users | 

## Methods

### NewObjectStorageUsersResp

`func NewObjectStorageUsersResp(osUsers []ObjectStorageUserRecord, ) *ObjectStorageUsersResp`

NewObjectStorageUsersResp instantiates a new ObjectStorageUsersResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageUsersRespWithDefaults

`func NewObjectStorageUsersRespWithDefaults() *ObjectStorageUsersResp`

NewObjectStorageUsersRespWithDefaults instantiates a new ObjectStorageUsersResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsUsers

`func (o *ObjectStorageUsersResp) GetOsUsers() []ObjectStorageUserRecord`

GetOsUsers returns the OsUsers field if non-nil, zero value otherwise.

### GetOsUsersOk

`func (o *ObjectStorageUsersResp) GetOsUsersOk() (*[]ObjectStorageUserRecord, bool)`

GetOsUsersOk returns a tuple with the OsUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUsers

`func (o *ObjectStorageUsersResp) SetOsUsers(v []ObjectStorageUserRecord)`

SetOsUsers sets OsUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


