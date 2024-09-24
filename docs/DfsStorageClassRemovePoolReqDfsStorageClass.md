# DfsStorageClassRemovePoolReqDfsStorageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolIds** | **[]int64** | pool ids to remove | 
**ToDefault** | **bool** | whether to move the pool to the default class after it is removed | 

## Methods

### NewDfsStorageClassRemovePoolReqDfsStorageClass

`func NewDfsStorageClassRemovePoolReqDfsStorageClass(poolIds []int64, toDefault bool, ) *DfsStorageClassRemovePoolReqDfsStorageClass`

NewDfsStorageClassRemovePoolReqDfsStorageClass instantiates a new DfsStorageClassRemovePoolReqDfsStorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsStorageClassRemovePoolReqDfsStorageClassWithDefaults

`func NewDfsStorageClassRemovePoolReqDfsStorageClassWithDefaults() *DfsStorageClassRemovePoolReqDfsStorageClass`

NewDfsStorageClassRemovePoolReqDfsStorageClassWithDefaults instantiates a new DfsStorageClassRemovePoolReqDfsStorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolIds

`func (o *DfsStorageClassRemovePoolReqDfsStorageClass) GetPoolIds() []int64`

GetPoolIds returns the PoolIds field if non-nil, zero value otherwise.

### GetPoolIdsOk

`func (o *DfsStorageClassRemovePoolReqDfsStorageClass) GetPoolIdsOk() (*[]int64, bool)`

GetPoolIdsOk returns a tuple with the PoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIds

`func (o *DfsStorageClassRemovePoolReqDfsStorageClass) SetPoolIds(v []int64)`

SetPoolIds sets PoolIds field to given value.


### GetToDefault

`func (o *DfsStorageClassRemovePoolReqDfsStorageClass) GetToDefault() bool`

GetToDefault returns the ToDefault field if non-nil, zero value otherwise.

### GetToDefaultOk

`func (o *DfsStorageClassRemovePoolReqDfsStorageClass) GetToDefaultOk() (*bool, bool)`

GetToDefaultOk returns a tuple with the ToDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDefault

`func (o *DfsStorageClassRemovePoolReqDfsStorageClass) SetToDefault(v bool)`

SetToDefault sets ToDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


