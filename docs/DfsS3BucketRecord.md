# DfsS3BucketRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AllUserPermission** | Pointer to **string** |  | [optional] 
**AuthUserPermission** | Pointer to **string** |  | [optional] 
**BucketPolicy** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DataVerify** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DfsGatewayGroup** | Pointer to [**DfsGatewayGroupNestview**](DfsGatewayGroupNestview.md) |  | [optional] 
**DfsRootfs** | Pointer to [**DfsRootfsNestview**](DfsRootfsNestview.md) |  | [optional] 
**EnableEtag** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**FSUserNestview**](FSUserNestview.md) |  | [optional] 
**OwnerPermission** | Pointer to **string** |  | [optional] 
**Path** | Pointer to [**DfsPathNestview**](DfsPathNestview.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Samples** | Pointer to [**[]DfsS3BucketStat**](DfsS3BucketStat.md) |  | [optional] 

## Methods

### NewDfsS3BucketRecord

`func NewDfsS3BucketRecord() *DfsS3BucketRecord`

NewDfsS3BucketRecord instantiates a new DfsS3BucketRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsS3BucketRecordWithDefaults

`func NewDfsS3BucketRecordWithDefaults() *DfsS3BucketRecord`

NewDfsS3BucketRecordWithDefaults instantiates a new DfsS3BucketRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *DfsS3BucketRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *DfsS3BucketRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *DfsS3BucketRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *DfsS3BucketRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAllUserPermission

`func (o *DfsS3BucketRecord) GetAllUserPermission() string`

GetAllUserPermission returns the AllUserPermission field if non-nil, zero value otherwise.

### GetAllUserPermissionOk

`func (o *DfsS3BucketRecord) GetAllUserPermissionOk() (*string, bool)`

GetAllUserPermissionOk returns a tuple with the AllUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUserPermission

`func (o *DfsS3BucketRecord) SetAllUserPermission(v string)`

SetAllUserPermission sets AllUserPermission field to given value.

### HasAllUserPermission

`func (o *DfsS3BucketRecord) HasAllUserPermission() bool`

HasAllUserPermission returns a boolean if a field has been set.

### GetAuthUserPermission

`func (o *DfsS3BucketRecord) GetAuthUserPermission() string`

GetAuthUserPermission returns the AuthUserPermission field if non-nil, zero value otherwise.

### GetAuthUserPermissionOk

`func (o *DfsS3BucketRecord) GetAuthUserPermissionOk() (*string, bool)`

GetAuthUserPermissionOk returns a tuple with the AuthUserPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserPermission

`func (o *DfsS3BucketRecord) SetAuthUserPermission(v string)`

SetAuthUserPermission sets AuthUserPermission field to given value.

### HasAuthUserPermission

`func (o *DfsS3BucketRecord) HasAuthUserPermission() bool`

HasAuthUserPermission returns a boolean if a field has been set.

### GetBucketPolicy

`func (o *DfsS3BucketRecord) GetBucketPolicy() string`

GetBucketPolicy returns the BucketPolicy field if non-nil, zero value otherwise.

### GetBucketPolicyOk

`func (o *DfsS3BucketRecord) GetBucketPolicyOk() (*string, bool)`

GetBucketPolicyOk returns a tuple with the BucketPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPolicy

`func (o *DfsS3BucketRecord) SetBucketPolicy(v string)`

SetBucketPolicy sets BucketPolicy field to given value.

### HasBucketPolicy

`func (o *DfsS3BucketRecord) HasBucketPolicy() bool`

HasBucketPolicy returns a boolean if a field has been set.

### GetCluster

`func (o *DfsS3BucketRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DfsS3BucketRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DfsS3BucketRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DfsS3BucketRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *DfsS3BucketRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DfsS3BucketRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DfsS3BucketRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DfsS3BucketRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDataVerify

`func (o *DfsS3BucketRecord) GetDataVerify() bool`

GetDataVerify returns the DataVerify field if non-nil, zero value otherwise.

### GetDataVerifyOk

`func (o *DfsS3BucketRecord) GetDataVerifyOk() (*bool, bool)`

GetDataVerifyOk returns a tuple with the DataVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVerify

`func (o *DfsS3BucketRecord) SetDataVerify(v bool)`

SetDataVerify sets DataVerify field to given value.

### HasDataVerify

`func (o *DfsS3BucketRecord) HasDataVerify() bool`

HasDataVerify returns a boolean if a field has been set.

### GetDescription

`func (o *DfsS3BucketRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DfsS3BucketRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DfsS3BucketRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DfsS3BucketRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDfsGatewayGroup

`func (o *DfsS3BucketRecord) GetDfsGatewayGroup() DfsGatewayGroupNestview`

GetDfsGatewayGroup returns the DfsGatewayGroup field if non-nil, zero value otherwise.

### GetDfsGatewayGroupOk

`func (o *DfsS3BucketRecord) GetDfsGatewayGroupOk() (*DfsGatewayGroupNestview, bool)`

GetDfsGatewayGroupOk returns a tuple with the DfsGatewayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsGatewayGroup

`func (o *DfsS3BucketRecord) SetDfsGatewayGroup(v DfsGatewayGroupNestview)`

SetDfsGatewayGroup sets DfsGatewayGroup field to given value.

### HasDfsGatewayGroup

`func (o *DfsS3BucketRecord) HasDfsGatewayGroup() bool`

HasDfsGatewayGroup returns a boolean if a field has been set.

### GetDfsRootfs

`func (o *DfsS3BucketRecord) GetDfsRootfs() DfsRootfsNestview`

GetDfsRootfs returns the DfsRootfs field if non-nil, zero value otherwise.

### GetDfsRootfsOk

`func (o *DfsS3BucketRecord) GetDfsRootfsOk() (*DfsRootfsNestview, bool)`

GetDfsRootfsOk returns a tuple with the DfsRootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsRootfs

`func (o *DfsS3BucketRecord) SetDfsRootfs(v DfsRootfsNestview)`

SetDfsRootfs sets DfsRootfs field to given value.

### HasDfsRootfs

`func (o *DfsS3BucketRecord) HasDfsRootfs() bool`

HasDfsRootfs returns a boolean if a field has been set.

### GetEnableEtag

`func (o *DfsS3BucketRecord) GetEnableEtag() bool`

GetEnableEtag returns the EnableEtag field if non-nil, zero value otherwise.

### GetEnableEtagOk

`func (o *DfsS3BucketRecord) GetEnableEtagOk() (*bool, bool)`

GetEnableEtagOk returns a tuple with the EnableEtag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEtag

`func (o *DfsS3BucketRecord) SetEnableEtag(v bool)`

SetEnableEtag sets EnableEtag field to given value.

### HasEnableEtag

`func (o *DfsS3BucketRecord) HasEnableEtag() bool`

HasEnableEtag returns a boolean if a field has been set.

### GetId

`func (o *DfsS3BucketRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsS3BucketRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsS3BucketRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsS3BucketRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DfsS3BucketRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DfsS3BucketRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DfsS3BucketRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DfsS3BucketRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *DfsS3BucketRecord) GetOwner() FSUserNestview`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DfsS3BucketRecord) GetOwnerOk() (*FSUserNestview, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DfsS3BucketRecord) SetOwner(v FSUserNestview)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DfsS3BucketRecord) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerPermission

