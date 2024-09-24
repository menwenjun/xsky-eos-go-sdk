# DiskSamplesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskSamples** | [**[]DiskStat**](DiskStat.md) | disk samples | 

## Methods

### NewDiskSamplesResp

`func NewDiskSamplesResp(diskSamples []DiskStat, ) *DiskSamplesResp`

NewDiskSamplesResp instantiates a new DiskSamplesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskSamplesRespWithDefaults

`func NewDiskSamplesRespWithDefaults() *DiskSamplesResp`

NewDiskSamplesRespWithDefaults instantiates a new DiskSamplesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskSamples

`func (o *DiskSamplesResp) GetDiskSamples() []DiskStat`

GetDiskSamples returns the DiskSamples field if non-nil, zero value otherwise.

### GetDiskSamplesOk

`func (o *DiskSamplesResp) GetDiskSamplesOk() (*[]DiskStat, bool)`

GetDiskSamplesOk returns a tuple with the DiskSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSamples

`func (o *DiskSamplesResp) SetDiskSamples(v []DiskStat)`

SetDiskSamples sets DiskSamples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


