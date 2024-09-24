# OSStorageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePoolIds** | Pointer to **[]int64** |  | [optional] 
**ClassId** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InactivePoolIds** | Pointer to **[]int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OsPolicy** | Pointer to [**ObjectStoragePolicy**](ObjectStoragePolicy.md) |  | [optional] 
**OsPolicyId** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOSStorageClass

`func NewOSStorageClass() *OSStorageClass`

NewOSStorageClass instantiates a new OSStorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSStorageClassWithDefaults

`func NewOSStorageClassWithDefaults() *OSStorageClass`

NewOSStorageClassWithDefaults instantiates a new OSStorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePoolIds

`func (o *OSStorageClass) GetActivePoolIds() []int64`

GetActivePoolIds returns the ActivePoolIds field if non-nil, zero value otherwise.

### GetActivePoolIdsOk

`func (o *OSStorageClass) GetActivePoolIdsOk() (*[]int64, bool)`

GetActivePoolIdsOk returns a tuple with the ActivePoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePoolIds

`func (o *OSStorageClass) SetActivePoolIds(v []int64)`

SetActivePoolIds sets ActivePoolIds field to given value.

### HasActivePoolIds

`func (o *OSStorageClass) HasActivePoolIds() bool`

HasActivePoolIds returns a boolean if a field has been set.

### GetClassId

`func (o *OSStorageClass) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OSStorageClass) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OSStorageClass) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *OSStorageClass) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetCluster

`func (o *OSStorageClass) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSStorageClass) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSStorageClass) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSStorageClass) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSStorageClass) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSStorageClass) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSStorageClass) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSStorageClass) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *OSStorageClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OSStorageClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OSStorageClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OSStorageClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *OSStorageClass) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSStorageClass) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSStorageClass) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSStorageClass) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInactivePoolIds

`func (o *OSStorageClass) GetInactivePoolIds() []int64`

GetInactivePoolIds returns the InactivePoolIds field if non-nil, zero value otherwise.

### GetInactivePoolIdsOk

`func (o *OSStorageClass) GetInactivePoolIdsOk() (*[]int64, bool)`

GetInactivePoolIdsOk returns a tuple with the InactivePoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivePoolIds

`func (o *OSStorageClass) SetInactivePoolIds(v []int64)`

SetInactivePoolIds sets InactivePoolIds field to given value.

### HasInactivePoolIds

`func (o *OSStorageClass) HasInactivePoolIds() bool`

HasInactivePoolIds returns a boolean if a field has been set.

### GetName

`func (o *OSStorageClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSStorageClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSStorageClass) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OSStorageClass) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsPolicy

`func (o *OSStorageClass) GetOsPolicy() ObjectStoragePolicy`

GetOsPolicy returns the OsPolicy field if non-nil, zero value otherwise.

### GetOsPolicyOk

`func (o *OSStorageClass) GetOsPolicyOk() (*ObjectStoragePolicy, bool)`

GetOsPolicyOk returns a tuple with the OsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPolicy

`func (o *OSStorageClass) SetOsPolicy(v ObjectStoragePolicy)`

SetOsPolicy sets OsPolicy field to given value.

### HasOsPolicy

`func (o *OSStorageClass) HasOsPolicy() bool`

HasOsPolicy returns a boolean if a field has been set.

### GetOsPolicyId

`func (o *OSStorageClass) GetOsPolicyId() int64`

GetOsPolicyId returns the OsPolicyId field if non-nil, zero value otherwise.

### GetOsPolicyIdOk

`func (o *OSStorageClass) GetOsPolicyIdOk() (*int64, bool)`

GetOsPolicyIdOk returns a tuple with the OsPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPolicyId

`func (o *OSStorageClass) SetOsPolicyId(v int64)`

SetOsPolicyId sets OsPolicyId field to given value.

### HasOsPolicyId

`func (o *OSStorageClass) HasOsPolicyId() bool`

HasOsPolicyId returns a boolean if a field has been set.

### GetStatus

`func (o *OSStorageClass) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSStorageClass) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSStorageClass) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSStorageClass) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *OSStorageClass) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSStorageClass) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSStorageClass) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSStorageClass) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


