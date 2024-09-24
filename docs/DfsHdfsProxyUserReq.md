# DfsHdfsProxyUserReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | host node that allows access to hdfs through proxy | [optional] 
**ProxyUserName** | Pointer to **string** | proxy user name of hdfs | [optional] 
**User** | Pointer to **string** | users allowed to proxy | [optional] 
**UserGroup** | Pointer to **string** | users group allowed to proxy | [optional] 

## Methods

### NewDfsHdfsProxyUserReq

`func NewDfsHdfsProxyUserReq() *DfsHdfsProxyUserReq`

NewDfsHdfsProxyUserReq instantiates a new DfsHdfsProxyUserReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsHdfsProxyUserReqWithDefaults

`func NewDfsHdfsProxyUserReqWithDefaults() *DfsHdfsProxyUserReq`

NewDfsHdfsProxyUserReqWithDefaults instantiates a new DfsHdfsProxyUserReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *DfsHdfsProxyUserReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DfsHdfsProxyUserReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DfsHdfsProxyUserReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DfsHdfsProxyUserReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetProxyUserName

`func (o *DfsHdfsProxyUserReq) GetProxyUserName() string`

GetProxyUserName returns the ProxyUserName field if non-nil, zero value otherwise.

### GetProxyUserNameOk

`func (o *DfsHdfsProxyUserReq) GetProxyUserNameOk() (*string, bool)`

GetProxyUserNameOk returns a tuple with the ProxyUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUserName

`func (o *DfsHdfsProxyUserReq) SetProxyUserName(v string)`

SetProxyUserName sets ProxyUserName field to given value.

### HasProxyUserName

`func (o *DfsHdfsProxyUserReq) HasProxyUserName() bool`

HasProxyUserName returns a boolean if a field has been set.

### GetUser

`func (o *DfsHdfsProxyUserReq) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DfsHdfsProxyUserReq) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DfsHdfsProxyUserReq) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DfsHdfsProxyUserReq) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserGroup

`func (o *DfsHdfsProxyUserReq) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *DfsHdfsProxyUserReq) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *DfsHdfsProxyUserReq) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *DfsHdfsProxyUserReq) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


