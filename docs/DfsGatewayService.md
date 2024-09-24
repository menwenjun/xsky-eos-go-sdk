# DfsGatewayService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsGateway** | Pointer to [**DfsGateway**](DfsGateway.md) |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsGatewayService

`func NewDfsGatewayService() *DfsGatewayService`

NewDfsGatewayService instantiates a new DfsGatewayService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayServiceWithDefaults

`func NewDfsGatewayServiceWithDefaults() *DfsGatewayService`

NewDfsGatewayServiceWithDefaults instantiates a new DfsGatewayService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsGatewayService) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsGatewayService) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsGatewayService) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsGatewayService) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsGatewayService) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsGatewayService) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsGatewayService) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsGatewayService) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsGateway

`func (o *DfsGatewayService) GetDfsGateway() DfsGateway`

GetDfsGateway returns the DfsGateway field if non-nil, zero value otherwise.

### GetDfsGatewayOk

`func (o *DfsGatewayService) GetDfsGatewayOk() (*DfsGateway, bool)`

GetDfsGatewayOk returns a tuple with the DfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateway

`func (o *DfsGatewayService) SetDfsGateway(v DfsGateway)`

SetDfsGateway sets DfsGateway field to given value.

### HasDfsGateway

`func (o *DfsGatewayService) HasDfsGateway() bool`

HasDfsGateway returns a boolean if a field has been set.

### GetHost

`func (o *DfsGatewayService) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DfsGatewayService) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DfsGatewayService) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *DfsGatewayService) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DfsGatewayService) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsGatewayService) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsGatewayService) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsGatewayService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DfsGatewayService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsGatewayService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsGatewayService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsGatewayService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPid

`func (o *DfsGatewayService) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *DfsGatewayService) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *DfsGatewayService) SetPid(v int64)`

SetPid sets Pid field to given value.

### HasPid

`func (o *DfsGatewayService) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetStatus

`func (o *DfsGatewayService) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsGatewayService) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsGatewayService) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsGatewayService) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsGatewayService) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsGatewayService) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsGatewayService) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsGatewayService) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


