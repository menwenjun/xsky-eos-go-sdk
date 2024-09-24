# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGroup** | Pointer to [**ClientGroupNestview**](ClientGroupNestview.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewClient

`func NewClient() *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientGroup

`func (o *Client) GetClientGroup() ClientGroupNestview`

GetClientGroup returns the ClientGroup field if non-nil, zero value otherwise.

### GetClientGroupOk

`func (o *Client) GetClientGroupOk() (*ClientGroupNestview, bool)`

GetClientGroupOk returns a tuple with the ClientGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGroup

`func (o *Client) SetClientGroup(v ClientGroupNestview)`

SetClientGroup sets ClientGroup field to given value.

### HasClientGroup

`func (o *Client) HasClientGroup() bool`

HasClientGroup returns a boolean if a field has been set.

### GetCluster

`func (o *Client) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Client) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Client) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Client) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCode

`func (o *Client) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Client) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Client) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Client) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreate

`func (o *Client) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *Client) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *Client) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *Client) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetId

`func (o *Client) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Client) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Client) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Client) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Client) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Client) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Client) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Client) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *Client) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Client) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Client) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Client) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


