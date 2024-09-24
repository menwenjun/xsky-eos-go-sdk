# ObjectStorageUserCreateReqInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**BucketQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ExternalBucketQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**ExternalBucketQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**ExternalUserQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**ExternalUserQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**Keys** | Pointer to [**[]ObjectStorageKey**](ObjectStorageKey.md) |  | [optional] 
**LocalBucketQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**LocalBucketQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**LocalUserQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**LocalUserQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**LocationConstraintEnabled** | Pointer to **bool** |  | [optional] 
**MaxBuckets** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**OpMask** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **int64** |  | [optional] 
**PolicyPollingEnabled** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to [**OSUserPropertiesReq**](OSUserPropertiesReq.md) |  | [optional] 
**UserQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**UserQuotaMaxSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewObjectStorageUserCreateReqInfo

`func NewObjectStorageUserCreateReqInfo(name string, ) *ObjectStorageUserCreateReqInfo`

NewObjectStorageUserCreateReqInfo instantiates a new ObjectStorageUserCreateReqInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageUserCreateReqInfoWithDefaults

`func NewObjectStorageUserCreateReqInfoWithDefaults() *ObjectStorageUserCreateReqInfo`

NewObjectStorageUserCreateReqInfoWithDefaults instantiates a new ObjectStorageUserCreateReqInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) GetBucketQuotaMaxObjects() int64`

GetBucketQuotaMaxObjects returns the BucketQuotaMaxObjects field if non-nil, zero value otherwise.

### GetBucketQuotaMaxObjectsOk

`func (o *ObjectStorageUserCreateReqInfo) GetBucketQuotaMaxObjectsOk() (*int64, bool)`

GetBucketQuotaMaxObjectsOk returns a tuple with the BucketQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) SetBucketQuotaMaxObjects(v int64)`

SetBucketQuotaMaxObjects sets BucketQuotaMaxObjects field to given value.

### HasBucketQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) HasBucketQuotaMaxObjects() bool`

HasBucketQuotaMaxObjects returns a boolean if a field has been set.

### GetBucketQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) GetBucketQuotaMaxSize() int64`

GetBucketQuotaMaxSize returns the BucketQuotaMaxSize field if non-nil, zero value otherwise.

### GetBucketQuotaMaxSizeOk

`func (o *ObjectStorageUserCreateReqInfo) GetBucketQuotaMaxSizeOk() (*int64, bool)`

GetBucketQuotaMaxSizeOk returns a tuple with the BucketQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) SetBucketQuotaMaxSize(v int64)`

SetBucketQuotaMaxSize sets BucketQuotaMaxSize field to given value.

### HasBucketQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) HasBucketQuotaMaxSize() bool`

HasBucketQuotaMaxSize returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStorageUserCreateReqInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStorageUserCreateReqInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStorageUserCreateReqInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStorageUserCreateReqInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ObjectStorageUserCreateReqInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ObjectStorageUserCreateReqInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ObjectStorageUserCreateReqInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ObjectStorageUserCreateReqInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *ObjectStorageUserCreateReqInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ObjectStorageUserCreateReqInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ObjectStorageUserCreateReqInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ObjectStorageUserCreateReqInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalBucketQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) GetExternalBucketQuotaMaxObjects() int64`

GetExternalBucketQuotaMaxObjects returns the ExternalBucketQuotaMaxObjects field if non-nil, zero value otherwise.

### GetExternalBucketQuotaMaxObjectsOk

`func (o *ObjectStorageUserCreateReqInfo) GetExternalBucketQuotaMaxObjectsOk() (*int64, bool)`

GetExternalBucketQuotaMaxObjectsOk returns a tuple with the ExternalBucketQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBucketQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) SetExternalBucketQuotaMaxObjects(v int64)`

SetExternalBucketQuotaMaxObjects sets ExternalBucketQuotaMaxObjects field to given value.

### HasExternalBucketQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) HasExternalBucketQuotaMaxObjects() bool`

HasExternalBucketQuotaMaxObjects returns a boolean if a field has been set.

### GetExternalBucketQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) GetExternalBucketQuotaMaxSize() int64`

GetExternalBucketQuotaMaxSize returns the ExternalBucketQuotaMaxSize field if non-nil, zero value otherwise.

### GetExternalBucketQuotaMaxSizeOk

`func (o *ObjectStorageUserCreateReqInfo) GetExternalBucketQuotaMaxSizeOk() (*int64, bool)`

GetExternalBucketQuotaMaxSizeOk returns a tuple with the ExternalBucketQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBucketQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) SetExternalBucketQuotaMaxSize(v int64)`

SetExternalBucketQuotaMaxSize sets ExternalBucketQuotaMaxSize field to given value.

### HasExternalBucketQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) HasExternalBucketQuotaMaxSize() bool`

HasExternalBucketQuotaMaxSize returns a boolean if a field has been set.

### GetExternalUserQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) GetExternalUserQuotaMaxObjects() int64`

GetExternalUserQuotaMaxObjects returns the ExternalUserQuotaMaxObjects field if non-nil, zero value otherwise.

### GetExternalUserQuotaMaxObjectsOk

`func (o *ObjectStorageUserCreateReqInfo) GetExternalUserQuotaMaxObjectsOk() (*int64, bool)`

GetExternalUserQuotaMaxObjectsOk returns a tuple with the ExternalUserQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) SetExternalUserQuotaMaxObjects(v int64)`

SetExternalUserQuotaMaxObjects sets ExternalUserQuotaMaxObjects field to given value.

### HasExternalUserQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) HasExternalUserQuotaMaxObjects() bool`

HasExternalUserQuotaMaxObjects returns a boolean if a field has been set.

### GetExternalUserQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) GetExternalUserQuotaMaxSize() int64`

GetExternalUserQuotaMaxSize returns the ExternalUserQuotaMaxSize field if non-nil, zero value otherwise.

### GetExternalUserQuotaMaxSizeOk

`func (o *ObjectStorageUserCreateReqInfo) GetExternalUserQuotaMaxSizeOk() (*int64, bool)`

GetExternalUserQuotaMaxSizeOk returns a tuple with the ExternalUserQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) SetExternalUserQuotaMaxSize(v int64)`

SetExternalUserQuotaMaxSize sets ExternalUserQuotaMaxSize field to given value.

### HasExternalUserQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) HasExternalUserQuotaMaxSize() bool`

HasExternalUserQuotaMaxSize returns a boolean if a field has been set.

### GetKeys

`func (o *ObjectStorageUserCreateReqInfo) GetKeys() []ObjectStorageKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ObjectStorageUserCreateReqInfo) GetKeysOk() (*[]ObjectStorageKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ObjectStorageUserCreateReqInfo) SetKeys(v []ObjectStorageKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ObjectStorageUserCreateReqInfo) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetLocalBucketQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) GetLocalBucketQuotaMaxObjects() int64`

GetLocalBucketQuotaMaxObjects returns the LocalBucketQuotaMaxObjects field if non-nil, zero value otherwise.

### GetLocalBucketQuotaMaxObjectsOk

`func (o *ObjectStorageUserCreateReqInfo) GetLocalBucketQuotaMaxObjectsOk() (*int64, bool)`

GetLocalBucketQuotaMaxObjectsOk returns a tuple with the LocalBucketQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalBucketQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) SetLocalBucketQuotaMaxObjects(v int64)`

SetLocalBucketQuotaMaxObjects sets LocalBucketQuotaMaxObjects field to given value.

### HasLocalBucketQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) HasLocalBucketQuotaMaxObjects() bool`

HasLocalBucketQuotaMaxObjects returns a boolean if a field has been set.

### GetLocalBucketQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) GetLocalBucketQuotaMaxSize() int64`

