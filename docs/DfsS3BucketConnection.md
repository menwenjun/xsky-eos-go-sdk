# DfsS3BucketConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to [**DfsS3BucketNestview**](DfsS3BucketNestview.md) |  | [optional] 
**ClientIp** | Pointer to **string** |  | [optional] 
**ClientPort** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**ConnectedAt** | Pointer to **time.Time** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsGateway** | Pointer to [**DfsGatewayNestview**](DfsGatewayNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ProtocolVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsS3BucketConnection

`func NewDfsS3BucketConnection() *DfsS3BucketConnection`

NewDfsS3BucketConnection instantiates a new DfsS3BucketConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsS3BucketConnectionWithDefaults

`func NewDfsS3BucketConnectionWithDefaults() *DfsS3BucketConnection`

NewDfsS3BucketConnectionWithDefaults instantiates a new DfsS3BucketConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *DfsS3BucketConnection) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *DfsS3BucketConnection) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *DfsS3BucketConnection) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *DfsS3BucketConnection) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetBucket

`func (o *DfsS3BucketConnection) GetBucket() DfsS3BucketNestview`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *DfsS3BucketConnection) GetBucketOk() (*DfsS3BucketNestview, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *DfsS3BucketConnection) SetBucket(v DfsS3BucketNestview)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *DfsS3BucketConnection) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetClientIp

`func (o *DfsS3BucketConnection) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *DfsS3BucketConnection) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *DfsS3BucketConnection) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *DfsS3BucketConnection) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientPort

`func (o *DfsS3BucketConnection) GetClientPort() int64`

GetClientPort returns the ClientPort field if non-nil, zero value otherwise.

### GetClientPortOk

`func (o *DfsS3BucketConnection) GetClientPortOk() (*int64, bool)`

GetClientPortOk returns a tuple with the ClientPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPort

`func (o *DfsS3BucketConnection) SetClientPort(v int64)`

SetClientPort sets ClientPort field to given value.

### HasClientPort

`func (o *DfsS3BucketConnection) HasClientPort() bool`

HasClientPort returns a boolean if a field has been set.

### GetCluster

`func (o *DfsS3BucketConnection) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsS3BucketConnection) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsS3BucketConnection) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsS3BucketConnection) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConnectedAt

`func (o *DfsS3BucketConnection) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *DfsS3BucketConnection) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *DfsS3BucketConnection) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *DfsS3BucketConnection) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetCreate

`func (o *DfsS3BucketConnection) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsS3BucketConnection) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsS3BucketConnection) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsS3BucketConnection) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsGateway

`func (o *DfsS3BucketConnection) GetDfsGateway() DfsGatewayNestview`

GetDfsGateway returns the DfsGateway field if non-nil, zero value otherwise.

### GetDfsGatewayOk

`func (o *DfsS3BucketConnection) GetDfsGatewayOk() (*DfsGatewayNestview, bool)`

GetDfsGatewayOk returns a tuple with the DfsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGateway

`func (o *DfsS3BucketConnection) SetDfsGateway(v DfsGatewayNestview)`

SetDfsGateway sets DfsGateway field to given value.

### HasDfsGateway

`func (o *DfsS3BucketConnection) HasDfsGateway() bool`

HasDfsGateway returns a boolean if a field has been set.

### GetId

`func (o *DfsS3BucketConnection) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsS3BucketConnection) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsS3BucketConnection) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsS3BucketConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *DfsS3BucketConnection) GetProtocolVersion() string`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *DfsS3BucketConnection) GetProtocolVersionOk() (*string, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *DfsS3BucketConnection) SetProtocolVersion(v string)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *DfsS3BucketConnection) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *DfsS3BucketConnection) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DfsS3BucketConnection) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DfsS3BucketConnection) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DfsS3BucketConnection) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsS3BucketConnection) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsS3BucketConnection) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsS3BucketConnection) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsS3BucketConnection) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsername

`func (o *DfsS3BucketConnection) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DfsS3BucketConnection) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DfsS3BucketConnection) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DfsS3BucketConnection) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


