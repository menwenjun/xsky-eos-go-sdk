# OSBucketObjectStorageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchingInfo** | Pointer to [**RuleMatchingInfo**](RuleMatchingInfo.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**XAmzStorageClassEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewOSBucketObjectStorageClass

`func NewOSBucketObjectStorageClass() *OSBucketObjectStorageClass`

NewOSBucketObjectStorageClass instantiates a new OSBucketObjectStorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSBucketObjectStorageClassWithDefaults

`func NewOSBucketObjectStorageClassWithDefaults() *OSBucketObjectStorageClass`

NewOSBucketObjectStorageClassWithDefaults instantiates a new OSBucketObjectStorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchingInfo

`func (o *OSBucketObjectStorageClass) GetMatchingInfo() RuleMatchingInfo`

GetMatchingInfo returns the MatchingInfo field if non-nil, zero value otherwise.

### GetMatchingInfoOk

`func (o *OSBucketObjectStorageClass) GetMatchingInfoOk() (*RuleMatchingInfo, bool)`

GetMatchingInfoOk returns a tuple with the MatchingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingInfo

`func (o *OSBucketObjectStorageClass) SetMatchingInfo(v RuleMatchingInfo)`

SetMatchingInfo sets MatchingInfo field to given value.

### HasMatchingInfo

`func (o *OSBucketObjectStorageClass) HasMatchingInfo() bool`

HasMatchingInfo returns a boolean if a field has been set.

### GetType

`func (o *OSBucketObjectStorageClass) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OSBucketObjectStorageClass) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OSBucketObjectStorageClass) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OSBucketObjectStorageClass) HasType() bool`

HasType returns a boolean if a field has been set.

### GetXAmzStorageClassEnabled

`func (o *OSBucketObjectStorageClass) GetXAmzStorageClassEnabled() bool`

GetXAmzStorageClassEnabled returns the XAmzStorageClassEnabled field if non-nil, zero value otherwise.

### GetXAmzStorageClassEnabledOk

`func (o *OSBucketObjectStorageClass) GetXAmzStorageClassEnabledOk() (*bool, bool)`

GetXAmzStorageClassEnabledOk returns a tuple with the XAmzStorageClassEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAmzStorageClassEnabled

`func (o *OSBucketObjectStorageClass) SetXAmzStorageClassEnabled(v bool)`

SetXAmzStorageClassEnabled sets XAmzStorageClassEnabled field to given value.

### HasXAmzStorageClassEnabled

`func (o *OSBucketObjectStorageClass) HasXAmzStorageClassEnabled() bool`

HasXAmzStorageClassEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


