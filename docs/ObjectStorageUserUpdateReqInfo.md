# ObjectStorageUserUpdateReqInfo

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
**LocalBucketQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**LocalBucketQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**LocalUserQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**LocalUserQuotaMaxSize** | Pointer to **int64** |  | [optional] 
**LocationConstraintEnabled** | Pointer to **bool** |  | [optional] 
**MaxBuckets** | Pointer to **int64** |  | [optional] 
**OpMask** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **int64** |  | [optional] 
**PolicyPollingEnabled** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to [**OSUserPropertiesReq**](OSUserPropertiesReq.md) |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 
**UserQuotaMaxObjects** | Pointer to **int64** |  | [optional] 
**UserQuotaMaxSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewObjectStorageUserUpdateReqInfo

`func NewObjectStorageUserUpdateReqInfo() *ObjectStorageUserUpdateReqInfo`

NewObjectStorageUserUpdateReqInfo instantiates a new ObjectStorageUserUpdateReqInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageUserUpdateReqInfoWithDefaults

`func NewObjectStorageUserUpdateReqInfoWithDefaults() *ObjectStorageUserUpdateReqInfo`

NewObjectStorageUserUpdateReqInfoWithDefaults instantiates a new ObjectStorageUserUpdateReqInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) GetBucketQuotaMaxObjects() int64`

GetBucketQuotaMaxObjects returns the BucketQuotaMaxObjects field if non-nil, zero value otherwise.

### GetBucketQuotaMaxObjectsOk

`func (o *ObjectStorageUserUpdateReqInfo) GetBucketQuotaMaxObjectsOk() (*int64, bool)`

GetBucketQuotaMaxObjectsOk returns a tuple with the BucketQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) SetBucketQuotaMaxObjects(v int64)`

SetBucketQuotaMaxObjects sets BucketQuotaMaxObjects field to given value.

### HasBucketQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) HasBucketQuotaMaxObjects() bool`

HasBucketQuotaMaxObjects returns a boolean if a field has been set.

### GetBucketQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) GetBucketQuotaMaxSize() int64`

GetBucketQuotaMaxSize returns the BucketQuotaMaxSize field if non-nil, zero value otherwise.

### GetBucketQuotaMaxSizeOk

`func (o *ObjectStorageUserUpdateReqInfo) GetBucketQuotaMaxSizeOk() (*int64, bool)`

GetBucketQuotaMaxSizeOk returns a tuple with the BucketQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) SetBucketQuotaMaxSize(v int64)`

SetBucketQuotaMaxSize sets BucketQuotaMaxSize field to given value.

### HasBucketQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) HasBucketQuotaMaxSize() bool`

HasBucketQuotaMaxSize returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectStorageUserUpdateReqInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectStorageUserUpdateReqInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectStorageUserUpdateReqInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectStorageUserUpdateReqInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ObjectStorageUserUpdateReqInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ObjectStorageUserUpdateReqInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ObjectStorageUserUpdateReqInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ObjectStorageUserUpdateReqInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *ObjectStorageUserUpdateReqInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ObjectStorageUserUpdateReqInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ObjectStorageUserUpdateReqInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ObjectStorageUserUpdateReqInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalBucketQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) GetExternalBucketQuotaMaxObjects() int64`

GetExternalBucketQuotaMaxObjects returns the ExternalBucketQuotaMaxObjects field if non-nil, zero value otherwise.

### GetExternalBucketQuotaMaxObjectsOk

`func (o *ObjectStorageUserUpdateReqInfo) GetExternalBucketQuotaMaxObjectsOk() (*int64, bool)`

GetExternalBucketQuotaMaxObjectsOk returns a tuple with the ExternalBucketQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBucketQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) SetExternalBucketQuotaMaxObjects(v int64)`

SetExternalBucketQuotaMaxObjects sets ExternalBucketQuotaMaxObjects field to given value.

### HasExternalBucketQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) HasExternalBucketQuotaMaxObjects() bool`

HasExternalBucketQuotaMaxObjects returns a boolean if a field has been set.

### GetExternalBucketQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) GetExternalBucketQuotaMaxSize() int64`

GetExternalBucketQuotaMaxSize returns the ExternalBucketQuotaMaxSize field if non-nil, zero value otherwise.

### GetExternalBucketQuotaMaxSizeOk

`func (o *ObjectStorageUserUpdateReqInfo) GetExternalBucketQuotaMaxSizeOk() (*int64, bool)`

GetExternalBucketQuotaMaxSizeOk returns a tuple with the ExternalBucketQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBucketQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) SetExternalBucketQuotaMaxSize(v int64)`

SetExternalBucketQuotaMaxSize sets ExternalBucketQuotaMaxSize field to given value.

### HasExternalBucketQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) HasExternalBucketQuotaMaxSize() bool`

HasExternalBucketQuotaMaxSize returns a boolean if a field has been set.

### GetExternalUserQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) GetExternalUserQuotaMaxObjects() int64`

GetExternalUserQuotaMaxObjects returns the ExternalUserQuotaMaxObjects field if non-nil, zero value otherwise.

### GetExternalUserQuotaMaxObjectsOk

`func (o *ObjectStorageUserUpdateReqInfo) GetExternalUserQuotaMaxObjectsOk() (*int64, bool)`

GetExternalUserQuotaMaxObjectsOk returns a tuple with the ExternalUserQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) SetExternalUserQuotaMaxObjects(v int64)`

SetExternalUserQuotaMaxObjects sets ExternalUserQuotaMaxObjects field to given value.

### HasExternalUserQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) HasExternalUserQuotaMaxObjects() bool`

HasExternalUserQuotaMaxObjects returns a boolean if a field has been set.

### GetExternalUserQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) GetExternalUserQuotaMaxSize() int64`

GetExternalUserQuotaMaxSize returns the ExternalUserQuotaMaxSize field if non-nil, zero value otherwise.

### GetExternalUserQuotaMaxSizeOk

`func (o *ObjectStorageUserUpdateReqInfo) GetExternalUserQuotaMaxSizeOk() (*int64, bool)`

GetExternalUserQuotaMaxSizeOk returns a tuple with the ExternalUserQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) SetExternalUserQuotaMaxSize(v int64)`

SetExternalUserQuotaMaxSize sets ExternalUserQuotaMaxSize field to given value.

### HasExternalUserQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) HasExternalUserQuotaMaxSize() bool`

HasExternalUserQuotaMaxSize returns a boolean if a field has been set.

### GetLocalBucketQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) GetLocalBucketQuotaMaxObjects() int64`

GetLocalBucketQuotaMaxObjects returns the LocalBucketQuotaMaxObjects field if non-nil, zero value otherwise.

### GetLocalBucketQuotaMaxObjectsOk

`func (o *ObjectStorageUserUpdateReqInfo) GetLocalBucketQuotaMaxObjectsOk() (*int64, bool)`

GetLocalBucketQuotaMaxObjectsOk returns a tuple with the LocalBucketQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalBucketQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) SetLocalBucketQuotaMaxObjects(v int64)`

SetLocalBucketQuotaMaxObjects sets LocalBucketQuotaMaxObjects field to given value.

### HasLocalBucketQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) HasLocalBucketQuotaMaxObjects() bool`

HasLocalBucketQuotaMaxObjects returns a boolean if a field has been set.

### GetLocalBucketQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) GetLocalBucketQuotaMaxSize() int64`

GetLocalBucketQuotaMaxSize returns the LocalBucketQuotaMaxSize field if non-nil, zero value otherwise.

### GetLocalBucketQuotaMaxSizeOk

