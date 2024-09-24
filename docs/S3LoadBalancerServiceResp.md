# S3LoadBalancerServiceResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneIds** | **[]int64** | s3 load balancer groups | 

## Methods

### NewS3LoadBalancerServiceResp

`func NewS3LoadBalancerServiceResp(zoneIds []int64, ) *S3LoadBalancerServiceResp`

NewS3LoadBalancerServiceResp instantiates a new S3LoadBalancerServiceResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LoadBalancerServiceRespWithDefaults

`func NewS3LoadBalancerServiceRespWithDefaults() *S3LoadBalancerServiceResp`

NewS3LoadBalancerServiceRespWithDefaults instantiates a new S3LoadBalancerServiceResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneIds

`func (o *S3LoadBalancerServiceResp) GetZoneIds() []int64`

GetZoneIds returns the ZoneIds field if non-nil, zero value otherwise.

### GetZoneIdsOk

`func (o *S3LoadBalancerServiceResp) GetZoneIdsOk() (*[]int64, bool)`

GetZoneIdsOk returns a tuple with the ZoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIds

`func (o *S3LoadBalancerServiceResp) SetZoneIds(v []int64)`

SetZoneIds sets ZoneIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


