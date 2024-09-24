# OsdGroupStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **time.Time** |  | [optional] 
**MaxAvailableKbyte** | Pointer to **int64** |  | [optional] 
**MinAvailableKbyte** | Pointer to **int64** |  | [optional] 
**TotalKbyte** | Pointer to **int64** |  | [optional] 
**UsedKbyte** | Pointer to **int64** |  | [optional] 

## Methods

### NewOsdGroupStat

`func NewOsdGroupStat() *OsdGroupStat`

NewOsdGroupStat instantiates a new OsdGroupStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsdGroupStatWithDefaults

`func NewOsdGroupStatWithDefaults() *OsdGroupStat`

NewOsdGroupStatWithDefaults instantiates a new OsdGroupStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *OsdGroupStat) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *OsdGroupStat) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *OsdGroupStat) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *OsdGroupStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetMaxAvailableKbyte

`func (o *OsdGroupStat) GetMaxAvailableKbyte() int64`

GetMaxAvailableKbyte returns the MaxAvailableKbyte field if non-nil, zero value otherwise.

### GetMaxAvailableKbyteOk

`func (o *OsdGroupStat) GetMaxAvailableKbyteOk() (*int64, bool)`

GetMaxAvailableKbyteOk returns a tuple with the MaxAvailableKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvailableKbyte

`func (o *OsdGroupStat) SetMaxAvailableKbyte(v int64)`

SetMaxAvailableKbyte sets MaxAvailableKbyte field to given value.

### HasMaxAvailableKbyte

`func (o *OsdGroupStat) HasMaxAvailableKbyte() bool`

HasMaxAvailableKbyte returns a boolean if a field has been set.

### GetMinAvailableKbyte

`func (o *OsdGroupStat) GetMinAvailableKbyte() int64`

GetMinAvailableKbyte returns the MinAvailableKbyte field if non-nil, zero value otherwise.

### GetMinAvailableKbyteOk

`func (o *OsdGroupStat) GetMinAvailableKbyteOk() (*int64, bool)`

GetMinAvailableKbyteOk returns a tuple with the MinAvailableKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAvailableKbyte

`func (o *OsdGroupStat) SetMinAvailableKbyte(v int64)`

SetMinAvailableKbyte sets MinAvailableKbyte field to given value.

### HasMinAvailableKbyte

`func (o *OsdGroupStat) HasMinAvailableKbyte() bool`

HasMinAvailableKbyte returns a boolean if a field has been set.

### GetTotalKbyte

`func (o *OsdGroupStat) GetTotalKbyte() int64`

GetTotalKbyte returns the TotalKbyte field if non-nil, zero value otherwise.

### GetTotalKbyteOk

`func (o *OsdGroupStat) GetTotalKbyteOk() (*int64, bool)`

GetTotalKbyteOk returns a tuple with the TotalKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKbyte

`func (o *OsdGroupStat) SetTotalKbyte(v int64)`

SetTotalKbyte sets TotalKbyte field to given value.

### HasTotalKbyte

`func (o *OsdGroupStat) HasTotalKbyte() bool`

HasTotalKbyte returns a boolean if a field has been set.

### GetUsedKbyte

`func (o *OsdGroupStat) GetUsedKbyte() int64`

GetUsedKbyte returns the UsedKbyte field if non-nil, zero value otherwise.

### GetUsedKbyteOk

`func (o *OsdGroupStat) GetUsedKbyteOk() (*int64, bool)`

GetUsedKbyteOk returns a tuple with the UsedKbyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKbyte

`func (o *OsdGroupStat) SetUsedKbyte(v int64)`

SetUsedKbyte sets UsedKbyte field to given value.

### HasUsedKbyte

`func (o *OsdGroupStat) HasUsedKbyte() bool`

HasUsedKbyte returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


