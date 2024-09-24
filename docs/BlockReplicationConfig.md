# BlockReplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** |  | [optional] 
**LocalGateways** | Pointer to **[]string** |  | [optional] 
**RemoteGateways** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBlockReplicationConfig

`func NewBlockReplicationConfig() *BlockReplicationConfig`

NewBlockReplicationConfig instantiates a new BlockReplicationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockReplicationConfigWithDefaults

`func NewBlockReplicationConfigWithDefaults() *BlockReplicationConfig`

NewBlockReplicationConfigWithDefaults instantiates a new BlockReplicationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *BlockReplicationConfig) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *BlockReplicationConfig) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *BlockReplicationConfig) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *BlockReplicationConfig) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetLocalGateways

`func (o *BlockReplicationConfig) GetLocalGateways() []string`

GetLocalGateways returns the LocalGateways field if non-nil, zero value otherwise.

### GetLocalGatewaysOk

`func (o *BlockReplicationConfig) GetLocalGatewaysOk() (*[]string, bool)`

GetLocalGatewaysOk returns a tuple with the LocalGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGateways

`func (o *BlockReplicationConfig) SetLocalGateways(v []string)`

SetLocalGateways sets LocalGateways field to given value.

### HasLocalGateways

`func (o *BlockReplicationConfig) HasLocalGateways() bool`

HasLocalGateways returns a boolean if a field has been set.

### GetRemoteGateways

`func (o *BlockReplicationConfig) GetRemoteGateways() []string`

GetRemoteGateways returns the RemoteGateways field if non-nil, zero value otherwise.

### GetRemoteGatewaysOk

`func (o *BlockReplicationConfig) GetRemoteGatewaysOk() (*[]string, bool)`

GetRemoteGatewaysOk returns a tuple with the RemoteGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGateways

`func (o *BlockReplicationConfig) SetRemoteGateways(v []string)`

SetRemoteGateways sets RemoteGateways field to given value.

### HasRemoteGateways

`func (o *BlockReplicationConfig) HasRemoteGateways() bool`

HasRemoteGateways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


