# ClusterServerTimeResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerTime** | Pointer to **time.Time** | current server time | [optional] 

## Methods

### NewClusterServerTimeResp

`func NewClusterServerTimeResp() *ClusterServerTimeResp`

NewClusterServerTimeResp instantiates a new ClusterServerTimeResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerTimeRespWithDefaults

`func NewClusterServerTimeRespWithDefaults() *ClusterServerTimeResp`

NewClusterServerTimeRespWithDefaults instantiates a new ClusterServerTimeResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerTime

`func (o *ClusterServerTimeResp) GetServerTime() time.Time`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *ClusterServerTimeResp) GetServerTimeOk() (*time.Time, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *ClusterServerTimeResp) SetServerTime(v time.Time)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *ClusterServerTimeResp) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


