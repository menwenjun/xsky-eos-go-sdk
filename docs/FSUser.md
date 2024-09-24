# FSUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketLimit** | Pointer to **int64** |  | [optional] 
**BucketNum** | Pointer to **int64** |  | [optional] 
**BucketParentPath** | Pointer to [**DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**BucketPermission** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FsUserGroupNum** | Pointer to **int64** |  | [optional] 
**FsUserGroups** | Pointer to [**[]FSUserGroupNestview**](FSUserGroupNestview.md) |  | [optional] 
**GatewayGroup** | Pointer to [**DfsGatewayGroup**](DfsGatewayGroup.md) |  | [optional] 
**HdfsNum** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PolicyFilterNum** | Pointer to **int64** |  | [optional] 
**PrimaryGroup** | Pointer to [**FSUserGroupNestview**](FSUserGroupNestview.md) |  | [optional] 
**QuotaNum** | Pointer to **int64** |  | [optional] 
**S3Enabled** | Pointer to **bool** | dfs s3 support | [optional] 
**S3Status** | Pointer to **string** | indicate that s3 service is ok or not | [optional] 
**ShareNums** | Pointer to **map[string]int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewFSUser

`func NewFSUser() *FSUser`

NewFSUser instantiates a new FSUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFSUserWithDefaults

`func NewFSUserWithDefaults() *FSUser`

NewFSUserWithDefaults instantiates a new FSUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketLimit

`func (o *FSUser) GetBucketLimit() int64`

GetBucketLimit returns the BucketLimit field if non-nil, zero value otherwise.

### GetBucketLimitOk

`func (o *FSUser) GetBucketLimitOk() (*int64, bool)`

GetBucketLimitOk returns a tuple with the BucketLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketLimit

`func (o *FSUser) SetBucketLimit(v int64)`

SetBucketLimit sets BucketLimit field to given value.

### HasBucketLimit

`func (o *FSUser) HasBucketLimit() bool`

HasBucketLimit returns a boolean if a field has been set.

### GetBucketNum

`func (o *FSUser) GetBucketNum() int64`

GetBucketNum returns the BucketNum field if non-nil, zero value otherwise.

### GetBucketNumOk

`func (o *FSUser) GetBucketNumOk() (*int64, bool)`

GetBucketNumOk returns a tuple with the BucketNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketNum

`func (o *FSUser) SetBucketNum(v int64)`

SetBucketNum sets BucketNum field to given value.

### HasBucketNum

`func (o *FSUser) HasBucketNum() bool`

HasBucketNum returns a boolean if a field has been set.

### GetBucketParentPath

`func (o *FSUser) GetBucketParentPath() DfsPathNestview`

GetBucketParentPath returns the BucketParentPath field if non-nil, zero value otherwise.

### GetBucketParentPathOk

`func (o *FSUser) GetBucketParentPathOk() (*DfsPathNestview, bool)`

GetBucketParentPathOk returns a tuple with the BucketParentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketParentPath

`func (o *FSUser) SetBucketParentPath(v DfsPathNestview)`

SetBucketParentPath sets BucketParentPath field to given value.

### HasBucketParentPath

`func (o *FSUser) HasBucketParentPath() bool`

HasBucketParentPath returns a boolean if a field has been set.

### GetBucketPermission

`func (o *FSUser) GetBucketPermission() string`

GetBucketPermission returns the BucketPermission field if non-nil, zero value otherwise.

### GetBucketPermissionOk

`func (o *FSUser) GetBucketPermissionOk() (*string, bool)`

GetBucketPermissionOk returns a tuple with the BucketPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPermission

`func (o *FSUser) SetBucketPermission(v string)`

SetBucketPermission sets BucketPermission field to given value.

### HasBucketPermission

`func (o *FSUser) HasBucketPermission() bool`

HasBucketPermission returns a boolean if a field has been set.

### GetCluster

`func (o *FSUser) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *FSUser) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *FSUser) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *FSUser) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *FSUser) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FSUser) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FSUser) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FSUser) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *FSUser) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FSUser) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FSUser) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FSUser) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *FSUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FSUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FSUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FSUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFsUserGroupNum

`func (o *FSUser) GetFsUserGroupNum() int64`

GetFsUserGroupNum returns the FsUserGroupNum field if non-nil, zero value otherwise.

### GetFsUserGroupNumOk

`func (o *FSUser) GetFsUserGroupNumOk() (*int64, bool)`

GetFsUserGroupNumOk returns a tuple with the FsUserGroupNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroupNum

`func (o *FSUser) SetFsUserGroupNum(v int64)`

SetFsUserGroupNum sets FsUserGroupNum field to given value.

### HasFsUserGroupNum

`func (o *FSUser) HasFsUserGroupNum() bool`

HasFsUserGroupNum returns a boolean if a field has been set.

### GetFsUserGroups

`func (o *FSUser) GetFsUserGroups() []FSUserGroupNestview`

GetFsUserGroups returns the FsUserGroups field if non-nil, zero value otherwise.

### GetFsUserGroupsOk

`func (o *FSUser) GetFsUserGroupsOk() (*[]FSUserGroupNestview, bool)`

GetFsUserGroupsOk returns a tuple with the FsUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUserGroups

`func (o *FSUser) SetFsUserGroups(v []FSUserGroupNestview)`

SetFsUserGroups sets FsUserGroups field to given value.

### HasFsUserGroups

`func (o *FSUser) HasFsUserGroups() bool`

HasFsUserGroups returns a boolean if a field has been set.

### GetGatewayGroup

`func (o *FSUser) GetGatewayGroup() DfsGatewayGroup`

GetGatewayGroup returns the GatewayGroup field if non-nil, zero value otherwise.

### GetGatewayGroupOk

`func (o *FSUser) GetGatewayGroupOk() (*DfsGatewayGroup, bool)`

GetGatewayGroupOk returns a tuple with the GatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayGroup

`func (o *FSUser) SetGatewayGroup(v DfsGatewayGroup)`

SetGatewayGroup sets GatewayGroup field to given value.

### HasGatewayGroup

`func (o *FSUser) HasGatewayGroup() bool`

HasGatewayGroup returns a boolean if a field has been set.

### GetHdfsNum

`func (o *FSUser) GetHdfsNum() int64`

GetHdfsNum returns the HdfsNum field if non-nil, zero value otherwise.

### GetHdfsNumOk

`func (o *FSUser) GetHdfsNumOk() (*int64, bool)`

GetHdfsNumOk returns a tuple with the HdfsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsNum

`func (o *FSUser) SetHdfsNum(v int64)`

SetHdfsNum sets HdfsNum field to given value.

### HasHdfsNum

`func (o *FSUser) HasHdfsNum() bool`

HasHdfsNum returns a boolean if a field has been set.

### GetId

`func (o *FSUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FSUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FSUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FSUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FSUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FSUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FSUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FSUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyFilterNum

`func (o *FSUser) GetPolicyFilterNum() int64`

GetPolicyFilterNum returns the PolicyFilterNum field if non-nil, zero value otherwise.

### GetPolicyFilterNumOk

`func (o *FSUser) GetPolicyFilterNumOk() (*int64, bool)`

GetPolicyFilterNumOk returns a tuple with the PolicyFilterNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyFilterNum

`func (o *FSUser) SetPolicyFilterNum(v int64)`

SetPolicyFilterNum sets PolicyFilterNum field to given value.

### HasPolicyFilterNum

`func (o *FSUser) HasPolicyFilterNum() bool`

HasPolicyFilterNum returns a boolean if a field has been set.

### GetPrimaryGroup

`func (o *FSUser) GetPrimaryGroup() FSUserGroupNestview`

GetPrimaryGroup returns the PrimaryGroup field if non-nil, zero value otherwise.

### GetPrimaryGroupOk

`func (o *FSUser) GetPrimaryGroupOk() (*FSUserGroupNestview, bool)`

GetPrimaryGroupOk returns a tuple with the PrimaryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGroup

`func (o *FSUser) SetPrimaryGroup(v FSUserGroupNestview)`

SetPrimaryGroup sets PrimaryGroup field to given value.

### HasPrimaryGroup

`func (o *FSUser) HasPrimaryGroup() bool`

HasPrimaryGroup returns a boolean if a field has been set.

### GetQuotaNum

`func (o *FSUser) GetQuotaNum() int64`

GetQuotaNum returns the QuotaNum field if non-nil, zero value otherwise.

### GetQuotaNumOk

`func (o *FSUser) GetQuotaNumOk() (*int64, bool)`

GetQuotaNumOk returns a tuple with the QuotaNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaNum

`func (o *FSUser) SetQuotaNum(v int64)`

SetQuotaNum sets QuotaNum field to given value.

### HasQuotaNum

`func (o *FSUser) HasQuotaNum() bool`

HasQuotaNum returns a boolean if a field has been set.

### GetS3Enabled

`func (o *FSUser) GetS3Enabled() bool`

GetS3Enabled returns the S3Enabled field if non-nil, zero value otherwise.

### GetS3EnabledOk

`func (o *FSUser) GetS3EnabledOk() (*bool, bool)`

GetS3EnabledOk returns a tuple with the S3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Enabled

`func (o *FSUser) SetS3Enabled(v bool)`

SetS3Enabled sets S3Enabled field to given value.

### HasS3Enabled

`func (o *FSUser) HasS3Enabled() bool`

HasS3Enabled returns a boolean if a field has been set.

### GetS3Status

`func (o *FSUser) GetS3Status() string`

GetS3Status returns the S3Status field if non-nil, zero value otherwise.

### GetS3StatusOk

`func (o *FSUser) GetS3StatusOk() (*string, bool)`

GetS3StatusOk returns a tuple with the S3Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Status

`func (o *FSUser) SetS3Status(v string)`

SetS3Status sets S3Status field to given value.

### HasS3Status

`func (o *FSUser) HasS3Status() bool`

HasS3Status returns a boolean if a field has been set.

### GetShareNums

`func (o *FSUser) GetShareNums() map[string]int64`

GetShareNums returns the ShareNums field if non-nil, zero value otherwise.

### GetShareNumsOk

`func (o *FSUser) GetShareNumsOk() (*map[string]int64, bool)`

GetShareNumsOk returns a tuple with the ShareNums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareNums

`func (o *FSUser) SetShareNums(v map[string]int64)`

SetShareNums sets ShareNums field to given value.

### HasShareNums

`func (o *FSUser) HasShareNums() bool`

HasShareNums returns a boolean if a field has been set.

### GetUpdate

`func (o *FSUser) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *FSUser) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *FSUser) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *FSUser) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUserId

`func (o *FSUser) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FSUser) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FSUser) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FSUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVersion

`func (o *FSUser) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FSUser) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FSUser) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FSUser) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


