# ObjectStorageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**CnameEnabled** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**DnsNames** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IndexPool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**LifecycleEndOn** | Pointer to **string** |  | [optional] 
**LifecycleStartOn** | Pointer to **string** |  | [optional] 
**MultiZoneEnabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OriginPullHostIds** | Pointer to **[]int64** |  | [optional] 
**S3LbSystemUserAk** | Pointer to **string** |  | [optional] 
**S3LbSystemUserSk** | Pointer to **string** |  | [optional] 
**SearchEnabled** | Pointer to **bool** |  | [optional] 
**SecondMergenceEnabled** | Pointer to **bool** |  | [optional] 
**SecondMergenceEndOn** | Pointer to **string** |  | [optional] 
**SecondMergenceStartOn** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TieringHostIds** | Pointer to **[]int64** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 
**Samples** | Pointer to [**[]ObjectStorageStat**](ObjectStorageStat.md) |  | [optional] 

## Methods

### NewObjectStorageRecord

`func NewObjectStorageRecord() *ObjectStorageRecord`

NewObjectStorageRecord instantiates a new ObjectStorageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageRecordWithDefaults

`func NewObjectStorageRecordWithDefaults() *ObjectStorageRecord`

NewObjectStorageRecordWithDefaults instantiates a new ObjectStorageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *ObjectStorageRecord) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *ObjectStorageRecord) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *ObjectStorageRecord) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *ObjectStorageRecord) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetCluster

`func (o *ObjectStorageRecord) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ObjectStorageRecord) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ObjectStorageRecord) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ObjectStorageRecord) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCnameEnabled

`func (o *ObjectStorageRecord) GetCnameEnabled() bool`

GetCnameEnabled returns the CnameEnabled field if non-nil, zero value otherwise.

### GetCnameEnabledOk

`func (o *ObjectStorageRecord) GetCnameEnabledOk() (*bool, bool)`

GetCnameEnabledOk returns a tuple with the CnameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameEnabled

`func (o *ObjectStorageRecord) SetCnameEnabled(v bool)`

SetCnameEnabled sets CnameEnabled field to given value.

### HasCnameEnabled

`func (o *ObjectStorageRecord) HasCnameEnabled() bool`

HasCnameEnabled returns a boolean if a field has been set.

### GetCreate

`func (o *ObjectStorageRecord) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ObjectStorageRecord) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ObjectStorageRecord) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ObjectStorageRecord) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDnsNames

