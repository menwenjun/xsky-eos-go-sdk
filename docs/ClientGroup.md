# ClientGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPathNum** | Pointer to **int64** |  | [optional] 
**BlockVolumeNum** | Pointer to **int64** |  | [optional] 
**Chap** | Pointer to **bool** |  | [optional] 
**ClientNum** | Pointer to **int64** |  | [optional] 
**Clients** | Pointer to [**[]Client**](Client.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Iname** | Pointer to **string** |  | [optional] 
**Isecret** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewClientGroup

`func NewClientGroup() *ClientGroup`

NewClientGroup instantiates a new ClientGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientGroupWithDefaults

`func NewClientGroupWithDefaults() *ClientGroup`

NewClientGroupWithDefaults instantiates a new ClientGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPathNum

`func (o *ClientGroup) GetAccessPathNum() int64`

GetAccessPathNum returns the AccessPathNum field if non-nil, zero value otherwise.

### GetAccessPathNumOk

`func (o *ClientGroup) GetAccessPathNumOk() (*int64, bool)`

GetAccessPathNumOk returns a tuple with the AccessPathNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPathNum

`func (o *ClientGroup) SetAccessPathNum(v int64)`

SetAccessPathNum sets AccessPathNum field to given value.

### HasAccessPathNum

`func (o *ClientGroup) HasAccessPathNum() bool`

HasAccessPathNum returns a boolean if a field has been set.

### GetBlockVolumeNum

`func (o *ClientGroup) GetBlockVolumeNum() int64`

GetBlockVolumeNum returns the BlockVolumeNum field if non-nil, zero value otherwise.

### GetBlockVolumeNumOk

`func (o *ClientGroup) GetBlockVolumeNumOk() (*int64, bool)`

GetBlockVolumeNumOk returns a tuple with the BlockVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolumeNum

`func (o *ClientGroup) SetBlockVolumeNum(v int64)`

SetBlockVolumeNum sets BlockVolumeNum field to given value.

### HasBlockVolumeNum

`func (o *ClientGroup) HasBlockVolumeNum() bool`

HasBlockVolumeNum returns a boolean if a field has been set.

### GetChap

`func (o *ClientGroup) GetChap() bool`

GetChap returns the Chap field if non-nil, zero value otherwise.

### GetChapOk

`func (o *ClientGroup) GetChapOk() (*bool, bool)`

GetChapOk returns a tuple with the Chap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChap

`func (o *ClientGroup) SetChap(v bool)`

SetChap sets Chap field to given value.

### HasChap

`func (o *ClientGroup) HasChap() bool`

HasChap returns a boolean if a field has been set.

### GetClientNum

`func (o *ClientGroup) GetClientNum() int64`

GetClientNum returns the ClientNum field if non-nil, zero value otherwise.

### GetClientNumOk

`func (o *ClientGroup) GetClientNumOk() (*int64, bool)`

GetClientNumOk returns a tuple with the ClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNum

`func (o *ClientGroup) SetClientNum(v int64)`

SetClientNum sets ClientNum field to given value.

### HasClientNum

`func (o *ClientGroup) HasClientNum() bool`

HasClientNum returns a boolean if a field has been set.

### GetClients

`func (o *ClientGroup) GetClients() []Client`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ClientGroup) GetClientsOk() (*[]Client, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ClientGroup) SetClients(v []Client)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ClientGroup) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetCluster

`func (o *ClientGroup) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClientGroup) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClientGroup) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ClientGroup) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ClientGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ClientGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ClientGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ClientGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *ClientGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ClientGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClientGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIname

`func (o *ClientGroup) GetIname() string`

GetIname returns the Iname field if non-nil, zero value otherwise.

### GetInameOk

`func (o *ClientGroup) GetInameOk() (*string, bool)`

GetInameOk returns a tuple with the Iname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIname

`func (o *ClientGroup) SetIname(v string)`

SetIname sets Iname field to given value.

### HasIname

`func (o *ClientGroup) HasIname() bool`

HasIname returns a boolean if a field has been set.

### GetIsecret

`func (o *ClientGroup) GetIsecret() string`

GetIsecret returns the Isecret field if non-nil, zero value otherwise.

### GetIsecretOk

`func (o *ClientGroup) GetIsecretOk() (*string, bool)`

GetIsecretOk returns a tuple with the Isecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsecret

`func (o *ClientGroup) SetIsecret(v string)`

SetIsecret sets Isecret field to given value.

### HasIsecret

`func (o *ClientGroup) HasIsecret() bool`

HasIsecret returns a boolean if a field has been set.

### GetName

`func (o *ClientGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ClientGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClientGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClientGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClientGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ClientGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *ClientGroup) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ClientGroup) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ClientGroup) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ClientGroup) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUpdate

`func (o *ClientGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ClientGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ClientGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ClientGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


