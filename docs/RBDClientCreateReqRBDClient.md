# RBDClientCreateReqRBDClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminIp** | **string** | admin ip | 
**Port** | Pointer to **int64** | rbdc monitor listen port | [optional] 
**PublicIp** | Pointer to **string** | public ip | [optional] 
**PublicNetwork** | Pointer to **string** | public network | [optional] 
**Token** | Pointer to **string** | token to access rbd client | [optional] 
**Type** | Pointer to **string** | rbd client type (managed or standalone) | [optional] 

## Methods

### NewRBDClientCreateReqRBDClient

`func NewRBDClientCreateReqRBDClient(adminIp string, ) *RBDClientCreateReqRBDClient`

NewRBDClientCreateReqRBDClient instantiates a new RBDClientCreateReqRBDClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRBDClientCreateReqRBDClientWithDefaults

`func NewRBDClientCreateReqRBDClientWithDefaults() *RBDClientCreateReqRBDClient`

NewRBDClientCreateReqRBDClientWithDefaults instantiates a new RBDClientCreateReqRBDClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminIp

`func (o *RBDClientCreateReqRBDClient) GetAdminIp() string`

GetAdminIp returns the AdminIp field if non-nil, zero value otherwise.

### GetAdminIpOk

`func (o *RBDClientCreateReqRBDClient) GetAdminIpOk() (*string, bool)`

GetAdminIpOk returns a tuple with the AdminIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIp

`func (o *RBDClientCreateReqRBDClient) SetAdminIp(v string)`

SetAdminIp sets AdminIp field to given value.


### GetPort

`func (o *RBDClientCreateReqRBDClient) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RBDClientCreateReqRBDClient) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RBDClientCreateReqRBDClient) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *RBDClientCreateReqRBDClient) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPublicIp

`func (o *RBDClientCreateReqRBDClient) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *RBDClientCreateReqRBDClient) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *RBDClientCreateReqRBDClient) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *RBDClientCreateReqRBDClient) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPublicNetwork

`func (o *RBDClientCreateReqRBDClient) GetPublicNetwork() string`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *RBDClientCreateReqRBDClient) GetPublicNetworkOk() (*string, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *RBDClientCreateReqRBDClient) SetPublicNetwork(v string)`

SetPublicNetwork sets PublicNetwork field to given value.

### HasPublicNetwork

`func (o *RBDClientCreateReqRBDClient) HasPublicNetwork() bool`

HasPublicNetwork returns a boolean if a field has been set.

### GetToken

`func (o *RBDClientCreateReqRBDClient) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RBDClientCreateReqRBDClient) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RBDClientCreateReqRBDClient) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RBDClientCreateReqRBDClient) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *RBDClientCreateReqRBDClient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RBDClientCreateReqRBDClient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RBDClientCreateReqRBDClient) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RBDClientCreateReqRBDClient) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


