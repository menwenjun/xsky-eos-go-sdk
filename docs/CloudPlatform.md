# CloudPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**AttachedCloudVolumeNum** | Pointer to **int64** |  | [optional] 
**CloudInstanceNum** | Pointer to **int64** |  | [optional] 
**CloudVolumeNum** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExtraProperties** | Pointer to [**CloudExtraProperties**](CloudExtraProperties.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RelatedResourceData** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SyncTime** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudPlatform

`func NewCloudPlatform() *CloudPlatform`

NewCloudPlatform instantiates a new CloudPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPlatformWithDefaults

`func NewCloudPlatformWithDefaults() *CloudPlatform`

NewCloudPlatformWithDefaults instantiates a new CloudPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *CloudPlatform) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *CloudPlatform) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *CloudPlatform) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *CloudPlatform) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetAttachedCloudVolumeNum

`func (o *CloudPlatform) GetAttachedCloudVolumeNum() int64`

GetAttachedCloudVolumeNum returns the AttachedCloudVolumeNum field if non-nil, zero value otherwise.

### GetAttachedCloudVolumeNumOk

`func (o *CloudPlatform) GetAttachedCloudVolumeNumOk() (*int64, bool)`

GetAttachedCloudVolumeNumOk returns a tuple with the AttachedCloudVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedCloudVolumeNum

`func (o *CloudPlatform) SetAttachedCloudVolumeNum(v int64)`

SetAttachedCloudVolumeNum sets AttachedCloudVolumeNum field to given value.

### HasAttachedCloudVolumeNum

`func (o *CloudPlatform) HasAttachedCloudVolumeNum() bool`

HasAttachedCloudVolumeNum returns a boolean if a field has been set.

### GetCloudInstanceNum

`func (o *CloudPlatform) GetCloudInstanceNum() int64`

GetCloudInstanceNum returns the CloudInstanceNum field if non-nil, zero value otherwise.

### GetCloudInstanceNumOk

`func (o *CloudPlatform) GetCloudInstanceNumOk() (*int64, bool)`

GetCloudInstanceNumOk returns a tuple with the CloudInstanceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceNum

`func (o *CloudPlatform) SetCloudInstanceNum(v int64)`

SetCloudInstanceNum sets CloudInstanceNum field to given value.

### HasCloudInstanceNum

`func (o *CloudPlatform) HasCloudInstanceNum() bool`

HasCloudInstanceNum returns a boolean if a field has been set.

### GetCloudVolumeNum

`func (o *CloudPlatform) GetCloudVolumeNum() int64`

GetCloudVolumeNum returns the CloudVolumeNum field if non-nil, zero value otherwise.

### GetCloudVolumeNumOk

`func (o *CloudPlatform) GetCloudVolumeNumOk() (*int64, bool)`

GetCloudVolumeNumOk returns a tuple with the CloudVolumeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudVolumeNum

`func (o *CloudPlatform) SetCloudVolumeNum(v int64)`

SetCloudVolumeNum sets CloudVolumeNum field to given value.

### HasCloudVolumeNum

`func (o *CloudPlatform) HasCloudVolumeNum() bool`

HasCloudVolumeNum returns a boolean if a field has been set.

### GetCluster

`func (o *CloudPlatform) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CloudPlatform) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CloudPlatform) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *CloudPlatform) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *CloudPlatform) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *CloudPlatform) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *CloudPlatform) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *CloudPlatform) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *CloudPlatform) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudPlatform) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudPlatform) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudPlatform) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtraProperties

`func (o *CloudPlatform) GetExtraProperties() CloudExtraProperties`

GetExtraProperties returns the ExtraProperties field if non-nil, zero value otherwise.

### GetExtraPropertiesOk

`func (o *CloudPlatform) GetExtraPropertiesOk() (*CloudExtraProperties, bool)`

GetExtraPropertiesOk returns a tuple with the ExtraProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProperties

`func (o *CloudPlatform) SetExtraProperties(v CloudExtraProperties)`

SetExtraProperties sets ExtraProperties field to given value.

### HasExtraProperties

`func (o *CloudPlatform) HasExtraProperties() bool`

HasExtraProperties returns a boolean if a field has been set.

### GetId

`func (o *CloudPlatform) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudPlatform) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudPlatform) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CloudPlatform) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudPlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudPlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudPlatform) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudPlatform) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelatedResourceData

`func (o *CloudPlatform) GetRelatedResourceData() string`

GetRelatedResourceData returns the RelatedResourceData field if non-nil, zero value otherwise.

### GetRelatedResourceDataOk

`func (o *CloudPlatform) GetRelatedResourceDataOk() (*string, bool)`

GetRelatedResourceDataOk returns a tuple with the RelatedResourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceData

`func (o *CloudPlatform) SetRelatedResourceData(v string)`

SetRelatedResourceData sets RelatedResourceData field to given value.

### HasRelatedResourceData

`func (o *CloudPlatform) HasRelatedResourceData() bool`

HasRelatedResourceData returns a boolean if a field has been set.

### GetStatus

`func (o *CloudPlatform) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudPlatform) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudPlatform) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudPlatform) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyncTime

`func (o *CloudPlatform) GetSyncTime() time.Time`

GetSyncTime returns the SyncTime field if non-nil, zero value otherwise.

### GetSyncTimeOk

`func (o *CloudPlatform) GetSyncTimeOk() (*time.Time, bool)`

GetSyncTimeOk returns a tuple with the SyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTime

`func (o *CloudPlatform) SetSyncTime(v time.Time)`

SetSyncTime sets SyncTime field to given value.

### HasSyncTime

`func (o *CloudPlatform) HasSyncTime() bool`

HasSyncTime returns a boolean if a field has been set.

### GetType

`func (o *CloudPlatform) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudPlatform) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudPlatform) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudPlatform) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdate

`func (o *CloudPlatform) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *CloudPlatform) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *CloudPlatform) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *CloudPlatform) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUrl

`func (o *CloudPlatform) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudPlatform) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudPlatform) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudPlatform) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *CloudPlatform) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CloudPlatform) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CloudPlatform) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CloudPlatform) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUuid

`func (o *CloudPlatform) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CloudPlatform) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CloudPlatform) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CloudPlatform) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


