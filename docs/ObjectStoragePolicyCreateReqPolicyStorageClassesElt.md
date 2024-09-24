# ObjectStoragePolicyCreateReqPolicyStorageClassesElt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePoolIds** | **[]int64** |  | 
**Class** | Pointer to **int32** |  | [optional] 
**ClassId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InactivePoolIds** | Pointer to **[]int64** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewObjectStoragePolicyCreateReqPolicyStorageClassesElt

`func NewObjectStoragePolicyCreateReqPolicyStorageClassesElt(activePoolIds []int64, name string, ) *ObjectStoragePolicyCreateReqPolicyStorageClassesElt`

NewObjectStoragePolicyCreateReqPolicyStorageClassesElt instantiates a new ObjectStoragePolicyCreateReqPolicyStorageClassesElt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStoragePolicyCreateReqPolicyStorageClassesEltWithDefaults

`func NewObjectStoragePolicyCreateReqPolicyStorageClassesEltWithDefaults() *ObjectStoragePolicyCreateReqPolicyStorageClassesElt`

NewObjectStoragePolicyCreateReqPolicyStorageClassesEltWithDefaults instantiates a new ObjectStoragePolicyCreateReqPolicyStorageClassesElt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePoolIds

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetActivePoolIds() []int64`

GetActivePoolIds returns the ActivePoolIds field if non-nil, zero value otherwise.

### GetActivePoolIdsOk

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetActivePoolIdsOk() (*[]int64, bool)`

GetActivePoolIdsOk returns a tuple with the ActivePoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePoolIds

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetActivePoolIds(v []int64)`

SetActivePoolIds sets ActivePoolIds field to given value.


### GetClass

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetClass() int32`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetClassOk() (*int32, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetClass(v int32)`

SetClass sets Class field to given value.

### HasClass

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetClassId

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInactivePoolIds

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetInactivePoolIds() []int64`

GetInactivePoolIds returns the InactivePoolIds field if non-nil, zero value otherwise.

### GetInactivePoolIdsOk

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetInactivePoolIdsOk() (*[]int64, bool)`

GetInactivePoolIdsOk returns a tuple with the InactivePoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivePoolIds

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetInactivePoolIds(v []int64)`

SetInactivePoolIds sets InactivePoolIds field to given value.

### HasInactivePoolIds

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) HasInactivePoolIds() bool`

HasInactivePoolIds returns a boolean if a field has been set.

### GetName

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStoragePolicyCreateReqPolicyStorageClassesElt) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


