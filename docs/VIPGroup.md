# VIPGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Preempt** | Pointer to **bool** |  | [optional] 
**ResourceId** | Pointer to **int64** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVIPGroup

`func NewVIPGroup() *VIPGroup`

NewVIPGroup instantiates a new VIPGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVIPGroupWithDefaults

`func NewVIPGroupWithDefaults() *VIPGroup`

NewVIPGroupWithDefaults instantiates a new VIPGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *VIPGroup) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *VIPGroup) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *VIPGroup) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *VIPGroup) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCreate

`func (o *VIPGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *VIPGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *VIPGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *VIPGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *VIPGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VIPGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VIPGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VIPGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VIPGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VIPGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VIPGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VIPGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *VIPGroup) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VIPGroup) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VIPGroup) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VIPGroup) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPreempt

`func (o *VIPGroup) GetPreempt() bool`

GetPreempt returns the Preempt field if non-nil, zero value otherwise.

### GetPreemptOk

`func (o *VIPGroup) GetPreemptOk() (*bool, bool)`

GetPreemptOk returns a tuple with the Preempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreempt

`func (o *VIPGroup) SetPreempt(v bool)`

SetPreempt sets Preempt field to given value.

### HasPreempt

`func (o *VIPGroup) HasPreempt() bool`

HasPreempt returns a boolean if a field has been set.

### GetResourceId

`func (o *VIPGroup) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *VIPGroup) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *VIPGroup) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *VIPGroup) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *VIPGroup) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *VIPGroup) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *VIPGroup) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *VIPGroup) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *VIPGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VIPGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VIPGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VIPGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *VIPGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *VIPGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *VIPGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *VIPGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


