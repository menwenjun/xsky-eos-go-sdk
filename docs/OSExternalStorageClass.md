# OSExternalStorageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**ClassId** | Pointer to **string** |  | [optional] 
**ClassStatus** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalStoragePlatforms** | Pointer to [**[]OSExternalStoragePlatform**](OSExternalStoragePlatform.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OsPolicy** | Pointer to [**ObjectStoragePolicy**](ObjectStoragePolicy.md) |  | [optional] 
**OsPolicyId** | Pointer to **int64** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**PlatformType** | Pointer to **string** |  | [optional] 
**PrefixEnabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StoragePattern** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOSExternalStorageClass

`func NewOSExternalStorageClass() *OSExternalStorageClass`

NewOSExternalStorageClass instantiates a new OSExternalStorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSExternalStorageClassWithDefaults

`func NewOSExternalStorageClassWithDefaults() *OSExternalStorageClass`

NewOSExternalStorageClassWithDefaults instantiates a new OSExternalStorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *OSExternalStorageClass) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSExternalStorageClass) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSExternalStorageClass) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSExternalStorageClass) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetClassId

`func (o *OSExternalStorageClass) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OSExternalStorageClass) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OSExternalStorageClass) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *OSExternalStorageClass) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetClassStatus

`func (o *OSExternalStorageClass) GetClassStatus() string`

GetClassStatus returns the ClassStatus field if non-nil, zero value otherwise.

### GetClassStatusOk

`func (o *OSExternalStorageClass) GetClassStatusOk() (*string, bool)`

GetClassStatusOk returns a tuple with the ClassStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassStatus

`func (o *OSExternalStorageClass) SetClassStatus(v string)`

SetClassStatus sets ClassStatus field to given value.

### HasClassStatus

`func (o *OSExternalStorageClass) HasClassStatus() bool`

HasClassStatus returns a boolean if a field has been set.

### GetCluster

`func (o *OSExternalStorageClass) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSExternalStorageClass) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSExternalStorageClass) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSExternalStorageClass) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDescription

`func (o *OSExternalStorageClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OSExternalStorageClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OSExternalStorageClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OSExternalStorageClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalStoragePlatforms

`func (o *OSExternalStorageClass) GetExternalStoragePlatforms() []OSExternalStoragePlatform`

GetExternalStoragePlatforms returns the ExternalStoragePlatforms field if non-nil, zero value otherwise.

### GetExternalStoragePlatformsOk

`func (o *OSExternalStorageClass) GetExternalStoragePlatformsOk() (*[]OSExternalStoragePlatform, bool)`

GetExternalStoragePlatformsOk returns a tuple with the ExternalStoragePlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoragePlatforms

`func (o *OSExternalStorageClass) SetExternalStoragePlatforms(v []OSExternalStoragePlatform)`

SetExternalStoragePlatforms sets ExternalStoragePlatforms field to given value.

### HasExternalStoragePlatforms

`func (o *OSExternalStorageClass) HasExternalStoragePlatforms() bool`

HasExternalStoragePlatforms returns a boolean if a field has been set.

### GetId

`func (o *OSExternalStorageClass) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSExternalStorageClass) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSExternalStorageClass) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OSExternalStorageClass) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OSExternalStorageClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSExternalStorageClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSExternalStorageClass) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OSExternalStorageClass) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsPolicy

`func (o *OSExternalStorageClass) GetOsPolicy() ObjectStoragePolicy`

GetOsPolicy returns the OsPolicy field if non-nil, zero value otherwise.

### GetOsPolicyOk

`func (o *OSExternalStorageClass) GetOsPolicyOk() (*ObjectStoragePolicy, bool)`

GetOsPolicyOk returns a tuple with the OsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPolicy

`func (o *OSExternalStorageClass) SetOsPolicy(v ObjectStoragePolicy)`

SetOsPolicy sets OsPolicy field to given value.

### HasOsPolicy

`func (o *OSExternalStorageClass) HasOsPolicy() bool`

HasOsPolicy returns a boolean if a field has been set.

### GetOsPolicyId

`func (o *OSExternalStorageClass) GetOsPolicyId() int64`

GetOsPolicyId returns the OsPolicyId field if non-nil, zero value otherwise.

### GetOsPolicyIdOk

`func (o *OSExternalStorageClass) GetOsPolicyIdOk() (*int64, bool)`

GetOsPolicyIdOk returns a tuple with the OsPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPolicyId

`func (o *OSExternalStorageClass) SetOsPolicyId(v int64)`

SetOsPolicyId sets OsPolicyId field to given value.

### HasOsPolicyId

`func (o *OSExternalStorageClass) HasOsPolicyId() bool`

HasOsPolicyId returns a boolean if a field has been set.

### GetPlatform

`func (o *OSExternalStorageClass) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *OSExternalStorageClass) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *OSExternalStorageClass) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *OSExternalStorageClass) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformType

`func (o *OSExternalStorageClass) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *OSExternalStorageClass) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *OSExternalStorageClass) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *OSExternalStorageClass) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetPrefixEnabled

`func (o *OSExternalStorageClass) GetPrefixEnabled() bool`

GetPrefixEnabled returns the PrefixEnabled field if non-nil, zero value otherwise.

### GetPrefixEnabledOk

`func (o *OSExternalStorageClass) GetPrefixEnabledOk() (*bool, bool)`

GetPrefixEnabledOk returns a tuple with the PrefixEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixEnabled

`func (o *OSExternalStorageClass) SetPrefixEnabled(v bool)`

SetPrefixEnabled sets PrefixEnabled field to given value.

### HasPrefixEnabled

`func (o *OSExternalStorageClass) HasPrefixEnabled() bool`

HasPrefixEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *OSExternalStorageClass) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSExternalStorageClass) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSExternalStorageClass) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSExternalStorageClass) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStoragePattern

`func (o *OSExternalStorageClass) GetStoragePattern() []string`

GetStoragePattern returns the StoragePattern field if non-nil, zero value otherwise.

### GetStoragePatternOk

`func (o *OSExternalStorageClass) GetStoragePatternOk() (*[]string, bool)`

GetStoragePatternOk returns a tuple with the StoragePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePattern

`func (o *OSExternalStorageClass) SetStoragePattern(v []string)`

SetStoragePattern sets StoragePattern field to given value.

### HasStoragePattern

`func (o *OSExternalStorageClass) HasStoragePattern() bool`

HasStoragePattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


