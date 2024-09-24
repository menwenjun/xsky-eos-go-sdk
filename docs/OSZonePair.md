# OSZonePair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **int64** |  | [optional] 
**ClockDiff** | Pointer to **int64** |  | [optional] 
**Cluster** | Pointer to [**ClusterNestview**](ClusterNestview.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**SourceZoneUuid** | Pointer to **string** |  | [optional] 
**TargetZoneUuid** | Pointer to **string** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOSZonePair

`func NewOSZonePair() *OSZonePair`

NewOSZonePair instantiates a new OSZonePair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSZonePairWithDefaults

`func NewOSZonePairWithDefaults() *OSZonePair`

NewOSZonePairWithDefaults instantiates a new OSZonePair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *OSZonePair) GetID() int64`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *OSZonePair) GetIDOk() (*int64, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *OSZonePair) SetID(v int64)`

SetID sets ID field to given value.

### HasID

`func (o *OSZonePair) HasID() bool`

HasID returns a boolean if a field has been set.

### GetClockDiff

`func (o *OSZonePair) GetClockDiff() int64`

GetClockDiff returns the ClockDiff field if non-nil, zero value otherwise.

### GetClockDiffOk

`func (o *OSZonePair) GetClockDiffOk() (*int64, bool)`

GetClockDiffOk returns a tuple with the ClockDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockDiff

`func (o *OSZonePair) SetClockDiff(v int64)`

SetClockDiff sets ClockDiff field to given value.

### HasClockDiff

`func (o *OSZonePair) HasClockDiff() bool`

HasClockDiff returns a boolean if a field has been set.

### GetCluster

`func (o *OSZonePair) GetCluster() ClusterNestview`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OSZonePair) GetClusterOk() (*ClusterNestview, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OSZonePair) SetCluster(v ClusterNestview)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OSZonePair) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *OSZonePair) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OSZonePair) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OSZonePair) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OSZonePair) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetSourceZoneUuid

`func (o *OSZonePair) GetSourceZoneUuid() string`

GetSourceZoneUuid returns the SourceZoneUuid field if non-nil, zero value otherwise.

### GetSourceZoneUuidOk

`func (o *OSZonePair) GetSourceZoneUuidOk() (*string, bool)`

GetSourceZoneUuidOk returns a tuple with the SourceZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceZoneUuid

`func (o *OSZonePair) SetSourceZoneUuid(v string)`

SetSourceZoneUuid sets SourceZoneUuid field to given value.

### HasSourceZoneUuid

`func (o *OSZonePair) HasSourceZoneUuid() bool`

HasSourceZoneUuid returns a boolean if a field has been set.

### GetTargetZoneUuid

`func (o *OSZonePair) GetTargetZoneUuid() string`

GetTargetZoneUuid returns the TargetZoneUuid field if non-nil, zero value otherwise.

### GetTargetZoneUuidOk

`func (o *OSZonePair) GetTargetZoneUuidOk() (*string, bool)`

GetTargetZoneUuidOk returns a tuple with the TargetZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetZoneUuid

`func (o *OSZonePair) SetTargetZoneUuid(v string)`

SetTargetZoneUuid sets TargetZoneUuid field to given value.

### HasTargetZoneUuid

`func (o *OSZonePair) HasTargetZoneUuid() bool`

HasTargetZoneUuid returns a boolean if a field has been set.

### GetUpdate

`func (o *OSZonePair) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *OSZonePair) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *OSZonePair) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *OSZonePair) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


