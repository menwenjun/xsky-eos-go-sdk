# VolumeMigrateReqVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestPoolId** | **int64** | migrate destination pool | 
**MigrationStripes** | Pointer to **int64** |  | [optional] 

## Methods

### NewVolumeMigrateReqVolume

`func NewVolumeMigrateReqVolume(destPoolId int64, ) *VolumeMigrateReqVolume`

NewVolumeMigrateReqVolume instantiates a new VolumeMigrateReqVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeMigrateReqVolumeWithDefaults

`func NewVolumeMigrateReqVolumeWithDefaults() *VolumeMigrateReqVolume`

NewVolumeMigrateReqVolumeWithDefaults instantiates a new VolumeMigrateReqVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestPoolId

`func (o *VolumeMigrateReqVolume) GetDestPoolId() int64`

GetDestPoolId returns the DestPoolId field if non-nil, zero value otherwise.

### GetDestPoolIdOk

`func (o *VolumeMigrateReqVolume) GetDestPoolIdOk() (*int64, bool)`

GetDestPoolIdOk returns a tuple with the DestPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPoolId

`func (o *VolumeMigrateReqVolume) SetDestPoolId(v int64)`

SetDestPoolId sets DestPoolId field to given value.


### GetMigrationStripes

`func (o *VolumeMigrateReqVolume) GetMigrationStripes() int64`

GetMigrationStripes returns the MigrationStripes field if non-nil, zero value otherwise.

### GetMigrationStripesOk

`func (o *VolumeMigrateReqVolume) GetMigrationStripesOk() (*int64, bool)`

GetMigrationStripesOk returns a tuple with the MigrationStripes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationStripes

`func (o *VolumeMigrateReqVolume) SetMigrationStripes(v int64)`

SetMigrationStripes sets MigrationStripes field to given value.

### HasMigrationStripes

`func (o *VolumeMigrateReqVolume) HasMigrationStripes() bool`

HasMigrationStripes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


