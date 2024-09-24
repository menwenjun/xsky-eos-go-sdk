# PoolNestview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | Pointer to **string** |  | [optional] 
**EncryptEnabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumaApplyPolicy** | Pointer to **string** |  | [optional] 
**NumaBindPolicy** | Pointer to **string** |  | [optional] 
**NumaEnabled** | Pointer to **bool** |  | [optional] 
**PoolType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Stretched** | Pointer to **bool** |  | [optional] 

## Methods

### NewPoolNestview

`func NewPoolNestview() *PoolNestview`

NewPoolNestview instantiates a new PoolNestview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolNestviewWithDefaults

`func NewPoolNestviewWithDefaults() *PoolNestview`

NewPoolNestviewWithDefaults instantiates a new PoolNestview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionStatus

`func (o *PoolNestview) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *PoolNestview) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *PoolNestview) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.

### HasActionStatus

`func (o *PoolNestview) HasActionStatus() bool`

HasActionStatus returns a boolean if a field has been set.

### GetEncryptEnabled

`func (o *PoolNestview) GetEncryptEnabled() bool`

GetEncryptEnabled returns the EncryptEnabled field if non-nil, zero value otherwise.

### GetEncryptEnabledOk

`func (o *PoolNestview) GetEncryptEnabledOk() (*bool, bool)`

GetEncryptEnabledOk returns a tuple with the EncryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptEnabled

`func (o *PoolNestview) SetEncryptEnabled(v bool)`

SetEncryptEnabled sets EncryptEnabled field to given value.

### HasEncryptEnabled

`func (o *PoolNestview) HasEncryptEnabled() bool`

HasEncryptEnabled returns a boolean if a field has been set.

### GetId

`func (o *PoolNestview) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolNestview) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolNestview) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PoolNestview) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PoolNestview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolNestview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolNestview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolNestview) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumaApplyPolicy

`func (o *PoolNestview) GetNumaApplyPolicy() string`

GetNumaApplyPolicy returns the NumaApplyPolicy field if non-nil, zero value otherwise.

### GetNumaApplyPolicyOk

`func (o *PoolNestview) GetNumaApplyPolicyOk() (*string, bool)`

GetNumaApplyPolicyOk returns a tuple with the NumaApplyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaApplyPolicy

`func (o *PoolNestview) SetNumaApplyPolicy(v string)`

SetNumaApplyPolicy sets NumaApplyPolicy field to given value.

### HasNumaApplyPolicy

`func (o *PoolNestview) HasNumaApplyPolicy() bool`

HasNumaApplyPolicy returns a boolean if a field has been set.

### GetNumaBindPolicy

`func (o *PoolNestview) GetNumaBindPolicy() string`

GetNumaBindPolicy returns the NumaBindPolicy field if non-nil, zero value otherwise.

### GetNumaBindPolicyOk

`func (o *PoolNestview) GetNumaBindPolicyOk() (*string, bool)`

GetNumaBindPolicyOk returns a tuple with the NumaBindPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaBindPolicy

`func (o *PoolNestview) SetNumaBindPolicy(v string)`

SetNumaBindPolicy sets NumaBindPolicy field to given value.

### HasNumaBindPolicy

`func (o *PoolNestview) HasNumaBindPolicy() bool`

HasNumaBindPolicy returns a boolean if a field has been set.

### GetNumaEnabled

`func (o *PoolNestview) GetNumaEnabled() bool`

GetNumaEnabled returns the NumaEnabled field if non-nil, zero value otherwise.

### GetNumaEnabledOk

`func (o *PoolNestview) GetNumaEnabledOk() (*bool, bool)`

GetNumaEnabledOk returns a tuple with the NumaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaEnabled

`func (o *PoolNestview) SetNumaEnabled(v bool)`

SetNumaEnabled sets NumaEnabled field to given value.

### HasNumaEnabled

`func (o *PoolNestview) HasNumaEnabled() bool`

HasNumaEnabled returns a boolean if a field has been set.

### GetPoolType

`func (o *PoolNestview) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *PoolNestview) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *PoolNestview) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *PoolNestview) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetStatus

`func (o *PoolNestview) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoolNestview) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoolNestview) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PoolNestview) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStretched

`func (o *PoolNestview) GetStretched() bool`

GetStretched returns the Stretched field if non-nil, zero value otherwise.

### GetStretchedOk

`func (o *PoolNestview) GetStretchedOk() (*bool, bool)`

GetStretchedOk returns a tuple with the Stretched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStretched

`func (o *PoolNestview) SetStretched(v bool)`

SetStretched sets Stretched field to given value.

### HasStretched

`func (o *PoolNestview) HasStretched() bool`

HasStretched returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


