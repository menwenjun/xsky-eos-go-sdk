# DfsRootfsCreateReqRootfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePoolIds** | Pointer to **[]int64** | id array of active pools | [optional] 
**DefaultClassName** | Pointer to **string** | name of default class | [optional] 
**DefaultClassWritePolicy** | Pointer to **string** | write policy of default class | [optional] 
**Description** | Pointer to **string** | description of rootfs | [optional] 
**DfsAuditLog** | [**DfsRootfsCreateReqRootfsAuditLog**](DfsRootfsCreateReqRootfsAuditLog.md) |  | 
**MetadataClusterId** | **int64** | metadata cluster id | 
**Name** | **string** | name of rootfs | 
**PoolIds** | Pointer to **[]int64** | id array of pools | [optional] 
**PoolPolicies** | Pointer to [**[]PoolPolicy**](PoolPolicy.md) | active pool policy array of default class | [optional] 
**StorageClasses** | Pointer to [**[]CustomStorageClass**](CustomStorageClass.md) | CustomStorageClasses create custom dfs storage class | [optional] 

## Methods

### NewDfsRootfsCreateReqRootfs

`func NewDfsRootfsCreateReqRootfs(dfsAuditLog DfsRootfsCreateReqRootfsAuditLog, metadataClusterId int64, name string, ) *DfsRootfsCreateReqRootfs`

NewDfsRootfsCreateReqRootfs instantiates a new DfsRootfsCreateReqRootfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsRootfsCreateReqRootfsWithDefaults

`func NewDfsRootfsCreateReqRootfsWithDefaults() *DfsRootfsCreateReqRootfs`

NewDfsRootfsCreateReqRootfsWithDefaults instantiates a new DfsRootfsCreateReqRootfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePoolIds

`func (o *DfsRootfsCreateReqRootfs) GetActivePoolIds() []int64`

GetActivePoolIds returns the ActivePoolIds field if non-nil, zero value otherwise.

### GetActivePoolIdsOk

`func (o *DfsRootfsCreateReqRootfs) GetActivePoolIdsOk() (*[]int64, bool)`

GetActivePoolIdsOk returns a tuple with the ActivePoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePoolIds

`func (o *DfsRootfsCreateReqRootfs) SetActivePoolIds(v []int64)`

SetActivePoolIds sets ActivePoolIds field to given value.

### HasActivePoolIds

`func (o *DfsRootfsCreateReqRootfs) HasActivePoolIds() bool`

HasActivePoolIds returns a boolean if a field has been set.

### GetDefaultClassName

`func (o *DfsRootfsCreateReqRootfs) GetDefaultClassName() string`

GetDefaultClassName returns the DefaultClassName field if non-nil, zero value otherwise.

### GetDefaultClassNameOk

`func (o *DfsRootfsCreateReqRootfs) GetDefaultClassNameOk() (*string, bool)`

GetDefaultClassNameOk returns a tuple with the DefaultClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClassName

`func (o *DfsRootfsCreateReqRootfs) SetDefaultClassName(v string)`

SetDefaultClassName sets DefaultClassName field to given value.

### HasDefaultClassName

`func (o *DfsRootfsCreateReqRootfs) HasDefaultClassName() bool`

HasDefaultClassName returns a boolean if a field has been set.

### GetDefaultClassWritePolicy

`func (o *DfsRootfsCreateReqRootfs) GetDefaultClassWritePolicy() string`

GetDefaultClassWritePolicy returns the DefaultClassWritePolicy field if non-nil, zero value otherwise.

### GetDefaultClassWritePolicyOk

`func (o *DfsRootfsCreateReqRootfs) GetDefaultClassWritePolicyOk() (*string, bool)`

GetDefaultClassWritePolicyOk returns a tuple with the DefaultClassWritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClassWritePolicy

`func (o *DfsRootfsCreateReqRootfs) SetDefaultClassWritePolicy(v string)`

SetDefaultClassWritePolicy sets DefaultClassWritePolicy field to given value.

### HasDefaultClassWritePolicy

`func (o *DfsRootfsCreateReqRootfs) HasDefaultClassWritePolicy() bool`

