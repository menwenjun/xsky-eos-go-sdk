# FSActiveDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DfsGatewayGroupNum** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**MaxUid** | Pointer to **int64** |  | [optional] 
**MinUid** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Realm** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**VerifyTime** | Pointer to **time.Time** |  | [optional] 
**Workgroup** | Pointer to **string** |  | [optional] 

## Methods

### NewFSActiveDirectory

`func NewFSActiveDirectory() *FSActiveDirectory`

NewFSActiveDirectory instantiates a new FSActiveDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSActiveDirectoryWithDefaults

`func NewFSActiveDirectoryWithDefaults() *FSActiveDirectory`

NewFSActiveDirectoryWithDefaults instantiates a new FSActiveDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *FSActiveDirectory) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *FSActiveDirectory) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *FSActiveDirectory) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *FSActiveDirectory) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *FSActiveDirectory) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *FSActiveDirectory) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *FSActiveDirectory) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *FSActiveDirectory) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *FSActiveDirectory) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FSActiveDirectory) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FSActiveDirectory) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FSActiveDirectory) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDfsGatewayGroupNum

`func (o *FSActiveDirectory) GetDfsGatewayGroupNum() int64`

GetDfsGatewayGroupNum returns the DfsGatewayGroupNum field if non-nil, zero value otherwise.

### GetDfsGatewayGroupNumOk

`func (o *FSActiveDirectory) GetDfsGatewayGroupNumOk() (*int64, bool)`

GetDfsGatewayGroupNumOk returns a tuple with the DfsGatewayGroupNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroupNum

`func (o *FSActiveDirectory) SetDfsGatewayGroupNum(v int64)`

SetDfsGatewayGroupNum sets DfsGatewayGroupNum field to given value.

### HasDfsGatewayGroupNum

`func (o *FSActiveDirectory) HasDfsGatewayGroupNum() bool`

HasDfsGatewayGroupNum returns a boolean if a field has been set.

### GetId

`func (o *FSActiveDirectory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FSActiveDirectory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FSActiveDirectory) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FSActiveDirectory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *FSActiveDirectory) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *FSActiveDirectory) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *FSActiveDirectory) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *FSActiveDirectory) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMaxUid

`func (o *FSActiveDirectory) GetMaxUid() int64`

GetMaxUid returns the MaxUid field if non-nil, zero value otherwise.

### GetMaxUidOk

`func (o *FSActiveDirectory) GetMaxUidOk() (*int64, bool)`

GetMaxUidOk returns a tuple with the MaxUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUid

`func (o *FSActiveDirectory) SetMaxUid(v int64)`

SetMaxUid sets MaxUid field to given value.

### HasMaxUid

`func (o *FSActiveDirectory) HasMaxUid() bool`

HasMaxUid returns a boolean if a field has been set.

### GetMinUid

`func (o *FSActiveDirectory) GetMinUid() int64`

GetMinUid returns the MinUid field if non-nil, zero value otherwise.

### GetMinUidOk

`func (o *FSActiveDirectory) GetMinUidOk() (*int64, bool)`

GetMinUidOk returns a tuple with the MinUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUid

`func (o *FSActiveDirectory) SetMinUid(v int64)`

SetMinUid sets MinUid field to given value.

### HasMinUid

`func (o *FSActiveDirectory) HasMinUid() bool`

HasMinUid returns a boolean if a field has been set.

### GetName

`func (o *FSActiveDirectory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FSActiveDirectory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FSActiveDirectory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FSActiveDirectory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRealm

`func (o *FSActiveDirectory) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *FSActiveDirectory) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *FSActiveDirectory) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *FSActiveDirectory) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetStatus

`func (o *FSActiveDirectory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FSActiveDirectory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FSActiveDirectory) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FSActiveDirectory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *FSActiveDirectory) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *FSActiveDirectory) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *FSActiveDirectory) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *FSActiveDirectory) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUsername

`func (o *FSActiveDirectory) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FSActiveDirectory) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FSActiveDirectory) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FSActiveDirectory) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVerifyTime

`func (o *FSActiveDirectory) GetVerifyTime() time.Time`

GetVerifyTime returns the VerifyTime field if non-nil, zero value otherwise.

### GetVerifyTimeOk

`func (o *FSActiveDirectory) GetVerifyTimeOk() (*time.Time, bool)`

GetVerifyTimeOk returns a tuple with the VerifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTime

`func (o *FSActiveDirectory) SetVerifyTime(v time.Time)`

SetVerifyTime sets VerifyTime field to given value.

### HasVerifyTime

`func (o *FSActiveDirectory) HasVerifyTime() bool`

HasVerifyTime returns a boolean if a field has been set.

### GetWorkgroup

`func (o *FSActiveDirectory) GetWorkgroup() string`

GetWorkgroup returns the Workgroup field if non-nil, zero value otherwise.

### GetWorkgroupOk

`func (o *FSActiveDirectory) GetWorkgroupOk() (*string, bool)`

GetWorkgroupOk returns a tuple with the Workgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkgroup

`func (o *FSActiveDirectory) SetWorkgroup(v string)`

SetWorkgroup sets Workgroup field to given value.

### HasWorkgroup

`func (o *FSActiveDirectory) HasWorkgroup() bool`

HasWorkgroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


