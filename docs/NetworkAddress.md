# NetworkAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Create** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to [**HostNestview**](HostNestview.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Mask** | Pointer to **int64** |  | [optional] 
**NetworkInterface** | Pointer to [**NetworkInterfaceNestview**](NetworkInterfaceNestview.md) |  | [optional] 
**OspCluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**SlaveSwitchTime** | Pointer to **time.Time** |  | [optional] 
**SlaveSwitched** | Pointer to **bool** |  | [optional] 
**Update** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNetworkAddress

`func NewNetworkAddress() *NetworkAddress`

NewNetworkAddress instantiates a new NetworkAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkAddressWithDefaults

`func NewNetworkAddressWithDefaults() *NetworkAddress`

NewNetworkAddressWithDefaults instantiates a new NetworkAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *NetworkAddress) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *NetworkAddress) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *NetworkAddress) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *NetworkAddress) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCluster

`func (o *NetworkAddress) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *NetworkAddress) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *NetworkAddress) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *NetworkAddress) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreate

`func (o *NetworkAddress) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *NetworkAddress) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *NetworkAddress) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *NetworkAddress) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetHost

`func (o *NetworkAddress) GetHost() HostNestview`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworkAddress) GetHostOk() (*HostNestview, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworkAddress) SetHost(v HostNestview)`

SetHost sets Host field to given value.

### HasHost

`func (o *NetworkAddress) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *NetworkAddress) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkAddress) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkAddress) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *NetworkAddress) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NetworkAddress) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NetworkAddress) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NetworkAddress) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMask

`func (o *NetworkAddress) GetMask() int64`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *NetworkAddress) GetMaskOk() (*int64, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *NetworkAddress) SetMask(v int64)`

SetMask sets Mask field to given value.

### HasMask

`func (o *NetworkAddress) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetNetworkInterface

`func (o *NetworkAddress) GetNetworkInterface() NetworkInterfaceNestview`

GetNetworkInterface returns the NetworkInterface field if non-nil, zero value otherwise.

### GetNetworkInterfaceOk

`func (o *NetworkAddress) GetNetworkInterfaceOk() (*NetworkInterfaceNestview, bool)`

GetNetworkInterfaceOk returns a tuple with the NetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterface

`func (o *NetworkAddress) SetNetworkInterface(v NetworkInterfaceNestview)`

SetNetworkInterface sets NetworkInterface field to given value.

### HasNetworkInterface

`func (o *NetworkAddress) HasNetworkInterface() bool`

HasNetworkInterface returns a boolean if a field has been set.

### GetOspCluster

`func (o *NetworkAddress) GetOspCluster() Cluster`

GetOspCluster returns the OspCluster field if non-nil, zero value otherwise.

### GetOspClusterOk

`func (o *NetworkAddress) GetOspClusterOk() (*Cluster, bool)`

GetOspClusterOk returns a tuple with the OspCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspCluster

`func (o *NetworkAddress) SetOspCluster(v Cluster)`

SetOspCluster sets OspCluster field to given value.

### HasOspCluster

`func (o *NetworkAddress) HasOspCluster() bool`

HasOspCluster returns a boolean if a field has been set.

### GetRoles

`func (o *NetworkAddress) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *NetworkAddress) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *NetworkAddress) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *NetworkAddress) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSlaveSwitchTime

`func (o *NetworkAddress) GetSlaveSwitchTime() time.Time`

GetSlaveSwitchTime returns the SlaveSwitchTime field if non-nil, zero value otherwise.

### GetSlaveSwitchTimeOk

`func (o *NetworkAddress) GetSlaveSwitchTimeOk() (*time.Time, bool)`

GetSlaveSwitchTimeOk returns a tuple with the SlaveSwitchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveSwitchTime

`func (o *NetworkAddress) SetSlaveSwitchTime(v time.Time)`

SetSlaveSwitchTime sets SlaveSwitchTime field to given value.

### HasSlaveSwitchTime

`func (o *NetworkAddress) HasSlaveSwitchTime() bool`

HasSlaveSwitchTime returns a boolean if a field has been set.

### GetSlaveSwitched

`func (o *NetworkAddress) GetSlaveSwitched() bool`

GetSlaveSwitched returns the SlaveSwitched field if non-nil, zero value otherwise.

### GetSlaveSwitchedOk

`func (o *NetworkAddress) GetSlaveSwitchedOk() (*bool, bool)`

GetSlaveSwitchedOk returns a tuple with the SlaveSwitched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveSwitched

`func (o *NetworkAddress) SetSlaveSwitched(v bool)`

SetSlaveSwitched sets SlaveSwitched field to given value.

### HasSlaveSwitched

`func (o *NetworkAddress) HasSlaveSwitched() bool`

HasSlaveSwitched returns a boolean if a field has been set.

### GetUpdate

`func (o *NetworkAddress) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *NetworkAddress) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *NetworkAddress) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *NetworkAddress) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


