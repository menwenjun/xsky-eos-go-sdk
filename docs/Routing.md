# Routing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Lbs** | Pointer to [**[]LoadBalancerInfo**](LoadBalancerInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewRouting

`func NewRouting() *Routing`

NewRouting instantiates a new Routing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingWithDefaults

`func NewRoutingWithDefaults() *Routing`

NewRoutingWithDefaults instantiates a new Routing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *Routing) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *Routing) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *Routing) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *Routing) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetId

`func (o *Routing) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Routing) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Routing) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Routing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLbs

`func (o *Routing) GetLbs() []LoadBalancerInfo`

GetLbs returns the Lbs field if non-nil, zero value otherwise.

### GetLbsOk

`func (o *Routing) GetLbsOk() (*[]LoadBalancerInfo, bool)`

GetLbsOk returns a tuple with the Lbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbs

`func (o *Routing) SetLbs(v []LoadBalancerInfo)`

SetLbs sets Lbs field to given value.

### HasLbs

`func (o *Routing) HasLbs() bool`

HasLbs returns a boolean if a field has been set.

### GetName

`func (o *Routing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Routing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Routing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Routing) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *Routing) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Routing) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Routing) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Routing) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetStatus

`func (o *Routing) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Routing) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Routing) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Routing) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