`func (o *DfsS3BucketRecord) GetOwnerPermission() string`

GetOwnerPermission returns the OwnerPermission field if non-nil, zero value otherwise.

### GetOwnerPermissionOk

`func (o *DfsS3BucketRecord) GetOwnerPermissionOk() (*string, bool)`

GetOwnerPermissionOk returns a tuple with the OwnerPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPermission

`func (o *DfsS3BucketRecord) SetOwnerPermission(v string)`

SetOwnerPermission sets OwnerPermission field to given value.

### HasOwnerPermission

`func (o *DfsS3BucketRecord) HasOwnerPermission() bool`

HasOwnerPermission returns a boolean if a field has been set.

### GetPath

`func (o *DfsS3BucketRecord) GetPath() DfsPathNestview`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DfsS3BucketRecord) GetPathOk() (*DfsPathNestview, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DfsS3BucketRecord) SetPath(v DfsPathNestview)`

SetPath sets Path field to given value.

### HasPath

`func (o *DfsS3BucketRecord) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStatus

`func (o *DfsS3BucketRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DfsS3BucketRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DfsS3BucketRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DfsS3BucketRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdate

`func (o *DfsS3BucketRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DfsS3BucketRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DfsS3BucketRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DfsS3BucketRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *DfsS3BucketRecord) GetSamples() []DfsS3BucketStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *DfsS3BucketRecord) GetSamplesOk() (*[]DfsS3BucketStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *DfsS3BucketRecord) SetSamples(v []DfsS3BucketStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *DfsS3BucketRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


