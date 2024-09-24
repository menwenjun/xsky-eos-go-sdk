# DpBlockSnapshotRecoveryJobCreateReqJobVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**PoolId** | **int64** |  | 

## Methods

### NewDpBlockSnapshotRecoveryJobCreateReqJobVolume

`func NewDpBlockSnapshotRecoveryJobCreateReqJobVolume(name string, poolId int64, ) *DpBlockSnapshotRecoveryJobCreateReqJobVolume`

NewDpBlockSnapshotRecoveryJobCreateReqJobVolume instantiates a new DpBlockSnapshotRecoveryJobCreateReqJobVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockSnapshotRecoveryJobCreateReqJobVolumeWithDefaults

`func NewDpBlockSnapshotRecoveryJobCreateReqJobVolumeWithDefaults() *DpBlockSnapshotRecoveryJobCreateReqJobVolume`

NewDpBlockSnapshotRecoveryJobCreateReqJobVolumeWithDefaults instantiates a new DpBlockSnapshotRecoveryJobCreateReqJobVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobVolume) GetFormat() int64`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobVolume) GetFormatOk() (*int64, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobVolume) SetFormat(v int64)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobVolume) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobVolume) SetName(v string)`

SetName sets Name field to given value.


### GetPoolId

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobVolume) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobVolume) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *DpBlockSnapshotRecoveryJobCreateReqJobVolume) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


