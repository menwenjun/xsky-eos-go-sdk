# DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainName** | **string** | chain name | 
**MasterClusterFsId** | **string** | master cluster fs id | 
**MasterGateway** | **string** | gateway ip | 
**MasterPairId** | **int64** | master pair id | 
**MasterVolumeGroupId** | **int64** | master volume group id | 
**MasterVolumeGroupName** | **string** | master volume group name | 
**PolicyCron** | **string** | policy cron | 
**SlaveGateway** | **string** | gateway ip | 

## Methods

### NewDpVolumeGroupSnapshotReplicationPairCreateReqGroupPair

`func NewDpVolumeGroupSnapshotReplicationPairCreateReqGroupPair(chainName string, masterClusterFsId string, masterGateway string, masterPairId int64, masterVolumeGroupId int64, masterVolumeGroupName string, policyCron string, slaveGateway string, ) *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair`

NewDpVolumeGroupSnapshotReplicationPairCreateReqGroupPair instantiates a new DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpVolumeGroupSnapshotReplicationPairCreateReqGroupPairWithDefaults

`func NewDpVolumeGroupSnapshotReplicationPairCreateReqGroupPairWithDefaults() *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair`

NewDpVolumeGroupSnapshotReplicationPairCreateReqGroupPairWithDefaults instantiates a new DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainName

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetChainName() string`

GetChainName returns the ChainName field if non-nil, zero value otherwise.

### GetChainNameOk

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetChainNameOk() (*string, bool)`

GetChainNameOk returns a tuple with the ChainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainName

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) SetChainName(v string)`

SetChainName sets ChainName field to given value.


### GetMasterClusterFsId

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetMasterClusterFsId() string`

GetMasterClusterFsId returns the MasterClusterFsId field if non-nil, zero value otherwise.

### GetMasterClusterFsIdOk

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetMasterClusterFsIdOk() (*string, bool)`

GetMasterClusterFsIdOk returns a tuple with the MasterClusterFsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterClusterFsId

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) SetMasterClusterFsId(v string)`

SetMasterClusterFsId sets MasterClusterFsId field to given value.


### GetMasterGateway

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetMasterGateway() string`

GetMasterGateway returns the MasterGateway field if non-nil, zero value otherwise.

### GetMasterGatewayOk

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetMasterGatewayOk() (*string, bool)`

GetMasterGatewayOk returns a tuple with the MasterGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterGateway

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) SetMasterGateway(v string)`

SetMasterGateway sets MasterGateway field to given value.


### GetMasterPairId

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetMasterPairId() int64`

GetMasterPairId returns the MasterPairId field if non-nil, zero value otherwise.

### GetMasterPairIdOk

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetMasterPairIdOk() (*int64, bool)`

GetMasterPairIdOk returns a tuple with the MasterPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPairId

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) SetMasterPairId(v int64)`

SetMasterPairId sets MasterPairId field to given value.


### GetMasterVolumeGroupId

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetMasterVolumeGroupId() int64`

GetMasterVolumeGroupId returns the MasterVolumeGroupId field if non-nil, zero value otherwise.

### GetMasterVolumeGroupIdOk

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetMasterVolumeGroupIdOk() (*int64, bool)`

GetMasterVolumeGroupIdOk returns a tuple with the MasterVolumeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVolumeGroupId

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) SetMasterVolumeGroupId(v int64)`

SetMasterVolumeGroupId sets MasterVolumeGroupId field to given value.


### GetMasterVolumeGroupName

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetMasterVolumeGroupName() string`

GetMasterVolumeGroupName returns the MasterVolumeGroupName field if non-nil, zero value otherwise.

### GetMasterVolumeGroupNameOk

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetMasterVolumeGroupNameOk() (*string, bool)`

GetMasterVolumeGroupNameOk returns a tuple with the MasterVolumeGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVolumeGroupName

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) SetMasterVolumeGroupName(v string)`

SetMasterVolumeGroupName sets MasterVolumeGroupName field to given value.


### GetPolicyCron

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetPolicyCron() string`

GetPolicyCron returns the PolicyCron field if non-nil, zero value otherwise.

### GetPolicyCronOk

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetPolicyCronOk() (*string, bool)`

GetPolicyCronOk returns a tuple with the PolicyCron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCron

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) SetPolicyCron(v string)`

SetPolicyCron sets PolicyCron field to given value.


### GetSlaveGateway

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetSlaveGateway() string`

GetSlaveGateway returns the SlaveGateway field if non-nil, zero value otherwise.

### GetSlaveGatewayOk

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) GetSlaveGatewayOk() (*string, bool)`

GetSlaveGatewayOk returns a tuple with the SlaveGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveGateway

`func (o *DpVolumeGroupSnapshotReplicationPairCreateReqGroupPair) SetSlaveGateway(v string)`

SetSlaveGateway sets SlaveGateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


