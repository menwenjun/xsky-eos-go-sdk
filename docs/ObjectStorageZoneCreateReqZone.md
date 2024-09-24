# ObjectStorageZoneCreateReqZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | os zone alias | [optional] 
**Master** | **bool** | is a master zone | 
**RemoteClusterId** | **int64** | remote cluster id | 

## Methods

### NewObjectStorageZoneCreateReqZone

`func NewObjectStorageZoneCreateReqZone(master bool, remoteClusterId int64, ) *ObjectStorageZoneCreateReqZone`

NewObjectStorageZoneCreateReqZone instantiates a new ObjectStorageZoneCreateReqZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageZoneCreateReqZoneWithDefaults

`func NewObjectStorageZoneCreateReqZoneWithDefaults() *ObjectStorageZoneCreateReqZone`

NewObjectStorageZoneCreateReqZoneWithDefaults instantiates a new ObjectStorageZoneCreateReqZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ObjectStorageZoneCreateReqZone) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ObjectStorageZoneCreateReqZone) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ObjectStorageZoneCreateReqZone) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ObjectStorageZoneCreateReqZone) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetMaster

`func (o *ObjectStorageZoneCreateReqZone) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ObjectStorageZoneCreateReqZone) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ObjectStorageZoneCreateReqZone) SetMaster(v bool)`

SetMaster sets Master field to given value.


### GetRemoteClusterId

`func (o *ObjectStorageZoneCreateReqZone) GetRemoteClusterId() int64`

GetRemoteClusterId returns the RemoteClusterId field if non-nil, zero value otherwise.

### GetRemoteClusterIdOk

`func (o *ObjectStorageZoneCreateReqZone) GetRemoteClusterIdOk() (*int64, bool)`

GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterId

`func (o *ObjectStorageZoneCreateReqZone) SetRemoteClusterId(v int64)`

SetRemoteClusterId sets RemoteClusterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


