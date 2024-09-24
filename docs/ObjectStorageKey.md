# ObjectStorageKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Reserved** | Pointer to **bool** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to [**ObjectStorageUserNestview**](ObjectStorageUserNestview.md) |  | [optional] 

## Methods

### NewObjectStorageKey

`func NewObjectStorageKey() *ObjectStorageKey`

NewObjectStorageKey instantiates a new ObjectStorageKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageKeyWithDefaults

`func NewObjectStorageKeyWithDefaults() *ObjectStorageKey`

NewObjectStorageKeyWithDefaults instantiates a new ObjectStorageKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *ObjectStorageKey) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ObjectStorageKey) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ObjectStorageKey) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ObjectStorageKey) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetCluster

`func (o *ObjectStorageKey) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageKey) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageKey) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageKey) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageKey) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageKey) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageKey) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageKey) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageKey) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageKey) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageKey) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReserved

`func (o *ObjectStorageKey) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *ObjectStorageKey) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *ObjectStorageKey) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *ObjectStorageKey) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetSecretKey

`func (o *ObjectStorageKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ObjectStorageKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ObjectStorageKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ObjectStorageKey) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageKey) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageKey) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ObjectStorageKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObjectStorageKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObjectStorageKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ObjectStorageKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *ObjectStorageKey) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ObjectStorageKey) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ObjectStorageKey) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ObjectStorageKey) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUser

`func (o *ObjectStorageKey) GetUser() ObjectStorageUserNestview`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ObjectStorageKey) GetUserOk() (*ObjectStorageUserNestview, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ObjectStorageKey) SetUser(v ObjectStorageUserNestview)`

SetUser sets User field to given value.

### HasUser

`func (o *ObjectStorageKey) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


