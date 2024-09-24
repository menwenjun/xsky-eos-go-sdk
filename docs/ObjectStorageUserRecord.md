# ObjectStorageUserRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketNum** | Pointer to **int64** |  | [optional] 
**BucketQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**BucketQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ExternalBucketQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**ExternalBucketQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**ExternalUserQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**ExternalUserQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LocalBucketQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**LocalBucketQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**LocalUserQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**LocalUserQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**LocationConstraintEnabled** | Pointer to **bool** |  | [optional] 
**MaxBuckets** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OpMask** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**ObjectStorageUserNestview**](ObjectStorageUserNestview.md) |  | [optional] 
**Policy** | Pointer to [**ObjectStoragePolicyNestview**](ObjectStoragePolicyNestview.md) |  | [optional] 
**PolicyPollingEnabled** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to [**OSUserProperties**](OSUserProperties.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**UserQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**UserQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**Samples** | Pointer to [**[]ObjectStorageUserStat**](ObjectStorageUserStat.md) |  | [optional] 

## Methods

### NewObjectStorageUserRecord

`func NewObjectStorageUserRecord() *ObjectStorageUserRecord`

NewObjectStorageUserRecord instantiates a new ObjectStorageUserRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageUserRecordWithDefaults

`func NewObjectStorageUserRecordWithDefaults() *ObjectStorageUserRecord`

NewObjectStorageUserRecordWithDefaults instantiates a new ObjectStorageUserRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketNum

`func (o *ObjectStorageUserRecord) GetBucketNum() int64`

GetBucketNum returns the BucketNum field if non-nil, zero value otherwise.

### GetBucketNumOk

`func (o *ObjectStorageUserRecord) GetBucketNumOk() (*int64, bool)`

GetBucketNumOk returns a tuple with the BucketNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketNum

`func (o *ObjectStorageUserRecord) SetBucketNum(v int64)`

SetBucketNum sets BucketNum field to given value.

### HasBucketNum

`func (o *ObjectStorageUserRecord) HasBucketNum() bool`

HasBucketNum returns a boolean if a field has been set.

### GetBucketQuotaMaxObjects

`func (o *ObjectStorageUserRecord) GetBucketQuotaMaxObjects() int64`

GetBucketQuotaMaxObjects returns the BucketQuotaMaxObjects field if non-nil, zero value otherwise.

### GetBucketQuotaMaxObjectsOk

`func (o *ObjectStorageUserRecord) GetBucketQuotaMaxObjectsOk() (*int64, bool)`

GetBucketQuotaMaxObjectsOk returns a tuple with the BucketQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketQuotaMaxObjects

`func (o *ObjectStorageUserRecord) SetBucketQuotaMaxObjects(v int64)`

SetBucketQuotaMaxObjects sets BucketQuotaMaxObjects field to given value.

### HasBucketQuotaMaxObjects

`func (o *ObjectStorageUserRecord) HasBucketQuotaMaxObjects() bool`

HasBucketQuotaMaxObjects returns a boolean if a field has been set.

### GetBucketQuotaMaxSize

`func (o *ObjectStorageUserRecord) GetBucketQuotaMaxSize() int64`

GetBucketQuotaMaxSize returns the BucketQuotaMaxSize field if non-nil, zero value otherwise.

### GetBucketQuotaMaxSizeOk

`func (o *ObjectStorageUserRecord) GetBucketQuotaMaxSizeOk() (*int64, bool)`

GetBucketQuotaMaxSizeOk returns a tuple with the BucketQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketQuotaMaxSize

`func (o *ObjectStorageUserRecord) SetBucketQuotaMaxSize(v int64)`

SetBucketQuotaMaxSize sets BucketQuotaMaxSize field to given value.

### HasBucketQuotaMaxSize

`func (o *ObjectStorageUserRecord) HasBucketQuotaMaxSize() bool`

HasBucketQuotaMaxSize returns a boolean if a field has been set.

### GetCluster

`func (o *ObjectStorageUserRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageUserRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageUserRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageUserRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageUserRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageUserRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageUserRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageUserRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStorageUserRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStorageUserRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStorageUserRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStorageUserRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ObjectStorageUserRecord) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ObjectStorageUserRecord) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ObjectStorageUserRecord) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ObjectStorageUserRecord) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *ObjectStorageUserRecord) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ObjectStorageUserRecord) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ObjectStorageUserRecord) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ObjectStorageUserRecord) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalBucketQuotaMaxObjects

`func (o *ObjectStorageUserRecord) GetExternalBucketQuotaMaxObjects() int64`

GetExternalBucketQuotaMaxObjects returns the ExternalBucketQuotaMaxObjects field if non-nil, zero value otherwise.

### GetExternalBucketQuotaMaxObjectsOk

`func (o *ObjectStorageUserRecord) GetExternalBucketQuotaMaxObjectsOk() (*int64, bool)`

GetExternalBucketQuotaMaxObjectsOk returns a tuple with the ExternalBucketQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBucketQuotaMaxObjects

`func (o *ObjectStorageUserRecord) SetExternalBucketQuotaMaxObjects(v int64)`

SetExternalBucketQuotaMaxObjects sets ExternalBucketQuotaMaxObjects field to given value.

### HasExternalBucketQuotaMaxObjects

`func (o *ObjectStorageUserRecord) HasExternalBucketQuotaMaxObjects() bool`

HasExternalBucketQuotaMaxObjects returns a boolean if a field has been set.

### GetExternalBucketQuotaMaxSize

`func (o *ObjectStorageUserRecord) GetExternalBucketQuotaMaxSize() int64`

GetExternalBucketQuotaMaxSize returns the ExternalBucketQuotaMaxSize field if non-nil, zero value otherwise.

### GetExternalBucketQuotaMaxSizeOk

`func (o *ObjectStorageUserRecord) GetExternalBucketQuotaMaxSizeOk() (*int64, bool)`

GetExternalBucketQuotaMaxSizeOk returns a tuple with the ExternalBucketQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBucketQuotaMaxSize

`func (o *ObjectStorageUserRecord) SetExternalBucketQuotaMaxSize(v int64)`

SetExternalBucketQuotaMaxSize sets ExternalBucketQuotaMaxSize field to given value.

### HasExternalBucketQuotaMaxSize

`func (o *ObjectStorageUserRecord) HasExternalBucketQuotaMaxSize() bool`

HasExternalBucketQuotaMaxSize returns a boolean if a field has been set.

### GetExternalUserQuotaMaxObjects

`func (o *ObjectStorageUserRecord) GetExternalUserQuotaMaxObjects() int64`

GetExternalUserQuotaMaxObjects returns the ExternalUserQuotaMaxObjects field if non-nil, zero value otherwise.

### GetExternalUserQuotaMaxObjectsOk

`func (o *ObjectStorageUserRecord) GetExternalUserQuotaMaxObjectsOk() (*int64, bool)`

GetExternalUserQuotaMaxObjectsOk returns a tuple with the ExternalUserQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserQuotaMaxObjects

`func (o *ObjectStorageUserRecord) SetExternalUserQuotaMaxObjects(v int64)`

SetExternalUserQuotaMaxObjects sets ExternalUserQuotaMaxObjects field to given value.

### HasExternalUserQuotaMaxObjects

`func (o *ObjectStorageUserRecord) HasExternalUserQuotaMaxObjects() bool`

HasExternalUserQuotaMaxObjects returns a boolean if a field has been set.

### GetExternalUserQuotaMaxSize

`func (o *ObjectStorageUserRecord) GetExternalUserQuotaMaxSize() int64`

GetExternalUserQuotaMaxSize returns the ExternalUserQuotaMaxSize field if non-nil, zero value otherwise.

### GetExternalUserQuotaMaxSizeOk

`func (o *ObjectStorageUserRecord) GetExternalUserQuotaMaxSizeOk() (*int64, bool)`

GetExternalUserQuotaMaxSizeOk returns a tuple with the ExternalUserQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserQuotaMaxSize

`func (o *ObjectStorageUserRecord) SetExternalUserQuotaMaxSize(v int64)`

SetExternalUserQuotaMaxSize sets ExternalUserQuotaMaxSize field to given value.

### HasExternalUserQuotaMaxSize

`func (o *ObjectStorageUserRecord) HasExternalUserQuotaMaxSize() bool`

HasExternalUserQuotaMaxSize returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageUserRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageUserRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageUserRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageUserRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalBucketQuotaMaxObjects

`func (o *ObjectStorageUserRecord) GetLocalBucketQuotaMaxObjects() int64`

GetLocalBucketQuotaMaxObjects returns the LocalBucketQuotaMaxObjects field if non-nil, zero value otherwise.

### GetLocalBucketQuotaMaxObjectsOk

`func (o *ObjectStorageUserRecord) GetLocalBucketQuotaMaxObjectsOk() (*int64, bool)`

GetLocalBucketQuotaMaxObjectsOk returns a tuple with the LocalBucketQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalBucketQuotaMaxObjects

`func (o *ObjectStorageUserRecord) SetLocalBucketQuotaMaxObjects(v int64)`

SetLocalBucketQuotaMaxObjects sets LocalBucketQuotaMaxObjects field to given value.

### HasLocalBucketQuotaMaxObjects

`func (o *ObjectStorageUserRecord) HasLocalBucketQuotaMaxObjects() bool`

HasLocalBucketQuotaMaxObjects returns a boolean if a field has been set.

### GetLocalBucketQuotaMaxSize

`func (o *ObjectStorageUserRecord) GetLocalBucketQuotaMaxSize() int64`

GetLocalBucketQuotaMaxSize returns the LocalBucketQuotaMaxSize field if non-nil, zero value otherwise.

### GetLocalBucketQuotaMaxSizeOk

`func (o *ObjectStorageUserRecord) GetLocalBucketQuotaMaxSizeOk() (*int64, bool)`

GetLocalBucketQuotaMaxSizeOk returns a tuple with the LocalBucketQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalBucketQuotaMaxSize

`func (o *ObjectStorageUserRecord) SetLocalBucketQuotaMaxSize(v int64)`

SetLocalBucketQuotaMaxSize sets LocalBucketQuotaMaxSize field to given value.

### HasLocalBucketQuotaMaxSize

`func (o *ObjectStorageUserRecord) HasLocalBucketQuotaMaxSize() bool`

HasLocalBucketQuotaMaxSize returns a boolean if a field has been set.

### GetLocalUserQuotaMaxObjects

`func (o *ObjectStorageUserRecord) GetLocalUserQuotaMaxObjects() int64`

GetLocalUserQuotaMaxObjects returns the LocalUserQuotaMaxObjects field if non-nil, zero value otherwise.

### GetLocalUserQuotaMaxObjectsOk

`func (o *ObjectStorageUserRecord) GetLocalUserQuotaMaxObjectsOk() (*int64, bool)`

GetLocalUserQuotaMaxObjectsOk returns a tuple with the LocalUserQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserQuotaMaxObjects

`func (o *ObjectStorageUserRecord) SetLocalUserQuotaMaxObjects(v int64)`

SetLocalUserQuotaMaxObjects sets LocalUserQuotaMaxObjects field to given value.

### HasLocalUserQuotaMaxObjects

`func (o *ObjectStorageUserRecord) HasLocalUserQuotaMaxObjects() bool`

HasLocalUserQuotaMaxObjects returns a boolean if a field has been set.

### GetLocalUserQuotaMaxSize

`func (o *ObjectStorageUserRecord) GetLocalUserQuotaMaxSize() int64`

GetLocalUserQuotaMaxSize returns the LocalUserQuotaMaxSize field if non-nil, zero value otherwise.

### GetLocalUserQuotaMaxSizeOk

`func (o *ObjectStorageUserRecord) GetLocalUserQuotaMaxSizeOk() (*int64, bool)`

GetLocalUserQuotaMaxSizeOk returns a tuple with the LocalUserQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserQuotaMaxSize

`func (o *ObjectStorageUserRecord) SetLocalUserQuotaMaxSize(v int64)`

SetLocalUserQuotaMaxSize sets LocalUserQuotaMaxSize field to given value.

### HasLocalUserQuotaMaxSize

`func (o *ObjectStorageUserRecord) HasLocalUserQuotaMaxSize() bool`

HasLocalUserQuotaMaxSize returns a boolean if a field has been set.

### GetLocationConstraintEnabled

`func (o *ObjectStorageUserRecord) GetLocationConstraintEnabled() bool`

GetLocationConstraintEnabled returns the LocationConstraintEnabled field if non-nil, zero value otherwise.

### GetLocationConstraintEnabledOk

`func (o *ObjectStorageUserRecord) GetLocationConstraintEnabledOk() (*bool, bool)`

GetLocationConstraintEnabledOk returns a tuple with the LocationConstraintEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationConstraintEnabled

`func (o *ObjectStorageUserRecord) SetLocationConstraintEnabled(v bool)`

SetLocationConstraintEnabled sets LocationConstraintEnabled field to given value.

### HasLocationConstraintEnabled

`func (o *ObjectStorageUserRecord) HasLocationConstraintEnabled() bool`

HasLocationConstraintEnabled returns a boolean if a field has been set.

### GetMaxBuckets

`func (o *ObjectStorageUserRecord) GetMaxBuckets() int64`

GetMaxBuckets returns the MaxBuckets field if non-nil, zero value otherwise.

### GetMaxBucketsOk

`func (o *ObjectStorageUserRecord) GetMaxBucketsOk() (*int64, bool)`

GetMaxBucketsOk returns a tuple with the MaxBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBuckets

`func (o *ObjectStorageUserRecord) SetMaxBuckets(v int64)`

SetMaxBuckets sets MaxBuckets field to given value.

### HasMaxBuckets

`func (o *ObjectStorageUserRecord) HasMaxBuckets() bool`

HasMaxBuckets returns a boolean if a field has been set.

### GetName

`func (o *ObjectStorageUserRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageUserRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageUserRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectStorageUserRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpMask

`func (o *ObjectStorageUserRecord) GetOpMask() string`

GetOpMask returns the OpMask field if non-nil, zero value otherwise.

### GetOpMaskOk

`func (o *ObjectStorageUserRecord) GetOpMaskOk() (*string, bool)`

GetOpMaskOk returns a tuple with the OpMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpMask

`func (o *ObjectStorageUserRecord) SetOpMask(v string)`

SetOpMask sets OpMask field to given value.

### HasOpMask

`func (o *ObjectStorageUserRecord) HasOpMask() bool`

HasOpMask returns a boolean if a field has been set.

### GetParent

`func (o *ObjectStorageUserRecord) GetParent() ObjectStorageUserNestview`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ObjectStorageUserRecord) GetParentOk() (*ObjectStorageUserNestview, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ObjectStorageUserRecord) SetParent(v ObjectStorageUserNestview)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ObjectStorageUserRecord) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPolicy

