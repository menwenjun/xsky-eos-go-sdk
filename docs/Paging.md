# Paging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**DurationBegin** | Pointer to **time.Time** |  | [optional] 
**DurationEnd** | Pointer to **time.Time** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 

## Methods

### NewPaging

`func NewPaging() *Paging`

NewPaging instantiates a new Paging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingWithDefaults

`func NewPagingWithDefaults() *Paging`

NewPagingWithDefaults instantiates a new Paging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Paging) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Paging) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Paging) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *Paging) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDuration

`func (o *Paging) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Paging) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Paging) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Paging) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationBegin

`func (o *Paging) GetDurationBegin() time.Time`

GetDurationBegin returns the DurationBegin field if non-nil, zero value otherwise.

### GetDurationBeginOk

`func (o *Paging) GetDurationBeginOk() (*time.Time, bool)`

GetDurationBeginOk returns a tuple with the DurationBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationBegin

`func (o *Paging) SetDurationBegin(v time.Time)`

SetDurationBegin sets DurationBegin field to given value.

### HasDurationBegin

`func (o *Paging) HasDurationBegin() bool`

HasDurationBegin returns a boolean if a field has been set.

### GetDurationEnd

`func (o *Paging) GetDurationEnd() time.Time`

GetDurationEnd returns the DurationEnd field if non-nil, zero value otherwise.

### GetDurationEndOk

`func (o *Paging) GetDurationEndOk() (*time.Time, bool)`

GetDurationEndOk returns a tuple with the DurationEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationEnd

`func (o *Paging) SetDurationEnd(v time.Time)`

SetDurationEnd sets DurationEnd field to given value.

### HasDurationEnd

`func (o *Paging) HasDurationEnd() bool`

HasDurationEnd returns a boolean if a field has been set.

### GetLimit

`func (o *Paging) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Paging) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Paging) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Paging) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *Paging) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Paging) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Paging) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *Paging) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetOffset

`func (o *Paging) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Paging) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Paging) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Paging) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPeriod

`func (o *Paging) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Paging) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Paging) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Paging) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetTotalCount

`func (o *Paging) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Paging) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Paging) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Paging) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTruncated

`func (o *Paging) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *Paging) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *Paging) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *Paging) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


