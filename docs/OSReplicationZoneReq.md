# OSReplicationZoneReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsRemotePolicyUuid** | **string** |  | 
**Readonly** | Pointer to **bool** |  | [optional] 

## Methods

### NewOSReplicationZoneReq

`func NewOSReplicationZoneReq(osRemotePolicyUuid string, ) *OSReplicationZoneReq`

NewOSReplicationZoneReq instantiates a new OSReplicationZoneReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSReplicationZoneReqWithDefaults

`func NewOSReplicationZoneReqWithDefaults() *OSReplicationZoneReq`

NewOSReplicationZoneReqWithDefaults instantiates a new OSReplicationZoneReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsRemotePolicyUuid

`func (o *OSReplicationZoneReq) GetOsRemotePolicyUuid() string`

GetOsRemotePolicyUuid returns the OsRemotePolicyUuid field if non-nil, zero value otherwise.

### GetOsRemotePolicyUuidOk

`func (o *OSReplicationZoneReq) GetOsRemotePolicyUuidOk() (*string, bool)`

GetOsRemotePolicyUuidOk returns a tuple with the OsRemotePolicyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRemotePolicyUuid

`func (o *OSReplicationZoneReq) SetOsRemotePolicyUuid(v string)`

SetOsRemotePolicyUuid sets OsRemotePolicyUuid field to given value.


### GetReadonly

`func (o *OSReplicationZoneReq) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *OSReplicationZoneReq) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *OSReplicationZoneReq) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *OSReplicationZoneReq) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


