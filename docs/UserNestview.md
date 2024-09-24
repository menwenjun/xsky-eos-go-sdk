# UserNestview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | id of user | [optional] 
**Name** | Pointer to **string** | name of user | [optional] 
**SkipGuide** | Pointer to **bool** | skip Guide | [optional] 

## Methods

### NewUserNestview

`func NewUserNestview() *UserNestview`

NewUserNestview instantiates a new UserNestview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNestviewWithDefaults

`func NewUserNestviewWithDefaults() *UserNestview`

NewUserNestviewWithDefaults instantiates a new UserNestview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserNestview) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserNestview) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserNestview) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserNestview) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserNestview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserNestview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserNestview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserNestview) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSkipGuide

`func (o *UserNestview) GetSkipGuide() bool`

GetSkipGuide returns the SkipGuide field if non-nil, zero value otherwise.

### GetSkipGuideOk

`func (o *UserNestview) GetSkipGuideOk() (*bool, bool)`

GetSkipGuideOk returns a tuple with the SkipGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipGuide

`func (o *UserNestview) SetSkipGuide(v bool)`

SetSkipGuide sets SkipGuide field to given value.

### HasSkipGuide

`func (o *UserNestview) HasSkipGuide() bool`

HasSkipGuide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