GetLocalBucketQuotaMaxSize returns the LocalBucketQuotaMaxSize field if non-nil, zero value otherwise.

### GetLocalBucketQuotaMaxSizeOk

`func (o *ObjectStorageUserCreateReqInfo) GetLocalBucketQuotaMaxSizeOk() (*int64, bool)`

GetLocalBucketQuotaMaxSizeOk returns a tuple with the LocalBucketQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalBucketQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) SetLocalBucketQuotaMaxSize(v int64)`

SetLocalBucketQuotaMaxSize sets LocalBucketQuotaMaxSize field to given value.

### HasLocalBucketQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) HasLocalBucketQuotaMaxSize() bool`

HasLocalBucketQuotaMaxSize returns a boolean if a field has been set.

### GetLocalUserQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) GetLocalUserQuotaMaxObjects() int64`

GetLocalUserQuotaMaxObjects returns the LocalUserQuotaMaxObjects field if non-nil, zero value otherwise.

### GetLocalUserQuotaMaxObjectsOk

`func (o *ObjectStorageUserCreateReqInfo) GetLocalUserQuotaMaxObjectsOk() (*int64, bool)`

GetLocalUserQuotaMaxObjectsOk returns a tuple with the LocalUserQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) SetLocalUserQuotaMaxObjects(v int64)`

SetLocalUserQuotaMaxObjects sets LocalUserQuotaMaxObjects field to given value.

### HasLocalUserQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) HasLocalUserQuotaMaxObjects() bool`

HasLocalUserQuotaMaxObjects returns a boolean if a field has been set.

### GetLocalUserQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) GetLocalUserQuotaMaxSize() int64`

GetLocalUserQuotaMaxSize returns the LocalUserQuotaMaxSize field if non-nil, zero value otherwise.

### GetLocalUserQuotaMaxSizeOk

`func (o *ObjectStorageUserCreateReqInfo) GetLocalUserQuotaMaxSizeOk() (*int64, bool)`

GetLocalUserQuotaMaxSizeOk returns a tuple with the LocalUserQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) SetLocalUserQuotaMaxSize(v int64)`

SetLocalUserQuotaMaxSize sets LocalUserQuotaMaxSize field to given value.

### HasLocalUserQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) HasLocalUserQuotaMaxSize() bool`

HasLocalUserQuotaMaxSize returns a boolean if a field has been set.

### GetLocationConstraintEnabled

`func (o *ObjectStorageUserCreateReqInfo) GetLocationConstraintEnabled() bool`

GetLocationConstraintEnabled returns the LocationConstraintEnabled field if non-nil, zero value otherwise.

### GetLocationConstraintEnabledOk

`func (o *ObjectStorageUserCreateReqInfo) GetLocationConstraintEnabledOk() (*bool, bool)`

GetLocationConstraintEnabledOk returns a tuple with the LocationConstraintEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationConstraintEnabled

`func (o *ObjectStorageUserCreateReqInfo) SetLocationConstraintEnabled(v bool)`

SetLocationConstraintEnabled sets LocationConstraintEnabled field to given value.

### HasLocationConstraintEnabled

`func (o *ObjectStorageUserCreateReqInfo) HasLocationConstraintEnabled() bool`

HasLocationConstraintEnabled returns a boolean if a field has been set.

### GetMaxBuckets

`func (o *ObjectStorageUserCreateReqInfo) GetMaxBuckets() int64`

GetMaxBuckets returns the MaxBuckets field if non-nil, zero value otherwise.

### GetMaxBucketsOk

`func (o *ObjectStorageUserCreateReqInfo) GetMaxBucketsOk() (*int64, bool)`

GetMaxBucketsOk returns a tuple with the MaxBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBuckets

`func (o *ObjectStorageUserCreateReqInfo) SetMaxBuckets(v int64)`

SetMaxBuckets sets MaxBuckets field to given value.

### HasMaxBuckets

