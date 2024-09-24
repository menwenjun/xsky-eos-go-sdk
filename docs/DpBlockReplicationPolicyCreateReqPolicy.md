# DpBlockReplicationPolicyCreateReqPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DpSiteId** | **int64** | protection site id | 
**Name** | **string** | name | 
**TimeoutSeconds** | **int64** | replication timeout seconds | 

## Methods

### NewDpBlockReplicationPolicyCreateReqPolicy

`func NewDpBlockReplicationPolicyCreateReqPolicy(dpSiteId int64, name string, timeoutSeconds int64, ) *DpBlockReplicationPolicyCreateReqPolicy`

NewDpBlockReplicationPolicyCreateReqPolicy instantiates a new DpBlockReplicationPolicyCreateReqPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockReplicationPolicyCreateReqPolicyWithDefaults

`func NewDpBlockReplicationPolicyCreateReqPolicyWithDefaults() *DpBlockReplicationPolicyCreateReqPolicy`

NewDpBlockReplicationPolicyCreateReqPolicyWithDefaults instantiates a new DpBlockReplicationPolicyCreateReqPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpSiteId

`func (o *DpBlockReplicationPolicyCreateReqPolicy) GetDpSiteId() int64`

GetDpSiteId returns the DpSiteId field if non-nil, zero value otherwise.

### GetDpSiteIdOk

`func (o *DpBlockReplicationPolicyCreateReqPolicy) GetDpSiteIdOk() (*int64, bool)`

GetDpSiteIdOk returns a tuple with the DpSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpSiteId

`func (o *DpBlockReplicationPolicyCreateReqPolicy) SetDpSiteId(v int64)`

SetDpSiteId sets DpSiteId field to given value.


### GetName

`func (o *DpBlockReplicationPolicyCreateReqPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DpBlockReplicationPolicyCreateReqPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DpBlockReplicationPolicyCreateReqPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetTimeoutSeconds

`func (o *DpBlockReplicationPolicyCreateReqPolicy) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *DpBlockReplicationPolicyCreateReqPolicy) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *DpBlockReplicationPolicyCreateReqPolicy) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


