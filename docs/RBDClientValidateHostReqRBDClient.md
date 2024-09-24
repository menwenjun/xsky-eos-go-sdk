# RBDClientValidateHostReqRBDClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminIp** | **string** | admin ip | 
**Port** | Pointer to **int64** | rbdc monitor listen port | [optional] 
**Token** | Pointer to **string** | storage server or storage client | [optional] 

## Methods

### NewRBDClientValidateHostReqRBDClient

`func NewRBDClientValidateHostReqRBDClient(adminIp string, ) *RBDClientValidateHostReqRBDClient`

NewRBDClientValidateHostReqRBDClient instantiates a new RBDClientValidateHostReqRBDClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRBDClientValidateHostReqRBDClientWithDefaults

`func NewRBDClientValidateHostReqRBDClientWithDefaults() *RBDClientValidateHostReqRBDClient`

NewRBDClientValidateHostReqRBDClientWithDefaults instantiates a new RBDClientValidateHostReqRBDClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminIp

`func (o *RBDClientValidateHostReqRBDClient) GetAdminIp() string`

GetAdminIp returns the AdminIp field if non-nil, zero value otherwise.

### GetAdminIpOk

`func (o *RBDClientValidateHostReqRBDClient) GetAdminIpOk() (*string, bool)`

GetAdminIpOk returns a tuple with the AdminIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIp

`func (o *RBDClientValidateHostReqRBDClient) SetAdminIp(v string)`

SetAdminIp sets AdminIp field to given value.


### GetPort

`func (o *RBDClientValidateHostReqRBDClient) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RBDClientValidateHostReqRBDClient) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RBDClientValidateHostReqRBDClient) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *RBDClientValidateHostReqRBDClient) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetToken

`func (o *RBDClientValidateHostReqRBDClient) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RBDClientValidateHostReqRBDClient) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RBDClientValidateHostReqRBDClient) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RBDClientValidateHostReqRBDClient) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


