# DpBlockAsyncReplicationPairCreateReqPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainName** | **string** | chain name | 
**GroupPairId** | **int64** | group pair id | 
**MasterClusterFsId** | **string** | master cluster fs id | 
**MasterGateway** | **string** | gateway ip | 
**MasterPairId** | **int64** | master pair id | 
**MasterPoolId** | **int64** | master pool id | 
**MasterPoolName** | **string** | master pool name | 
**MasterVolumeId** | **int64** | master volume id | 
**MasterVolumeName** | **string** | master volume name | 
**MasterVolumeSize** | **int64** | volume size | 
**MasterVolumeUuid** | Pointer to **string** | master volume uuid | [optional] 
**PolicyCron** | **string** | policy cron | 
**SlaveGateway** | **string** | gateway ip | 
**SlavePoolId** | **int64** | slave pool id | 
**SlavePoolName** | **string** | slave pool name | 
**SlaveVolumeName** | **string** | slave volume name | 

## Methods

### NewDpBlockAsyncReplicationPairCreateReqPair

`func NewDpBlockAsyncReplicationPairCreateReqPair(chainName string, groupPairId int64, masterClusterFsId string, masterGateway string, masterPairId int64, masterPoolId int64, masterPoolName string, masterVolumeId int64, masterVolumeName string, masterVolumeSize int64, policyCron string, slaveGateway string, slavePoolId int64, slavePoolName string, slaveVolumeName string, ) *DpBlockAsyncReplicationPairCreateReqPair`

NewDpBlockAsyncReplicationPairCreateReqPair instantiates a new DpBlockAsyncReplicationPairCreateReqPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpBlockAsyncReplicationPairCreateReqPairWithDefaults

`func NewDpBlockAsyncReplicationPairCreateReqPairWithDefaults() *DpBlockAsyncReplicationPairCreateReqPair`

NewDpBlockAsyncReplicationPairCreateReqPairWithDefaults instantiates a new DpBlockAsyncReplicationPairCreateReqPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainName

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetChainName() string`

GetChainName returns the ChainName field if non-nil, zero value otherwise.

### GetChainNameOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetChainNameOk() (*string, bool)`

GetChainNameOk returns a tuple with the ChainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainName

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetChainName(v string)`

SetChainName sets ChainName field to given value.


### GetGroupPairId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetGroupPairId() int64`

GetGroupPairId returns the GroupPairId field if non-nil, zero value otherwise.

### GetGroupPairIdOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetGroupPairIdOk() (*int64, bool)`

GetGroupPairIdOk returns a tuple with the GroupPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPairId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetGroupPairId(v int64)`

SetGroupPairId sets GroupPairId field to given value.


### GetMasterClusterFsId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterClusterFsId() string`

GetMasterClusterFsId returns the MasterClusterFsId field if non-nil, zero value otherwise.

### GetMasterClusterFsIdOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterClusterFsIdOk() (*string, bool)`

GetMasterClusterFsIdOk returns a tuple with the MasterClusterFsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterClusterFsId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetMasterClusterFsId(v string)`

SetMasterClusterFsId sets MasterClusterFsId field to given value.


### GetMasterGateway

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterGateway() string`

GetMasterGateway returns the MasterGateway field if non-nil, zero value otherwise.

### GetMasterGatewayOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterGatewayOk() (*string, bool)`

GetMasterGatewayOk returns a tuple with the MasterGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterGateway

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetMasterGateway(v string)`

SetMasterGateway sets MasterGateway field to given value.


### GetMasterPairId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterPairId() int64`

GetMasterPairId returns the MasterPairId field if non-nil, zero value otherwise.

### GetMasterPairIdOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterPairIdOk() (*int64, bool)`

GetMasterPairIdOk returns a tuple with the MasterPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPairId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetMasterPairId(v int64)`

SetMasterPairId sets MasterPairId field to given value.


### GetMasterPoolId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterPoolId() int64`

GetMasterPoolId returns the MasterPoolId field if non-nil, zero value otherwise.

### GetMasterPoolIdOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterPoolIdOk() (*int64, bool)`

GetMasterPoolIdOk returns a tuple with the MasterPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPoolId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetMasterPoolId(v int64)`

SetMasterPoolId sets MasterPoolId field to given value.


### GetMasterPoolName

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterPoolName() string`

GetMasterPoolName returns the MasterPoolName field if non-nil, zero value otherwise.

### GetMasterPoolNameOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterPoolNameOk() (*string, bool)`

GetMasterPoolNameOk returns a tuple with the MasterPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPoolName

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetMasterPoolName(v string)`

SetMasterPoolName sets MasterPoolName field to given value.


### GetMasterVolumeId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterVolumeId() int64`

GetMasterVolumeId returns the MasterVolumeId field if non-nil, zero value otherwise.

### GetMasterVolumeIdOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterVolumeIdOk() (*int64, bool)`

GetMasterVolumeIdOk returns a tuple with the MasterVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVolumeId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetMasterVolumeId(v int64)`

SetMasterVolumeId sets MasterVolumeId field to given value.


### GetMasterVolumeName

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterVolumeName() string`

GetMasterVolumeName returns the MasterVolumeName field if non-nil, zero value otherwise.

### GetMasterVolumeNameOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterVolumeNameOk() (*string, bool)`

GetMasterVolumeNameOk returns a tuple with the MasterVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVolumeName

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetMasterVolumeName(v string)`

SetMasterVolumeName sets MasterVolumeName field to given value.


### GetMasterVolumeSize

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterVolumeSize() int64`

GetMasterVolumeSize returns the MasterVolumeSize field if non-nil, zero value otherwise.

### GetMasterVolumeSizeOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterVolumeSizeOk() (*int64, bool)`

GetMasterVolumeSizeOk returns a tuple with the MasterVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVolumeSize

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetMasterVolumeSize(v int64)`

SetMasterVolumeSize sets MasterVolumeSize field to given value.


### GetMasterVolumeUuid

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterVolumeUuid() string`

GetMasterVolumeUuid returns the MasterVolumeUuid field if non-nil, zero value otherwise.

### GetMasterVolumeUuidOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetMasterVolumeUuidOk() (*string, bool)`

GetMasterVolumeUuidOk returns a tuple with the MasterVolumeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVolumeUuid

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetMasterVolumeUuid(v string)`

SetMasterVolumeUuid sets MasterVolumeUuid field to given value.

### HasMasterVolumeUuid

`func (o *DpBlockAsyncReplicationPairCreateReqPair) HasMasterVolumeUuid() bool`

HasMasterVolumeUuid returns a boolean if a field has been set.

### GetPolicyCron

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetPolicyCron() string`

GetPolicyCron returns the PolicyCron field if non-nil, zero value otherwise.

### GetPolicyCronOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetPolicyCronOk() (*string, bool)`

GetPolicyCronOk returns a tuple with the PolicyCron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCron

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetPolicyCron(v string)`

SetPolicyCron sets PolicyCron field to given value.


### GetSlaveGateway

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetSlaveGateway() string`

GetSlaveGateway returns the SlaveGateway field if non-nil, zero value otherwise.

### GetSlaveGatewayOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetSlaveGatewayOk() (*string, bool)`

GetSlaveGatewayOk returns a tuple with the SlaveGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveGateway

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetSlaveGateway(v string)`

SetSlaveGateway sets SlaveGateway field to given value.


### GetSlavePoolId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetSlavePoolId() int64`

GetSlavePoolId returns the SlavePoolId field if non-nil, zero value otherwise.

### GetSlavePoolIdOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetSlavePoolIdOk() (*int64, bool)`

GetSlavePoolIdOk returns a tuple with the SlavePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlavePoolId

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetSlavePoolId(v int64)`

SetSlavePoolId sets SlavePoolId field to given value.


### GetSlavePoolName

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetSlavePoolName() string`

GetSlavePoolName returns the SlavePoolName field if non-nil, zero value otherwise.

### GetSlavePoolNameOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetSlavePoolNameOk() (*string, bool)`

GetSlavePoolNameOk returns a tuple with the SlavePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlavePoolName

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetSlavePoolName(v string)`

SetSlavePoolName sets SlavePoolName field to given value.


### GetSlaveVolumeName

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetSlaveVolumeName() string`

GetSlaveVolumeName returns the SlaveVolumeName field if non-nil, zero value otherwise.

### GetSlaveVolumeNameOk

`func (o *DpBlockAsyncReplicationPairCreateReqPair) GetSlaveVolumeNameOk() (*string, bool)`

GetSlaveVolumeNameOk returns a tuple with the SlaveVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveVolumeName

`func (o *DpBlockAsyncReplicationPairCreateReqPair) SetSlaveVolumeName(v string)`

SetSlaveVolumeName sets SlaveVolumeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


