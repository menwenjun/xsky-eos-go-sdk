# XDCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | path of dfs xdcache | 
**RootfsId** | **int64** | dfs rootfs id | 

## Methods

### NewXDCache

`func NewXDCache(path string, rootfsId int64, ) *XDCache`

NewXDCache instantiates a new XDCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXDCacheWithDefaults

`func NewXDCacheWithDefaults() *XDCache`

NewXDCacheWithDefaults instantiates a new XDCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *XDCache) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *XDCache) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *XDCache) SetPath(v string)`

SetPath sets Path field to given value.


### GetRootfsId

`func (o *XDCache) GetRootfsId() int64`

GetRootfsId returns the RootfsId field if non-nil, zero value otherwise.

### GetRootfsIdOk

`func (o *XDCache) GetRootfsIdOk() (*int64, bool)`

GetRootfsIdOk returns a tuple with the RootfsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfsId

`func (o *XDCache) SetRootfsId(v int64)`

SetRootfsId sets RootfsId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