`func (o *ObjectStorageUserRecord) GetPolicy() ObjectStoragePolicyNestview`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ObjectStorageUserRecord) GetPolicyOk() (*ObjectStoragePolicyNestview, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ObjectStorageUserRecord) SetPolicy(v ObjectStoragePolicyNestview)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ObjectStorageUserRecord) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPolicyPollingEnabled

`func (o *ObjectStorageUserRecord) GetPolicyPollingEnabled() bool`

GetPolicyPollingEnabled returns the PolicyPollingEnabled field if non-nil, zero value otherwise.

### GetPolicyPollingEnabledOk

`func (o *ObjectStorageUserRecord) GetPolicyPollingEnabledOk() (*bool, bool)`

GetPolicyPollingEnabledOk returns a tuple with the PolicyPollingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyPollingEnabled

`func (o *ObjectStorageUserRecord) SetPolicyPollingEnabled(v bool)`

SetPolicyPollingEnabled sets PolicyPollingEnabled field to given value.

### HasPolicyPollingEnabled

`func (o *ObjectStorageUserRecord) HasPolicyPollingEnabled() bool`

HasPolicyPollingEnabled returns a boolean if a field has been set.

### GetProperties

`func (o *ObjectStorageUserRecord) GetProperties() OSUserProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ObjectStorageUserRecord) GetPropertiesOk() (*OSUserProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ObjectStorageUserRecord) SetProperties(v OSUserProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ObjectStorageUserRecord) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageUserRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageUserRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageUserRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageUserRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuspended

`func (o *ObjectStorageUserRecord) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ObjectStorageUserRecord) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ObjectStorageUserRecord) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ObjectStorageUserRecord) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetUpdate

