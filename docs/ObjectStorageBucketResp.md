# ObjectStorageBucketResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsBucket** | [**ObjectStorageBucketRecord**](ObjectStorageBucketRecord.md) |  | 

## Methods

### NewObjectStorageBucketResp

`func NewObjectStorageBucketResp(osBucket ObjectStorageBucketRecord, ) *ObjectStorageBucketResp`

NewObjectStorageBucketResp instantiates a new ObjectStorageBucketResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketRespWithDefaults

`func NewObjectStorageBucketRespWithDefaults() *ObjectStorageBucketResp`

NewObjectStorageBucketRespWithDefaults instantiates a new ObjectStorageBucketResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsBucket

`func (o *ObjectStorageBucketResp) GetOsBucket() ObjectStorageBucketRecord`

GetOsBucket returns the OsBucket field if non-nil, zero value otherwise.

### GetOsBucketOk

`func (o *ObjectStorageBucketResp) GetOsBucketOk() (*ObjectStorageBucketRecord, bool)`

GetOsBucketOk returns a tuple with the OsBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBucket

`func (o *ObjectStorageBucketResp) SetOsBucket(v ObjectStorageBucketRecord)`

SetOsBucket sets OsBucket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