`func (o *ObjectStorageUserUpdateReqInfo) GetLocalBucketQuotaMaxSizeOk() (*int64, bool)`

GetLocalBucketQuotaMaxSizeOk returns a tuple with the LocalBucketQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalBucketQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) SetLocalBucketQuotaMaxSize(v int64)`

SetLocalBucketQuotaMaxSize sets LocalBucketQuotaMaxSize field to given value.

### HasLocalBucketQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) HasLocalBucketQuotaMaxSize() bool`

HasLocalBucketQuotaMaxSize returns a boolean if a field has been set.

### GetLocalUserQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) GetLocalUserQuotaMaxObjects() int64`

GetLocalUserQuotaMaxObjects returns the LocalUserQuotaMaxObjects field if non-nil, zero value otherwise.

### GetLocalUserQuotaMaxObjectsOk

`func (o *ObjectStorageUserUpdateReqInfo) GetLocalUserQuotaMaxObjectsOk() (*int64, bool)`

GetLocalUserQuotaMaxObjectsOk returns a tuple with the LocalUserQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) SetLocalUserQuotaMaxObjects(v int64)`

SetLocalUserQuotaMaxObjects sets LocalUserQuotaMaxObjects field to given value.

### HasLocalUserQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) HasLocalUserQuotaMaxObjects() bool`

HasLocalUserQuotaMaxObjects returns a boolean if a field has been set.

### GetLocalUserQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) GetLocalUserQuotaMaxSize() int64`

GetLocalUserQuotaMaxSize returns the LocalUserQuotaMaxSize field if non-nil, zero value otherwise.

### GetLocalUserQuotaMaxSizeOk

`func (o *ObjectStorageUserUpdateReqInfo) GetLocalUserQuotaMaxSizeOk() (*int64, bool)`

GetLocalUserQuotaMaxSizeOk returns a tuple with the LocalUserQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) SetLocalUserQuotaMaxSize(v int64)`

SetLocalUserQuotaMaxSize sets LocalUserQuotaMaxSize field to given value.

### HasLocalUserQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) HasLocalUserQuotaMaxSize() bool`

HasLocalUserQuotaMaxSize returns a boolean if a field has been set.

### GetLocationConstraintEnabled

`func (o *ObjectStorageUserUpdateReqInfo) GetLocationConstraintEnabled() bool`

GetLocationConstraintEnabled returns the LocationConstraintEnabled field if non-nil, zero value otherwise.

### GetLocationConstraintEnabledOk

`func (o *ObjectStorageUserUpdateReqInfo) GetLocationConstraintEnabledOk() (*bool, bool)`

GetLocationConstraintEnabledOk returns a tuple with the LocationConstraintEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationConstraintEnabled

`func (o *ObjectStorageUserUpdateReqInfo) SetLocationConstraintEnabled(v bool)`

SetLocationConstraintEnabled sets LocationConstraintEnabled field to given value.

### HasLocationConstraintEnabled

`func (o *ObjectStorageUserUpdateReqInfo) HasLocationConstraintEnabled() bool`

HasLocationConstraintEnabled returns a boolean if a field has been set.

### GetMaxBuckets

`func (o *ObjectStorageUserUpdateReqInfo) GetMaxBuckets() int64`

GetMaxBuckets returns the MaxBuckets field if non-nil, zero value otherwise.

### GetMaxBucketsOk

`func (o *ObjectStorageUserUpdateReqInfo) GetMaxBucketsOk() (*int64, bool)`

GetMaxBucketsOk returns a tuple with the MaxBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBuckets

`func (o *ObjectStorageUserUpdateReqInfo) SetMaxBuckets(v int64)`

SetMaxBuckets sets MaxBuckets field to given value.

### HasMaxBuckets

`func (o *ObjectStorageUserUpdateReqInfo) HasMaxBuckets() bool`

HasMaxBuckets returns a boolean if a field has been set.

### GetOpMask

