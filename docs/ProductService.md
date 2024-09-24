# ProductService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiredTime** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProductService

`func NewProductService() *ProductService`

NewProductService instantiates a new ProductService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductServiceWithDefaults

`func NewProductServiceWithDefaults() *ProductService`

NewProductServiceWithDefaults instantiates a new ProductService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiredTime

`func (o *ProductService) GetExpiredTime() time.Time`

GetExpiredTime returns the ExpiredTime field if non-nil, zero value otherwise.

### GetExpiredTimeOk

`func (o *ProductService) GetExpiredTimeOk() (*time.Time, bool)`

GetExpiredTimeOk returns a tuple with the ExpiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTime

`func (o *ProductService) SetExpiredTime(v time.Time)`

SetExpiredTime sets ExpiredTime field to given value.

### HasExpiredTime

`func (o *ProductService) HasExpiredTime() bool`

HasExpiredTime returns a boolean if a field has been set.

### GetName

`func (o *ProductService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *ProductService) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProductService) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ProductService) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ProductService) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


