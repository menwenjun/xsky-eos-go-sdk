# ObjectStorageBucketsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsBuckets** | [**[]ObjectStorageBucketRecord**](ObjectStorageBucketRecord.md) | records of object storage bucket | 

## Methods

### NewObjectStorageBucketsResp

`func NewObjectStorageBucketsResp(osBuckets []ObjectStorageBucketRecord, ) *ObjectStorageBucketsResp`

NewObjectStorageBucketsResp instantiates a new ObjectStorageBucketsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageBucketsRespWithDefaults

`func NewObjectStorageBucketsRespWithDefaults() *ObjectStorageBucketsResp`

NewObjectStorageBucketsRespWithDefaults instantiates a new ObjectStorageBucketsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsBuckets

`func (o *ObjectStorageBucketsResp) GetOsBuckets() []ObjectStorageBucketRecord`

GetOsBuckets returns the OsBuckets field if non-nil, zero value otherwise.

### GetOsBucketsOk

`func (o *ObjectStorageBucketsResp) GetOsBucketsOk() (*[]ObjectStorageBucketRecord, bool)`

GetOsBucketsOk returns a tuple with the OsBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBuckets

`func (o *ObjectStorageBucketsResp) SetOsBuckets(v []ObjectStorageBucketRecord)`

SetOsBuckets sets OsBuckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


