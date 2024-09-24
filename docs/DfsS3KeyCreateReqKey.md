# DfsS3KeyCreateReqKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | access key | 
**SecretKey** | Pointer to **string** | secret key | [optional] 
**UserId** | **int64** | user id | 

## Methods

### NewDfsS3KeyCreateReqKey

`func NewDfsS3KeyCreateReqKey(accessKey string, userId int64, ) *DfsS3KeyCreateReqKey`

NewDfsS3KeyCreateReqKey instantiates a new DfsS3KeyCreateReqKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsS3KeyCreateReqKeyWithDefaults

`func NewDfsS3KeyCreateReqKeyWithDefaults() *DfsS3KeyCreateReqKey`

NewDfsS3KeyCreateReqKeyWithDefaults instantiates a new DfsS3KeyCreateReqKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *DfsS3KeyCreateReqKey) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *DfsS3KeyCreateReqKey) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *DfsS3KeyCreateReqKey) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *DfsS3KeyCreateReqKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *DfsS3KeyCreateReqKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *DfsS3KeyCreateReqKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *DfsS3KeyCreateReqKey) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUserId

`func (o *DfsS3KeyCreateReqKey) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DfsS3KeyCreateReqKey) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DfsS3KeyCreateReqKey) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


