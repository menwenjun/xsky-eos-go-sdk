# DfsSMBSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** |  | [optional] 
**ClientPort** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**ConnectedAt** | Pointer to **time.Time** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsGateway** | Pointer to [**DfsGatewayNestview**](DfsGatewayNestview.md) |  | [optional] 
**DfsSmbShare** | Pointer to [**DfsSMBShareNestview**](DfsSMBShareNestview.md) |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ProtocolVersion** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsSMBSession

`func NewDfsSMBSession() *DfsSMBSession`

NewDfsSMBSession instantiates a new DfsSMBSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsSMBSessionWithDefaults

`func NewDfsSMBSessionWithDefaults() *DfsSMBSession`

NewDfsSMBSessionWithDefaults instantiates a new DfsSMBSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *DfsSMBSession) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *DfsSMBSession) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *DfsSMBSession) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *DfsSMBSession) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientPort

`func (o *DfsSMBSession) GetClientPort() int64`

GetClientPort returns the ClientPort field if non-nil, zero value otherwise.

### GetClientPortOk

`func (o *DfsSMBSession) GetClientPortOk() (*int64, bool)`

GetClientPortOk returns a tuple with the ClientPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPort

`func (o *DfsSMBSession) SetClientPort(v int64)`

SetClientPort sets ClientPort field to given value.

### HasClientPort

`func (o *DfsSMBSession) HasClientPort() bool`

HasClientPort returns a boolean if a field has been set.

### GetCluster

`func (o *DfsSMBSession) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsSMBSession) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsSMBSession) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsSMBSession) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnectedAt

`func (o *DfsSMBSession) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *DfsSMBSession) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *DfsSMBSession) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *DfsSMBSession) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetCreate

`func (o *DfsSMBSession) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsSMBSession) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsSMBSession) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsSMBSession) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsGateway

`func (o *DfsSMBSession) GetDfsGateway() DfsGatewayNestview`

GetDfsGateway returns the DfsGateway field if non-nil, zero value otherwise.

### GetDfsGatewayOk

`func (o *DfsSMBSession) GetDfsGatewayOk() (*DfsGatewayNestview, bool)`

GetDfsGatewayOk returns a tuple with the DfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateway

`func (o *DfsSMBSession) SetDfsGateway(v DfsGatewayNestview)`

SetDfsGateway sets DfsGateway field to given value.

### HasDfsGateway

`func (o *DfsSMBSession) HasDfsGateway() bool`

HasDfsGateway returns a boolean if a field has been set.

### GetDfsSmbShare

`func (o *DfsSMBSession) GetDfsSmbShare() DfsSMBShareNestview`

GetDfsSmbShare returns the DfsSmbShare field if non-nil, zero value otherwise.

### GetDfsSmbShareOk

`func (o *DfsSMBSession) GetDfsSmbShareOk() (*DfsSMBShareNestview, bool)`

GetDfsSmbShareOk returns a tuple with the DfsSmbShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsSmbShare

`func (o *DfsSMBSession) SetDfsSmbShare(v DfsSMBShareNestview)`

SetDfsSmbShare sets DfsSmbShare field to given value.

### HasDfsSmbShare

`func (o *DfsSMBSession) HasDfsSmbShare() bool`

HasDfsSmbShare returns a boolean if a field has been set.

### GetGroup

`func (o *DfsSMBSession) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DfsSMBSession) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DfsSMBSession) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DfsSMBSession) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *DfsSMBSession) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsSMBSession) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsSMBSession) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsSMBSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *DfsSMBSession) GetProtocolVersion() string`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *DfsSMBSession) GetProtocolVersionOk() (*string, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *DfsSMBSession) SetProtocolVersion(v string)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *DfsSMBSession) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsSMBSession) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsSMBSession) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsSMBSession) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsSMBSession) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsername

`func (o *DfsSMBSession) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DfsSMBSession) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DfsSMBSession) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DfsSMBSession) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


