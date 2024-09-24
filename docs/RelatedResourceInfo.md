# RelatedResourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Outsides** | Pointer to [**[]OutsideResourceInfo**](OutsideResourceInfo.md) |  | [optional] 
**RelatedResources** | Pointer to **[]string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 

## Methods

### NewRelatedResourceInfo

`func NewRelatedResourceInfo() *RelatedResourceInfo`

NewRelatedResourceInfo instantiates a new RelatedResourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedResourceInfoWithDefaults

`func NewRelatedResourceInfoWithDefaults() *RelatedResourceInfo`

NewRelatedResourceInfoWithDefaults instantiates a new RelatedResourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelatedResourceInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelatedResourceInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelatedResourceInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RelatedResourceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOutsides

`func (o *RelatedResourceInfo) GetOutsides() []OutsideResourceInfo`

GetOutsides returns the Outsides field if non-nil, zero value otherwise.

### GetOutsidesOk

`func (o *RelatedResourceInfo) GetOutsidesOk() (*[]OutsideResourceInfo, bool)`

GetOutsidesOk returns a tuple with the Outsides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsides

`func (o *RelatedResourceInfo) SetOutsides(v []OutsideResourceInfo)`

SetOutsides sets Outsides field to given value.

### HasOutsides

`func (o *RelatedResourceInfo) HasOutsides() bool`

HasOutsides returns a boolean if a field has been set.

### GetRelatedResources

`func (o *RelatedResourceInfo) GetRelatedResources() []string`

GetRelatedResources returns the RelatedResources field if non-nil, zero value otherwise.

### GetRelatedResourcesOk

`func (o *RelatedResourceInfo) GetRelatedResourcesOk() (*[]string, bool)`

GetRelatedResourcesOk returns a tuple with the RelatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResources

`func (o *RelatedResourceInfo) SetRelatedResources(v []string)`

SetRelatedResources sets RelatedResources field to given value.

### HasRelatedResources

`func (o *RelatedResourceInfo) HasRelatedResources() bool`

HasRelatedResources returns a boolean if a field has been set.

### GetResource

`func (o *RelatedResourceInfo) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *RelatedResourceInfo) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *RelatedResourceInfo) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *RelatedResourceInfo) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