`func (o *ObjectStorageUserUpdateReqInfo) GetOpMask() string`

GetOpMask returns the OpMask field if non-nil, zero value otherwise.

### GetOpMaskOk

`func (o *ObjectStorageUserUpdateReqInfo) GetOpMaskOk() (*string, bool)`

GetOpMaskOk returns a tuple with the OpMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpMask

`func (o *ObjectStorageUserUpdateReqInfo) SetOpMask(v string)`

SetOpMask sets OpMask field to given value.

### HasOpMask

`func (o *ObjectStorageUserUpdateReqInfo) HasOpMask() bool`

HasOpMask returns a boolean if a field has been set.

### GetPolicyId

`func (o *ObjectStorageUserUpdateReqInfo) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ObjectStorageUserUpdateReqInfo) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ObjectStorageUserUpdateReqInfo) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ObjectStorageUserUpdateReqInfo) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyPollingEnabled

`func (o *ObjectStorageUserUpdateReqInfo) GetPolicyPollingEnabled() bool`

GetPolicyPollingEnabled returns the PolicyPollingEnabled field if non-nil, zero value otherwise.

### GetPolicyPollingEnabledOk

`func (o *ObjectStorageUserUpdateReqInfo) GetPolicyPollingEnabledOk() (*bool, bool)`

GetPolicyPollingEnabledOk returns a tuple with the PolicyPollingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyPollingEnabled

`func (o *ObjectStorageUserUpdateReqInfo) SetPolicyPollingEnabled(v bool)`

SetPolicyPollingEnabled sets PolicyPollingEnabled field to given value.

### HasPolicyPollingEnabled

`func (o *ObjectStorageUserUpdateReqInfo) HasPolicyPollingEnabled() bool`

HasPolicyPollingEnabled returns a boolean if a field has been set.

### GetProperties

`func (o *ObjectStorageUserUpdateReqInfo) GetProperties() OSUserPropertiesReq`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ObjectStorageUserUpdateReqInfo) GetPropertiesOk() (*OSUserPropertiesReq, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ObjectStorageUserUpdateReqInfo) SetProperties(v OSUserPropertiesReq)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ObjectStorageUserUpdateReqInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSuspended

`func (o *ObjectStorageUserUpdateReqInfo) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ObjectStorageUserUpdateReqInfo) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ObjectStorageUserUpdateReqInfo) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ObjectStorageUserUpdateReqInfo) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetUserQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) GetUserQuotaMaxObjects() int64`

GetUserQuotaMaxObjects returns the UserQuotaMaxObjects field if non-nil, zero value otherwise.

### GetUserQuotaMaxObjectsOk

`func (o *ObjectStorageUserUpdateReqInfo) GetUserQuotaMaxObjectsOk() (*int64, bool)`

GetUserQuotaMaxObjectsOk returns a tuple with the UserQuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) SetUserQuotaMaxObjects(v int64)`

SetUserQuotaMaxObjects sets UserQuotaMaxObjects field to given value.

### HasUserQuotaMaxObjects

`func (o *ObjectStorageUserUpdateReqInfo) HasUserQuotaMaxObjects() bool`

HasUserQuotaMaxObjects returns a boolean if a field has been set.

### GetUserQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) GetUserQuotaMaxSize() int64`

GetUserQuotaMaxSize returns the UserQuotaMaxSize field if non-nil, zero value otherwise.

### GetUserQuotaMaxSizeOk

`func (o *ObjectStorageUserUpdateReqInfo) GetUserQuotaMaxSizeOk() (*int64, bool)`

GetUserQuotaMaxSizeOk returns a tuple with the UserQuotaMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) SetUserQuotaMaxSize(v int64)`

SetUserQuotaMaxSize sets UserQuotaMaxSize field to given value.

### HasUserQuotaMaxSize

`func (o *ObjectStorageUserUpdateReqInfo) HasUserQuotaMaxSize() bool`

HasUserQuotaMaxSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


