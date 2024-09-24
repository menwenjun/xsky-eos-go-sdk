# S3Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketLimit** | Pointer to **int64** |  | [optional] 
**BucketPath** | Pointer to **string** |  | [optional] 
**BucketPermission** | Pointer to **string** |  | [optional] 
**GatewayGroupId** | Pointer to **int64** |  | [optional] 
**S3Enabled** | Pointer to **bool** |  | [optional] 
**S3Keys** | Pointer to [**[]S3Key**](S3Key.md) |  | [optional] 

## Methods

### NewS3Req

`func NewS3Req() *S3Req`

NewS3Req instantiates a new S3Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ReqWithDefaults

`func NewS3ReqWithDefaults() *S3Req`

NewS3ReqWithDefaults instantiates a new S3Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketLimit

`func (o *S3Req) GetBucketLimit() int64`

GetBucketLimit returns the BucketLimit field if non-nil, zero value otherwise.

### GetBucketLimitOk

`func (o *S3Req) GetBucketLimitOk() (*int64, bool)`

GetBucketLimitOk returns a tuple with the BucketLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketLimit

`func (o *S3Req) SetBucketLimit(v int64)`

SetBucketLimit sets BucketLimit field to given value.

### HasBucketLimit

`func (o *S3Req) HasBucketLimit() bool`

HasBucketLimit returns a boolean if a field has been set.

### GetBucketPath

`func (o *S3Req) GetBucketPath() string`

GetBucketPath returns the BucketPath field if non-nil, zero value otherwise.

### GetBucketPathOk

`func (o *S3Req) GetBucketPathOk() (*string, bool)`

GetBucketPathOk returns a tuple with the BucketPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPath

`func (o *S3Req) SetBucketPath(v string)`

SetBucketPath sets BucketPath field to given value.

### HasBucketPath

`func (o *S3Req) HasBucketPath() bool`

HasBucketPath returns a boolean if a field has been set.

### GetBucketPermission

`func (o *S3Req) GetBucketPermission() string`

GetBucketPermission returns the BucketPermission field if non-nil, zero value otherwise.

### GetBucketPermissionOk

`func (o *S3Req) GetBucketPermissionOk() (*string, bool)`

GetBucketPermissionOk returns a tuple with the BucketPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPermission

`func (o *S3Req) SetBucketPermission(v string)`

SetBucketPermission sets BucketPermission field to given value.

### HasBucketPermission

`func (o *S3Req) HasBucketPermission() bool`

HasBucketPermission returns a boolean if a field has been set.

### GetGatewayGroupId

`func (o *S3Req) GetGatewayGroupId() int64`

GetGatewayGroupId returns the GatewayGroupId field if non-nil, zero value otherwise.

### GetGatewayGroupIdOk

`func (o *S3Req) GetGatewayGroupIdOk() (*int64, bool)`

GetGatewayGroupIdOk returns a tuple with the GatewayGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayGroupId

`func (o *S3Req) SetGatewayGroupId(v int64)`

SetGatewayGroupId sets GatewayGroupId field to given value.

### HasGatewayGroupId

`func (o *S3Req) HasGatewayGroupId() bool`

HasGatewayGroupId returns a boolean if a field has been set.

### GetS3Enabled

`func (o *S3Req) GetS3Enabled() bool`

GetS3Enabled returns the S3Enabled field if non-nil, zero value otherwise.

### GetS3EnabledOk

`func (o *S3Req) GetS3EnabledOk() (*bool, bool)`

GetS3EnabledOk returns a tuple with the S3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Enabled

`func (o *S3Req) SetS3Enabled(v bool)`

SetS3Enabled sets S3Enabled field to given value.

### HasS3Enabled

`func (o *S3Req) HasS3Enabled() bool`

HasS3Enabled returns a boolean if a field has been set.

### GetS3Keys

`func (o *S3Req) GetS3Keys() []S3Key`

GetS3Keys returns the S3Keys field if non-nil, zero value otherwise.

### GetS3KeysOk

`func (o *S3Req) GetS3KeysOk() (*[]S3Key, bool)`

GetS3KeysOk returns a tuple with the S3Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Keys

`func (o *S3Req) SetS3Keys(v []S3Key)`

SetS3Keys sets S3Keys field to given value.

### HasS3Keys

`func (o *S3Req) HasS3Keys() bool`

HasS3Keys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


