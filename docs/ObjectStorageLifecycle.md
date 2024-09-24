# ObjectStorageLifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Rules** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewObjectStorageLifecycle

`func NewObjectStorageLifecycle() *ObjectStorageLifecycle`

NewObjectStorageLifecycle instantiates a new ObjectStorageLifecycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageLifecycleWithDefaults

`func NewObjectStorageLifecycleWithDefaults() *ObjectStorageLifecycle`

NewObjectStorageLifecycleWithDefaults instantiates a new ObjectStorageLifecycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ObjectStorageLifecycle) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageLifecycle) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageLifecycle) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageLifecycle) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageLifecycle) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageLifecycle) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageLifecycle) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageLifecycle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRules

`func (o *ObjectStorageLifecycle) GetRules() []map[string]interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ObjectStorageLifecycle) GetRulesOk() (*[]map[string]interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ObjectStorageLifecycle) SetRules(v []map[string]interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *ObjectStorageLifecycle) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageLifecycle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageLifecycle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageLifecycle) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageLifecycle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


