# AlertInfoGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertInfo** | Pointer to [**AlertInfo**](AlertInfo.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ResourceId** | Pointer to **int64** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAlertInfoGroup

`func NewAlertInfoGroup() *AlertInfoGroup`

NewAlertInfoGroup instantiates a new AlertInfoGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertInfoGroupWithDefaults

`func NewAlertInfoGroupWithDefaults() *AlertInfoGroup`

NewAlertInfoGroupWithDefaults instantiates a new AlertInfoGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertInfo

`func (o *AlertInfoGroup) GetAlertInfo() AlertInfo`

GetAlertInfo returns the AlertInfo field if non-nil, zero value otherwise.

### GetAlertInfoOk

`func (o *AlertInfoGroup) GetAlertInfoOk() (*AlertInfo, bool)`

GetAlertInfoOk returns a tuple with the AlertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertInfo

`func (o *AlertInfoGroup) SetAlertInfo(v AlertInfo)`

SetAlertInfo sets AlertInfo field to given value.

### HasAlertInfo

`func (o *AlertInfoGroup) HasAlertInfo() bool`

HasAlertInfo returns a boolean if a field has been set.

### GetCreate

`func (o *AlertInfoGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AlertInfoGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AlertInfoGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AlertInfoGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetGroup

`func (o *AlertInfoGroup) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AlertInfoGroup) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AlertInfoGroup) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AlertInfoGroup) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *AlertInfoGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertInfoGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertInfoGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AlertInfoGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceId

`func (o *AlertInfoGroup) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AlertInfoGroup) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AlertInfoGroup) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AlertInfoGroup) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *AlertInfoGroup) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AlertInfoGroup) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AlertInfoGroup) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AlertInfoGroup) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *AlertInfoGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertInfoGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertInfoGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertInfoGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *AlertInfoGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AlertInfoGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AlertInfoGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AlertInfoGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


