# OSBucketQosUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxBandwidthBytes** | **int64** |  | 
**RequestsPerSecond** | **int64** |  | 
**SuddenBandwidthBytes** | Pointer to **int64** |  | [optional] 

## Methods

### NewOSBucketQosUserRequest

`func NewOSBucketQosUserRequest(maxBandwidthBytes int64, requestsPerSecond int64, ) *OSBucketQosUserRequest`

NewOSBucketQosUserRequest instantiates a new OSBucketQosUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSBucketQosUserRequestWithDefaults

`func NewOSBucketQosUserRequestWithDefaults() *OSBucketQosUserRequest`

NewOSBucketQosUserRequestWithDefaults instantiates a new OSBucketQosUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxBandwidthBytes

`func (o *OSBucketQosUserRequest) GetMaxBandwidthBytes() int64`

GetMaxBandwidthBytes returns the MaxBandwidthBytes field if non-nil, zero value otherwise.

### GetMaxBandwidthBytesOk

`func (o *OSBucketQosUserRequest) GetMaxBandwidthBytesOk() (*int64, bool)`

GetMaxBandwidthBytesOk returns a tuple with the MaxBandwidthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBandwidthBytes

`func (o *OSBucketQosUserRequest) SetMaxBandwidthBytes(v int64)`

SetMaxBandwidthBytes sets MaxBandwidthBytes field to given value.


### GetRequestsPerSecond

`func (o *OSBucketQosUserRequest) GetRequestsPerSecond() int64`

GetRequestsPerSecond returns the RequestsPerSecond field if non-nil, zero value otherwise.

### GetRequestsPerSecondOk

`func (o *OSBucketQosUserRequest) GetRequestsPerSecondOk() (*int64, bool)`

GetRequestsPerSecondOk returns a tuple with the RequestsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPerSecond

`func (o *OSBucketQosUserRequest) SetRequestsPerSecond(v int64)`

SetRequestsPerSecond sets RequestsPerSecond field to given value.


### GetSuddenBandwidthBytes

`func (o *OSBucketQosUserRequest) GetSuddenBandwidthBytes() int64`

GetSuddenBandwidthBytes returns the SuddenBandwidthBytes field if non-nil, zero value otherwise.

### GetSuddenBandwidthBytesOk

`func (o *OSBucketQosUserRequest) GetSuddenBandwidthBytesOk() (*int64, bool)`

GetSuddenBandwidthBytesOk returns a tuple with the SuddenBandwidthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuddenBandwidthBytes

`func (o *OSBucketQosUserRequest) SetSuddenBandwidthBytes(v int64)`

SetSuddenBandwidthBytes sets SuddenBandwidthBytes field to given value.

### HasSuddenBandwidthBytes

`func (o *OSBucketQosUserRequest) HasSuddenBandwidthBytes() bool`

HasSuddenBandwidthBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


