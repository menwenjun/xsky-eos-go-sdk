# ObjectStorageKeyCreateReqKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** |  | 
**SecretKey** | Pointer to **string** |  | [optional] 
**UserId** | **int64** |  | 

## Methods

### NewObjectStorageKeyCreateReqKey

`func NewObjectStorageKeyCreateReqKey(accessKey string, userId int64, ) *ObjectStorageKeyCreateReqKey`

NewObjectStorageKeyCreateReqKey instantiates a new ObjectStorageKeyCreateReqKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageKeyCreateReqKeyWithDefaults

`func NewObjectStorageKeyCreateReqKeyWithDefaults() *ObjectStorageKeyCreateReqKey`

NewObjectStorageKeyCreateReqKeyWithDefaults instantiates a new ObjectStorageKeyCreateReqKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *ObjectStorageKeyCreateReqKey) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ObjectStorageKeyCreateReqKey) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ObjectStorageKeyCreateReqKey) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *ObjectStorageKeyCreateReqKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ObjectStorageKeyCreateReqKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ObjectStorageKeyCreateReqKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ObjectStorageKeyCreateReqKey) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUserId

`func (o *ObjectStorageKeyCreateReqKey) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ObjectStorageKeyCreateReqKey) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ObjectStorageKeyCreateReqKey) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