`func (o *ObjectStorageRecord) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *ObjectStorageRecord) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *ObjectStorageRecord) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *ObjectStorageRecord) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.

### GetId

`func (o *ObjectStorageRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStorageRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStorageRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStorageRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndexPool

`func (o *ObjectStorageRecord) GetIndexPool() PoolNestview`

GetIndexPool returns the IndexPool field if non-nil, zero value otherwise.

### GetIndexPoolOk

`func (o *ObjectStorageRecord) GetIndexPoolOk() (*PoolNestview, bool)`

GetIndexPoolOk returns a tuple with the IndexPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPool

`func (o *ObjectStorageRecord) SetIndexPool(v PoolNestview)`

SetIndexPool sets IndexPool field to given value.

### HasIndexPool

`func (o *ObjectStorageRecord) HasIndexPool() bool`

HasIndexPool returns a boolean if a field has been set.

### GetLifecycleEndOn

`func (o *ObjectStorageRecord) GetLifecycleEndOn() string`

GetLifecycleEndOn returns the LifecycleEndOn field if non-nil, zero value otherwise.

### GetLifecycleEndOnOk

`func (o *ObjectStorageRecord) GetLifecycleEndOnOk() (*string, bool)`

GetLifecycleEndOnOk returns a tuple with the LifecycleEndOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleEndOn

`func (o *ObjectStorageRecord) SetLifecycleEndOn(v string)`

SetLifecycleEndOn sets LifecycleEndOn field to given value.

### HasLifecycleEndOn

`func (o *ObjectStorageRecord) HasLifecycleEndOn() bool`

HasLifecycleEndOn returns a boolean if a field has been set.

### GetLifecycleStartOn

`func (o *ObjectStorageRecord) GetLifecycleStartOn() string`

GetLifecycleStartOn returns the LifecycleStartOn field if non-nil, zero value otherwise.

### GetLifecycleStartOnOk

`func (o *ObjectStorageRecord) GetLifecycleStartOnOk() (*string, bool)`

GetLifecycleStartOnOk returns a tuple with the LifecycleStartOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStartOn

`func (o *ObjectStorageRecord) SetLifecycleStartOn(v string)`

SetLifecycleStartOn sets LifecycleStartOn field to given value.

### HasLifecycleStartOn

`func (o *ObjectStorageRecord) HasLifecycleStartOn() bool`

HasLifecycleStartOn returns a boolean if a field has been set.

### GetMultiZoneEnabled

`func (o *ObjectStorageRecord) GetMultiZoneEnabled() bool`

GetMultiZoneEnabled returns the MultiZoneEnabled field if non-nil, zero value otherwise.

### GetMultiZoneEnabledOk

`func (o *ObjectStorageRecord) GetMultiZoneEnabledOk() (*bool, bool)`

GetMultiZoneEnabledOk returns a tuple with the MultiZoneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiZoneEnabled

`func (o *ObjectStorageRecord) SetMultiZoneEnabled(v bool)`

SetMultiZoneEnabled sets MultiZoneEnabled field to given value.

### HasMultiZoneEnabled

`func (o *ObjectStorageRecord) HasMultiZoneEnabled() bool`

HasMultiZoneEnabled returns a boolean if a field has been set.

### GetName

`func (o *ObjectStorageRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectStorageRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginPullHostIds

`func (o *ObjectStorageRecord) GetOriginPullHostIds() []int64`

GetOriginPullHostIds returns the OriginPullHostIds field if non-nil, zero value otherwise.

### GetOriginPullHostIdsOk

`func (o *ObjectStorageRecord) GetOriginPullHostIdsOk() (*[]int64, bool)`

GetOriginPullHostIdsOk returns a tuple with the OriginPullHostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPullHostIds

`func (o *ObjectStorageRecord) SetOriginPullHostIds(v []int64)`

SetOriginPullHostIds sets OriginPullHostIds field to given value.

### HasOriginPullHostIds

`func (o *ObjectStorageRecord) HasOriginPullHostIds() bool`

HasOriginPullHostIds returns a boolean if a field has been set.

### GetS3LbSystemUserAk

`func (o *ObjectStorageRecord) GetS3LbSystemUserAk() string`

GetS3LbSystemUserAk returns the S3LbSystemUserAk field if non-nil, zero value otherwise.

### GetS3LbSystemUserAkOk

`func (o *ObjectStorageRecord) GetS3LbSystemUserAkOk() (*string, bool)`

GetS3LbSystemUserAkOk returns a tuple with the S3LbSystemUserAk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3LbSystemUserAk

`func (o *ObjectStorageRecord) SetS3LbSystemUserAk(v string)`

SetS3LbSystemUserAk sets S3LbSystemUserAk field to given value.

### HasS3LbSystemUserAk

`func (o *ObjectStorageRecord) HasS3LbSystemUserAk() bool`

HasS3LbSystemUserAk returns a boolean if a field has been set.

### GetS3LbSystemUserSk

`func (o *ObjectStorageRecord) GetS3LbSystemUserSk() string`

GetS3LbSystemUserSk returns the S3LbSystemUserSk field if non-nil, zero value otherwise.

### GetS3LbSystemUserSkOk

`func (o *ObjectStorageRecord) GetS3LbSystemUserSkOk() (*string, bool)`

GetS3LbSystemUserSkOk returns a tuple with the S3LbSystemUserSk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3LbSystemUserSk

`func (o *ObjectStorageRecord) SetS3LbSystemUserSk(v string)`

SetS3LbSystemUserSk sets S3LbSystemUserSk field to given value.

### HasS3LbSystemUserSk

`func (o *ObjectStorageRecord) HasS3LbSystemUserSk() bool`

HasS3LbSystemUserSk returns a boolean if a field has been set.

### GetSearchEnabled

`func (o *ObjectStorageRecord) GetSearchEnabled() bool`

GetSearchEnabled returns the SearchEnabled field if non-nil, zero value otherwise.

### GetSearchEnabledOk

`func (o *ObjectStorageRecord) GetSearchEnabledOk() (*bool, bool)`

GetSearchEnabledOk returns a tuple with the SearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEnabled

`func (o *ObjectStorageRecord) SetSearchEnabled(v bool)`

SetSearchEnabled sets SearchEnabled field to given value.

### HasSearchEnabled

`func (o *ObjectStorageRecord) HasSearchEnabled() bool`

HasSearchEnabled returns a boolean if a field has been set.

### GetSecondMergenceEnabled

`func (o *ObjectStorageRecord) GetSecondMergenceEnabled() bool`

GetSecondMergenceEnabled returns the SecondMergenceEnabled field if non-nil, zero value otherwise.

### GetSecondMergenceEnabledOk

`func (o *ObjectStorageRecord) GetSecondMergenceEnabledOk() (*bool, bool)`

GetSecondMergenceEnabledOk returns a tuple with the SecondMergenceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondMergenceEnabled

`func (o *ObjectStorageRecord) SetSecondMergenceEnabled(v bool)`

SetSecondMergenceEnabled sets SecondMergenceEnabled field to given value.

### HasSecondMergenceEnabled

`func (o *ObjectStorageRecord) HasSecondMergenceEnabled() bool`

HasSecondMergenceEnabled returns a boolean if a field has been set.

### GetSecondMergenceEndOn

`func (o *ObjectStorageRecord) GetSecondMergenceEndOn() string`

GetSecondMergenceEndOn returns the SecondMergenceEndOn field if non-nil, zero value otherwise.

### GetSecondMergenceEndOnOk

`func (o *ObjectStorageRecord) GetSecondMergenceEndOnOk() (*string, bool)`

GetSecondMergenceEndOnOk returns a tuple with the SecondMergenceEndOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondMergenceEndOn

`func (o *ObjectStorageRecord) SetSecondMergenceEndOn(v string)`

SetSecondMergenceEndOn sets SecondMergenceEndOn field to given value.

### HasSecondMergenceEndOn

`func (o *ObjectStorageRecord) HasSecondMergenceEndOn() bool`

HasSecondMergenceEndOn returns a boolean if a field has been set.

### GetSecondMergenceStartOn

`func (o *ObjectStorageRecord) GetSecondMergenceStartOn() string`

GetSecondMergenceStartOn returns the SecondMergenceStartOn field if non-nil, zero value otherwise.

### GetSecondMergenceStartOnOk

`func (o *ObjectStorageRecord) GetSecondMergenceStartOnOk() (*string, bool)`

GetSecondMergenceStartOnOk returns a tuple with the SecondMergenceStartOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondMergenceStartOn

`func (o *ObjectStorageRecord) SetSecondMergenceStartOn(v string)`

SetSecondMergenceStartOn sets SecondMergenceStartOn field to given value.

### HasSecondMergenceStartOn

`func (o *ObjectStorageRecord) HasSecondMergenceStartOn() bool`

HasSecondMergenceStartOn returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectStorageRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStorageRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTieringHostIds

`func (o *ObjectStorageRecord) GetTieringHostIds() []int64`

GetTieringHostIds returns the TieringHostIds field if non-nil, zero value otherwise.

### GetTieringHostIdsOk

`func (o *ObjectStorageRecord) GetTieringHostIdsOk() (*[]int64, bool)`

GetTieringHostIdsOk returns a tuple with the TieringHostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieringHostIds

`func (o *ObjectStorageRecord) SetTieringHostIds(v []int64)`

SetTieringHostIds sets TieringHostIds field to given value.

### HasTieringHostIds

`func (o *ObjectStorageRecord) HasTieringHostIds() bool`

HasTieringHostIds returns a boolean if a field has been set.

### GetUpdate

`func (o *ObjectStorageRecord) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ObjectStorageRecord) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ObjectStorageRecord) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ObjectStorageRecord) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSamples

`func (o *ObjectStorageRecord) GetSamples() []ObjectStorageStat`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ObjectStorageRecord) GetSamplesOk() (*[]ObjectStorageStat, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ObjectStorageRecord) SetSamples(v []ObjectStorageStat)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ObjectStorageRecord) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


