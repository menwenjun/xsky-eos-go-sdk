# DfsGatewayServiceRecord

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
**Samples** | Pointer to [**[]DfsGatewayServiceStat**](DfsGatewayServiceStat.md) |  | [optional] 

## Methods

### NewDfsGatewayServiceRecord

`func NewDfsGatewayServiceRecord() *DfsGatewayServiceRecord`

NewDfsGatewayServiceRecord instantiates a new DfsGatewayServiceRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsGatewayServiceRecordWithDefaults

`func NewDfsGatewayServiceRecordWithDefaults() *DfsGatewayServiceRecord`

NewDfsGatewayServiceRecordWithDefaults instantiates a new DfsGatewayServiceRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsGatewayServiceRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsGatewayServiceRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsGatewayServiceRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsGatewayServiceRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsGatewayServiceRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsGatewayServiceRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsGatewayServiceRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsGatewayServiceRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsGateway

`func (o *DfsGatewayServiceRecord) GetDfsGateway() DfsGateway`

GetDfsGateway returns the DfsGateway field if non-nil, zero value otherwise.

### GetDfsGatewayOk

`func (o *DfsGatewayServiceRecord) GetDfsGatewayOk() (*DfsGateway, bool)`

GetDfsGatewayOk returns a tuple with the DfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateway

`func (o *DfsGatewayServiceRecord) SetDfsGateway(v DfsGateway)`

SetDfsGateway sets DfsGateway field to given value.

### HasDfsGateway

`func (o *DfsGatewayServiceRecord) HasDfsGateway() bool`

HasDfsGateway returns a boolean if a field has been set.

### GetHost

`func (o *DfsGatewayServiceRecord) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DfsGatewayServiceRecord) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DfsGatewayServiceRecord) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *DfsGatewayServiceRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DfsGatewayServiceRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsGatewayServiceRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsGatewayServiceRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsGatewayServiceRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DfsGatewayServiceRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsGatewayServiceRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsGatewayServiceRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsGatewayServiceRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPid

`func (o *DfsGatewayServiceRecord) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *DfsGatewayServiceRecord) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *DfsGatewayServiceRecord) SetPid(v int64)`

SetPid sets Pid field to given value.

### HasPid

`func (o *DfsGatewayServiceRecord) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetStatus

`func (o *DfsGatewayServiceRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsGatewayServiceRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsGatewayServiceRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsGatewayServiceRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsGatewayServiceRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsGatewayServiceRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsGatewayServiceRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsGatewayServiceRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *DfsGatewayServiceRecord) GetSamples() []DfsGatewayServiceStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DfsGatewayServiceRecord) GetSamplesOk() (*[]DfsGatewayServiceStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DfsGatewayServiceRecord) SetSamples(v []DfsGatewayServiceStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *DfsGatewayServiceRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