HasDefaultClassWritePolicy returns a boolean if a field has been set.

### GetDescription

`func (o *DfsRootfsCreateReqRootfs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsRootfsCreateReqRootfs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsRootfsCreateReqRootfs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsRootfsCreateReqRootfs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsAuditLog

`func (o *DfsRootfsCreateReqRootfs) GetDfsAuditLog() DfsRootfsCreateReqRootfsAuditLog`

GetDfsAuditLog returns the DfsAuditLog field if non-nil, zero value otherwise.

### GetDfsAuditLogOk

`func (o *DfsRootfsCreateReqRootfs) GetDfsAuditLogOk() (*DfsRootfsCreateReqRootfsAuditLog, bool)`

GetDfsAuditLogOk returns a tuple with the DfsAuditLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsAuditLog

`func (o *DfsRootfsCreateReqRootfs) SetDfsAuditLog(v DfsRootfsCreateReqRootfsAuditLog)`

SetDfsAuditLog sets DfsAuditLog field to given value.


### GetMetadataClusterId

`func (o *DfsRootfsCreateReqRootfs) GetMetadataClusterId() int64`

GetMetadataClusterId returns the MetadataClusterId field if non-nil, zero value otherwise.

### GetMetadataClusterIdOk

`func (o *DfsRootfsCreateReqRootfs) GetMetadataClusterIdOk() (*int64, bool)`

GetMetadataClusterIdOk returns a tuple with the MetadataClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataClusterId

`func (o *DfsRootfsCreateReqRootfs) SetMetadataClusterId(v int64)`

SetMetadataClusterId sets MetadataClusterId field to given value.


### GetName

`func (o *DfsRootfsCreateReqRootfs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsRootfsCreateReqRootfs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsRootfsCreateReqRootfs) SetName(v string)`

SetName sets Name field to given value.


### GetPoolIds

`func (o *DfsRootfsCreateReqRootfs) GetPoolIds() []int64`

GetPoolIds returns the PoolIds field if non-nil, zero value otherwise.

### GetPoolIdsOk

`func (o *DfsRootfsCreateReqRootfs) GetPoolIdsOk() (*[]int64, bool)`

GetPoolIdsOk returns a tuple with the PoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIds

`func (o *DfsRootfsCreateReqRootfs) SetPoolIds(v []int64)`

SetPoolIds sets PoolIds field to given value.

### HasPoolIds

`func (o *DfsRootfsCreateReqRootfs) HasPoolIds() bool`

HasPoolIds returns a boolean if a field has been set.

### GetPoolPolicies

`func (o *DfsRootfsCreateReqRootfs) GetPoolPolicies() []PoolPolicy`

GetPoolPolicies returns the PoolPolicies field if non-nil, zero value otherwise.

### GetPoolPoliciesOk

`func (o *DfsRootfsCreateReqRootfs) GetPoolPoliciesOk() (*[]PoolPolicy, bool)`

GetPoolPoliciesOk returns a tuple with the PoolPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPolicies

`func (o *DfsRootfsCreateReqRootfs) SetPoolPolicies(v []PoolPolicy)`

SetPoolPolicies sets PoolPolicies field to given value.

### HasPoolPolicies

`func (o *DfsRootfsCreateReqRootfs) HasPoolPolicies() bool`

HasPoolPolicies returns a boolean if a field has been set.

### GetStorageClasses

`func (o *DfsRootfsCreateReqRootfs) GetStorageClasses() []CustomStorageClass`

GetStorageClasses returns the StorageClasses field if non-nil, zero value otherwise.

### GetStorageClassesOk

`func (o *DfsRootfsCreateReqRootfs) GetStorageClassesOk() (*[]CustomStorageClass, bool)`

GetStorageClassesOk returns a tuple with the StorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClasses

`func (o *DfsRootfsCreateReqRootfs) SetStorageClasses(v []CustomStorageClass)`

SetStorageClasses sets StorageClasses field to given value.

### HasStorageClasses

`func (o *DfsRootfsCreateReqRootfs) HasStorageClasses() bool`

HasStorageClasses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


