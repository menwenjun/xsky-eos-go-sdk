# RemoteCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**BlockReplicationNum** | Pointer to **int64** |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**FsId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OsZoneNum** | Pointer to **int64** |  | [optional] 
**RemoteClusterId** | Pointer to **int64** |  | [optional] 
**SnapshotReplicationNum** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewRemoteCluster

`func NewRemoteCluster() *RemoteCluster`

NewRemoteCluster instantiates a new RemoteCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteClusterWithDefaults

`func NewRemoteClusterWithDefaults() *RemoteCluster`

NewRemoteClusterWithDefaults instantiates a new RemoteCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *RemoteCluster) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *RemoteCluster) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *RemoteCluster) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *RemoteCluster) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetBlockReplicationNum

`func (o *RemoteCluster) GetBlockReplicationNum() int64`

GetBlockReplicationNum returns the BlockReplicationNum field if non-nil, zero value otherwise.

### GetBlockReplicationNumOk

`func (o *RemoteCluster) GetBlockReplicationNumOk() (*int64, bool)`

GetBlockReplicationNumOk returns a tuple with the BlockReplicationNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReplicationNum

`func (o *RemoteCluster) SetBlockReplicationNum(v int64)`

SetBlockReplicationNum sets BlockReplicationNum field to given value.

### HasBlockReplicationNum

`func (o *RemoteCluster) HasBlockReplicationNum() bool`

HasBlockReplicationNum returns a boolean if a field has been set.

### GetConnected

`func (o *RemoteCluster) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *RemoteCluster) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *RemoteCluster) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *RemoteCluster) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCreate

`func (o *RemoteCluster) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *RemoteCluster) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *RemoteCluster) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *RemoteCluster) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetFsId

`func (o *RemoteCluster) GetFsId() string`

GetFsId returns the FsId field if non-nil, zero value otherwise.

### GetFsIdOk

`func (o *RemoteCluster) GetFsIdOk() (*string, bool)`

GetFsIdOk returns a tuple with the FsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsId

`func (o *RemoteCluster) SetFsId(v string)`

SetFsId sets FsId field to given value.

### HasFsId

`func (o *RemoteCluster) HasFsId() bool`

HasFsId returns a boolean if a field has been set.

### GetId

`func (o *RemoteCluster) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteCluster) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteCluster) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RemoteCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsZoneNum

`func (o *RemoteCluster) GetOsZoneNum() int64`

GetOsZoneNum returns the OsZoneNum field if non-nil, zero value otherwise.

### GetOsZoneNumOk

`func (o *RemoteCluster) GetOsZoneNumOk() (*int64, bool)`

GetOsZoneNumOk returns a tuple with the OsZoneNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsZoneNum

`func (o *RemoteCluster) SetOsZoneNum(v int64)`

SetOsZoneNum sets OsZoneNum field to given value.

### HasOsZoneNum

`func (o *RemoteCluster) HasOsZoneNum() bool`

HasOsZoneNum returns a boolean if a field has been set.

### GetRemoteClusterId

`func (o *RemoteCluster) GetRemoteClusterId() int64`

GetRemoteClusterId returns the RemoteClusterId field if non-nil, zero value otherwise.

### GetRemoteClusterIdOk

`func (o *RemoteCluster) GetRemoteClusterIdOk() (*int64, bool)`

GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterId

`func (o *RemoteCluster) SetRemoteClusterId(v int64)`

SetRemoteClusterId sets RemoteClusterId field to given value.

### HasRemoteClusterId

`func (o *RemoteCluster) HasRemoteClusterId() bool`

HasRemoteClusterId returns a boolean if a field has been set.

### GetSnapshotReplicationNum

`func (o *RemoteCluster) GetSnapshotReplicationNum() int64`

GetSnapshotReplicationNum returns the SnapshotReplicationNum field if non-nil, zero value otherwise.

### GetSnapshotReplicationNumOk

`func (o *RemoteCluster) GetSnapshotReplicationNumOk() (*int64, bool)`

GetSnapshotReplicationNumOk returns a tuple with the SnapshotReplicationNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotReplicationNum

`func (o *RemoteCluster) SetSnapshotReplicationNum(v int64)`

SetSnapshotReplicationNum sets SnapshotReplicationNum field to given value.

### HasSnapshotReplicationNum

`func (o *RemoteCluster) HasSnapshotReplicationNum() bool`

HasSnapshotReplicationNum returns a boolean if a field has been set.

### GetStatus

`func (o *RemoteCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemoteCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *RemoteCluster) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *RemoteCluster) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *RemoteCluster) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *RemoteCluster) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUrl

`func (o *RemoteCluster) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RemoteCluster) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RemoteCluster) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RemoteCluster) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVersion

`func (o *RemoteCluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RemoteCluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RemoteCluster) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RemoteCluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


