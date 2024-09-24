# DfsHdfsProxyUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsHdfs** | Pointer to [**DfsHdfsNestview**](DfsHdfsNestview.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ProxyUserName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**UserGroup** | Pointer to **string** |  | [optional] 

## Methods

### NewDfsHdfsProxyUser

`func NewDfsHdfsProxyUser() *DfsHdfsProxyUser`

NewDfsHdfsProxyUser instantiates a new DfsHdfsProxyUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsHdfsProxyUserWithDefaults

`func NewDfsHdfsProxyUserWithDefaults() *DfsHdfsProxyUser`

NewDfsHdfsProxyUserWithDefaults instantiates a new DfsHdfsProxyUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsHdfsProxyUser) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsHdfsProxyUser) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsHdfsProxyUser) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsHdfsProxyUser) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsHdfsProxyUser) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsHdfsProxyUser) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsHdfsProxyUser) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsHdfsProxyUser) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsHdfs

`func (o *DfsHdfsProxyUser) GetDfsHdfs() DfsHdfsNestview`

GetDfsHdfs returns the DfsHdfs field if non-nil, zero value otherwise.

### GetDfsHdfsOk

`func (o *DfsHdfsProxyUser) GetDfsHdfsOk() (*DfsHdfsNestview, bool)`

GetDfsHdfsOk returns a tuple with the DfsHdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfs

`func (o *DfsHdfsProxyUser) SetDfsHdfs(v DfsHdfsNestview)`

SetDfsHdfs sets DfsHdfs field to given value.

### HasDfsHdfs

`func (o *DfsHdfsProxyUser) HasDfsHdfs() bool`

HasDfsHdfs returns a boolean if a field has been set.

### GetHost

`func (o *DfsHdfsProxyUser) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DfsHdfsProxyUser) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DfsHdfsProxyUser) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DfsHdfsProxyUser) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DfsHdfsProxyUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsHdfsProxyUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsHdfsProxyUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsHdfsProxyUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProxyUserName

`func (o *DfsHdfsProxyUser) GetProxyUserName() string`

GetProxyUserName returns the ProxyUserName field if non-nil, zero value otherwise.

### GetProxyUserNameOk

`func (o *DfsHdfsProxyUser) GetProxyUserNameOk() (*string, bool)`

GetProxyUserNameOk returns a tuple with the ProxyUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUserName

`func (o *DfsHdfsProxyUser) SetProxyUserName(v string)`

SetProxyUserName sets ProxyUserName field to given value.

### HasProxyUserName

`func (o *DfsHdfsProxyUser) HasProxyUserName() bool`

HasProxyUserName returns a boolean if a field has been set.

### GetStatus

`func (o *DfsHdfsProxyUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsHdfsProxyUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsHdfsProxyUser) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsHdfsProxyUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsHdfsProxyUser) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsHdfsProxyUser) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsHdfsProxyUser) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsHdfsProxyUser) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUser

`func (o *DfsHdfsProxyUser) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DfsHdfsProxyUser) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DfsHdfsProxyUser) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DfsHdfsProxyUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserGroup

`func (o *DfsHdfsProxyUser) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *DfsHdfsProxyUser) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *DfsHdfsProxyUser) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *DfsHdfsProxyUser) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


