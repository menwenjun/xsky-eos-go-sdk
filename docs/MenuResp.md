# MenuResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Menu** | **string** | JSONTextField defines field represented as JSON in API, while stored as string in db | 

## Methods

### NewMenuResp

`func NewMenuResp(menu string, ) *MenuResp`

NewMenuResp instantiates a new MenuResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuRespWithDefaults

`func NewMenuRespWithDefaults() *MenuResp`

NewMenuRespWithDefaults instantiates a new MenuResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMenu

`func (o *MenuResp) GetMenu() string`

GetMenu returns the Menu field if non-nil, zero value otherwise.

### GetMenuOk

`func (o *MenuResp) GetMenuOk() (*string, bool)`

GetMenuOk returns a tuple with the Menu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenu

`func (o *MenuResp) SetMenu(v string)`

SetMenu sets Menu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


