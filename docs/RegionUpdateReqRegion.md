# RegionUpdateReqRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | retion name | 
**StatsReservedDays** | Pointer to **int64** | region stats reserved day | [optional] 

## Methods

### NewRegionUpdateReqRegion

`func NewRegionUpdateReqRegion(name string, ) *RegionUpdateReqRegion`

NewRegionUpdateReqRegion instantiates a new RegionUpdateReqRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionUpdateReqRegionWithDefaults

`func NewRegionUpdateReqRegionWithDefaults() *RegionUpdateReqRegion`

NewRegionUpdateReqRegionWithDefaults instantiates a new RegionUpdateReqRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegionUpdateReqRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegionUpdateReqRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegionUpdateReqRegion) SetName(v string)`

SetName sets Name field to given value.


### GetStatsReservedDays

`func (o *RegionUpdateReqRegion) GetStatsReservedDays() int64`

GetStatsReservedDays returns the StatsReservedDays field if non-nil, zero value otherwise.

### GetStatsReservedDaysOk

`func (o *RegionUpdateReqRegion) GetStatsReservedDaysOk() (*int64, bool)`

GetStatsReservedDaysOk returns a tuple with the StatsReservedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsReservedDays

`func (o *RegionUpdateReqRegion) SetStatsReservedDays(v int64)`

SetStatsReservedDays sets StatsReservedDays field to given value.

### HasStatsReservedDays

`func (o *RegionUpdateReqRegion) HasStatsReservedDays() bool`

HasStatsReservedDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


