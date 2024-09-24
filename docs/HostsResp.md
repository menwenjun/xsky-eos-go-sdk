# HostsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | [**[]HostRecord**](HostRecord.md) | hosts | 

## Methods

### NewHostsResp

`func NewHostsResp(hosts []HostRecord, ) *HostsResp`

NewHostsResp instantiates a new HostsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostsRespWithDefaults

`func NewHostsRespWithDefaults() *HostsResp`

NewHostsRespWithDefaults instantiates a new HostsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *HostsResp) GetHosts() []HostRecord`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *HostsResp) GetHostsOk() (*[]HostRecord, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *HostsResp) SetHosts(v []HostRecord)`

SetHosts sets Hosts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


