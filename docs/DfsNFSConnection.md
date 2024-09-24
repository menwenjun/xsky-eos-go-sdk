# DfsNFSConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsGateway** | Pointer to [**DfsGatewayNestview**](DfsGatewayNestview.md) |  | [optional] 
**DfsNfsShare** | Pointer to [**DfsNFSShareNestview**](DfsNFSShareNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**NfsVersion** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsNFSConnection

`func NewDfsNFSConnection() *DfsNFSConnection`

NewDfsNFSConnection instantiates a new DfsNFSConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsNFSConnectionWithDefaults

`func NewDfsNFSConnectionWithDefaults() *DfsNFSConnection`

NewDfsNFSConnectionWithDefaults instantiates a new DfsNFSConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *DfsNFSConnection) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *DfsNFSConnection) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *DfsNFSConnection) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *DfsNFSConnection) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetCluster

`func (o *DfsNFSConnection) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsNFSConnection) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsNFSConnection) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsNFSConnection) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsNFSConnection) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsNFSConnection) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsNFSConnection) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsNFSConnection) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsGateway

`func (o *DfsNFSConnection) GetDfsGateway() DfsGatewayNestview`

GetDfsGateway returns the DfsGateway field if non-nil, zero value otherwise.

### GetDfsGatewayOk

`func (o *DfsNFSConnection) GetDfsGatewayOk() (*DfsGatewayNestview, bool)`

GetDfsGatewayOk returns a tuple with the DfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateway

`func (o *DfsNFSConnection) SetDfsGateway(v DfsGatewayNestview)`

SetDfsGateway sets DfsGateway field to given value.

### HasDfsGateway

`func (o *DfsNFSConnection) HasDfsGateway() bool`

HasDfsGateway returns a boolean if a field has been set.

### GetDfsNfsShare

`func (o *DfsNFSConnection) GetDfsNfsShare() DfsNFSShareNestview`

GetDfsNfsShare returns the DfsNfsShare field if non-nil, zero value otherwise.

### GetDfsNfsShareOk

`func (o *DfsNFSConnection) GetDfsNfsShareOk() (*DfsNFSShareNestview, bool)`

GetDfsNfsShareOk returns a tuple with the DfsNfsShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsNfsShare

`func (o *DfsNFSConnection) SetDfsNfsShare(v DfsNFSShareNestview)`

SetDfsNfsShare sets DfsNfsShare field to given value.

### HasDfsNfsShare

`func (o *DfsNFSConnection) HasDfsNfsShare() bool`

HasDfsNfsShare returns a boolean if a field has been set.

### GetId

`func (o *DfsNFSConnection) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsNFSConnection) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsNFSConnection) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsNFSConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNfsVersion

`func (o *DfsNFSConnection) GetNfsVersion() string`

GetNfsVersion returns the NfsVersion field if non-nil, zero value otherwise.

### GetNfsVersionOk

`func (o *DfsNFSConnection) GetNfsVersionOk() (*string, bool)`

GetNfsVersionOk returns a tuple with the NfsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsVersion

`func (o *DfsNFSConnection) SetNfsVersion(v string)`

SetNfsVersion sets NfsVersion field to given value.

### HasNfsVersion

`func (o *DfsNFSConnection) HasNfsVersion() bool`

HasNfsVersion returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsNFSConnection) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsNFSConnection) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsNFSConnection) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsNFSConnection) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


