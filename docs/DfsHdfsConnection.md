# DfsHdfsConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** |  | [optional] 
**ClientPort** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**ConnectedAt** | Pointer to **time.Time** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsGateway** | Pointer to [**DfsGatewayNestview**](DfsGatewayNestview.md) |  | [optional] 
**DfsGatewayZone** | Pointer to [**DfsGatewayZoneNestview**](DfsGatewayZoneNestview.md) |  | [optional] 
**Hdfs** | Pointer to [**DfsHdfsNestview**](DfsHdfsNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ProtocolVersion** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsHdfsConnection

`func NewDfsHdfsConnection() *DfsHdfsConnection`

NewDfsHdfsConnection instantiates a new DfsHdfsConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsHdfsConnectionWithDefaults

`func NewDfsHdfsConnectionWithDefaults() *DfsHdfsConnection`

NewDfsHdfsConnectionWithDefaults instantiates a new DfsHdfsConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *DfsHdfsConnection) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *DfsHdfsConnection) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *DfsHdfsConnection) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *DfsHdfsConnection) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientPort

`func (o *DfsHdfsConnection) GetClientPort() int64`

GetClientPort returns the ClientPort field if non-nil, zero value otherwise.

### GetClientPortOk

`func (o *DfsHdfsConnection) GetClientPortOk() (*int64, bool)`

GetClientPortOk returns a tuple with the ClientPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPort

`func (o *DfsHdfsConnection) SetClientPort(v int64)`

SetClientPort sets ClientPort field to given value.

### HasClientPort

`func (o *DfsHdfsConnection) HasClientPort() bool`

HasClientPort returns a boolean if a field has been set.

### GetCluster

`func (o *DfsHdfsConnection) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsHdfsConnection) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsHdfsConnection) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsHdfsConnection) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnectedAt

`func (o *DfsHdfsConnection) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *DfsHdfsConnection) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *DfsHdfsConnection) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *DfsHdfsConnection) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetCreate

`func (o *DfsHdfsConnection) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsHdfsConnection) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsHdfsConnection) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsHdfsConnection) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsGateway

`func (o *DfsHdfsConnection) GetDfsGateway() DfsGatewayNestview`

GetDfsGateway returns the DfsGateway field if non-nil, zero value otherwise.

### GetDfsGatewayOk

`func (o *DfsHdfsConnection) GetDfsGatewayOk() (*DfsGatewayNestview, bool)`

GetDfsGatewayOk returns a tuple with the DfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateway

`func (o *DfsHdfsConnection) SetDfsGateway(v DfsGatewayNestview)`

SetDfsGateway sets DfsGateway field to given value.

### HasDfsGateway

`func (o *DfsHdfsConnection) HasDfsGateway() bool`

HasDfsGateway returns a boolean if a field has been set.

### GetDfsGatewayZone

`func (o *DfsHdfsConnection) GetDfsGatewayZone() DfsGatewayZoneNestview`

GetDfsGatewayZone returns the DfsGatewayZone field if non-nil, zero value otherwise.

### GetDfsGatewayZoneOk

`func (o *DfsHdfsConnection) GetDfsGatewayZoneOk() (*DfsGatewayZoneNestview, bool)`

GetDfsGatewayZoneOk returns a tuple with the DfsGatewayZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayZone

`func (o *DfsHdfsConnection) SetDfsGatewayZone(v DfsGatewayZoneNestview)`

SetDfsGatewayZone sets DfsGatewayZone field to given value.

### HasDfsGatewayZone

`func (o *DfsHdfsConnection) HasDfsGatewayZone() bool`

HasDfsGatewayZone returns a boolean if a field has been set.

### GetHdfs

`func (o *DfsHdfsConnection) GetHdfs() DfsHdfsNestview`

GetHdfs returns the Hdfs field if non-nil, zero value otherwise.

### GetHdfsOk

`func (o *DfsHdfsConnection) GetHdfsOk() (*DfsHdfsNestview, bool)`

GetHdfsOk returns a tuple with the Hdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfs

`func (o *DfsHdfsConnection) SetHdfs(v DfsHdfsNestview)`

SetHdfs sets Hdfs field to given value.

### HasHdfs

`func (o *DfsHdfsConnection) HasHdfs() bool`

HasHdfs returns a boolean if a field has been set.

### GetId

`func (o *DfsHdfsConnection) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsHdfsConnection) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsHdfsConnection) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsHdfsConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *DfsHdfsConnection) GetProtocolVersion() string`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *DfsHdfsConnection) GetProtocolVersionOk() (*string, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *DfsHdfsConnection) SetProtocolVersion(v string)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *DfsHdfsConnection) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsHdfsConnection) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsHdfsConnection) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsHdfsConnection) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsHdfsConnection) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsername

`func (o *DfsHdfsConnection) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DfsHdfsConnection) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DfsHdfsConnection) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DfsHdfsConnection) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


