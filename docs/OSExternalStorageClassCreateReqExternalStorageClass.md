# OSExternalStorageClassCreateReqExternalStorageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ExternalStoragePlatforms** | [**[]OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt**](OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt.md) |  | 
**Name** | **string** |  | 
**OsPolicyId** | **int64** |  | 
**Platform** | **string** |  | 
**PlatformType** | **string** |  | 
**PrefixEnabled** | Pointer to **bool** |  | [optional] 
**StoragePattern** | Pointer to **[]string** | SliceStringField defines slice string field | [optional] 

## Methods

### NewOSExternalStorageClassCreateReqExternalStorageClass

`func NewOSExternalStorageClassCreateReqExternalStorageClass(classId string, externalStoragePlatforms []OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt, name string, osPolicyId int64, platform string, platformType string, ) *OSExternalStorageClassCreateReqExternalStorageClass`

NewOSExternalStorageClassCreateReqExternalStorageClass instantiates a new OSExternalStorageClassCreateReqExternalStorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSExternalStorageClassCreateReqExternalStorageClassWithDefaults

`func NewOSExternalStorageClassCreateReqExternalStorageClassWithDefaults() *OSExternalStorageClassCreateReqExternalStorageClass`

NewOSExternalStorageClassCreateReqExternalStorageClassWithDefaults instantiates a new OSExternalStorageClassCreateReqExternalStorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetDescription

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalStoragePlatforms

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetExternalStoragePlatforms() []OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt`

GetExternalStoragePlatforms returns the ExternalStoragePlatforms field if non-nil, zero value otherwise.

### GetExternalStoragePlatformsOk

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetExternalStoragePlatformsOk() (*[]OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt, bool)`

GetExternalStoragePlatformsOk returns a tuple with the ExternalStoragePlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoragePlatforms

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetExternalStoragePlatforms(v []OSExternalStorageClassCreateReqExternalStorageClassExternalStoragePlatformsElt)`

SetExternalStoragePlatforms sets ExternalStoragePlatforms field to given value.


### GetName

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetName(v string)`

SetName sets Name field to given value.


### GetOsPolicyId

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetOsPolicyId() int64`

GetOsPolicyId returns the OsPolicyId field if non-nil, zero value otherwise.

### GetOsPolicyIdOk

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetOsPolicyIdOk() (*int64, bool)`

GetOsPolicyIdOk returns a tuple with the OsPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPolicyId

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetOsPolicyId(v int64)`

SetOsPolicyId sets OsPolicyId field to given value.


### GetPlatform

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlatformType

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.


### GetPrefixEnabled

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPrefixEnabled() bool`

GetPrefixEnabled returns the PrefixEnabled field if non-nil, zero value otherwise.

### GetPrefixEnabledOk

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetPrefixEnabledOk() (*bool, bool)`

GetPrefixEnabledOk returns a tuple with the PrefixEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixEnabled

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetPrefixEnabled(v bool)`

SetPrefixEnabled sets PrefixEnabled field to given value.

### HasPrefixEnabled

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) HasPrefixEnabled() bool`

HasPrefixEnabled returns a boolean if a field has been set.

### GetStoragePattern

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetStoragePattern() []string`

GetStoragePattern returns the StoragePattern field if non-nil, zero value otherwise.

### GetStoragePatternOk

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) GetStoragePatternOk() (*[]string, bool)`

GetStoragePatternOk returns a tuple with the StoragePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePattern

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) SetStoragePattern(v []string)`

SetStoragePattern sets StoragePattern field to given value.

### HasStoragePattern

`func (o *OSExternalStorageClassCreateReqExternalStorageClass) HasStoragePattern() bool`

HasStoragePattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


