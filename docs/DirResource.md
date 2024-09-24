# DirResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dir** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**[]DirResourceExisted**](DirResourceExisted.md) |  | [optional] 

## Methods

### NewDirResource

`func NewDirResource() *DirResource`

NewDirResource instantiates a new DirResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirResourceWithDefaults

`func NewDirResourceWithDefaults() *DirResource`

NewDirResourceWithDefaults instantiates a new DirResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDir

`func (o *DirResource) GetDir() string`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *DirResource) GetDirOk() (*string, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *DirResource) SetDir(v string)`

SetDir sets Dir field to given value.

### HasDir

`func (o *DirResource) HasDir() bool`

HasDir returns a boolean if a field has been set.

### GetResources

`func (o *DirResource) GetResources() []DirResourceExisted`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DirResource) GetResourcesOk() (*[]DirResourceExisted, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DirResource) SetResources(v []DirResourceExisted)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DirResource) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


