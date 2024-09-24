# ConsistentSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**SnapName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewConsistentSnapshot

`func NewConsistentSnapshot() *ConsistentSnapshot`

NewConsistentSnapshot instantiates a new ConsistentSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsistentSnapshotWithDefaults

`func NewConsistentSnapshotWithDefaults() *ConsistentSnapshot`

NewConsistentSnapshotWithDefaults instantiates a new ConsistentSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ConsistentSnapshot) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ConsistentSnapshot) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ConsistentSnapshot) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ConsistentSnapshot) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ConsistentSnapshot) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ConsistentSnapshot) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ConsistentSnapshot) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ConsistentSnapshot) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *ConsistentSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsistentSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsistentSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsistentSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ConsistentSnapshot) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsistentSnapshot) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsistentSnapshot) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConsistentSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSnapName

`func (o *ConsistentSnapshot) GetSnapName() string`

GetSnapName returns the SnapName field if non-nil, zero value otherwise.

### GetSnapNameOk

`func (o *ConsistentSnapshot) GetSnapNameOk() (*string, bool)`

GetSnapNameOk returns a tuple with the SnapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapName

`func (o *ConsistentSnapshot) SetSnapName(v string)`

SetSnapName sets SnapName field to given value.

### HasSnapName

`func (o *ConsistentSnapshot) HasSnapName() bool`

HasSnapName returns a boolean if a field has been set.

### GetStatus

`func (o *ConsistentSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConsistentSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConsistentSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConsistentSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


