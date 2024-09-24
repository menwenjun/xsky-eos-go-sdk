# DfsTierPoolPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**DfsStorageClass** | Pointer to [**DfsTierNestview**](DfsTierNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Pool** | Pointer to [**PoolNestview**](PoolNestview.md) |  | [optional] 
**Threshold** | Pointer to **int64** |  | [optional] 

## Methods

### NewDfsTierPoolPolicy

`func NewDfsTierPoolPolicy() *DfsTierPoolPolicy`

NewDfsTierPoolPolicy instantiates a new DfsTierPoolPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDfsTierPoolPolicyWithDefaults

`func NewDfsTierPoolPolicyWithDefaults() *DfsTierPoolPolicy`

NewDfsTierPoolPolicyWithDefaults instantiates a new DfsTierPoolPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DfsTierPoolPolicy) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DfsTierPoolPolicy) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DfsTierPoolPolicy) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DfsTierPoolPolicy) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDfsStorageClass

`func (o *DfsTierPoolPolicy) GetDfsStorageClass() DfsTierNestview`

GetDfsStorageClass returns the DfsStorageClass field if non-nil, zero value otherwise.

### GetDfsStorageClassOk

`func (o *DfsTierPoolPolicy) GetDfsStorageClassOk() (*DfsTierNestview, bool)`

GetDfsStorageClassOk returns a tuple with the DfsStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsStorageClass

`func (o *DfsTierPoolPolicy) SetDfsStorageClass(v DfsTierNestview)`

SetDfsStorageClass sets DfsStorageClass field to given value.

### HasDfsStorageClass

`func (o *DfsTierPoolPolicy) HasDfsStorageClass() bool`

HasDfsStorageClass returns a boolean if a field has been set.

### GetId

`func (o *DfsTierPoolPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DfsTierPoolPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DfsTierPoolPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DfsTierPoolPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPool

`func (o *DfsTierPoolPolicy) GetPool() PoolNestview`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *DfsTierPoolPolicy) GetPoolOk() (*PoolNestview, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *DfsTierPoolPolicy) SetPool(v PoolNestview)`

SetPool sets Pool field to given value.

### HasPool

`func (o *DfsTierPoolPolicy) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetThreshold

`func (o *DfsTierPoolPolicy) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *DfsTierPoolPolicy) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *DfsTierPoolPolicy) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *DfsTierPoolPolicy) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