`func (o *ObjectStorageUserRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ObjectStorageUserRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ObjectStorageUserRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ObjectStorageUserRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetUserQuotaMaxObjects

`func (o *ObjectStorageUserRecord) GetUserQuotaMaxObjects() int64`

GetUserQuotaMaxObjects returns the UserQuotaMaxObjects field if non-nil, zero value otherwise.

### GetUserQuotaMaxObjectsOk

`func (o *ObjectStorageUserRecord) GetUserQuotaMaxObjectsOk() (*int64, bool)`

GetUserQuotaMaxObjectsOk returns a tuple with the UserQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuotaMaxObjects

`func (o *ObjectStorageUserRecord) SetUserQuotaMaxObjects(v int64)`

SetUserQuotaMaxObjects sets UserQuotaMaxObjects field to given value.

### HasUserQuotaMaxObjects

`func (o *ObjectStorageUserRecord) HasUserQuotaMaxObjects() bool`

HasUserQuotaMaxObjects returns a boolean if a field has been set.

### GetUserQuotaMaxSize

`func (o *ObjectStorageUserRecord) GetUserQuotaMaxSize() int64`

GetUserQuotaMaxSize returns the UserQuotaMaxSize field if non-nil, zero value otherwise.

### GetUserQuotaMaxSizeOk

`func (o *ObjectStorageUserRecord) GetUserQuotaMaxSizeOk() (*int64, bool)`

GetUserQuotaMaxSizeOk returns a tuple with the UserQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuotaMaxSize

`func (o *ObjectStorageUserRecord) SetUserQuotaMaxSize(v int64)`

SetUserQuotaMaxSize sets UserQuotaMaxSize field to given value.

### HasUserQuotaMaxSize

`func (o *ObjectStorageUserRecord) HasUserQuotaMaxSize() bool`

HasUserQuotaMaxSize returns a boolean if a field has been set.

### GetSamples

`func (o *ObjectStorageUserRecord) GetSamples() []ObjectStorageUserStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ObjectStorageUserRecord) GetSamplesOk() (*[]ObjectStorageUserStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ObjectStorageUserRecord) SetSamples(v []ObjectStorageUserStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ObjectStorageUserRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


