# AlertInfoBatchDeleteResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateAfter** | Pointer to **time.Time** |  | [optional] 
**CreateBefore** | Pointer to **time.Time** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**UnresolvedCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewAlertInfoBatchDeleteResp

`func NewAlertInfoBatchDeleteResp() *AlertInfoBatchDeleteResp`

NewAlertInfoBatchDeleteResp instantiates a new AlertInfoBatchDeleteResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertInfoBatchDeleteRespWithDefaults

`func NewAlertInfoBatchDeleteRespWithDefaults() *AlertInfoBatchDeleteResp`

NewAlertInfoBatchDeleteRespWithDefaults instantiates a new AlertInfoBatchDeleteResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateAfter

`func (o *AlertInfoBatchDeleteResp) GetCreateAfter() time.Time`

GetCreateAfter returns the CreateAfter field if non-nil, zero value otherwise.

### GetCreateAfterOk

`func (o *AlertInfoBatchDeleteResp) GetCreateAfterOk() (*time.Time, bool)`

GetCreateAfterOk returns a tuple with the CreateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAfter

`func (o *AlertInfoBatchDeleteResp) SetCreateAfter(v time.Time)`

SetCreateAfter sets CreateAfter field to given value.

### HasCreateAfter

`func (o *AlertInfoBatchDeleteResp) HasCreateAfter() bool`

HasCreateAfter returns a boolean if a field has been set.

### GetCreateBefore

`func (o *AlertInfoBatchDeleteResp) GetCreateBefore() time.Time`

GetCreateBefore returns the CreateBefore field if non-nil, zero value otherwise.

### GetCreateBeforeOk

`func (o *AlertInfoBatchDeleteResp) GetCreateBeforeOk() (*time.Time, bool)`

GetCreateBeforeOk returns a tuple with the CreateBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBefore

`func (o *AlertInfoBatchDeleteResp) SetCreateBefore(v time.Time)`

SetCreateBefore sets CreateBefore field to given value.

### HasCreateBefore

`func (o *AlertInfoBatchDeleteResp) HasCreateBefore() bool`

HasCreateBefore returns a boolean if a field has been set.

### GetTotal

`func (o *AlertInfoBatchDeleteResp) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AlertInfoBatchDeleteResp) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AlertInfoBatchDeleteResp) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AlertInfoBatchDeleteResp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnresolvedCount

`func (o *AlertInfoBatchDeleteResp) GetUnresolvedCount() int64`

GetUnresolvedCount returns the UnresolvedCount field if non-nil, zero value otherwise.

### GetUnresolvedCountOk

`func (o *AlertInfoBatchDeleteResp) GetUnresolvedCountOk() (*int64, bool)`

GetUnresolvedCountOk returns a tuple with the UnresolvedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedCount

`func (o *AlertInfoBatchDeleteResp) SetUnresolvedCount(v int64)`

SetUnresolvedCount sets UnresolvedCount field to given value.

### HasUnresolvedCount

`func (o *AlertInfoBatchDeleteResp) HasUnresolvedCount() bool`

HasUnresolvedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


