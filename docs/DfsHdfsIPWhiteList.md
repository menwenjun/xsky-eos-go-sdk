# DfsHdfsIPWhiteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsHdfs** | Pointer to [**DfsHdfsNestview**](DfsHdfsNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ips** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDfsHdfsIPWhiteList

`func NewDfsHdfsIPWhiteList() *DfsHdfsIPWhiteList`

NewDfsHdfsIPWhiteList instantiates a new DfsHdfsIPWhiteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsHdfsIPWhiteListWithDefaults

`func NewDfsHdfsIPWhiteListWithDefaults() *DfsHdfsIPWhiteList`

NewDfsHdfsIPWhiteListWithDefaults instantiates a new DfsHdfsIPWhiteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DfsHdfsIPWhiteList) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsHdfsIPWhiteList) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsHdfsIPWhiteList) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsHdfsIPWhiteList) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsHdfsIPWhiteList) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsHdfsIPWhiteList) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsHdfsIPWhiteList) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsHdfsIPWhiteList) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsHdfs

`func (o *DfsHdfsIPWhiteList) GetDfsHdfs() DfsHdfsNestview`

GetDfsHdfs returns the DfsHdfs field if non-nil, zero value otherwise.

### GetDfsHdfsOk

`func (o *DfsHdfsIPWhiteList) GetDfsHdfsOk() (*DfsHdfsNestview, bool)`

GetDfsHdfsOk returns a tuple with the DfsHdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsHdfs

`func (o *DfsHdfsIPWhiteList) SetDfsHdfs(v DfsHdfsNestview)`

SetDfsHdfs sets DfsHdfs field to given value.

### HasDfsHdfs

`func (o *DfsHdfsIPWhiteList) HasDfsHdfs() bool`

HasDfsHdfs returns a boolean if a field has been set.

### GetId

`func (o *DfsHdfsIPWhiteList) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsHdfsIPWhiteList) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsHdfsIPWhiteList) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsHdfsIPWhiteList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIps

`func (o *DfsHdfsIPWhiteList) GetIps() string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *DfsHdfsIPWhiteList) GetIpsOk() (*string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *DfsHdfsIPWhiteList) SetIps(v string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *DfsHdfsIPWhiteList) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetPermission

`func (o *DfsHdfsIPWhiteList) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *DfsHdfsIPWhiteList) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *DfsHdfsIPWhiteList) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *DfsHdfsIPWhiteList) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetStatus

`func (o *DfsHdfsIPWhiteList) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsHdfsIPWhiteList) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsHdfsIPWhiteList) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsHdfsIPWhiteList) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsHdfsIPWhiteList) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsHdfsIPWhiteList) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsHdfsIPWhiteList) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsHdfsIPWhiteList) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


