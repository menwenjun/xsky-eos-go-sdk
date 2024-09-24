# ClientGroupUpdateReqClientGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to [**[]ClientGroupUpdateReqClientGroupClientsElt**](ClientGroupUpdateReqClientGroupClientsElt.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewClientGroupUpdateReqClientGroup

`func NewClientGroupUpdateReqClientGroup() *ClientGroupUpdateReqClientGroup`

NewClientGroupUpdateReqClientGroup instantiates a new ClientGroupUpdateReqClientGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientGroupUpdateReqClientGroupWithDefaults

`func NewClientGroupUpdateReqClientGroupWithDefaults() *ClientGroupUpdateReqClientGroup`

NewClientGroupUpdateReqClientGroupWithDefaults instantiates a new ClientGroupUpdateReqClientGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *ClientGroupUpdateReqClientGroup) GetClients() []ClientGroupUpdateReqClientGroupClientsElt`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ClientGroupUpdateReqClientGroup) GetClientsOk() (*[]ClientGroupUpdateReqClientGroupClientsElt, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ClientGroupUpdateReqClientGroup) SetClients(v []ClientGroupUpdateReqClientGroupClientsElt)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ClientGroupUpdateReqClientGroup) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetDescription

`func (o *ClientGroupUpdateReqClientGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientGroupUpdateReqClientGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientGroupUpdateReqClientGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientGroupUpdateReqClientGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ClientGroupUpdateReqClientGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientGroupUpdateReqClientGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientGroupUpdateReqClientGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientGroupUpdateReqClientGroup) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


