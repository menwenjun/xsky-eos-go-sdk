# DfsFTPSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** |  | [optional] 
**ClientPort** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**ConnectedAt** | Pointer to **time.Time** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsFtpShare** | Pointer to [**DfsFTPShareNestview**](DfsFTPShareNestview.md) |  | [optional] 
**DfsGateway** | Pointer to [**DfsGatewayNestview**](DfsGatewayNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsFTPSession

`func NewDfsFTPSession() *DfsFTPSession`

NewDfsFTPSession instantiates a new DfsFTPSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsFTPSessionWithDefaults

`func NewDfsFTPSessionWithDefaults() *DfsFTPSession`

NewDfsFTPSessionWithDefaults instantiates a new DfsFTPSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *DfsFTPSession) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *DfsFTPSession) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *DfsFTPSession) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *DfsFTPSession) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientPort

`func (o *DfsFTPSession) GetClientPort() int64`

GetClientPort returns the ClientPort field if non-nil, zero value otherwise.

### GetClientPortOk

`func (o *DfsFTPSession) GetClientPortOk() (*int64, bool)`

GetClientPortOk returns a tuple with the ClientPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPort

`func (o *DfsFTPSession) SetClientPort(v int64)`

SetClientPort sets ClientPort field to given value.

### HasClientPort

`func (o *DfsFTPSession) HasClientPort() bool`

HasClientPort returns a boolean if a field has been set.

### GetCluster

`func (o *DfsFTPSession) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsFTPSession) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsFTPSession) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsFTPSession) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnectedAt

`func (o *DfsFTPSession) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *DfsFTPSession) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *DfsFTPSession) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *DfsFTPSession) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetCreate

`func (o *DfsFTPSession) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsFTPSession) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsFTPSession) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsFTPSession) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsFtpShare

`func (o *DfsFTPSession) GetDfsFtpShare() DfsFTPShareNestview`

GetDfsFtpShare returns the DfsFtpShare field if non-nil, zero value otherwise.

### GetDfsFtpShareOk

`func (o *DfsFTPSession) GetDfsFtpShareOk() (*DfsFTPShareNestview, bool)`

GetDfsFtpShareOk returns a tuple with the DfsFtpShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsFtpShare

`func (o *DfsFTPSession) SetDfsFtpShare(v DfsFTPShareNestview)`

SetDfsFtpShare sets DfsFtpShare field to given value.

### HasDfsFtpShare

`func (o *DfsFTPSession) HasDfsFtpShare() bool`

HasDfsFtpShare returns a boolean if a field has been set.

### GetDfsGateway

`func (o *DfsFTPSession) GetDfsGateway() DfsGatewayNestview`

GetDfsGateway returns the DfsGateway field if non-nil, zero value otherwise.

### GetDfsGatewayOk

`func (o *DfsFTPSession) GetDfsGatewayOk() (*DfsGatewayNestview, bool)`

GetDfsGatewayOk returns a tuple with the DfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateway

`func (o *DfsFTPSession) SetDfsGateway(v DfsGatewayNestview)`

SetDfsGateway sets DfsGateway field to given value.

### HasDfsGateway

`func (o *DfsFTPSession) HasDfsGateway() bool`

HasDfsGateway returns a boolean if a field has been set.

### GetId

`func (o *DfsFTPSession) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsFTPSession) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsFTPSession) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsFTPSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsFTPSession) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsFTPSession) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsFTPSession) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsFTPSession) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsername

`func (o *DfsFTPSession) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DfsFTPSession) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DfsFTPSession) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DfsFTPSession) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