`func (o *ObjectStorageUserCreateReqInfo) HasMaxBuckets() bool`

HasMaxBuckets returns a boolean if a field has been set.

### GetName

`func (o *ObjectStorageUserCreateReqInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStorageUserCreateReqInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStorageUserCreateReqInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOpMask

`func (o *ObjectStorageUserCreateReqInfo) GetOpMask() string`

GetOpMask returns the OpMask field if non-nil, zero value otherwise.

### GetOpMaskOk

`func (o *ObjectStorageUserCreateReqInfo) GetOpMaskOk() (*string, bool)`

GetOpMaskOk returns a tuple with the OpMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpMask

`func (o *ObjectStorageUserCreateReqInfo) SetOpMask(v string)`

SetOpMask sets OpMask field to given value.

### HasOpMask

`func (o *ObjectStorageUserCreateReqInfo) HasOpMask() bool`

HasOpMask returns a boolean if a field has been set.

### GetPolicyId

`func (o *ObjectStorageUserCreateReqInfo) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ObjectStorageUserCreateReqInfo) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ObjectStorageUserCreateReqInfo) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ObjectStorageUserCreateReqInfo) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyPollingEnabled

`func (o *ObjectStorageUserCreateReqInfo) GetPolicyPollingEnabled() bool`

GetPolicyPollingEnabled returns the PolicyPollingEnabled field if non-nil, zero value otherwise.

### GetPolicyPollingEnabledOk

`func (o *ObjectStorageUserCreateReqInfo) GetPolicyPollingEnabledOk() (*bool, bool)`

GetPolicyPollingEnabledOk returns a tuple with the PolicyPollingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyPollingEnabled

`func (o *ObjectStorageUserCreateReqInfo) SetPolicyPollingEnabled(v bool)`

SetPolicyPollingEnabled sets PolicyPollingEnabled field to given value.

### HasPolicyPollingEnabled

`func (o *ObjectStorageUserCreateReqInfo) HasPolicyPollingEnabled() bool`

HasPolicyPollingEnabled returns a boolean if a field has been set.

### GetProperties

`func (o *ObjectStorageUserCreateReqInfo) GetProperties() OSUserPropertiesReq`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ObjectStorageUserCreateReqInfo) GetPropertiesOk() (*OSUserPropertiesReq, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ObjectStorageUserCreateReqInfo) SetProperties(v OSUserPropertiesReq)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ObjectStorageUserCreateReqInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetUserQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) GetUserQuotaMaxObjects() int64`

GetUserQuotaMaxObjects returns the UserQuotaMaxObjects field if non-nil, zero value otherwise.

### GetUserQuotaMaxObjectsOk

`func (o *ObjectStorageUserCreateReqInfo) GetUserQuotaMaxObjectsOk() (*int64, bool)`

GetUserQuotaMaxObjectsOk returns a tuple with the UserQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) SetUserQuotaMaxObjects(v int64)`

SetUserQuotaMaxObjects sets UserQuotaMaxObjects field to given value.

### HasUserQuotaMaxObjects

`func (o *ObjectStorageUserCreateReqInfo) HasUserQuotaMaxObjects() bool`

HasUserQuotaMaxObjects returns a boolean if a field has been set.

### GetUserQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) GetUserQuotaMaxSize() int64`

GetUserQuotaMaxSize returns the UserQuotaMaxSize field if non-nil, zero value otherwise.

### GetUserQuotaMaxSizeOk

`func (o *ObjectStorageUserCreateReqInfo) GetUserQuotaMaxSizeOk() (*int64, bool)`

GetUserQuotaMaxSizeOk returns a tuple with the UserQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) SetUserQuotaMaxSize(v int64)`

SetUserQuotaMaxSize sets UserQuotaMaxSize field to given value.

### HasUserQuotaMaxSize

`func (o *ObjectStorageUserCreateReqInfo) HasUserQuotaMaxSize() bool`

HasUserQuotaMaxSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


