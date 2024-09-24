# FSUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FsUsers** | Pointer to [**[]FSUserNestview**](FSUserNestview.md) | Not reload default, because it will cause a lot of memory in action log | [optional] 
**GroupId** | Pointer to **int64** |  | [optional] 
**HdfsNum** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PolicyFilterNum** | Pointer to **int64** |  | [optional] 
**QuotaNum** | Pointer to **int64** |  | [optional] 
**ShareNums** | Pointer to **map[string]int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**UserNum** | Pointer to **int64** |  | [optional] 

## Methods

### NewFSUserGroup

`func NewFSUserGroup() *FSUserGroup`

NewFSUserGroup instantiates a new FSUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSUserGroupWithDefaults

`func NewFSUserGroupWithDefaults() *FSUserGroup`

NewFSUserGroupWithDefaults instantiates a new FSUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *FSUserGroup) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *FSUserGroup) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *FSUserGroup) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *FSUserGroup) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *FSUserGroup) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *FSUserGroup) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *FSUserGroup) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *FSUserGroup) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *FSUserGroup) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FSUserGroup) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FSUserGroup) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FSUserGroup) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *FSUserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FSUserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FSUserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FSUserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFsUsers

`func (o *FSUserGroup) GetFsUsers() []FSUserNestview`

GetFsUsers returns the FsUsers field if non-nil, zero value otherwise.

### GetFsUsersOk

`func (o *FSUserGroup) GetFsUsersOk() (*[]FSUserNestview, bool)`

GetFsUsersOk returns a tuple with the FsUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUsers

`func (o *FSUserGroup) SetFsUsers(v []FSUserNestview)`

SetFsUsers sets FsUsers field to given value.

### HasFsUsers

`func (o *FSUserGroup) HasFsUsers() bool`

HasFsUsers returns a boolean if a field has been set.

### GetGroupId

`func (o *FSUserGroup) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FSUserGroup) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FSUserGroup) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *FSUserGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHdfsNum

`func (o *FSUserGroup) GetHdfsNum() int64`

GetHdfsNum returns the HdfsNum field if non-nil, zero value otherwise.

### GetHdfsNumOk

`func (o *FSUserGroup) GetHdfsNumOk() (*int64, bool)`

GetHdfsNumOk returns a tuple with the HdfsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsNum

`func (o *FSUserGroup) SetHdfsNum(v int64)`

SetHdfsNum sets HdfsNum field to given value.

### HasHdfsNum

`func (o *FSUserGroup) HasHdfsNum() bool`

HasHdfsNum returns a boolean if a field has been set.

### GetId

`func (o *FSUserGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FSUserGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FSUserGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FSUserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FSUserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FSUserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FSUserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FSUserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyFilterNum

`func (o *FSUserGroup) GetPolicyFilterNum() int64`

GetPolicyFilterNum returns the PolicyFilterNum field if non-nil, zero value otherwise.

### GetPolicyFilterNumOk

`func (o *FSUserGroup) GetPolicyFilterNumOk() (*int64, bool)`

GetPolicyFilterNumOk returns a tuple with the PolicyFilterNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyFilterNum

`func (o *FSUserGroup) SetPolicyFilterNum(v int64)`

SetPolicyFilterNum sets PolicyFilterNum field to given value.

### HasPolicyFilterNum

`func (o *FSUserGroup) HasPolicyFilterNum() bool`

HasPolicyFilterNum returns a boolean if a field has been set.

### GetQuotaNum

`func (o *FSUserGroup) GetQuotaNum() int64`

GetQuotaNum returns the QuotaNum field if non-nil, zero value otherwise.

### GetQuotaNumOk

`func (o *FSUserGroup) GetQuotaNumOk() (*int64, bool)`

GetQuotaNumOk returns a tuple with the QuotaNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaNum

`func (o *FSUserGroup) SetQuotaNum(v int64)`

SetQuotaNum sets QuotaNum field to given value.

### HasQuotaNum

`func (o *FSUserGroup) HasQuotaNum() bool`

HasQuotaNum returns a boolean if a field has been set.

### GetShareNums

`func (o *FSUserGroup) GetShareNums() map[string]int64`

GetShareNums returns the ShareNums field if non-nil, zero value otherwise.

### GetShareNumsOk

`func (o *FSUserGroup) GetShareNumsOk() (*map[string]int64, bool)`

GetShareNumsOk returns a tuple with the ShareNums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareNums

`func (o *FSUserGroup) SetShareNums(v map[string]int64)`

SetShareNums sets ShareNums field to given value.

### HasShareNums

`func (o *FSUserGroup) HasShareNums() bool`

HasShareNums returns a boolean if a field has been set.

### GetUpdate

`func (o *FSUserGroup) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *FSUserGroup) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *FSUserGroup) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *FSUserGroup) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUserNum

`func (o *FSUserGroup) GetUserNum() int64`

GetUserNum returns the UserNum field if non-nil, zero value otherwise.

### GetUserNumOk

`func (o *FSUserGroup) GetUserNumOk() (*int64, bool)`

GetUserNumOk returns a tuple with the UserNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNum

`func (o *FSUserGroup) SetUserNum(v int64)`

SetUserNum sets UserNum field to given value.

### HasUserNum

`func (o *FSUserGroup) HasUserNum() bool`

HasUserNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


